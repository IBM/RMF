/**
* (C) Copyright IBM Corp. 2023, 2025.
* (C) Copyright Rocket Software, Inc. 2023-2025.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package dds

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"golang.org/x/sync/singleflight"
)

const UpdateInterval = 15 * time.Minute
const DefaultTimeOffset = 0
const DefaultMinTime = 100

const IndexPath = "/gpm/index"
const RootPath = "/gpm/root"
const ContainedPath = "/gpm/contained"
const PerformPath = "/gpm/perform"
const XslHeadersPath = "/gpm/include/reptrans.xsl"
const FullReportPath = "/gpm/rmfm3"

var MayHaveExt = map[string]bool{
	IndexPath:      true,
	RootPath:       true,
	ContainedPath:  true,
	PerformPath:    true,
	FullReportPath: true,
}
var ErrParse = errors.New("unable to parse DDS response")
var ErrUnauthorized = errors.New("not authorized to access DDS")

type Client struct {
	baseUrl    string
	username   string
	password   string
	httpClient *http.Client
	headerMap  *HeaderMap
	timeData   *TimeData
	useXmlExt  atomic.Bool

	stopChan  chan struct{}
	closeOnce sync.Once
	waitGroup sync.WaitGroup
	rwMutex   sync.RWMutex
	single    singleflight.Group
}

func NewClient(baseUrl string, username string, password string, timeout int, tlsSkipVerify bool, disableCompression bool) *Client {
	client := &Client{
		baseUrl:  strings.TrimRight(baseUrl, "/"),
		username: username,
		password: password,
		httpClient: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				DisableCompression: disableCompression,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: tlsSkipVerify, // #nosec G402
				},
			},
		},
		stopChan: make(chan struct{}),
	}
	client.waitGroup.Add(1)
	go client.sync()
	return client
}

func (c *Client) sync() {
	logger := log.Logger.With("func", "sync")
	ticker := time.NewTicker(UpdateInterval)
	defer c.waitGroup.Done()
	defer ticker.Stop()
	defer func() {
		if r := recover(); r != nil {
			logger.Debug("DDS background sync stopped", "error", r)
		}
	}()
	logger.Debug("DDS background sync started")
	c.updateCache()
	for {
		select {
		case <-ticker.C:
			c.updateCache()
		case <-c.stopChan:
			logger.Debug("DDS background sync stopped")
			return
		}
	}
}

func (c *Client) updateCache() {
	c.updateTimeData()
	c.updateHeaders()
}

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		close(c.stopChan)
		c.waitGroup.Wait()
		c.httpClient.CloseIdleConnections()
	})
}

func (c *Client) GetByRequest(r *Request) (*Response, error) {
	path, params, err := r.pathWithParams(c.GetCachedTimeOffset())
	if err != nil {
		return nil, err
	}
	return c.Get(path, params...)
}

func (c *Client) Get(path string, params ...string) (*Response, error) {
	var response Response
	data, err := c.GetRaw(path, params...)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrParse, err)
	}
	return &response, nil
}

func (c *Client) GetRaw(path string, params ...string) ([]byte, error) {
	logger := log.Logger.With("func", "GetRaw")

	_, mayHaveExt := MayHaveExt[path]
	useXmlExt := c.useXmlExt.Load()
	ofs := 0
	if useXmlExt {
		ofs = 1
	}

	path = strings.TrimLeft(path, "/")
	values := url.Values{}
	for i := 0; i < len(params)-1; i += 2 {
		if i+1 >= len(params) {
			return nil, errors.New("params must be key-value pairs of strings")
		}
		values.Add(params[i], params[i+1])
	}

	// Newer versions of DDS have JSON endpoints without the XML extension.
	// Also, it may run in a mode not fully compliant with HTTP emulating DDS v1 behavior.
	// We need to support either variation.
	for i := range 2 {
		if mayHaveExt && (i+ofs)%2 == 1 {
			path += ".xml"
		}
		fullURL := fmt.Sprintf("%s/%s?%s", c.baseUrl, path, values.Encode())
		// nolint:noctx
		req, err := http.NewRequest(http.MethodGet, fullURL, http.NoBody)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Accept", "application/json")
		if strings.TrimSpace(c.username) != "" {
			req.SetBasicAuth(c.username, c.password)
		}
		logger.Debug("requesting DDS data", "url", fullURL)
		response, err := c.httpClient.Do(req)
		if err != nil {
			logger.Debug("failed to fetch DDS data", "error", err)
			return nil, err
		}
		logger.Debug("fetched DDS data", "url", fullURL, "status", response.Status)
		mediaType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
		if response.StatusCode == http.StatusUnauthorized {
			return nil, ErrUnauthorized
		} else if mayHaveExt && i == 0 &&
			( // DDS v2 version without JSON only endpoints (no extention)
			response.StatusCode == http.StatusNotFound ||
				// Future DDS v2 version without JSON support for old XML endpoints
				response.StatusCode == http.StatusNotAcceptable ||
				// DDS v1 compatible mode
				err != nil || mediaType != "application/json") {
			c.useXmlExt.Store(!useXmlExt)
			continue
		} else if response.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("unexpected HTTP status (%s)", response.Status)
		}
		defer response.Body.Close()
		return io.ReadAll(response.Body)
	}
	return nil, errors.New("unexpected return while requesting DDS data")
}

func (c *Client) GetRawIndex(ctx context.Context) ([]byte, error) {
	return c.GetRaw(IndexPath)
}

func (c *Client) GetRoot(ctx context.Context) (*Response, error) {
	return c.Get(RootPath)
}

func (c *Client) GetRawContained(ctx context.Context, resource string) ([]byte, error) {
	return c.GetRaw(ContainedPath, "resource", resource)
}

func (c *Client) GetCachedTimeOffset() time.Duration {
	timeData := c.ensureTimeData()
	if timeData != nil {
		return timeData.LocalStart.Sub(timeData.UTCStart.Time)
	}
	return DefaultTimeOffset
}

func (c *Client) ensureTimeData() *TimeData {
	c.rwMutex.RLock()
	timeData := c.timeData
	c.rwMutex.RUnlock()
	if timeData == nil {
		timeData = c.updateTimeData()
	}
	return timeData
}

func (c *Client) updateTimeData() *TimeData {
	logger := log.Logger.With("func", "updateTimeData")
	result, _, _ := c.single.Do("timeData", func() (any, error) {
		response, err := c.Get(PerformPath, "resource", ",,SYSPLEX", "id", "8D0D50")
		if err != nil {
			logger.Error("unable to fetch DDS time data", "error", err)
			return nil, err
		}
		timeData := response.Reports[0].TimeData
		if timeData == nil {
			logger.Error("unable to fetch DDS time data", "error", "no time data in DDS response")
			return nil, err
		}
		c.rwMutex.Lock()
		c.timeData = timeData
		c.rwMutex.Unlock()
		logger.Debug("DDS time data updated")
		return timeData, nil
	})
	if result != nil {
		return result.(*TimeData)
	}
	return nil
}

func (c *Client) GetCachedMintime() time.Duration {
	timeData := c.ensureTimeData()
	minTime := DefaultMinTime
	if timeData != nil && timeData.MinTime.Value != 0 {
		minTime = timeData.MinTime.Value
	}
	return time.Duration(minTime) * time.Second
}

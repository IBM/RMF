/**
* (C) Copyright IBM Corp. 2023, 2024.
* (C) Copyright Rocket Software, Inc. 2023-2024.
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
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
)

const UpdateInterval = 15 * time.Minute
const DefaultTimeOffset = 0

const IndexPath = "/gpm/index.xml"
const RootPath = "/gpm/root.xml"
const ContainedPath = "/gpm/contained.xml"
const PerformPath = "/gpm/perform.xml"
const XslHeadersPath = "/gpm/include/reptrans.xsl"
const FullReportPath = "/gpm/rmfm3.xml"

var ErrParse = errors.New("unable to parse DDS response")
var ErrUnauthorized = errors.New("not authorized to access DDS")

type Client struct {
	baseUrl    string
	username   string
	password   string
	httpClient *http.Client
	headerMap  *HeaderMap
	timeData   *TimeData

	stopChan  chan struct{}
	closeOnce sync.Once
	waitGroup sync.WaitGroup
	rwMutex   sync.RWMutex
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
	c.updateTimeData()
	c.GetHeaders()
	for {
		select {
		case <-ticker.C:
			c.updateTimeData()
			c.GetHeaders()
		case <-c.stopChan:
			logger.Debug("DDS background sync stopped")
			return
		}
	}
}

func (c *Client) Close() {
	c.closeOnce.Do(func() {
		close(c.stopChan)
		c.waitGroup.Wait()
		c.httpClient.CloseIdleConnections()
	})
}

func (c *Client) Get(ctx context.Context, path string, params ...string) (*Response, error) {
	var response Response
	data, err := c.GetRaw(ctx, path, params...)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrParse, err)
	}
	return &response, nil
}

func (c *Client) GetRaw(ctx context.Context, path string, params ...string) ([]byte, error) {
	logger := log.Logger.With("func", "GetRaw")
	path = strings.TrimLeft(path, "/")
	values := url.Values{}
	for i := 0; i < len(params)-1; i += 2 {
		if i+1 >= len(params) {
			return nil, errors.New("params must be key-value pairs of strings")
		}
		values.Add(params[i], params[i+1])
	}
	fullURL := fmt.Sprintf("%s/%s?%s", c.baseUrl, path, values.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, http.NoBody)
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
	if response.StatusCode == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status (%s)", response.Status)
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

func (c *Client) GetRawIndex(ctx context.Context) ([]byte, error) {
	return c.GetRaw(ctx, IndexPath)
}

func (c *Client) GetRoot(ctx context.Context) (*Response, error) {
	return c.Get(ctx, RootPath)
}

func (c *Client) GetRawContained(ctx context.Context, resource string) ([]byte, error) {
	return c.GetRaw(ctx, ContainedPath, "resource", resource)
}

func (c *Client) GetCachedTimeOffset() time.Duration {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	if c.timeData == nil {
		c.updateTimeData()
	}
	if c.timeData != nil {
		return c.timeData.LocalStart.Sub(c.timeData.UTCStart.Time)
	}
	return DefaultTimeOffset
}

func (c *Client) updateTimeData() {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	logger := log.Logger.With("func", "getTimeOffset")
	response, err := c.Get(context.Background(), PerformPath, "resource", ",,SYSPLEX", "id", "8D0D50")
	if err != nil {
		logger.Error("unable to fetch DDS time data", "error", err)
	}
	timeData := response.Reports[0].TimeData
	if timeData == nil {
		logger.Error("unable to fetch DDS time data", "error", "no time data in DDS response")
	}
	c.timeData = timeData
	logger.Debug("DDS time data updated")
}

func (c *Client) GetCachedMintime() int {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	if c.timeData == nil {
		c.updateTimeData()
	}
	if c.timeData != nil {
		return c.timeData.MinTime.Value
	}
	return 0
}

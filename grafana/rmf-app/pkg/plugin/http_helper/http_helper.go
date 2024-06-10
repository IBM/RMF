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

package http_helper

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

const DefaultHttpTimeout = 60

type HttpHelper struct {
}

func (h *HttpHelper) GetBaseUrl(em *typ.DatasourceEndpointModel) string {
	return em.URL
}

func (h *HttpHelper) GetQueryType(qm *typ.QueryModel) string {
	var resultQueryType string
	if strings.Trim(qm.SelectedQuery, "") != "" {
		splitStringSlice := strings.SplitAfter(qm.SelectedQuery, ".")
		if len(splitStringSlice) > 1 {
			vt := splitStringSlice[1]
			if strings.ToLower(vt) == "report." {
				resultQueryType = "report"
			} else {
				resultQueryType = "gauge"
			}
		}
	}
	return resultQueryType
}

func (h *HttpHelper) GetHttpUrlFromContext(qm *typ.QueryModel, em *typ.DatasourceEndpointModel) (string, error) {
	urlPath := ""
	if h.GetQueryType(qm) == "report" {
		urlPath = "/gpm/rmfm3.xml?"
	} else {
		urlPath = "/gpm/perform.xml?"
	}
	timeRange, err := h.GetTimeRange(qm, em)
	if err != nil {
		return "", err
	}
	return h.GetBaseUrl(em) + urlPath + qm.SelectedResource.Value + "&range=" + timeRange, nil
}

func (h *HttpHelper) GetHttpUrlFromQuery(query string, em *typ.DatasourceEndpointModel) string {
	urlPath := "/gpm/contained.xml?resource=" + query
	return h.GetBaseUrl(em) + urlPath
}

func (h *HttpHelper) GetHttpUrlForRoot(em *typ.DatasourceEndpointModel) string {
	return h.GetBaseUrl(em) + "/gpm/root.xml"
}

func (h *HttpHelper) GetHttpUrlForIndex(em *typ.DatasourceEndpointModel) string {
	return h.GetBaseUrl(em) + "/gpm/index.xml"
}

func (h *HttpHelper) GetHttpUrlForTimezoneOffset(em *typ.DatasourceEndpointModel) string {
	return h.GetBaseUrl(em) + "/gpm/perform.xml?resource=,,SYSPLEX&id=8D0D50"
}

func (h *HttpHelper) ExecuteHttpGet(queryURL string, em *typ.DatasourceEndpointModel) (*http.Response, error) {
	return executeHttpGetRequest(queryURL, em)
}

func (h *HttpHelper) GetHttpUrlForReportXsl(em *typ.DatasourceEndpointModel) string {
	return h.GetBaseUrl(em) + "/gpm/include/reptrans.xsl"
}

func (h *HttpHelper) GetXslFileContents(queryURL string, em *typ.DatasourceEndpointModel) (string, error) {
	resp, err := executeHttpGetRequestInternal(queryURL, em, true)
	if err != nil {
		return "", fmt.Errorf("GET error in GetXslFileContents(): %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status error in GetXslFileContents(): %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body in GetXslFileContents(): %w", err)
	}

	return string(data), nil
}

func executeHttpGetRequest(queryURL string, em *typ.DatasourceEndpointModel) (*http.Response, error) {
	return executeHttpGetRequestInternal(queryURL, em, false)
}

func executeHttpGetRequestInternal(queryURL string, em *typ.DatasourceEndpointModel, isXmlFileRequest bool) (*http.Response, error) {
	var client *http.Client

	// TODO: pass the proper context
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, queryURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	if isXmlFileRequest {
		req.Header.Add("Accept", "text/xml")
	} else {
		req.Header.Add("Accept", "application/json")
	}

	client = &http.Client{
		Timeout: time.Duration(em.IntTimeout) * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: em.TlsSkipVerify, // #nosec G402
			},
		},
	}
	if strings.TrimSpace(em.UserName) != "" {
		req.SetBasicAuth(em.UserName, em.Password)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not complete execution of executeHttpGetRequestInternal() - data source not reachable - error=%w", err)
	} else if res.StatusCode == 400 { // Bad request
		return nil, errors.New("bad request (Status Code 400) in executeHttpGetRequestInternal(). please check the data source configuration")
	} else if res.StatusCode == 401 { // Bad request
		return nil, errors.New("unauthorized (Status Code 401) in executeHttpGetRequestInternal(). please check the data source configuration")
	}

	return res, err
}

func (h *HttpHelper) GetTimeRange(qm *typ.QueryModel, em *typ.DatasourceEndpointModel) (string, error) {
	const DATETIMELAYOUT = "20060102150405"
	var (
		serverFromTime, serverToTime time.Time
	)
	if qm.SelectedVisualisationType == "TimeSeries" {
		if qm.AbsoluteTimeSelected {
			serverFromTime = qm.TimeSeriesTimeRangeFrom.Add(qm.ServerTimeData.ServerTimezoneOffset)
			serverToTime = qm.TimeSeriesTimeRangeFrom.Add(qm.ServerTimeData.ServerTimezoneOffset)
		} else {
			serverFromTime = qm.TimeSeriesTimeRangeTo.Add(qm.ServerTimeData.ServerTimezoneOffset)
			serverToTime = qm.TimeSeriesTimeRangeTo.Add(qm.ServerTimeData.ServerTimezoneOffset)
		}
	} else {
		serverFromTime = qm.TimeRangeFrom.Add(qm.ServerTimeData.ServerTimezoneOffset)
		serverToTime = qm.TimeRangeTo.Add(qm.ServerTimeData.ServerTimezoneOffset)
	}
	return serverFromTime.Format(DATETIMELAYOUT) + "," + serverToTime.Format(DATETIMELAYOUT), nil
}

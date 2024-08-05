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

package types

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

const TimeSeriesType = "TimeSeries"

type QueryModel struct {
	SelectedQuery string `json:"selectedQuery"`
	// FIXME: it contains also metric ID and needs to be re-parsed, e.g. id=8D21B0&resource=,,SYSPLEX
	SelectedResource          SelectedResource `json:"selectedResource"`
	RefreshRequired           bool             `json:"refreshRequired"`
	AbsoluteTimeSelected      bool             `json:"absoluteTimeSelected"`
	DashboardUid              string           `json:"dashboardUid"`
	SelectedVisualisationType string           `json:"selectedVisualisationType"`
	RMFPanelId                string           `json:"rmfPanelGuid"`
	TimeRangeFrom             time.Time        // 'From' time converted to UTC
	TimeRangeTo               time.Time        // 'To' time converted to UTC
	TimeSeriesTimeRangeFrom   time.Time        // 'From' Time converted to (DDS) local server time for timeseries
	TimeSeriesTimeRangeTo     time.Time        // 'To' Time converted to (DDS) local server time for timeseries
	ServerTimeData            DDSTimeData
}

func FromDataQuery(query backend.DataQuery) (*QueryModel, error) {
	var qm QueryModel
	if err := json.Unmarshal(query.JSON, &qm); err != nil {
		return nil, err
	}
	qm.TimeRangeFrom, qm.TimeRangeTo = query.TimeRange.From.UTC(), query.TimeRange.To.UTC()
	return &qm, nil
}

func (q *QueryModel) Copy() *QueryModel {
	copy := *q
	return &copy
}

func (qm *QueryModel) getQueryType() string {
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

func (qm *QueryModel) getTimeRange() string {
	var (
		serverFromTime, serverToTime time.Time
	)
	if qm.SelectedVisualisationType == TimeSeriesType {
		if qm.AbsoluteTimeSelected {
			serverFromTime = qm.TimeSeriesTimeRangeFrom.Add(qm.ServerTimeData.TimeOffset)
			serverToTime = qm.TimeSeriesTimeRangeFrom.Add(qm.ServerTimeData.TimeOffset)
		} else {
			serverFromTime = qm.TimeSeriesTimeRangeTo.Add(qm.ServerTimeData.TimeOffset)
			serverToTime = qm.TimeSeriesTimeRangeTo.Add(qm.ServerTimeData.TimeOffset)
		}
	} else {
		serverFromTime = qm.TimeRangeFrom.Add(qm.ServerTimeData.TimeOffset)
		serverToTime = qm.TimeRangeTo.Add(qm.ServerTimeData.TimeOffset)
	}
	return serverFromTime.Format(dds.DateTimeFormat) + "," + serverToTime.Format(dds.DateTimeFormat)
}

func (qm *QueryModel) GetPathWithParams() (string, []string) {
	var path string
	if qm.getQueryType() == "report" {
		path = dds.FullReportPath
	} else {
		path = dds.PerformPath
	}
	paramList := []string{"range", qm.getTimeRange()}
	// FIXME: process errors
	params, _ := url.ParseQuery(qm.SelectedResource.Value)
	for key, values := range params {
		paramList = append(paramList, key, strings.Join(values, ","))
	}
	return path, paramList
}

func (q *QueryModel) CacheKey() []byte {
	return []byte(q.SelectedResource.Value)
}

type DDSTimeData struct {
	LocalStartTime time.Time     // The local start time of the DDS Server
	LocalEndTime   time.Time     // The local end time of the local DDS Server
	LocalPrevTime  time.Time     // The previous time (current - gathererInterval) of the local DDS Server
	LocalNextTime  time.Time     // The next time (current + gathererInterval) of the local DDS Server
	UTCStartTime   time.Time     // The UTC start time corresponding to the LocalStartTime
	UTCEndTime     time.Time     // The UTC end time corresponding to the LocalEndTime
	MinTime        int           // GathererInterval: Time to wait before fetching data chunks from server (usually set at 100 (secs)).
	TimeOffset     time.Duration // The timezone offset value from UTC time
}

type SelectedResource struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

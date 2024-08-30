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

type ResponseStatus struct {
	TimeOffset time.Duration // The timezone offset value from UTC time
	Mintime    int
	LocalPrev  time.Time
	LocalNext  time.Time

	CurrentTime time.Time
}

func (rs *ResponseStatus) Update(other *ResponseStatus) {
	rs.TimeOffset = other.TimeOffset
	rs.Mintime = other.Mintime
	rs.LocalPrev = other.LocalPrev
	rs.LocalNext = other.LocalNext
	rs.CurrentTime = other.CurrentTime
}

func (rs *ResponseStatus) UpdateFromTimeData(timeData *dds.TimeData) {
	rs.TimeOffset = timeData.LocalStart.Sub(timeData.UTCStart.Time)
	rs.Mintime = timeData.MinTime.Value
	rs.LocalPrev = timeData.LocalPrev.Time
	rs.LocalNext = timeData.LocalNext.Time
	//ensure CurrentTime inside interval
	//currentMiddle = S+(E-S)/2
	currentMiddle := timeData.LocalStart.Time.Add(
		time.Duration(
			timeData.LocalEnd.Time.Sub(timeData.LocalStart.Time).Nanoseconds() / 2,
		),
	)
	currentMiddle = currentMiddle.Add(-1 * rs.TimeOffset)
	rs.CurrentTime = currentMiddle
}

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

	ResponseStatus
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

func (qm *QueryModel) getTime() string {
	var time time.Time
	if qm.SelectedVisualisationType == TimeSeriesType {
		time = qm.CurrentTime.Add(qm.TimeOffset)
	} else {
		time = qm.TimeRangeFrom.Add(qm.TimeOffset)
	}
	return time.Format(dds.DateTimeFormat)
}

func (qm *QueryModel) GetPathWithParams() (string, []string) {
	var path string
	if qm.getQueryType() == "report" {
		path = dds.FullReportPath
	} else {
		path = dds.PerformPath
	}
	paramList := []string{"time", qm.getTime()}
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

type SelectedResource struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

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
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type Request struct {
	Resource  string
	TimeRange data.TimeRange
	Batched   bool
}

func NewRequest(res string, from time.Time, to time.Time, step time.Duration) *Request {
	q := Request{Resource: res, TimeRange: data.TimeRange{From: from, To: to}, Batched: false}
	q.Align(step)
	return &q
}

func NewBatchRequest(res string, t time.Time, step time.Duration) *Request {
	startHour := t.Truncate(step)
	endHour := startHour.Add(step)
	q := Request{Resource: res, TimeRange: data.TimeRange{From: startHour, To: endHour}, Batched: true}
	return &q
}

func (r *Request) Align(d time.Duration) {
	r.TimeRange.From = r.TimeRange.From.Truncate(d)
	t := r.TimeRange.To.Truncate(d)
	if t.Equal(r.TimeRange.From) || t.Before(r.TimeRange.To) {
		r.TimeRange.To = t.Add(d)
	}
}

func (r *Request) Add(d time.Duration) {
	r.TimeRange.From = r.TimeRange.From.Add(d)
	r.TimeRange.To = r.TimeRange.To.Add(d)
}

func (r *Request) String() string {
	return fmt.Sprintf("%s [%s - %s]", r.Resource, r.TimeRange.From, r.TimeRange.To)
}

func (r *Request) formatRange(timeOfs time.Duration) string {
	from := r.TimeRange.From
	to := r.TimeRange.To
	return from.Add(timeOfs).Format(DateTimeFormat) + "," + to.Add(timeOfs).Format(DateTimeFormat)
}

func (r *Request) pathWithParams(timeOfs time.Duration) (string, []string, error) {
	path := ""
	rawParams, err := url.ParseQuery(r.Resource)
	if err != nil {
		return "", nil, err
	}
	params := make([]string, 0, 1)
	for key, values := range rawParams {
		if key == "report" {
			path = FullReportPath
		}
		params = append(params, key, strings.Join(values, ";"))
	}
	if path == "" {
		path = PerformPath
	}
	var rangeParam string = "range"
	if r.Batched {
		path = TimeSeriesPath
		timeOfs = 0
		rangeParam = "rangeUtc"
	}
	params = append(params, rangeParam, r.formatRange(timeOfs))
	return path, params, nil
}

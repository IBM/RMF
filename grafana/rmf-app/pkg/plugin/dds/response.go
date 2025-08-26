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
	"fmt"
	"strings"
	"time"
)

const DateTimeFormat = "20060102150405"

const ReportFormat = "report"
const ListFormat = "list"
const SingleFormat = "single"

const NumericColType = "N"
const TextColType = "T"

var SupportedFormats = map[string]bool{
	ReportFormat: true,
	ListFormat:   true,
	SingleFormat: true,
}
var AcceptableMessages = map[string]bool{
	"GPM0501I": true, // Data may be inconsistent due to a change of the WLM service policy or IPS
	"GPM0502I": true, // DDS has returned partial data because time gaps have been detected
	"GPM0504I": true, // Data may be inconsistent due to a change of the WLM service policy
	"GPM0505I": true, // Data may be inconsistent due to a change of the RMF cycle time
	"GPM0506I": true, // IPL/REBOOT detected in the specified time range
	"GPM0507I": true, // DDS could not retrieve valid data for the specified date and time range
	"GPM0709I": true, // Filter has caused no data to be returned
}

type Response struct {
	Reports []Report `json:"report"`
}

type Report struct {
	TimeData *TimeData
	Metric   *Metric
	Message  *Message
	Caption  Caption
	Resource *Resource `json:"resource"`
	Headers  struct {
		Cols []Col `json:"col"`
	} `json:"columnHeaders"`
	Rows []Row `json:"row"`
}

func (r Report) GetRowNames() []string {
	var names []string
	if r.Rows != nil {
		for _, row := range r.Rows {
			names = append(names, row.Cols[0])
		}
	}
	return names
}

type TimeData struct {
	// FIXME: don't use these in report headers: they are in DDS timezone. Remove from the mapping.
	LocalStart DateTime `json:"localStart"`
	LocalEnd   DateTime `json:"localEnd"`
	LocalPrev  DateTime `json:"localPrev"`
	LocalNext  DateTime `json:"localNext"`
	UTCStart   DateTime `json:"utcStart"`
	UTCEnd     DateTime `json:"utcEnd"`
	NumSamples int      `json:"numSamples"`
	NumSystems *int     `json:"numSystems,omitempty"`
	MinTime    struct {
		Unit  string `json:"unit"`
		Value int    `json:"value"`
	} `json:"gathererInterval"`
	DataRange struct {
		Unit  string `json:"unit"`
		Value int    `json:"value"`
	} `json:"dataRange"`
}

type Resource struct {
	Reslabel string `json:"reslabel"`
	Restype  string `json:"restype"`
}

func (r Resource) GetName() string {
	ss := strings.Split(r.Reslabel, ",")
	if len(ss) > 1 {
		return ss[1]
	}
	return r.Reslabel
}

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(DateTimeFormat, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*dt = DateTime{Time: parsed}
	return nil
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if dt == nil {
		return nil, nil
	}
	return []byte(`"` + dt.Format(DateTimeFormat) + `"`), nil
}

type Message struct {
	Id          string
	Severity    int
	Description string
}

func (msg Message) Error() string {
	return fmt.Sprintf("DDS error: %s (severity %d). %s", msg.Id, msg.Severity, msg.Description)
}

type Metric struct {
	Id          string
	Description string
	Format      string
	NumCols     int
	Filter      string
}

type Col struct {
	Id   string `json:"value"`
	Type string `json:"type"`
}

type Row struct {
	Cols []string `json:"col"`
}

type Caption struct {
	Vars []struct {
		Name  string
		Value string
	} `json:"var"`
}

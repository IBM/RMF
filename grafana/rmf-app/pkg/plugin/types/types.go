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
	"fmt"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type DatasourceEndpointModel struct {
	Server             string `json:"path"`
	Port               string `json:"port"`
	SSL                bool   `json:"ssl"`
	UserName           string `json:"userName"`
	Password           string `json:"password"`
	VerifyInsecureCert bool   `json:"skipVerify"`
	DatasourceUid      string
}

type Panel struct {
	PanelGuid         string
	QueryModels       []*QueryModel
	DataLastFetchedAt time.Time // Keeps track of when the last data was fetched.
}

type QueryModel struct {
	Datasource                Datasource       `json:"datasource"`
	SelectedQuery             string           `json:"selectedQuery"`
	SelectedResource          SelectedResource `json:"selectedResource"`
	RefreshRequired           bool             `json:"refreshRequired"`
	AbsoluteTimeSelected      bool             `json:"absoluteTimeSelected"`
	DashboardUid              string           `json:"dashboardUid"`
	SelectedVisualisationType string           `json:"selectedVisualisationType"`
	RMFPanelId                string           `json:"rmfPanelGuid"`
	RefID                     string
	TimeRangeFrom             time.Time // 'From' time converted to UTC
	TimeRangeTo               time.Time // 'To' time converted to UTC
	TimeSeriesTimeRangeFrom   time.Time // 'From' Time converted to (DDS) local server time for timeseries
	TimeSeriesTimeRangeTo     time.Time // 'To' Time converted to (DDS) local server time for timeseries
	UniquePath                string    // A Unique guid that is used as the streaming PATH
	ServiceCallInProgress     bool      // True if a service call is in progress, else false
	ServerTimeData            DDSTimeData
	Url                       string // Stores the DDS Url invoked to get response data
}

type DDSResponse struct {
	Reports []struct {
		Metric struct {
			Format string
		}
		Rows []struct {
			Cols []string `json:"col"`
		} `json:"row"`
	} `json:"report"`
}

type DDSMessage struct {
	Id          string
	Severity    int
	Description string
}

func (msg *DDSMessage) String() string {
	return fmt.Sprintf("%s (Sev: %d) %s", msg.Id, msg.Severity, msg.Description)
}

type DDSTimeData struct {
	LocalStartTime       time.Time     // The local start time of the DDS Server
	LocalEndTime         time.Time     // The local end time of the local DDS Server
	LocalPrevTime        time.Time     // The previous time (current - gathererInterval) of the local DDS Server
	LocalNextTime        time.Time     // The next time (current + gathererInterval) of the local DDS Server
	UTCStartTime         time.Time     // The UTC start time corresponding to the LocalStartTime
	UTCEndTime           time.Time     // The UTC end time corresponding to the LocalEndTime
	ServiceCallInterval  float64       // GathererInterval: Time to wait before fetching data chunks from server (usually set at 100 (secs)).
	ServerTimezoneOffset time.Duration // The timezone offset value from UTC time
}

type IntervalOffset struct {
	ServiceCallInterval  float64       // GathererInterval: Time to wait before fetching data chunks from server (usually set at 100 (secs)).
	ServerTimezoneOffset time.Duration // The timezone offset value from UTC time
}

type Datasource struct {
	Type string `json:"type"`
	Uid  string `json:"uid"`
}

type SelectedResource struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ColHeaderXslMap struct {
	HeaderID   string
	HeaderText string
}

// Below structures are used exclusively for caching

type CacheItems struct {
	CacheItemKey    string
	CacheItemValues []CacheItemValue
}

type CacheItemValue struct {
	ValueKey       time.Time
	Value          data.Frame
	ServerTimeData DDSTimeData
}

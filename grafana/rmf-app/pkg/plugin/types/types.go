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
	// Conventional Grafana HTTP config (see the `DataSourceHttpSettings` UI element)
	// TODO: the type is to get rid of. We need to re-use one HTTP client on DS level.
	URL           string
	IntTimeout    int
	RawTimeout    string `json:"timeout"`
	TlsSkipVerify bool   `json:"tlsSkipVerify"`
	// Legacy: custom RMF settings. We should ge rid of it at some point.
	Server   string `json:"path"`
	Port     string `json:"port"`
	SSL      bool   `json:"ssl"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	// NB: the meaning of JSON field is inverted. If set, we do verify certificates.
	VerifyInsecureCert bool `json:"skipVerify"`
	DatasourceUid      string
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
	TimeRangeFrom             time.Time        // 'From' time converted to UTC
	TimeRangeTo               time.Time        // 'To' time converted to UTC
	TimeSeriesTimeRangeFrom   time.Time        // 'From' Time converted to (DDS) local server time for timeseries
	TimeSeriesTimeRangeTo     time.Time        // 'To' Time converted to (DDS) local server time for timeseries
	ServerTimeData            DDSTimeData
	Url                       string // Stores the DDS Url invoked to get response data
}

func (query *QueryModel) Copy() *QueryModel {
	copy := *query
	return &copy
}

func (query *QueryModel) CacheKey() []byte {
	return []byte(query.Datasource.Uid + "/" + query.SelectedResource.Value)
}

type VariableQueryRequest struct {
	Query string `json:"query"`
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
	LocalStartTime time.Time     // The local start time of the DDS Server
	LocalEndTime   time.Time     // The local end time of the local DDS Server
	LocalPrevTime  time.Time     // The previous time (current - gathererInterval) of the local DDS Server
	LocalNextTime  time.Time     // The next time (current + gathererInterval) of the local DDS Server
	UTCStartTime   time.Time     // The UTC start time corresponding to the LocalStartTime
	UTCEndTime     time.Time     // The UTC end time corresponding to the LocalEndTime
	MinTime        float64       // GathererInterval: Time to wait before fetching data chunks from server (usually set at 100 (secs)).
	TimeOffset     time.Duration // The timezone offset value from UTC time
}

type IntervalTimeData struct {
	MinTime    float64
	TimeOffset time.Duration // Time offset from UTC time
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

type ValueError struct {
	Value int
	Err   error
}

func NewValueError(value int, err error) *ValueError {
	return &ValueError{
		Value: value,
		Err:   err,
	}
}

func (ve *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", ve.Err)
}

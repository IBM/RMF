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

package main

type Response struct {
	Reports                []Report                 `json:"report"`
	Server                 *ServerInfo              `json:"server,omitempty"`
	Messages               []Message                `json:"message"`
	MetricList             []MetricListEntry        `json:"metricList"`
	AttributeList          []any                    `json:"attributeList"`
	FilterInstancesList    []any                    `json:"filterInstancesList"`
	Postprocessor          []any                    `json:"postprocessor"`
	WorkscopeList          []any                    `json:"workscopeList"`
	ContainedResourcesList []ContainedResourcesList `json:"containedResourcesList"`
}

type ServerInfo struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	Platform      string `json:"platform"`
	Functionality string `json:"functionality"`
}

type Report struct {
	TimeData      *TimeData      `json:"timeData,omitempty"`
	Metric        *Metric        `json:"metric,omitempty"`
	Message       *Message       `json:"message,omitempty"`
	Caption       *Caption       `json:"caption,omitempty"`
	Resource      *Resource      `json:"resource,omitempty"`
	ColumnHeaders *ColumnHeaders `json:"columnHeaders,omitempty"`
	Rows          []Row          `json:"row,omitempty"`
}

type TimeData struct {
	LocalStart       string       `json:"localStart"`
	LocalEnd         string       `json:"localEnd"`
	UTCStart         string       `json:"utcStart"`
	UTCEnd           string       `json:"utcEnd"`
	LocalPrev        string       `json:"localPrev"`
	LocalNext        string       `json:"localNext"`
	NumSamples       int          `json:"numSamples"`
	NumSystems       *int         `json:"numSystems,omitempty"`
	GathererInterval TimeInterval `json:"gathererInterval"`
	DataRange        TimeInterval `json:"dataRange"`
}

type TimeInterval struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Metric struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Format      string `json:"format"`
	NumCols     int    `json:"numcols"`
	Unit        string `json:"unit,omitempty"`
	Filter      string `json:"filter,omitempty"`
	ListType    string `json:"listtype,omitempty"`
	Workscope   string `json:"workscope,omitempty"`
}

type Message struct {
	Id          string `json:"id,omitempty"`
	Severity    int    `json:"severity,omitempty"`
	Description string `json:"description,omitempty"`
}

type Caption struct {
	Vars []CaptionVar `json:"var,omitempty"`
}

type CaptionVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Resource struct {
	Reslabel   string `json:"reslabel"`
	Restype    string `json:"restype,omitempty"`
	Expandable string `json:"expandable,omitempty"`
	Icon       string `json:"icon,omitempty"`
}

type ColumnHeaders struct {
	Cols []Col `json:"col"`
}

type Col struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Row struct {
	Cols []string `json:"col"`
}

type ContainedResourcesList struct {
	Contained *ContainedResources `json:"contained,omitempty"`
	Resource  *Resource           `json:"resource,omitempty"`
}

type ContainedResources struct {
	Resources []Resource `json:"resource"`
}

// MetricListEntry represents one entry in the /gpm/index response.
type MetricListEntry struct {
	ResourceType string            `json:"resourceType"`
	Metrics      []IndexMetricItem `json:"metric"`
}

type IndexMetricItem struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Format      string `json:"format"`
}

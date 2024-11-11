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

package frame

import (
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func getFrameName(qm *QueryModel) string {
	var resultFrameName string
	if strings.Trim(qm.SelectedQuery, " ") != "" {
		splitStringSlice := strings.SplitAfter(qm.SelectedQuery, ".")
		if len(splitStringSlice) > 1 {
			vt := splitStringSlice[1]
			if strings.ToLower(vt) == "report." {
				resultFrameName = splitStringSlice[1] + splitStringSlice[2]
			} else {
				if strings.Contains(vt, "{") {
					resultFrameName = vt[:strings.Index(vt, "{")]
				} else {
					resultFrameName = vt
				}
			}
		}
	}
	return strings.Trim(resultFrameName, " ")
}

// getFrameLabels builds labels based on DDS metric name and type
func getFrameLabels(metricType string, queryName string) data.Labels {
	labels := data.Labels{}
	if metricType == "list" {
		labels["metric"], _ = splitQueryName(queryName)
	}
	return labels
}

// splitQueryName splits metric name into short metric name and resource type.
// For example, `% AAP by job` becomes pair `% AAP` and `job`
func splitQueryName(queryName string) (string, string) {
	shortQueryName, itemType, _ := strings.Cut(queryName, " by ")
	return strings.TrimSpace(shortQueryName), strings.TrimSpace(itemType)
}

// SeriesFields stores info about fields that appeared in time series.
type SeriesFields map[string]FieldInfo
type FieldInfo struct {
	Time   time.Time
	Labels data.Labels
}

// SyncFieldNames adds Fields having names from the map with `nil` values if they are not present
// in the current frame.
// Required for frames streamed as time series.
// If we send a frame without a field name that was there in previous frames of the same time series,
// values for the field in those frames will be discarded by frontend.
func SyncFieldNames(seriesFields SeriesFields, frame *data.Frame, frameTime time.Time) {
	fieldNames := map[string]bool{}
	for _, field := range frame.Fields {
		seriesFields[field.Name] = FieldInfo{Time: frameTime, Labels: field.Labels}
		fieldNames[field.Name] = true
	}
	for key := range seriesFields {
		if _, ok := fieldNames[key]; !ok {
			newField := data.NewField(key, seriesFields[key].Labels, []*float64{nil})
			frame.Fields = append(frame.Fields, newField)
		}
	}
}

// RemoveOldFieldNames cuts off field names older than the given time.
func RemoveOldFieldNames(fieldMap SeriesFields, cutoffTime time.Time) {
	for name, fieldInfo := range fieldMap {
		if fieldInfo.Time.Before(cutoffTime) {
			delete(fieldMap, name)
		}
	}
}

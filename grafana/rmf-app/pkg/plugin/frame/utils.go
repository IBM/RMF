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

package frame

import (
	"errors"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

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
	frameRows := 1
	if len(frame.Fields) > 0 {
		frameRows = frame.Fields[0].Len()
	}
	for _, field := range frame.Fields {
		seriesFields[field.Name] = FieldInfo{Time: frameTime, Labels: field.Labels}
		fieldNames[field.Name] = true
	}
	for key := range seriesFields {
		if _, ok := fieldNames[key]; !ok {
			newField := data.NewField(key, seriesFields[key].Labels, make([]*float64, frameRows))
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

func MergeInto(dst *data.Frame, src *data.Frame) (*data.Frame, error) {
	if dst == nil {
		dst = data.NewFrame(src.Name)
	}
	if src == nil {
		return dst, nil
	}
	dstLen, err := dst.RowLen()
	if err != nil {
		return nil, err
	}
	srcLen, err := src.RowLen()
	if err != nil {
		return nil, err
	}
	for _, field2 := range src.Fields {
		field1, _ := dst.FieldByName(field2.Name)
		if field1 == nil {
			switch field2.Type() {
			case data.FieldTypeTime:
				field1 = data.NewField(field2.Name, field2.Labels, make([]time.Time, dstLen))
			case data.FieldTypeNullableFloat64:
				field1 = data.NewField(field2.Name, field2.Labels, make([]*float64, dstLen))
			default:
				return nil, errors.New("unsupported field type")
			}
			dst.Fields = append(dst.Fields, field1)
		}
		for i := range srcLen {
			field1.Append(field2.At(i))
		}
	}
	for _, field1 := range dst.Fields {
		if field2, _ := src.FieldByName(field1.Name); field2 == nil {
			for range srcLen {
				field1.Append(nil)
			}
		}
	}
	return dst, nil
}

func GetDuration(f *data.Frame) time.Duration {
	if f == nil {
		return 0
	}
	if len(f.Fields) == 0 {
		return 0
	}
	timeField := f.Fields[0]
	if timeField.Type() != data.FieldTypeTime {
		return 0
	}
	var minTime time.Time
	var maxTime time.Time
	for i := 0; i < timeField.Len(); i++ {
		t, ok := timeField.At(i).(time.Time)
		if ok {
			if minTime.IsZero() || t.Before(minTime) {
				minTime = t
			}
			if maxTime.IsZero() || t.After(maxTime) {
				maxTime = t
			}
		}
	}
	return maxTime.Sub(minTime)
}

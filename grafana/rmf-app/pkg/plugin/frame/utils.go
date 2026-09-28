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
	"fmt"
	"slices"
	"sort"
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
			case data.FieldTypeNullableString:
				field1 = data.NewField(field2.Name, field2.Labels, make([]*string, dstLen))
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

// MergeTimeSeries merges two time-series frames using timestamps as unique keys.
// When both frames contain the same timestamp, non-nil values from src replace
// corresponding values from dst.
func MergeTimeSeries(dst *data.Frame, src *data.Frame) (*data.Frame, error) {
	if dst == nil && src == nil {
		return nil, nil
	}

	frames := []*data.Frame{dst, src}
	fieldTemplates := make(map[string]*data.Field)
	rows := make(map[time.Time]map[string]any)
	var fieldOrder []string
	var timeFieldName string
	frameName := ""

	for _, frame := range frames {
		if frame == nil {
			continue
		}
		if frameName == "" {
			frameName = frame.Name
		}

		rowCount, err := frame.RowLen()
		if err != nil {
			return nil, err
		}

		var timeField *data.Field
		for _, field := range frame.Fields {
			if field.Type() == data.FieldTypeTime {
				if timeField != nil {
					return nil, errors.New("frame contains multiple time fields")
				}
				timeField = field
			}

			existing, found := fieldTemplates[field.Name]
			if found {
				if existing.Type() != field.Type() {
					return nil, fmt.Errorf("field %q has conflicting types", field.Name)
				}
			} else {
				fieldTemplates[field.Name] = field
				fieldOrder = append(fieldOrder, field.Name)
			}
		}

		if timeField == nil {
			return nil, errors.New("frame does not contain a time field")
		}
		if timeFieldName != "" && timeFieldName != timeField.Name {
			return nil, errors.New("frames have different time field names")
		}
		timeFieldName = timeField.Name

		for rowIndex := range rowCount {
			timestamp, ok := timeField.At(rowIndex).(time.Time)
			if !ok {
				return nil, fmt.Errorf(
					"invalid timestamp in field %q at row %d",
					timeField.Name,
					rowIndex,
				)
			}

			if rows[timestamp] == nil {
				rows[timestamp] = make(map[string]any)
			}

			for _, field := range frame.Fields {
				if field.Name == timeFieldName {
					continue
				}
				value := field.At(rowIndex)
				if value != nil {
					rows[timestamp][field.Name] = value
				}
			}
		}
	}

	timestamps := make([]time.Time, 0, len(rows))
	for timestamp := range rows {
		timestamps = append(timestamps, timestamp)
	}
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})

	result := data.NewFrame(frameName)
	for _, fieldName := range fieldOrder {
		template := fieldTemplates[fieldName]

		var field *data.Field
		switch template.Type() {
		case data.FieldTypeTime:
			field = data.NewField(fieldName, template.Labels, make([]time.Time, 0, len(timestamps)))
		case data.FieldTypeNullableFloat64:
			field = data.NewField(fieldName, template.Labels, make([]*float64, 0, len(timestamps)))
		case data.FieldTypeNullableString:
			field = data.NewField(fieldName, template.Labels, make([]*string, 0, len(timestamps)))
		default:
			return nil, fmt.Errorf("unsupported field type %s", template.Type())
		}

		field.SetConfig(template.Config)
		result.Fields = append(result.Fields, field)
	}

	for _, timestamp := range timestamps {
		for _, field := range result.Fields {
			if field.Name == timeFieldName {
				field.Append(timestamp)
			} else {
				field.Append(rows[timestamp][field.Name])
			}
		}
	}

	return result, nil
}

func CopyReportField(field *data.Field, length int) *data.Field {
	var newField *data.Field
	t := field.Type()
	switch t {
	case data.FieldTypeNullableFloat64:
		newField = data.NewField(field.Name, field.Labels, []*float64{})
	case data.FieldTypeNullableString:
		newField = data.NewField(field.Name, field.Labels, []*string{})
	default:
		newField = data.NewField(field.Name, field.Labels, []string{})
	}
	newField.SetConfig(field.Config)
	length = slices.Min([]int{length, field.Len()})
	for i := 0; i < length; i++ {
		newField.Append(field.At(i))
	}
	return newField
}

func GetFrameTable(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	for _, field := range f.Fields {
		if !strings.HasPrefix(field.Name, CaptionPrefix) &&
			!strings.HasPrefix(field.Name, BannerPrefix) {
			var newField = CopyReportField(field, field.Len())
			newField.Delete(0)
			newFrame.Fields = append(newFrame.Fields, newField)
		}
	}
	return &newFrame
}

func GetFrameCaption(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	newFrame.Fields = append(newFrame.Fields, data.NewField("Key", nil, []string{}))
	newFrame.Fields = append(newFrame.Fields, data.NewField("Value", nil, []string{}))
	for _, field := range f.Fields {
		if strings.HasPrefix(field.Name, CaptionPrefix) {
			key := strings.TrimPrefix(field.Name, CaptionPrefix)
			value := getStringAt(field, 0)
			newFrame.Fields[0].Append(key)
			newFrame.Fields[1].Append(value)
		}
	}
	return &newFrame
}

func GetFrameBanner(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	newFrame.Fields = append(newFrame.Fields, data.NewField("Key", nil, []string{}))
	newFrame.Fields = append(newFrame.Fields, data.NewField("Value", nil, []string{}))
	for _, field := range f.Fields {
		if strings.HasPrefix(field.Name, BannerPrefix) {
			key := strings.TrimPrefix(field.Name, BannerPrefix)
			value := getStringAt(field, 0)
			newFrame.Fields[0].Append(key)
			newFrame.Fields[1].Append(value)
		}
	}
	return &newFrame
}

func getStringAt(field *data.Field, index int) string {
	value := ""
	if field.Len() > index {
		v := field.At(index)
		if v != nil {
			if s, ok := v.(string); ok {
				value = s
			} else if s, ok := v.(*string); ok {
				if s != nil {
					value = *s
				}
			} else if f, ok := v.(float64); ok {
				value = fmt.Sprintf("%f", f)
			}
		}
	}
	return value
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

func GetMaxTime(f *data.Frame) time.Time {
	if f == nil {
		return time.Time{}
	}
	if len(f.Fields) == 0 {
		return time.Time{}
	}
	timeField := f.Fields[0]
	if timeField.Type() != data.FieldTypeTime {
		return time.Time{}
	}
	var maxTime time.Time
	for i := 0; i < timeField.Len(); i++ {
		t, ok := timeField.At(i).(time.Time)
		if ok {
			if maxTime.IsZero() || t.After(maxTime) {
				maxTime = t
			}
		}
	}
	return maxTime
}

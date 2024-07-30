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
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

func Build(ddsResponse *dds.Response, headers dds.HeaderMap, queryModel *typ.QueryModel) (*data.Frame, error) {
	logger := log.Logger.With("func", "Build")

	reportsNum := len(ddsResponse.Reports)
	if reportsNum == 0 {
		return nil, errors.New("no reports in DDS response")
	} else if reportsNum > 1 {
		return nil, fmt.Errorf("too many reports (%d) in DDS response", reportsNum)
	}
	report := ddsResponse.Reports[0]
	if report.TimeData == nil {
		return nil, errors.New("no time data in DDS response")
	}
	if report.Metric == nil {
		return nil, errors.New("no metric data in DDS response")
	}
	if _, ok := dds.SupportedFormats[report.Metric.Format]; !ok {
		return nil, fmt.Errorf("unsupported data format (%s) in DDS response", report.Metric.Format)
	}

	if message := report.Message; message != nil {
		// FIXME: format in the conventional way
		_, ok := dds.AcceptableMessages[message.Id]
		if !ok {
			return nil, fmt.Errorf("DDS error: %s", message)
		} else {
			logger.Warn(fmt.Sprintf("DDS error: %s", message))
		}
	}

	format := report.Metric.Format
	queryModel.ServerTimeData = copyTimeData(report.TimeData)
	var newFrame *data.Frame
	var err error
	if format == dds.ReportFormat {
		newFrame, err = buildForReport(&report, headers, queryModel)
		newFrame.Meta = &data.FrameMeta{Custom: report.TimeData}
	} else {
		newFrame = buildForMetric(&report, queryModel)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to build frame: %w", err)
	}
	return newFrame, nil
}

func copyTimeData(ddsTimeData *dds.TimeData) typ.DDSTimeData {
	// FIXME: this is redundant, just re-use dds.TimeData
	resultTimeData := typ.DDSTimeData{}
	resultTimeData.LocalStartTime = ddsTimeData.LocalStart.Time
	resultTimeData.LocalEndTime = ddsTimeData.LocalEnd.Time
	resultTimeData.LocalPrevTime = ddsTimeData.LocalPrev.Time
	resultTimeData.LocalNextTime = ddsTimeData.LocalNext.Time
	resultTimeData.UTCStartTime = ddsTimeData.UTCStart.Time
	resultTimeData.UTCEndTime = ddsTimeData.UTCEnd.Time
	// FIXME: sync with DDS client
	resultTimeData.TimeOffset = resultTimeData.LocalStartTime.Sub(resultTimeData.UTCStartTime)
	resultTimeData.MinTime = ddsTimeData.MintTime.Value

	// Convert all to UTC times.
	// FIXME: we already have UTC timestamps, the "local" time fields below are not really local.
	resultTimeData.LocalStartTime = resultTimeData.LocalStartTime.Add(-1 * resultTimeData.TimeOffset)
	resultTimeData.LocalEndTime = resultTimeData.LocalEndTime.Add(-1 * resultTimeData.TimeOffset)
	resultTimeData.LocalPrevTime = resultTimeData.LocalPrevTime.Add(-1 * resultTimeData.TimeOffset)
	resultTimeData.LocalNextTime = resultTimeData.LocalNextTime.Add(-1 * resultTimeData.TimeOffset)
	return resultTimeData
}

// buildForMetric parses JSON string and create a data frame either for time series or a regular one.
func buildForMetric(report *dds.Report, query *typ.QueryModel) *data.Frame {
	queryName := getFrameName(query)

	if query.SelectedVisualisationType == typ.TimeSeriesType {
		return buildWideForMetric(report, queryName)
	} else {
		return buildLongForMetric(report, queryName)
	}
}

// buildWideForMetric creates a time series data frame for a metric from pre-parsed DDS response.
// Grafana frame format: wide.
func buildWideForMetric(report *dds.Report, queryName string) *data.Frame {
	frameName := queryName
	timestamp := report.TimeData.UTCEnd.Time
	metricFormat := report.Metric.Format
	labels := getFrameLabels(metricFormat, queryName)
	resultFrame := data.NewFrame(frameName, data.NewField("time", nil, []time.Time{timestamp}))

	iterateMetricRows(report, queryName,
		func(name string, value *float64) {
			newField := data.NewField(name, labels, []*float64{value})
			resultFrame.Fields = append(resultFrame.Fields, newField)
		})

	// Built-in alerting requires either a frame in wide format with mandatory numeric fields,
	// or an empty one. However, empty frames won't work for time series views.
	// Solution for single type metric is to send nil values if there's no data.
	// For list type metrics, we don't have column names to do the same; it has to be fixed differently.
	if len(resultFrame.Fields) == 1 && metricFormat == "single" {
		newField := data.NewField(queryName, labels, []*float64{nil})
		resultFrame.Fields = append(resultFrame.Fields, newField)
	}

	return resultFrame
}

// buildLongForMetric creates a non time series data frame for a metric from pre-parsed DDS response.
// Grafana frame format: long.
func buildLongForMetric(report *dds.Report, queryName string) *data.Frame {
	metricFormat := report.Metric.Format
	nameField := "metric"
	timestamp := report.TimeData.UTCEnd.Time
	valField := queryName
	if metricFormat == "list" {
		valField, nameField = splitQueryName(queryName)
		if nameField == "" {
			nameField = queryName
			valField = "value"
		}
	}

	resultFrame := data.NewFrame("",
		data.NewField("time", nil, []time.Time{}),
		data.NewField(nameField, nil, []string{}),
		data.NewField(valField, nil, []*float64{}),
	)

	iterateMetricRows(report, queryName,
		func(name string, value *float64) {
			resultFrame.Fields[0].Append(timestamp)
			resultFrame.Fields[1].Append(name)
			resultFrame.Fields[2].Append(value)
		})

	return resultFrame
}

// iterateMetricRows parses metric key-value pairs and passes them to `process` while iterating over rows.
func iterateMetricRows(report *dds.Report, defaultName string, process func(name string, value *float64)) {
	for _, jsonRow := range report.Rows {
		cols := jsonRow.Cols
		name, rawValue := cols[0], cols[1]
		if name == "*NoData*" {
			continue
		}
		if len(jsonRow.Cols) == 3 {
			name += "[" + cols[2] + "]"
		}
		if name == "" {
			name = defaultName
		}
		var value *float64
		floatValue, err := strconv.ParseFloat(rawValue, 64)
		// rawValue can contain different kinds of text meaning n/a: NaN, blank value, Deact etc.
		if rawValue != "NaN" && err == nil {
			value = &floatValue
		}
		process(name, value)
	}
}

func buildForReport(report *dds.Report, headers dds.HeaderMap, qm *typ.QueryModel) (*data.Frame, error) {
	resultFrame := data.NewFrame(getFrameName(qm))
	// Add the regular metrics values as fields (either float64 or string values)
	var trackingIndexCol, trackingIndexRow int
	reportName := report.Metric.Id

	// Get a reference to columnHeaders Json property
	cols := report.Headers.Cols
	for colIndex, col := range cols {
		var stringSlice []string
		var floatSlice []float64

		rows := report.Rows

		// Some columns have type="N" (meaning Numeric) but contain "" values. In those cases it will be considered string.
		// FIXME: it causes all kind of issues with sorting in UI.
		treatColumnAsString := shouldColumnBeTreatedAsString(rows, colIndex)

		for rowIndex, row := range rows {
			currentValue := row.Cols[colIndex]
			if treatColumnAsString {
				stringSlice = append(stringSlice, currentValue)
			} else {
				f2, err := strconv.ParseFloat(currentValue, 64)
				if err != nil {
					return nil, fmt.Errorf("could not convert value to float in ConstructReportFrameFromJson(). Error=%w", err)
				}
				floatSlice = append(floatSlice, f2)
			}
			trackingIndexRow = rowIndex
		}
		headerText := headers.Get(reportName, col.Id)
		if treatColumnAsString {
			newField := data.NewField(headerText, nil, stringSlice)
			resultFrame.Fields = append(resultFrame.Fields, newField)
		} else {
			newField := data.NewField(headerText, nil, floatSlice)
			resultFrame.Fields = append(resultFrame.Fields, newField)
		}
		trackingIndexCol = colIndex
	}

	trackingIndexCol++
	trackingIndexRow++

	// Add the captions (these go as headers in the header table for Reports, not to be confused with col headers of reports)
	for _, variable := range report.Caption.Vars {
		fieldHeaderText := "Header::" + headers.Get(reportName, variable.Name)
		fieldValue := variable.Value
		fieldValues := make([]string, trackingIndexRow) // This slice will be having 'trackingIndexRow' rows
		fieldValues[0] = fieldValue                     // Set only the 0th value
		// Create a new field
		newField := data.NewField(fieldHeaderText, nil, fieldValues)
		resultFrame.Fields = append(resultFrame.Fields, newField)
		trackingIndexCol++
	}
	return resultFrame, nil
}

func shouldColumnBeTreatedAsString(rows []dds.Row, colIndex int) bool {
	treatAsString := false
	for _, row := range rows {
		currentValue := row.Cols[colIndex]
		_, err := strconv.ParseFloat(currentValue, 64)
		if err != nil {
			treatAsString = true
			break
		}
	}
	return treatAsString
}

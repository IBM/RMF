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
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
)

const BannerPrefix = "Banner::"
const CaptionPrefix = "Caption::"
const ReportDateFormat = "01/02/2006 15:04:05"

func TaggedFrame(t time.Time, tag string) *data.Frame {
	return data.NewFrame(
		"",
		data.NewField("time", nil, []time.Time{t}),
		data.NewField(tag, nil, []*float64{nil}),
	)
}

func NoDataFrame(t time.Time) *data.Frame {
	return data.NewFrame(
		"",
		data.NewField("time", nil, []time.Time{t}),
	)
}

func validateResponse(ddsResponse *dds.Response) error {
	logger := log.Logger.With("func", "Build")

	reportsNum := len(ddsResponse.Reports)
	if reportsNum == 0 {
		return errors.New("no reports in DDS response")
	} else if reportsNum > 1 {
		return fmt.Errorf("too many reports (%d) in DDS response", reportsNum)
	}
	report := ddsResponse.Reports[0]
	if message := report.Message; message != nil {
		if _, ok := dds.AcceptableMessages[message.Id]; !ok {
			return message
		} else {
			logger.Debug(message.Error())
		}
	}
	if report.TimeData == nil {
		return errors.New("no time data in DDS response")
	}
	if report.Metric == nil {
		return errors.New("no metric data in DDS response")
	}
	if _, ok := dds.SupportedFormats[report.Metric.Format]; !ok {
		return fmt.Errorf("unsupported data format (%s) in DDS response", report.Metric.Format)
	}
	return nil
}

func Build(ddsResponse *dds.Response, headers *dds.HeaderMap, wide bool) (*data.Frame, error) {
	err := validateResponse(ddsResponse)
	if err != nil {
		return nil, err
	}
	report := ddsResponse.Reports[0]

	format := report.Metric.Format
	frameName := strings.Trim(report.Metric.Description, " ")
	var newFrame *data.Frame

	if format == dds.ReportFormat {
		newFrame = buildForReport(&report, headers, frameName)
	} else if wide {
		return buildWideForMetric(&report, frameName), nil
	} else {
		return buildLongForMetric(&report, frameName), nil
	}
	return newFrame, nil
}

func Build2(ddsResponse *dds.Response, headers *dds.HeaderMap) ([]*data.Frame, error) {
	err := validateResponse(ddsResponse)
	if err != nil {
		return nil, err
	}
	report := ddsResponse.Reports[0]
	return buildForReport2(&report, headers), nil
}

// buildWideForMetric creates a time series data frame for a metric from pre-parsed DDS response.
// Grafana frame format: wide.
func buildWideForMetric(report *dds.Report, frameName string) *data.Frame {
	timestamp := report.TimeData.UTCEnd.Time
	metricFormat := report.Metric.Format
	labels := getFrameLabels(metricFormat, frameName)
	resultFrame := data.NewFrame(frameName, data.NewField("time", nil, []time.Time{timestamp}))

	iterateMetricRows(report, frameName,
		func(name string, value *float64) {
			newField := data.NewField(name, labels, []*float64{value})
			resultFrame.Fields = append(resultFrame.Fields, newField)
		})

	// Built-in alerting requires either a frame in wide format with mandatory numeric fields,
	// or an empty one. However, empty frames won't work for time series views.
	// Solution for single type metric is to send nil values if there's no data.
	// For list type metrics, we don't have column names to do the same; it has to be fixed differently.
	if len(resultFrame.Fields) == 1 && metricFormat == "single" {
		newField := data.NewField(frameName, labels, []*float64{nil})
		resultFrame.Fields = append(resultFrame.Fields, newField)
	}

	return resultFrame
}

// buildLongForMetric creates a non time series data frame for a metric from pre-parsed DDS response.
// Grafana frame format: long.
func buildLongForMetric(report *dds.Report, frameName string) *data.Frame {
	metricFormat := report.Metric.Format
	nameField := "metric"
	timestamp := report.TimeData.UTCEnd.Time
	valField := frameName
	if metricFormat == "list" {
		valField, nameField = splitQueryName(frameName)
		if nameField == "" {
			nameField = frameName
			valField = "value"
		}
	}

	resultFrame := data.NewFrame("",
		data.NewField("time", nil, []time.Time{}),
		data.NewField(nameField, nil, []string{}),
		data.NewField(valField, nil, []*float64{}),
	)

	iterateMetricRows(report, frameName,
		func(name string, value *float64) {
			resultFrame.Fields[0].Append(timestamp)
			resultFrame.Fields[1].Append(name)
			resultFrame.Fields[2].Append(value)
		})

	return resultFrame
}

// iterateMetricRows parses metric key-value pairs and passes them to `process` while iterating over rows.
func iterateMetricRows(report *dds.Report, defaultName string, process func(name string, value *float64)) {
	colMap := map[string]bool{}
	var sb strings.Builder
	for _, jsonRow := range report.Rows {
		cols := jsonRow.Cols
		name, rawValue := cols[0], cols[1]
		if name == "*NoData*" {
			continue
		}
		sb.Reset()
		sb.WriteString(name)
		if len(jsonRow.Cols) >= 3 {
			sb.WriteString("[")
			sb.WriteString(cols[2])
			if len(jsonRow.Cols) >= 4 && cols[3] != "" {
				sb.WriteString(".")
				sb.WriteString(cols[3])
			}
			sb.WriteString("]")
		}
		if sb.Len() == 0 {
			sb.WriteString(defaultName)
		}
		colName := sb.String()
		if _, ok := colMap[colName]; ok {
			continue
		}
		colMap[colName] = true
		process(colName, parseFloat(rawValue))
	}
}

func buildForReport(report *dds.Report, headers *dds.HeaderMap, frameName string) *data.Frame {
	logger := log.Logger.With("func", "buildForReport")
	frame := data.NewFrame(frameName)
	reportName := report.Metric.Id

	for i, col := range report.Headers.Cols {
		header := headers.Get(reportName, col.Id)
		var field *data.Field
		// The first value is a dummy to fit in captions and header.
		if col.Type == dds.NumericColType {
			field = data.NewField(header, nil, []*float64{nil})
		} else {
			if col.Type != dds.TextColType {
				logger.Warn("unsupported column type, considering as string", "type", col.Type)
			}
			field = data.NewField(header, nil, []string{""})
		}
		for _, row := range report.Rows {
			rawValue := row.Cols[i]
			if col.Type == dds.NumericColType {
				field.Append(parseFloat(rawValue))
			} else {
				field.Append(rawValue)
			}
		}
		frame.Fields = append(frame.Fields, field)
	}

	buildField := func(name string, prefix string, value string) *data.Field {
		// All the frames must have the same number of rows.
		values := make([]*string, 1+len(report.Rows))
		values[0] = &value
		field := data.NewField(prefix+name, nil, values)
		return field
	}

	timeData := report.TimeData

	field := buildField("Samples", BannerPrefix, strconv.Itoa(timeData.NumSamples))
	frame.Fields = append(frame.Fields, field)
	if timeData.NumSystems != nil {
		field = buildField("Systems", BannerPrefix, strconv.Itoa(*timeData.NumSystems))
		frame.Fields = append(frame.Fields, field)
	}
	field = buildField("Time range", BannerPrefix,
		fmt.Sprintf("%s - %s",
			timeData.LocalStart.Format(ReportDateFormat),
			timeData.LocalEnd.Format(ReportDateFormat)))
	frame.Fields = append(frame.Fields, field)
	field = buildField("Interval", BannerPrefix,
		timeData.LocalEnd.Sub(timeData.LocalStart.Time).String())
	frame.Fields = append(frame.Fields, field)

	for _, caption := range report.Caption.Vars {
		name := headers.Get(reportName, caption.Name)
		field := buildField(name, CaptionPrefix, caption.Value)
		frame.Fields = append(frame.Fields, field)
	}

	return frame
}

func buildForReport2(report *dds.Report, headers *dds.HeaderMap) []*data.Frame {
	reportName := report.Metric.Id
	captionFrame := data.NewFrame(reportName + "_CAP")
	intervalFrame := data.NewFrame(reportName + "_INT")
	reportFrame := data.NewFrame(reportName)

	for i, col := range report.Headers.Cols {
		header := headers.Get(reportName, col.Id)
		var field *data.Field = nil
		for _, row := range report.Rows {
			rawValue := row.Cols[i]
			if field == nil {
				if col.Type == dds.NumericColType {
					field = data.NewField(header, nil, []*float64{parseFloat(rawValue)})
				} else {
					field = data.NewField(header, nil, []string{rawValue})
				}
				continue
			}
			if col.Type == dds.NumericColType {
				field.Append(parseFloat(rawValue))
			} else {
				field.Append(rawValue)
			}
		}
		reportFrame.Fields = append(reportFrame.Fields, field)
	}

	buildField := func(name string, prefix string, value string) *data.Field {
		// All the frames must have the same number of rows.
		values := make([]*string, 1+len(report.Rows))
		values[0] = &value
		field := data.NewField(prefix+name, nil, values)
		return field
	}

	timeData := report.TimeData

	field := buildField("Samples", BannerPrefix, strconv.Itoa(timeData.NumSamples))
	intervalFrame.Fields = append(intervalFrame.Fields, field)
	if timeData.NumSystems != nil {
		field = buildField("Systems", BannerPrefix, strconv.Itoa(*timeData.NumSystems))
		intervalFrame.Fields = append(intervalFrame.Fields, field)
	}
	field = buildField("Time range", BannerPrefix,
		fmt.Sprintf("%s - %s",
			timeData.LocalStart.Format(ReportDateFormat),
			timeData.LocalEnd.Format(ReportDateFormat)))
	intervalFrame.Fields = append(intervalFrame.Fields, field)
	field = buildField("Interval", BannerPrefix,
		timeData.LocalEnd.Sub(timeData.LocalStart.Time).String())
	intervalFrame.Fields = append(intervalFrame.Fields, field)

	for _, caption := range report.Caption.Vars {
		name := headers.Get(reportName, caption.Name)
		field := buildField(name, CaptionPrefix, caption.Value)
		captionFrame.Fields = append(captionFrame.Fields, field)
	}

	result := make([]*data.Frame, 3)
	result[0] = reportFrame
	result[1] = intervalFrame
	result[2] = captionFrame
	return result
}

func parseFloat(value string) *float64 {
	// Value can contain different kinds of text meaning n/a: NaN, blank value, Deact, etc.
	if parsed, err := strconv.ParseFloat(value, 64); err == nil && !math.IsNaN(parsed) {
		return &parsed
	}
	return nil
}

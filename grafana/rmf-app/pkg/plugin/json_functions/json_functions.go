/**
* (C) Copyright IBM Corp. 2023.
* (C) Copyright Rocket Software, Inc. 2023.
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

package json_functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	ffns "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame_functions"
	httphlpr "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/http_helper"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	xmlf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/xml_functions"

	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/tidwall/gjson"
)

var _headerXslMap map[string][]typ.ColHeaderXslMap = nil

func GetJsonObject(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetJsonPropertyValue(jsonStr string, propertyPath string) interface{} {
	jval := gjson.Get(jsonStr, propertyPath)
	return jval.Value()
}

func GetJsonPropertyValueAsNumber(jsonStr string, propertyPath string) float64 {
	return gjson.Get(jsonStr, propertyPath).Num
}

func GetDataFormat(jsonStr string) (string, error) {
	dataFormat := gjson.Get(jsonStr, "report.0.metric.format").String()
	if dataFormat == "" {
		return "", fmt.Errorf("could not get data format in GetDataFormat()")
	}
	return dataFormat, nil
}

func GetMessageInResponse(jsonStr string) *typ.DDSMessage {
	var message *typ.DDSMessage
	messageNode := GetJsonPropertyValue(jsonStr, "report.0.message")
	if messageNode != nil {
		messageItem := messageNode.(map[string]interface{})
		if messageItem != nil {
			message = new(typ.DDSMessage)
			if messageItem["id"] != nil {
				message.Id = messageItem["id"].(string)
			}
			if messageItem["severity"] != nil {
				message.Severity = int(messageItem["severity"].(float64))
			}
			if messageItem["description"] != nil {
				message.Description = messageItem["description"].(string)
			}
		}
	}
	return message
}

func FetchIntervalAndOffset(jsonStr string) (typ.IntervalOffset, error) {
	var (
		resultTimeData typ.IntervalOffset
		err            error
	)

	localStartTime, err := time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.localStart").(string))
	if err != nil {
		return resultTimeData, err
	}

	UTCStartTime, err := time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.utcStart").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.ServerTimezoneOffset = localStartTime.Sub(UTCStartTime)
	resultTimeData.ServiceCallInterval = float64(GetJsonPropertyValueAsNumber(jsonStr, "report.0.timeData.gathererInterval.value"))

	return resultTimeData, nil
}

func FetchServerTimeConfig(jsonStr string) (typ.DDSTimeData, error) {
	var (
		resultTimeData typ.DDSTimeData
		err            error
	)

	resultTimeData.LocalStartTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.localStart").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.LocalEndTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.localEnd").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.LocalPrevTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.localPrev").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.LocalNextTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.localNext").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.UTCStartTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.utcStart").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.UTCEndTime, err = time.Parse("20060102150405", GetJsonPropertyValue(jsonStr, "report.0.timeData.utcEnd").(string))
	if err != nil {
		return resultTimeData, err
	}

	resultTimeData.ServerTimezoneOffset = resultTimeData.LocalStartTime.Sub(resultTimeData.UTCStartTime)

	resultTimeData.ServiceCallInterval = float64(GetJsonPropertyValueAsNumber(jsonStr, "report.0.timeData.gathererInterval.value"))

	// Convert all to UTC times.
	resultTimeData.LocalStartTime = resultTimeData.LocalStartTime.Add(-1 * resultTimeData.ServerTimezoneOffset)
	resultTimeData.LocalEndTime = resultTimeData.LocalEndTime.Add(-1 * resultTimeData.ServerTimezoneOffset)
	resultTimeData.LocalPrevTime = resultTimeData.LocalPrevTime.Add(-1 * resultTimeData.ServerTimezoneOffset)
	resultTimeData.LocalNextTime = resultTimeData.LocalNextTime.Add(-1 * resultTimeData.ServerTimezoneOffset)
	return resultTimeData, nil
}

// MetricFrameFromJson parses JSON string and create a data frame either for time series or a regular one.
func MetricFrameFromJson(jsonStr string, query *typ.QueryModel, isTimeSeries bool) (*data.Frame, error) {
	var ddsResponse typ.DDSResponse
	if err := json.Unmarshal([]byte(jsonStr), &ddsResponse); err != nil {
		return nil, fmt.Errorf("could not parse JSON in MetricFrameFromJson(): Error = %v", err)
	}
	if len(ddsResponse.Reports) == 0 {
		return nil, errors.New("unexpected data in MetricFrameFromJson(): Error = no report sections")
	}

	diff := query.ServerTimeData.LocalEndTime.Sub(query.ServerTimeData.LocalStartTime) / 2
	frameTimestamp := query.ServerTimeData.LocalStartTime.Add(diff)
	queryName := ffns.GetFrameName(query)

	if isTimeSeries {
		return ConstructMetricTSFrame(&ddsResponse, queryName, frameTimestamp), nil
	} else {
		return ConstructMetricFrame(&ddsResponse, queryName, frameTimestamp), nil
	}
}

// ConstructMetricTSFrame creates a timeseries data frame for a metric from pre-parsed DDS response.
func ConstructMetricTSFrame(ddsResponse *typ.DDSResponse, queryName string, timestamp time.Time) *data.Frame {
	frameName := ""
	if ddsResponse.Reports[0].Metric.Format == "list" {
		frameName = queryName
	}
	resultFrame := data.NewFrame(frameName, data.NewField("timestamp", nil, []time.Time{timestamp}))

	IterateMetricRows(ddsResponse, queryName,
		func(name string, value *float64) {
			newField := data.NewField(name, nil, []*float64{value})
			resultFrame.Fields = append(resultFrame.Fields, newField)
		})

	return resultFrame
}

// ConstructMetricFrame creates a non-timeseries data frame for a metric from pre-parsed DDS response.
func ConstructMetricFrame(ddsResponse *typ.DDSResponse, queryName string, timestamp time.Time) *data.Frame {
	metricFormat := ddsResponse.Reports[0].Metric.Format
	nameField := "name"
	valField := ""
	if metricFormat == "list" {
		nameField = queryName
		valField = "value"
		if queryName != "" {
			splitStringSlice := strings.SplitAfter(queryName, "by")
			if len(splitStringSlice) > 1 {
				valField = strings.TrimSpace(splitStringSlice[0])
				valField = strings.TrimRight(valField, "by")
				valField = strings.TrimSpace(valField)
				nameField = strings.TrimSpace(splitStringSlice[1])
			}
		}
	}

	resultFrame := data.NewFrame("",
		data.NewField("timestamp", nil, []time.Time{}),
		data.NewField(nameField, nil, []string{}),
		data.NewField(valField, nil, []*float64{}),
	)

	IterateMetricRows(ddsResponse, queryName,
		func(name string, value *float64) {
			resultFrame.Fields[0].Append(timestamp)
			resultFrame.Fields[1].Append(name)
			resultFrame.Fields[2].Append(value)
			if metricFormat == "single" {
				resultFrame.Fields[2].Name = name
			}
		})

	return resultFrame
}

// IterateMetricRows parses metric key-value pairs and passes them to `process` while iterating over rows.
func IterateMetricRows(ddsResponse *typ.DDSResponse, defaultName string, process func(name string, value *float64)) {
	for _, jsonRow := range ddsResponse.Reports[0].Rows {
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

func ConstructMetadataFromJson(jsonStr string) map[string]interface{} {
	metaData := GetJsonPropertyValue(jsonStr, "report.0.timeData").(map[string]interface{})
	return metaData
}

func ConstructReportFrameFromJson(jsonStr string,
	queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {

	// writeToFile(jsonStr)
	resultFrame := data.NewFrame(ffns.GetFrameName(queryModel))

	// The below function will
	headerInfoList, err := getUpdatedHeaderInfoList(jsonStr, endpointModel, queryModel.SelectedQuery)
	if err != nil {
		return resultFrame, fmt.Errorf("could not get headerInfoList in ConstructReportFrameFromJson(). Error=%v", err)
	}

	// Add the regular metrics values as fields (either float64 or string values)
	var trackingIndexCol, trackingIndexRow int

	// Get a reference to columnHeaders Json property
	colHeaders := GetJsonPropertyValue(jsonStr, "report.0.columnHeaders").(map[string]interface{})["col"].([]interface{})

	for colIndex := 0; colIndex < len(colHeaders); colIndex++ {
		var stringSlice []string
		var floatSlice []float64

		// Stop gap solution. To be removed.
		headerInfoListDict := headerInfoList[colIndex].(map[string]interface{})
		if headerInfoListDict["HeaderID"] == "CASPTYPM" {
			headerInfoListDict["Type"] = "T"
		}

		rows := GetJsonPropertyValue(jsonStr, "report.0.row").([]interface{})

		// Some columns have type="N" (meaning Numeric) but contain "" values. In those cases it will be considered string.
		treatColumnAsString := shouldColumnBeTreatedAsString(rows, colIndex)

		for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
			currentValue := rows[rowIndex].(map[string]interface{})["col"].([]interface{})[colIndex].(string)
			if treatColumnAsString {
				stringSlice = append(stringSlice, currentValue)
			} else {
				f2, err := strconv.ParseFloat(currentValue, 64)
				if err != nil {
					return nil, fmt.Errorf("could not convert value to float in ConstructReportFrameFromJson(). Error=%v", err)
				}
				floatSlice = append(floatSlice, f2)
			}
			trackingIndexRow = rowIndex
		}
		if treatColumnAsString {
			newField := data.NewField(headerInfoListDict["HeaderText"].(string), nil, stringSlice)
			resultFrame.Fields = append(resultFrame.Fields, newField)
		} else {
			newField := data.NewField(headerInfoListDict["HeaderText"].(string), nil, floatSlice)
			resultFrame.Fields = append(resultFrame.Fields, newField)
		}
		trackingIndexCol = colIndex
	}

	trackingIndexCol++
	trackingIndexRow++

	var varList []interface{}
	captionNode := GetJsonPropertyValue(jsonStr, "report.0.caption")
	if captionNode != nil {
		varList = captionNode.((map[string]interface{}))["var"].([]interface{})
	}

	// Add the captions (these go as headers in the header table for Reports, not to be confused with col headers of reports)
	for _, caption := range varList {
		fieldValue := ""
		headerInfoListDict := headerInfoList[trackingIndexCol].(map[string]interface{})
		fieldHeaderText := "Header::" + headerInfoListDict["HeaderText"].(string)
		val, ok := caption.(map[string]interface{})["value"]
		if ok {
			fieldValue = val.(string)
		}
		fieldValues := make([]string, trackingIndexRow) // This slice will be having 'trackingIndexRow' rows
		fieldValues[0] = fieldValue                     // Set only the 0th value
		// Create a new field
		newField := data.NewField(fieldHeaderText, nil, fieldValues)
		resultFrame.Fields = append(resultFrame.Fields, newField)
		trackingIndexCol++
	}
	return resultFrame, nil
}

func getUpdatedHeaderInfoList(jsonStr string, em *typ.DatasourceEndpointModel, selectedQuery string) ([]interface{}, error) {
	var xmlFns xmlf.XmlFunctions
	var httpHlpr httphlpr.HttpHelper
	var colHeaderInfoList []interface{}

	// Form the headerXslMap that contains the headerText (when lookedup with headerId)
	if _headerXslMap == nil {
		xslUrl := httpHlpr.GetHttpUrlForReportXsl(em)
		xslFileContent, err := httpHlpr.GetXslFileContents(xslUrl, em)
		if err != nil {
			return nil, fmt.Errorf("could not get xslFileContent in getUpdatedHeaderInfoList. Error=%v", err)
		}
		headerXslMap, err := xmlFns.GetColHeaderXslMap(xslFileContent)
		if err != nil {
			return nil, fmt.Errorf("could not unmarshal jsonStr data in getUpdatedHeaderInfoList. Error=%v", err)
		}
		_headerXslMap = headerXslMap
	}

	// Get a reference to columnHeaders Json property
	colHeaders := GetJsonPropertyValue(jsonStr, "report.0.columnHeaders").((map[string]interface{}))["col"].([]interface{})

	// Fetch and update the column header info
	for _, header := range colHeaders {
		header = updateHeaderTextWithXslMapValue(selectedQuery, header)
		colHeaderInfoList = append(colHeaderInfoList, header)
	}

	// to be commented for prod. for debugging only
	// writeToFile(jsonStr)

	// Get a reference to caption(s) Json property
	var varList []interface{}
	captionNode := GetJsonPropertyValue(jsonStr, "report.0.caption")
	if captionNode != nil {
		varList = captionNode.((map[string]interface{}))["var"].([]interface{})
	}

	// Fetch and update the caption info
	for indx := 0; indx < len(varList); indx++ {
		captionInfo := updateHeaderTextWithXslMapValue(selectedQuery, varList[indx])
		colHeaderInfoList = append(colHeaderInfoList, captionInfo)
	}
	return colHeaderInfoList, nil
}

func updateHeaderTextWithXslMapValue(selectedQuery string, colHeaderInfo interface{}) interface{} {
	reportName := "NA"
	if strings.Contains(strings.Trim(strings.ToUpper(selectedQuery), ""), "REPORT.STORCR") {
		reportName = "STORCR"
	} else if strings.Contains(strings.Trim(strings.ToUpper(selectedQuery), ""), "REPORT.STORC") {
		reportName = "STORC"
	}
	resultMap := _headerXslMap[reportName]

	var updatedColHeaderInfo = make(map[string]interface{})
	if colHeaderInfo.(map[string]interface{})["name"] != nil {
		updatedColHeaderInfo["HeaderID"] = colHeaderInfo.(map[string]interface{})["name"]
		updatedColHeaderInfo["HeaderText"] = colHeaderInfo.(map[string]interface{})["name"]
		updatedColHeaderInfo["Value"] = colHeaderInfo.(map[string]interface{})["value"]
		updatedColHeaderInfo["Type"] = "T"
	} else {
		updatedColHeaderInfo["HeaderID"] = colHeaderInfo.(map[string]interface{})["value"]
		updatedColHeaderInfo["HeaderText"] = colHeaderInfo.(map[string]interface{})["value"]
		updatedColHeaderInfo["Value"] = colHeaderInfo.(map[string]interface{})["value"]
		updatedColHeaderInfo["Type"] = colHeaderInfo.(map[string]interface{})["type"]
	}
	for _, vl := range resultMap {
		if vl.HeaderID == updatedColHeaderInfo["HeaderID"] {
			updatedColHeaderInfo["HeaderText"] = vl.HeaderText
			break
		}
	}
	return updatedColHeaderInfo
}

func shouldColumnBeTreatedAsString(rows []interface{}, colIndex int) bool {
	treatAsString := false
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		currentValue := rows[rowIndex].(map[string]interface{})["col"].([]interface{})[colIndex].(string)
		_, err := strconv.ParseFloat(currentValue, 64)
		if err != nil {
			treatAsString = true
			break
		}
	}
	return treatAsString
}

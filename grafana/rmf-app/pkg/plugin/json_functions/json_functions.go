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
	resultFormat := gjson.Get(jsonStr, "report.0.metric.format").String()
	// result formats can be single, report, report_single, list, list_single
	if resultFormat != "" {
		if resultFormat == "single" {
			return resultFormat, nil
		} else if resultFormat == "report" || resultFormat == "list" {
			rowElementCount := GetJsonPropertyValue(jsonStr, "report.0.row.#")
			if rowElementCount == 1 {
				if resultFormat == "report" {
					resultFormat = "report_single"
				} else {
					resultFormat = "list_single"
				}
			}
		}
	} else {
		return "", fmt.Errorf("could not get data format in GetDataFormat()")
	}
	return resultFormat, nil
}

func GetErrorInResponse(jsonStr string) string {
	var returnResult string
	messageNode := GetJsonPropertyValue(jsonStr, "report.0.message")
	if messageNode != nil {
		messageItem := messageNode.(map[string]interface{})
		if messageItem != nil {
			if messageItem["id"] != nil {
				returnResult = returnResult + messageItem["id"].(string)
				if messageItem["severity"] != nil {
					returnResult = returnResult + " (Sev: " + fmt.Sprintf("%d", int(messageItem["severity"].(float64))) + ") "
					if messageItem["description"] != nil {
						returnResult = returnResult + messageItem["description"].(string)
					}
				}
			}
		}
	}
	return returnResult
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

func ConstructSingleValueFrameFromJson(jsonStr string,
	queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	var timestampList []time.Time
	var fieldList []string
	var valueList []float64
	var valueStrList []string

	resultFrame := data.NewFrame("")

	cols := GetJsonPropertyValue(jsonStr, "report.0.row").([]interface{})
	for colIndex := 0; colIndex < len(cols); colIndex++ {
		// var newField *data.Field
		currentCols := cols[colIndex].(map[string]interface{})["col"].([]interface{})
		fieldName := currentCols[0].(string)
		fieldValue := currentCols[1].(string)
		// fieldLabels := map[string]string{"name": fieldValue}
		fieldValueN, err := strconv.ParseFloat(fieldValue, 64)

		if len(fieldName) < 1 {
			fieldName = ffns.GetFrameName(queryModel)
		}

		timestampList = append(timestampList, queryModel.TimeRangeFrom)
		newTimeField := data.NewField("timestamp", nil, timestampList)
		resultFrame.Fields = append(resultFrame.Fields, newTimeField)

		fieldList = append(fieldList, fieldName)
		newNameField := data.NewField("Name", nil, fieldList)
		resultFrame.Fields = append(resultFrame.Fields, newNameField)

		if err != nil {
			fieldValues := []string{fieldValue}
			// newField = data.NewField(fieldName, fieldLabels, fieldValues)
			valueStrList = append(valueStrList, fieldValues[0])
			newValField := data.NewField(fieldName, nil, valueStrList)
			resultFrame.Fields = append(resultFrame.Fields, newValField)
		} else {
			fieldValues := []float64{fieldValueN}
			// newField = data.NewField(fieldName, fieldLabels, fieldValues)
			valueList = append(valueList, fieldValues[0])
			newValField := data.NewField(fieldName, nil, valueList)
			resultFrame.Fields = append(resultFrame.Fields, newValField)
		}

		// resultFrame.Fields = append(resultFrame.Fields, newField)
	}
	return resultFrame, nil
}

func ConstructTimeSeriesSingleValueFrameFromJson(jsonStr string,
	queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	var plottingTimeStamp time.Time // Stores the timestamp that is plotted in the timeline plugin.

	resultFrame := data.NewFrame("")

	// For relative timeseries (forward plotting), the plotting timestamp is the mid-time point of local start and end times
	if !queryModel.AbsoluteTimeSelected {
		diff := (queryModel.ServerTimeData.LocalEndTime.Sub(queryModel.ServerTimeData.LocalStartTime)) / 2
		plottingTimeStamp = queryModel.ServerTimeData.LocalStartTime.Add(diff)
	} else { // For absolute  the plotting timestamp is the TimeSeriesTimeRangeFrom timestamp.
		plottingTimeStamp = queryModel.TimeSeriesTimeRangeFrom
	}

	// Add the plotting timestamp field
	resultFrame.Fields = append(resultFrame.Fields,
		data.NewField("time", nil, []time.Time{plottingTimeStamp}))

	// Add the regular metrics values as fields (either float64 or string values)
	var fieldName string
	rows := GetJsonPropertyValue(jsonStr, "report.0.row").([]interface{})
	cols := rows[0].(map[string]interface{})["col"].([]interface{})
	fieldName = cols[0].(string)
	if fieldName == "" { // For single values fieldName can come blank. Then set it to frame name.
		fieldName = ffns.GetFrameName(queryModel)
	}
	fieldValue := cols[1].(string)
	fieldValueN, err := strconv.ParseFloat(fieldValue, 64)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve float value  in ConstructTimeSeriesSingleValueFrameFromJson(): Error = %v", err)
	} else if fieldValue == "NaN" {
		return nil, fmt.Errorf("could not create frame in ConstructTimeSeriesSingleValueFrameFromJson(). DDS returned NaN value")
	}
	fieldValues := []float64{fieldValueN}
	newField := data.NewField(fieldName, nil, fieldValues)
	resultFrame.Fields = append(resultFrame.Fields, newField)
	return resultFrame, nil
}

func ConstructTimeSeriesListFrameFromJson(jsonStr string,
	queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {

	var plottingTimeStamp time.Time
	resultFrame := data.NewFrame(ffns.GetFrameName(queryModel))

	// For relative timeseries (forward plotting), the plotting timestamp is the mid-time point of local start and end times
	if !queryModel.AbsoluteTimeSelected {
		diff := (queryModel.ServerTimeData.LocalEndTime.Sub(queryModel.ServerTimeData.LocalStartTime)) / 2
		plottingTimeStamp = queryModel.ServerTimeData.LocalStartTime.Add(diff)
	} else { // For absolute  the plotting timestamp is the TimeSeriesTimeRangeFrom timestamp.
		plottingTimeStamp = queryModel.TimeSeriesTimeRangeFrom
	}

	// Add the plotting timestamp field
	resultFrame.Fields = append(resultFrame.Fields,
		data.NewField("time", nil, []time.Time{plottingTimeStamp}))

	// Get the column count from the first row.
	rows := GetJsonPropertyValue(jsonStr, "report.0.row").([]interface{})

	var colCount int
	if len(rows) > 0 {
		colsRef := rows[0].(map[string]interface{})["col"]
		colCount = len(colsRef.([]interface{}))
	}

	// Add the regular metrics values as fields (either float64 or string values)
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		var fieldName string
		colsRef := rows[rowIndex].(map[string]interface{})["col"].([]interface{})
		// Sometimes a third column will actually give the field name with a qualifier
		if colCount == 3 {
			fieldName = colsRef[0].(string) + "[" + colsRef[2].(string) + "]"
		} else {
			fieldName = colsRef[0].(string)
		}

		fieldValue := colsRef[1].(string)
		fieldValueN, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve float value  in ConstructTimeSeriesListFrameFromJson(): Error = %v", err)
		} else if fieldValue == "NaN" {
			return nil, fmt.Errorf("could not create frame in ConstructTimeSeriesListFrameFromJson(). DDS returned NaN value")
		}
		fieldValues := []float64{fieldValueN}
		newField := data.NewField(fieldName, nil, fieldValues)
		resultFrame.Fields = append(resultFrame.Fields, newField)
	}
	return resultFrame, nil
}

func ConstructListFrameFromJson(jsonStr string,
	queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	var timestampList []time.Time
	var fieldList []string
	var valueList []float64

	resultFrame := data.NewFrame("")

	rows := GetJsonPropertyValue(jsonStr, "report.0.row").([]interface{})
	for rowIndex := 0; rowIndex < len(rows); rowIndex++ {
		var fieldName string
		cols := rows[rowIndex].(map[string]interface{})["col"].([]interface{})
		if len(cols) == 3 {
			fieldName = cols[0].(string) + "[" + cols[2].(string) + "]"
		} else {
			fieldName = cols[0].(string)
		}
		fieldValue := cols[1].(string)
		fieldValueN, err := strconv.ParseFloat(fieldValue, 64)
		if err != nil {
			return nil, fmt.Errorf("could not retrieve float value  in ConstructListFrameFromJson(): Error = %v", err)
		} else if fieldValue == "NaN" {
			return nil, fmt.Errorf("could not create frame in ConstructListFrameFromJson(). DDS returned NaN value")
		}
		fieldValues := []float64{fieldValueN}
		fieldList = append(fieldList, fieldName)
		valueList = append(valueList, fieldValues[0])
		timestampList = append(timestampList, queryModel.TimeRangeFrom)
	}

	if len(fieldList) > 0 {
		var queryName string = ffns.GetFrameName(queryModel)
		var nameField string = queryName
		var valField string = "Value"

		if strings.Trim(queryName, "") != "" {
			splitStringSlice := strings.SplitAfter(queryName, "by")
			if len(splitStringSlice) > 1 {
				valField = strings.TrimSpace(splitStringSlice[0])
				valField = strings.TrimRight(valField, "by")
				nameField = strings.TrimSpace(splitStringSlice[1])
			}
		}
		newTimeField := data.NewField("timestamp", nil, timestampList)
		resultFrame.Fields = append(resultFrame.Fields, newTimeField)

		newNameField := data.NewField(nameField, nil, fieldList)
		resultFrame.Fields = append(resultFrame.Fields, newNameField)

		newValField := data.NewField(valField, nil, valueList)
		resultFrame.Fields = append(resultFrame.Fields, newValField)
	}
	return resultFrame, nil
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

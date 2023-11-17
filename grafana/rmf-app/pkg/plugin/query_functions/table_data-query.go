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
package query_functions

import (
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	errh "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/error_handler"
	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	repo "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/repository"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

type TableDataQuery struct {
}

func (t *TableDataQuery) QueryForTableData(
	queryModel *typ.QueryModel,
	endpointModel *typ.DatasourceEndpointModel,
	errHandler *errh.ErrHandler) (*data.Frame, error) {

	var repos repo.Repository
	var jsonf jsonf.JsonFunctions

	// Get xml data response
	responseData, err := repos.ExecuteQueryAndGetResponse(queryModel, endpointModel)
	if err != nil {
		return nil, fmt.Errorf("error after calling ExecuteQueryAndGetResponse in QueryForTableData: responseData=%v & error=%v", responseData, err)
	}

	// Get the Json String
	jsonStr := string(responseData[:])
	if jsonStr == "" || jsonStr == "*No Data*" {
		return nil, fmt.Errorf("response json is blank/no data in QueryForTableData - error=%v", err)
	} else {
		errHandler.LogStatus(fmt.Sprintf("***Executed query: %s with url: %s and got response in QueryForTableData(): %s", queryModel.SelectedQuery, queryModel.Url, jsonStr))
	}

	// Check for any error contained in the response
	errorInResponse := jsonf.GetErrorInResponse(jsonStr)
	if errorInResponse != "" {
		return nil, fmt.Errorf("DDSError - " + errorInResponse)
	}

	// Get the dataformat
	resultDataFormat, err := jsonf.GetDataFormat(jsonStr)
	if err != nil {
		return nil, fmt.Errorf("could not get result dataformat QueryForTableData(): Error=%v", err)
	}

	// GathererInterval is used to wait 'n' secs before streaming a data chunk or while calling service to fetch data.
	// If we invoke the service again within this interval, the results will be returned from cache
	queryModel.ServerTimeData.ServiceCallInterval = float64(jsonf.GetJsonPropertyValueAsNumber(jsonStr, "report.0.timeData.gathererInterval.value"))

	// Compose the newFrame
	// resultDataFormat can be single, report, report_single, list, list_single collectively called table_data
	var newFrame *data.Frame = new(data.Frame)
	if resultDataFormat == "list" { // 'list' (ideal for RMF Chart)
		newFrame, err = jsonf.ConstructListFrameFromJson(jsonStr, queryModel, endpointModel)
	} else if resultDataFormat == "single" || resultDataFormat == "list_single" { // 'single' or 'list_single' (ideal for Gauge)
		newFrame, err = jsonf.ConstructSingleValueFrameFromJson(jsonStr, queryModel, endpointModel)
	} else if resultDataFormat == "report" { // 'report' (for RMF Report)
		newFrame, err = jsonf.ConstructReportFrameFromJson(jsonStr, queryModel, endpointModel)
	}
	if err != nil {
		return nil, fmt.Errorf("could not obtain frame QueryForTableData(): Error=%v", err)
	}
	// Get the meta data fro Json. The meta data info like numSamples is displayed on top of the header
	newFrame.Meta = &data.FrameMeta{
		Custom: jsonf.ConstructMetadataFromJson(jsonStr),
	}
	return newFrame, nil
}

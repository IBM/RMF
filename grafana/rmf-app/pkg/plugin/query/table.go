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

package query

import (
	"fmt"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	repo "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/repository"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

func GetTableFrame(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	logger := log.Logger.With("func", "GetTableFrame")
	var repos repo.Repository

	responseData, err := repos.ExecuteQueryAndGetResponse(queryModel, endpointModel)
	if err != nil {
		return nil, fmt.Errorf("error after calling ExecuteQueryAndGetResponse in QueryForTableData: responseData=%v & error=%w", responseData, err)
	}

	// Get the Json String
	jsonStr := string(responseData[:])
	if jsonStr == "" || jsonStr == "*No Data*" {
		return nil, fmt.Errorf("response json is blank/no data in QueryForTableData - Error=%w", err)
	} else {
		logger.Debug("executed query", "query", queryModel.SelectedQuery, "url", queryModel.Url)
	}

	// Check for any error contained in the response
	message := jsonf.GetMessageInResponse(jsonStr)
	if message != nil {
		return nil, fmt.Errorf("DDSError - %s", message)
	}

	// Get the data format
	resultDataFormat, err := jsonf.GetDataFormat(jsonStr)
	if err != nil {
		return nil, fmt.Errorf("could not get result dataformat QueryForTableData(): Error=%w", err)
	}

	queryModel.ServerTimeData, err = jsonf.FetchServerTimeConfig(jsonStr)
	if err != nil {
		return nil, fmt.Errorf("could not get ServerTimeData in QueryForTableData(): Error=%w", err)
	}

	// GathererInterval is used to wait 'n' secs before streaming a data chunk or while calling service to fetch data.
	// If we invoke the service again within this interval, the results will be returned from cache
	if queryModel.ServerTimeData, err = jsonf.FetchServerTimeConfig(jsonStr); err != nil {
		return nil, fmt.Errorf("could not get DDS time data in QueryForTableData(): Error=%w", err)
	}

	// Compose the newFrame
	// Expected data format is report, list, or single.
	var newFrame *data.Frame
	if resultDataFormat == "report" {
		newFrame, err = jsonf.ConstructReportFrameFromJson(jsonStr, queryModel, endpointModel)
	} else {
		newFrame, err = jsonf.MetricFrameFromJson(jsonStr, queryModel, false)
	}
	if err != nil {
		return nil, fmt.Errorf("could not obtain frame QueryForTableData(): Error=%w", err)
	}
	// Get the metadata from Json. The metadata info like numSamples is displayed on top of the header
	newFrame.Meta = &data.FrameMeta{
		Custom: jsonf.ConstructMetadataFromJson(jsonStr),
	}
	return newFrame, nil
}

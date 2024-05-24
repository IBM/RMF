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

package query_functions

import (
	"errors"
	"fmt"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/data"

	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	repo "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/repository"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

type TimeSeriesQuery struct {
}

var AcceptableMessages = map[string]bool{
	"GPM0501I": true, // Data may be inconsistent due to a change of the WLM service policy or IPS
	"GPM0502I": true, // DDS has returned partial data because time gaps have been detected
	"GPM0504I": true, // The data retrieved by DDS may be inconsistent due to a change of the WLM service policy
	"GPM0505I": true, // The data retrieved by DDS may be inconsistent due to a change of the RMF cycle time
	"GPM0506I": true, // IPL/REBOOT detected in the specified time range
	"GPM0507I": true, // DDS could not retrieve valid data for the specified date and time range
}

func (t *TimeSeriesQuery) QueryForTimeseriesDataFrame(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	logger := log.Logger.With("func", "QueryForTimeseriesDataFrame")
	var repos repo.Repository

	// Fetch data from server
	responseData, err := repos.ExecuteQueryAndGetResponse(queryModel, endpointModel)
	if err != nil {
		return nil, err
	}

	// Convert the Xml data to Json
	jsonStr := string(responseData[:])
	if jsonStr == "" || jsonStr == "*No Data*" {
		return nil, errors.New("response json is blank/no data in QueryForTimeseriesDataFrame")
	} else {
		logger.Debug("executed query", "query", queryModel.SelectedQuery, "url", queryModel.Url)
	}

	// Check for any error contained in the response
	message := jsonf.GetMessageInResponse(jsonStr)
	if message != nil {
		_, ok := AcceptableMessages[message.Id]
		if !ok {
			return nil, fmt.Errorf("DDS error: %s", message)
		} else {
			logger.Info(fmt.Sprintf("DDS error: %s", message))
		}
	}

	// GathererInterval is used to wait 'n' secs before streaming a data chunk or while calling service to fetch data.
	// If we invoke the service again within this interval, the results will be returned from cache
	queryModel.ServerTimeData, err = jsonf.FetchServerTimeConfig(jsonStr) // float64(jsonf.GetJsonPropertyValueAsNumber(jsonStr, "report.0.timeData.gathererInterval.value"))
	if err != nil {
		return nil, fmt.Errorf("could not get ServerTimeData in QueryForTimeseriesDataFrame(): Error=%w", err)
	}

	newFrame, err := jsonf.MetricFrameFromJson(jsonStr, queryModel, true)
	if err != nil {
		return nil, fmt.Errorf("could not obtain frame in QueryForTimeseriesDataFrame(): Error=%w", err)
	}
	return newFrame, nil
}

func (t *TimeSeriesQuery) SetTimeRange(queryModel *typ.QueryModel, plotAbsoluteReverse ...bool) {
	var plotReverse bool
	if len(plotAbsoluteReverse) > 0 {
		if plotAbsoluteReverse[0] {
			plotReverse = true
		}
	}

	// Set the Query Model's TimeSeriesTimeRangeFrom and TimeSeriesTimeRangeTo properties
	if queryModel.AbsoluteTimeSelected { // Absolute time
		if queryModel.ServerTimeData.ServiceCallInterval == 0 || queryModel.TimeSeriesTimeRangeFrom.IsZero() {
			fromTime := queryModel.TimeRangeFrom
			queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo = fromTime, fromTime
		} else {
			if plotReverse {
				localPrevTime := queryModel.ServerTimeData.LocalPrevTime
				queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo = localPrevTime, localPrevTime
			} else {
				addedTime := queryModel.TimeSeriesTimeRangeFrom.Add(time.Duration(time.Second * time.Duration(queryModel.ServerTimeData.ServiceCallInterval)))
				queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo = addedTime, addedTime
			}
		}
	} else { // Relative time
		if queryModel.ServerTimeData.ServiceCallInterval == 0 || queryModel.TimeSeriesTimeRangeTo.IsZero() {
			toTime := queryModel.TimeRangeTo
			queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo = toTime, toTime
		} else {
			localNextTime := queryModel.ServerTimeData.LocalNextTime
			queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo = localNextTime, localNextTime
		}
	}
}

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
	"errors"
	"fmt"
	"io"
	"strings"

	httphlper "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/http_helper"
	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

func FetchRootInfo(em *typ.DatasourceEndpointModel) (bool, error) {
	logger := log.Logger.With("func", "FetchRootInfo")
	var httphlpr httphlper.HttpHelper
	url := httphlpr.GetHttpUrlForRoot(em)
	responseData, err := httphlpr.ExecuteHttpGet(url, em)
	if err != nil {
		return false, fmt.Errorf("could not fetch health check configuration in FetchRootInfo() - error=%w", err)
	}
	defer responseData.Body.Close()
	if responseData.StatusCode == 401 { // Unauthorized
		return false, errors.New("authentication failed in FetchRootInfo(). possible invalid username/password combination")
	}
	response, err := io.ReadAll(responseData.Body)
	if err != nil {
		return false, fmt.Errorf("could in io.ReadAll() in FetchRootInfo() - error=%w", err)
	}

	jsonStr := string(response[:])
	logger.Debug("fetched root info and got response", "url", url, "response", jsonStr)
	if jsonStr == "" || jsonStr == "*No Data*" {
		return false, fmt.Errorf("response json is blank/no data in FetchRootInfo - error=%w", err)
	} else {
		jsonObj, err := jsonf.GetJsonObject(jsonStr)
		if err != nil {
			if strings.Contains(jsonStr, "<?xml version=\"1.0\" encoding=\"UTF-8\"?") {
				return false, typ.NewValueError(100, errors.New("unsupported version of DDS"))
			} else {
				return false, fmt.Errorf("response not valid in FetchRootInfo - error=%w", err)
			}
		}
		return len(jsonObj) > 0, nil
	}
}

func FetchMetricsFromIndex(em *typ.DatasourceEndpointModel) ([]byte, error) {
	var httphlpr httphlper.HttpHelper
	url := httphlpr.GetHttpUrlForIndex(em)
	responseData, err := httphlpr.ExecuteHttpGet(url, em)
	if err != nil {
		return nil, fmt.Errorf("could not fetch metrics from index, url= %s in FetchMetricsFromIndex() - error= %w", url, err)
	}
	defer responseData.Body.Close()
	response, err := io.ReadAll(responseData.Body)
	if err != nil {
		return nil, fmt.Errorf("error in io.ReadAll() in FetchMetricsFromIndex() - Error=%w", err)
	}
	return response, nil
}

func FetchIntervalAndOffset(em *typ.DatasourceEndpointModel) (typ.IntervalTimeData, error) {
	// Fetch the Http Url from endpoint model
	// <proto>://<host>:<port>/gpm/perform.xml?resource=,,SYSPLEX&id=8D0D50
	var (
		httphlpr       httphlper.HttpHelper
		resultTimeData typ.IntervalTimeData
	)
	url := httphlpr.GetHttpUrlForTimezoneOffset(em)

	// Execute the request and fetch the response
	response, err := httphlpr.ExecuteHttpGet(url, em)
	if err != nil {
		return resultTimeData, err
	}

	// Close the response body once all operations are over
	defer response.Body.Close()

	// read the response body
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return resultTimeData, fmt.Errorf("could not read http repsonse body in FetchServerTimeConfiguration(). error=%w", err)
	}

	jsonStr := string(responseData[:])
	if jsonStr == "" || jsonStr == "*No Data*" {
		return resultTimeData, errors.New("response json is blank/no data in FetchServerTimeConfiguration")
	}

	resultTimeData, err = jsonf.FetchIntervalAndOffset(jsonStr)
	if err != nil {
		return resultTimeData, errors.New("could not fetch server local time and svc call interval in FetchServerTimeConfiguration")
	}
	return resultTimeData, nil
}

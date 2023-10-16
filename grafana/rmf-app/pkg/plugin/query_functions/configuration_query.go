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
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	errh "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/error_handler"
	httphlper "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/http_helper"
	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

type ConfigurationQuery struct {
}

func (c *ConfigurationQuery) FetchRootInfo(
	em *typ.DatasourceEndpointModel,
	errHandler *errh.ErrHandler) (bool, error) {
	var httphlpr httphlper.HttpHelper
	url := httphlpr.GetHttpUrlForRoot(em)
	responseData, err := httphlpr.ExecuteHttpGet(url, em)
	if err != nil {
		return false, fmt.Errorf("could not fetch health check configuration in FetchRootInfo() - error=%v", err)
	}
	defer responseData.Body.Close()
	if responseData.StatusCode == 401 { // Unauthorized
		return false, fmt.Errorf("authentication failed in FetchRootInfo(). possible invalid username/password combination")
	}
	response, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		return false, fmt.Errorf("could in ioutil.ReadAll() in FetchRootInfo() - error=%v", err)
	}

	jsonStr := string(response[:])
	errHandler.LogStatus(fmt.Sprintf("\n***Fetched Root Info for url: %s and got response in FetchRootInfo(): %s", url, jsonStr))
	if jsonStr == "" || jsonStr == "*No Data*" {
		return false, fmt.Errorf("response json is blank/no data in FetchRootInfo - error=%v", err)
	} else {
		var jfuncs jsonf.JsonFunctions
		jsonObj, err := jfuncs.GetJsonObject(jsonStr)
		if err != nil {
			if strings.Contains(jsonStr, "<?xml version=\"1.0\" encoding=\"UTF-8\"?") {
				return false, typ.NewValueError(100, errors.New("unsupported version of DDS"))
			} else {
				return false, fmt.Errorf("response not valid in FetchRootInfo - error=%v", err)
			}
		}
		return len(jsonObj) > 0, nil
	}
}

func (c *ConfigurationQuery) FetchMetricsFromIndex(em *typ.DatasourceEndpointModel) ([]byte, error) {
	var httphlpr httphlper.HttpHelper
	url := httphlpr.GetHttpUrlForIndex(em)
	responseData, err := httphlpr.ExecuteHttpGet(url, em)
	if err != nil {
		return nil, fmt.Errorf("could not fetch metrics from index, url= %s in FetchMetricsFromIndex() - error= %v", url, err)
	}
	defer responseData.Body.Close()
	response, err := ioutil.ReadAll(responseData.Body)
	if err != nil {
		return nil, fmt.Errorf("error in ioutil.ReadAll() in FetchMetricsFromIndex() - error=%v", err)
	}
	return response, nil
}

func (c *ConfigurationQuery) FetchIntervalAndOffset(em *typ.DatasourceEndpointModel) (typ.IntervalOffset, error) {
	// Fetch the Http Url from endpoint model
	// <proto>://<host>:<port>/gpm/perform.xml?resource=,,SYSPLEX&id=8D0D50
	var (
		httphlpr       httphlper.HttpHelper
		resultTimeData typ.IntervalOffset
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
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return resultTimeData, fmt.Errorf("could not read http repsonse body in FetchServerTimeConfiguration(). error:=%v", err)
	}

	jsonStr := string(responseData[:])
	if jsonStr == "" || jsonStr == "*No Data*" {
		return resultTimeData, fmt.Errorf("response json is blank/no data in FetchServerTimeConfiguration")
	}

	var jfuncs jsonf.JsonFunctions
	resultTimeData, err = jfuncs.FetchIntervalAndOffset(jsonStr)
	if err != nil {
		return resultTimeData, fmt.Errorf("could not fetch server local time and svc call interval in FetchServerTimeConfiguration")
	}
	return resultTimeData, nil
}

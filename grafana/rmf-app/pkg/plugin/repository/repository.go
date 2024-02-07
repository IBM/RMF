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

package repository

import (
	"fmt"
	"io"
	"net/http"

	httphelpr "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/http_helper"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

type Repository struct {
}

func (r *Repository) ExecuteQueryAndGetResponse(queryModel *typ.QueryModel,
	endpointModel *typ.DatasourceEndpointModel) ([]byte, error) {
	// Get the http Response from service
	httpResponse, err := fetchDataFromServer(queryModel, endpointModel)
	if err != nil {
		return nil, fmt.Errorf("could not fetch data from server in ExecuteQueryAndGetResponse(). error:=%v", err)
	} else if httpResponse.StatusCode == 400 {
		return nil, fmt.Errorf("error fetching data from server in ExecuteQueryAndGetResponse(). possibly invalid query")
	}

	// Close the response body once all operations are over
	defer httpResponse.Body.Close()

	// read the response body
	responseData, err := io.ReadAll(httpResponse.Body)
	if responseData == nil && err != nil {
		return nil, fmt.Errorf("could not read http repsonse body in ExecuteQueryAndGetResponse(). error:=%v", err)
	}
	return responseData, nil
}

func fetchDataFromServer(qm *typ.QueryModel, em *typ.DatasourceEndpointModel) (*http.Response, error) {
	var httphlpr httphelpr.HttpHelper

	// Fetch the Http Url from query model and context
	url, err := httphlpr.GetHttpUrlFromContext(qm, em)
	if err != nil {
		return nil, err
	}
	// Store the queryModel
	qm.Url = url

	// Execute the query url
	response, err := httphlpr.ExecuteHttpGet(url, em)
	return response, err
}

func (r *Repository) ExecuteForVariableQuery(query string, em *typ.DatasourceEndpointModel) ([]byte, error) {
	// Get the http Response from service
	httpResponse, err := fetchDataForVariableQuery(query, em)
	if err != nil {
		return nil, fmt.Errorf("could not get httpResponse in ExecuteForVariableQuery(). Error=%v", err)
	}

	// Close the response body once all operations are over
	defer httpResponse.Body.Close()

	// read the response body
	responseData, err := io.ReadAll(httpResponse.Body)
	return responseData, err
}

func fetchDataForVariableQuery(query string, em *typ.DatasourceEndpointModel) (*http.Response, error) {
	var httphlpr httphelpr.HttpHelper

	// Fetch the Http Url from query model and endpoint
	url := httphlpr.GetHttpUrlFromQuery(query, em)

	// Execute the request and fetch the response
	response, err := httphlpr.ExecuteHttpGet(url, em)

	// defer response.Body.Close()
	return response, err
}

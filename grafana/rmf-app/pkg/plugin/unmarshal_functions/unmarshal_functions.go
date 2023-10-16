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
package unmarshal_functions

import (
	"encoding/json"
	"strings"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	guuid "github.com/google/uuid"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

type UnmarshalFunctions struct {
}

func (u *UnmarshalFunctions) UnmarshalQueryModel(pCtx backend.PluginContext, query backend.DataQuery) (*typ.QueryModel, error) {
	// Unmarshal the query JSON into our queryModel.
	var qm typ.QueryModel

	// Unmarshal both the JSON(s)
	err := json.Unmarshal(query.JSON, &qm)
	if err != nil {
		return nil, err
	}

	// Attach other fields such as TimeRange.From, TimeRange.To
	attachOtherQueryFields(pCtx, &qm, query)

	return &qm, nil
}

func (u *UnmarshalFunctions) UnmarshalEndpointModel(pCtx backend.PluginContext) (*typ.DatasourceEndpointModel, error) {
	// Unmarshal the pluginContext JSON into our endpointModel.
	var endpointModel typ.DatasourceEndpointModel

	err := json.Unmarshal(pCtx.DataSourceInstanceSettings.JSONData, &endpointModel)
	if err != nil {
		return nil, err
	}
	if pCtx.DataSourceInstanceSettings.DecryptedSecureJSONData != nil {
		val, ok := pCtx.DataSourceInstanceSettings.DecryptedSecureJSONData["password"]
		if ok {
			endpointModel.Password = val
		}
	}
	endpointModel.DatasourceUid = pCtx.DataSourceInstanceSettings.UID
	return &endpointModel, nil
}

func (u *UnmarshalFunctions) UnmarshalQueryAndEndpointModel(query backend.DataQuery, pCtx backend.PluginContext) (*typ.QueryModel, *typ.DatasourceEndpointModel, error) {
	qm, err := u.UnmarshalQueryModel(pCtx, query)
	if err != nil {
		return nil, nil, err
	}
	ep, err := u.UnmarshalEndpointModel(pCtx)
	if err != nil {
		return nil, nil, err
	}
	return qm, ep, nil
}

func attachOtherQueryFields(pCtx backend.PluginContext, qm *typ.QueryModel, query backend.DataQuery) {
	qm.TimeRangeFrom, qm.TimeRangeTo = query.TimeRange.From.UTC(), query.TimeRange.To.UTC()
}

func getGuidString() string {
	guidStr := guuid.New().String()
	guidStr = strings.ReplaceAll(guidStr, "-", "")
	return guidStr
}

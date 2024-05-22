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

package unmarshal_functions

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	http "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/http_helper"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
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
	// Unmarshal the pluginContext JSON into our em.
	var em typ.DatasourceEndpointModel
	dsSettings := pCtx.DataSourceInstanceSettings

	err := json.Unmarshal(pCtx.DataSourceInstanceSettings.JSONData, &em)
	if err != nil {
		return nil, err
	}
	if em.Server != "" || em.Port != "" {
		// Data source in legacy format
		protocol := "http"
		if em.SSL {
			protocol = "https"
		}
		em.URL = fmt.Sprintf(
			"%s://%s:%s", protocol, strings.TrimSpace(em.Server), strings.TrimSpace(em.Port))
		em.IntTimeout = http.DefaultHttpTimeout
		em.TlsSkipVerify = !em.VerifyInsecureCert
		if pCtx.DataSourceInstanceSettings.DecryptedSecureJSONData != nil {
			val, ok := pCtx.DataSourceInstanceSettings.DecryptedSecureJSONData["password"]
			if ok {
				em.Password = val
			}
		}
	} else {
		// Data source in conventional Grafana format
		em.URL = dsSettings.URL
		em.IntTimeout, err = strconv.Atoi(em.RawTimeout)
		if err != nil {
			em.IntTimeout = http.DefaultHttpTimeout
		}
		if dsSettings.BasicAuthEnabled {
			em.UserName = dsSettings.BasicAuthUser
			em.Password = dsSettings.DecryptedSecureJSONData["basicAuthPassword"]
		}
	}
	em.DatasourceUid = pCtx.DataSourceInstanceSettings.UID
	return &em, nil
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

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

package cache

import (
	"encoding/json"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	"github.com/VictoriaMetrics/fastcache"
)

const channelCacheSize = 128 * 1024 * 1024

var channelCache *fastcache.Cache

func init() {
	channelCache = fastcache.New(channelCacheSize)
}

func GetChannelQuery(path string) (*typ.QueryModel, error) {
	var query typ.QueryModel
	queryBytes := channelCache.Get(nil, []byte(path))
	err := json.Unmarshal(queryBytes, &query)
	return &query, err
}

func SetChannelQuery(path string, query *typ.QueryModel) error {
	queryBytes, err := json.Marshal(*query)
	if err == nil {
		channelCache.Set([]byte(path), queryBytes)
	}
	return err
}

func HasChannelQuery(path string) bool {
	return channelCache.Has([]byte(path))
}

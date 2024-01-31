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

package frame_functions

import (
	"strings"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

func GetFrameName(qm *typ.QueryModel) string {
	var resultFrameName string
	if strings.Trim(qm.SelectedQuery, "") != "" {
		splitStringSlice := strings.SplitAfter(qm.SelectedQuery, ".")
		if len(splitStringSlice) > 1 {
			vt := splitStringSlice[1]
			if strings.ToLower(vt) == "report." {
				resultFrameName = splitStringSlice[1] + splitStringSlice[2]
			} else {
				if strings.Contains(vt, "{") {
					resultFrameName = vt[:strings.Index(vt, "{")]
				} else {
					resultFrameName = vt
				}
			}
		}
	}
	return strings.Trim(resultFrameName, " ")
}

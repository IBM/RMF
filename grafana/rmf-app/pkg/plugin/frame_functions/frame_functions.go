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
	"time"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"

	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type FrameFunctions struct{}

func (fc *FrameFunctions) GetFrameName(qm *typ.QueryModel) string {
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

func (fc *FrameFunctions) RemoveOutOfRangeDatapoints(queryModel *typ.QueryModel, newFrame *data.Frame) *data.Frame {
	timeDifference := queryModel.TimeRangeTo.Sub(queryModel.TimeRangeFrom)
	earliestTimeToInclude := queryModel.TimeSeriesTimeRangeTo.Add(-timeDifference * 2) // We subtract 2 time differences.

	// Loop through the rows and remove those rows that fall before 'earliestTimeToInclude' value
	for rowIndex := 0; rowIndex < newFrame.Rows(); rowIndex++ {
		currentTimeValue := newFrame.At(0, rowIndex)
		if currentTimeValue.(time.Time).Before(earliestTimeToInclude) {
			newFrame.DeleteRow(rowIndex)
		}
	}
	// Get the meta data fro Json. The meta data info like numSamples is displayed on top of the header
	newFrame.Meta = &data.FrameMeta{
		Custom: earliestTimeToInclude,
	}

	return newFrame
}

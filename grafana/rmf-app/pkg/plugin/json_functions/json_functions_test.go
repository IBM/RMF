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

package json_functions_test

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	jsonf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/json_functions"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name          string
	Description   string
	Skip          bool
	QueryModel    typ.QueryModel
	DdsResponse   json.RawMessage
	IsTimeSeries  bool
	ExpectedFrame json.RawMessage
	ExpectedError string
}

func LoadTestCases(t *testing.T) []TestCase {
	var testCases []TestCase
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	jsonFile, err := os.ReadFile(filepath.Join(dir, "testdata/MetricFrameFromJson.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(jsonFile, &testCases); err != nil {
		t.Fatal(err)
	}
	return testCases
}

func TestMetricFrameFromJson(t *testing.T) {
	for _, testCase := range LoadTestCases(t) {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Skip {
				t.Skip("skipped because skip flag is set")
			}
			var expectedJson bytes.Buffer
			err := json.Indent(&expectedJson, testCase.ExpectedFrame, "", "  ")
			if assert.Nil(t, err, "fail to indent") {
				frame, err := jsonf.MetricFrameFromJson(string(testCase.DdsResponse), &testCase.QueryModel, testCase.IsTimeSeries)
				if err == nil {
					actualJson, _ := json.MarshalIndent(frame, "", "  ")
					assert.Equal(t, expectedJson.String(), string(actualJson), "frames are not identical")
				} else {
					assert.Equal(t, testCase.ExpectedError, err.Error(), "unexpected error message")
				}
			}
		})
	}
}

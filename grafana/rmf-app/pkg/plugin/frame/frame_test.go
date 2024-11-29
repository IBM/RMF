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

package frame

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
)

type TestCase struct {
	Name          string
	Description   string
	Skip          bool
	QueryModel    QueryModel
	DdsResponse   *dds.Response
	ExpectedFrame json.RawMessage
	ExpectedError string
}

func LoadTestCases(t *testing.T) []TestCase {
	var testCases []TestCase
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	jsonFile, err := os.ReadFile(filepath.Join(dir, "testdata/frames.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(jsonFile, &testCases); err != nil {
		t.Fatal(err)
	}
	return testCases
}

func TestFrame(t *testing.T) {
	for _, testCase := range LoadTestCases(t) {
		t.Run(testCase.Name, func(t *testing.T) {
			if testCase.Skip {
				t.Skip("skipped because skip flag is set")
			}
			var expectedJson bytes.Buffer
			err := json.Indent(&expectedJson, testCase.ExpectedFrame, "", "  ")
			if assert.NoError(t, err, "failed to indent") {
				frame, err := Build(testCase.DdsResponse, nil, &testCase.QueryModel)
				if err == nil {
					actualJson, _ := json.MarshalIndent(frame, "", "  ")
					assert.JSONEq(t, expectedJson.String(), string(actualJson), "frames are not identical")
				} else {
					assert.Equal(t, testCase.ExpectedError, err.Error(), "unexpected error message")
				}
			}
		})
	}
}

/**
* (C) Copyright IBM Corp. 2023, 2025.
* (C) Copyright Rocket Software, Inc. 2023-2025.
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

package dds

import (
	"encoding/xml"
	"regexp"
	"strings"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
)

const NoReport = ""

var XslConditionRe = regexp.MustCompile(`\$(\w*)\s*=\s*'([^']*)'`)

// HeaderMap is a map of report names to maps of column ids to column display names.
// E.g. {"STORCR": {"CSXNAME": "Job Name"}} or {NoReport: {"CADPDEVN": "4-Digit Device Number"}
// when column id is not bound to a report in the XSL mapping directly.
type HeaderMap map[string]map[string]string

// Get returns the display name of a column in a report based on the column id.
// If no display name is found, the column id itself is returned.
func (m HeaderMap) Get(report string, colId string) string {
	if _, ok := m[report]; !ok {
		report = NoReport
		if _, ok := m[report]; !ok {
			return colId
		}
	}
	if colName, ok := m[report][colId]; ok {
		return colName
	}
	return colId
}

type XslHeaderMap struct {
	Template struct {
		Choose XslChoose `xml:"choose"`
	} `xml:"template"`
}

type XslChoose struct {
	When []struct {
		Test   string    `xml:"test,attr"`
		Text   string    `xml:",chardata"`
		Choose XslChoose `xml:"choose"`
	} `xml:"when"`
}

func (c *Client) GetCachedHeaders() *HeaderMap {
	c.rwMutex.RLock()
	headerMap := c.headerMap
	c.rwMutex.RUnlock()
	if headerMap != nil {
		return headerMap
	}
	return c.updateHeaders()
}

func (c *Client) updateHeaders() *HeaderMap {
	logger := log.Logger.With("func", "GetHeaderMap")
	result, _, _ := c.single.Do("headers", func() (any, error) {
		headers := HeaderMap{}
		raw, err := c.GetRaw(XslHeadersPath)
		if err != nil {
			logger.Error("failed to fetch XSL header map", "error", err)
			return &headers, err
		}
		var xslHeaders XslHeaderMap
		if err = xml.Unmarshal(raw, &xslHeaders); err != nil {
			logger.Error("failed to unmarshal XLS header map", "error", err)
			return &headers, err
		}
		buildHeaders(headers, NoReport, xslHeaders.Template.Choose)
		c.rwMutex.Lock()
		c.headerMap = &headers
		c.rwMutex.Unlock()
		logger.Debug("header map updated")
		return &headers, nil
	})
	return result.(*HeaderMap)
}

func buildHeaders(res HeaderMap, report string, choose XslChoose) {
	logger := log.Logger.With("func", "buildHeaders")
	for _, when := range choose.When {
		condition := strings.TrimSpace(when.Test)
		match := XslConditionRe.FindStringSubmatch(condition)
		if len(match) > 0 {
			key, value := match[1], match[2]
			switch key {
			case "var":
				if _, ok := res[report]; !ok {
					res[report] = make(map[string]string)
				}
				res[report][value] = strings.TrimSpace(when.Text)
			case "report":
				buildHeaders(res, value, when.Choose)
			default:
				logger.Error("unexpected condition key in XSL header map", "key", key)
			}
		} else {
			logger.Error("unexpected condition in XSL header map", "condition", condition)
		}
		buildHeaders(res, report, when.Choose)
	}
}

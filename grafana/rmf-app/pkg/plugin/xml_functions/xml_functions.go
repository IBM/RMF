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

package xml_functions

import (
	"fmt"
	"regexp"
	"strings"

	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"

	"github.com/antchfx/xmlquery"
)

type XmlFunctions struct {
}

func (x *XmlFunctions) GetColHeaderXslMap(xslContent string) (map[string][]typ.ColHeaderXslMap, error) {
	resultMap := make(map[string][]typ.ColHeaderXslMap)
	doc, err := xmlquery.Parse(strings.NewReader(xslContent))
	if err != nil {
		return nil, fmt.Errorf("could not Parse data in GetColHeaderXslMap(). Error=%w", err)
	}

	// Attach the headers where report = 'NA'
	xPath := "//xsl:stylesheet/xsl:template/xsl:choose/xsl:when[contains(@test,'$var')]"
	normalMaps := getHeaderInfoSlice(doc, xPath)
	resultMap["NA"] = normalMaps

	// Attach the headers where report = 'STORCR'
	xPath = "//xsl:stylesheet/xsl:template/xsl:choose/xsl:when[contains(@test,\"$report = 'STORCR'\")]/xsl:choose/xsl:when"
	report1Maps := getHeaderInfoSlice(doc, xPath)
	resultMap["STORCR"] = report1Maps

	// Attach the headers where report = 'STORC'
	xPath = "//xsl:stylesheet/xsl:template/xsl:choose/xsl:when[contains(@test,\"$report = 'STORC'\")]/xsl:choose/xsl:when"
	report2Maps := getHeaderInfoSlice(doc, xPath)
	resultMap["STORC"] = report2Maps

	return resultMap, err
}

func getHeaderInfoSlice(doc *xmlquery.Node, xPath string) []typ.ColHeaderXslMap {
	var resultList []typ.ColHeaderXslMap //nolint:prealloc
	nodes := xmlquery.Find(doc, xPath)
	for _, node := range nodes {
		var mp typ.ColHeaderXslMap
		attr := extractHeaderId(node.SelectAttr("test"))
		inText := node.InnerText()
		mp.HeaderID = attr
		mp.HeaderText = inText
		resultList = append(resultList, mp)
	}
	return resultList
}

func extractHeaderId(s string) string {
	re := regexp.MustCompile(`'(.*)'`)
	match := re.FindStringSubmatch(s)
	return match[1]
}

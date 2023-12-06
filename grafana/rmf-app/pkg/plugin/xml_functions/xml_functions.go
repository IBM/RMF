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
		return nil, fmt.Errorf("could not Parse data in GetColHeaderXslMap(). Error=%v", err)
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

func (x *XmlFunctions) GetDataFormat(inputXml string) (string, error) {
	doc, err := xmlquery.Parse(strings.NewReader(inputXml))
	if err != nil || strings.Trim(inputXml, " ") == "" {
		return "", fmt.Errorf("error parsing xml in GetDataFormat() - error=%v", err)
	}

	xPath := "//report/metric/format"
	node := xmlquery.FindOne(doc, xPath)
	var innerText string
	if node != nil {
		innerText = node.InnerText() // possible values: single, report or list
		if innerText == "report" {
			xPath = "//report/row"
		} else if innerText == "list" {
			xPath = "//row"
		} else {
			return innerText, nil
		}

		nodeList := xmlquery.Find(doc.SelectElement("ddsml"), xPath)
		if len(nodeList) == 1 && innerText == "report" {
			return "report_single", nil
		} else if len(nodeList) == 1 && innerText == "list" {
			return "list_single", nil
		}
		return innerText, nil
	}
	return "", fmt.Errorf("the result data xml is invalid in GetDataFormat() - error=%v", err)
}

func getHeaderInfoSlice(doc *xmlquery.Node, xPath string) []typ.ColHeaderXslMap {
	var resultList []typ.ColHeaderXslMap
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

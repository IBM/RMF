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

package main

import (
	_ "embed"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

//go:embed reptrans.xsl
var reptransXSL string

const dateTimeFormat = "20060102150405"

// mockConfig holds the configurable mock topology.
type mockConfig struct {
	SysplexName string
	Systems     []string
	TimeOffset  time.Duration // offset from UTC (e.g., 0 for UTC, -5h for EST)
	Interval    int           // gatherer interval in seconds
}

var defaultConfig = mockConfig{
	SysplexName: "MOCKPLEX",
	Systems:     []string{"SYS1", "SYS2", "SYS3"},
	TimeOffset:  0,
	Interval:    60,
}

// metricDef defines a known metric and how to generate its data.
type metricDef struct {
	Id          string
	Description string
	Format      string // "single", "list", "report"
	Unit        string
	Columns     []Col
	Generator   func(cfg *mockConfig, t time.Time) []Row
}

var knownMetrics = map[string]metricDef{
	"8D0D60": {
		Id:          "8D0D60",
		Description: "# users by MVS image",
		Format:      "list",
		Unit:        "count",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPNUSR", Type: "N"}},
		Generator:   genUsersByImage,
	},
	"8D0460": {
		Id:          "8D0460",
		Description: "% CPU utilization (CP)",
		Format:      "single",
		Unit:        "%",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPCPUA", Type: "N"}},
		Generator:   genCPUUtilization,
	},
	"8D24B0": {
		Id:          "8D24B0",
		Description: "% effective logical utilization (CP)",
		Format:      "single",
		Unit:        "%",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPELUA", Type: "N"}},
		Generator:   genEffLogicalUtil,
	},
	"8D0F00": {
		Id:          "8D0F00",
		Description: "execution velocity",
		Format:      "single",
		Unit:        "%",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPEVEL", Type: "N"}},
		Generator:   genExecVelocity,
	},
	"8D1200": {
		Id:          "8D1200",
		Description: "transaction ended rate",
		Format:      "single",
		Unit:        "count/sec",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPTENR", Type: "N"}},
		Generator:   genTxnEndedRate,
	},
	"8D0100": {
		Id:          "8D0100",
		Description: "total LPAR CPU utilization",
		Format:      "list",
		Unit:        "%",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPLPCA", Type: "N"}},
		Generator:   genLPARCPU,
	},
	"8D1800": {
		Id:          "8D1800",
		Description: "real storage utilization",
		Format:      "single",
		Unit:        "%",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPRSTU", Type: "N"}},
		Generator:   genStorageUtil,
	},
	"8D2000": {
		Id:          "8D2000",
		Description: "paging rate",
		Format:      "single",
		Unit:        "pages/sec",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "CADPPGRT", Type: "N"}},
		Generator:   genPagingRate,
	},
}

// timeWobble creates a deterministic-ish value that varies with time, seeded by a name.
func timeWobble(t time.Time, name string, base, amplitude float64) float64 {
	// Use minute-of-day + name hash for pseudo-random but reproducible wobble
	minuteOfDay := float64(t.Hour()*60 + t.Minute())
	nameHash := float64(0)
	for _, c := range name {
		nameHash += float64(c)
	}
	val := base + amplitude*math.Sin((minuteOfDay+nameHash)/60.0*math.Pi)
	// Add a small random jitter
	val += (rand.Float64() - 0.5) * amplitude * 0.1
	return math.Max(0, val)
}

func genUsersByImage(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		count := int(timeWobble(t, sys, 500, 200))
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%d", count)}})
	}
	return rows
}

func genCPUUtilization(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		pct := timeWobble(t, sys+"cpu", 45, 25)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", math.Min(pct, 100))}})
	}
	return rows
}

func genEffLogicalUtil(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		pct := timeWobble(t, sys+"elu", 35, 20)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", math.Min(pct, 100))}})
	}
	return rows
}

func genExecVelocity(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		pct := timeWobble(t, sys+"ev", 75, 15)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", math.Min(pct, 100))}})
	}
	return rows
}

func genTxnEndedRate(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		rate := timeWobble(t, sys+"txn", 150, 80)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.1f", rate)}})
	}
	return rows
}

func genLPARCPU(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		pct := timeWobble(t, sys+"lpar", 55, 30)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", math.Min(pct, 100))}})
	}
	return rows
}

func genStorageUtil(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		pct := timeWobble(t, sys+"stor", 65, 15)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", math.Min(pct, 100))}})
	}
	return rows
}

func genPagingRate(cfg *mockConfig, t time.Time) []Row {
	rows := make([]Row, 0, len(cfg.Systems))
	for _, sys := range cfg.Systems {
		rate := timeWobble(t, sys+"pag", 20, 15)
		rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.1f", rate)}})
	}
	return rows
}

// genUnknownMetric generates plausible data for any metric ID not explicitly defined.
func genUnknownMetric(cfg *mockConfig, metricID string, t time.Time) metricDef {
	return metricDef{
		Id:          metricID,
		Description: fmt.Sprintf("metric %s", metricID),
		Format:      "list",
		Unit:        "",
		Columns:     []Col{{Value: "CADPNAME", Type: "T"}, {Value: "VALUE", Type: "N"}},
		Generator: func(cfg *mockConfig, t time.Time) []Row {
			rows := make([]Row, 0, len(cfg.Systems))
			for _, sys := range cfg.Systems {
				val := timeWobble(t, sys+metricID, 50, 30)
				rows = append(rows, Row{Cols: []string{sys, fmt.Sprintf("%.2f", val)}})
			}
			return rows
		},
	}
}

// makeTimeData builds a TimeData block for the given reference time.
func makeTimeData(cfg *mockConfig, refTime time.Time, isSysplex bool) *TimeData {
	localTime := refTime.Add(cfg.TimeOffset)
	// Truncate to interval boundary
	intervalDur := time.Duration(cfg.Interval) * time.Second
	localTime = localTime.Truncate(intervalDur)
	utcTime := localTime.Add(-cfg.TimeOffset)

	localStart := localTime.Format(dateTimeFormat)
	localEnd := localTime.Add(intervalDur).Format(dateTimeFormat)
	utcStart := utcTime.Format(dateTimeFormat)
	utcEnd := utcTime.Add(intervalDur).Format(dateTimeFormat)
	localPrev := localTime.Add(-intervalDur / 2).Format(dateTimeFormat)
	localNext := localTime.Add(intervalDur + intervalDur/2).Format(dateTimeFormat)

	td := &TimeData{
		LocalStart: localStart,
		LocalEnd:   localEnd,
		UTCStart:   utcStart,
		UTCEnd:     utcEnd,
		LocalPrev:  localPrev,
		LocalNext:  localNext,
		NumSamples: 1,
		GathererInterval: TimeInterval{
			Unit:  "SECONDS",
			Value: cfg.Interval,
		},
		DataRange: TimeInterval{
			Unit:  "SECONDS",
			Value: cfg.Interval,
		},
	}

	if isSysplex {
		n := len(cfg.Systems)
		td.NumSystems = &n
	}

	return td
}

// makePerformResponse builds a full perform response for a given metric and resource.
func makePerformResponse(cfg *mockConfig, metricID, resource string, refTime time.Time) *Response {
	md, ok := knownMetrics[metricID]
	if !ok {
		md = genUnknownMetric(cfg, metricID, refTime)
	}

	isSysplex := strings.Contains(strings.ToUpper(resource), "SYSPLEX")
	restype := "MVS_IMAGE"
	if isSysplex {
		restype = "SYSPLEX"
	}

	// Parse the resource label; default to sysplex if empty-ish
	reslabel := resource
	if reslabel == "" || reslabel == ",,SYSPLEX" {
		reslabel = fmt.Sprintf(",%s,SYSPLEX", cfg.SysplexName)
		restype = "SYSPLEX"
	}

	var rows []Row
	if md.Generator != nil {
		rows = md.Generator(cfg, refTime)
	}

	report := Report{
		TimeData: makeTimeData(cfg, refTime, isSysplex),
		Metric: &Metric{
			Id:          md.Id,
			Description: md.Description,
			Format:      md.Format,
			NumCols:     len(md.Columns),
			Unit:        md.Unit,
			Filter:      "HI=20;ORD=VD",
			ListType:    "M",
			Workscope:   ",,G",
		},
		Resource: &Resource{
			Reslabel: reslabel,
			Restype:  restype,
		},
		ColumnHeaders: &ColumnHeaders{Cols: md.Columns},
		Rows:          rows,
	}

	return &Response{
		Reports:  []Report{report},
		Server:   defaultServer(),
		Messages: []Message{},
	}
}

func defaultServer() *ServerInfo {
	return &ServerInfo{
		Name:          "RMF-DDS-Mock",
		Version:       "03.02.00",
		Platform:      "z/OS",
		Functionality: "3650",
	}
}

// makeRootResponse builds the /gpm/root response.
// If a root.json snapshot was loaded, it is returned directly.
func makeRootResponse(cfg *mockConfig) *Response {
	if rootSnapshot != nil {
		return rootSnapshot
	}
	return &Response{
		ContainedResourcesList: []ContainedResourcesList{
			{
				Contained: &ContainedResources{
					Resources: []Resource{
						{
							Reslabel:   fmt.Sprintf(",%s,SYSPLEX", cfg.SysplexName),
							Restype:    "SYSPLEX",
							Expandable: "YES",
							Icon:       "rmfsyspl.gif",
						},
					},
				},
				Resource: &Resource{Reslabel: ",,root", Restype: "root"},
			},
		},
		Server:              defaultServer(),
		Reports:             []Report{},
		Messages:            []Message{},
		MetricList:          []MetricListEntry{},
		AttributeList:       []interface{}{},
		FilterInstancesList: []interface{}{},
		Postprocessor:       []interface{}{},
		WorkscopeList:       []interface{}{},
	}
}

// makeContainedResponse builds a /gpm/contained response for a given resource.
// If a contained.json snapshot was loaded, it is returned for SYSPLEX/root resources.
func makeContainedResponse(cfg *mockConfig, resource string) *Response {
	var resources []Resource

	resUpper := strings.ToUpper(resource)
	if strings.Contains(resUpper, "SYSPLEX") || strings.Contains(resUpper, "ROOT") {
		if containedSnapshot != nil {
			return containedSnapshot
		}
		// Return MVS images (systems) under the sysplex
		for _, sys := range cfg.Systems {
			resources = append(resources, Resource{
				Reslabel:   fmt.Sprintf(",%s,MVS_IMAGE", sys),
				Restype:    "MVS_IMAGE",
				Expandable: "YES",
				Icon:       "rmfmvsi.gif",
			})
		}
	} else {
		// For any other resource, return some plausible children
		resources = []Resource{
			{Reslabel: ",SUBSYS1,SUBSYSTEM", Restype: "SUBSYSTEM"},
			{Reslabel: ",SUBSYS2,SUBSYSTEM", Restype: "SUBSYSTEM"},
		}
	}

	return &Response{
		ContainedResourcesList: []ContainedResourcesList{
			{
				Contained: &ContainedResources{
					Resources: resources,
				},
			},
		},
		Server:   defaultServer(),
		Reports:  []Report{},
		Messages: []Message{},
	}
}

// makeIndexResponse builds the /gpm/index response with a representative metric catalog.
// This is only used as a fallback when no index.json snapshot is available.
func makeIndexResponse(cfg *mockConfig) *Response {
	sysplexMetrics := []IndexMetricItem{
		{Id: "8D0D60", Description: "# users by MVS image", Format: "list"},
		{Id: "8D0460", Description: "% CPU utilization (CP)", Format: "single"},
		{Id: "8D24B0", Description: "% effective logical utilization (CP)", Format: "single"},
		{Id: "8D0F00", Description: "execution velocity", Format: "single"},
		{Id: "8D1200", Description: "transaction ended rate", Format: "single"},
		{Id: "8D0100", Description: "total LPAR CPU utilization", Format: "list"},
		{Id: "8D1800", Description: "real storage utilization", Format: "single"},
		{Id: "8D2000", Description: "paging rate", Format: "single"},
	}

	mvsMetrics := []IndexMetricItem{
		{Id: "8D0460", Description: "% CPU utilization (CP)", Format: "single"},
		{Id: "8D24B0", Description: "% effective logical utilization (CP)", Format: "single"},
		{Id: "8D0F00", Description: "execution velocity", Format: "single"},
		{Id: "8D1200", Description: "transaction ended rate", Format: "single"},
		{Id: "8D1800", Description: "real storage utilization", Format: "single"},
		{Id: "8D2000", Description: "paging rate", Format: "single"},
		{Id: "8D3000", Description: "I/O rate", Format: "single"},
		{Id: "8D3100", Description: "channel busy", Format: "list"},
	}

	return &Response{
		MetricList: []MetricListEntry{
			{ResourceType: "SYSPLEX", Metrics: sysplexMetrics},
			{ResourceType: "MVS_IMAGE", Metrics: mvsMetrics},
		},
		Server:   defaultServer(),
		Reports:  []Report{},
		Messages: []Message{},
	}
}

// reportDef defines a known full report (rmfm3) and how to generate its data.
type reportDef struct {
	Id          string
	Description string
	Columns     []Col
	Caption     []CaptionVar
	Generator   func(cfg *mockConfig, t time.Time) []Row
}

// knownReports holds report definitions. All reports are auto-registered
// from snapshots at startup by InitSnapshotGenerators (see generators.go).
// Unknown reports requested at runtime fall back to genUnknownReport.
var knownReports = map[string]reportDef{}

// makeReportResponse builds a /gpm/rmfm3 response with format "report".
func makeReportResponse(cfg *mockConfig, reportName, resource string, refTime time.Time) *Response {
	rd, ok := knownReports[strings.ToUpper(reportName)]
	if !ok {
		rd = genUnknownReport(reportName)
	}

	reslabel := resource
	restype := "MVS_IMAGE"
	if reslabel == "" {
		reslabel = fmt.Sprintf(",%s,MVS_IMAGE", cfg.Systems[0])
	}
	if strings.Contains(strings.ToUpper(reslabel), "SYSPLEX") {
		restype = "SYSPLEX"
	}

	var rows []Row
	if rd.Generator != nil {
		rows = rd.Generator(cfg, refTime)
	}

	td := makeTimeData(cfg, refTime, false)
	td.NumSamples = 60 // reports typically aggregate many samples

	// Build caption. For SYSSUM, update date/time dynamically.
	var caption *Caption
	if len(rd.Caption) > 0 {
		capVars := make([]CaptionVar, len(rd.Caption))
		copy(capVars, rd.Caption)
		for i := range capVars {
			if capVars[i].Name == "SUMIDTVC" || capVars[i].Name == "SUMADTVC" {
				capVars[i].Value = refTime.Format("01/02/06 15:04:05")
			}
		}
		caption = &Caption{Vars: capVars}
	}

	report := Report{
		TimeData: td,
		Metric: &Metric{
			Id:          rd.Id,
			Description: rd.Description,
			Format:      "report",
			NumCols:     len(rd.Columns),
		},
		Resource: &Resource{
			Reslabel: reslabel,
			Restype:  restype,
		},
		ColumnHeaders: &ColumnHeaders{Cols: rd.Columns},
		Rows:          rows,
		Caption:       caption,
	}

	return &Response{
		Reports:  []Report{report},
		Server:   defaultServer(),
		Messages: []Message{},
	}
}

// genUnknownReport generates a plausible report-format response for unrecognized report names.
func genUnknownReport(reportName string) reportDef {
	return reportDef{
		Id:          strings.ToUpper(reportName),
		Description: fmt.Sprintf("%s (Report)", strings.ToUpper(reportName)),
		Columns: []Col{
			{Value: "NAME", Type: "T"},
			{Value: "STATUS", Type: "T"},
			{Value: "VALUE1", Type: "N"},
			{Value: "VALUE2", Type: "N"},
		},
		Caption: []CaptionVar{
			{Name: "RPTNAME", Value: strings.ToUpper(reportName)},
			{Name: "RPTTYPE", Value: "Monitor III"},
			{Name: "RPTSYS", Value: "MOCK"},
		},
		Generator: func(cfg *mockConfig, t time.Time) []Row {
			var rows []Row
			for i, sys := range cfg.Systems {
				v1 := timeWobble(t, sys+reportName+"1", 50, 30)
				v2 := timeWobble(t, sys+reportName+"2", 100, 60)
				status := "Active"
				if i%3 == 2 {
					status = "Idle"
				}
				rows = append(rows, Row{Cols: []string{
					sys, status,
					fmt.Sprintf("%.2f", v1),
					fmt.Sprintf("%.2f", v2),
				}})
			}
			return rows
		},
	}
}

// parseRange parses the DDS range parameter "YYYYMMDDHHMMSS,YYYYMMDDHHMMSS"
// and returns the start time. If parsing fails, returns current time.
func parseRange(rangeParam string) time.Time {
	if rangeParam == "" {
		return time.Now().UTC()
	}
	parts := strings.SplitN(rangeParam, ",", 2)
	t, err := time.Parse(dateTimeFormat, parts[0])
	if err != nil {
		return time.Now().UTC()
	}
	return t
}

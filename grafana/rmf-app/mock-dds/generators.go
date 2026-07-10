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
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// snapshotDir is the directory containing real DDS JSON snapshots.
// It is set once at startup from the -data flag or defaults to "testdata"
// relative to the working directory.
var snapshotDir = "data"

// ---------------------------------------------------------------------------
// Snapshot cache: loaded once on first access
// ---------------------------------------------------------------------------

// snapshotReport holds pre-parsed report data from a JSON snapshot.
type snapshotReport struct {
	Rows     []Row        // original rows
	ColTypes []string     // "N" or "T" per column, derived from columnHeaders
	Caption  []CaptionVar // caption variables (if any)
	// Metadata from the snapshot for auto-registering unknown reports
	ReportMeta *Metric // metric section from the snapshot
	Columns    []Col   // columnHeaders from the snapshot
}

// snapshotMetric holds pre-parsed metric data from a JSON snapshot.
type snapshotMetric struct {
	Rows     []Row
	ColTypes []string
	// Metadata from the snapshot for auto-registering unknown metrics
	MetricMeta *Metric // metric section from the snapshot
	Columns    []Col   // columnHeaders from the snapshot
}

var (
	reportSnapshots   map[string]*snapshotReport // keyed by uppercase report name
	metricSnapshots   map[string]*snapshotMetric // keyed by metric ID (e.g. "8D0D60")
	rootSnapshot      *Response                  // loaded from root.json if available
	containedSnapshot *Response                  // loaded from contained.json if available
	indexSnapshotRaw  []byte                     // raw bytes from index.json, served as-is
)

// ensureSnapshotsLoaded loads all snapshot files once.
func ensureSnapshotsLoaded() {
	if reportSnapshots != nil {
		return
	}
	reportSnapshots = make(map[string]*snapshotReport)
	metricSnapshots = make(map[string]*snapshotMetric)

	entries, err := os.ReadDir(snapshotDir)
	if err != nil {
		log.Printf("WARNING: cannot read snapshot dir %q: %v (generators will return empty data)", snapshotDir, err)
		return
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		path := filepath.Join(snapshotDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("WARNING: cannot read %s: %v", path, err)
			continue
		}

		var resp Response
		if err := json.Unmarshal(data, &resp); err != nil {
			log.Printf("WARNING: cannot parse %s: %v", path, err)
			continue
		}

		name := strings.TrimSuffix(e.Name(), ".json")

		if strings.HasPrefix(name, "report_") {
			reportName := strings.ToUpper(strings.TrimPrefix(name, "report_"))
			sr := extractSnapshotReport(&resp)
			if sr != nil {
				reportSnapshots[reportName] = sr
				log.Printf("  loaded report snapshot: %s (%d rows)", reportName, len(sr.Rows))
			}
		} else if strings.HasPrefix(name, "metric_") {
			metricID := strings.ToUpper(strings.TrimPrefix(name, "metric_"))
			// Strip resource-type suffixes like "_MVSIMAGE" from the key
			metricID = strings.TrimSuffix(metricID, "_MVSIMAGE")
			sm := extractSnapshotMetric(&resp)
			if sm != nil {
				metricSnapshots[metricID] = sm
				log.Printf("  loaded metric snapshot: %s (%d rows)", metricID, len(sm.Rows))
			}
		}
	}

	// Load structural snapshots (root, contained).
	for _, sf := range []struct {
		file string
		dest **Response
	}{
		{"root.json", &rootSnapshot},
		{"contained.json", &containedSnapshot},
	} {
		if data, err := os.ReadFile(filepath.Join(snapshotDir, sf.file)); err == nil {
			var resp Response
			if err := json.Unmarshal(data, &resp); err != nil {
				log.Printf("WARNING: cannot parse %s: %v", sf.file, err)
			} else {
				*sf.dest = &resp
				log.Printf("  loaded snapshot: %s", sf.file)
			}
		}
	}

	// Load index.json as raw bytes so it can be served verbatim
	// (the Response struct doesn't capture all index.json fields).
	if data, err := os.ReadFile(filepath.Join(snapshotDir, "index.json")); err == nil {
		// Validate it's valid JSON, then store the raw bytes.
		if !json.Valid(data) {
			log.Printf("WARNING: index.json is not valid JSON")
		} else {
			indexSnapshotRaw = data
			log.Printf("  loaded snapshot: index.json (raw, %d bytes)", len(data))
		}
	}

	log.Printf("Snapshots loaded: %d reports, %d metrics", len(reportSnapshots), len(metricSnapshots))
}

func extractSnapshotReport(resp *Response) *snapshotReport {
	if len(resp.Reports) == 0 {
		return nil
	}
	rpt := resp.Reports[0]

	// Skip snapshots that are error/NoData responses
	if rpt.Message != nil && rpt.Message.Severity > 0 {
		return nil
	}
	if len(rpt.Rows) > 0 && len(rpt.Rows[0].Cols) > 0 && rpt.Rows[0].Cols[0] == "*NoData*" {
		return nil
	}

	var colTypes []string
	var cols []Col
	if rpt.ColumnHeaders != nil {
		for _, c := range rpt.ColumnHeaders.Cols {
			colTypes = append(colTypes, c.Type)
			cols = append(cols, c)
		}
	}
	var caption []CaptionVar
	if rpt.Caption != nil {
		caption = rpt.Caption.Vars
	}
	return &snapshotReport{
		Rows:       rpt.Rows,
		ColTypes:   colTypes,
		Caption:    caption,
		ReportMeta: rpt.Metric,
		Columns:    cols,
	}
}

func extractSnapshotMetric(resp *Response) *snapshotMetric {
	if len(resp.Reports) == 0 {
		return nil
	}
	rpt := resp.Reports[0]

	// Skip snapshots that are error/NoData responses
	if rpt.Message != nil && rpt.Message.Severity > 0 {
		return nil
	}
	// Skip if rows contain only *NoData*
	if len(rpt.Rows) > 0 && len(rpt.Rows[0].Cols) > 0 && rpt.Rows[0].Cols[0] == "*NoData*" {
		return nil
	}

	var colTypes []string
	var cols []Col
	if rpt.ColumnHeaders != nil {
		for _, c := range rpt.ColumnHeaders.Cols {
			colTypes = append(colTypes, c.Type)
			cols = append(cols, c)
		}
	} else if rpt.Metric != nil && rpt.Metric.NumCols > 0 {
		// Real DDS metric responses don't include columnHeaders.
		// Infer types: first column is text (label), rest are numeric.
		for i := 0; i < rpt.Metric.NumCols; i++ {
			if i == 0 {
				colTypes = append(colTypes, "T")
				cols = append(cols, Col{Value: "Name", Type: "T"})
			} else {
				colTypes = append(colTypes, "N")
				cols = append(cols, Col{Value: fmt.Sprintf("Value%d", i), Type: "N"})
			}
		}
	}
	return &snapshotMetric{
		Rows:       rpt.Rows,
		ColTypes:   colTypes,
		MetricMeta: rpt.Metric,
		Columns:    cols,
	}
}

// ---------------------------------------------------------------------------
// Core wobble engine: take snapshot rows and vary numeric values
// ---------------------------------------------------------------------------

// wobbleRows clones the snapshot rows and applies time-based wobble to all
// columns whose type is "N" (numeric). Text columns are kept as-is.
func wobbleRows(rows []Row, colTypes []string, t time.Time) []Row {
	result := make([]Row, len(rows))
	for i, row := range rows {
		cols := make([]string, len(row.Cols))
		copy(cols, row.Cols)

		for j, val := range cols {
			if val == "" {
				continue
			}
			// Only wobble numeric columns
			if j < len(colTypes) && colTypes[j] == "N" {
				cols[j] = wobbleNumericString(val, t, i, j)
			}
		}
		result[i] = Row{Cols: cols}
	}
	return result
}

// wobbleNumericString parses a numeric string, applies a small time-varying
// perturbation, and formats it back in the same style (int or float).
func wobbleNumericString(s string, t time.Time, rowIdx, colIdx int) string {
	s = strings.TrimSpace(s)
	if s == "" || s == "N/A" || s == "****" {
		return s
	}

	isFloat := strings.Contains(s, ".")

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return s // not a number, leave as-is
	}

	if f == 0 {
		return s // don't wobble zeros
	}

	// Deterministic seed from row/col position + time (changes each second)
	secondOfDay := float64(t.Hour()*3600 + t.Minute()*60 + t.Second())
	seed := secondOfDay + float64(rowIdx*1000+colIdx*7)

	// ±15% amplitude relative to the original value
	amplitude := math.Abs(f) * 0.15
	wobble := amplitude * math.Sin(seed/30.0*math.Pi)
	// Add per-request jitter (±50% of amplitude, i.e. ±7.5% of value)
	wobble += (rand.Float64() - 0.5) * amplitude

	result := f + wobble

	if !isFloat {
		// Integer formatting: preserve sign, round
		return strconv.Itoa(int(math.Round(result)))
	}

	// Detect decimal places from original
	parts := strings.Split(s, ".")
	decimals := len(parts[1])
	return strconv.FormatFloat(result, 'f', decimals, 64)
}

// ---------------------------------------------------------------------------
// Report generators backed by snapshots
// ---------------------------------------------------------------------------

// snapshotReportGenerator returns a generator function that serves rows from
// a pre-loaded snapshot, with numeric wobble applied at serve time.
// If no snapshot exists for the given report name, returns nil.
func snapshotReportGenerator(reportName string) func(cfg *mockConfig, t time.Time) []Row {
	ensureSnapshotsLoaded()
	sr, ok := reportSnapshots[strings.ToUpper(reportName)]
	if !ok || len(sr.Rows) == 0 {
		return nil // no snapshot or empty data — keep the fallback generator
	}
	// Capture sr in closure
	return func(cfg *mockConfig, t time.Time) []Row {
		return wobbleRows(sr.Rows, sr.ColTypes, t)
	}
}

// snapshotMetricGenerator returns a generator function for metrics.
func snapshotMetricGenerator(metricID string) func(cfg *mockConfig, t time.Time) []Row {
	ensureSnapshotsLoaded()
	sm, ok := metricSnapshots[strings.ToUpper(metricID)]
	if !ok || len(sm.Rows) == 0 {
		return nil // no snapshot or empty data — keep the fallback generator
	}
	return func(cfg *mockConfig, t time.Time) []Row {
		return wobbleRows(sm.Rows, sm.ColTypes, t)
	}
}

// ---------------------------------------------------------------------------
// Initialization: wire snapshot generators into knownReports / knownMetrics
// ---------------------------------------------------------------------------

// InitSnapshotGenerators should be called once at startup (after flag parsing)
// to override the generators in knownReports and knownMetrics with
// snapshot-backed versions wherever a snapshot file is available.
func InitSnapshotGenerators() {
	ensureSnapshotsLoaded()

	overridden := 0
	for name, rd := range knownReports {
		gen := snapshotReportGenerator(name)
		if gen != nil {
			rd.Generator = gen
			// Also update caption and columns from the snapshot
			sr := reportSnapshots[strings.ToUpper(name)]
			if sr != nil {
				if len(sr.Caption) > 0 {
					rd.Caption = sr.Caption
				}
				if len(sr.Columns) > 0 {
					rd.Columns = sr.Columns
				}
			}
			knownReports[name] = rd
			overridden++
		}
	}
	log.Printf("Overrode %d report generators with snapshot data", overridden)

	overridden = 0
	for id, md := range knownMetrics {
		gen := snapshotMetricGenerator(id)
		if gen != nil {
			md.Generator = gen
			knownMetrics[id] = md
			overridden++
		}
	}
	log.Printf("Overrode %d metric generators with snapshot data", overridden)

	// Second pass: register metrics from snapshots that aren't in knownMetrics yet
	added := 0
	for id, sm := range metricSnapshots {
		if _, exists := knownMetrics[id]; exists {
			continue // already handled above
		}
		if sm.MetricMeta == nil || len(sm.Rows) == 0 {
			continue
		}

		// Build columns from snapshot
		columns := sm.Columns
		if len(columns) == 0 {
			// Fallback: create generic columns based on numcols
			numCols := sm.MetricMeta.NumCols
			if numCols == 0 {
				numCols = 2
			}
			columns = make([]Col, numCols)
			columns[0] = Col{Value: "Name", Type: "T"}
			for j := 1; j < numCols; j++ {
				columns[j] = Col{Value: fmt.Sprintf("Value%d", j), Type: "N"}
			}
		}

		capturedSM := sm // capture for closure
		knownMetrics[id] = metricDef{
			Id:          sm.MetricMeta.Id,
			Description: sm.MetricMeta.Description,
			Format:      sm.MetricMeta.Format,
			Unit:        sm.MetricMeta.Unit,
			Columns:     columns,
			Generator: func(cfg *mockConfig, t time.Time) []Row {
				return wobbleRows(capturedSM.Rows, capturedSM.ColTypes, t)
			},
		}
		added++
	}
	log.Printf("Added %d new metric entries from snapshots", added)

	// Third pass: register reports from snapshots that aren't in knownReports yet
	added = 0
	for name, sr := range reportSnapshots {
		if _, exists := knownReports[name]; exists {
			continue // already handled above
		}
		if len(sr.Rows) == 0 {
			continue
		}

		// Build columns from snapshot
		columns := sr.Columns
		if len(columns) == 0 {
			// Fallback: create generic columns based on first row width
			numCols := 2
			if len(sr.Rows) > 0 {
				numCols = len(sr.Rows[0].Cols)
			}
			columns = make([]Col, numCols)
			columns[0] = Col{Value: "Name", Type: "T"}
			for j := 1; j < numCols; j++ {
				columns[j] = Col{Value: fmt.Sprintf("Value%d", j), Type: "N"}
			}
		}

		// Derive ID and description from metadata or the report name
		rid := strings.ToUpper(name)
		rdesc := "Report " + name
		if sr.ReportMeta != nil {
			if sr.ReportMeta.Id != "" {
				rid = sr.ReportMeta.Id
			}
			if sr.ReportMeta.Description != "" {
				rdesc = sr.ReportMeta.Description
			}
		}

		capturedSR := sr // capture for closure
		knownReports[name] = reportDef{
			Id:          rid,
			Description: rdesc,
			Columns:     columns,
			Caption:     sr.Caption,
			Generator: func(cfg *mockConfig, t time.Time) []Row {
				return wobbleRows(capturedSR.Rows, capturedSR.ColTypes, t)
			},
		}
		added++
	}
	log.Printf("Added %d new report entries from snapshots", added)
}

// SnapshotTopology extracts the sysplex name and system list from loaded
// root.json and contained.json snapshots. Returns empty strings/nil when
// no snapshots are available.
func SnapshotTopology() (sysplex string, systems []string) {
	if rootSnapshot == nil {
		return "", nil
	}
	for _, crl := range rootSnapshot.ContainedResourcesList {
		if crl.Contained == nil {
			continue
		}
		for _, r := range crl.Contained.Resources {
			if r.Restype == "SYSPLEX" {
				parts := strings.Split(r.Reslabel, ",")
				if len(parts) >= 2 {
					sysplex = parts[1]
					break
				}
			}
		}
		if sysplex != "" {
			break
		}
	}

	if containedSnapshot != nil {
		for _, crl := range containedSnapshot.ContainedResourcesList {
			if crl.Contained == nil {
				continue
			}
			for _, r := range crl.Contained.Resources {
				if r.Restype == "MVS_IMAGE" {
					parts := strings.Split(r.Reslabel, ",")
					if len(parts) >= 2 {
						systems = append(systems, parts[1])
					}
				}
			}
		}
	}

	return sysplex, systems
}

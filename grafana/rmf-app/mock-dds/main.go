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
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	port := flag.Int("port", 8803, "port to listen on")
	dataDir := flag.String("data", "data", "directory with JSON snapshot files (from fetch-data tool)")
	flag.Parse()

	// Load real-data snapshots from disk (if available) to back the generators.
	snapshotDir = *dataDir
	InitSnapshotGenerators()

	cfg := defaultConfig

	// Override config with topology from snapshots (root.json / contained.json).
	if sysplex, systems := SnapshotTopology(); sysplex != "" {
		cfg.SysplexName = sysplex
		log.Printf("Config: sysplex=%s (from snapshot)", sysplex)
		if len(systems) > 0 {
			cfg.Systems = systems
			log.Printf("Config: systems=%v (from snapshot)", systems)
		}
	}

	cfgp := &cfg

	mux := http.NewServeMux()

	// Register handlers for all DDS endpoints.
	// Each path is registered with and without the .xml extension
	// to support both DDS v1 and v2 client modes.
	registerJSON(mux, "/gpm/root", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, makeRootResponse(cfgp))
	})

	registerJSON(mux, "/gpm/index", func(w http.ResponseWriter, r *http.Request) {
		if indexSnapshotRaw != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(indexSnapshotRaw)
			return
		}
		writeJSON(w, makeIndexResponse(cfgp))
	})

	registerJSON(mux, "/gpm/contained", func(w http.ResponseWriter, r *http.Request) {
		resource := r.URL.Query().Get("resource")
		writeJSON(w, makeContainedResponse(cfgp, resource))
	})

	registerJSON(mux, "/gpm/perform", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		handlePerform(w, r, cfgp)
	})

	registerJSON(mux, "/gpm/rmfm3", func(w http.ResponseWriter, r *http.Request) {
		handleReport(w, r, cfgp)
	})

	// XSL headers endpoint returns XML, not JSON
	mux.HandleFunc("/gpm/include/reptrans.xsl", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		w.Header().Set("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, reptransXSL)
	})

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("DDS mock server starting on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

// handlePerform handles /gpm/perform requests (single/list format metrics).
func handlePerform(w http.ResponseWriter, r *http.Request, cfg *mockConfig) {
	q := r.URL.Query()
	resource := q.Get("resource")
	metricID := q.Get("id")
	rangeParam := q.Get("range")

	if metricID == "" {
		metricID = "8D0D60" // default to user count
	}

	refTime := parseRange(rangeParam)
	resp := makePerformResponse(cfg, metricID, resource, refTime)
	writeJSON(w, resp)
}

// handleReport handles /gpm/rmfm3 requests (full report format).
func handleReport(w http.ResponseWriter, r *http.Request, cfg *mockConfig) {
	q := r.URL.Query()
	resource := q.Get("resource")
	reportName := q.Get("report")
	rangeParam := q.Get("range")

	if reportName == "" {
		reportName = "SYSSUM" // default report
	}

	refTime := parseRange(rangeParam)
	resp := makeReportResponse(cfg, reportName, resource, refTime)
	writeJSON(w, resp)
}

// registerJSON registers a handler at the given path and also at path + ".xml",
// both returning Content-Type: application/json.
func registerJSON(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		handler(w, r)
	}
	mux.HandleFunc(path, wrapped)
	mux.HandleFunc(path+".xml", wrapped)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		log.Printf("error encoding response: %v", err)
	}
}

func logRequest(r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.String())
}

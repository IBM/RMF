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

package plugin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/httpclient"
)

const DefaultHttpTimeout = 60
const DefaultCacheSizeMB = 1024
const MinimalCacheSizeMB = 128
const MinBatchRequestMinutes = 10
const MaxBatchRequestMinutes = 120
const DefaultBatchRequestMinutes = MaxBatchRequestMinutes

type Config struct {
	URL                 string
	Timeout             int
	CacheSize           int
	BatchRequestMinutes int
	Username            string
	Password            string // #nosec G117
	JSON                struct {
		// Conventional Grafana HTTP config (see the `DataSourceHttpSettings` UI element)
		TimeoutRaw         string `json:"timeout"`
		TlsSkipVerify      bool   `json:"tlsSkipVerify"`
		DisableCompression bool   `json:"disableCompression"`
		// Custom RMF settings.
		CacheSizeRaw string `json:"cacheSize"`
		// Legacy custom RMF settings. We should get rid of these at some point.
		Server               *string `json:"path"`
		Port                 string  `json:"port"`
		SSL                  bool    `json:"ssl"`
		Username             string  `json:"userName"`
		Password             string  `json:"password"`   // #nosec G117
		SSLVerify            bool    `json:"skipVerify"` // NB: the meaning of JSON field is inverted.
		OmegamonDs           string  `json:"omegamonDs"`
		BatchRequestInterval string  `json:"batchRequestInterval"`
	}
}

func (ds *RMFDatasource) getConfig(ctx context.Context, settings backend.DataSourceInstanceSettings) (*Config, *httpclient.Options, error) {
	var config Config
	logger := log.Logger.With("func", "getConfig")
	err := json.Unmarshal(settings.JSONData, &config.JSON)
	if err != nil {
		return nil, nil, err
	}
	if config.JSON.Server != nil {
		// Data source in legacy format
		protocol := "http"
		if config.JSON.SSL {
			protocol = "https"
		}
		config.URL = fmt.Sprintf("%s://%s:%s", protocol, *config.JSON.Server, config.JSON.Port)
		config.Timeout = DefaultHttpTimeout
		config.JSON.TlsSkipVerify = !config.JSON.SSLVerify
		config.Username = config.JSON.Username
		if val, ok := settings.DecryptedSecureJSONData["password"]; ok {
			config.Password = val
		}
	} else {
		// Data source in conventional Grafana format
		config.URL = settings.URL
		if config.Timeout, err = strconv.Atoi(config.JSON.TimeoutRaw); err != nil {
			config.Timeout = DefaultHttpTimeout
		}
		if settings.BasicAuthEnabled {
			config.Username = settings.BasicAuthUser
			if val, ok := settings.DecryptedSecureJSONData["basicAuthPassword"]; ok {
				config.Password = val
			} else if val, ok := settings.DecryptedSecureJSONData["password"]; ok {
				// Password still may be in old format: it can't be converted in frontend, only reset.
				config.Password = val
			}
		}
	}
	if config.CacheSize, err = strconv.Atoi(config.JSON.CacheSizeRaw); err != nil {
		logger.Warn("cache size is not valid, applying default", "cacheSizeRaw", config.JSON.CacheSizeRaw)
		config.CacheSize = DefaultCacheSizeMB
	}
	if config.CacheSize < MinimalCacheSizeMB {
		logger.Warn("cache size is not small, using minimal value", "cacheSize", config.CacheSize)
		config.CacheSize = MinimalCacheSizeMB
	}

	httpOpts, err := settings.HTTPClientOptions(ctx)
	if err != nil {
		return nil, nil, err
	}
	if httpOpts.Timeouts == nil {
		defaults := httpclient.DefaultTimeoutOptions
		httpOpts.Timeouts = &defaults
	}
	httpOpts.Timeouts.Timeout = time.Duration(config.Timeout) * time.Second
	if config.JSON.TlsSkipVerify {
		if httpOpts.TLS == nil {
			httpOpts.TLS = &httpclient.TLSOptions{}
		}
		httpOpts.TLS.InsecureSkipVerify = true
	}
	if strings.TrimSpace(config.Username) != "" {
		httpOpts.BasicAuth = &httpclient.BasicAuthOptions{
			User:     config.Username,
			Password: config.Password,
		}
	}
	if config.JSON.DisableCompression {
		httpOpts.ConfigureTransport = func(_ httpclient.Options, transport *http.Transport) {
			transport.DisableCompression = true
		}
	}
	config.URL = strings.TrimRight(config.URL, "/")

	if config.BatchRequestMinutes, err = strconv.Atoi(config.JSON.BatchRequestInterval); err != nil {
		logger.Warn("batch request interval is not valid, applying default", "batchRequestInterval", config.JSON.BatchRequestInterval)
		config.BatchRequestMinutes = DefaultBatchRequestMinutes
	}
	if config.BatchRequestMinutes < MinBatchRequestMinutes {
		logger.Warn("batch request interval is too small, using minimal value", "batchRequestInterval", config.BatchRequestMinutes)
		config.BatchRequestMinutes = MinBatchRequestMinutes
	}
	if config.BatchRequestMinutes > MaxBatchRequestMinutes {
		logger.Warn("batch request interval is too large, using maximal value", "batchRequestInterval", config.BatchRequestMinutes)
		config.BatchRequestMinutes = MaxBatchRequestMinutes
	}

	return &config, &httpOpts, nil
}

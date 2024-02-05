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

package config_helper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"gopkg.in/yaml.v3"
)

type PluginConfig struct {
	Logging struct {
		TraceErrors bool `yaml:"trace_errors"`
		TraceCalls  bool `yaml:"trace_calls"`
	} `yaml:"logging"`
	CacheSettings struct {
		MetricsCacheSizeMB        int `yaml:"metrics_cache_size_mb"`
		QueryModelCacheSizeMB     int `yaml:"query_model_cache_size_mb"`
		IntervalOffsetCacheSizeMB int `yaml:"interval_offset_cache_size_mb"`
	} `yaml:"cache_settings"`
}

func NewPluginConfig() *PluginConfig {
	pluginConfig := PluginConfig{}

	executablePath, err := os.Executable()
	if err != nil {
		panic("could not get executablePath in NewPluginConfig()")
	}

	// Get the directory containing the executable
	executableDir := filepath.Dir(executablePath)
	yamlFilePath := filepath.Join(executableDir, "rmf-plugin-config.yml")

	if fileContent, err := os.ReadFile(yamlFilePath); err != nil {
		// Provide default values for the pluginConfig props
		pluginConfig.Logging.TraceErrors = true
		pluginConfig.Logging.TraceCalls = false
		pluginConfig.CacheSettings.MetricsCacheSizeMB = 512
		pluginConfig.CacheSettings.QueryModelCacheSizeMB = 512
		pluginConfig.CacheSettings.IntervalOffsetCacheSizeMB = 64
	} else {
		if err := yaml.Unmarshal(fileContent, &pluginConfig); err != nil {
			log.DefaultLogger.Info(fmt.Sprintf("could not get unmarshal filContent in NewPluginConfig(): details: %v ", err), nil)
		}
	}
	return &pluginConfig
}

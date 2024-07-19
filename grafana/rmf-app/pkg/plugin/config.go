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

package plugin

import (
	"os"
	"path/filepath"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"gopkg.in/yaml.v3"
)

const DefaultCacheSizeMB = 1024

var pluginConfig *PluginConfig

// TODO: it has to be defined on the data source level along with other options and editable via UI.
type PluginConfig struct {
	cache struct {
		size int `yaml:"size-mb"`
	} `yaml:"cache"`
}

func init() {
	pluginConfig = NewPluginConfig()
}

func NewPluginConfig() *PluginConfig {
	logger := log.Logger.With("func", "NewPluginConfig")
	pluginConfig := PluginConfig{}

	executablePath, err := os.Executable()
	if err != nil {
		panic("could not get executablePath in NewPluginConfig()")
	}

	executableDir := filepath.Dir(executablePath)
	yamlFilePath := filepath.Join(executableDir, "rmf-plugin-config.yml")

	if fileContent, err := os.ReadFile(yamlFilePath); err != nil {
		pluginConfig.cache.size = DefaultCacheSizeMB
	} else {
		if err := yaml.Unmarshal(fileContent, &pluginConfig); err != nil {
			logger.Info("could not get unmarshal filContent", "error", err)
			pluginConfig.cache.size = DefaultCacheSizeMB
		}
	}
	return &pluginConfig
}

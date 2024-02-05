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

package error_handler

import (
	"embed"
	"fmt"
	"runtime/debug"
	"time"

	idgen "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/id_generator"
	plugincnfg "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/plugin_config"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"gopkg.in/yaml.v3"
)

//go:embed error_codes.yml
var f embed.FS

type ErrHandler struct {
	idGenerator  idgen.IdGenerator
	errCodesDict map[string]string
	PluginConfig *plugincnfg.PluginConfig
}

func NewErrHandler(plgCnfg *plugincnfg.PluginConfig) *ErrHandler {
	errHandler := ErrHandler{idGenerator: *idgen.NewIdGenerator(),
		errCodesDict: make(map[string]string), PluginConfig: plgCnfg}
	if err := errHandler.makeErrorCodesDictionary(); err != nil {
		panic("backend plugin error: could not create error codes dictionary.")
	}
	return &errHandler
}

func (e *ErrHandler) makeErrorCodesDictionary() error {
	yfile, err := f.ReadFile("error_codes.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yfile, &e.errCodesDict)
	if err != nil {
		return err
	}
	return nil
}

// Default Panic handler for all panics raised by inner packages/components
func (e *ErrHandler) HandleErrors() {
	if r := recover(); r != nil {
		e.LogError("S", "ERR_INTERNAL_ERROR", fmt.Errorf("recovered from unexpected panic - details: %v, stackTrace: %s", r, string(debug.Stack())))
	}
}

// Logs any messages as Info with the default logger
func (e *ErrHandler) LogStatus(msg string, args ...interface{}) {
	// Log only if EnableTrace is true in rmf-plugin-config.yml
	if e.PluginConfig.Logging.TraceCalls {
		log.DefaultLogger.Info(msg, args)
	}
}

func (e *ErrHandler) LogError(severity string, errCode string, err error) {
	if err := e.LogErrorAndReturnErrorInfo(severity, errCode, err); err != nil {
		log.DefaultLogger.Info("LogError() completed.")
	}
}

func (e *ErrHandler) LogErrorAndReturnErrorInfo(severity string, errCode string, err error) error {
	const ERR_ID_LENGTH = 10
	var (
		returnErrorToUser error
		errDesc           string
	)

	errUniqueId := e.idGenerator.GenerateUniqueId(ERR_ID_LENGTH)

	if errCode == "ERR_DDS_ERROR" {
		errDesc = err.Error()
		returnErrorToUser = fmt.Errorf("%s", errDesc)
	} else {
		errDesc = e.errCodesDict[errCode]
		errorMsgToUser := fmt.Sprintf("%s. Please quote error id: %s for further diagnosis of this issue (severity=%s)", errDesc, errUniqueId, severity)
		returnErrorToUser = fmt.Errorf(errorMsgToUser)
	}
	errorFullMsg := fmt.Sprintf("\n\n****ERROR @ Time (UTC): %s*****", time.Now().UTC().String())
	errorFullMsg = errorFullMsg + fmt.Sprintf("\nError logged for error id: %s, severity: %s, err_code: %s, details: %v, stackTrace: %s\n", errUniqueId, severity, errCode, err, string(debug.Stack()))
	errorFullMsg = errorFullMsg + "****ERROR End*******\n"

	// Log only if EnableTrace is true in rmf-plugin-config.yml
	if e.PluginConfig.Logging.TraceErrors {
		log.DefaultLogger.Error(errorFullMsg)
	}

	return returnErrorToUser
}

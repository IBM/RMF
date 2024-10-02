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

package log

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"runtime/debug"
	"strings"

	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

var Logger = log.DefaultLogger

type ErrorCode int

const (
	InternalError ErrorCode = iota
	FrameError
	InputError
	ConnectionError
)

var ErrorCodeMap = map[ErrorCode]string{
	InternalError:   "An internal error occurred in the IBM RMF Datasource plugin. Please contact your administrator and check the logs.",
	FrameError:      "Unable to fetch data from server.",
	InputError:      "The input provided is invalid.",
	ConnectionError: "DDS connection failed.",
}

const errorIdLen = 10

var errorIdAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateErrorId() (string, error) {
	res := make([]rune, errorIdLen)
	for i := 0; i < errorIdLen; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(errorIdAlphabet))))
		if err != nil {
			return "n/a", err
		}
		res[i] = errorIdAlphabet[int(num.Int64())]
	}
	return string(res), nil
}

// ErrorWithId logs an error with a unique id and returns a message with the same code.
// The message is passed to frontend; the id can be used to identify corresponding backend error.
func ErrorWithId(logger log.Logger, errCode ErrorCode, msg string, args ...interface{}) error {
	errorId, err := generateErrorId()
	if err != nil {
		Logger.Error("unable to generate error id", "error", err, "func", "ErrorWithId")
	}
	args = append(args, "errorId", errorId)
	logger.Error(msg, args...)

	userErrDesc := ErrorCodeMap[errCode]
	return fmt.Errorf("%s (error id %s)", strings.Trim(userErrDesc, "."), errorId)
}

func FrameErrorWithId(logger log.Logger, err error) error {
	return ErrorWithId(logger, FrameError, "failed to get frame", "error", err)
}

func LogAndRecover(logger log.Logger) {
	if r := recover(); r != nil {
		stack := string(debug.Stack())
		logger.Error("recovered from panic", "error", r)
		for _, line := range strings.Split(stack, "\n") {
			logger.Error("recovered from panic", "stack", line)
		}
	}
}

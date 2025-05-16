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
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"github.com/google/uuid"
)

const sep = "-"

// See "github.com/grafana/grafana-plugin-sdk-go/live" for path id limitations
var maxChannelLength = 160
var illegalChar = regexp.MustCompile(`[^A-z0-9/=.]`)
var escapedChar = regexp.MustCompile(`_(\d+)_`)

func encodeChannelPath(res string, from time.Time, to time.Time, absolute bool, interval time.Duration) (string, error) {
	res = illegalChar.ReplaceAllStringFunc(
		res,
		func(s string) string {
			return fmt.Sprintf("_%d_", int(s[0]))
		})
	absFlag := "0"
	if absolute {
		absFlag = "1"
	}
	path := res +
		sep + strconv.FormatInt(from.Unix(), 10) +
		sep + strconv.FormatInt(to.Unix(), 10) +
		sep + absFlag +
		sep + strconv.FormatInt(int64(interval.Seconds()), 10)
	// Channel path has to be unique to avoid streaming issues
	// Append at least 8 char long random part
	if len(path) >= maxChannelLength-8 {
		return path, errors.New("too long resource value")
	}
	path += sep + uuid.NewString()[:min(160-len(path), 36)]
	return path, nil
}

func decodeChannelPath(path string) (string, time.Time, time.Time, bool, time.Duration, error) {
	logger := log.Logger.With("func", "decodeChannelPath")
	var (
		res      string
		from     time.Time
		to       time.Time
		absolute bool
		interval time.Duration
	)
	parts := strings.Split(path, sep)
	if len(parts) < 6 {
		return res, from, to, absolute, interval, errors.New("invalid number of elements in path")
	}
	res = escapedChar.ReplaceAllStringFunc(parts[0], func(match string) string {
		submatches := escapedChar.FindStringSubmatch(match)
		if len(submatches) > 1 {
			code, err := strconv.Atoi(submatches[1])
			if err == nil && code >= 0 && code <= unicode.MaxRune {
				return string(rune(code))
			}
			logger.Error("unable to decode resource value", "error", err, "match", match)
		}
		return match
	})
	if timestamp, err := strconv.ParseInt(parts[1], 10, 64); err != nil {
		return res, from, to, absolute, interval, err
	} else {
		from = time.Unix(timestamp, 0)
	}
	if timestamp, err := strconv.ParseInt(parts[2], 10, 64); err != nil {
		return res, from, to, absolute, interval, err
	} else {
		to = time.Unix(timestamp, 0)
	}
	if parts[3] == "1" {
		absolute = true
	}
	if d, err := strconv.ParseInt(parts[4], 10, 64); err != nil {
		return res, from, to, absolute, interval, err
	} else {
		interval = time.Duration(d) * time.Second
	}
	return res, from, to, absolute, interval, nil
}

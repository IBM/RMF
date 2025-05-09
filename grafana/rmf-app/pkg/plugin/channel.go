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
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const Sep = ":"

func encodeChannelPath(res string, from time.Time, to time.Time, absolute bool, interval time.Duration) string {
	absFlag := "0"
	if absolute {
		absFlag = "1"
	}
	path := res +
		Sep + strconv.FormatInt(from.Unix(), 10) +
		Sep + strconv.FormatInt(to.Unix(), 10) +
		Sep + absFlag +
		Sep + strconv.FormatInt(int64(interval.Seconds()), 10) +
		Sep + uuid.NewString()[:8]
	return base64.StdEncoding.EncodeToString([]byte(path))
}

func decodeChannelPath(b string) (string, time.Time, time.Time, bool, time.Duration, error) {
	var (
		res      string
		from     time.Time
		to       time.Time
		absolute bool
		interval time.Duration
	)
	path, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return res, from, to, absolute, interval, err
	}
	parts := strings.Split(string(path), Sep)
	if len(parts) != 6 {
		return res, from, to, absolute, interval, errors.New("invalid number of elements")
	}
	res = parts[0]
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

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

package cache

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type Cache struct {
	cache *fastcache.Cache
}

func NewFrameCache(size int) *Cache {
	return &Cache{cache: fastcache.New(size * 1024 * 1024)}
}

func (fc *Cache) Reset() {
	fc.cache.Reset()
}

func Key(r *dds.Request, wide bool) []byte {
	format := "long"
	if wide {
		format = "wide"
	}
	return []byte(fmt.Sprintf("%s[%s]@%d-%d", r.Resource, format, r.TimeRange.From.UnixMilli(), r.TimeRange.To.UnixMilli()))
}

func (fc *Cache) GetFrame(r *dds.Request, wide bool) *data.Frame {
	logger := log.Logger.With("func", "GetFrame")
	var frame data.Frame
	key := Key(r, wide)
	buf := fc.cache.GetBig(nil, key)
	if buf != nil {
		err := json.Unmarshal(buf, &frame)
		if err != nil {
			logger.Error("Unmarshal error", "err", err, "key", key)
			return nil
		} else {
			return &frame
		}
	}
	return nil
}

func (fc *Cache) SaveFrame(f *data.Frame, r *dds.Request, wide bool) error {
	key := Key(r, wide)
	frame := fc.GetFrame(r, wide)
	if frame != nil {
		return nil
	}
	val, err := json.Marshal(&f)
	if err != nil {
		return err
	}
	fc.cache.SetBig(key, val)
	return nil
}

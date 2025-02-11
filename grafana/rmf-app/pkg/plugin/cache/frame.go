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

package cache

import (
	"encoding/json"
	"errors"
	"math"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type CacheItemValue struct {
	ValueKey time.Time
	Value    data.Frame

	frame.ResponseStatus
}

type FrameCache struct {
	cache *fastcache.Cache
}

func NewFrameCache(size int) *FrameCache {
	return &FrameCache{cache: fastcache.New(size * 1024 * 1024)}
}

func (fc *FrameCache) getCacheItemValue(key []byte) *CacheItemValue {
	logger := log.Logger.With("func", "getCacheItemValue")
	byteCacheItemValue := fc.cache.GetBig(nil, key)
	if byteCacheItemValue != nil {
		var cacheItemValue CacheItemValue
		err := json.Unmarshal(byteCacheItemValue, &cacheItemValue)
		if err != nil {
			logger.Debug("Unmarshal error", "err", err, "key", string(key))
			return nil
		}
		return &cacheItemValue
	}
	return nil
}

func (fc *FrameCache) GetFrame(qm *frame.QueryModel) (*data.Frame, error) {
	var (
		resultframe *data.Frame
	)
	matchedCacheItem := fc.getCacheItemValue(qm.CacheKey())
	if matchedCacheItem != nil {
		diffInSecs := int(matchedCacheItem.ValueKey.Sub(qm.CurrentTime).Seconds())
		if int(math.Abs(float64(diffInSecs))) <= int(matchedCacheItem.Mintime) {
			qm.Update(&matchedCacheItem.ResponseStatus)
			resultframe = &matchedCacheItem.Value
		}
	} else {
		return nil, errors.New("frame not found in cache in GetFrame()")
	}
	return resultframe, nil
}

func (fc *FrameCache) SaveFrame(frame *data.Frame, qm *frame.QueryModel) error {
	logger := log.Logger.With("func", "SaveFrame")

	var cacheItemValue *CacheItemValue
	cacheItemValue = fc.getCacheItemValue(qm.CacheKey())
	if cacheItemValue != nil {
		if cacheItemValue.CurrentTime.Equal(qm.CurrentTime) {
			logger.Debug("cache item already exist", "key", string(qm.CacheKey()))
			return nil
		}
	}
	var newCacheItemValue CacheItemValue = fc.createCacheItemValue(frame, qm)

	if cacheItemValueBytes, err := json.Marshal(&newCacheItemValue); err != nil {
		return err
	} else {
		fc.cache.SetBig(qm.CacheKey(), cacheItemValueBytes)
		logger.Debug("cache item added", "key", string(qm.CacheKey()))
	}
	return nil
}

func (fc *FrameCache) createCacheItemValue(frame *data.Frame, qm *frame.QueryModel) CacheItemValue {
	var (
		cacheItemValue CacheItemValue
	)
	cacheItemValue.ValueKey = qm.CurrentTime
	cacheItemValue.Value = *frame
	cacheItemValue.Update(&qm.ResponseStatus)
	return cacheItemValue
}

func (fc *FrameCache) Reset() {
	fc.cache.Reset()
}

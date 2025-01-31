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
	"sort"
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

func (fc *FrameCache) getCacheItemValues(key []byte) ([]CacheItemValue, error) {
	var cacheItemValues []CacheItemValue
	byteCacheItemValues := fc.cache.GetBig(nil, key)
	if byteCacheItemValues == nil {
		return cacheItemValues, errors.New("no cache item")
	} else {
		err := json.Unmarshal(byteCacheItemValues, &cacheItemValues)
		if err != nil {
			return nil, err
		}
	}
	return cacheItemValues, nil
}

func (fc *FrameCache) getFilteredCacheItemValues(cacheItemValues []CacheItemValue, queryModel *frame.QueryModel, plotAbsoluteReverse ...bool) []CacheItemValue {
	var filteredCacheItemValues []CacheItemValue
	var plotReverse bool
	if len(plotAbsoluteReverse) > 0 && plotAbsoluteReverse[0] {
		plotReverse = true
	}
	// Set the Query Model's TimeSeriesTimeRangeFrom and TimeSeriesTimeRangeTo properties
	if queryModel.AbsoluteTimeSelected { // Absolute time
		for _, cacheItem := range cacheItemValues {
			if plotReverse { // Are we plotting the reverse absolute time for a relative timeline?
				if cacheItem.ValueKey.Before(queryModel.CurrentTime.Add(time.Duration(queryModel.Mintime))) {
					filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
				}
			} else {
				if cacheItem.ValueKey.After(queryModel.CurrentTime) {
					filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
				}
			}
		}
	} else { // Relative time
		for _, cacheItem := range cacheItemValues {
			if cacheItem.ValueKey.After(queryModel.CurrentTime) {
				filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
			}
		}
	}
	return filteredCacheItemValues
}

func (fc *FrameCache) GetFrame(qm *frame.QueryModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "GetFrame")
	var (
		resultframe             *data.Frame
		filteredCacheItemValues []CacheItemValue
	)
	cacheItemValues, err := fc.getCacheItemValues(qm.CacheKey())
	if err != nil {
		logger.Debug("cache item values not obtained", "error", err, "key", string(qm.CacheKey()))
	}
	if len(cacheItemValues) > 0 {
		filteredCacheItemValues = fc.getFilteredCacheItemValues(cacheItemValues, qm, plotAbsoluteReverse...)
		if len(filteredCacheItemValues) > 0 {
			var matchedCacheItem *CacheItemValue
			if len(plotAbsoluteReverse) > 0 && plotAbsoluteReverse[0] {
				matchedCacheItem = &filteredCacheItemValues[len(filteredCacheItemValues)-1]
			} else {
				matchedCacheItem = &filteredCacheItemValues[0]
			}
			diffInSecs := int(matchedCacheItem.ValueKey.Sub(qm.CurrentTime).Seconds())
			if int(math.Abs(float64(diffInSecs))) <= int(matchedCacheItem.Mintime) {
				qm.Update(&matchedCacheItem.ResponseStatus)
				resultframe = &matchedCacheItem.Value
			}
		} else {
			return nil, errors.New("frame not found in cache in GetFrame()")
		}
	} else {
		return nil, errors.New("frame not found in cache in GetFrame()")
	}
	return resultframe, nil
}

func (fc *FrameCache) SaveFrame(frame *data.Frame, qm *frame.QueryModel) error {
	logger := log.Logger.With("func", "SaveFrame")

	cacheItemValues, err := fc.getCacheItemValues(qm.CacheKey())
	if len(cacheItemValues) > 0 {
		for _, cacheItemValue := range cacheItemValues {
			if cacheItemValue.CurrentTime.Equal(qm.CurrentTime) {
				logger.Debug("cache item already exist", "key", string(qm.CacheKey()))
				return nil
			}
		}
	}
	if err != nil {
		logger.Debug("cache item values not obtained", "error", err, "key", string(qm.CacheKey()))
	}
	cacheItemValue := fc.createCacheItemValue(frame, qm)
	cacheItemValues = append(cacheItemValues, cacheItemValue)
	cacheItemValues = fc.sortCacheItemValuesSlice(cacheItemValues)

	if cacheItemValueBytes, err := json.Marshal(cacheItemValues); err != nil {
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

func (fc *FrameCache) sortCacheItemValuesSlice(cacheItemValues []CacheItemValue) []CacheItemValue {
	// Sort the array
	sort.Slice(cacheItemValues, func(i, j int) bool {
		return cacheItemValues[i].ValueKey.Before(cacheItemValues[j].ValueKey)
	})
	return cacheItemValues
}

func (fc *FrameCache) Reset() {
	fc.cache.Reset()
}

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

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type FrameCache struct {
	cache *fastcache.Cache
}

func NewFrameCache(size int) *FrameCache {
	return &FrameCache{cache: fastcache.New(size * 1024 * 1024)}
}

func (fc *FrameCache) getCacheItemValues(key []byte) ([]typ.CacheItemValue, error) {
	var cacheItemValues []typ.CacheItemValue
	byteCacheItemValues := fc.cache.GetBig(nil, key)
	if byteCacheItemValues == nil {
		return cacheItemValues, errors.New("could not obtain cache item values in getCacheItemValues()")
	} else {
		err := json.Unmarshal(byteCacheItemValues, &cacheItemValues)
		if err != nil {
			return nil, err
		}
	}
	return cacheItemValues, nil
}

func (fc *FrameCache) getFilteredCacheItemValues(cacheItemValues []typ.CacheItemValue, queryModel *typ.QueryModel, plotAbsoluteReverse ...bool) []typ.CacheItemValue {
	var filteredCacheItemValues []typ.CacheItemValue
	var plotReverse bool
	if len(plotAbsoluteReverse) > 0 {
		if plotAbsoluteReverse[0] {
			plotReverse = true
		}
	}
	// Set the Query Model's TimeSeriesTimeRangeFrom and TimeSeriesTimeRangeTo properties
	if queryModel.AbsoluteTimeSelected { // Absolute time
		for _, cacheItem := range cacheItemValues {
			if plotReverse { // Are we plotting the reverse absolute time for a relative timeline?
				if cacheItem.ValueKey.Before(queryModel.TimeSeriesTimeRangeFrom.Add(time.Duration(queryModel.ServerTimeData.MinTime))) {
					filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
				}
			} else {
				if cacheItem.ValueKey.After(queryModel.TimeSeriesTimeRangeFrom) {
					filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
				}
			}
		}
	} else { // Relative time
		for _, cacheItem := range cacheItemValues {
			if cacheItem.ValueKey.After(queryModel.TimeSeriesTimeRangeFrom) {
				filteredCacheItemValues = append(filteredCacheItemValues, cacheItem)
			}
		}
	}
	return filteredCacheItemValues
}

func (fc *FrameCache) GetFrame(qm *typ.QueryModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "GetFrame")
	var (
		resultframe             *data.Frame
		filteredCacheItemValues []typ.CacheItemValue
	)
	cacheItemValues, err := fc.getCacheItemValues(qm.CacheKey())
	if err != nil {
		logger.Info("cache item values not obtained", "error", err)
	}
	if len(cacheItemValues) > 0 {
		if len(plotAbsoluteReverse) > 0 {
			filteredCacheItemValues = fc.getFilteredCacheItemValues(cacheItemValues, qm, plotAbsoluteReverse[0])
		} else {
			filteredCacheItemValues = fc.getFilteredCacheItemValues(cacheItemValues, qm)
		}
		if len(filteredCacheItemValues) > 0 {
			var matchedCacheItem *typ.CacheItemValue
			if len(plotAbsoluteReverse) > 0 {
				matchedCacheItem = &filteredCacheItemValues[len(filteredCacheItemValues)-1]
			} else {
				matchedCacheItem = &filteredCacheItemValues[0]
			}
			diffInSecs := int(matchedCacheItem.ValueKey.Sub(qm.TimeSeriesTimeRangeFrom).Seconds())
			if int(math.Abs(float64(diffInSecs))) <= int(matchedCacheItem.ServerTimeData.MinTime) {
				qm.ServerTimeData = matchedCacheItem.ServerTimeData
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

func (fc *FrameCache) SaveFrame(frame *data.Frame, qm *typ.QueryModel) error {
	logger := log.Logger.With("func", "SaveFrame")

	cacheItemValues, err := fc.getCacheItemValues(qm.CacheKey())
	if err != nil {
		logger.Info("cache item values not obtained", "error", err)
	}
	cacheItemValue := fc.createCacheItemValue(frame, qm)
	cacheItemValues = append(cacheItemValues, cacheItemValue)
	cacheItemValues = fc.sortCacheItemValuesSlice(cacheItemValues)

	if cacheItemValueBytes, err := json.Marshal(cacheItemValues); err != nil {
		return err
	} else {
		fc.cache.SetBig(qm.CacheKey(), cacheItemValueBytes)
	}
	return nil
}

func (fc *FrameCache) createCacheItemValue(frame *data.Frame, qm *typ.QueryModel) typ.CacheItemValue {
	var (
		cacheItemValue typ.CacheItemValue
	)
	if qm.AbsoluteTimeSelected {
		cacheItemValue.ValueKey = qm.TimeSeriesTimeRangeFrom
	} else {
		cacheItemValue.ValueKey = qm.TimeSeriesTimeRangeTo
	}
	cacheItemValue.Value = *frame
	cacheItemValue.ServerTimeData = qm.ServerTimeData
	return cacheItemValue
}

func (fc *FrameCache) sortCacheItemValuesSlice(cacheItemValues []typ.CacheItemValue) []typ.CacheItemValue {
	// Sort the array
	sort.Slice(cacheItemValues, func(i, j int) bool {
		return cacheItemValues[i].ValueKey.Before(cacheItemValues[j].ValueKey)
	})
	return cacheItemValues
}

func (fc *FrameCache) Reset() {
	fc.cache.Reset()
}

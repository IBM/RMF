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

package cache_functions

import (
	"encoding/json"
	"errors"
	"math"
	"sort"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	plugincnfg "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/plugin_config"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type RMFCache struct {
	metricsCache        *fastcache.Cache
	queryModelCache     *fastcache.Cache
	intervalOffsetCache *fastcache.Cache
}

const (
	METRICS = iota
	QUERYMODEL
	INTERVALOFFSET
)

func NewRMFCache(pluginCnfg *plugincnfg.PluginConfig) *RMFCache {
	mCache := fastcache.New(pluginCnfg.CacheSettings.MetricsCacheSizeMB * 1024 * 1024)
	qCache := fastcache.New(pluginCnfg.CacheSettings.QueryModelCacheSizeMB * 1024 * 1024)
	iCache := fastcache.New(pluginCnfg.CacheSettings.IntervalOffsetCacheSizeMB * 1024 * 1024)
	return &RMFCache{
		metricsCache:        mCache,
		queryModelCache:     qCache,
		intervalOffsetCache: iCache,
	}
}

func (c *RMFCache) Close() {
	c.metricsCache.Reset()
	c.queryModelCache.Reset()
	c.intervalOffsetCache.Reset()
}

func (c *RMFCache) Set(key string, val []byte, cacheIdentifier int) {
	if cacheIdentifier == INTERVALOFFSET {
		c.intervalOffsetCache.SetBig([]byte(key), val)
	} else if cacheIdentifier == QUERYMODEL {
		c.queryModelCache.SetBig([]byte(key), val)
	} else {
		c.metricsCache.SetBig([]byte(key), val)
	}
}

func (c *RMFCache) Get(key string, cacheIdentifier int) []byte {
	var result []byte
	if cacheIdentifier == INTERVALOFFSET {
		return c.intervalOffsetCache.GetBig(result, []byte(key))
	} else if cacheIdentifier == QUERYMODEL {
		return c.queryModelCache.GetBig(result, []byte(key))
	} else {
		return c.metricsCache.GetBig(result, []byte(key))
	}
}

func (c *RMFCache) Delete(key string, cacheIdentifier int) error {
	byteKey := []byte(key)
	if cacheIdentifier == INTERVALOFFSET {
		c.intervalOffsetCache.Del(byteKey)
	} else if cacheIdentifier == QUERYMODEL {
		c.queryModelCache.Del(byteKey)
	} else {
		c.metricsCache.Del(byteKey)
	}
	return nil
}

func (c *RMFCache) SaveQueryModel(queryModel *typ.QueryModel) error {
	cacheKey := queryModel.UniquePath
	if qm, err := json.Marshal(queryModel); err != nil {
		return err
	} else {
		c.Set(cacheKey, qm, QUERYMODEL)
	}
	return nil
}

func (c *RMFCache) GetQueryModel(key string) (*typ.QueryModel, error) {
	var qm typ.QueryModel
	uqm := c.Get(key, QUERYMODEL)
	err := json.Unmarshal(uqm, &qm)
	if err != nil {
		return nil, err
	}
	return &qm, nil
}

func (c *RMFCache) getCacheItemValues(key string) ([]typ.CacheItemValue, error) {
	var cacheItemValues []typ.CacheItemValue
	byteCacheItemValues := c.Get(key, METRICS)
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

func (c *RMFCache) getFilteredCacheItemValues(cacheItemValues []typ.CacheItemValue, queryModel *typ.QueryModel, plotAbsoluteReverse ...bool) []typ.CacheItemValue {
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
				if cacheItem.ValueKey.Before(queryModel.TimeSeriesTimeRangeFrom.Add(time.Duration(queryModel.ServerTimeData.ServiceCallInterval))) {
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

func (c *RMFCache) GetFrame(qm *typ.QueryModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "GetFrame")
	var (
		resultframe *data.Frame
		cacheKey    string
		// err                     error
		filteredCacheItemValues []typ.CacheItemValue
	)
	cacheKey = c.getFrameCacheKey(qm)
	cacheItemValues, err := c.getCacheItemValues(cacheKey)
	if err != nil {
		logger.Info("cache item values not obtained", "error", err)
	}
	if len(cacheItemValues) > 0 {
		if len(plotAbsoluteReverse) > 0 {
			filteredCacheItemValues = c.getFilteredCacheItemValues(cacheItemValues, qm, plotAbsoluteReverse[0])
		} else {
			filteredCacheItemValues = c.getFilteredCacheItemValues(cacheItemValues, qm)
		}
		if len(filteredCacheItemValues) > 0 {
			var matchedCacheItem *typ.CacheItemValue
			if len(plotAbsoluteReverse) > 0 {
				matchedCacheItem = &filteredCacheItemValues[len(filteredCacheItemValues)-1]
			} else {
				matchedCacheItem = &filteredCacheItemValues[0]
			}
			diffInSecs := int(matchedCacheItem.ValueKey.Sub(qm.TimeSeriesTimeRangeFrom).Seconds())
			if int(math.Abs(float64(diffInSecs))) <= int(matchedCacheItem.ServerTimeData.ServiceCallInterval) {
				qm.ServerTimeData = matchedCacheItem.ServerTimeData
				resultframe = &matchedCacheItem.Value
			}
		}
	} else {
		return nil, errors.New("frame not found in cache in GetFrame()")
	}
	return resultframe, nil
}

func (c *RMFCache) SaveFrame(frame *data.Frame, qm *typ.QueryModel) error {
	logger := log.Logger.With("func", "SaveFrame")

	cacheKey := c.getFrameCacheKey(qm)
	cacheItemValues, err := c.getCacheItemValues(cacheKey)
	if err != nil {
		logger.Info("cache item values not obtained", "error", err)
	}
	cacheItemValue := c.createCacheItemValue(frame, qm)
	cacheItemValues = append(cacheItemValues, cacheItemValue)
	cacheItemValues = c.sortCacheItemValuesSlice(cacheItemValues)

	if cacheItemValueBytes, err := json.Marshal(cacheItemValues); err != nil {
		return err
	} else {
		c.Set(cacheKey, cacheItemValueBytes, METRICS)
	}
	return nil
}

func (c *RMFCache) createCacheItemValue(frame *data.Frame, qm *typ.QueryModel) typ.CacheItemValue {
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

func (c *RMFCache) sortCacheItemValuesSlice(cacheItemValues []typ.CacheItemValue) []typ.CacheItemValue {
	// Sort the array
	sort.Slice(cacheItemValues, func(i, j int) bool {
		return cacheItemValues[i].ValueKey.Before(cacheItemValues[j].ValueKey)
	})
	return cacheItemValues
}

func (c *RMFCache) DeleteFrame(qm *typ.QueryModel) error {
	cacheKey := c.getFrameCacheKey(qm)
	return c.Delete(cacheKey, METRICS)
}

func (c *RMFCache) GetIntervalAndOffset(cacheKey string) (typ.IntervalOffset, error) {
	var resultIntervalOffset typ.IntervalOffset
	if byteIntervalOffset := c.Get(cacheKey, INTERVALOFFSET); byteIntervalOffset == nil {
		return resultIntervalOffset, errors.New("could not get interval and offset from cache in GetIntervalAndOffset()")
	} else {
		err := json.Unmarshal(byteIntervalOffset, &resultIntervalOffset)
		if err != nil {
			return resultIntervalOffset, err
		}
	}
	return resultIntervalOffset, nil
}

func (c *RMFCache) SaveIntervalAndOffset(cacheKey string, intervalOffset typ.IntervalOffset) error {
	if byteIntervalOffset, err := json.Marshal(intervalOffset); err != nil {
		return err
	} else {
		c.Set(cacheKey, byteIntervalOffset, INTERVALOFFSET)
	}
	return nil
}

func (c *RMFCache) getFrameCacheKey(qm *typ.QueryModel) string {
	resultCacheKey := qm.SelectedResource.Value + "/" + qm.Datasource.Uid // + "/" + c.formatTime(qm.TimeSeriesTimeRangeFrom)
	return resultCacheKey
}

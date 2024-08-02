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
	"sync"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/query"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
)

const UpdateInterval = 15 * time.Minute
const DefaultMinTime = 100
const DefaultTimeOffset = 0

var DefaultTimeData = typ.IntervalTimeData{MinTime: DefaultMinTime, TimeOffset: DefaultTimeOffset}

type DdsOptions struct {
	mutex    sync.RWMutex
	timeData typ.IntervalTimeData
}

func (ds *RMFDatasource) watchDdsOptions() {
	logger := log.Logger.With("func", "watchDdsOptions")
	ticker := time.NewTicker(UpdateInterval)
	defer ds.waitGroup.Done()
	defer ticker.Stop()
	logger.Debug("DDS options watcher started")
	for {
		select {
		case <-ticker.C:
			ds.getDdsTimeData()
		case <-ds.stopChan:
			logger.Debug("DDS options watcher stopped")
			return
		}
	}
}

func (ds *RMFDatasource) getCachedDdsTimeData() typ.IntervalTimeData {
	ds.ddsOptions.mutex.RLock()
	timeData := ds.ddsOptions.timeData
	ds.ddsOptions.mutex.RUnlock()
	if timeData.MinTime == 0 {
		return ds.getDdsTimeData()
	}
	return timeData
}

func (ds *RMFDatasource) getDdsTimeData() typ.IntervalTimeData {
	logger := log.Logger.With("func", "getDdsTimeData")
	latest, err := query.FetchIntervalAndOffset(ds.endpoint)
	if err != nil {
		logger.Error("unable to retrieve DDS timezone offset", "error", err)
		return DefaultTimeData
	}
	ds.ddsOptions.mutex.Lock()
	defer ds.ddsOptions.mutex.Unlock()
	ds.ddsOptions.timeData = latest
	return ds.ddsOptions.timeData
}

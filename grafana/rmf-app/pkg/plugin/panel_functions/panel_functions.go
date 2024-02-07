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

package panel_functions

import (
	"fmt"
	"math"
	"strings"
	"time"

	cf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache_functions"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"

	"github.com/google/uuid"
)

type PanelFunctions struct {
}

func (pf *PanelFunctions) GetRMFPanelIDs(qm *typ.QueryModel, pCtx backend.PluginContext) (string, string) {
	rmfPanelId := strings.ReplaceAll(qm.RMFPanelId, "-", "")
	if pCtx.User != nil {
		rmfPanelId = pCtx.User.Login + "/" + rmfPanelId
	} else { // TO DO: Need to check whether we can get a userid when User object is nil.
		rmfPanelId = "admin" + "/" + rmfPanelId
	}
	uniquePath := rmfPanelId + "/" + qm.RefID

	// Replace all strings in the TSO Id having #, $ and @ symbols because channel path cannot contain these characters.
	uniquePath = replaceSpecialCharacters(uniquePath)

	return rmfPanelId, uniquePath
}

func replaceSpecialCharacters(uniquePath string) string {
	replacer := strings.NewReplacer("#", "-", "$", "-", "@", "-")
	return replacer.Replace(uniquePath)
}

func (pf *PanelFunctions) SaveQueryModelInCache(rmfCache *cf.RMFCache, qm *typ.QueryModel) {
	if err := rmfCache.SaveQueryModel(qm); err != nil {
		log.DefaultLogger.Info(fmt.Sprintf("could not save queryModel in cache (SaveQueryModelInCache()): details: %v ", err), nil)
	}
}

func (pf *PanelFunctions) GetMatchedQueryModelFromCache(rmfCache *cf.RMFCache, panelUniquePath string) *typ.QueryModel {
	if queryModel, err := rmfCache.GetQueryModel(panelUniquePath); err != nil {
		return nil
	} else {
		return queryModel
	}
}

func (pf *PanelFunctions) SaveFrameInCache(rmfCache *cf.RMFCache, qm *typ.QueryModel, frame *data.Frame) error {
	if err := rmfCache.SaveFrame(frame, qm); err != nil {
		return err
	}
	return nil
}

func (pf *PanelFunctions) GetFrameFromCache(cache *cf.RMFCache, qm *typ.QueryModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	var (
		frame *data.Frame
		err   error
	)
	if len(plotAbsoluteReverse) > 0 {
		frame, err = cache.GetFrame(qm, plotAbsoluteReverse[0])
	} else {
		frame, err = cache.GetFrame(qm)
	}
	if frame == nil || err != nil {
		return nil, fmt.Errorf("frame not found in cache in GetFrameFromCache()")
	} else {
		return frame, nil
	}
}

func (pf *PanelFunctions) GetIntervalAndOffsetFromCache(rmfCache *cf.RMFCache, cacheKey string) (typ.IntervalOffset, error) {
	return rmfCache.GetIntervalAndOffset(cacheKey)
}

func (pf *PanelFunctions) SaveIntervalAndOffsetInCache(rmfCache *cf.RMFCache, cacheKey string, intervalOffset typ.IntervalOffset) error {
	if err := rmfCache.SaveIntervalAndOffset(cacheKey, intervalOffset); err != nil {
		return err
	} else {
		return nil
	}
}

func (pf *PanelFunctions) UpdateServiceCallInProgressStatus(rmfCache *cf.RMFCache, qm *typ.QueryModel, isServiceCallInProgress bool) error {
	cacheKey := qm.UniquePath
	if queryModel, err := rmfCache.GetQueryModel(cacheKey); err != nil {
		return fmt.Errorf("could not update service call progress status in UpdateServiceCallInProgressStatus()")
	} else {
		queryModel.ServiceCallInProgress = isServiceCallInProgress
		if err := rmfCache.SaveQueryModel(queryModel); err != nil {
			return fmt.Errorf("could not save queryModel in UpdateServiceCallInProgressStatus()")
		}
	}
	return nil
}

func (pf *PanelFunctions) IsCurrentRequestWithinTimeLimit(rmfCache *cf.RMFCache, qm *typ.QueryModel) bool {
	var requestWithinTimeLimit bool
	cacheKey := qm.UniquePath
	if queryModel, err := rmfCache.GetQueryModel(cacheKey); err != nil {
		return false
	} else {
		addedDuration := time.Duration(queryModel.ServerTimeData.ServiceCallInterval) * time.Second
		addedTime := queryModel.TimeSeriesTimeRangeFrom.Add(addedDuration)
		if !queryModel.AbsoluteTimeSelected {
			requestWithinTimeLimit = queryModel.TimeRangeTo.Before(addedTime)
		} else {
			requestWithinTimeLimit = queryModel.TimeRangeFrom.Before(addedTime)
		}
	}
	return requestWithinTimeLimit
}

func (pf *PanelFunctions) PanelRefreshRequired(rmfCache *cf.RMFCache, qm *typ.QueryModel) bool {
	cacheKey := qm.UniquePath
	qModel, err := rmfCache.GetQueryModel(cacheKey)
	if err != nil {
		return false
	} else {
		return qModel.Datasource.Uid != qm.Datasource.Uid ||
			qModel.SelectedQuery != qm.SelectedQuery ||
			qModel.AbsoluteTimeSelected != qm.AbsoluteTimeSelected ||
			qModel.SelectedVisualisationType != qm.SelectedVisualisationType ||
			(qModel.AbsoluteTimeSelected && (qModel.TimeRangeFrom != qm.TimeRangeFrom || qModel.TimeRangeTo != qm.TimeRangeTo))
	}
}

func (pf *PanelFunctions) DeleteQueryModel(rmfCache *cf.RMFCache, uniquePath string) {
	if err := rmfCache.Delete(uniquePath, cf.QUERYMODEL); err != nil {
		log.DefaultLogger.Info(fmt.Sprintf("queryModel not in cache: DeleteQueryModel(): details: %v ", err), nil)
	}
}

func (pf *PanelFunctions) GetDataRangeAndTimeRange(rmfCache *cf.RMFCache, qm *typ.QueryModel) (float64, time.Time, time.Time) {
	cacheKey := qm.UniquePath
	if queryModel, err := rmfCache.GetQueryModel(cacheKey); err != nil {
		return 0, qm.TimeRangeFrom, qm.TimeRangeTo
	} else {
		return queryModel.ServerTimeData.ServiceCallInterval, queryModel.TimeSeriesTimeRangeFrom, queryModel.TimeSeriesTimeRangeTo
	}
}

func (pf *PanelFunctions) HasIntervalExceededEndTime(queryModel *typ.QueryModel) bool {
	addedDuration := time.Duration(queryModel.ServerTimeData.ServiceCallInterval) * time.Second
	addedTime := queryModel.TimeSeriesTimeRangeFrom.Add(addedDuration)
	return addedTime.After(queryModel.TimeRangeTo)
}

func (pf *PanelFunctions) GetNewGuid() string {
	newGuid := uuid.New()
	return strings.Replace(newGuid.String(), "-", "", -1)
}

func (pf *PanelFunctions) GetIterationsForRelativePlotting(qm *typ.QueryModel) (int, error) {
	currentTimeUTC := time.Now().UTC()
	difference := qm.TimeSeriesTimeRangeTo.Sub(currentTimeUTC)
	differenceInSecs := int(math.Abs(difference.Seconds()))
	if qm.ServerTimeData.ServiceCallInterval == 0 {
		return 0, fmt.Errorf("ServiceCallInterval must not be zero in GetIterationsForRelativePlotting()")
	}
	result := int(differenceInSecs / int(qm.ServerTimeData.ServiceCallInterval))
	if result == 0 {
		result = 1 //We need to invoke the svc atleast once. So return 1.
	}
	return result, nil
}

func (pf *PanelFunctions) GetIterationsForReverseAbsPlotting(qm *typ.QueryModel) (int, error) {
	difference := qm.TimeSeriesTimeRangeTo.Sub(qm.ServerTimeData.LocalPrevTime)
	differenceInSecs := int(math.Abs(difference.Seconds()))
	if qm.ServerTimeData.ServiceCallInterval == 0 {
		return 0, fmt.Errorf("ServiceCallInterval must not be zero in GetIterationsForRelativePlotting()")
	}
	result := int(differenceInSecs / int(qm.ServerTimeData.ServiceCallInterval))
	if result == 0 {
		result = 1 //We need to invoke the svc atleast once. So return 1.
	} else if result >= 1 {
		// Adjust local previous time so that it not more than ServiceCallInterval.
		// Meaning the difference between TimeSeriesTimeRangeTo and localPrevTime must not be greater than ServiceCallInterval seconds
		qm.ServerTimeData.LocalPrevTime = qm.TimeSeriesTimeRangeTo.Add(-1 * time.Duration(qm.ServerTimeData.ServiceCallInterval) * time.Second)
	}
	return result, nil
}

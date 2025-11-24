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
	"context"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

func (ds *RMFDatasource) getFrame(r *dds.Request, wide bool) (*data.Frame, error) {
	key := cache.FrameKey(r, wide)
	result, err, _ := ds.single.Do(string(key), func() (interface{}, error) {
		ddsResponse, err := ds.ddsClient.GetByRequest(r)
		if err != nil {
			return nil, err
		}
		headers := ds.ddsClient.GetCachedHeaders()
		f, err := frame.Build(ddsResponse, headers, wide)
		if err != nil {
			return nil, err
		}
		return f, nil
	})
	if result != nil {
		return result.(*data.Frame), err
	} else {
		return nil, err
	}
}

// getStep calculates the most appropriate time series step.
// There's no ideal solution. We assume that it aligns with one hour.
// If it doesn't, streaming will still work, but some queries will miss cache.
func getStep(mintime time.Duration, limit time.Duration) time.Duration {
	n := 3600 / int(mintime.Seconds())
	step := time.Hour // The maximum possible
	for i := 1; i <= n; i++ {
		if n%i == 0 && time.Duration(i)*mintime >= limit {
			step = time.Duration(i) * mintime
			break
		}
	}
	return step
}

// getCachedTSFrames fetches all sequentional time series frames from cache and merges it into one frame.
// It has to syncronize time series field set passed in `fields`.
func (ds *RMFDatasource) getCachedTSFrames(r *dds.Request, stop time.Time, step time.Duration, fields frame.SeriesFields) (*data.Frame, time.Duration, error) {
	var (
		f    *data.Frame
		jump time.Duration
		err  error
	)
	// Create a copy of the original request - don't alter it
	cr := dds.NewRequest(r.Resource, r.TimeRange.From, r.TimeRange.To, step)
	for r.TimeRange.To.Before(stop) {
		next := ds.frameCache.Get(cr, true)
		if next == nil {
			break
		}
		frame.SyncFieldNames(fields, next, r.TimeRange.To)
		f, err = frame.MergeInto(f, next)
		if err != nil {
			return nil, jump, err
		}
		cr.Add(step)
		jump += step
	}
	return f, jump, err
}

func (ds *RMFDatasource) getCachedReportFrames(r *dds.Request) *data.Frame {
	return ds.frameCache.Get(r, true)
}

func (ds *RMFDatasource) setCachedReportFrames(f *data.Frame, r *dds.Request) {
	logger := log.Logger.With("func", "setCachedReportFrames")
	if err := ds.frameCache.Set(f, r, true); err != nil {
		logger.Error("failed to save data in cache", "request", r.String(), "reason", err)
	}
}

func (ds *RMFDatasource) serveTSFrame(ctx context.Context, sender *backend.StreamSender, fields frame.SeriesFields, r *dds.Request, hist bool) error {
	logger := log.Logger.With("func", "serveTSFrame")
	var f *data.Frame
	var err error

	for {
		if err = ctx.Err(); err != nil {
			return err
		}
		if !hist {
			d := time.Until(r.TimeRange.To.Add(SdsDelay))
			logger.Debug("sleeping", "request", r.String(), "duration", d.String())
			time.Sleep(d)
		}
		logger.Debug("executing query", "request", r.String())
		f, err = ds.getFrame(r, true)
		if err != nil {
			logger.Error("failed to get data", "request", r.String(), "reason", err)
			f = frame.NoDataFrame(r.TimeRange.To)
		} else {
			if !hist {
				t, ok := f.Fields[0].At(0).(time.Time)
				if !ok || t.Before(r.TimeRange.To) {
					logger.Debug("mintime is not ready yet")
					time.Sleep(SdsDelay)
					continue
				}
			}
			if err := ds.frameCache.Set(f, r, true); err != nil {
				logger.Error("failed to save data in cache", "request", r.String(), "reason", err)
			}
		}
		break
	}
	// No data was returned by DDS yet by any previous request
	if len(f.Fields) < 2 && len(fields) == 0 {
		return nil
	}
	frame.SyncFieldNames(fields, f, r.TimeRange.To)
	if err := sender.SendFrame(f, data.IncludeAll); err != nil {
		return err
	}
	return nil
}

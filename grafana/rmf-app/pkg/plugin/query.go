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
	"github.com/google/uuid"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/live"
)

func (ds *RMFDatasource) getFrame(r *dds.Request, wide bool) (*data.Frame, error) {
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
}

func (ds *RMFDatasource) getFrameCached(r *dds.Request, wide bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "getFrameCached")
	key := cache.FrameKey(r, wide)

	result, err, _ := ds.single.Do(string(key), func() (interface{}, error) {
		f := ds.frameCache.Get(r, wide)
		// Fetch from the DDS Server and then save to cache if required.
		if f == nil {
			f, err := ds.getFrame(r, wide)
			if err != nil {
				return nil, err
			} else {
				// Probably the requested mintime is not ready yet, don't cache it
				// We still can use it in non-timeseries views
				t, ok := f.Fields[0].At(0).(time.Time)
				if !ok || t.Before(r.TimeRange.To) {
					return f, nil
				}
				if err = ds.frameCache.Set(f, r, wide); err != nil {
					return nil, err
				}
			}
			return f, nil
		} else {
			logger.Debug("cached value exists", "key", key)
			logger.Warn("cached value exists", "key", r.String())
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

type RequestParams struct {
	Resource struct {
		Value string `json:"value"`
	} `json:"selectedResource"`
	AbsoluteTime bool   `json:"absoluteTimeSelected"`
	VisType      string `json:"selectedVisualisationType"`
}

func (ds *RMFDatasource) getFirstTSFrame(params *RequestParams, tr backend.TimeRange, step time.Duration) (*data.Frame, error) {
	var (
		f   *data.Frame
		err error
	)
	res := params.Resource.Value
	absolute := params.AbsoluteTime

	r := dds.NewRequest(res, tr.From, tr.From, step)
	fields := frame.SeriesFields{}
	for !absolute && r.TimeRange.To.Before(time.Now().Add(-SdsDelay)) || absolute && r.TimeRange.To.Before(tr.To) {
		next := ds.frameCache.Get(r, true)
		if next == nil {
			break
		}
		frame.SyncFieldNames(fields, next, r.TimeRange.To)
		f, err = frame.MergeInto(f, next)
		if err != nil {
			return nil, err
		}
		r.Add(step)
	}
	if f == nil {
		f = frame.TaggedFrame(tr.From, "No data yet...")
	}
	channel := live.Channel{
		Scope:     live.ScopeDatasource,
		Namespace: ds.uid,
		Path:      uuid.NewString(),
	}
	cachedChannel := cache.Channel{
		Resource:  res,
		TimeRange: backend.TimeRange{From: r.TimeRange.From, To: tr.To},
		Absolute:  absolute,
		Step:      step,
		Fields:    fields,
	}
	if err = ds.channelCache.Set(channel.Path, &cachedChannel); err != nil {
		return nil, err
	}
	f.SetMeta(&data.FrameMeta{Channel: channel.String()})
	return f, nil
}

func (ds *RMFDatasource) serveNextTSFrame(ctx context.Context, sender *backend.StreamSender, fields frame.SeriesFields, r *dds.Request, hist bool) error {
	logger := log.Logger.With("func", "queryNextTSFrame")
	var f *data.Frame
	var err error

	for {
		if err = ctx.Err(); err != nil {
			return err
		}
		if !hist {
			d := time.Until(r.TimeRange.To.Add(SdsDelay))
			time.Sleep(d)
		}
		logger.Debug("executing query", "request", r.String())
		f, err = ds.getFrameCached(r, true)
		if err != nil {
			logger.Error("failed to get data", "request", r.String(), "reason", err)
			f = frame.NoDataFrame(r.TimeRange.To)
		}
		if !hist {
			t, ok := f.Fields[0].At(0).(time.Time)
			if !ok || t.Before(r.TimeRange.To) {
				logger.Debug("mintime is not ready yet")
				time.Sleep(SdsDelay)
				continue
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

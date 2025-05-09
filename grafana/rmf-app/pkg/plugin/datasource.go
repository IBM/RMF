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
	"encoding/json"
	"errors"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/live"
	"golang.org/x/sync/singleflight"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/dds"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
)

// Make sure RMFDatasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin
// in runtime. Plugin should implement only interfaces which are required for a
// particular task.
var (
	_ instancemgmt.InstanceDisposer = (*RMFDatasource)(nil)
	_ backend.CheckHealthHandler    = (*RMFDatasource)(nil)
	_ backend.CallResourceHandler   = (*RMFDatasource)(nil)
	_ backend.QueryDataHandler      = (*RMFDatasource)(nil)
	_ backend.StreamHandler         = (*RMFDatasource)(nil)
)

const SdsDelay = 5 * time.Second
const TimeSeriesType = "TimeSeries"

type RMFDatasource struct {
	uid       string
	name      string
	cache     *cache.Cache
	ddsClient *dds.Client
	single    singleflight.Group
}

// NewRMFDatasource creates a new instance of the RMF datasource.
func NewRMFDatasource(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	logger := log.Logger.With("func", "NewRMFDatasource")
	ds := &RMFDatasource{uid: settings.UID, name: settings.Name}
	config, err := ds.getConfig(settings)
	if err != nil {
		logger.Error("failed to get config", "error", err)
		return nil, err
	}
	// nolint:contextcheck
	ds.ddsClient = dds.NewClient(config.URL, config.Username, config.Password, config.Timeout,
		config.JSON.TlsSkipVerify, config.JSON.DisableCompression)
	ds.cache = cache.NewFrameCache(config.CacheSize)
	logger.Info("initialized a datasource",
		"uid", settings.UID, "name", settings.Name,
		"url", config.URL, "timeout", config.Timeout, "cacheSize", config.CacheSize,
		"username", config.Username, "tlsSkipVerify", config.JSON.TlsSkipVerify)
	return ds, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewRMFClient factory function.
func (ds *RMFDatasource) Dispose() {
	logger := log.Logger.With("func", "Dispose")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)
	ds.cache.Reset()
	ds.ddsClient.Close()
	logger.Info("disposed datasource", "uid", ds.uid, "name", ds.name)
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (ds *RMFDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (retRes *backend.CheckHealthResult, _ error) {

	logger := log.Logger.With("func", "CheckHealth")

	// Recover from any panic so as to not bring down this backend datasource
	defer func() {
		if r := recover(); r != nil {
			message := log.ErrorWithId(logger, log.InternalError, "recovered from panic", "error", r, "stack", string(debug.Stack())).Error()
			retRes = &backend.CheckHealthResult{Status: backend.HealthStatusError, Message: message}
		}
	}()

	var (
		message string
		status  backend.HealthStatus
	)

	_, err := ds.ddsClient.GetRoot(ctx)
	if err != nil {
		status = backend.HealthStatusError
		if errors.Is(err, dds.ErrUnauthorized) {
			message = "Unauthorized. Make sure the credentials are correct."
		} else if errors.Is(err, dds.ErrParse) {
			message = "Unsupported version of DDS."
		} else {
			message = log.ErrorWithId(logger, log.ConnectionError, "couldn't fetch root info", "error", err).Error()
		}
	} else {
		status = backend.HealthStatusOk
		message = "Data source is working."
	}
	return &backend.CheckHealthResult{Status: status, Message: message}, nil
}

type VariableQueryRequest struct {
	Query string `json:"query"`
}

func (ds *RMFDatasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	logger := log.Logger.With("func", "CallResource")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)
	switch req.Path {
	// FIXME: it's a contained.xml request for M3 resource tree. Re-factor accordingly.
	case "variablequery":
		// Extract the query parameter from the POST request
		jsonStr := string(req.Body)
		varRequest := VariableQueryRequest{}
		err := json.Unmarshal([]byte(jsonStr), &varRequest)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not unmarshal data", "error", err)
		}
		ddsResource := varRequest.Query
		if len(strings.TrimSpace(ddsResource)) == 0 {
			return log.ErrorWithId(logger, log.InputError, "variable query cannot be blank")
		}

		data, err := ds.ddsClient.GetRawContained(ctx, ddsResource)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch data", "query", ddsResource, "error", err)
		} else {
			logger.Debug("executed variable query and got response", "query", ddsResource)
		}
		return sender.Send(&backend.CallResourceResponse{Status: http.StatusOK, Body: data})
	// FIXME: it's a metrics index request. Re-factor accordingly.
	case "autopopulate":
		metricsIndex, err := ds.ddsClient.GetRawIndex(ctx)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch (autopopulate) metrics", "error", err)
		} else {
			logger.Debug("executed autopopulate and got response")
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   metricsIndex,
		})
	default:
		return sender.Send(&backend.CallResourceResponse{Status: http.StatusNotFound, Body: nil})
	}
}

type RequestParams struct {
	Resource struct {
		Value string `json:"value"`
	} `json:"selectedResource"`
	AbsoluteTime bool   `json:"absoluteTimeSelected"`
	VisType      string `json:"selectedVisualisationType"`
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (ds *RMFDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (qr *backend.QueryDataResponse, errRet error) {
	logger := log.Logger.With("func", "QueryData")
	qr = backend.NewQueryDataResponse()

	// Recover from any panic to prevent bringing down this backend datasource.
	defer func() {
		if r := recover(); r != nil {
			// Assign error to the first incomplete query: that's where the panic occurred.
			err := log.ErrorWithId(logger, log.InternalError, "recovered from panic", "error", r, "stack", string(debug.Stack()))
			for _, query := range req.Queries {
				if _, ok := qr.Responses[query.RefID]; !ok {
					qr.Responses[query.RefID] = backend.DataResponse{
						Status: backend.StatusInternal,
						Error:  err,
					}
					return
				}
			}
			qr = nil
			errRet = err
		}
	}()

	type ResponseWithId struct {
		refId    string
		response *backend.DataResponse
	}
	var wg sync.WaitGroup
	responseChan := make(chan ResponseWithId, len(req.Queries))

	for _, query := range req.Queries {
		wg.Add(1)

		go func(q backend.DataQuery) {
			defer wg.Done()

			var response *backend.DataResponse
			var params RequestParams
			err := json.Unmarshal(q.JSON, &params)

			if err != nil {
				response = &backend.DataResponse{Status: backend.StatusBadRequest, Error: err}
			} else {
				mintime := ds.ddsClient.GetCachedMintime()
				if params.VisType == TimeSeriesType {
					// Initialize time series stream
					from := q.TimeRange.From
					f := frame.TaggedFrame(from, "No data yet...")
					path := encodeChannelPath(params.Resource.Value, from, q.TimeRange.To, params.AbsoluteTime, q.Interval)
					channel := live.Channel{
						Scope:     live.ScopeDatasource,
						Namespace: req.PluginContext.DataSourceInstanceSettings.UID,
						Path:      path,
					}
					f.SetMeta(&data.FrameMeta{Channel: channel.String()})
					response = &backend.DataResponse{Frames: data.Frames{f}}
				} else {
					// Query non-timeseries data
					r := dds.NewRequest(params.Resource.Value, q.TimeRange.From, q.TimeRange.To, mintime)
					response = &backend.DataResponse{}
					// FIXME: doesn't it need to be cached?
					if newFrame, err := ds.getFrame(ctx, r, false); err != nil {
						// nolint:errorlint
						if cause, ok := errors.Unwrap(err).(*dds.Message); ok {
							response.Error = cause
							response.Status = backend.StatusBadRequest
						} else {
							response.Error = log.FrameErrorWithId(logger, err)
							response.Status = backend.StatusInternal
						}
					} else if newFrame != nil {
						response.Frames = append(response.Frames, newFrame)
					}
				}
			}
			responseChan <- ResponseWithId{refId: q.RefID, response: response}
		}(query)

	}

	go func() {
		wg.Wait()
		close(responseChan)
	}()
	for respWithId := range responseChan {
		qr.Responses[respWithId.refId] = *respWithId.response
	}
	return qr, nil
}

// RunStream is called once for any open channel. Results are shared with everyone
// subscribed to the same channel.
func (ds *RMFDatasource) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	logger := log.Logger.With("func", "RunStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	res, from, to, absolute, interval, err := decodeChannelPath(string(req.Path))
	if err != nil {
		logger.Error("unable to decode channel path", "err", err)
		return nil
	}

	// Calculate the most appropriate interval length, i.e. time series step.
	// There's no ideal solution. We assume that it aligns with one hour.
	// If it doesn't, streaming will still work, but some queries will miss cache.
	mintime := ds.ddsClient.GetCachedMintime()
	n := 3600 / int(mintime.Seconds())
	step := time.Hour // The maximum possible
	for i := 1; i <= n; i++ {
		if n%i == 0 && time.Duration(i)*mintime >= interval {
			step = time.Duration(i) * mintime
			break
		}
	}
	logger.Debug("starting streaming", "step", step.String(), "interval", interval.String(), "path", req.Path)

	r := dds.NewRequest(res, from, from, step)
	seriesFields := frame.SeriesFields{}

	// Stream historical part of time series
	for {
		if err := ctx.Err(); err != nil {
			logger.Info("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		if !absolute && r.TimeRange.To.After(time.Now().Add(-SdsDelay)) || absolute && r.TimeRange.To.After(to) {
			logger.Debug("finished with historical data", "request", r.String(), "path", req.Path)
			break
		}
		logger.Debug("executing query", "request", r.String())
		f, err := ds.getFrameCached(ctx, r, true)
		if err != nil {
			logger.Error("failed to get data", "request", r.String(), "reason", err, "path", req.Path)
			f = frame.NoDataFrame(r.TimeRange.To)
		}
		// No data was returned by DDS yet by this and any previous request
		if len(f.Fields) < 2 && len(seriesFields) == 0 {
			r.Add(step)
			continue
		}
		frame.SyncFieldNames(seriesFields, f, r.TimeRange.To)
		if err := sender.SendFrame(f, data.IncludeAll); err != nil {
			logger.Info("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		r.Add(step)
	}
	if !absolute {
		// Stream live data as it's being collected
		for {
			if err := ctx.Err(); err != nil {
				logger.Info("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			d := time.Until(r.TimeRange.To.Add(SdsDelay))
			logger.Debug("waiting for the next mintime", "duration", d.String(), "path", req.Path)
			time.Sleep(d)

			f, err := ds.getFrameCached(ctx, r, true)
			if err != nil {
				logger.Error("failed to get data", "request", r.String(), "reason", err, "path", req.Path)
				f = frame.NoDataFrame(r.TimeRange.To)
			}

			t, ok := f.Fields[0].At(0).(time.Time)
			if !ok || t.Before(r.TimeRange.To) {
				logger.Debug("mintime is not ready yet", "path", req.Path)
				time.Sleep(SdsDelay)
				continue
			}
			// No data was returned by DDS yet by any previous request
			if len(f.Fields) < 2 && len(seriesFields) == 0 {
				r.Add(step)
				continue
			}
			frame.SyncFieldNames(seriesFields, f, r.TimeRange.To)
			if err := sender.SendFrame(f, data.IncludeAll); err != nil {
				logger.Info("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			r.Add(step)
		}
	}
	if len(seriesFields) == 0 {
		// There is no data at all, send a dummy frame without fields to reflect it in UI
		f := data.NewFrame("")
		if err := sender.SendFrame(f, data.IncludeAll); err != nil {
			logger.Info("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
	}
	logger.Info("streaming stopped", "reason", "all the data sent", "path", req.Path)
	return nil
}

func (ds *RMFDatasource) getFrame(ctx context.Context, r *dds.Request, wide bool) (*data.Frame, error) {
	ddsResponse, err := ds.ddsClient.GetByRequest(ctx, r)
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

func (ds *RMFDatasource) getFrameCached(ctx context.Context, r *dds.Request, wide bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "getFrameCached")
	key := cache.Key(r, wide)

	result, err, _ := ds.single.Do(string(key), func() (interface{}, error) {
		f := ds.cache.GetFrame(r, wide)
		// Fetch from the DDS Server and then save to cache if required.
		if f == nil {
			f, err := ds.getFrame(ctx, r, wide)
			if err != nil {
				return nil, err
			} else {
				// Probably the requested mintime is not ready yet, don't cache it
				// We still can use it in non-timeseries views
				t, ok := f.Fields[0].At(0).(time.Time)
				if !ok || t.Before(r.TimeRange.To) {
					return f, nil
				}
				if err = ds.cache.SaveFrame(f, r, wide); err != nil {
					return nil, err
				}
			}
			return f, nil
		} else {
			logger.Debug("cached value exists", "key", key)
		}
		return f, nil
	})
	if result != nil {
		return result.(*data.Frame), err
	} else {
		return nil, err
	}
}

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (ds *RMFDatasource) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	return &backend.SubscribeStreamResponse{Status: backend.SubscribeStreamStatusOK}, nil
}

// PublishStream is called when a client sends a message to the stream.
func (d *RMFDatasource) PublishStream(_ context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	return &backend.PublishStreamResponse{Status: backend.PublishStreamStatusPermissionDenied}, nil
}

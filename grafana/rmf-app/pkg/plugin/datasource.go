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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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

const SdsDelay = 5

type RMFDatasource struct {
	uid          string
	name         string
	channelCache *cache.ChannelCache
	frameCache   *cache.FrameCache
	ddsClient    *dds.Client
	single       singleflight.Group
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
	ds.channelCache = cache.NewChannelCache(ChannelCacheSizeMB)
	ds.frameCache = cache.NewFrameCache(config.CacheSize)
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
	ds.channelCache.Reset()
	ds.frameCache.Reset()
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
			qm, err := frame.NewQueryModel(q)
			if err != nil {
				if errors.Is(err, frame.ErrBlankResource) {
					response = &backend.DataResponse{Status: backend.StatusOK}
				} else {
					response = &backend.DataResponse{Status: backend.StatusBadRequest, Error: err}
				}
			} else {
				// nolint:contextcheck
				qm.TimeOffset = ds.ddsClient.GetCachedTimeOffset()
				// nolint:contextcheck
				qm.Mintime = ds.ddsClient.GetCachedMintime()
				if qm.SelectedVisualisationType == frame.TimeSeriesType {
					response = ds.queryTimeSeries(ctx, req.PluginContext, qm)
				} else {
					// FIXME: it's not actually table data. Just not time series.
					response = ds.queryTableData(ctx, qm)
				}
				if response == nil {
					err = log.ErrorWithId(logger, log.InternalError, "query response is nil")
					response = &backend.DataResponse{Status: backend.StatusInternal, Error: err}
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

func (ds *RMFDatasource) queryTimeSeries(ctx context.Context, pCtx backend.PluginContext, query *frame.QueryModel) *backend.DataResponse {
	logger := log.Logger.With("func", "queryTimeSeries")

	var (
		newFrame     *data.Frame
		err          error
		dataResponse *backend.DataResponse = &backend.DataResponse{}
	)

	setQueryTimeRange(query, false)
	if latestNotReady(query.CurrentTime, query.Mintime) {
		logger.Debug("interval not yet ready, step back", "time", query.CurrentTime.String())
		moveNextPrevious(query, false)
	}
	if newFrame, err = ds.getFrameFromCacheOrServer(ctx, query); err != nil {
		// nolint:errorlint
		if cause, ok := errors.Unwrap(err).(*dds.Message); ok {
			dataResponse.Error = cause
			dataResponse.Status = backend.StatusBadRequest
		} else {
			dataResponse.Error = log.FrameErrorWithId(logger, err)
			dataResponse.Status = backend.StatusInternal
		}
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
		if err := ds.createChannelForStreaming(pCtx, query, newFrame); err != nil {
			dataResponse.Error = err
		}
	}
	return dataResponse
}

func (ds *RMFDatasource) createChannelForStreaming(pCtx backend.PluginContext, query *frame.QueryModel, firstFrame *data.Frame) error {
	channelPath := uuid.New().String()
	channel := live.Channel{
		Scope:     live.ScopeDatasource,
		Namespace: pCtx.DataSourceInstanceSettings.UID,
		Path:      channelPath,
	}
	firstFrame.SetMeta(&data.FrameMeta{Channel: channel.String()})
	query.SeriesFields = frame.SeriesFields{}
	frame.SyncFieldNames(query.SeriesFields, firstFrame, query.CurrentTime)
	return ds.channelCache.SetChannelQuery(channelPath, query)
}

// RunStream is called once for any open channel.  Results are shared with everyone
// subscribed to the same channel.
func (ds *RMFDatasource) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	logger := log.Logger.With("func", "RunStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	var err error
	query, err := ds.channelCache.GetChannelQuery(req.Path)
	if err != nil {
		return err
	}
	logger.Debug("RunStream", "path", req.Path, "query", query.SelectedQuery, "dashboard", query.DashboardUid, "absoluteTime", query.AbsoluteTimeSelected)
	// Stream absolute or relative timeline data
	if query.AbsoluteTimeSelected {
		err = ds.streamDataForAbsoluteRange(ctx, req, sender, query)
	} else {
		err = ds.streamDataForRelativeRange(ctx, req, sender, query)
	}
	return err
}

func (ds *RMFDatasource) streamDataForAbsoluteRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *frame.QueryModel) error {
	var waitTime time.Duration
	logger := log.Logger.With("func", "streamDataForAbsoluteRange")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Set wait time to 1/100th of a second
	waitTime = (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	err := ds.streamDataAbsolute(ctx, req, sender, matchedQueryModel, waitTime)
	if err != nil {
		return err
	}

	return nil
}

func (ds *RMFDatasource) streamDataForRelativeRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *frame.QueryModel) error {
	logger := log.Logger.With("func", "streamDataForRelativeRange")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Set wait time to 'ServiceCallInterval' for relative and 1/100th of a second for historical
	waitTime := (time.Second * time.Duration(matchedQueryModel.Mintime))
	histWaitTime := (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	err := ds.streamDataRelative(ctx, req, sender, matchedQueryModel, &waitTime, &histWaitTime)
	if err != nil {
		return err
	}
	return nil
}

func (ds *RMFDatasource) streamDataAbsolute(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *frame.QueryModel, waitTime time.Duration) error {
	logger := log.Logger.With("func", "streamDataAbsolute")
	var (
		newFrame *data.Frame
		err      error
	)
	histTicker := time.NewTicker(waitTime)
	seriesFields := matchedQueryModel.SeriesFields

	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			logger.Debug("closing stream: Done.", "reason", err, "path", req.Path)
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if matchedQueryModel.TimeRangeExceeded() {
				histTicker.Stop()
				logger.Debug("closing stream", "reason", "finished with historical data", "path", req.Path, "CurrentTime", matchedQueryModel.CurrentTime.String(), "TimeRangeFrom", matchedQueryModel.TimeRangeFrom, "TimeRangeTo", matchedQueryModel.TimeRangeTo.String())
				return nil
			}
			setQueryTimeRange(matchedQueryModel)
			if latestNotReady(matchedQueryModel.CurrentTime, matchedQueryModel.Mintime) {
				logger.Debug("interval not yet ready", "time", matchedQueryModel.CurrentTime.String())
				continue
			}
			// Send new data periodically.
			logger.Debug("executing query", "query", matchedQueryModel.SelectedQuery, "current", matchedQueryModel.CurrentTime, "to", matchedQueryModel.TimeRangeTo)
			if newFrame, err = ds.getFrameFromCacheOrServer(ctx, matchedQueryModel); err != nil {
				moveNextPrevious(matchedQueryModel, true)
				return log.ErrorWithId(logger, log.InternalError, "could not get new frame", "error", err)
			}
			if matchedQueryModel.CurrentTime.Equal(matchedQueryModel.LastTime) {
				moveNextPrevious(matchedQueryModel, true)
				logger.Debug("skip frame due to duplication", "time", matchedQueryModel.CurrentTime.String())
				continue
			}
			matchedQueryModel.LastTime = matchedQueryModel.CurrentTime
			frame.SyncFieldNames(seriesFields, newFrame, matchedQueryModel.CurrentTime)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to send frame", "error", err)
			}
			err = ds.channelCache.SetChannelQuery(req.Path, matchedQueryModel)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
			}
		}
	}
}

func (ds *RMFDatasource) streamDataRelative(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *frame.QueryModel, waitTime *time.Duration, histWaitTime *time.Duration) error {
	logger := log.Logger.With("func", "streamDataRelative")
	var newFrame *data.Frame
	// FIXME: tickers are not suitable for the streaming.
	// Time for the next request should be calculated based on the time of the latest response.
	// Also, requests for historical and current data should be synchronized.
	mainTicker := time.NewTicker(*waitTime)
	histTicker := time.NewTicker(*histWaitTime)
	seriesFields := matchedQueryModel.SeriesFields
	duration := matchedQueryModel.TimeRangeTo.Sub(matchedQueryModel.TimeRangeFrom)

	histQueryModel, err := ds.channelCache.GetChannelQuery(req.Path + "/h")
	if err != nil {
		histQueryModel = matchedQueryModel.Copy()
		histQueryModel.AbsoluteTimeSelected = true
	}

	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			logger.Debug("closing stream: Done.", "reason", err, "path", req.Path)
			// Stop tickers to enable garbage collection of resources
			mainTicker.Stop()
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if histQueryModel.TimeRangeExceeded() {
				histTicker.Stop()
				logger.Debug("finished with historical data", "path", req.Path, "CurrentTime", histQueryModel.CurrentTime.String(), "TimeRangeFrom", matchedQueryModel.TimeRangeFrom.String(), "TimeRangeTo", matchedQueryModel.TimeRangeTo.String())
				continue
			}
			setQueryTimeRange(histQueryModel, true)
			logger.Debug("executing query for historical data", "query", histQueryModel.SelectedQuery, "current", histQueryModel.CurrentTime, "from", histQueryModel.TimeRangeFrom)
			// Fetch the data
			if latestNotReady(histQueryModel.CurrentTime, histQueryModel.Mintime) {
				logger.Debug("interval not yet ready", "time", histQueryModel.CurrentTime.String())
				continue
			}
			if newFrame, err = ds.getFrameFromCacheOrServer(ctx, histQueryModel); err != nil {
				moveNextPrevious(histQueryModel, false)
				return log.ErrorWithId(logger, log.InternalError, "could not get new frame for historical data", "error", err)
			}
			if histQueryModel.CurrentTime.Equal(histQueryModel.LastTime) {
				logger.Debug("skip frame due to duplication", "time", histQueryModel.CurrentTime.String())
				moveNextPrevious(histQueryModel, false)
				continue
			}
			histQueryModel.LastTime = histQueryModel.CurrentTime
			frame.SyncFieldNames(seriesFields, newFrame, histQueryModel.CurrentTime)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to send frame for historical data", "error", err)
			}
			err = ds.channelCache.SetChannelQuery(req.Path+"/h", matchedQueryModel)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
			}
		case <-mainTicker.C:
			var numberOfIterations int
			if numberOfIterations, err = getIterationsForRelativePlotting(matchedQueryModel); err != nil {
				return err
			}
			logger.Debug("executing query for relative data", "query", matchedQueryModel.SelectedQuery, "iterations", numberOfIterations)
			for counter := 0; counter < numberOfIterations; counter++ {
				setQueryTimeRange(matchedQueryModel)
				if latestNotReady(matchedQueryModel.CurrentTime, matchedQueryModel.Mintime) {
					logger.Debug("interval not yet ready", "time", matchedQueryModel.CurrentTime.String())
					continue
				}
				logger.Debug("executing query", "query", matchedQueryModel.SelectedQuery, "current", matchedQueryModel.CurrentTime)
				if newFrame, err = ds.getFrameFromCacheOrServer(ctx, matchedQueryModel); err != nil {
					moveNextPrevious(matchedQueryModel, true)
					return log.ErrorWithId(logger, log.InternalError, "could not get new frame for relative data", "error", err)
				}
				if matchedQueryModel.CurrentTime.Equal(matchedQueryModel.LastTime) {
					moveNextPrevious(matchedQueryModel, true)
					logger.Debug("skip frame due to duplication", "time", matchedQueryModel.CurrentTime.String())
					continue
				}
				matchedQueryModel.LastTime = matchedQueryModel.CurrentTime
				frame.RemoveOldFieldNames(seriesFields, matchedQueryModel.CurrentTime.Add(-duration))
				frame.SyncFieldNames(seriesFields, newFrame, histQueryModel.CurrentTime)
				err = sender.SendFrame(newFrame, data.IncludeAll)
				if err != nil {
					return log.ErrorWithId(logger, log.InternalError, "failed to send frame for relative data", "error", err)
				}
				// Save the query model in cache
				err = ds.channelCache.SetChannelQuery(req.Path, matchedQueryModel)
				if err != nil {
					return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
				}
			}
		}
	}
}

func (ds *RMFDatasource) getFrame(ctx context.Context, queryModel *frame.QueryModel) (*data.Frame, error) {
	path, params := queryModel.GetPathWithParams()
	ddsResponse, err := ds.ddsClient.Get(ctx, path, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to get DDS response: %w", err)
	}
	// nolint:contextcheck
	newFrame, err := frame.Build(ddsResponse, ds.ddsClient.GetCachedHeaders(), queryModel)
	if err != nil {
		return nil, fmt.Errorf("failed to construct frame: %w", err)
	}
	return newFrame, nil
}

func (ds *RMFDatasource) getFrameFromCacheOrServer(ctx context.Context, queryModel *frame.QueryModel) (*data.Frame, error) {
	logger := log.Logger.With("func", "getFrameFromCacheOrServer")
	key := string(queryModel.CacheKey())
	result, err, _ := ds.single.Do(key, func() (interface{}, error) {
		var (
			newFrame *data.Frame
			err      error
		)
		newFrame = ds.frameCache.GetFrame(queryModel)
		// Fetch from the DDS Server and then save to cache if required.
		if newFrame == nil {
			newFrame, err = ds.getFrame(ctx, queryModel)
			if err != nil {
				return nil, err
			} else {
				if err = ds.frameCache.SaveFrame(newFrame, queryModel); err != nil {
					return nil, err
				}
			}
		} else {
			logger.Debug("cached value exist", "key", key)
		}
		return newFrame, nil
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
	logger := log.Logger.With("func", "SubscribeStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	status := backend.SubscribeStreamStatusPermissionDenied
	if ds.channelCache.HasChannelQuery(req.Path) {
		status = backend.SubscribeStreamStatusOK
	}
	return &backend.SubscribeStreamResponse{Status: status}, nil
}

// PublishStream is called when a client sends a message to the stream.
func (d *RMFDatasource) PublishStream(_ context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	logger := log.Logger.With("func", "PublishStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Do not allow publishing at all.
	return &backend.PublishStreamResponse{Status: backend.PublishStreamStatusPermissionDenied}, nil
}

func (ds *RMFDatasource) queryTableData(ctx context.Context, qm *frame.QueryModel) *backend.DataResponse {
	logger := log.Logger.With("func", "queryTableData")
	dataResponse := &backend.DataResponse{}
	// FIXME: doesn't it need to be cached?
	if newFrame, err := ds.getFrame(ctx, qm); err != nil {
		// nolint:errorlint
		if cause, ok := errors.Unwrap(err).(*dds.Message); ok {
			dataResponse.Error = cause
			dataResponse.Status = backend.StatusBadRequest
		} else {
			dataResponse.Error = log.FrameErrorWithId(logger, err)
			dataResponse.Status = backend.StatusInternal
		}
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	return dataResponse
}

func getIterationsForRelativePlotting(qm *frame.QueryModel) (int, error) {
	currentTimeUTC := time.Now().UTC()
	difference := qm.CurrentTime.Sub(currentTimeUTC)
	differenceInSecs := int(math.Abs(difference.Seconds()))
	if qm.Mintime == 0 {
		return 0, errors.New("ServiceCallInterval must not be zero in GetIterationsForRelativePlotting()")
	}
	differenceInSecs -= qm.Mintime / 2
	differenceInSecs -= SdsDelay
	result := int(differenceInSecs / int(qm.Mintime))
	if result == 0 {
		// FIXME: it's not necessarily true.
		result = 1 //We need to invoke the svc at least once. So return 1.
	}
	return result, nil
}

func setQueryTimeRange(queryModel *frame.QueryModel, plotAbsoluteReverse ...bool) {
	var plotReverse bool
	if len(plotAbsoluteReverse) > 0 {
		if plotAbsoluteReverse[0] {
			plotReverse = true
		}
	}

	// Set the Query Model's TimeSeriesTimeRangeFrom and TimeSeriesTimeRangeTo properties
	if queryModel.AbsoluteTimeSelected { // Absolute time
		if queryModel.Mintime == 0 || queryModel.CurrentTime.IsZero() {
			fromTime := queryModel.TimeRangeFrom
			queryModel.CurrentTime = queryModel.AdjustRealtime(fromTime, queryModel.Mintime)
		} else {
			if plotReverse {
				localPrevTime := queryModel.LocalPrev.Add(-1 * queryModel.TimeOffset)
				queryModel.CurrentTime = queryModel.AdjustRealtime(localPrevTime, queryModel.Mintime)
			} else {
				addedTime := queryModel.CurrentTime.Add(time.Duration(time.Second * time.Duration(queryModel.Mintime)))
				queryModel.CurrentTime = queryModel.AdjustRealtime(addedTime, queryModel.Mintime)
			}
		}
	} else { // Relative time
		if queryModel.Mintime == 0 || queryModel.CurrentTime.IsZero() {
			toTime := queryModel.TimeRangeTo
			queryModel.CurrentTime = queryModel.AdjustRealtime(toTime, queryModel.Mintime)
		} else {
			localNextTime := queryModel.LocalNext.Add(-1 * queryModel.TimeOffset)
			queryModel.CurrentTime = queryModel.AdjustRealtime(localNextTime, queryModel.Mintime)
		}
	}
}

func latestNotReady(t time.Time, m int) bool {
	var now time.Time = time.Now()
	return t.Add(time.Second*time.Duration(m/2) + SdsDelay).After(now)
}

func moveNextPrevious(qm *frame.QueryModel, next bool) {
	if next {
		qm.CurrentTime = qm.CurrentTime.Add(time.Duration(qm.Mintime) * time.Second)
		qm.LocalNext = qm.LocalNext.Add(time.Duration(qm.Mintime) * time.Second)
	} else {
		qm.CurrentTime = qm.CurrentTime.Add(-1 * time.Duration(qm.Mintime) * time.Second)
		qm.LocalPrev = qm.LocalPrev.Add(-1 * time.Duration(qm.Mintime) * time.Second)
	}
}

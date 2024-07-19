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

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache"
	framef "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame_functions"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	qf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/query"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	umf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/unmarshal_functions"
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

type RMFDatasource struct {
	endpoint   *typ.DatasourceEndpointModel
	ddsOptions DdsOptions
	stopChan   chan struct{}
	waitGroup  sync.WaitGroup
}

// NewRMFDatasource creates a new instance of the RMF datasource.
func NewRMFDatasource(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	ds := &RMFDatasource{
		stopChan: make(chan struct{}),
	}
	endpoint, err := umf.UnmarshalEndpointModel(settings)
	if err != nil {
		return nil, err
	}
	ds.endpoint = endpoint
	ds.waitGroup.Add(1)
	go ds.watchDdsOptions()
	return ds, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewRMFClient factory function.
func (ds *RMFDatasource) Dispose() {
	logger := log.Logger.With("func", "Dispose")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)
	close(ds.stopChan)
	ds.waitGroup.Wait()
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (ds *RMFDatasource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (retRes *backend.CheckHealthResult, _ error) {

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

	res, err := qf.FetchRootInfo(ds.endpoint)
	if err != nil {
		if errors.As(err, new(*typ.ValueError)) {
			message = "Unsupported version of DDS."
		} else {
			message = log.ErrorWithId(logger, log.ConnectionError, "couldn't fetch root info", "error", err).Error()
		}
		return &backend.CheckHealthResult{
				Status:  backend.HealthStatusError,
				Message: message},
			nil
	}
	if res {
		status = backend.HealthStatusOk
		message = "Data source is working."
	} else {
		status = backend.HealthStatusError
		message = "Unknown error. Please contact your administrator and check the logs."
	}
	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil
}

func (ds *RMFDatasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	logger := log.Logger.With("func", "CallResource")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)
	switch req.Path {
	case "variablequery":
		// Extract the query parameter from the POST request
		jsonStr := string(req.Body)
		varRequest := typ.VariableQueryRequest{}
		err := json.Unmarshal([]byte(jsonStr), &varRequest)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not unmarshal data", "error", err)
		}
		ddsQuery := varRequest.Query
		if len(strings.TrimSpace(ddsQuery)) == 0 {
			return log.ErrorWithId(logger, log.InputError, "variable query cannot be blank")
		}

		// Execute the query and return as body
		data, err := qf.GetVariable(ddsQuery, ds.endpoint)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch xml data", "query", ddsQuery, "error", err)
		} else {
			logger.Debug("executed variable query and got response", "query", ddsQuery, "response", data)
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   data,
		})
	case "autopopulate":
		metricsData, err := qf.FetchMetricsFromIndex(ds.endpoint)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch (autopopulate) metrics", "error", err)
		} else {
			logger.Debug("executed autopopulate and got response")
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   metricsData,
		})
	default:
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusNotFound,
			Body:   nil,
		})
	}
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (ds *RMFDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	logger := log.Logger.With("func", "QueryData")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	queryDataResponse := backend.NewQueryDataResponse()
	for _, query := range req.Queries {
		var response *backend.DataResponse
		ddsQuery, err := umf.UnmarshalQueryModel(req.PluginContext, query)
		if err != nil || ddsQuery.SelectedQuery == "" {
			continue
		}
		timeData := ds.getCachedDdsTimeData()
		ddsQuery.ServerTimeData.TimeOffset = timeData.TimeOffset
		ddsQuery.ServerTimeData.MinTime = timeData.MinTime
		if ddsQuery.SelectedVisualisationType == "TimeSeries" {
			response = ds.queryTimeSeries(req.PluginContext, ddsQuery)
		} else {
			response = ds.queryTableData(ddsQuery)
		}
		if response != nil {
			queryDataResponse.Responses[query.RefID] = *response
		}
	}
	return queryDataResponse, nil
}

func (ds *RMFDatasource) queryTimeSeries(pCtx backend.PluginContext, query *typ.QueryModel) *backend.DataResponse {
	logger := log.Logger.With("func", "queryTimeSeries")

	var (
		newFrame     *data.Frame
		err          error
		dataResponse *backend.DataResponse = &backend.DataResponse{}
	)

	newFrame, err = ds.getFrameFromCacheOrServer(query, false)
	if err != nil {
		dataResponse.Error = log.FrameErrorWithId(logger, err)
	}
	if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
		if err := ds.createChannelForStreaming(pCtx, query, newFrame); err != nil {
			dataResponse.Error = err
		}
	}
	return dataResponse
}

func (ds *RMFDatasource) createChannelForStreaming(pCtx backend.PluginContext, query *typ.QueryModel, frame *data.Frame) error {
	channelPath := uuid.New().String()
	channel := live.Channel{
		Scope:     live.ScopeDatasource,
		Namespace: pCtx.DataSourceInstanceSettings.UID,
		Path:      channelPath,
	}
	frame.SetMeta(&data.FrameMeta{Channel: channel.String()})
	return cache.SetChannelQuery(channelPath, query)
}

// RunStream is called once for any open channel.  Results are shared with everyone
// subscribed to the same channel.
func (ds *RMFDatasource) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	logger := log.Logger.With("func", "RunStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	var err error
	query, err := cache.GetChannelQuery(req.Path)
	if err != nil {
		return err
	}
	// Stream absolute or relative timeline data
	if query.AbsoluteTimeSelected {
		err = ds.streamDataForAbsoluteRange(ctx, req, sender, query)
	} else {
		err = ds.streamDataForRelativeRange(ctx, req, sender, query)
	}
	return err
}

func (ds *RMFDatasource) streamDataForAbsoluteRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
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

func (ds *RMFDatasource) streamDataForRelativeRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
	logger := log.Logger.With("func", "streamDataForRelativeRange")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Set wait time to 'ServiceCallInterval' for relative and 1/100th of a second for historical
	waitTime := (time.Second * time.Duration(matchedQueryModel.ServerTimeData.MinTime))
	histWaitTime := (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	err := ds.streamDataRelative(ctx, req, sender, matchedQueryModel, &waitTime, &histWaitTime)
	if err != nil {
		return err
	}
	return nil
}

func (ds *RMFDatasource) streamDataAbsolute(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime time.Duration) error {
	logger := log.Logger.With("func", "streamDataAbsolute")
	var (
		newFrame *data.Frame
		err      error
	)
	histTicker := time.NewTicker(waitTime)
	seriesFields := framef.SeriesFields{}

	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			logger.Debug("closing stream", "reason", err, "path", req.Path)
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if matchedQueryModel.TimeSeriesTimeRangeTo.After(matchedQueryModel.TimeRangeTo) {
				histTicker.Stop()
				logger.Debug("closing stream", "reason", "finished with historical data", "path", req.Path)
				return nil
			}
			// Send new data periodically.
			logger.Debug("executing query", "query", matchedQueryModel.SelectedQuery)
			if newFrame, err = ds.getFrameFromCacheOrServer(matchedQueryModel); err != nil {
				return log.ErrorWithId(logger, log.InternalError, "could not get new frame", "error", err)
			}
			framef.SyncFieldNames(seriesFields, newFrame, matchedQueryModel.TimeSeriesTimeRangeFrom)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to send frame", "error", err)
			}
			err = cache.SetChannelQuery(req.Path, matchedQueryModel)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
			}
		}
	}
}

func (ds *RMFDatasource) streamDataRelative(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime *time.Duration, histWaitTime *time.Duration) error {
	logger := log.Logger.With("func", "streamDataRelative")
	var newFrame *data.Frame
	mainTicker := time.NewTicker(*waitTime)
	histTicker := time.NewTicker(*histWaitTime)
	seriesFields := framef.SeriesFields{}
	duration := matchedQueryModel.TimeRangeTo.Sub(matchedQueryModel.TimeRangeFrom)

	histQueryModel, err := cache.GetChannelQuery(req.Path + "/h")
	if err != nil {
		histQueryModel = matchedQueryModel.Copy()
		histQueryModel.AbsoluteTimeSelected = true
	}

	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			logger.Debug("closing stream", "reason", err, "path", req.Path)
			// Stop tickers to enable garbage collection of resources
			mainTicker.Stop()
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if histQueryModel.TimeSeriesTimeRangeFrom.Before(matchedQueryModel.TimeRangeFrom) { //TimeRangeFrom
				histTicker.Stop()
				logger.Debug("finished with historical data", "path", req.Path)
				continue
			}
			logger.Debug("executing query for historical data", "query", histQueryModel.SelectedQuery)
			// Fetch the data
			if newFrame, err = ds.getFrameFromCacheOrServer(histQueryModel, true); err != nil {
				return log.ErrorWithId(logger, log.InternalError, "could not get new frame for historical data", "error", err)
			}
			framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to send frame for historical data", "error", err)
			}
			err = cache.SetChannelQuery(req.Path+"/h", matchedQueryModel)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
			}
		case <-mainTicker.C: // Relative Data: Wait for a (usually) larger interval (i.e. gathererInterval) and query DDS
			var numberOfIterations int
			if numberOfIterations, err = getIterationsForRelativePlotting(matchedQueryModel); err != nil {
				return err
			}
			logger.Debug("executing query for relative data", "query", matchedQueryModel.SelectedQuery, "iterations", numberOfIterations)
			for counter := 0; counter < numberOfIterations; counter++ {
				if newFrame, err = ds.getFrameFromCacheOrServer(matchedQueryModel); err != nil {
					return log.ErrorWithId(logger, log.InternalError, "could not get new frame for relative data", "error", err)
				}
				framef.RemoveOldFieldNames(seriesFields, matchedQueryModel.TimeSeriesTimeRangeFrom.Add(-duration))
				framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
				err = sender.SendFrame(newFrame, data.IncludeAll)
				if err != nil {
					return log.ErrorWithId(logger, log.InternalError, "failed to send frame for relative data", "error", err)
				}
				// Save the query model in cache
				err = cache.SetChannelQuery(req.Path, matchedQueryModel)
				if err != nil {
					return log.ErrorWithId(logger, log.InternalError, "failed to save frame in cache", "error", err)
				}
			}
		}
	}
}

func (ds *RMFDatasource) getFrameFromCacheOrServer(queryModel *typ.QueryModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "getFrameFromCacheOrServer")
	var (
		newFrame *data.Frame
		err      error
	)

	// For relative timeseries (forward only) don't look in the cache
	if !queryModel.AbsoluteTimeSelected {
		qf.SetQueryTimeRange(queryModel)
	} else {
		// If frame exists in cache for the query get it from there
		// else get frame from the server through a service call
		if len(plotAbsoluteReverse) > 0 {
			qf.SetQueryTimeRange(queryModel, plotAbsoluteReverse[0])
			newFrame, _ = cache.GetFrame(queryModel, plotAbsoluteReverse[0])
		} else {
			qf.SetQueryTimeRange(queryModel)
			newFrame, _ = cache.GetFrame(queryModel)
		}
	}

	// Fetch from the DDS Server and then save to cache if required.
	if newFrame == nil {
		newFrame, err = qf.GetTimeSeriesFrame(queryModel, ds.endpoint)
		if err != nil {
			return nil, log.FrameErrorWithId(logger, err)
		} else {
			if err = cache.SaveFrame(newFrame, queryModel); err != nil {
				return nil, log.ErrorWithId(logger, log.InternalError, "could not save frame", "error", err)
			}
		}
	}
	return newFrame, nil
}

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (ds *RMFDatasource) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	logger := log.Logger.With("func", "SubscribeStream")
	// Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	status := backend.SubscribeStreamStatusPermissionDenied
	if cache.HasChannelQuery(req.Path) {
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

func (ds *RMFDatasource) queryTableData(query *typ.QueryModel) *backend.DataResponse {
	logger := log.Logger.With("func", "queryTableData")
	dataResponse := &backend.DataResponse{}
	if newFrame, err := qf.GetTableFrame(query, ds.endpoint); err != nil {
		dataResponse.Error = log.FrameErrorWithId(logger, err)
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	return dataResponse
}

func getIterationsForRelativePlotting(qm *typ.QueryModel) (int, error) {
	currentTimeUTC := time.Now().UTC()
	difference := qm.TimeSeriesTimeRangeTo.Sub(currentTimeUTC)
	differenceInSecs := int(math.Abs(difference.Seconds()))
	if qm.ServerTimeData.MinTime == 0 {
		return 0, errors.New("ServiceCallInterval must not be zero in GetIterationsForRelativePlotting()")
	}
	result := int(differenceInSecs / int(qm.ServerTimeData.MinTime))
	if result == 0 {
		result = 1 //We need to invoke the svc atleast once. So return 1.
	}
	return result, nil
}

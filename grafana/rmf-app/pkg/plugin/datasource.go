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
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/live"

	cf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache_functions"
	framef "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame_functions"
	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/log"
	pnlf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/panel_functions"
	plugincnfg "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/plugin_config"
	qf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/query_functions"
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
	Cache         *cf.RMFCache
	EndpointModel *typ.DatasourceEndpointModel
	PluginConfig  *plugincnfg.PluginConfig
}

// NewRMFDatasource creates a new datasource instance
func NewRMFDatasource(_ context.Context, _ backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	// TODO: config plugin via Grafana settings
	pluginCnfg := plugincnfg.NewPluginConfig()
	return &RMFDatasource{
		Cache:        cf.NewRMFCache(pluginCnfg),
		PluginConfig: pluginCnfg,
	}, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewRMFClient factory function.
func (d *RMFDatasource) Dispose() {
	// Clean up datasource instance resources.
	d.Cache.Close()
	d.EndpointModel = nil
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *RMFDatasource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (retRes *backend.CheckHealthResult, _ error) {

	logger := log.Logger.With("func", "CheckHealth")

	// Recover from any panic so as to not bring down this backend datasource
	defer func() {
		if r := recover(); r != nil {
			message := log.ErrorWithId(logger, log.InternalError, "recovered from panic", "error", r, "stack", string(debug.Stack())).Error()
			retRes = &backend.CheckHealthResult{Status: backend.HealthStatusError, Message: message}
		}
	}()

	var (
		unmFns    umf.UnmarshalFunctions
		message   string
		status    backend.HealthStatus
		confQuery qf.ConfigurationQuery
	)

	// Unmarshal the endpoint model
	endpointModel, err := unmFns.UnmarshalEndpointModel(req.PluginContext)
	if err != nil {
		message = log.ErrorWithId(logger, log.InternalError, "could not unmarshal endpointModel", "error", err).Error()
		return &backend.CheckHealthResult{
				Status:  backend.HealthStatusError,
				Message: message},
			nil
	}

	res, err := confQuery.FetchRootInfo(endpointModel)
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

type VariableQueryRequest struct {
	Query string `json:"query"`
}

func (d *RMFDatasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	logger := log.Logger.With("func", "CallResource")
	var unmFns umf.UnmarshalFunctions

	//Unmarshal the endpoint model
	endpointModel, err := unmFns.UnmarshalEndpointModel(req.PluginContext)
	if err != nil {
		return log.ErrorWithId(logger, log.InternalError, "could not unmarshal endpointModel model", "error", err)
	}

	switch req.Path {
	case "variablequery":
		var variableQuery qf.VariableQuery

		// Extract the query parameter from the POST request
		jsonStr := string(req.Body)
		data := VariableQueryRequest{}
		err := json.Unmarshal([]byte(jsonStr), &data)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not unmarshal data", "error", err)
		}
		svcQuery := data.Query
		if len(strings.TrimSpace(svcQuery)) == 0 {
			return log.ErrorWithId(logger, log.InputError, "variable query cannot be blank")
		}

		// Execute the query and return as body
		xmlData, err := variableQuery.ExecuteQuery(svcQuery, endpointModel)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch xml data", "query", svcQuery, "error", err)
		} else {
			logger.Debug("executed variable query and got response", "query", svcQuery, "response", xmlData)
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   xmlData,
		})
	case "autopopulate":
		var confQuery qf.ConfigurationQuery
		metricsData, err := confQuery.FetchMetricsFromIndex(endpointModel)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not fetch (autopopulate) metrics", "error", err)
		} else {
			logger.Debug("executed autopopulate and got response", "response", metricsData)
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
func (d *RMFDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {

	// create queryDataResponse struct
	queryDataResponse := backend.NewQueryDataResponse()

	if len(req.Queries) == 0 {
		return queryDataResponse, nil
	}

	// loop over queries and execute them individually.
	for _, q := range req.Queries {
		dataResponse := d.query(ctx, req.PluginContext, q)
		if dataResponse != nil {
			queryDataResponse.Responses[q.RefID] = *dataResponse
		}
	}

	return queryDataResponse, nil
}

func (d *RMFDatasource) query(_ context.Context, pCtx backend.PluginContext, q backend.DataQuery) *backend.DataResponse {
	var (
		pnlfuncs          pnlf.PanelFunctions
		unmFns            umf.UnmarshalFunctions
		confQuery         qf.ConfigurationQuery
		intervalAndOffset typ.IntervalOffset
	)

	logger := log.Logger.With("func", "query")
	var dataResponse *backend.DataResponse
	dataResponseChnl := make(chan *backend.DataResponse)

	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	//Unmarshal the input models
	queryModel, endpointModel, err := unmFns.UnmarshalQueryAndEndpointModel(q, pCtx)
	if err != nil || queryModel.SelectedQuery == "" {
		return nil
	}

	// Change the PanelId by appending the user name. Unique path is used as part of channelid for streaming
	queryModel.RMFPanelId, queryModel.UniquePath = pnlfuncs.GetRMFPanelIDs(queryModel, pCtx)

	// Set the server time data and save in cache for subsequent retrieval
	if intervalAndOffset, err = pnlfuncs.GetIntervalAndOffsetFromCache(d.Cache, queryModel.Datasource.Uid); err != nil {
		intervalAndOffset, err = confQuery.FetchIntervalAndOffset(endpointModel)
		if err != nil {
			err = log.ErrorWithId(logger, log.ConnectionError, "could not fetch Interval and Offset", "error", err)
		} else {
			err = pnlfuncs.SaveIntervalAndOffsetInCache(d.Cache, queryModel.Datasource.Uid, intervalAndOffset)
			if err != nil {
				err = log.ErrorWithId(logger, log.InternalError, "could not save Interval and Offset", "error", err)
			}
		}
	}

	if err != nil {
		go func() {
			dataResponse = &backend.DataResponse{}
			dataResponse.Error = err
			dataResponseChnl <- dataResponse
		}()
	} else {
		// Save in queryModel
		queryModel.ServerTimeData.ServiceCallInterval = intervalAndOffset.ServiceCallInterval
		queryModel.ServerTimeData.ServerTimezoneOffset = intervalAndOffset.ServerTimezoneOffset

		if queryModel.SelectedVisualisationType == "TimeSeries" {
			go d.query_timeseries(pCtx, queryModel, endpointModel, dataResponseChnl)
		} else {
			go d.query_tabledata(queryModel, endpointModel, dataResponseChnl)
		}
	}
	dataResponse = <-dataResponseChnl
	return dataResponse
}

func (d *RMFDatasource) query_timeseries(pCtx backend.PluginContext, queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, dataResponseChnl chan *backend.DataResponse) {
	logger := log.Logger.With("func", "query_timeseries")
	var (
		newFrame     *data.Frame
		err          error
		dataResponse *backend.DataResponse = &backend.DataResponse{}
	)

	if newFrame, err = d.query_timeseries_internal(pCtx, queryModel, endpointModel); err != nil {
		dataResponse.Error = log.FrameErrorWithId(logger, err)
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	dataResponseChnl <- dataResponse
}

func (d *RMFDatasource) query_timeseries_internal(pCtx backend.PluginContext, queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {

	logger := log.Logger.With("func", "query_timeseries_internal")

	defer log.LogAndRecover(logger)

	// Insert QueryModel in Panels if not present. A Panel is an abstraction that stores its QueryModels underneath.
	// A Panel is created to ensure that a single service call is active at any given point of time within a panel
	// because RMF is not currently capable of handling concurrent service calls from a single panel to fetch metrics.
	var (
		// frameFunctions ff.FrameFunctions
		pnlfuncs pnlf.PanelFunctions
		newFrame *data.Frame
		err      error
	)

	// Do any final updates before returning.
	defer func() {
		d.EndpointModel = endpointModel
		pnlfuncs.SaveQueryModelInCache(d.Cache, queryModel)
	}()

	newFrame, err = d.fetchTimeseriesData(queryModel, endpointModel)
	if err != nil {
		return nil, log.FrameErrorWithId(logger, err)
	}
	d.createChannelForStreaming(pCtx, queryModel, newFrame)

	return newFrame, nil
}

func (d *RMFDatasource) createChannelForStreaming(pCtx backend.PluginContext, queryModel *typ.QueryModel, newFrame *data.Frame) {
	var pnlfn pnlf.PanelFunctions
	guid := pnlfn.GetNewGuid()

	// Create the channel
	// If query called with streaming on then return a channel
	// to subscribe on a client-side and consume updates from a plugin.
	channel := live.Channel{
		Scope:     live.ScopeDatasource,
		Namespace: pCtx.DataSourceInstanceSettings.UID, //"stream", // "dashboard", // pCtx.DataSourceInstanceSettings.UID,
		Path:      queryModel.UniquePath + "/" + guid,
	}
	newFrame.SetMeta(&data.FrameMeta{
		Channel: channel.String(),
	})
}

func (d *RMFDatasource) fetchTimeseriesData(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	logger := log.Logger.With("func", "fetchTimeseriesData")
	var (
		pnlfuncs pnlf.PanelFunctions
		newFrame *data.Frame
		err      error
	)

	defer log.LogAndRecover(logger)

	// Has either the datasource, query or interval or (from and to date) range changed? Then a panel refresh is required
	panelRefreshRequired := pnlfuncs.PanelRefreshRequired(d.Cache, queryModel)
	if panelRefreshRequired && queryModel.AbsoluteTimeSelected {
		pnlfuncs.DeleteQueryModel(d.Cache, queryModel.UniquePath) // Delete (any) existing queryModel
	} else { // Relative time
		pnlfuncs.DeleteQueryModel(d.Cache, queryModel.UniquePath)      // Delete (any) existing queryModel
		pnlfuncs.DeleteQueryModel(d.Cache, queryModel.UniquePath+"/H") // Delete (any) historical queryModel (also)
	}

	// If frame exists in cache for the query get it from there else get frame from the server through a service call
	if newFrame, err = d.getFrameFromCacheOrServer(queryModel, endpointModel); err != nil {
		return nil, log.FrameErrorWithId(logger, err)
	}

	return newFrame, nil
}

// RunStream is called once for any open channel.  Results are shared with everyone
// subscribed to the same channel.
func (d *RMFDatasource) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	logger := log.Logger.With("func", "RunStream")
	var (
		pnlfuncs pnlf.PanelFunctions
		err      error
	)

	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// The matched query model can be an actual query model or a temp query model
	reqPath := req.Path[:strings.LastIndex(req.Path, "/")]
	matchedQueryModel := pnlfuncs.GetMatchedQueryModelFromCache(d.Cache, reqPath)

	// Stream absolute or relative timeline data
	if matchedQueryModel.AbsoluteTimeSelected {
		err = d.streamDataForAbsoluteRange(ctx, req, sender, matchedQueryModel)
	} else {
		err = d.streamDataForRelativeRange(ctx, req, sender, matchedQueryModel)
	}
	if err != nil {
		return err
	}
	return nil
}

func (d *RMFDatasource) streamDataForAbsoluteRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
	var waitTime time.Duration
	logger := log.Logger.With("func", "streamDataForAbsoluteRange")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Set wait time to 1/100th of a second
	waitTime = (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	err := d.streamDataAbsolute(ctx, req, sender, matchedQueryModel, waitTime)
	if err != nil {
		return err
	}

	return nil
}

func (d *RMFDatasource) streamDataForRelativeRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
	logger := log.Logger.With("func", "streamDataForRelativeRange")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Set wait time to 'ServiceCallInterval' for relative and 1/100th of a second for historical
	waitTime := (time.Second * time.Duration(matchedQueryModel.ServerTimeData.ServiceCallInterval))
	histWaitTime := (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	_, err := d.streamDataRelative(ctx, req, sender, matchedQueryModel, &waitTime, &histWaitTime)
	if err != nil {
		return err
	}
	return nil
}

func (d *RMFDatasource) streamDataAbsolute(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime time.Duration) error {
	logger := log.Logger.With("func", "streamDataAbsolute")
	var (
		pnlfuncs pnlf.PanelFunctions
		newFrame *data.Frame
		err      error
	)
	histTicker := time.NewTicker(waitTime)
	seriesFields := framef.SeriesFields{}

	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			logger.Debug("closing stream", "reason", err, "path", req.Path)
			pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if matchedQueryModel.TimeSeriesTimeRangeTo.After(matchedQueryModel.TimeRangeTo) {
				histTicker.Stop()
				return nil
			}
			// Send new data periodically.
			logger.Debug("executing query", "query", matchedQueryModel.SelectedQuery)
			if newFrame, err = d.getFrameFromCacheOrServer(matchedQueryModel, d.EndpointModel); err != nil {
				return log.ErrorWithId(logger, log.InternalError, "could not get new frame", "error", err)
			}
			framef.SyncFieldNames(seriesFields, newFrame, matchedQueryModel.TimeSeriesTimeRangeFrom)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return log.ErrorWithId(logger, log.InternalError, "failed to send frame", "error", err)
			}
			// Save the query model in cache
			pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
		}
	}
}

func (d *RMFDatasource) streamDataRelative(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime *time.Duration, histWaitTime *time.Duration) (bool, error) {
	logger := log.Logger.With("func", "streamDataRelative")
	var (
		pnlfuncs       pnlf.PanelFunctions
		newFrame       *data.Frame
		err            error
		histQueryModel typ.QueryModel
	)
	mainTicker := time.NewTicker(*waitTime)
	histTicker := time.NewTicker(*histWaitTime)
	seriesFields := framef.SeriesFields{}
	duration := matchedQueryModel.TimeRangeTo.Sub(matchedQueryModel.TimeRangeFrom)

	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			logger.Debug("closing stream", "reason", err, "path", req.Path)
			// Stop tickers to enable garbage collection of resources
			mainTicker.Stop()
			histTicker.Stop()
			return true, err
		case <-histTicker.C: // Historical Data: Wait for a smaller interval and query DDS
			var numberOfIterations int
			if numberOfIterations, err = pnlfuncs.GetIterationsForReverseAbsPlotting(matchedQueryModel); err != nil {
				return false, err
			}
			histQueryModel = d.getHistoricalQueryModel(matchedQueryModel)
			if histQueryModel.TimeSeriesTimeRangeFrom.Before(matchedQueryModel.TimeRangeFrom) { //TimeRangeFrom
				histTicker.Stop()
			} else {
				logger.Debug("executing query for historical data", "query", histQueryModel.SelectedQuery)
				for counter := 0; counter < numberOfIterations; counter++ {
					// Fetch the data
					if newFrame, err = d.getFrameFromCacheOrServer(&histQueryModel, d.EndpointModel, true); err != nil {
						return false, log.ErrorWithId(logger, log.InternalError, "could not get new frame for historical data", "error", err)
					}
					framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
					err = sender.SendFrame(newFrame, data.IncludeAll)
					if err != nil {
						return false, log.ErrorWithId(logger, log.InternalError, "failed to send frame for historical data", "error", err)
					}
					pnlfuncs.SaveQueryModelInCache(d.Cache, &histQueryModel)
				}
			}
		case <-mainTicker.C: // Relative Data: Wait for a (usually) larger interval (i.e. gathererInterval) and query DDS
			var numberOfIterations int
			if numberOfIterations, err = pnlfuncs.GetIterationsForRelativePlotting(matchedQueryModel); err != nil {
				return false, err
			}
			logger.Debug("executing query for relative data", "query", matchedQueryModel.SelectedQuery, "iterations", numberOfIterations)
			for counter := 0; counter < numberOfIterations; counter++ {
				if newFrame, err = d.getFrameFromCacheOrServer(matchedQueryModel, d.EndpointModel); err != nil {
					return false, log.ErrorWithId(logger, log.InternalError, "could not get new frame for relative data", "error", err)
				}
				framef.RemoveOldFieldNames(seriesFields, matchedQueryModel.TimeSeriesTimeRangeFrom.Add(-duration))
				framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
				err = sender.SendFrame(newFrame, data.IncludeAll)
				if err != nil {
					return false, log.ErrorWithId(logger, log.InternalError, "failed to send frame for relative data", "error", err)
				}
				// Save the query model in cache
				pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
			}
		}
	}
}

func (d *RMFDatasource) getFrameFromCacheOrServer(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
	logger := log.Logger.With("func", "getFrameFromCacheOrServer")
	var (
		newFrame        *data.Frame
		timeSeriesQuery qf.TimeSeriesQuery
		err             error
		pnlfuncs        pnlf.PanelFunctions
	)

	frameNotFoundInCache := false

	// For relative timeseries (forward only) don't look in the cache
	if !queryModel.AbsoluteTimeSelected {
		timeSeriesQuery.SetTimeRange(queryModel)
	} else {
		// If frame exists in cache for the query get it from there
		// else get frame from the server through a service call
		if len(plotAbsoluteReverse) > 0 {
			timeSeriesQuery.SetTimeRange(queryModel, plotAbsoluteReverse[0])
			newFrame, err = pnlfuncs.GetFrameFromCache(d.Cache, queryModel, plotAbsoluteReverse[0])
		} else {
			timeSeriesQuery.SetTimeRange(queryModel)
			newFrame, err = pnlfuncs.GetFrameFromCache(d.Cache, queryModel)
		}
		if err != nil {
			frameNotFoundInCache = true
		}
	}

	// Fetch from the DDS Server and then save to cache if required.
	if (!queryModel.AbsoluteTimeSelected) || frameNotFoundInCache {
		newFrame, err = timeSeriesQuery.QueryForTimeseriesDataFrame(queryModel, endpointModel)
		if err != nil {
			return nil, log.FrameErrorWithId(logger, err)
		} else {
			if err = pnlfuncs.SaveFrameInCache(d.Cache, queryModel, newFrame); err != nil {
				return nil, log.ErrorWithId(logger, log.InternalError, "could not save frame", "error", err)
			}
		}
	}
	return newFrame, nil
}

func (d *RMFDatasource) getHistoricalQueryModel(matchedQueryModel *typ.QueryModel) typ.QueryModel {
	var (
		resultQueryModel typ.QueryModel
		pnlfuncs         pnlf.PanelFunctions
	)
	matchedHQueryModel := pnlfuncs.GetMatchedQueryModelFromCache(d.Cache, matchedQueryModel.UniquePath+"/H")

	if matchedHQueryModel == nil {
		resultQueryModel.AbsoluteTimeSelected = true
		resultQueryModel.Datasource = matchedQueryModel.Datasource
		resultQueryModel.SelectedVisualisationType = matchedQueryModel.SelectedVisualisationType
		resultQueryModel.TimeRangeFrom = matchedQueryModel.TimeRangeFrom
		resultQueryModel.TimeRangeTo = matchedQueryModel.TimeRangeTo
		resultQueryModel.TimeSeriesTimeRangeFrom = matchedQueryModel.TimeSeriesTimeRangeFrom
		resultQueryModel.TimeSeriesTimeRangeTo = matchedQueryModel.TimeSeriesTimeRangeTo
		resultQueryModel.SelectedQuery = matchedQueryModel.SelectedQuery
		resultQueryModel.SelectedResource = matchedQueryModel.SelectedResource
		resultQueryModel.UniquePath = matchedQueryModel.UniquePath + "/H" //means 'Historical'
		resultQueryModel.ServerTimeData = matchedQueryModel.ServerTimeData
	} else {
		resultQueryModel = *matchedHQueryModel
	}
	return resultQueryModel
}

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (d *RMFDatasource) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	logger := log.Logger.With("func", "SubscribeStream")
	var pnlfuncs pnlf.PanelFunctions
	status := backend.SubscribeStreamStatusPermissionDenied

	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	reqPath := req.Path[:strings.LastIndex(req.Path, "/")]
	matchedQueryModel := pnlfuncs.GetMatchedQueryModelFromCache(d.Cache, reqPath)

	if matchedQueryModel != nil { // we got a matched QueryModel? Then we are ok.
		// Allow subscribing only on expected path.
		status = backend.SubscribeStreamStatusOK
	}
	return &backend.SubscribeStreamResponse{
		Status: status,
	}, nil
}

// PublishStream is called when a client sends a message to the stream.
func (d *RMFDatasource) PublishStream(_ context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	logger := log.Logger.With("func", "PublishStream")
	//Recover from any panic so as to not bring down this backend datasource
	defer log.LogAndRecover(logger)

	// Do not allow publishing at all.
	return &backend.PublishStreamResponse{
		Status: backend.PublishStreamStatusPermissionDenied,
	}, nil
}

func (d *RMFDatasource) query_tabledata(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, dataResponseChnl chan *backend.DataResponse) {
	logger := log.Logger.With("func", "query_tabledata")
	// Insert QueryModel in Panels if not present. A Panel is an abstraction that stores its QueryModels underneath.
	// A Panel is created to ensure that a single service call is active at any given point of time within a panel
	// because RMF is not currently capable of handling concurrent service calls from a single panel to fetch metrics.
	var (
		pnlfuncs     pnlf.PanelFunctions
		dataResponse *backend.DataResponse = &backend.DataResponse{}
	)

	pnlfuncs.SaveQueryModelInCache(d.Cache, queryModel)

	// Do any final updates before returning.
	defer func() {
		d.EndpointModel = endpointModel
		pnlfuncs.SaveQueryModelInCache(d.Cache, queryModel)
		if err := pnlfuncs.UpdateServiceCallInProgressStatus(d.Cache, queryModel, false); err != nil {
			logger.Info("could not update service call progress status (false)", "error", err)
		}
	}()
	if err := pnlfuncs.UpdateServiceCallInProgressStatus(d.Cache, queryModel, true); err != nil {
		logger.Info("could not update service call progress status (true)", "error", err)
	}

	var tableDataQuery qf.TableDataQuery

	// Query the table data
	if newFrame, err := tableDataQuery.QueryForTableData(queryModel, endpointModel); err != nil {
		dataResponse.Error = log.FrameErrorWithId(logger, err)
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	dataResponseChnl <- dataResponse
}

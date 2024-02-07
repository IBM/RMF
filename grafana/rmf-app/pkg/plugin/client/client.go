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

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"github.com/grafana/grafana-plugin-sdk-go/live"

	cf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/cache_functions"
	errh "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/error_handler"
	framef "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame_functions"
	pnlf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/panel_functions"
	plugincnfg "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/plugin_config"
	qf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/query_functions"
	typ "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/types"
	umf "github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/unmarshal_functions"
)

// RMFClient is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type RMFClient struct {
	Cache         *cf.RMFCache
	EndpointModel *typ.DatasourceEndpointModel
	ErrHandler    *errh.ErrHandler
	PluginConfig  *plugincnfg.PluginConfig
}

func NewRMFClient() RMFClient {
	pluginCnfg := plugincnfg.NewPluginConfig()
	return RMFClient{
		Cache:        cf.NewRMFCache(pluginCnfg),
		ErrHandler:   errh.NewErrHandler(pluginCnfg),
		PluginConfig: pluginCnfg,
	}
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewRMFClient factory function.
func (d *RMFClient) Dispose() {
	// Clean up datasource instance resources.
	d.Cache.Close()
	d.EndpointModel = nil
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *RMFClient) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	var unmFns umf.UnmarshalFunctions
	var message string
	var status backend.HealthStatus = backend.HealthStatusError

	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

	//Unmarshal the endpoint model
	endpointModel, err := unmFns.UnmarshalEndpointModel(req.PluginContext)
	if err != nil {
		return nil, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not unmarshal endpointModel model in CheckHealth(). Error= %v", err))
	}

	var confQuery qf.ConfigurationQuery
	res, err := confQuery.FetchRootInfo(endpointModel, d.ErrHandler)
	if err != nil {
		if _, ok := err.(*typ.ValueError); ok {
			message = "Unsupported version of DDS"
		} else {
			err = d.ErrHandler.LogErrorAndReturnErrorInfo("E", "ERR_DATASRC_CONNECTION_FAILED", err)
			message = err.Error()
		}
	}
	if res {
		status = backend.HealthStatusOk
		message = "Data source is working"
	}
	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil
}

type VariableQueryRequest struct {
	Query string `json:"query"`
}

func (d *RMFClient) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	var unmFns umf.UnmarshalFunctions

	//Unmarshal the endpoint model
	endpointModel, err := unmFns.UnmarshalEndpointModel(req.PluginContext)
	if err != nil {
		return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not unmarshal endpointModel model in CallResource(). Error= %v", err))
	}

	switch req.Path {
	case "variablequery":
		var variableQuery qf.VariableQuery

		// Extract the query parameter from the POST request
		jsonStr := string(req.Body)
		data := VariableQueryRequest{}
		err := json.Unmarshal([]byte(jsonStr), &data)
		if err != nil {
			return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not unmarshal jsonStr data in CallResource(). Error= %v", err))
		}
		svcQuery := data.Query
		if len(strings.TrimSpace(svcQuery)) == 0 {
			return d.ErrHandler.LogErrorAndReturnErrorInfo("I", "ERR_INVALID_INPUT", fmt.Errorf("variable query cannot be blank"))
		}

		// Execute the query and return as body
		xmlData, err := variableQuery.ExecuteQuery(svcQuery, endpointModel)
		if err != nil {
			return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not fetch xml data in CallResource() for query: %s. error:= %v", svcQuery, err))
		} else {
			d.ErrHandler.LogStatus(fmt.Sprintf("***Executed variable query: %s and got response in CallResource(): %s", svcQuery, xmlData))
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   xmlData,
		})
	case "autopopulate":
		var confQuery qf.ConfigurationQuery
		metricsData, err := confQuery.FetchMetricsFromIndex(endpointModel)
		if err != nil {
			return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not fetch (autopopulate) metrics in CallResource(). error: %v", err))
		} else {
			d.ErrHandler.LogStatus(fmt.Sprintf("***Executed autopopulate (FetchMetricsFromIndex) and got response in CallResource(): %v", metricsData))
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
func (d *RMFClient) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {

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

func (d *RMFClient) query(ctx context.Context, pCtx backend.PluginContext, q backend.DataQuery) *backend.DataResponse {
	var (
		pnlfuncs          pnlf.PanelFunctions
		unmFns            umf.UnmarshalFunctions
		confQuery         qf.ConfigurationQuery
		intervalAndOffset typ.IntervalOffset
	)

	var dataResponse *backend.DataResponse
	dataResponseChnl := make(chan *backend.DataResponse)

	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

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
			err = d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_UNABLE_TO_FETCH_DATA", fmt.Errorf("could not fetch Interval and Offset in query(). Error=%v", err))
		} else {
			err = pnlfuncs.SaveIntervalAndOffsetInCache(d.Cache, queryModel.Datasource.Uid, intervalAndOffset)
			if err != nil {
				err = d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not save Interval and Offset in query(). Error=%v", err))
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

func (d *RMFClient) query_timeseries(pCtx backend.PluginContext, queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, dataResponseChnl chan *backend.DataResponse) {
	var (
		newFrame     *data.Frame
		err          error
		dataResponse *backend.DataResponse = &backend.DataResponse{}
	)

	if newFrame, err = d.query_timeseries_internal(pCtx, queryModel, endpointModel); err != nil {
		dataResponse.Error = d.getErrorInternal(err)
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	dataResponseChnl <- dataResponse
}

func (d *RMFClient) query_timeseries_internal(pCtx backend.PluginContext, queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {

	defer d.ErrHandler.HandleErrors()

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
		return nil, d.getErrorInternal(err)
	}
	d.createChannelForStreaming(pCtx, queryModel, newFrame)

	return newFrame, nil
}

func (d *RMFClient) createChannelForStreaming(pCtx backend.PluginContext, queryModel *typ.QueryModel, newFrame *data.Frame) {
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

func (d *RMFClient) fetchTimeseriesData(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel) (*data.Frame, error) {
	var (
		pnlfuncs pnlf.PanelFunctions
		newFrame *data.Frame
		err      error
	)

	defer d.ErrHandler.HandleErrors()

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
		return nil, d.getErrorInternal(err)
	}

	return newFrame, nil
}

// RunStream is called once for any open channel.  Results are shared with everyone
// subscribed to the same channel.
func (d *RMFClient) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	var (
		pnlfuncs pnlf.PanelFunctions
		err      error
	)

	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

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

func (d *RMFClient) streamDataForAbsoluteRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
	var waitTime time.Duration

	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

	// Set wait time to 1/100th of a second
	waitTime = (time.Second * time.Duration(1)) / 100

	// Stream data frames periodically till stream closed by Grafana.
	err := d.streamDataAbsolute(ctx, req, sender, matchedQueryModel, waitTime)
	if err != nil {
		return err
	}

	return nil
}

func (d *RMFClient) streamDataForRelativeRange(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel) error {
	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

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

func (d *RMFClient) streamDataAbsolute(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime time.Duration) error {
	var (
		pnlfuncs pnlf.PanelFunctions
		newFrame *data.Frame
		err      error
	)
	histTicker := time.NewTicker(waitTime)
	seriesFields := map[string]time.Time{}
	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			d.ErrHandler.LogStatus("**Stream closed: Reason: ", err, ": path", req.Path)
			pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
			histTicker.Stop()
			return err
		case <-histTicker.C:
			if matchedQueryModel.TimeSeriesTimeRangeTo.After(matchedQueryModel.TimeRangeTo) {
				histTicker.Stop()
				return nil
			}
			// Send new data periodically.
			d.ErrHandler.LogStatus("*RunStream executing for ", matchedQueryModel.SelectedQuery)
			if newFrame, err = d.getFrameFromCacheOrServer(matchedQueryModel, d.EndpointModel); err != nil {
				return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not get new frame in streamDataAbsolute(). error=%v", err))
			}
			framef.SyncFieldNames(seriesFields, newFrame, matchedQueryModel.TimeSeriesTimeRangeFrom)
			err = sender.SendFrame(newFrame, data.IncludeAll)
			if err != nil {
				return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("error sending frame in streamDataAbsolute(). error:%v", err))
			}
			// Save the query model in cache
			pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
		}
	}
}

func (d *RMFClient) streamDataRelative(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender, matchedQueryModel *typ.QueryModel, waitTime *time.Duration, histWaitTime *time.Duration) (bool, error) {
	var (
		pnlfuncs       pnlf.PanelFunctions
		newFrame       *data.Frame
		err            error
		histQueryModel typ.QueryModel
	)
	mainTicker := time.NewTicker(*waitTime)
	histTicker := time.NewTicker(*histWaitTime)
	seriesFields := map[string]time.Time{}
	duration := matchedQueryModel.TimeRangeTo.Sub(matchedQueryModel.TimeRangeFrom)

	for {
		select {
		case <-ctx.Done(): // Did the client cancel out?
			err := ctx.Err()
			d.ErrHandler.LogStatus("*Stream closed and stopping timers: ", err, ": path", req.Path)
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
				d.ErrHandler.LogStatus("*RunStream (Historical) executing for ", histQueryModel.SelectedQuery)
				for counter := 0; counter < numberOfIterations; counter++ {
					// Fetch the data
					if newFrame, err = d.getFrameFromCacheOrServer(&histQueryModel, d.EndpointModel, true); err != nil {
						return false, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not get new frame for absolute data (historical) in streamDataRelative(). error=%v", err))
					}
					framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
					err = sender.SendFrame(newFrame, data.IncludeAll)
					if err != nil {
						return false, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("error sending frame for absolute data (historical, SendFrame) in streamDataRelative(). error:%v", err))
					}
					pnlfuncs.SaveQueryModelInCache(d.Cache, &histQueryModel)
				}
			}
		case <-mainTicker.C: // Relative Data: Wait for a (usually) larger interval (i.e. gathererInterval) and query DDS
			var numberOfIterations int
			if numberOfIterations, err = pnlfuncs.GetIterationsForRelativePlotting(matchedQueryModel); err != nil {
				return false, err
			}
			d.ErrHandler.LogStatus(fmt.Sprintf("*RunStream executing for %s (%d time(s))", matchedQueryModel.SelectedQuery, numberOfIterations))
			for counter := 0; counter < numberOfIterations; counter++ {
				if newFrame, err = d.getFrameFromCacheOrServer(matchedQueryModel, d.EndpointModel); err != nil {
					return false, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not get new frame for relative data in streamDataRelative(). error=%v", err))
				}
				framef.RemoveOldFieldNames(seriesFields, matchedQueryModel.TimeSeriesTimeRangeFrom.Add(-duration))
				framef.SyncFieldNames(seriesFields, newFrame, histQueryModel.TimeSeriesTimeRangeFrom)
				err = sender.SendFrame(newFrame, data.IncludeAll)
				if err != nil {
					return false, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("error sending frame for relative data (forward, SendFrame) in streamDataRelative(). error:%v", err))
				}
				// Save the query model in cache
				pnlfuncs.SaveQueryModelInCache(d.Cache, matchedQueryModel)
			}
		}
	}
}

func (d *RMFClient) getFrameFromCacheOrServer(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, plotAbsoluteReverse ...bool) (*data.Frame, error) {
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
		newFrame, err = timeSeriesQuery.QueryForTimeseriesDataFrame(queryModel, endpointModel, d.ErrHandler)
		if err != nil {
			return nil, d.getErrorInternal(err)
		} else {
			if err = pnlfuncs.SaveFrameInCache(d.Cache, queryModel, newFrame); err != nil {
				return nil, d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("could not save frame in getFrameFromCacheOrServer(): Error=%v", err))
			}
		}
	}
	return newFrame, nil
}

func (d *RMFClient) getHistoricalQueryModel(matchedQueryModel *typ.QueryModel) typ.QueryModel {
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
func (d *RMFClient) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	var pnlfuncs pnlf.PanelFunctions
	status := backend.SubscribeStreamStatusPermissionDenied

	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

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
func (d *RMFClient) PublishStream(_ context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	//Recover from any panic so as to not bring down this backend datasource
	defer d.ErrHandler.HandleErrors()

	// Do not allow publishing at all.
	return &backend.PublishStreamResponse{
		Status: backend.PublishStreamStatusPermissionDenied,
	}, nil
}

func (d *RMFClient) query_tabledata(queryModel *typ.QueryModel, endpointModel *typ.DatasourceEndpointModel, dataResponseChnl chan *backend.DataResponse) {
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
			log.DefaultLogger.Info(fmt.Sprintf("could not update service call progress status (false): details: %v ", err), nil)
		}
	}()
	if err := pnlfuncs.UpdateServiceCallInProgressStatus(d.Cache, queryModel, true); err != nil {
		log.DefaultLogger.Info(fmt.Sprintf("could not update service call progress status (true): details: %v ", err), nil)
	}

	var tableDataQuery qf.TableDataQuery

	// Query the table data
	if newFrame, err := tableDataQuery.QueryForTableData(queryModel, endpointModel, d.ErrHandler); err != nil {
		dataResponse.Error = d.getErrorInternal(err)
	} else if newFrame != nil {
		dataResponse.Frames = append(dataResponse.Frames, newFrame)
	}
	dataResponseChnl <- dataResponse
}

func (d *RMFClient) getErrorInternal(err error) error {
	if strings.Contains(err.Error(), "DDSError") {
		return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_DDS_ERROR", err)
	} else {
		return d.ErrHandler.LogErrorAndReturnErrorInfo("S", "ERR_INTERNAL_ERROR", fmt.Errorf("error while retrieving new frame in query_timeseries_internal(): Error=%v", err))
	}
}

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
	"fmt"
	"net/http"
	"regexp"
	"runtime/debug"
	"slices"
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

const ChannelCacheSizeMB = 64
const SdsDelay = 5 * time.Second
const TimeSeriesType = "TimeSeries"
const QueryPattern = `^([A-Za-z_][A-Za-z0-9_]*)\(([^)]*)\)$` // e.g., banner(resource), table(resource), caption(resource)

type RMFDatasource struct {
	uid          string
	name         string
	channelCache *cache.ChannelCache
	frameCache   *cache.FrameCache
	ddsClient    *dds.Client
	single       singleflight.Group
	omegamonDs   string
	queryMatcher *regexp.Regexp
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
	ds.omegamonDs = config.JSON.OmegamonDs
	logger.Info("initialized a datasource",
		"uid", settings.UID, "name", settings.Name,
		"url", config.URL, "timeout", config.Timeout, "cacheSize", config.CacheSize,
		"username", config.Username, "tlsSkipVerify", config.JSON.TlsSkipVerify)
	ds.queryMatcher = regexp.MustCompile(QueryPattern)
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
		varRequest := VariableQueryRequest{}
		err := json.Unmarshal(req.Body, &varRequest)
		if err != nil {
			return log.ErrorWithId(logger, log.InternalError, "could not unmarshal data", "error", err)
		}
		q := varRequest.Query
		q = strings.Trim(q, " ")
		q = strings.ToLower(q)

		switch q {
		case "sysplex":
			data, _ := ds.sysplexContainedJson()
			return sender.Send(&backend.CallResourceResponse{Status: http.StatusOK, Body: data})
		case "systems":
			data, _ := ds.systemsContainedJson()
			return sender.Send(&backend.CallResourceResponse{Status: http.StatusOK, Body: data})
		case "omegamonds":
			data, _ := ds.omegamonContainedJson()
			return sender.Send(&backend.CallResourceResponse{Status: http.StatusOK, Body: data})
		}

		// Extract the query parameter from the POST request
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

func (ds *RMFDatasource) sysplexContainedJson() ([]byte, error) {
	sysplex := ds.ddsClient.GetSysplex()
	return toContainedJson([]string{sysplex})
}

func (ds *RMFDatasource) systemsContainedJson() ([]byte, error) {
	systems := ds.ddsClient.GetSystems()
	slices.Sort(systems)
	return toContainedJson(systems)
}

func (ds *RMFDatasource) omegamonContainedJson() ([]byte, error) {
	return toContainedJson([]string{ds.omegamonDs})
}

func toContainedJson(resources []string) ([]byte, error) {
	type Resource struct {
		Reslabel string `json:"reslabel"`
	}
	type Contained struct {
		Resource []Resource `json:"resource"`
	}
	type ContainedResource struct {
		Contained Contained `json:"contained"`
	}
	type Ddsml struct {
		ContainedResourcesList []ContainedResource `json:"containedResourcesList"`
	}

	contained := Contained{Resource: []Resource{}}
	for _, res := range resources {
		contained.Resource = append(contained.Resource, Resource{Reslabel: "," + res + ","})
	}

	var containedResource = ContainedResource{
		Contained: contained,
	}

	result := Ddsml{
		ContainedResourcesList: []ContainedResource{containedResource},
	}
	return json.Marshal(result)
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

		go func(q *backend.DataQuery) {
			defer wg.Done()

			var response *backend.DataResponse
			var params RequestParams
			err := json.Unmarshal(q.JSON, &params)

			if err != nil {
				response = &backend.DataResponse{Status: backend.StatusBadRequest, Error: err}
			} else if params.Resource.Value == "" {
				response = &backend.DataResponse{Status: backend.StatusOK}
			} else {
				mintime := ds.ddsClient.GetCachedMintime()
				if params.VisType == TimeSeriesType {
					// Initialize time series stream
					step := getStep(mintime, q.Interval)
					fields := frame.SeriesFields{}
					start := q.TimeRange.From.UTC()
					r := dds.NewRequest(params.Resource.Value, start, start, step)
					f, jump, err := ds.getCachedTSFrames(r, q.TimeRange.To.UTC(), step, fields)
					if f == nil || err != nil {
						f = frame.TaggedFrame(start, "No data yet...")
					}
					channel := live.Channel{
						Scope:     live.ScopeDatasource,
						Namespace: ds.uid,
						Path:      uuid.NewString(),
					}
					cachedChannel := cache.Channel{
						Resource:  params.Resource.Value,
						TimeRange: backend.TimeRange{From: start.Add(jump), To: q.TimeRange.To.UTC()},
						Absolute:  params.AbsoluteTime,
						Step:      step,
						Fields:    fields,
					}
					err = ds.channelCache.Set(channel.Path, &cachedChannel)
					if err != nil {
						response = &backend.DataResponse{Status: backend.StatusInternal, Error: err}
					} else {
						f.SetMeta(&data.FrameMeta{Channel: channel.String()})
						response = &backend.DataResponse{Frames: data.Frames{f}}
					}
				} else {
					// Query non-timeseries data
					queryKind, query := ds.parseQuery(params.Resource.Value)
					r := dds.NewRequest(query, q.TimeRange.From.UTC(), q.TimeRange.To.UTC(), mintime)
					response = &backend.DataResponse{}
					newFrame := ds.getCachedReportFrames(r)
					if newFrame == nil {
						newFrame, err = ds.getFrame(r, false)
					}
					if err != nil {
						var msg *dds.Message
						if errors.As(err, &msg) {
							response.Error = err
							response.Status = backend.StatusBadRequest
						} else {
							response.Error = log.FrameErrorWithId(logger, err)
							response.Status = backend.StatusInternal
						}
					} else if newFrame != nil {
						ds.setCachedReportFrames(newFrame, r)
						switch queryKind {
						case "banner":
							newFrame = ds.getFrameBanner(newFrame)
						case "caption":
							newFrame = ds.getFrameCaption(newFrame)
						case "table":
							newFrame = ds.getFrameTable(newFrame)
						}
						response.Frames = append(response.Frames, newFrame)
					}
				}
			}
			responseChan <- ResponseWithId{refId: q.RefID, response: response}
		}(&query)

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

	// res, from, to, absolute, interval, err := decodeChannelPath(string(req.Path))
	c, err := ds.channelCache.Get(req.Path)
	if err != nil {
		logger.Error("unable to find channel", "err", err)
		return nil
	}
	res := c.Resource
	step := c.Step
	absolute := c.Absolute
	from := c.TimeRange.From
	to := c.TimeRange.To
	fields := c.Fields

	logger.Debug("starting streaming", "step", step.String(), "path", req.Path)
	r := dds.NewRequest(res, from, from, step)

	// Stream historical part of time series
	stop := to
	for {
		if !absolute {
			stop = time.Now().Add(-SdsDelay)
		}
		if r.TimeRange.To.After(stop) {
			logger.Debug("finished with historical data", "request", r.String(), "path", req.Path)
			break
		}
		f, jump, err := ds.getCachedTSFrames(r, stop, step, fields)
		if err != nil {
			logger.Info("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		if f != nil {
			if err := sender.SendFrame(f, data.IncludeAll); err != nil {
				logger.Info("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			r.Add(jump)
			continue
		}
		if err := ds.serveTSFrame(ctx, sender, fields, r, true); err != nil {
			logger.Info("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		r.Add(step)
	}
	if !absolute {
		// Stream live data as it's being collected
		for {
			if err := ds.serveTSFrame(ctx, sender, fields, r, false); err != nil {
				logger.Info("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			r.Add(step)
		}
	} else if len(fields) == 0 {
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

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (ds *RMFDatasource) SubscribeStream(_ context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	return &backend.SubscribeStreamResponse{Status: backend.SubscribeStreamStatusOK}, nil
}

// PublishStream is called when a client sends a message to the stream.
func (d *RMFDatasource) PublishStream(_ context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	return &backend.PublishStreamResponse{Status: backend.PublishStreamStatusPermissionDenied}, nil
}

func (d *RMFDatasource) parseQuery(resource string) (string, string) {
	matches := d.queryMatcher.FindStringSubmatch(resource)
	if len(matches) == 3 {
		return strings.ToLower(matches[1]), matches[2]
	}
	return "", resource
}

func copyReportField(field *data.Field, length int) *data.Field {
	var newField *data.Field
	t := field.Type()
	switch t {
	case data.FieldTypeNullableFloat64:
		newField = data.NewField(field.Name, field.Labels, []*float64{})
	case data.FieldTypeNullableString:
		newField = data.NewField(field.Name, field.Labels, []*string{})
	default:
		newField = data.NewField(field.Name, field.Labels, []string{})
	}
	newField.SetConfig(field.Config)
	length = slices.Min([]int{length, field.Len()})
	for i := 0; i < length; i++ {
		newField.Append(field.At(i))
	}
	return newField
}

func (ds *RMFDatasource) getFrameTable(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	for _, field := range f.Fields {
		if !strings.HasPrefix(field.Name, frame.CaptionPrefix) &&
			!strings.HasPrefix(field.Name, frame.BannerPrefix) {
			var newField = copyReportField(field, field.Len())
			newField.Delete(0)
			newFrame.Fields = append(newFrame.Fields, newField)
		}
	}
	return &newFrame
}

func (ds *RMFDatasource) getFrameCaption(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	newFrame.Fields = append(newFrame.Fields, data.NewField("Key", nil, []string{}))
	newFrame.Fields = append(newFrame.Fields, data.NewField("Value", nil, []string{}))
	for _, field := range f.Fields {
		if strings.HasPrefix(field.Name, frame.CaptionPrefix) {
			key := strings.TrimPrefix(field.Name, frame.CaptionPrefix)
			value := getStringAt(field, 0)
			newFrame.Fields[0].Append(key)
			newFrame.Fields[1].Append(value)
		}
	}
	return &newFrame
}

func (ds *RMFDatasource) getFrameBanner(f *data.Frame) *data.Frame {
	var newFrame data.Frame
	newFrame.Fields = append(newFrame.Fields, data.NewField("Key", nil, []string{}))
	newFrame.Fields = append(newFrame.Fields, data.NewField("Value", nil, []string{}))
	for _, field := range f.Fields {
		if strings.HasPrefix(field.Name, frame.BannerPrefix) {
			key := strings.TrimPrefix(field.Name, frame.BannerPrefix)
			value := getStringAt(field, 0)
			newFrame.Fields[0].Append(key)
			newFrame.Fields[1].Append(value)
		}
	}
	return &newFrame
}

func getStringAt(field *data.Field, index int) string {
	value := ""
	if field.Len() > index {
		v := field.At(index)
		if v != nil {
			if s, ok := v.(string); ok {
				value = s
			} else if s, ok := v.(*string); ok {
				if s != nil {
					value = *s
				}
			} else if f, ok := v.(float64); ok {
				value = fmt.Sprintf("%f", f)
			}
		}
	}
	return value
}

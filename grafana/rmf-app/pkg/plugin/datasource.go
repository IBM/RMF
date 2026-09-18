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
const DDS_BATCH_REQUESTS_FUNCTIONALITY_MASK = 0x1000

type RMFDatasource struct {
	uid                  string
	name                 string
	channelCache         *cache.ChannelCache
	frameCache           *cache.FrameCache
	ddsClient            *dds.Client
	single               singleflight.Group
	omegamonDs           string
	queryMatcher         *regexp.Regexp
	batchRequestInterval int
	ddsSync              time.Duration
}

// NewRMFDatasource creates a new instance of the RMF datasource.
func NewRMFDatasource(ctx context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	logger := log.Logger.With("func", "NewRMFDatasource")
	ds := &RMFDatasource{uid: settings.UID, name: settings.Name}
	config, httpOpts, err := ds.getConfig(ctx, settings)
	if err != nil {
		logger.Error("failed to get config", "error", err)
		return nil, err
	}
	ds.ddsClient, err = dds.NewClient(config.URL, *httpOpts)
	if err != nil {
		logger.Error("failed to create DDS client", "error", err)
		return nil, err
	}
	ds.channelCache = cache.NewChannelCache(ChannelCacheSizeMB)
	ds.frameCache = cache.NewFrameCache(config.CacheSize)
	ds.omegamonDs = config.JSON.OmegamonDs
	ds.batchRequestInterval = config.BatchRequestMinutes
	ds.ddsSync = time.Duration(config.SyncMinute) * time.Minute
	logger.Debug("initialized a datasource",
		"uid", settings.UID, "name", settings.Name,
		"url", config.URL, "timeout", config.Timeout, "cacheSize", config.CacheSize,
		"username", config.Username, "tlsSkipVerify", config.JSON.TlsSkipVerify,
		"batchRequestInterval", ds.batchRequestInterval,
		"syncMinute", config.SyncMinute,
	)
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
	logger.Debug("disposed datasource", "uid", ds.uid, "name", ds.name)
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

func (ds *RMFDatasource) Align(step time.Duration, t time.Time) time.Time {
	periodMillis := step.Milliseconds()
	anchorMillis := ds.ddsSync.Milliseconds()
	tMillis := t.UnixMilli()
	steps := floorDiv(tMillis-anchorMillis, periodMillis)
	result := time.UnixMilli(anchorMillis + steps*periodMillis)
	log.Logger.Debug("Aligning time", "step", step, "t", t, "result", result)
	return result
}

func (ds *RMFDatasource) AlignRange(step time.Duration, s time.Time, e time.Time) (time.Time, time.Time) {
	sa := ds.Align(step, s)
	ea := ds.Align(step, e)
	if ea.Equal(sa) || ea.Before(sa) {
		ea = sa.Add(step)
	}
	log.Logger.Debug("Aligning range", "step", step, "s", s, "e", e, "sa", sa, "ea", ea)
	return sa, ea
}

func (ds *RMFDatasource) AlignBatch(start time.Time, end time.Time) (time.Time, time.Time) {
	step := time.Duration(ds.batchRequestInterval) * time.Minute
	return ds.AlignBatchStep(start, end, step)
}

func (ds *RMFDatasource) AlignBatchStep(start time.Time, end time.Time, step time.Duration) (time.Time, time.Time) {
	truncate := min(step, time.Hour)
	s := start.Truncate(truncate)
	e := s.Add(step)
	if e.After(end) {
		e = end
	}
	log.Logger.Debug("Aligning batch", "step", step, "start", start, "end", end, "s", s, "e", e)
	return s, e
}

// floorDiv is integer division that rounds toward negative infinity,
// matching Java's Math.floorDiv (Go's / truncates toward zero).
func floorDiv(a, b int64) int64 {
	q := a / b
	if (a%b != 0) && ((a < 0) != (b < 0)) {
		q--
	}
	return q
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
					start = ds.Align(step, start)
					r := dds.NewRequest(params.Resource.Value, start, start)
					f, jump, err := ds.getCachedTSFrames(r, q.TimeRange.To.UTC(), step, fields)
					if ds.supportsBatchRequests() {
						span := step
						step = time.Duration(ds.batchRequestInterval) * time.Minute
						s, e := ds.AlignBatch(q.TimeRange.From.UTC(), q.TimeRange.To.UTC())
						r = dds.NewBatchRequest(params.Resource.Value, s, e, span)
						f2, _, err2 := ds.getCachedTSFrames(r, q.TimeRange.To.UTC(), step, fields)
						if f == nil || err != nil {
							f = f2
						} else if f2 != nil && err2 == nil {
							step = frame.GetDuration(f2)
							f, err = frame.MergeInto(f, f2)
						}
					}
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
						Interval:  q.Interval,
						Span:      r.Span,
						Mintime:   mintime,
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
					from, to := ds.AlignRange(mintime, q.TimeRange.From.UTC(), q.TimeRange.To.UTC())
					r := dds.NewRequest(query, from, to)
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
							newFrame = frame.GetFrameBanner(newFrame)
						case "caption":
							newFrame = frame.GetFrameCaption(newFrame)
						case "table":
							newFrame = frame.GetFrameTable(newFrame)
						default:
							if strings.Contains(query, "report=") {
								newFrame = frame.GetFrameTable(newFrame)
							}
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
	span := c.Span
	mintime := c.Mintime
	interval := c.Interval
	absolute := c.Absolute
	from := c.TimeRange.From
	to := c.TimeRange.To
	fields := c.Fields

	logger.Debug("starting streaming", "step", step.String(), "path", req.Path)
	var r *dds.Request
	if ds.supportsBatchRequests() {
		s, e := ds.AlignBatch(from, to)
		r = dds.NewBatchRequest(res, s, e, span)
	} else {
		from, to = ds.AlignRange(mintime, from, from)
		r = dds.NewRequest(res, from, to)
	}

	// Stream historical part of time series
	stop := to
	for {
		if !absolute {
			stop = time.Now().Add(-SdsDelay)
		}
		if r.Batched {
			if r.TimeRange.From.After(stop) {
				logger.Debug("batch finished with historical data", "request", r.String(), "path", req.Path)
				break
			}
		} else {
			if r.TimeRange.To.After(stop) {
				logger.Debug("finished with historical data", "request", r.String(), "path", req.Path)
				break
			}
		}
		f, _, err := ds.getCachedTSFrames(r, stop, step, fields)
		if err != nil {
			logger.Debug("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		if f != nil {
			if err := sender.SendFrame(f, data.IncludeAll); err != nil {
				logger.Debug("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			r.Add(frame.GetDuration(f))
			continue
		}
		if err := ds.serveTSFrame(ctx, sender, fields, r, true); err != nil {
			if gpme, ok := errors.AsType[*dds.GpmError](err); ok && gpme.Id == dds.MESSAGE_ID_NOT_ENOUGTH_MEMORY {
				logger.Debug("GPM0555I: reduce step", "step", step.Minutes())
				step = step / 2
				if step < MinBatchRequestMinutes*time.Minute {
					logger.Info("streaming stopped", "reason", "step is too small", "path", req.Path, "step", step.Minutes(), "error", err)
					return nil
				}
				s, e := ds.AlignBatchStep(r.TimeRange.From, r.TimeRange.To, step)
				r = dds.NewBatchRequest(res, s, e, span)
				continue
			}
			logger.Debug("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
		r.Add(step)
	}
	if !absolute {
		// Stream live data as it's being collected
		if r.Batched {
			step = getStep(interval, mintime)
			start, end := ds.AlignRange(mintime, stop, stop)
			r = dds.NewRequest(res, start, end)
		}
		for {
			if err := ds.serveTSFrame(ctx, sender, fields, r, false); err != nil {
				logger.Debug("streaming stopped", "reason", err, "path", req.Path)
				return nil
			}
			r.Add(step)
		}
	} else if len(fields) == 0 {
		// There is no data at all, send a dummy frame without fields to reflect it in UI
		f := data.NewFrame("")
		if err := sender.SendFrame(f, data.IncludeAll); err != nil {
			logger.Debug("streaming stopped", "reason", err, "path", req.Path)
			return nil
		}
	}
	logger.Debug("streaming stopped", "reason", "all the data sent", "path", req.Path)
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

func (ds *RMFDatasource) supportsBatchRequests() bool {
	fl := ds.ddsClient.GetFunctionality()
	return fl&DDS_BATCH_REQUESTS_FUNCTIONALITY_MASK == DDS_BATCH_REQUESTS_FUNCTIONALITY_MASK
}

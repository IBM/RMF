/**
* (C) Copyright IBM Corp. 2023.
* (C) Copyright Rocket Software, Inc. 2023.
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

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/client"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/models"
)

// Make sure RMFBackendDatasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. In this example datasource instance implements backend.QueryDataHandler,
// backend.CheckHealthHandler, backend.StreamHandler interfaces. Plugin should not
// implement all these interfaces - only those which are required for a particular task.
// For example if plugin does not need streaming functionality then you are free to remove
// methods that implement backend.StreamHandler. Implementing instancemgmt.InstanceDisposer
// is useful to clean up resources used by previous datasource instance when a new datasource
// instance created upon datasource settings changed.
var (
	_ backend.QueryDataHandler      = (*RMFBackendDatasource)(nil)
	_ backend.CheckHealthHandler    = (*RMFBackendDatasource)(nil)
	_ backend.StreamHandler         = (*RMFBackendDatasource)(nil)
	_ instancemgmt.InstanceDisposer = (*RMFBackendDatasource)(nil)
	_ backend.CallResourceHandler   = (*RMFBackendDatasource)(nil)
)

// RMFBackendDatasource is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type RMFBackendDatasource struct {
	Ctx      context.Context
	Settings models.Settings
	client   client.RMFClient
}

// NewRMFBackendDatasource creates a new datasource instance.
func NewRMFBackendDatasource(ctx context.Context, settings models.Settings) (instancemgmt.Instance, error) {
	return &RMFBackendDatasource{
		Ctx:      ctx,
		Settings: settings,
		client:   client.NewRMFClient(),
	}, nil
}

// NewRMFDataSourceInstance creates a new instance
func NewRMFDataSourceInstance(_ context.Context, settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	datasourceSettings, err := models.LoadSettings(settings)
	if err != nil {
		return nil, err
	}

	datasourceSettings.CachingEnabled = true

	gh, err := NewRMFBackendDatasource(context.Background(), datasourceSettings)
	if err != nil {
		return nil, err
	}
	return gh, nil
}

// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewRMFBackendDatasource factory function.
func (d *RMFBackendDatasource) Dispose() {
	// Clean up datasource instance resources.
	d.client.Dispose()
}

func (d *RMFBackendDatasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	return d.client.CallResource(ctx, req, sender)
}

// CheckHealth handles health checks sent from Grafana to the plugin.
// The main use case for these health checks is the test button on the
// datasource configuration page which allows users to verify that
// a datasource is working as expected.
func (d *RMFBackendDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	return d.client.CheckHealth(ctx, req)
}

// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *RMFBackendDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	return d.client.QueryData(ctx, req)
}

// RunStream is called once for any open channel.  Results are shared with everyone
// subscribed to the same channel.
func (d *RMFBackendDatasource) RunStream(ctx context.Context, req *backend.RunStreamRequest, sender *backend.StreamSender) error {
	return d.client.RunStream(ctx, req, sender)
}

// SubscribeStream is called when a client wants to connect to a stream. This callback
// allows sending the first message.
func (d *RMFBackendDatasource) SubscribeStream(ctx context.Context, req *backend.SubscribeStreamRequest) (*backend.SubscribeStreamResponse, error) {
	return d.client.SubscribeStream(ctx, req)
}

// PublishStream is called when a client sends a message to the stream.
func (d *RMFBackendDatasource) PublishStream(ctx context.Context, req *backend.PublishStreamRequest) (*backend.PublishStreamResponse, error) {
	return d.client.PublishStream(ctx, req)
}

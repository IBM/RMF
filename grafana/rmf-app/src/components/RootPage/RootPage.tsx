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
import React, { PureComponent } from 'react';
import { AppRootProps, NavModelItem } from '@grafana/data';
import { config, getBackendSrv } from '@grafana/runtime';
import { Alert } from '@grafana/ui';
import { ApplicationName, ApplicationSubTitle, DataSourceType } from '../../constants';
import { GlobalSettings, RmfDataSourceInstanceSettings } from '../../types';
import { DataSourceList } from '../DataSourceList';

/**
 * Properties
 */
interface Props extends AppRootProps<GlobalSettings> {}

/**
 * State
 */
interface State {
  /**
   * Data sources
   *
   * @type {RmfDataSourceInstanceSettings[]}
   */
  dataSources: RmfDataSourceInstanceSettings[];

  /**
   * Loading
   *
   * @type {boolean}
   */
  loading: boolean;
}

/**
 * Root Page
 */
export class RootPage extends PureComponent<Props, State> {
  /**
   * Default state
   */
  state: State = {
    loading: true,
    dataSources: [],
  };

  /**
   * Mount
   */
  async componentDidMount() {
    this.updateNav();

    /**
     * Get data sources
     */
    const dataSources = await getBackendSrv()
      .get('/api/datasources')
      .then((result: RmfDataSourceInstanceSettings[]) => {
        return result.filter((ds: RmfDataSourceInstanceSettings) => {
          return ds.type === DataSourceType.RMFTYPE;
        });
      });

    /**
     * Workaround, until reload function will be added to DataSourceSrv
     *
     * @see https://github.com/grafana/grafana/issues/30728
     * @see https://github.com/grafana/grafana/issues/29809
     */
    await getBackendSrv()
      .get('/api/frontend/settings')
      .then((settings: any) => {
        if (!settings.datasources) {
          return;
        }

        /**
         * Set data sources
         */
        config.datasources = settings.datasources;
        config.defaultDatasource = settings.defaultDatasource;
      });

    /**
     * Check supported commands for RMF Data Sources
     */
    await Promise.all(
      dataSources.map(async (ds: RmfDataSourceInstanceSettings) => {
        ds.commands = [];

        /**
         * Get Data Source
         */
        // const rmf = await getDataSourceSrv().get(ds.name);

        /**
         * Execute query
         */
        // const dsQuery = rmf.query({
        //   targets: [{ refId: 'A', query: RmfCommand.COMMAND }],
        // } as DataQueryRequest<RmfQuery>) as unknown;

        // const query = lastValueFrom(dsQuery as Observable<DataQueryResponse>);
        // if (!query) {
        //   return;
        // }

        /**
         * Get available commands
         */
        // await query
        //   .then((response: DataQueryResponse) => response.data)
        //   .then((data: DataQueryResponseData[]) =>
        //     data.forEach((item: DataQueryResponseData) => {
        //       item.fields.forEach((field: Field) => {
        //         ds.commands.push(
        //           ...field.values
        //             .toArray()
        //             .filter((value: string) => value.match(/\S+\.\S+|INFO/i))
        //             .map((value) => value.toUpperCase())
        //         );
        //       });
        //     })
        //   )
        //   .catch(() => {});
      })
    );

    /**
     * Set state
     */
    this.setState({
      dataSources,
      loading: false,
    });
  }

  /**
   * Navigation
   */
  updateNav() {
    const { path, onNavChanged, meta } = this.props;
    const tabs: NavModelItem[] = [];

    /**
     * Home
     */
    tabs.push({
      text: 'Home',
      url: path,
      id: 'home',
      icon: 'home',
      active: true,
    });

    /**
     * Header
     */
    const node = {
      text: ApplicationName,
      img: meta.info.logos.large,
      subTitle: ApplicationSubTitle,
      url: path,
      children: tabs,
    };

    /**
     * Update the page header
     */
    onNavChanged({
      node: node,
      main: node,
    });
  }

  /**
   * Render
   */
  render() {
    const { loading, dataSources } = this.state;

    /**
     * Loading
     */
    if (loading) {
      return (
        <Alert title="Loading..." severity="info">
          <p>Loading time depends on the number of configured data sources.</p>
        </Alert>
      );
    }

    return <DataSourceList dataSources={dataSources} />;
  }
}

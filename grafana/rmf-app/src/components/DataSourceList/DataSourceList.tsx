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
import React, { FC, useCallback } from 'react';
import { getBackendSrv, locationService } from '@grafana/runtime';
import { Alert, Button, HorizontalGroup, LinkButton, VerticalGroup } from '@grafana/ui';
import { DataSourceName, DataSourceType } from '../../constants';
import { HighAvailability, RMFCube } from '../../icons';
import { RmfDataSourceInstanceSettings } from '../../types';

/**
 * Properties
 */
interface Props {
  /**
   * Data sources
   *
   * @type {RmfDataSourceInstanceSettings[]}
   */
  dataSources: RmfDataSourceInstanceSettings[];
}

/**
 * Get unique name for a new data source
 * @param dataSources
 */
const getNewDataSourceName = (dataSources: RmfDataSourceInstanceSettings[]) => {
  let postfix = 1;
  const name = DataSourceName.RMFNAME;

  /**
   * Check if exists
   */
  if (!dataSources.some((dataSource) => dataSource.name === name)) {
    return name;
  }

  while (dataSources.some((dataSource) => dataSource.name === `${name}-${postfix}`)) {
    postfix++;
  }

  return `${name}-${postfix}`;
};

export const DataSourceList: FC<Props> = ({ dataSources }) => {
  const addNewDataSource = useCallback(() => {
    getBackendSrv()
      .post('/api/datasources', {
        name: getNewDataSourceName(dataSources),
        type: DataSourceType.RMFTYPE,
        access: 'proxy',
      })
      .then(({ datasource }) => {
        locationService.push(`/datasources/edit/${datasource.uid}`);
      });
  }, [dataSources]);

  /**
   * Check if any data sources was added
   */
  if (dataSources.length === 0) {
    return (
      <div>
        <div className="page-action-bar">
          <div className="page-action-bar__spacer" />
          <Button onClick={addNewDataSource} icon="plus" variant="secondary">
            Add RMF Data Source
          </Button>
        </div>
        <Alert title="Please add RMF Data Sources." severity="info">
          <p>You can add as many RMF data sources as you want.</p>
        </Alert>
      </div>
    );
  }

  /**
   * Return
   */
  return (
    <div>
      <div className="page-action-bar">
        <div className="page-action-bar__spacer" />
        <Button onClick={addNewDataSource} icon="plus" variant="secondary">
          Add RMF Data Source
        </Button>
      </div>

      <section className="card-section card-list-layout-list">
        <ol className="card-list">
          {dataSources.map((rmfds, index) => {
            const title = 'Working as expected';
            const fill = '#DC382D';
            const url = rmfds.url ? rmfds.url : 'Not specified';

            return (
              <li className="card-item-wrapper" key={index} aria-label="check-card">
                <a className="card-item" href={`#`}>
                  <HorizontalGroup justify="space-between">
                    <HorizontalGroup justify="flex-start">
                      <div>
                        <RMFCube size={32} fill={fill} title={title} />
                      </div>
                      <VerticalGroup>
                        <div className="card-item-name">{rmfds.name}</div>
                        <div className="card-item-sub-name">{url}</div>
                      </VerticalGroup>
                    </HorizontalGroup>

                    <HorizontalGroup justify="flex-end">
                      {(rmfds.jsonData.path || rmfds.jsonData.port) && (
                        <VerticalGroup>
                          <div>
                            {rmfds.jsonData.path}:{rmfds.jsonData.port}
                          </div>
                        </VerticalGroup>
                      )}
                      {rmfds.jsonData['client']?.match(/cluster|sentinel/) && (
                        <div>
                          <HighAvailability size={24} fill={fill} />
                        </div>
                      )}
                    </HorizontalGroup>

                    <HorizontalGroup justify="flex-end">
                      <div>
                        <LinkButton
                          variant="secondary"
                          href={`datasources/edit/${rmfds.uid}`}
                          title="Data Source Settings"
                          icon="sliders-v-alt"
                        >
                          Settings
                        </LinkButton>
                      </div>
                    </HorizontalGroup>
                  </HorizontalGroup>
                </a>
              </li>
            );
          })}
        </ol>
      </section>
    </div>
  );
};

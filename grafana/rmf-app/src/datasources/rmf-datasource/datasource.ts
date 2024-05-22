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
import { DataSourceInstanceSettings, MetricFindValue, ScopedVars } from '@grafana/data';
import { DataSourceWithBackend, getBackendSrv, getTemplateSrv } from '@grafana/runtime';
import { getResourceName, queryValidation } from './common/common.helper';
import { ConfigSettings } from './common/configSettings';
import { RMFDataSourceJsonData, RMFQuery, resourceBaseData } from './common/types';

export class DataSource extends DataSourceWithBackend<RMFQuery, RMFDataSourceJsonData> {
  dashBoardRefreshed: boolean;
  dashboardName: string;

  constructor(instanceSettings: DataSourceInstanceSettings<RMFDataSourceJsonData>) {
    super(instanceSettings);
    this.dashBoardRefreshed = true;
    this.dashboardName = '';
  }

  applyTemplateVariables(query: RMFQuery, scopedVars: ScopedVars): Record<string, any> {
    let currentTime = new Date();
    let newquery = getTemplateSrv().replace(query.selectedResource?.label, scopedVars);
    let CurrentDashbaordName = getTemplateSrv().replace('${__dashboard.uid}', scopedVars);

    let timestamp = getTemplateSrv().replace('${__to}', scopedVars);
    let toDate = new Date(parseInt(timestamp));

    var seconds = (currentTime.getTime() - toDate.getTime()) / 1000;
    let isRelative = seconds < 61

    let dashboardChanged: boolean;
    dashboardChanged = false;
    if (CurrentDashbaordName !== '' && this.dashboardName !== CurrentDashbaordName) {
      this.dashboardName = CurrentDashbaordName;
      dashboardChanged = true;
    }

    const interpolatedQuery: RMFQuery = {
      ...query,
      selectedResource: {
        label: query.selectedResource?.label,
        value: newquery,
      },
      refreshRequired: this.dashBoardRefreshed || dashboardChanged,
      absoluteTimeSelected: !isRelative,
      dashboardUid: CurrentDashbaordName,
    };

    this.dashBoardRefreshed = false;

    return interpolatedQuery;
  }

  async metricFindQuery(query: any, options?: any): Promise<MetricFindValue[]> {
    let queryResult = queryValidation(query, resourceBaseData);
    if (queryResult.result === false) {
      throw new Error('Variable: ' + queryResult.errorMessage);
    }

    let resourceString = getTemplateSrv().replace(queryResult.resourceCommand, options.scopedVars);
    let urlPathLoader = ConfigSettings.UrlSettings.END_TAG + encodeURIComponent(resourceString) + ConfigSettings.UrlSettings.END_TAG;
    let id = this.id;

    return new Promise((resolve) => {
      const metricFindValue = this.loadDataFromService(urlPathLoader, id, options)
        .then((resp: any) => {
          const result = JSON.parse(resp.data);
          let resNames = getResourceName(result);
          const retResult: MetricFindValue[] = [];
          resNames.map((row: any, index: number) => {
            let itemName = '';
            try {
              // itemName = row[queryResult.columnName.toLowerCase()][0].trim();
              itemName = row[queryResult.columnName.toLowerCase()].trim();
              let itemNameParts = itemName.split(',');
              if (itemNameParts[1].trim() === '*') {
                itemName = itemNameParts[2].trim(); // Ex: ULQ01,*,ALL_WLM_RESOURCE_GROUPS
              } else {
                itemName = itemNameParts[1].trim(); // Ex: ULQ01,BATCHLOW,WLM_RESOURCE_GROUP
              }
              if (
                queryResult.filterTypes.trim() !== '' &&
                queryResult.filterTypes.indexOf(row.restype.toUpperCase().trim()) !== -1
              ) {
                retResult.push({ text: itemName });
              } else if (queryResult.filterTypes.trim() === '') {
                retResult.push({ text: itemName });
              }
            } catch (errorObj) { }
          });
          return retResult;
        })
        .catch((err) => {
          throw new Error(err.message + ', [resource=' + queryResult.resourceCommand + ']');
        });
      return resolve(metricFindValue);
    });
  }

  async loadDataFromService(urlPath: string, id?: number, options?: any) {
    return new Promise((resolve, reject) =>
      getBackendSrv()
        .fetch({
          method: 'post',
          headers: {
            Accept: 'application/text',
            'Content-Type': 'application/text',
          },
          url: `/api/datasources/${id}/resources/variablequery`,
          responseType: 'text',
          data: JSON.stringify({ query: urlPath }),
        })
        .subscribe(
          (resp) => {
            if (resp.data) {
              if (options !== undefined && options === 'headres') {
              } else {
                resolve({
                  data: resp.data,
                });
              }
            } else {
              reject({ status: 'failure', message: 'Test connection failed.' });
            }
          },
          (err) => {
            reject({ status: 'failure', message: err.data.message });
          }
        )
    );
  }
}

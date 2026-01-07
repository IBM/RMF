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
import { ConfigSettings } from './configSettings';
import { FieldValidationResilt, MetricItem, QueryValidationResult, visualisationType } from './types';

export const getRootXmlPath = (instanceSettings: any): string => {
  if (instanceSettings !== null && instanceSettings !== undefined && instanceSettings.jsonData !== undefined) {
    let rootXmlPath = 'http://' + (instanceSettings!.jsonData.path ? instanceSettings.jsonData.path?.toString() : '');
    rootXmlPath =
      rootXmlPath +
      (instanceSettings!.jsonData?.port !== undefined && instanceSettings!.jsonData?.port.toString().length > 0
        ? ':' + instanceSettings.jsonData.port?.toString()
        : '');
    rootXmlPath = rootXmlPath + ConfigSettings.UrlSettings.ROOT_URL;
    return rootXmlPath;
  } else {
    return '';
  }
};

export const getPath = (rootXmlPath: string): string => {
  if (rootXmlPath !== null && rootXmlPath !== undefined && rootXmlPath !== '') {
    return rootXmlPath.replace(ConfigSettings.UrlSettings.ROOT_URL, '');
  } else {
    return '';
  }
};

export const getResource = (result: any): any[] => {
  // let resource = result !== undefined ? result.ddsml['metric-list'] : '';
  let resource = result !== undefined ? result.metricList : '';
  if (resource === null || resource === undefined) {
    resource = [];
  }
  return resource;
};

export const loadDataToDictionary = (resourceList: any[]): any => {
  let returnResult: any = {};
  if (resourceList !== null && resourceList !== undefined && resourceList.length > 0) {
    resourceList.map((item: any) => {
      let metricList: MetricItem[] = [];
      let resourceTypeKey = item.resource.restype;
      if (item.metric !== null && item.metric !== undefined && item.metric.length > 0) {
        item.metric.map((metricItem: any) => {
          metricList.push({ value: metricItem['id'], name: metricItem.description });
        });
      }
      returnResult[resourceTypeKey] = metricList;
    });
  }
  return returnResult;
};

export const getSelectedResource = (Props: any): string => {
  let result = '';
  if (
    Props !== undefined &&
    Props.query !== undefined &&
    Props.query.selectedResource !== undefined &&
    Props.query.selectedResource.value !== undefined
  ) {
    result = Props.query.selectedResource.value;
  }
  return result;
};

export const isItFirstDotPosition = (str: string, selectionEnd: number): boolean => {
  if (str.toString().indexOf('.') === selectionEnd) {
    return true;
  }
  return false;
};

export const getEnableTimeSeriesStatus = (Props: any): boolean => {
  let result: boolean;
  if (
    Props !== undefined &&
    Props.query !== undefined &&
    Props.query.selectedVisualisationType !== undefined &&
    Props.query.selectedVisualisationType === visualisationType.timeseries
  ) {
    result = true;
  } else {
    result = false;
  }
  return result;
};

export const getVisualisationType = (selectedQuery: string, enableTimeSeries: boolean): visualisationType => {
  if (enableTimeSeries) {
    return visualisationType.timeseries;
  } else {
    return selectedQuery.toString().toUpperCase().indexOf('.REPORT.') !== -1
      ? visualisationType.table
      : visualisationType.bargauge;
  }
};

export const generateUUID = () => {
  let currentTime = new Date().getTime(),
    currentTimePerform = (performance && performance.now && performance.now() * 1000) || 0;
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, (xChar) => {
    let randomNumber = Math.random() * 16;
    if (currentTime > 0) {
      randomNumber = (currentTime + randomNumber) % 16 | 0;
      currentTime = Math.floor(currentTime / 16);
    } else {
      randomNumber = (currentTimePerform + randomNumber) % 16 | 0;
      currentTimePerform = Math.floor(currentTimePerform / 16);
    }
    return (xChar === 'x' ? randomNumber : (randomNumber & 0x7) | 0x8).toString(16);
  });
};

export const queryValidation = (query: string, resourceBaseData: any): QueryValidationResult => {
  let queryResult: QueryValidationResult = {
    result: false,
    resourceCommand: '',
    filterTypes: '',
    columnName: '',
    errorMessage: '',
  };
  //TODO remove when deprecated
  if ("sysplex" == query.toLowerCase() ||
      "systems" == query.toLowerCase() ||
      "omegamonds" == query.toLowerCase()) {
    queryResult.result = true;
    queryResult.resourceCommand = query;
    queryResult.columnName = "RESLABEL";
    return queryResult;
  }
  let result = true;
  if (query === '' || query.length < 20) {
    result = false;
  } else {
    let selectFound = false;
    let columnFound = false;
    let fromFound = false;
    let modelFound = false;
    let whereFound = false;
    let resourceFound = false;
    let filterFound = true;
    let ulqFound = true;
    let resNameFound = true;
    let resTypeFound = true;
    let chkUnwantedFieldStatus: FieldValidationResilt = { result: true, strResult: '' };
    // query = query.toLocaleUpperCase().trim();
    query = query.replace(/\r?\n|\r/g, '  ').trim();
    // query = query.replace(/['"]+/g, '').trim();
    // /^\"[*.]\"/g;
    query = query.replace(/\s\s+/g, ' ').trim();
    query = query.replace(/\s{2,}/g, ' ').trim();
    query = query.replace(/\=\s/g, '=').trim();
    query = query.replace(/\s=/g, '=').trim();
    query = query.replace(/\,\s/g, ',').trim();
    let queryParts = query.split(' ');
    if (queryParts.length > 10 || queryParts.length < 7) {
      result = false;
    } else {
      selectFound = queryParts[0].toLocaleUpperCase().trim() === 'SELECT';
      columnFound = queryParts[1].trim().length > 1;
      queryResult.columnName = columnFound ? 'RES' + queryParts[1].toLocaleUpperCase().trim() : '';
      fromFound = queryParts[2].toLocaleUpperCase().trim() === 'FROM';
      modelFound =
        queryParts[3].toLocaleUpperCase().trim().length > 1 && queryParts[3].toLocaleUpperCase().trim() === 'RESOURCE';
      whereFound = queryParts[4].toLocaleUpperCase().trim() === 'WHERE';

      let ulq = '';
      let resource_name = '*';
      let resource_type = '';
      let filter = '';
      if (selectFound && columnFound && fromFound && modelFound && whereFound) {
        chkUnwantedFieldStatus = chkUnwantedField('FILTER,ULQ,NAME,TYPE, AND', queryParts);
        if (chkUnwantedFieldStatus.result === true) {
          if (fiendLHSField('FILTER', queryParts)) {
            let retRsFilter: FieldValidationResilt = fieldValidation('FILTER', queryParts);
            !retRsFilter.result ? (filterFound = false) : (filter = retRsFilter.strResult);
          }
          if (fiendLHSField('ULQ', queryParts)) {
            let retRsULQ: FieldValidationResilt = fieldValidation('ULQ', queryParts);
            !retRsULQ.result ? (ulqFound = false) : (ulq = retRsULQ.strResult);
          }
          if (fiendLHSField('NAME', queryParts)) {
            let retRsName: FieldValidationResilt = fieldValidation('NAME', queryParts);
            !retRsName.result ? (resNameFound = false) : (resource_name = retRsName.strResult);
            resource_name = retRsName.result && resource_name === '' ? '*' : resource_name;
          }
          if (fiendLHSField('TYPE', queryParts)) {
            let retRsType: FieldValidationResilt = fieldValidation('TYPE', queryParts);
            !retRsType.result ? (resTypeFound = false) : (resource_type = retRsType.strResult);

            resourceFound = true;
            if (resource_name !== '*' || (resource_name === '*' && resource_type.indexOf('ALL_') > -1)) {
              queryResult.resourceCommand = ulq + ',' + resource_name + ',' + resource_type;
            } else {
              let path = getAllResKeyPath(resourceBaseData, resource_type);
              if (path === '') {
                resourceFound = false;
              } else {
                queryResult.resourceCommand = path.replace('<0>', ulq.trim());
              }
            }
            queryResult.filterTypes = filter;
          }
        } else {
          queryResult.errorMessage = chkUnwantedFieldStatus.strResult;
        }
      }
      result =
        selectFound &&
        columnFound &&
        fromFound &&
        modelFound &&
        whereFound &&
        resourceFound &&
        filterFound &&
        ulqFound &&
        resNameFound &&
        resTypeFound;

      queryResult.errorMessage =
        queryResult.errorMessage +
        getErrorMessage(
          selectFound,
          columnFound,
          fromFound,
          modelFound,
          whereFound,
          resourceFound,
          filterFound,
          ulqFound,
          resNameFound,
          resTypeFound
        );
    }
  }
  if (!result) {
    queryResult.errorMessage =
      queryResult.errorMessage.length > 0
        ? 'SQL syntax errors on [' + queryResult.errorMessage + ']'
        : 'SQL syntax errors';
  }
  queryResult.result = result;
  return queryResult;
};

export const chkUnwantedField = (fields: string, queryParts: string[]): FieldValidationResilt => {
  let result = true;
  let fieldlist = fields.split(',');
  let andRequired: Boolean = false; // AND will define between each query where parameters
  let resultMessage = '';
  if (fieldlist.length > 0 && queryParts.length > 4) {
    queryParts.forEach((item, index) => {
      if (index > 4) {
        if (!andRequired && item.toLocaleUpperCase().trim() !== 'AND') {
          let lhs = item.trim().split('=')[0].trim();
          const fieldtoSearch = (element: string) => {
            if (element.trim().split('=')[0].toLocaleUpperCase().trim() === lhs.toLocaleUpperCase()) {
              return true;
            } else {
              return false;
            }
          };
          let idx = fieldlist.findIndex(fieldtoSearch);
          result = idx === -1 ? false : result;
          resultMessage = resultMessage + result ? item.trim() : '';
        } else if (
          (!andRequired && item.toLocaleUpperCase().trim() === 'AND') ||
          (andRequired && item.toLocaleUpperCase().trim() !== 'AND')
        ) {
          result = false;
          resultMessage = resultMessage + result ? 'AND' : '';
        }
        andRequired = !andRequired;
      }
    });
  }
  return { result: result, strResult: result ? '' : resultMessage };
};

const fieldValidation = (field: string, queryParts: string[]): FieldValidationResilt => {
  let result = true;
  let strResult = '';
  const fieldtoSearch = (element: string) => element.toLocaleUpperCase().indexOf(field.toLocaleUpperCase() + '=') > -1;
  let index = queryParts.findIndex(fieldtoSearch);
  if (index < 0 || queryParts[index].trim().split('=')[0].toLocaleUpperCase().trim() !== field.toLocaleUpperCase()) {
    result = false;
  } else {
    strResult = queryParts[index].trim().split('=')[1].trim();
    if (strResult.match(/\".*\"/g) !== null || strResult.match(/\'.*\'/g) !== null) {
      strResult = strResult.replace(/['"]+/g, '').trim();
    } else {
      result = false;
    }
  }
  return { result: result, strResult: strResult };
};

const fiendLHSField = (field: string, queryParts: string[]): boolean => {
  const fieldtoSearch = (element: string) => {
    if (
      element.trim().indexOf('=') > -1 &&
      element.trim().split('=')[0].toLocaleUpperCase().trim() === field.toLocaleUpperCase()
    ) {
      return true;
    } else {
      return false;
    }
  };
  let index = queryParts.findIndex(fieldtoSearch);
  return index > -1;
};

const getErrorMessage = (
  selectFound: Boolean,
  columnFound: Boolean,
  fromFound: Boolean,
  modelFound: Boolean,
  whereFound: Boolean,
  resourceFound: Boolean,
  filterFound: Boolean,
  ulqFound: Boolean,
  resNameFound: Boolean,
  resTypeFound: Boolean
): string => {
  let resultMessage = '';
  let result =
    selectFound &&
    columnFound &&
    fromFound &&
    modelFound &&
    whereFound &&
    resourceFound &&
    filterFound &&
    ulqFound &&
    resNameFound &&
    resTypeFound;
  if (!result) {
    resultMessage = resultMessage + selectFound ? ',SELECT' : '';
    resultMessage = resultMessage + columnFound ? ',<Column name>' : '';
    resultMessage = resultMessage + fromFound ? ',FROM' : '';
    resultMessage = resultMessage + modelFound ? ',<model name>' : '';
    resultMessage = resultMessage + whereFound ? ',WHERE' : '';
    resultMessage = resultMessage + filterFound ? ',FILTER' : '';
    resultMessage = resultMessage + ulqFound ? ',ULQ' : '';
    resultMessage = resultMessage + resNameFound ? ',Name' : '';
    resultMessage = resultMessage + resTypeFound ? ',TYPE' : '';
    resultMessage = resultMessage + resourceFound ? ',<Type name>' : '';
  }
  return resultMessage;
};

export const getAllResKeyPath = (parent: any, resource: string, path = ''): string => {
  Object.entries(parent).forEach(([keyIndex, KeyValue]) => {
    if (keyIndex.toUpperCase() !== 'RESOURCE_PATH' && keyIndex.toUpperCase() !== 'RESOURCE_ALL_PATH') {
      if (keyIndex.trim().toUpperCase() === resource.trim().toUpperCase()) {
        path = (KeyValue as any).RESOURCE_ALL_PATH !== undefined ? (KeyValue as any).RESOURCE_ALL_PATH : '';
        return path;
      }
      path = getAllResKeyPath(KeyValue as any, resource, path);
    }
    return path;
  });
  return path;
};

export const getResKeyPath = (parent: any, resource: string, path = '') => {
  Object.entries(parent).forEach(([keyIndex, value]) => {
    if (keyIndex.toUpperCase() !== 'RESOURCE_PATH' && keyIndex.toUpperCase() !== 'RESOURCE_ALL_PATH') {
      if (keyIndex.trim().toUpperCase() === resource.trim().toUpperCase()) {
        path = (value as any).RESOURCE_PATH !== undefined ? (value as any).RESOURCE_PATH : '';
        return path;
      }
      path = getResKeyPath(value as any, resource, path);
    }
    return path;
  });
  return path;
};

export const getResourceName = (result: any): any => {
  let resource = result !== undefined ? result.containedResourcesList[0].contained.resource : '';
  if (resource === null || resource === undefined || resource.length < 1) {
    resource = [];
  }
  return Array.isArray(resource) ? resource : [resource];
};

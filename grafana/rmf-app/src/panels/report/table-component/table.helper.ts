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
import {
  applyFieldOverrides,
  applyRawFieldOverrides,
  ConfigOverrideRule,
  DataFrame,
  Field,
  FieldColor,
  FieldConfigSource,
  PanelData,
  ThresholdsConfig,
} from '@grafana/data';
import { config } from '@grafana/runtime';
import { TableFieldOptions } from '@grafana/ui';
import { DataListItem, defaultThresholds, headerSplitKey, TableBanner } from '../types';

// const ThemeContext = React.createContext(getTheme());
// let useTheme = (): GrafanaTheme => {
//   let ThemeContextMock: React.Context<GrafanaTheme> | null = null;
//   return useContext(ThemeContextMock || ThemeContext);
// };
// const theme = useTheme();

export const InitFrameData = (data: PanelData): DataFrame[] => {
  let frameData: DataFrame[] = [
    {
      fields: [],
      length: 0,
    },
  ];

  if (
    data !== undefined &&
    data.series !== undefined &&
    data.series.length > 0 &&
    data.series[0].fields !== undefined
  ) {
    frameData = [
      {
        fields: data.series[0].fields,
        length: data.series[0].fields[0].values.length,
      } as DataFrame,
    ] as DataFrame[];
  }

  return frameData;
};
export const applyNearestPercentage = (field: Field, maxVal: number): Field => {
  if (field.type === 'number') {
    field.config.min = 0;
    (field.state as any).range.min = 0;
    field.config.max = maxVal;
    (field.state as any).range.max = field.config.max;
  }
  return field;
};

export const applySelectedDefaultsAndOverrides = (
  options: any,
  fieldConfig: FieldConfigSource,
  data: DataFrame[]
): DataFrame[] => {
  let result: DataFrame[] = applyRawFieldOverrides(data);
  let removeHeadersIndex = -1;
  removeHeadersIndex = result[0].fields.findIndex((col: any) => col.name.indexOf(headerSplitKey) !== -1);
  while (removeHeadersIndex > -1) {
    result[0].fields.splice(removeHeadersIndex, 1);
    removeHeadersIndex = result[0].fields.findIndex((col: any) => col.name.indexOf(headerSplitKey) !== -1);
  }

  // First apply default settings
  result[0].fields.map((field: Field) => {
    if (fieldConfig.defaults.thresholds !== undefined) {
      field.config.thresholds = fieldConfig.defaults.thresholds;
    }
    if (field.config.thresholds === undefined) {
      field.config = { ...field.config, ...{ thresholds: defaultThresholds } };
    }
    field.config.custom = {
      align: 'auto',
      filterable: fieldConfig.defaults.custom.filterable,
      cellOptions: fieldConfig.defaults.custom.cellOptions.type ? fieldConfig.defaults.custom.cellOptions : { type: fieldConfig.defaults.custom.cellOptions },
    } as TableFieldOptions;
  });

  // apply overrides
  if (fieldConfig.overrides && fieldConfig.overrides.length > 0) {
    fieldConfig.overrides.map((ovItem: ConfigOverrideRule) => {
      if (ovItem.matcher.id === 'byName') {
        ovItem.properties.map((ovrProp) => {
          if (ovrProp.id === 'custom.cellOptions') {
            result[0].fields.map((field: Field, index: number) => {
              if (field.name === ovItem.matcher.options) {
                field = applyNearestPercentage(field, 100);
                field.config.custom = {
                  cellOptions: ovrProp.value.type ? ovrProp.value : { type: ovrProp.value },
                } as TableFieldOptions;
              }
            });
          }
          if (ovrProp.id === 'custom.filterable') {
            result[0].fields.map((field: Field, index: number) => {
              if (field.name === ovItem.matcher.options) {
                if (field.config.custom !== undefined && field.config.custom.cellOptions !== undefined) {
                  field.config.custom = {
                    cellOptions: field.config.custom.cellOptions,
                    filterable: ovrProp.value,
                  } as TableFieldOptions;
                } else {
                  field.config.custom = {
                    filterable: ovrProp.value,
                  } as TableFieldOptions;
                }
              }
            });
          }
          if (ovrProp.id === 'color') {
            result[0].fields.map((field: Field) => {
              if (field.name === ovItem.matcher.options && ovrProp && ovrProp.value && ovrProp.value.mode) {
                field.config.color = {
                  mode: ovrProp.value.mode,
                  fixedColor: ovrProp.value.fixedColor,
                } as FieldColor;
              }
            });
          }
          if (ovrProp.id === 'thresholds') {
            result[0].fields.map((field: Field) => {
              if (field.name === ovItem.matcher.options && ovrProp && ovrProp.value && ovrProp.value.mode) {
                field.config.thresholds = {
                  mode: ovrProp.value.mode,
                  steps: ovrProp.value.steps,
                } as ThresholdsConfig;
              }
            });
          }
        });
      }
    });
  }
  return result;
};

export const prepareBannerToDisplay = (data: any): TableBanner => {
  let bannerItem: TableBanner = {};
  // This is for old rmf-datasource support : will be remove
  if (data.series && data.series[0] && data.series[0].meta && data.series[0].meta.tableBanner) {
    bannerItem = data.series[0].meta.tableBanner as TableBanner;
  }
  // ----------
  if (data.series && data.series[0] && data.series[0].meta && data.series[0].meta.custom) {
    bannerItem = {
      samples: data.series[0].meta.custom['numSamples'],
      system:
        data.series[0].meta.custom['numSystems'] !== undefined && data.series[0].meta.custom['numSystems'] !== ''
          ? data.series[0].meta.custom['numSystems']
          : '',
      timeRange:
        data.series[0].meta.custom['localEnd'] !== undefined && data.series[0].meta.custom['localEnd'] !== ''
          ? stringToDateString(data.series[0].meta.custom['localStart']) +
            ' - ' +
            stringToDateString(data.series[0].meta.custom['localEnd'])
          : stringToDateString(data.series[0].meta.custom['localStart']),
      range:
        data.series[0].meta.custom.dataRange !== undefined &&
        data.series[0].meta.custom.dataRange.value !== undefined &&
        data.series[0].meta.custom.dataRange.value > 0
          ? data.series[0].meta.custom.dataRange.value
          : '',
    };
  }
  return bannerItem;
};

export const stringToDateString = (serviceResult: String): string => {
  //convert result from ex 20230501071820 YYYYMMDDHHMMSS to MM/DD/YYYY HH:MM:SS
  let result =
    serviceResult.toString().substring(4, 6) +
    '/' +
    serviceResult.toString().substring(6, 8) +
    '/' +
    serviceResult.toString().substring(0, 4) +
    '/ ' +
    serviceResult.toString().substring(8, 10) +
    ':' +
    serviceResult.toString().substring(10, 12) +
    ':' +
    serviceResult.toString().substring(12, 14);
  return result;
};

export const prepareCaptionListToDisplay = (data: DataFrame[]): DataListItem[] => {
  let captionList: DataListItem[] = [];
  let result = applyRawFieldOverrides(data);
  let removeHeadersIndex = result[0].fields.filter((col: any) => col.name.indexOf(headerSplitKey) !== -1);

  if (removeHeadersIndex && removeHeadersIndex.length > 0) {
    removeHeadersIndex.map((field: any) => {
      if (field && field.config && field.config.displayName && field.config.displayName.length > 0) {
        captionList.push({ name: field.config.displayName, value: field.values.buffer[0] });
      } else {
        captionList.push({ name: field.name.replace(headerSplitKey, ''), value: field.values.buffer[0] });
      }
    });
  }
  return captionList;
};

export const applyFieldOverridesForBarGauge = (finalData: DataFrame[]): DataFrame[] => {
  return applyFieldOverrides({
    data: finalData,
    fieldConfig: {
      overrides: [],
      defaults: {},
    },
    replaceVariables: (value: string) => value,
    theme: config.theme2 as any,
  });
};

export const getPaginationFlagFromFieldConfig = (fieldConfig: FieldConfigSource): boolean => {
  let result: boolean;
  if (
    fieldConfig !== undefined &&
    fieldConfig.defaults !== undefined &&
    fieldConfig.defaults.custom !== undefined &&
    fieldConfig.defaults.custom.enablePagination !== undefined
  ) {
    result = fieldConfig.defaults.custom.enablePagination;
  } else {
    result = false;
  }
  return result;
};

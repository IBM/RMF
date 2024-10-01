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
import { defaultThresholds, CAPTION_PREFIX, BANNER_PREFIX, ReportData } from '../types';

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

export const applySelectedDefaultsAndOverrides = (options: any, fieldConfig: FieldConfigSource, data: DataFrame[]): ReportData => {
  // FIXME: send banner, captions and table data in different frames.
  let result = applyRawFieldOverrides(data);
  let bannerFields: Field[] = [];
  let captionFields: Field[] = [];
  let tableFields: Field[] = []
  for (let i = 0; i < result[0].fields.length; i++) {
    let field = result[0].fields[i]
    if (field.name.startsWith(BANNER_PREFIX)) {
      bannerFields.push(field);
    } else if (field.name.startsWith(CAPTION_PREFIX)) {
      captionFields.push(field);
    } else {
      tableFields.push(field);
    }
  }
  result[0].fields = tableFields;

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

  // Apply overrides
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
  return { bannerFields: bannerFields, captionFields: captionFields, tableData: result};
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

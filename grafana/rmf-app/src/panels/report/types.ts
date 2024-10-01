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
import { SelectableValue, ThresholdsConfig, ThresholdsMode, Field, DataFrame } from '@grafana/data';
import { TableCellDisplayMode } from '@grafana/ui';

export interface ReportData {
  bannerFields: Field[];
  captionFields: Field[];
  tableData: DataFrame[];
}

export interface FieldProps {
  fields: Field[]
}

export const BANNER_PREFIX= 'Banner::';
export const CAPTION_PREFIX = 'Caption::';

export const displayModes: Array<SelectableValue<any>> = [
  { value: TableCellDisplayMode.Auto, label: 'Auto' },
  { value: TableCellDisplayMode.ColorText, label: 'Colored text' },
  { value: TableCellDisplayMode.ColorBackground, label: 'Colored background' },
  { value: TableCellDisplayMode.Gauge, label: 'Gauge' },
  { value: TableCellDisplayMode.JSONView, label: 'JSON View' },
  { value: TableCellDisplayMode.Image, label: 'Image' },
]

export const defaultThresholds: ThresholdsConfig = {
  steps: [
    { value: -Infinity, color: 'green', state: 'HighHigh' },
    { value: 70, color: 'orange', state: 'OK' },
    { value: 90, color: 'red', state: 'High' },
  ],
  mode: ThresholdsMode.Absolute,
};

export interface TableBanner {
  samples?: string;
  systems?: string;
  timeRange?: string;
  range?: string;
}

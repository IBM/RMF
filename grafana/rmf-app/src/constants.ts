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

import appPluginJson from './plugin.json';
import dataSourcePluginJson from './datasources/rmf-datasource/plugin.json';

export const APP_NAME = appPluginJson.name;
export const APP_DESC = appPluginJson.info.description;
export const APP_LOGO = appPluginJson.info.logos.large;
export const APP_LOGO_URL = `public/plugins/${appPluginJson.id}/${appPluginJson.info.logos.large}`;
export const APP_BASE_URL = `/a/${appPluginJson.id}`;
export const DATA_SOURCE_TYPE = dataSourcePluginJson.id;
export const DATA_SOURCE_NAME = dataSourcePluginJson.name;
export const DDS_OPEN_METRICS_DOC_URL =
  'https://www.ibm.com/docs/en/zos/3.1.0?topic=functions-setting-up-distributed-data-server-zos';

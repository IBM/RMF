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
import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
// import { ConfigEditor } from './config-editor/config-editor.component';
import { RMFQuery, RMFDataSourceJsonData } from './common/types';
import { QueryEditorAutoCompleteComponent } from './query-editor/queryeditor.parser.component';
import ConfigEditor from './config-editor/config-editor.component';

export const plugin = new DataSourcePlugin<DataSource, RMFQuery, RMFDataSourceJsonData>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditorAutoCompleteComponent);

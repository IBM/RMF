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
import { PanelPlugin } from '@grafana/data';
import { displayModes } from './types';
import { MainComponent } from './main.component';

export const plugin = new PanelPlugin<{}>(MainComponent)
  .setPanelOptions((builder) => {
    return builder;
  })
  .useFieldConfig({
    useCustomConfig: (builder) => {
      return builder
        .addSelect({
          path: 'cellOptions',
          name: 'Cell Type',
          description: 'Color text, background, show as gauge, etc',
          settings: {
            // @ts-ignore
            options: displayModes
          },
          defaultValue: 'Auto'
        })
        .addBooleanSwitch({
          path: 'enablePagination',
          name: 'Enable pagination',
          description: '',
          defaultValue: false,
        })
        .addBooleanSwitch({
          path: 'filterable',
          name: 'Column filter',
          description: 'Enables/disables field filters in table',
          defaultValue: false,
        });
    },
  });

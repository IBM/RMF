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
import { AppPlugin } from '@grafana/data';
import { Config, Root } from './components';
import { GlobalSettings } from './types';

/**
 * Application Plugin
 */
export const plugin = new AppPlugin<GlobalSettings>().setRootPage(Root).addConfigPage({
  title: 'Config',
  body: Config,
  id: 'config',
});

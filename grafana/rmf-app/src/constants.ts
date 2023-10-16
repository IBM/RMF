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
/**
 * Data Source types
 */
export enum DataSourceType {
  RMFTYPE = 'ibm-rmf-datasource',
}

/**
 * New Data Source names
 */
export enum DataSourceName {
  RMFNAME = 'IBM RMF for z/OS',
}

/**
 * RMF commands
 */
export enum RmfCommand {
  COMMAND = 'command',
}

/**
 * Client Type Values
 */
export enum ClientTypeValue {
  CLUSTER = 'cluster',
  SENTINEL = 'sentinel',
  SOCKET = 'socket',
  STANDALONE = 'standalone',
}

/**
 * Application root page
 */
export const ApplicationRoot = '/a/ibm-rmf';

/**
 * Application
 */
export const ApplicationName = 'IBM RMF';
export const ApplicationSubTitle = 'IBM RMF Plugins Manager';

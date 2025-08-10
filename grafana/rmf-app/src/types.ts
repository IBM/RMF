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
import { DataQuery, DataSourceInstanceSettings, DataSourceJsonData } from '@grafana/data';

/**
 * Global Settings
 */
export type GlobalSettings = object;

/**
 * SVG
 */
export interface SVGProps extends React.HTMLAttributes<SVGElement> {
  /**
   * Size
   *
   * @type {number}
   */
  size: number;

  /**
   * Fill color
   *
   * @type {string}
   */
  fill?: string;

  /**
   * Title
   *
   * @type {string}
   */
  title?: string;

  /**
   * Class Name
   *
   * @type {string}
   */
  className?: string;
}

/**
 * Options configured for each DataSource instance
 */
export interface RmfDataSourceOptions extends DataSourceJsonData {
  path: string;
  port: string;
  /**
   * Pool Size
   *
   * @type {number}
   */
  poolSize: number;

  /**
   * Timeout
   *
   * @type {number}
   */
  timeout: number;

  /**
   * Pool Ping Interval
   *
   * @type {number}
   */
  pingInterval: number;

  /**
   * Pool Pipeline Window
   *
   * @type {number}
   */
  pipelineWindow: number;

  /**
   * TLS Authentication
   *
   * @type {boolean}
   */
  tlsAuth: boolean;

  /**
   * TLS Skip Verify
   *
   * @type {boolean}
   */
  tlsSkipVerify: boolean;

  /**
   * Sentinel Master group name
   *
   * @type {string}
   */
  sentinelName: string;

  /**
   * ACL enabled
   *
   * @type {boolean}
   */
  acl: boolean;

  /**
   * CLI disabled
   *
   * @type {boolean}
   */
  cliDisabled: boolean;

  /**
   * ACL Username
   *
   * @type {string}
   */
  user: string;
}

/**
 * Instance Settings
 */
export interface RmfDataSourceInstanceSettings extends DataSourceInstanceSettings<RmfDataSourceOptions> {
  /**
   * Commands
   *
   * @type {string[]}
   */
  commands: string[];
}

/**
 * RMF Query
 */
export interface RmfQuery extends DataQuery {
  /**
   * Query command
   *
   * @type {string}
   */
  query?: string;

  /**
   * RMF Command type
   *
   * @type {string}
   */
  type?: string;

  /**
   * RMF Command
   *
   * @type {string}
   */
  command?: string;

  /**
   * RMF Section
   *
   * @type {string}
   */
  section?: string;

  /**
   * CLI or Raw mode
   *
   * @type {boolean}
   */
  cli?: boolean;
}

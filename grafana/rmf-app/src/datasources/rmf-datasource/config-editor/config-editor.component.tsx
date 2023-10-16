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
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { LegacyForms, PanelContainer } from '@grafana/ui';
import React, { ChangeEvent, PureComponent, SyntheticEvent } from 'react';
import { RMFDataSourceJsonData as RMFDataSourceOptions, RMFDatasourceSecureJsonData } from '../common/types';

require('./config-editor.component.css');

const { FormField, Switch } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<RMFDataSourceOptions, RMFDatasourceSecureJsonData> { }

interface State { }

export default class ConfigEditor extends PureComponent<Props, State> {
  constructor(props: DataSourcePluginOptionsEditorProps<RMFDataSourceOptions, RMFDatasourceSecureJsonData>) {
    super(props);
    const { options, onOptionsChange } = props;
    const { jsonData, secureJsonData } = options;

    if (jsonData === undefined || jsonData.path === undefined) {
      const jsonData = {
        ...options.jsonData,
      };
      onOptionsChange({ ...options, jsonData });
    }
    if (secureJsonData === undefined) {
      const secureJsonData = {
        ...options.secureJsonData,
      };
      onOptionsChange({ ...options, secureJsonData });
    }
  }

  onPathChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      path: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onPortChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      port: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onSSLCheckChange = (event: SyntheticEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    // remove username and password if SSL is unselected
    const jsonData = {
      ...options.jsonData,
      ssl: event.currentTarget.checked,
      userName: '',
      skipVerify: event.currentTarget.checked,
    };
    const secureJsonData = {
      ...options.secureJsonData,
      password: '',
    };
    onOptionsChange({...options, jsonData: jsonData, secureJsonData: secureJsonData});
  };

  onSkipVerifySCheckChange = (event: SyntheticEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    // remove username and password if SSL is unselected
    const jsonData = {
      ...options.jsonData,
      skipVerify: (options.jsonData.ssl && options.jsonData.ssl === true) ? event.currentTarget.checked : false,
    };
    onOptionsChange({ ...options, jsonData });
  };

  onUsernameChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      userName: options.jsonData.ssl ? event.target.value : '',
    };
    onOptionsChange({ ...options, jsonData });
  };

  onPasswordChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: {
        ...options.secureJsonData,
        password: options.jsonData.ssl ? event.target.value : '',
      },
    });
  };

  render() {
    const { options } = this.props;
    const { jsonData, secureJsonData = {}, secureJsonFields = {} } = options;

    return (
      <div className="gf-form-group">
        <label className="width-12 auth-label">Server Information:</label>
        <div className="gf-form">
          <FormField
            label="Server"
            labelWidth={6}
            inputWidth={20}
            onChange={this.onPathChange}
            value={jsonData.path || ''}
          />
        </div>
        <div className="gf-form">
          <FormField
            label="Port"
            labelWidth={6}
            inputWidth={20}
            onChange={this.onPortChange}
            value={jsonData.port || ''}
          />
        </div>
        <PanelContainer>
          <label className="width-12 auth-label">Basic Authentication:</label>
          <div className="gf-form">
            <Switch
              label="Use SSL"
              checked={jsonData.ssl || false}
              tooltip="Enable secure communications"
              onChange={this.onSSLCheckChange}
            />
          </div>
          <div className="gf-form">
            <Switch
              label="Verify the server's certificate chain/host name"
              checked={jsonData.skipVerify === undefined ? false : jsonData.skipVerify}
              tooltip="If disabled the Datasource accepts any certificate presented by the server and any host name in that certificate. This is considered insecure."
              onChange={this.onSkipVerifySCheckChange}
              disabled={!jsonData.ssl}
            />
          </div>
          <div className="gf-form">
            <FormField
              label="Username"
              labelWidth={6}
              inputWidth={20}
              onChange={this.onUsernameChange}
              value={jsonData.userName || ''}
              disabled={!jsonData.ssl}
            />
            <FormField
              type="password"
              label="Password"
              labelWidth={6}
              inputWidth={20}
              onChange={this.onPasswordChange}
              value={
                secureJsonData && (secureJsonData as any).password !== undefined ? (secureJsonData as any).password : ''
              }
              placeholder={
                secureJsonFields && (secureJsonFields as any).password !== undefined
                  ? jsonData.ssl
                    ? '*********************'
                    : ''
                  : ''
              }
              disabled={!jsonData.ssl}
            />
          </div>
        </PanelContainer>
      </div>
    );
  }
}

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
import React, { PureComponent } from 'react';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { FieldValidationMessage, InlineField, InlineSwitch, LegacyForms } from '@grafana/ui';
import { RMFDataSourceSettings, RMFDataSourceJsonData, RMFDataSourceSecureJsonData } from '../common/types';

require('./config-editor.component.css');

const { FormField } = LegacyForms;
const FIELD_LABEL_WIDTH = 10;
const FIELD_INPUT_WIDTH= 20;
const SWITCH_LABEL_WIDTH = 20;
const DEFAULT_HTTP_TIMEOUT = "60";
const DEFAULT_CACHE_SIZE = "1024";

type Props = DataSourcePluginOptionsEditorProps<RMFDataSourceJsonData, RMFDataSourceSecureJsonData>;

interface State {
  urlError?: string;
  httpTimeoutError?: string
  basicAuthUserError?: string
  cacheSizeError?: string
}


export default class ConfigEditor extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    const {options, onOptionsChange} = props;
    const {jsonData, secureJsonData} = options;
    this.state = {}

    if (jsonData?.path !== undefined || jsonData?.port !== undefined) {
      // Convert from the legacy format if possible
      options.url = `http${jsonData?.ssl? "s": ""}://${jsonData.path}${jsonData?.port? ":" + jsonData.port: ""}`;
      let username = (jsonData?.userName || "").trim();
      if (username) {
        options.basicAuth = true;
        options.basicAuthUser = username;
      }
      options.jsonData = {
        timeout: DEFAULT_HTTP_TIMEOUT,
        tlsSkipVerify: jsonData?.skipVerify || false
      };
      options.secureJsonData = { basicAuthPassword: secureJsonData?.password || "" }
    } else {
      // Initialize jsonData and secureJsonData if needed
      options.jsonData = {
        timeout: jsonData?.timeout || DEFAULT_HTTP_TIMEOUT,
        cacheSize: jsonData?.cacheSize|| DEFAULT_CACHE_SIZE,
        tlsSkipVerify: jsonData?.tlsSkipVerify || false
      };
      options.secureJsonData = { basicAuthPassword: secureJsonData?.basicAuthPassword || "" }
    }
    onOptionsChange({...options});
  }

  urlRegex = /^(http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$/

  validateUrl = (url: string) => {
    if (this.urlRegex.test(url)) {
      this.setState({ urlError: undefined });
    } else {
      this.setState({ urlError: 'Invalid DDS URL' });
    }
  }

  validateHttpTimeout = (value: string) => {
    let numValue= Number(value)
    if (numValue > 0 && Number.isInteger(numValue)) {
      this.setState({ httpTimeoutError: undefined });
    } else {
      this.setState({ httpTimeoutError: 'Timeout must be a positive integer' });
    }
  }

  validateUsername = (value: string) => {
    if (value?.trim().length > 0) {
      this.setState({ basicAuthUserError: undefined });
    } else {
      this.setState({ basicAuthUserError: 'Username cannot be blank' });
    }
  }

  validateCacheSize = (value: string) => {
    let numValue= Number(value)
    if (numValue >= 128 && Number.isInteger(numValue)) {
      this.setState({ cacheSizeError: undefined });
    } else {
      this.setState({ cacheSizeError: 'Cache size must be ≥ 128' });
    }
  }

  updateSettings = (updates: Partial<RMFDataSourceSettings>) => {
    const { options, onOptionsChange } = this.props;
    onOptionsChange(
      { ...options, ...updates,
        jsonData: {...options.jsonData, ...updates.jsonData},
        secureJsonData: {...options.secureJsonData, ...updates.secureJsonData},
      }
    );
  };

  render() {
    const {options} = this.props;
    const {
      urlError,
      httpTimeoutError,
      basicAuthUserError,
      cacheSizeError
    } = this.state;

    return (
      <div className="gf-form-group">
        <h3 className="page-heading">HTTP</h3>
        <div className="gf-form-group">
          <div className="gf-form">
            <FormField
              label={'DDS URL'}
              labelWidth={FIELD_LABEL_WIDTH}
              tooltip="DDS URL needs to be accessible from the grafana backend/server."
              placeholder="https://dds-server:8803"
              inputWidth={FIELD_INPUT_WIDTH}
              value={options.url}
              onChange={(event) => {
                this.updateSettings({url: event.currentTarget.value})
              }}
              onBlur={(event) => {
                this.validateUrl(event.currentTarget.value)
              }}
            />
            {urlError && (
              <FieldValidationMessage horizontal={true}>{urlError}</FieldValidationMessage>
            )}
          </div>
          <div className="gf-form">
            <FormField
              label="Timeout"
              labelWidth={FIELD_LABEL_WIDTH}
              tooltip="HTTP request timeout in seconds"
              placeholder="Timeout in seconds"
              inputWidth={FIELD_INPUT_WIDTH}
              value={options.jsonData?.timeout}
              onChange={(event) => {
                this.updateSettings({jsonData: {timeout: event.currentTarget.value}})
              }}
              onBlur={(event) => {
                this.validateHttpTimeout(event.currentTarget.value)
              }}
            />
            {httpTimeoutError && (
              <FieldValidationMessage horizontal={true}>{httpTimeoutError}</FieldValidationMessage>
            )}
          </div>
        </div>
        <h3 className="page-heading">Auth</h3>
        <div className="gf-form-group">
          <div className="gf-form-inline">
            <InlineField label="Basic Auth"
                         labelWidth={SWITCH_LABEL_WIDTH}
                         tooltip="Use the basic authentication method">
              <InlineSwitch
                value={options.basicAuth}
                onChange={(event) => {
                  this.updateSettings(
                    {
                      basicAuth: event.currentTarget.checked,
                      basicAuthUser: "",
                      secureJsonData: {basicAuthPassword: ""}
                    })
                  this.setState({basicAuthUserError: undefined});
                }}
              />
            </InlineField>
          </div>
          <div className="gf-form-inline">
            <InlineField label="Skip TLS Verify"
                         labelWidth={SWITCH_LABEL_WIDTH}
                         tooltip="Skip TLS certificate validation">
              <InlineSwitch
                value={options.jsonData?.tlsSkipVerify}
                onChange={(event) => {
                  this.updateSettings({jsonData: {tlsSkipVerify: event.currentTarget.checked}})
                }}
              />
            </InlineField>
          </div>
        </div>

        {(options.basicAuth) && (
          <>
            <h6>Basic Auth Details</h6>
            <div className="gf-form-group">
              <div className="gf-form">
                <FormField
                  label="User"
                  labelWidth={FIELD_LABEL_WIDTH}
                  placeholder="User"
                  inputWidth={FIELD_INPUT_WIDTH}
                  value={options.basicAuthUser}
                  onChange={(event) => {
                    this.updateSettings({basicAuthUser: event.currentTarget.value})
                  }}
                  onBlur={(event) => {
                    this.validateUsername(event.currentTarget.value)
                  }}
                />
                {basicAuthUserError && (
                  <FieldValidationMessage
                    horizontal={true}>{basicAuthUserError}</FieldValidationMessage>
                )}
              </div>
              <div className="gf-form">
                <FormField
                  label="Password"
                  labelWidth={FIELD_LABEL_WIDTH}
                  type="password"
                  placeholder="Password"
                  inputWidth={FIELD_INPUT_WIDTH}
                  value={options.secureJsonData?.basicAuthPassword}
                  onChange={(event) => {
                    this.updateSettings({secureJsonData: {basicAuthPassword: event.currentTarget.value}})
                  }}
                />
              </div>
            </div>
          </>
        )}

        <h3 className="page-heading">Caching</h3>
        <div className="gf-form-group">
          <div className="gf-form">
            <FormField
              label="Size"
              labelWidth={FIELD_LABEL_WIDTH}
              tooltip="Cache size in MB for the data source"
              placeholder="≥ 128"
              inputWidth={FIELD_INPUT_WIDTH}
              value={options.jsonData?.cacheSize}
              onChange={(event) => {
                this.updateSettings({jsonData: {cacheSize: event.currentTarget.value}})
              }}
              onBlur={(event) => {
                this.validateCacheSize(event.currentTarget.value)
              }}
            />
            {cacheSizeError && (
              <FieldValidationMessage horizontal={true}>{cacheSizeError}</FieldValidationMessage>
            )}
          </div>
        </div>

      </div>
    );
  }
}

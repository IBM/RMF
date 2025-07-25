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
import { FieldValidationMessage, InlineField, InlineSwitch, LegacyForms, SecretInput } from '@grafana/ui';
import { RMFDataSourceSettings, RMFDataSourceJsonData, RMFDataSourceSecureJsonData } from '../common/types';

require('./config-editor.component.css');

const { FormField } = LegacyForms;
const FIELD_LABEL_WIDTH = 10;
const FIELD_INPUT_WIDTH = 20;
const INLINE_LABEL_WIDTH = 20;
const PASSWORD_SET_WIDTH = 30;
const PASSWORD_EDIT_WIDTH = 40;
const DEFAULT_HTTP_TIMEOUT = '60';
const DEFAULT_CACHE_SIZE = '1024';
const MINIMAL_CACHE_SIZE = 128;

type Props = DataSourcePluginOptionsEditorProps<RMFDataSourceJsonData, RMFDataSourceSecureJsonData>;

interface State {
  urlError?: string;
  httpTimeoutError?: string;
  basicAuthUserError?: string;
  cacheSizeError?: string;
}
// TODO: somehow prometheus can validate fields from "run and test" in v11
export default class ConfigEditor extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);

    const { options, onOptionsChange } = props;
    const { jsonData } = options;
    this.state = {};

    if (jsonData?.path !== undefined || jsonData?.port !== undefined) {
      // Convert from the legacy format if possible
      options.url = `http${jsonData.ssl ? 's' : ''}://${jsonData.path}${jsonData.port ? ':' + jsonData.port : ''}`;
      const username = (jsonData.userName || '').trim();
      if (username) {
        options.basicAuth = true;
        options.basicAuthUser = username;
      }
      options.jsonData = {
        timeout: DEFAULT_HTTP_TIMEOUT,
        tlsSkipVerify: !(jsonData.skipVerify !== undefined ? jsonData.skipVerify : true), // NB: the meaning of skipVerify is inverted.
        disableCompression: jsonData.disableCompression ?? false,
      };
    } else {
      // Initialize jsonData if needed
      options.jsonData = {
        timeout: jsonData?.timeout || DEFAULT_HTTP_TIMEOUT,
        cacheSize: jsonData?.cacheSize || DEFAULT_CACHE_SIZE,
        tlsSkipVerify: jsonData?.tlsSkipVerify || false,
        disableCompression: jsonData?.disableCompression ?? false,
      };
    }
    onOptionsChange({ ...options });
  }

  urlRegex = /^(http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?$/;

  validateUrl = (url: string) => {
    if (this.urlRegex.test(url)) {
      this.setState({ urlError: undefined });
    } else {
      this.setState({ urlError: 'Invalid DDS URL' });
    }
  };

  validateHttpTimeout = (value: string) => {
    let numValue = Number(value);
    if (numValue > 0 && Number.isInteger(numValue)) {
      this.setState({ httpTimeoutError: undefined });
    } else {
      this.setState({ httpTimeoutError: 'Timeout must be a positive integer' });
    }
  };

  validateUsername = (value: string) => {
    if (value?.trim().length > 0) {
      this.setState({ basicAuthUserError: undefined });
    } else {
      this.setState({ basicAuthUserError: 'Username cannot be blank' });
    }
  };

  validateCacheSize = (value: string) => {
    let numValue = Number(value);
    if (numValue >= MINIMAL_CACHE_SIZE && Number.isInteger(numValue)) {
      this.setState({ cacheSizeError: undefined });
    } else {
      this.setState({ cacheSizeError: 'Cache size must be ≥ ' + MINIMAL_CACHE_SIZE });
    }
  };

  updateSettings = (updates: Partial<RMFDataSourceSettings>) => {
    const { options, onOptionsChange } = this.props;
    onOptionsChange({
      ...options,
      ...updates,
      jsonData: { ...options.jsonData, ...updates.jsonData },
    });
  };

  setPassword = (password: string) => {
    const { options, onOptionsChange } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: { basicAuthPassword: password },
      secureJsonFields: {},
    });
  };

  resetPassword = () => {
    const { options, onOptionsChange } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: { basicAuthPassword: '' },
      secureJsonFields: { basicAuthPassword: false },
    });
  };

  render() {
    const { options } = this.props;
    const { urlError, httpTimeoutError, basicAuthUserError, cacheSizeError } = this.state;
    const isPasswordSet = options.secureJsonFields?.basicAuthPassword || options.secureJsonFields?.password || false;

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
                this.updateSettings({ url: event.currentTarget.value });
              }}
              onBlur={(event) => {
                this.validateUrl(event.currentTarget.value);
              }}
            />
            {urlError && <FieldValidationMessage horizontal={true}>{urlError}</FieldValidationMessage>}
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
                this.updateSettings({ jsonData: { timeout: event.currentTarget.value } });
              }}
              onBlur={(event) => {
                this.validateHttpTimeout(event.currentTarget.value);
              }}
            />
            {httpTimeoutError && <FieldValidationMessage horizontal={true}>{httpTimeoutError}</FieldValidationMessage>}
          </div>
          <div className="gf-form-inline">
            <InlineField label="Compression" labelWidth={INLINE_LABEL_WIDTH} tooltip="Request HTTP compression">
              <InlineSwitch
                // Negative flags aren't good for UX.
                // However, it's more consistent to keep it consistent with Golang DisableCompression option in the backend.
                value={!options.jsonData?.disableCompression}
                onChange={(event) => {
                  this.updateSettings({ jsonData: { disableCompression: !event.currentTarget.checked } });
                }}
              />
            </InlineField>
          </div>
        </div>
        <h3 className="page-heading">Auth</h3>
        <div className="gf-form-group">
          <div className="gf-form-inline">
            <InlineField
              label="Basic Auth"
              labelWidth={INLINE_LABEL_WIDTH}
              tooltip="Use the basic authentication method"
            >
              <InlineSwitch
                value={options.basicAuth}
                onChange={(event) => {
                  this.updateSettings({
                    basicAuth: event.currentTarget.checked,
                    basicAuthUser: '',
                    secureJsonData: { basicAuthPassword: '' },
                  });
                  this.setState({ basicAuthUserError: undefined });
                }}
              />
            </InlineField>
          </div>
          <div className="gf-form-inline">
            <InlineField
              label="Skip TLS Verify"
              labelWidth={INLINE_LABEL_WIDTH}
              tooltip="Skip TLS certificate validation"
            >
              <InlineSwitch
                value={options.jsonData?.tlsSkipVerify}
                onChange={(event) => {
                  this.updateSettings({ jsonData: { tlsSkipVerify: event.currentTarget.checked } });
                }}
              />
            </InlineField>
          </div>
        </div>

        {options.basicAuth && (
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
                    this.updateSettings({ basicAuthUser: event.currentTarget.value });
                  }}
                  onBlur={(event) => {
                    this.validateUsername(event.currentTarget.value);
                  }}
                />
                {basicAuthUserError && (
                  <FieldValidationMessage horizontal={true}>{basicAuthUserError}</FieldValidationMessage>
                )}
              </div>
              <div className="gf-form">
                <InlineField label="Password" labelWidth={INLINE_LABEL_WIDTH}>
                  <SecretInput
                    label="Password"
                    // labelWidth={FIELD_LABEL_WIDTH}
                    type="password"
                    placeholder="Password"
                    width={isPasswordSet ? PASSWORD_SET_WIDTH : PASSWORD_EDIT_WIDTH}
                    isConfigured={isPasswordSet}
                    onChange={(event) => this.setPassword(event.currentTarget.value)}
                    onReset={() => this.resetPassword()}
                  />
                </InlineField>
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
                this.updateSettings({ jsonData: { cacheSize: event.currentTarget.value } });
              }}
              onBlur={(event) => {
                this.validateCacheSize(event.currentTarget.value);
              }}
            />
            {cacheSizeError && <FieldValidationMessage horizontal={true}>{cacheSizeError}</FieldValidationMessage>}
          </div>
        </div>
      </div>
    );
  }
}

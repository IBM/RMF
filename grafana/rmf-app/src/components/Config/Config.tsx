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
import React, { PureComponent } from 'react';
import { AppPluginMeta, PluginConfigPageProps } from '@grafana/data';
import { BackendSrv, getBackendSrv, getLocationSrv } from '@grafana/runtime';
import { Button } from '@grafana/ui';
import { APP_NAME, APP_BASE_URL } from '../../constants';
import { GlobalSettings } from '../../types';

type Props = PluginConfigPageProps<AppPluginMeta<GlobalSettings>>;
interface State {
  isEnabled: boolean;
}

export class Config extends PureComponent<Props, State> {
  private backendSrv: BackendSrv = getBackendSrv();

  constructor(props: Props) {
    super(props);
    this.state = {
      isEnabled: false,
    };
  }

  componentDidMount(): void {
    this.setState(() => ({
      isEnabled: this.props.plugin.meta?.enabled ? true : false,
    }));
  }

  goHome = (): void => {
    getLocationSrv().update({
      path: APP_BASE_URL,
      partial: false,
    });
  };

  updatePluginSettings = (settings: { enabled: boolean; jsonData: unknown; pinned: boolean }): Promise<undefined> => {
    return this.backendSrv.post(`api/plugins/${this.props.plugin.meta.id}/settings`, settings);
  };

  onDisable = () => {
    this.updatePluginSettings({ enabled: false, jsonData: {}, pinned: false }).then(() => {
      window.location.reload();
    });
  };

  onEnable = () => {
    this.updatePluginSettings({ enabled: true, jsonData: {}, pinned: true }).then(() => {
      window.location.reload();
    });
  };

  render() {
    const { isEnabled } = this.state;

    return (
      <>
        {!isEnabled && <p>Click below to enable the application.</p>}
        <p>Go to the {APP_NAME} App to install sample dashboards and get started.</p>
        <div className="gf-form gf-form-button-row">
          {isEnabled ? (
            <Button variant="destructive" onClick={this.onDisable}>
              Disable
            </Button>
          ) : (
            <Button onClick={this.onEnable}>Enable</Button>
          )}
          <Button disabled={!isEnabled} onClick={this.goHome}>
            Go to {APP_NAME} App
          </Button>
        </div>
      </>
    );
  }
}

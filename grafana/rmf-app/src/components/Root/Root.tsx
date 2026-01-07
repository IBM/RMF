/**
 * (C) Copyright IBM Corp. 2023, 2025.
 * (C) Copyright Rocket Software, Inc. 2023-2025.
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
import { PanelContainer, Button, TextLink, Box, Icon, Alert, InlineSwitch, Input, InlineField, DropzoneFile } from '@grafana/ui';
import { locationService, getBackendSrv, getDataSourceSrv, getAppEvents } from '@grafana/runtime';
import { AppRootProps, AppEvents } from '@grafana/data';

import { DDS_OPEN_METRICS_DOC_URL, DATA_SOURCE_TYPE, APP_LOGO_URL, FALCON_AS_DASHBOARD, FALCON_SYS_DASHBOARD, PM_LOGO_URL } from '../../constants';
import { GlobalSettings } from '../../types';
import { DASHBOARDS as DDS_DASHBOARDS } from '../../dashboards/dds';
import { DASHBOARDS as PROM_DASHBOARDS } from '../../dashboards/prometheus';
import { findFolder, deleteFolder, installDashboards, findDashboard } from './utils';
import { FolderStatus, Operation, OperCode, OperStatus, FalconStatus, PmStatus } from './types';
import { StatusIcon } from './StatusIcon';
import { Space } from './Space';
import { Header } from './Header';
import { FileDropzone } from '@grafana/ui';
import { parsePmDatasources, parsePmImportFileToDashboard } from './PmImport';

const DDS_FOLDER_UID = 'ibm-rmf-dds';
const DDS_FOLDER_NAME = 'IBM RMF (DDS)';
const PROM_FOLDER_UID = 'ibm-rmf-prometheus';
const PROM_FOLDER_NAME = 'IBM RMF (Prometheus)';
const PM_FOLDER_UID = 'ibm-rmf-pm';
const PM_FOLDER_NAME = 'IBM RMF (PM)';
const DATASOURCE_API = '/api/datasources';

interface Props extends AppRootProps<GlobalSettings> {}

interface State {
  dds: FolderStatus;
  prom: FolderStatus;
  falcon: FalconStatus;
  pm: PmStatus;
}

export class Root extends PureComponent<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      dds: {
        folderPath: DDS_FOLDER_NAME,
        installed: false,
        operation: { code: OperCode.None, status: OperStatus.None },
      },
      prom: {
        folderPath: PROM_FOLDER_NAME,
        installed: false,
        operation: { code: OperCode.None, status: OperStatus.None },
      },
      falcon: {
        enabled: false,
        asDashboard: FALCON_AS_DASHBOARD,
        sysDashboard: FALCON_SYS_DASHBOARD,
      },
      pm: {
        folderPath: PM_FOLDER_NAME,
        installed: false,
        operation: { code: OperCode.None, status: OperStatus.None },
      },
    };
  }

  async componentDidMount() {
    await this.updateFolderState();
  }

  updateFolderState = async () => {
    try {
      const ddsFolderPath = await findFolder(DDS_FOLDER_UID);
      const promFolderPath = await findFolder(PROM_FOLDER_UID);
      const pmFolderPath = await findFolder(PM_FOLDER_UID);
      const asDashboard = await findDashboard("Job CPU Details", ["omegamon", "zos", "lpar", "cpu"]);
      const sysDashboard = await findDashboard("z/OS Enterprise Overview", ["omegamon", "zos", "enterprise"]);
      this.setState((prevState) => ({
        dds: {
          ...prevState.dds,
          installed: ddsFolderPath !== undefined,
          folderPath: ddsFolderPath || DDS_FOLDER_NAME,
        },
        prom: {
          ...prevState.prom,
          installed: promFolderPath !== undefined,
          folderPath: promFolderPath || PROM_FOLDER_NAME,
        },
        falcon: {
          ...prevState.falcon,
          asDashboard: asDashboard !== undefined ? asDashboard : FALCON_AS_DASHBOARD,
          sysDashboard: sysDashboard !== undefined ? sysDashboard : FALCON_SYS_DASHBOARD,
        },
        pm: {
          ...prevState.pm,
          installed: pmFolderPath !== undefined,
          folderPath: pmFolderPath || PM_FOLDER_NAME,
        },
      }));
    } catch (error) {
      console.error('failed to update state', error);
    }
  };

  createDataSource = async () => {
    const { datasource } = await getBackendSrv().post(DATASOURCE_API, {
      type: DATA_SOURCE_TYPE,
      access: 'proxy',
    });
    await getDataSourceSrv().reload();
    locationService.push(`/datasources/edit/${datasource.uid}`);
  };

  goToFolder = async (folderUid: string, newInstall: boolean) => {
    if (newInstall) {
      // Don't use locationService: the only way to reload folder tree is to reload the page.
      // Official Grafana plugins do the same trick.
      window.location.assign(`/dashboards/f/${folderUid}`);
    } else {
      locationService.push(`/dashboards/f/${folderUid}`);
    }
  };

  setFolderState = (isDds: boolean, operation: Operation) => {
    this.setState((prevState) => ({
      ...prevState,
      [isDds ? 'dds' : 'prom']: {
        ...prevState[isDds ? 'dds' : 'prom'],
        operation: operation,
      },
    }));
  };

  processFolder = async (folderUid: string, operCode: OperCode, falcon: FalconStatus) => {
    const isDds = folderUid === DDS_FOLDER_UID;
    const defaultFolderName = isDds ? DDS_FOLDER_NAME : PROM_FOLDER_NAME;
    const dashboards = isDds ? DDS_DASHBOARDS : PROM_DASHBOARDS;

    this.setFolderState(isDds, {
      code: operCode,
      status: OperStatus.InProgress,
    });

    try {
      if (operCode === OperCode.Reset || operCode === OperCode.Delete) {
        await deleteFolder(folderUid);
      }
      if (operCode === OperCode.Reset || operCode === OperCode.Install) {
        await installDashboards(folderUid, defaultFolderName, dashboards, falcon);
      }
      this.setFolderState(isDds, {
        code: operCode,
        status: OperStatus.Done,
      });
    } catch (error) {
      this.setFolderState(isDds, {
        code: operCode,
        status: OperStatus.Error,
      });
      const appEvents = getAppEvents();
      appEvents.publish({
        type: AppEvents.alertError.name,
        payload: [`Unexpected error: ${error instanceof Error ? error.message : 'Unknown error'}`],
      });
    } finally {
      // There seems to be no way to refresh the folder tree without reloading the page.
      // When a folder deleted or created,
      // it won't be visible in dashboard tree until page is reloaded.
      await this.updateFolderState();
    }
  };

  importDashboard = async (folderUid: string, operCode: OperCode, dashboards: any[]) => {
    //const isDds = folderUid === DDS_FOLDER_UID;
    const defaultFolderName = PM_FOLDER_NAME;

    /*this.setFolderState(isDds, {
      code: operCode,
      status: OperStatus.InProgress,
    });*/
    var err;
    try {
      if (operCode === OperCode.Reset || operCode === OperCode.Delete) {
        await deleteFolder(folderUid);
      }
      if (operCode === OperCode.Reset || operCode === OperCode.Install) {
        await installDashboards(folderUid, defaultFolderName, dashboards, {enabled: false} as FalconStatus);
      }
      /*this.setFolderState(isDds, {
        code: operCode,
        status: OperStatus.Done,
      });*/
    } catch (error) {
      /*this.setFolderState(isDds, {
        code: operCode,
        status: OperStatus.Error,
      });*/
      const appEvents = getAppEvents();
      appEvents.publish({
        type: AppEvents.alertError.name,
        payload: [`Unexpected error: ${error instanceof Error ? error.message : 'Unknown error'}`],
      });
      err = error;
    } finally {
      // There seems to be no way to refresh the folder tree without reloading the page.
      // When a folder deleted or created,
      // it won't be visible in dashboard tree until page is reloaded.
      await this.updateFolderState();
    }
    if (!err) {
      const appEvents = getAppEvents();
      appEvents.publish({
        type: AppEvents.alertSuccess.name,
        payload: [`Dashboard imported: ${dashboards.length ? dashboards[0].title : 'Unknown'}`],
      });
    }
  };

  prepareDatasources = async (content: string | ArrayBuffer | null) : Promise<Map<string, string>> => {
    const datasources = await getDataSourceSrv().getList({type: DATA_SOURCE_TYPE});
    let nameUid = new Map<string, string>();
    try {
      const pmDatasources = parsePmDatasources(content);
      for (const pmDs of pmDatasources) {
        const existingDs = datasources.find(ds => ds.name === pmDs.name);
        if (!existingDs) {
          const { datasource } = await getBackendSrv().post(DATASOURCE_API, {
            type: DATA_SOURCE_TYPE,
            access: 'proxy',
            name: pmDs.name,
            url: `${pmDs.https ? 'https' : 'http'}://${pmDs.host}:${pmDs.port}`,
          });
          nameUid.set(pmDs.name, datasource.uid);
          getDataSourceSrv().reload();
        } else {
          nameUid.set(pmDs.name, existingDs.uid);
        }
      }
    } catch (error) {
      const appEvents = getAppEvents();
      appEvents.publish({
        type: AppEvents.alertError.name,
        payload: [`Unexpected error: ${error instanceof Error ? error.message : 'Unknown error'}`],
      });
    }
    return nameUid;
  }

  render() {
    const { dds, prom, falcon, pm } = this.state;
    const isBusy = dds.operation.status === OperStatus.InProgress || prom.operation.status === OperStatus.InProgress;

    return (
      // It's ScrollContainer, but it's available only in Grafana v12+
      <div style={{ overflow: 'auto', height: '100%', margin: '32px' }}>
        <Header />
        <Space layout={'block'} v={2} />
        {/* Use PanelContainer for Alert to avoid extra gaps on bigger screens */}
        <PanelContainer style={{ border: 'none' }}>
          <Alert title="What happens when I click 'Install Dashboards'?" severity="info" style={{ margin: '0' }}>
            Clicking on install will create a folder with a reserved UID and add sample dashboards to the folder. Both
            the folder and the dashboards are not owned by the plugin and can be managed by users.
            <br />
            <b>NB:</b> after the installation, you may need to reload the page to see the folders under the{' '}
            <i>Dashboards</i> section.
          </Alert>
        </PanelContainer>
        <Space layout={'block'} v={1} />
        <PanelContainer style={{ border: 'none' }}>
          <Alert title="What happens when I click 'Update / Reset'?" severity="info" style={{ margin: '0' }}>
            The destination folder having the reserved UID is deleted, and the sample dashboards are re-installed.
            <br />
            <b>NB:</b> if you change reserved folder title or move it, it will still have the same UID and will be
            reset.
          </Alert>
        </PanelContainer>
        <Space layout={'block'} v={3} />
        <Box backgroundColor={'secondary'}>
          <PanelContainer style={{ backgroundColor: 'rgba(255, 255, 255, 0)', width: '100%', padding: '2rem' }}>
            <h4>
              <img style={{ width: '48px', height: '48px' }} src={APP_LOGO_URL} alt="logo for IBM RMF" />
              <Space layout={'inline'} h={2} />
              DDS Sample Dashboards
            </h4>
            <br />
            <p>The plugin provides a way to access RMF Distributed Data Server (DDS) directly.</p>
            <p>Create an IBM RMF data source and install the sample dashboards to explore.</p>
            <p>
              Destination folder: <i>Dashboards / {dds.folderPath}</i>
              <Space layout={'inline'} h={2} />
              <span style={{ color: 'gray' }}>[UID=&apos;{DDS_FOLDER_UID}&apos;]</span>
            </p>
            <p>
              <p>
                Link with IBM Z OMEGAMON Web UI dashboards:
                <Space layout={'inline'} h={2} />
                <InlineSwitch transparent={true} defaultChecked={falcon.enabled} 
                  onChange={e => {
                    falcon.enabled = e.currentTarget.checked;
                    this.updateFolderState();
                  }} 
                />
              </p>
              {falcon.enabled && (
                <>
                  <p>
                    <InlineField label="Address Space Details Dashboard:" labelWidth={30}>
                      <Input type="url" width={61} defaultValue={falcon.asDashboard} 
                        onChange={e => {
                          falcon.asDashboard = e.target.value;
                        }}
                      />
                    </InlineField>
                  </p>
                  <p>
                    <InlineField label="LPAR Details Dashboard:" labelWidth={30}>
                      <Input type="url" width={61} defaultValue={falcon.sysDashboard} 
                        onChange={e => {
                          falcon.sysDashboard = e.target.value;
                        }}
                      />
                    </InlineField>
                  </p>
                  <p>
                    <Space layout={'inline'} h={1} />
                    Note: You must associate RMF data source with OMEGAMON data source (see RMF data source edit page)
                  </p>
                </>
              )}
            </p>
            <Button disabled={isBusy} variant="primary" fill="outline" icon={'plus'} onClick={this.createDataSource}>
              Create Data Source
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy}
              variant="primary"
              fill="outline"
              icon={'apps'}
              onClick={
                dds.installed
                  ? () =>
                      this.goToFolder(
                        DDS_FOLDER_UID,
                        dds.operation.code === OperCode.Install || dds.operation.code === OperCode.Reset
                      )
                  : async () => {
                      await this.processFolder(DDS_FOLDER_UID, OperCode.Install, falcon);
                    }
              }
            >
              {dds.installed ? 'Go to ' : 'Install '}
              DDS Dashboards
              <Space h={1} />
              <StatusIcon code={dds.operation.code === OperCode.Install ? dds.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy || !dds.installed}
              variant="success"
              fill="outline"
              icon={'process'}
              onClick={async () => {
                await this.processFolder(DDS_FOLDER_UID, OperCode.Reset, falcon);
              }}
            >
              Update / Reset
              <Space h={1} />
              <StatusIcon code={dds.operation.code === OperCode.Reset ? dds.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy || !dds.installed}
              variant="destructive"
              fill="outline"
              icon={'trash-alt'}
              onClick={async () => {
                await this.processFolder(DDS_FOLDER_UID, OperCode.Delete, falcon);
              }}
            >
              Delete
              <Space h={1} />
              <StatusIcon code={dds.operation.code === OperCode.Delete ? dds.operation.status : OperStatus.None} />
            </Button>
          </PanelContainer>
        </Box>
        <Space layout={'block'} v={3} />
        <Box backgroundColor={'secondary'}>
          <PanelContainer style={{ backgroundColor: 'rgba(255, 255, 255, 0)', width: '100%', padding: '2rem' }}>
            <h4>
              <Icon style={{ color: '#e6522c' }} size={'xxxl'} name={'gf-prometheus'} />
              <Space layout={'inline'} h={2} /> Prometheus Sample Dashboards
            </h4>
            <br />
            <p>
              RMF Distributed Data Server (DDS) exposes Monitor III metrics in OpenMetrics format which allows to feed
              the data into 3rd party monitoring systems such as Prometheus, VictoriaMetrics, and others.
            </p>
            <p>
              Follow{' '}
              <TextLink href={DDS_OPEN_METRICS_DOC_URL} color="link" inline={true} external>
                RMF DDS documentation
              </TextLink>{' '}
              to set it up and install the Prometheus-based sample dashboards to explore.
            </p>
            <p>
              Destination folder: <i>Dashboards / {prom.folderPath}</i>
              <Space layout={'inline'} h={2} />
              <span style={{ color: 'gray' }}>[UID=&apos;{PROM_FOLDER_UID}&apos;]</span>
            </p>
            <Button
              disabled={isBusy}
              variant="primary"
              fill="outline"
              icon={'apps'}
              onClick={
                prom.installed
                  ? () =>
                      this.goToFolder(
                        PROM_FOLDER_UID,
                        prom.operation.code === OperCode.Install || prom.operation.code === OperCode.Reset
                      )
                  : async () => {
                      await this.processFolder(PROM_FOLDER_UID, OperCode.Install, falcon);
                    }
              }
            >
              {prom.installed ? 'Go to ' : 'Install '}
              Prometheus Dashboards
              <Space h={1} />
              <StatusIcon code={prom.operation.code === OperCode.Install ? prom.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy || !prom.installed}
              variant="success"
              fill="outline"
              icon={'process'}
              onClick={async () => {
                await this.processFolder(PROM_FOLDER_UID, OperCode.Reset, falcon);
              }}
            >
              Update / Reset
              <Space h={1} />
              <StatusIcon code={prom.operation.code === OperCode.Reset ? prom.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy || !prom.installed}
              variant="destructive"
              fill="outline"
              icon={'trash-alt'}
              onClick={async () => {
                await this.processFolder(PROM_FOLDER_UID, OperCode.Delete, falcon);
              }}
            >
              Delete
              <Space h={1} />
              <StatusIcon code={prom.operation.code === OperCode.Delete ? prom.operation.status : OperStatus.None} />
            </Button>
          </PanelContainer>
        </Box>
        <Space layout={'block'} v={3} />
        <Box backgroundColor={'secondary'}>
          <PanelContainer style={{ backgroundColor: 'rgba(255, 255, 255, 0)', width: '100%', padding: '2rem' }}>
            <h4>
              <img style={{ width: '48px', height: '48px' }} src={PM_LOGO_URL} alt="logo for IBM RMF" />
              <Space layout={'inline'} h={2} /> Import Dashboards from RMF Performance Monitoring File
            </h4>
            <br />
            <p>
              You can import dashboards .po files. Drag and drop the files onto the drop zone below.
            </p>
            <p>
              Destination folder: <i>Dashboards / {pm.folderPath}</i>
              <Space layout={'inline'} h={2} />
              <span style={{ color: 'gray' }}>[UID=&apos;{PM_FOLDER_UID}&apos;]</span>
            </p>
            <Button
              disabled={isBusy || !pm.installed}
              variant="primary"
              fill="outline"
              icon={'apps'}
              onClick={
                () => {
                      getDataSourceSrv().reload();
                      this.goToFolder(
                        PM_FOLDER_UID,
                        pm.operation.code === OperCode.Install || pm.operation.code === OperCode.Reset
                      );
                }
              }
            >
              Go to PM Dashboards
              <Space h={1} />
              <StatusIcon code={pm.operation.code === OperCode.Install ? pm.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Button
              disabled={isBusy || !pm.installed}
              variant="destructive"
              fill="outline"
              icon={'trash-alt'}
              onClick={async () => {
                await this.processFolder(PM_FOLDER_UID, OperCode.Delete, falcon);
              }}
            >
              Delete
              <Space h={1} />
              <StatusIcon code={pm.operation.code === OperCode.Delete ? pm.operation.status : OperStatus.None} />
            </Button>
            <Space layout={'inline'} h={2} />
            <Space layout={'block'} v={2} />
            <FileDropzone options={{accept:{"application/octet-stream": [".po"]}}} onLoad={async (result) => {
                const nameUid = await this.prepareDatasources(result);
                const dashboard = parsePmImportFileToDashboard(result, nameUid);
                await this.importDashboard(PM_FOLDER_UID, OperCode.Install, [dashboard]);
            }} fileListRenderer={(file: DropzoneFile, removeFile: (file: DropzoneFile) => void) => {
              return null;
            }} />
            <Space layout={'block'} v={2} />
            <p>
              <Space layout={'inline'} h={1} />
              Note: Ensure imported datasources have DDS Server credentials if required and fully qualified domain names or accessible IP addresses before using the dashboards.
            </p>
          </PanelContainer>
        </Box>
      </div>
    );
  }
}

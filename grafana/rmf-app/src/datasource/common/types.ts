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
import {
  DataQuery,
  DataSourceJsonData,
  DataSourceSettings,
  NullValueMode,
  SelectableValue,
  ThresholdsMode,
} from '@grafana/data';

export enum visualisationType {
  auto = 'auto',
  bargauge = 'bargauge',
  table = 'table',
  gauge = 'gauge',
  timeseries = 'TimeSeries',
}

export interface RMFQuery extends DataQuery {
  constant?: number;
  selectedResource?: SelectableValue<any>;
  selectedMetrics?: Array<SelectableValue<any>>;
  loadDataSource?: string;
  treeviewResourcesData?: TreeNode[];
  metricsDataBackup?: TreeNode[];
  reportDataBackup?: TreeNode[];
  reportOptionsData?: SelectItem[];
  selectedVisualisationType?: visualisationType;
  onLeaveClick: ((this: PermissionStatus, ev: Event) => any) | null;
  onExpandCollapsClick: ((this: PermissionStatus, data: TreeNode, ev: Event) => any) | null;
  id?: string;
  selectedType?: string;
  selectDropDownBoxValue?: any;
  autocompleteTextFieldOptions?: any;
  queryText?: string;
  selectedQuery?: string;
  editorSelectedResource?: string;
  enableTimeSeries?: boolean;
  serviceCallInprogres?: boolean;
  errorMessage?: string;
  rmfPanelGuid?: string;
  refreshRequired?: boolean;
  absoluteTimeSelected?: boolean;
  dashboardUid?: string;
}

export interface MetricItem {
  value: string;
  name: string;
}

export interface SelectItem {
  label: string;
  value: any;
  description?: string;
  imgUrl?: string;
}

export interface MyTreeViewData extends DataQuery {
  data?: TreeNode[];
  onLeaveClick?: ((ev: Event) => any) | null;
  onExpandCollapsClick?: ((data: SelectedTreeNode, flag: boolean) => any) | null;
  id?: string;
  displayCheckbox?: boolean;
}

export interface SelectedTreeNode {
  name: string;
  value: string;
  visualisationType: visualisationType;
  chkBoxSelected: boolean;
  desc?: string;
}

export interface TreeNode {
  name: string;
  value: string;
  index: number;
  visualisationType?: visualisationType;
  attributes: boolean;
  expandable: boolean;
  expanded: boolean;
  icon: string;
  restype: string;
  format: string;
  unit: string;
  displayChkBox: boolean;
  chkBoxSelected: boolean;
  children: TreeNode[];
}

export const defaultQuery: Partial<RMFQuery> = {
  constant: 6.5,
};

/**
 * These are options configured for each DataSource instance
 */
export interface RMFDataSourceJsonData extends DataSourceJsonData {
  // Conventional Grafana HTTP config (see the `DataSourceHttpSettings` UI element)
  tlsSkipVerify?: boolean;
  // We store it as a string because of complications of UI input validation
  timeout?: string;
  // Custom config RMF options
  cacheSize?: string;
  disableCompression?: boolean;
  // Legacy: custom RMF settings. We should ge rid of it at some point.
  path?: string;
  port?: string;
  ssl?: boolean;
  userName?: string;
  // NB: the meaning of that is inverted. If set, we do verify certificates.
  skipVerify?: boolean;
  omegamonDs?: string;
  batchRequestInterval?: string;
}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface RMFDataSourceSecureJsonData {
  // Conventional Grafana HTTP config (see the `DataSourceHttpSettings` UI element)
  basicAuthPassword?: string;
  // Legacy: custom RMF settings. We should ge rid of it at some point.
  userName?: string;
  password?: string;
}

export type RMFDataSourceSettings = DataSourceSettings<RMFDataSourceJsonData, RMFDataSourceSecureJsonData>;

export interface DashboardHistoryItem {
  dashboardId?: number;
  panelId?: number;
  queryItems: QueryItem[];
  inprogress: boolean;
}

export interface QueryItem {
  refId: string;
  startTime: string;
  endTime: string;
  nextTimeInterval: string;
  IntervalSec: number;
  timeList: number[];
  dataList: any[];
  sentItem: any;
  resource: string;
  workscope: string;
  filter: string;
  id: string;
}

export const defaultTableConfig = {
  decimals: 0,
  custom: {
    align: 'left',
    width: 80,
  },
  nullValueMode: NullValueMode.Ignore,
  thresholds: {
    mode: ThresholdsMode.Percentage,
    steps: [
      { value: -Infinity, color: 'green', state: 'HighHigh' },
      { value: 70, color: 'orange', state: 'OK' },
      { value: 90, color: 'red', state: 'High' },
    ],
  },
};

export interface PropTypes {
  Component?: any;
  defaultValue: string;
  disabled?: boolean;
  maxOptions?: number;
  onBlur?: any;
  onChange?: any;
  onKeyDown?: any;
  handleOnRequestOptions?: any;
  handleOnDropDownItemSelect?: any;
  handleOnKeyDown?: any;
  handleOnTextChange?: any;
  changeOnSelect?: any;
  options?: any;
  regex?: string;
  matchAny?: boolean;
  minchars: number;
  requestonlyifnooptions?: boolean;
  spaceremovers?: string[];
  spacer?: string;
  trigger?: any;
  value?: string;
  offsetx?: number;
  offsety?: number;
  passthroughenter?: boolean;
  caret?: any;
  helpervisible?: boolean;
  selection?: any;
  matchstart?: number;
  matchlength?: number;
  top?: number;
  left?: number;
}

//  trigger: ['@', '='],
export const autoCompleteDefaultProps: PropTypes = {
  Component: 'textarea',
  defaultValue: '',
  disabled: false,
  maxOptions: 0,
  regex: '^[#()/$%A-Za-z0-9\\-_ ]+$',
  matchAny: true,
  minchars: 0,
  requestonlyifnooptions: true,
  spaceremovers: [',', '.', '!', '?'],
  spacer: ' ',
  trigger: ['.'],
  offsetx: 0,
  offsety: 0,
  value: undefined,
  passthroughenter: false,
  caret: {},
  helpervisible: false,
  selection: 0,
  matchstart: 0,
  matchlength: 0,
  top: 0,
  left: 0,
};

export const lhsSign = '.';
export const rhsSign = '=';

export const resourceBaseData = {
  SYSPLEX: {
    RESOURCE_PATH: ',<0>,SYSPLEX',
    MVS_IMAGE: {
      RESOURCE_PATH: ',<0>,MVS_IMAGE',
      'I/O_SUBSYSTEM': {
        RESOURCE_ALL_PATH: '<0>,*,I/O_SUBSYSTEM',
        SSID: {
          RESOURCE_ALL_PATH: '<0>,*,ALL_SSIDS',
          RESOURCE_PATH: '<0>,<1>,SSID',
        },
        LOGICAL_CONTROL_UNIT: {
          RESOURCE_ALL_PATH: '<0>,*,ALL_LCUS',
          RESOURCE_PATH: '<0>,<1>,LOGICAL_CONTROL_UNIT',
        },
        CHANNEL_PATH: {
          RESOURCE_ALL_PATH: '<0>,*,ALL_CHANNELS',
          RESOURCE_PATH: '<0>,<1>,CHANNEL_PATH',
        },
        VOLUME: {
          RESOURCE_ALL_PATH: '<0>,*,ALL_VOLUMES',
          RESOURCE_PATH: '<0>,<1>,VOLUME',
        },
        CRYPTO: {
          RESOURCE_ALL_PATH: '<0>,*,CRYPTO',
          RESOURCE_PATH: '<0>,<1>,CRYPTO_CARD',
        },
        PCIE: {
          RESOURCE_ALL_PATH: '<0>,*,PCIE',
          RESOURCE_PATH: '<0>,*,PCIE_FUNCTION',
        },
        SCM: {
          RESOURCE_ALL_PATH: '<0>,*,SCM',
          RESOURCE_PATH: '<0>,<1>,SCM_CARD',
        },
        ZFS: {
          RESOURCE_ALL_PATH: '<0>,*,ZFS',
          RESOURCE_PATH: '<0>,<1>,AGGREGATE',
          FILESYSTEM: {
            RESOURCE_PATH: '<0>,<1>,AGGREGATE',
          },
        },
      },
      PROCESSOR: {
        RESOURCE_ALL_PATH: '<0>,*,PROCESSOR',
      },
      STORAGE: {
        RESOURCE_ALL_PATH: '<0>,*,STORAGE',
        AUXILIARY_STORAGE: {
          RESOURCE_ALL_PATH: '<0>,*,AUXILIARY_STORAGE',
        },
        CENTRAL_STORAGE: {
          RESOURCE_ALL_PATH: '<0>,*,CENTRAL_STORAGE',
          CSA: {
            RESOURCE_ALL_PATH: '<0>,*,CSA',
          },
          SQA: {
            RESOURCE_ALL_PATH: '<0>,*,SQA',
          },
          ECSA: {
            RESOURCE_ALL_PATH: '<0>,*,ECSA',
          },
          ESQA: {
            RESOURCE_ALL_PATH: '<0>,*,ESQA',
          },
        },
      },
      ENQUEUE: {
        RESOURCE_ALL_PATH: '<0>,*,ENQUEUE',
      },
      OPERATOR: {
        RESOURCE_ALL_PATH: '<0>,*,OPERATOR',
      },
      SW_SUBSYSTEMS: {
        RESOURCE_ALL_PATH: '<0>,*,SW_SUBSYSTEMS',
        JES: {
          RESOURCE_ALL_PATH: '<0>,*,JES',
        },
        XCF: {
          RESOURCE_ALL_PATH: '<0>,*,XCF',
        },
        HSM: {
          RESOURCE_ALL_PATH: '<0>,*,HSM',
        },
      },
    },
    CPC: {
      RESOURCE_PATH: ',<0>,CPC',
      LPAR: {
        RESOURCE_PATH: '<0>,<1>,LPAR',
      },
    },
    COUPLING_FACILITY: {
      RESOURCE_PATH: ',<0>,COUPLING_FACILITY',
      CF_STRUCTURE: {
        RESOURCE_PATH: '<0>,<1>,CF_STRUCTURE',
      },
    },
    WLM_ACTIVE_POLICY: {
      RESOURCE_PATH: ',<0>,WLM_ACTIVE_POLICY',
      WLM_WORKLOAD: {
        RESOURCE_ALL_PATH: '<0>,*,ALL_WLM_WORKLOADS',
        RESOURCE_PATH: '<0>,<1>,WLM_WORKLOAD',
        WLM_SERVICE_CLASS: {
          RESOURCE_PATH: '<0>,<1>,WLM_SERVICE_CLASS',
          WLM_SC_PERIOD: {
            RESOURCE_PATH: '<0>,<1>,WLM_SC_PERIOD',
          },
        },
      },
      WLM_REPORT_CLASS: {
        RESOURCE_ALL_PATH: '<0>,*,ALL_WLM_REPORT_CLASSES',
        RESOURCE_PATH: '<0>,<1>,WLM_REPORT_CLASS',
        WLM_RC_PERIOD: {
          RESOURCE_PATH: '<0>,<1>,WLM_RC_PERIOD',
        },
      },
      WLM_RESOURCE_GROUP: {
        RESOURCE_ALL_PATH: '<0>,*,ALL_WLM_RESOURCE_GROUPS',
        RESOURCE_PATH: '<0>,<1>,WLM_RESOURCE_GROUP',
      },
    },
  },
};

export interface QueryValidationResult {
  result: boolean;
  resourceCommand: string;
  filterTypes: string;
  columnName: string;
  errorMessage: string;
}

export interface FieldValidationResilt {
  result: boolean;
  strResult: string;
}

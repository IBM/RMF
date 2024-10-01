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
// import defaults from 'lodash/defaults';

import { QueryEditorProps } from '@grafana/data';
import { getBackendSrv, getTemplateSrv } from '@grafana/runtime';
import { RadioButtonGroup, Spinner, Switch } from '@grafana/ui';
import React, { PureComponent } from 'react';
import { AutocompleteTextField } from '../autocomplete-text/autocomplete-textfield';
import {
  generateUUID,
  getEnableTimeSeriesStatus,
  getResource,
  getSelectedGuid,
  getSelectedResource,
  getVisualisationType,
  isItFirstDotPosition,
  loadDataToDictionary,
} from '../common/common.helper';
import {
  autoCompleteDefaultProps,
  lhsSign,
  MetricItem,
  PropTypes,
  rhsSign,
  RMFDataSourceJsonData,
  RMFQuery,
} from '../common/types';
import { DataSource } from '../datasource';

import { Parser } from '../parser/core/parser';
import { GrammarResult } from '../parser/core/type';

require('./queryeditor.parser.component.css');

type Props = QueryEditorProps<DataSource, RMFQuery, RMFDataSourceJsonData>;
type state = RMFQuery;

let metricDict: any = {};
let rmfPanelGuid = '';

export class QueryEditorAutoCompleteComponent extends PureComponent<Props, state> {
  editorSelectedResource = '';
  enableTimeSeries = false;

  rbArgs = {
    disabled: false,
    disabledOptions: [''],
    size: 'md',
    fullWidth: true,
  };

  // queryEditor inputbox props
  autoComplDefProps: PropTypes;
  refInput: any;

  constructor(props: QueryEditorProps<DataSource, RMFQuery, RMFDataSourceJsonData>) {
    super(props);
    this.state = {
      onLeaveClick: null,
      onExpandCollapsClick: null,
      refId: '',
    };

    // Load all initial default value
    this.editorSelectedResource = getSelectedResource(props);
    rmfPanelGuid = getSelectedGuid(props);
    if (rmfPanelGuid === '') {
      rmfPanelGuid = generateUUID();
    }
    this.enableTimeSeries = getEnableTimeSeriesStatus(props);

    this.autoComplDefProps = { ...autoCompleteDefaultProps };
    this.autoComplDefProps.value = props.query?.selectedQuery ? props.query.selectedQuery : '';
    this.refInput = React.createRef();
    if (Object.keys(metricDict).length === 0) {
      this.loadDataFromService(props.datasource.id)
        .then((resp1: any) => {})
        .catch((err: any) => {
          this.setServiceCallInprogresState(false);
        });
    }
  }

  componentDidMount() {
    this.setState({ selectedType: 'report' });
    this.setState({ autocompleteTextFieldOptions: { [lhsSign]: [], [rhsSign]: [] } });
    this.setState({ queryText: '' });
    this.setState({ editorSelectedResource: this.editorSelectedResource });
    this.setState({ enableTimeSeries: this.enableTimeSeries });
    this.setState({ rmfPanelGuid: rmfPanelGuid });
  }

  async loadDataFromService(id: number) {
    return await new Promise((resolve, reject) =>
      getBackendSrv()
        .fetch({
          method: 'post',
          headers: {
            Accept: 'application/text',
            'Content-Type': 'application/text',
          },
          url: `/api/datasources/${id}/resources/autopopulate`,
          responseType: 'text',
        })
        .subscribe(
          (resp) => {
            if (resp.data) {
              let result = JSON.parse(resp.data as string);
              let resourceList: any[] = getResource(result);
              metricDict = loadDataToDictionary(resourceList);
              resolve(true);
            } else {
              reject(false);
            }
          },
          (err) => {
            this.setServiceCallInprogresState(false);
            reject(false);
          }
        )
    );
  }

  onTextChange = (val: string, e: any) => {
    this.autoComplDefProps.value = val !== undefined ? val : '';
  };

  onKeyDown = (e: any) => {
    if (e.code === 'Period') {
      // Predefine and test to load popup list
      let val = e.currentTarget.defaultValue.substr(0, e.target.selectionStart) + '.';
      if (isItFirstDotPosition(val, e.target.selectionEnd)) {
        let key = val.substring(0, e.target.selectionEnd).toString().toUpperCase();
        this.loadMetricsFromService(key);
      }
    }
  };

  loadMetricsFromService = (key: string): boolean => {
    let isFound: boolean;
    let lhsDropdownList: any[] = this.loadReportAndMetricsList();
    let rhsDropdownList: any[] = [];
    isFound = false;
    if (Object.keys(metricDict).length > 0 && key && key !== '' && metricDict[key] !== undefined) {
      metricDict[key].map((item: MetricItem) => {
        isFound = true;
        lhsDropdownList.push(item.name);
      });
    }

    if (rhsDropdownList.length === 0) {
      rhsDropdownList.push('No result found.');
    }
    rhsDropdownList = this.getDashboardVariables(rhsDropdownList);
    this.setState({ autocompleteTextFieldOptions: { [lhsSign]: lhsDropdownList, [rhsSign]: rhsDropdownList } });
    return isFound;
  };

  loadReportAndMetricsList = () => {
    let lhsDropdownList: any[] = ['REPORT'];
    let rhsDropdownList: any[] = [];
    rhsDropdownList = this.getDashboardVariables(rhsDropdownList);
    this.setState({ autocompleteTextFieldOptions: { [lhsSign]: lhsDropdownList, [rhsSign]: rhsDropdownList } });
    return lhsDropdownList;
  };

  getDashboardVariables(rhsDropdownList: any[]): any[] {
    getTemplateSrv()
      .getVariables()
      .filter((item) => item.type === 'query')
      .forEach((item) => {
        rhsDropdownList.push('$' + item.name);
      });
    return rhsDropdownList;
  }

  onAutoCompBlur = (e: any) => {
    const { onChange, query, onRunQuery } = this.props;
    let value =
      e !== null && e !== undefined && e.currentTarget !== undefined && e.currentTarget.value !== undefined
        ? e.currentTarget.value.trim()
        : '';
    this.autoComplDefProps.value =
      e !== null && e !== undefined && e.currentTarget !== undefined && e.currentTarget.value !== undefined
        ? e.currentTarget.value
        : '';
    if (value.length > 0) {
      let parser = new Parser(
        e !== null && e !== undefined && e.currentTarget !== undefined && e.currentTarget.value !== undefined
          ? e.currentTarget.value.toString().trim()
          : ''
      );
      let result: GrammarResult = parser.parse();
      if (result.errorFound) {
        this.setState({ errorMessage: result.errorMessage });
      } else {
        this.setState({ errorMessage: '' });
        let value =
          e !== null && e !== undefined && e.currentTarget !== undefined && e.currentTarget.value !== undefined
            ? e.currentTarget.value
            : '';
        let finalQueryToSave = this.replaceMetricNamewithID(result.query);
        if (value.length > 0) {
          onChange({
            ...query,
            selectedVisualisationType: getVisualisationType(value, this.enableTimeSeries),
            selectedQuery: value,
            selectedResource: { label: finalQueryToSave, value: finalQueryToSave },
            rmfPanelGuid: rmfPanelGuid,
          });
        }
        this.setState({ editorSelectedResource: finalQueryToSave });
        onRunQuery();
      }
    }
  };

  onSwitchChange = (e: any) => {
    const { onChange, query, onRunQuery } = this.props;
    this.enableTimeSeries =
      e !== null && e !== undefined && e.currentTarget !== undefined && e.currentTarget.checked !== undefined
        ? e.currentTarget.checked
        : false;
    onChange({
      ...query,
      selectedVisualisationType: getVisualisationType(
        this.autoComplDefProps?.value ? this.autoComplDefProps?.value : '',
        this.enableTimeSeries
      ),
    });
    this.setState({ enableTimeSeries: this.enableTimeSeries });
    onRunQuery();
  };

  replaceMetricNamewithID = (query: string): string => {
    let finalQuery: string = query;
    let resType = '';
    let metricName = '';
    let metricId = '';
    let queryParts = query.split('&');
    if (queryParts && queryParts.length > 0) {
      queryParts.map((item: string) => {
        if (item.substring(0, 3).toLocaleUpperCase() === 'ID=') {
          let innerQueryParts = item.split('=');
          metricName = innerQueryParts[1];
        }
        if (item.substring(0, 9).toLocaleUpperCase() === 'RESOURCE=') {
          let innerQueryParts = item.split(',');
          resType = innerQueryParts[2].toLocaleUpperCase();
        }
      });
      metricId = this.getMetricIDFromDictionary(resType, metricName);
      if (metricId !== '') {
        finalQuery = finalQuery.replace(metricName, metricId);
      }
    }
    return finalQuery;
  };

  getMetricIDFromDictionary = (resourceTypeKey: string, metricName: string): string => {
    let resultId = '';
    if (Object.keys(metricDict).length > 0 && resourceTypeKey !== '' && metricName.trim() !== '') {
      metricDict[resourceTypeKey].map((item: MetricItem) => {
        if (item.name.trim().toLocaleUpperCase() === metricName.trim().toLocaleUpperCase()) {
          resultId = item.value;
        }
      });
    }
    return resultId;
  };

  setServiceCallInprogresState = (status: boolean) => {
    this.autoComplDefProps.disabled = status;
    this.setState({ serviceCallInprogres: status });
    this.refInput.current.refInput.current.focus();
  };

  // Grfana templete will refresh only if state items are updated
  render() {
    return (
      <div id={'main+' + (this.props.query as any).refId}>
        <div>
          <span>
            <RadioButtonGroup
              disabledOptions={['resource']}
              options={[{ label: 'Resource:', value: 'resource', icon: 'angle-double-down' }]}
              size={this.rbArgs.size as any}
              value={'resource'}
            />
            <Spinner
              className={
                this.state && this.state.serviceCallInprogres ? 'autotext-spinner' : 'autotext-spinner hide-display'
              }
              size={20}
            />
            <label
              className={
                this.state && this.state.serviceCallInprogres ? 'autotext-spinner' : 'autotext-spinner hide-display'
              }
            >
              Loading...
            </label>
            <label
              className={
                this.state && this.state.errorMessage && this.state.errorMessage.trim() !== ''
                  ? 'autotext-error'
                  : 'autotext-error hide-display'
              }
            >
              Error: {this.state.errorMessage}
            </label>
          </span>
          <span>
            <AutocompleteTextField
              ref={this.refInput}
              onBlur={this.onAutoCompBlur}
              options={
                this.state && this.state.autocompleteTextFieldOptions
                  ? this.state?.autocompleteTextFieldOptions
                  : { [lhsSign]: [], [rhsSign]: [] }
              }
              handleOnRequestOptions={() => {}}
              handleOnDropDownItemSelect={() => {}}
              handleOnKeyDown={this.onKeyDown}
              handleOnTextChange={this.onTextChange}
              {...this.autoComplDefProps}
            />
            <span>
              <div className="rmf-form">
                <label className="gf-form-label">Time series</label>
                <div className="css-lohq6k-input-wrapper width-22 filltext">
                  <div className="css-1w5c5dq-input-inputWrapper rmf-switch">
                    <Switch label={'Time Series'} value={this.state.enableTimeSeries} onChange={this.onSwitchChange} />
                  </div>
                </div>
              </div>
              <div className="rmf-form-resource">
                <label className="gf-form-label width-9">Selected Resource</label>
                <div className="css-lohq6k-input-wrapper width-22 filltext">
                  <div className="css-1w5c5dq-input-inputWrapper filltext">
                    <input
                      className="css-nwgrif-input-input filltext"
                      type="text"
                      value={this.state && this.state.editorSelectedResource ? this.state.editorSelectedResource : ''}
                      disabled={true}
                    />
                  </div>
                </div>
              </div>
            </span>
          </span>
        </div>
      </div>
    );
  }
}

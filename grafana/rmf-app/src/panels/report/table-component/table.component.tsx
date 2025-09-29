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
import { Table } from '@grafana/ui';
import React, { useRef } from 'react';
import { BannerComponent } from '../banner-component/banner.component';
import { CaptionsComponent } from '../captions-component/captions.component';
import {
  InitFrameData,
  applySelectedDefaultsAndOverrides,
  getPaginationFlagFromFieldConfig,
} from './table.helper';
import {PanelProps} from "@grafana/data";
require('./table.component.css');

interface Props extends PanelProps<{}> { }

export const TableComponent: React.FC<Props> = ({ options, fieldConfig, data, width, height }) => {
  let frameData = InitFrameData(data);

  const reportData  = applySelectedDefaultsAndOverrides(options, fieldConfig, frameData);
  const tableData = reportData.tableData
  const captionData = reportData.captionFields
  //const finalTableData = applyFieldOverridesForBarGauge(reportData.tableData);
  const enablePagination: boolean = getPaginationFlagFromFieldConfig(fieldConfig);

  // Setting the scroll properties
  let divBannerRef = useRef<HTMLInputElement>(null);
  let divCaptionRef = useRef<HTMLInputElement>(null);
  let actTableHeight = height - (divBannerRef?.current?.offsetHeight ? divBannerRef?.current?.offsetHeight : 38);
  if (captionData && captionData.length > 0 && captionData.length > 6) {
    actTableHeight = actTableHeight - (divCaptionRef?.current?.offsetHeight ? divCaptionRef?.current?.offsetHeight : 110);
  } else if (captionData && captionData.length > 0 && captionData.length > 3) {
    actTableHeight = actTableHeight - (divCaptionRef?.current?.offsetHeight ? divCaptionRef?.current?.offsetHeight : 76);
  } else {
    actTableHeight = actTableHeight - (divCaptionRef?.current?.offsetHeight ? divCaptionRef?.current?.offsetHeight : 38);
  }


  return (
    <div>
      <div className="banner-section ">
        <BannerComponent fields={ reportData.bannerFields } />
      </div>
      {reportData.captionFields.length > 0 ? (
        <div className="caption-section">
          <CaptionsComponent fields={ reportData.captionFields } />
        </div>
      ) : (
        ''
      )}
      <div className="panel-table-container">
        {( tableData && tableData.length > 0 && tableData.fields && tableData.fields.length > 0) ? (
          <Table
            key={'dataTable'}
            data={tableData}
            height={actTableHeight}
            width={width}
            enablePagination={enablePagination}
          />
        ) : (
          <div className="nodata-label">No Data</div>
        )}
      </div>
    </div>
  );
};

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
import React from 'react';
import { Table } from '@grafana/ui';
import { DataFrame, PanelProps } from '@grafana/data';
import { TableBanner } from '../types';
import { CaptionsComponent } from '../captions-component/captions.component';
import { BannerComponent } from '../banner-component/banner.component';
import {
  applyFieldOverridesForBarGauge,
  prepareBannerToDisplay,
  prepareCaptionListToDisplay,
  applySelectedDefaultsAndOverrides,
  InitFrameData,
  getPaginationFlagFromFieldConfig,
} from './table.helper';
require('./table.component.css');

interface Props extends PanelProps<{}> {}

export const TableComponent: React.FC<Props> = ({ options, fieldConfig, data, width, height }) => {
  let bannerItems: TableBanner;
  let captionList: any[] = [];
  let frameData = InitFrameData(data);

  bannerItems = prepareBannerToDisplay(data);
  captionList = prepareCaptionListToDisplay(frameData);
  const finalData: DataFrame[] = applySelectedDefaultsAndOverrides(options, fieldConfig, frameData);
  const finalTableData = applyFieldOverridesForBarGauge(finalData);
  const enablePagination: boolean = getPaginationFlagFromFieldConfig(fieldConfig);

  // TODO: need to finalize later
  let actTableHeight = (height * 95) / 100;
  return (
    <div>
      <div className="banner-section ">
        <BannerComponent dataList={bannerItems} />
      </div>
      {captionList.length > 0 ? (
        <div className="caption-section">
          <CaptionsComponent dataList={captionList} />
        </div>
      ) : (
        ''
      )}
      <div className="panel-table-container">
        {(finalData && finalData.length > 0 && finalData[0].fields.length) > 0 ? (
          <Table
            key={'dataTable'}
            data={finalTableData[0]}
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

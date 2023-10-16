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
import React from 'react';
import { BannerData } from '../types';
require('./banner.component.css');

export const BannerComponent: React.FC<BannerData> = ({ dataList }) => {
  return (
    <span>
      <span>
        <span className="banner-item">Samples</span>:&nbsp;
        {dataList !== undefined && dataList.samples !== undefined ? dataList.samples : ''}
      </span>
      {dataList !== undefined && dataList.system !== undefined && dataList.system !== '' ? (
        <span>
          <span className="banner-item">&nbsp;&nbsp;&nbsp;&nbsp;System</span>:&nbsp;{dataList.system}
        </span>
      ) : (
        ''
      )}
      <span>
        <span className="banner-item">&nbsp;&nbsp;&nbsp;&nbsp;Time Range</span>:&nbsp;
        {dataList !== undefined && dataList.timeRange !== undefined ? dataList.timeRange : ''}
      </span>
      {dataList !== undefined && dataList.range !== undefined && dataList.range !== '' ? (
        <span>
          <span className="banner-item">&nbsp;&nbsp;&nbsp;&nbsp;Range</span>:&nbsp;{dataList.range}&nbsp;Seconds
        </span>
      ) : (
        ''
      )}
    </span>
  );
};

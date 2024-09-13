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
import { BANNER_PREFIX, FieldProps } from '../types';
require('./banner.component.css');


export const BannerComponent: React.FC<FieldProps> = ({ fields }) => {
  return (
    <span>
      { fields.map((field) => {
          const name = field.name.slice(BANNER_PREFIX.length)
          const values = field.values.buffer || field.values || [];
          return (
            <span key={ field.name }>
              <span className="banner-item">{ name }</span>:&nbsp;{ values[0] ?? 'N/A' }&nbsp;&nbsp;&nbsp;&nbsp;
            </span>
          );
        })
      }
    </span>
  )
};

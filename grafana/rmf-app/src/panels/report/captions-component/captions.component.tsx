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
import { CAPTION_PREFIX, FieldProps } from '../types';
require('./captions.component.css');


export const CaptionsComponent: React.FC<FieldProps> = ({ fields }) => {
  return (
    <span>
      {
        fields.map((field: any) => {
          const name = field.name.slice(CAPTION_PREFIX.length)
          const values = field.values.buffer || field.values || [];
          return (
            <div key={ field.name } className="captionitem">
              <span className="captionitemname">{ name }</span> : { values[0] ?? 'N/A'}
            </div>
          );
        })
      }
    </span>
  );
};

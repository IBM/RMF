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
import React from 'react';
import { APP_NAME, APP_DESC, APP_LOGO_URL } from '../../constants';

interface Props {}

// It's an ugly replacement of the implementation based on @grafana/scenes
// for the sake of compatibility with Grafana v10+
export const Header: React.FC<Props> = ({}) => {
  return (
    <>
      <div style={{ display: 'flex', alignItems: 'row' }}>
        <img
          src={APP_LOGO_URL}
          style={{ width: '32px', height: '32px', marginRight: '16px' }}
          alt={`logo for ${APP_NAME}`}
        />
        <h1 style={{ display: 'inline-block' }}>{APP_NAME}</h1>
      </div>
      <div style={{ position: 'relative', color: 'gray' }}>{APP_DESC}</div>
    </>
  );
};

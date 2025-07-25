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

interface SpaceProps {
  layout?: string;
  h?: number;
  v?: number;
}

// `Space` from @grafana/ui is available only in Grafana v11+
export const Space: React.FC<SpaceProps> = ({ layout = 'block', h = 0, v = 0 }) => {
  let style = {
    display: layout == 'inline' ? 'inline-block' : 'block',
    paddingRight: `${h * 8}px`,
    paddingBottom: `${v * 8}px`,
  };
  return <div style={style} />;
};

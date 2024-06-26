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
import React, { FC } from 'react';
import { SVGProps } from '../types';

/**
 * Multi Layer Security
 */
export const MultiLayerSecurity: FC<SVGProps> = ({ size, fill, ...rest }) => {
  return (
    <a target="_blank" rel="noreferrer" href="" title="Multi Layer Security enabled.">
      <svg
        version="1.1"
        id="MultiLayerSecurity"
        x="0px"
        y="0px"
        viewBox="0 0 32 32"
        width={size}
        height={size}
        {...rest}
      >
        <path
          fill={fill ? fill : '#DC382D'}
          d="M27,11.1L27,11.1L26.9,9c0-4.7-3.8-8.5-8.5-8.5c-0.8,0-1.7,0.1-2.5,0.4C11.5-0.5,6.8,2,5.4,6.5
	C5.2,7.3,5.1,8.2,5.1,9v2.1H0v16.6C0,30.1,1.9,32,4.3,32h23.4c2.4,0,4.3-1.9,4.3-4.3V11.1H27z M18.5,1.5c4.1,0,7.4,3.3,7.5,7.5v2.1
	h-4V9c0-3.1-1.7-5.9-4.4-7.4C17.9,1.6,18.2,1.5,18.5,1.5z M21,11.1H11V9c0-3.2,2-6,5-7c3,1.1,5,3.9,5,7V11.1L21,11.1z M6.1,9
	c0-4.1,3.3-7.4,7.4-7.5c0.3,0,0.6,0,1,0.1C11.7,3.1,10,5.9,10,9v2.1H6L6.1,9L6.1,9z M4.3,31C2.5,31,1,29.5,1,27.7V12.1h4v15.6
	C5,29,5.6,30.2,6.5,31H4.3z M9.3,31C7.4,31,6,29.5,6,27.7V12.1h20v15.6c0,1.8-1.5,3.3-3.3,3.3H9.3z M31,27.7c0,1.8-1.5,3.3-3.3,3.3
	h-2.3c1-0.8,1.6-2,1.6-3.3V12.1h4C31,12.1,31,27.7,31,27.7z M16,15.1c-1.8,0-3.3,1.5-3.3,3.3c0,1.1,0.6,2.2,1.5,2.8v3.4
	c0,1,0.8,1.8,1.8,1.9c1,0,1.8-0.8,1.9-1.8v-0.1v-3.4c1.5-1,2-3,1-4.6C18.2,15.7,17.1,15.1,16,15.1z M17.1,20.4l-0.3,0.1v4
	c0,0.4-0.3,0.8-0.8,0.8c-0.4,0-0.8-0.3-0.8-0.8c0,0,0,0,0-0.1v-4l-0.3-0.1c-1.1-0.6-1.6-2-1-3.1s2-1.6,3.1-1s1.6,2,1,3.1
	C17.8,19.9,17.5,20.2,17.1,20.4L17.1,20.4z"
        />
      </svg>
    </a>
  );
};

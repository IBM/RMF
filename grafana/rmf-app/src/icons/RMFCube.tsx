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
 * RMF cube
 */
export const RMFCube: FC<SVGProps> = ({ size, fill, title, ...rest }) => {
  return (
    <a
      target="_blank"
      rel="noreferrer"
      href=""
      title={title ? title : 'IBM RMF is a in-memory data structure store, used as a database, cache and data broker.'}
    >
      <svg version="1.1" id="RMFCube" x="0px" y="0px" viewBox="0 0 32 32" width={size} height={size} {...rest}>
        <path
          fill={fill ? fill : '#DC382D'}
          d="M31.7,16.3c-0.4-0.1-1.5-0.5-2.7-1c1.3-0.6,2.4-1,2.8-1.3c0.8-0.4,1.6-0.9,1.6-1.6c0-0.7-0.8-1-1.7-1.3
	c-0.6-0.2-1.7-0.6-3.1-1.2c1.5-0.6,2.6-1.2,3.2-1.4c0.8-0.4,1.5-0.9,1.5-1.6c0-0.7-0.8-1-1.6-1.3c-0.8-0.3-3.5-1.3-6.1-2.3
	c-2.6-1-5.3-2.1-6.1-2.4c-1.8-0.6-2.8-0.7-4.8,0.1c-1.9,0.8-11.5,4.5-13,5.1C1.1,6.4,0.3,6.8,0.3,7.5c0,0.7,0.8,1.2,1.6,1.6
	C2.4,9.3,3.6,9.8,5,10.4c-0.6,0.2-1.2,0.5-1.7,0.7c-0.7,0.3-1.2,0.5-1.5,0.6c-0.7,0.3-1.5,0.6-1.5,1.3c0,0.7,0.8,1.2,1.6,1.6
	c0.5,0.2,1.5,0.7,2.8,1.2l-0.3,0.1c-1.2,0.5-2.1,0.8-2.6,1c-0.7,0.3-1.5,0.7-1.5,1.3c0,0.6,0.6,1,1.6,1.6c1.3,0.7,6.6,2.8,10.1,4.3
	c1.2,0.5,2.2,0.9,2.6,1.1c0.8,0.4,1.4,0.5,2.1,0.5c0.8,0,1.7-0.3,2.8-0.9c0.9-0.5,4-1.8,6.8-3c2.5-1.1,4.8-2.1,5.6-2.5
	c0.8-0.4,1.6-0.9,1.6-1.6C33.3,17,32.5,16.6,31.7,16.3z M1.3,7.5C1.3,7.5,1.3,7.5,1.3,7.5c0,0,0.2-0.2,0.9-0.5
	c1.5-0.6,11.1-4.3,13-5.1c1.8-0.8,2.5-0.6,4-0.1c0.8,0.3,3.5,1.3,6.1,2.3c2.6,1,5.3,2.1,6.1,2.4C32.1,6.8,32.3,7,32.3,7
	c0,0.1-0.2,0.3-1,0.7c-0.8,0.4-3.1,1.4-5.6,2.5c-2.9,1.2-5.9,2.5-6.8,3c-1.9,1-2.7,0.9-4,0.3c-0.7-0.3-3.2-1.4-5.9-2.5
	C6.2,9.8,3.2,8.6,2.3,8.1C1.5,7.8,1.3,7.5,1.3,7.5z M1.3,13C1.3,13,1.3,13,1.3,13c0,0,0.2-0.2,0.8-0.5C2.4,12.5,3,12.3,3.7,12
	c0.8-0.3,1.7-0.7,2.7-1l0,0c0.8,0.3,1.6,0.6,2.4,1c2.7,1.1,5.2,2.1,5.9,2.4c0.8,0.4,1.4,0.5,2.1,0.5c0.8,0,1.7-0.3,2.8-0.9
	c0.9-0.5,4-1.8,6.8-3c0.4-0.2,0.8-0.3,1.2-0.5c1.7,0.7,3.4,1.3,4.1,1.5c0.6,0.2,0.9,0.4,0.9,0.4c-0.1,0.1-0.3,0.3-1,0.6
	c-0.8,0.4-3.1,1.4-5.6,2.5c-2.9,1.2-5.9,2.5-6.8,3c-1.9,1-2.7,0.9-4,0.3c-0.7-0.3-3.2-1.4-5.9-2.5c-2.9-1.2-5.9-2.4-6.8-2.8
	C1.5,13.3,1.3,13,1.3,13z M31.3,18.4c-0.8,0.4-3.1,1.4-5.6,2.5c-2.9,1.2-5.9,2.5-6.8,3c-1.9,1-2.7,0.9-4,0.3
	c-0.4-0.2-1.3-0.6-2.6-1.1c-3.3-1.3-8.7-3.6-10-4.2c-0.7-0.3-0.9-0.5-1-0.6c0.1-0.1,0.3-0.2,0.8-0.4c0.5-0.2,1.4-0.5,2.6-1L6,16.3
	c0.9,0.4,1.8,0.7,2.7,1.1c2.6,1.1,5.1,2.1,5.8,2.4c0.8,0.4,1.4,0.5,2.1,0.5c0.8,0,1.7-0.3,2.8-0.9c0.9-0.5,4-1.8,6.8-3
	c0.5-0.2,1-0.4,1.5-0.6c1.6,0.6,3.1,1.2,3.7,1.4c0.6,0.2,0.9,0.4,0.9,0.4C32.2,17.8,32,18,31.3,18.4z"
        />
      </svg>
    </a>
  );
};

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
import { css } from '@emotion/css';
import { config } from '@grafana/runtime';
import { APP_NAME, APP_DESC, APP_LOGO_URL } from '../../constants';

const getStyles = () => {
  const theme = config.theme2;
  return {
    headerContainer: css({
      display: 'flex',
      alignItems: 'center',
    }),
    headerLogo: css({
      width: theme.spacing(4),
      height: theme.spacing(4),
      marginRight: theme.spacing(2),
    }),
    headerTitle: css({
      display: 'inline-block',
    }),
    headerDescription: css({
      position: 'relative' as const,
      color: theme.colors.text.secondary,
    }),
  };
};

interface Props {}

export const Header: React.FC<Props> = ({}) => {
  const styles = getStyles();
  return (
    <>
      <div className={styles.headerContainer}>
        <img src={APP_LOGO_URL} className={styles.headerLogo} alt={`logo for ${APP_NAME}`} />
        <h1 className={styles.headerTitle}>{APP_NAME}</h1>
      </div>
      <div className={styles.headerDescription}>{APP_DESC}</div>
    </>
  );
};

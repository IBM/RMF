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
import { Routes, Route, Navigate } from 'react-router-dom';
import { AppRootProps } from '@grafana/data';
import {
  useSceneApp,
  SceneObjectBase,
  SceneObjectState,
  SceneApp,
  EmbeddedScene,
  SceneFlexLayout,
  SceneFlexItem,
  SceneAppPage,
} from '@grafana/scenes';
import { APP_BASE_URL, APP_NAME, APP_LOGO_URL, APP_DESC } from '../constants';
import { Installer } from './Installer';

export const PluginPropsContext = React.createContext<AppRootProps | null>(null);

class SceneItem extends SceneObjectBase<SceneObjectState> {
  static Component = Installer;
  public constructor(state?: Partial<SceneObjectState>) {
    super({ ...state });
  }
}

const getScene = () => {
  return new EmbeddedScene({
    body: new SceneFlexLayout({
      children: [
        new SceneFlexItem({
          body: new SceneItem(),
        }),
      ],
    }),
  });
};

const RootPage = new SceneAppPage({
  title: APP_NAME,
  titleImg: APP_LOGO_URL,
  subTitle: APP_DESC,
  url: APP_BASE_URL,
  getScene: getScene,
  routePath: '',
});

export function getPluginApp() {
  return new SceneApp({
    pages: [RootPage],
    urlSyncOptions: {
      updateUrlOnInit: true,
      createBrowserHistorySteps: true,
    },
  });
}

export function App(props: AppRootProps) {
  const pluginApp = useSceneApp(getPluginApp);
  return (
    <PluginPropsContext.Provider value={props}>
      <Routes>
        <Route path={''} element={<pluginApp.Component model={pluginApp} />} />
        <Route path={'*'} element={<Navigate to={APP_BASE_URL} replace />} />
      </Routes>
    </PluginPropsContext.Provider>
  );
}

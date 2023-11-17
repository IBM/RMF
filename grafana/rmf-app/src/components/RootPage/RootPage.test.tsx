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
import { shallow } from 'enzyme';
import React from 'react';
import { Observable } from 'rxjs';
import { AppPluginMeta, PluginType } from '@grafana/data';
import { Alert } from '@grafana/ui';
import { ApplicationName, ApplicationSubTitle } from '../../constants';
import { DataSourceList } from '../DataSourceList';
import { RootPage } from './RootPage';

/**
 * Meta
 */
const getMeta = (): AppPluginMeta => ({
  id: '',
  name: '',
  type: PluginType.app,
  module: '',
  baseUrl: '',
  info: {
    author: {} as any,
    description: '',
    logos: {
      large: '',
      small: '',
    },
    links: [],
    screenshots: [],
    updated: '',
    version: '',
  },
});

/**
 * DataSourceMock
 */
const getDataSourceMock = jest.fn().mockImplementation(() => Promise.resolve([]));

/**
 * RMFMock
 */
const rmfMock = {
  query: jest.fn().mockImplementation(
    () =>
      new Observable((subscriber) => {
        subscriber.next({
          data: [
            {
              fields: [
                {
                  values: {
                    toArray() {
                      return ['info', '2', '3'];
                    },
                  },
                },
              ],
              length: 1,
            },
          ],
        });
        subscriber.complete();
      })
  ),
};
/**
 * GetRmfMock
 */
const getRmfMock = jest.fn().mockImplementation(() => Promise.resolve(rmfMock));

/**
 * Mock @grafana/runtime
 */
jest.mock('@grafana/runtime', () => ({
  getBackendSrv: () => ({
    get: getDataSourceMock,
  }),
  getDataSourceSrv: () => ({
    get: getRmfMock,
  }),
}));

/**
 * RootPage
 */
describe('RootPage', () => {
  const meta = getMeta();
  const path = '/app';
  const onNavChangedMock = jest.fn();

  beforeAll(() => {
    Object.defineProperty(window, 'location', {
      value: { reload: jest.fn() },
    });
  });

  beforeEach(() => {
    onNavChangedMock.mockClear();
    getDataSourceMock.mockClear();
    getRmfMock.mockClear();
    rmfMock.query.mockClear();
  });

  /**
   * Mounting
   */
  describe('Mounting', () => {
    it('Should update navigation', () => {
      const wrapper = shallow<RootPage>(
        <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />
      );
      const testedMethod = jest.spyOn(wrapper.instance(), 'updateNav');
      wrapper.instance().componentDidMount();
      expect(testedMethod).toHaveBeenCalledTimes(1);
    });

    it('Should make get /api/datasources request', () => {
      const wrapper = shallow<RootPage>(
        <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />
      );
      wrapper.instance().componentDidMount();
      expect(getDataSourceMock).toHaveBeenCalledWith('/api/datasources');
    });

    // it('Should check supported commands', (done) => {
    //   getDataSourceMock.mockImplementationOnce(() =>
    //     Promise.resolve([
    //       {
    //         type: DataSourceType.RMFTYPE,
    //         name: 'IBM RMF for z/OS',
    //       },
    //     ])
    //   );
    //   const wrapper = shallow<RootPage>(
    //     <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />
    //   );
    //   wrapper.instance().componentDidMount();

    //   setImmediate(() => {
    //     expect(getRmfMock).toHaveBeenCalledWith('@grafana/runtime');
    //     expect(rmfMock.query).toHaveBeenCalledWith({ targets: [{ refId: 'A', query: RmfCommand.COMMAND }] });
    //     expect(wrapper.state().loading).toBeFalsy();
    //     expect(wrapper.state().dataSources).toEqual([
    //       {
    //         type: DataSourceType.RMFTYPE,
    //         name: 'IBM RMF for z/OS',
    //         commands: ['INFO'],
    //       },
    //     ]);
    //     done();
    //   });
    // });
  });

  /**
   * updateNav
   */
  describe('updateNav', () => {
    it('Should call onNavChanged prop', () => {
      const wrapper = shallow<RootPage>(
        <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />
      );
      wrapper.instance().updateNav();
      const node = {
        text: ApplicationName,
        img: meta.info.logos.large,
        subTitle: ApplicationSubTitle,
        url: path,
        children: [
          {
            text: 'Home',
            url: path,
            id: 'home',
            icon: 'home',
            active: true,
          },
        ],
      };
      expect(onNavChangedMock).toHaveBeenCalledWith({
        node: node,
        main: node,
      });
    });
  });

  /**
   * Rendering
   */
  describe('rendering', () => {
    it('Should show message if loading=true', (done) => {
      const wrapper = shallow<RootPage>(
        <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />
      );
      const loadingMessageComponent = wrapper.findWhere(
        (node) => node.is(Alert) && node.prop('title') === 'Loading...'
      );
      expect(loadingMessageComponent.exists()).toBeTruthy();
      wrapper.instance().componentDidMount();
      setImmediate(() => {
        const dataSourceListComponent = wrapper.findWhere((node) => node.is(DataSourceList));
        const loadingMessageComponent = wrapper.findWhere(
          (node) => node.is(Alert) && node.prop('title') === 'Loading...'
        );
        expect(loadingMessageComponent.exists()).not.toBeTruthy();
        expect(dataSourceListComponent.exists()).toBeTruthy();
        expect(dataSourceListComponent.prop('dataSources')).toEqual(wrapper.state().dataSources);
        done();
      });
    });

    it('If dataSource is unable to make query, should work correctly', async () => {
      const wrapper = shallow<RootPage>(
        <RootPage basename="" meta={meta} path={path} query={null as any} onNavChanged={onNavChangedMock} />,
        { disableLifecycleMethods: true }
      );

      await wrapper.instance().componentDidMount();

      const dataSourceListComponent = wrapper.findWhere((node) => node.is(DataSourceList));
      const loadingMessageComponent = wrapper.findWhere(
        (node) => node.is(Alert) && node.prop('title') === 'Loading...'
      );
      expect(loadingMessageComponent.exists()).not.toBeTruthy();
      expect(dataSourceListComponent.exists()).toBeTruthy();
    });
  });

  afterAll(() => {
    jest.resetAllMocks();
  });
});

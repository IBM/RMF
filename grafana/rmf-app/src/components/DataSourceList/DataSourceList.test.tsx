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
import { shallow, ShallowWrapper } from 'enzyme';
import { HighAvailability, MultiLayerSecurity, RMFCube } from 'icons';
import React from 'react';
import { Alert } from '@grafana/ui';
import { DataSourceType } from '../../constants';
import { DataSourceList } from './DataSourceList';

type ShallowComponent = ShallowWrapper<typeof DataSourceList>;

const backendSrvMock = {
  post: jest.fn(),
};

jest.mock('@grafana/runtime', () => ({
  getBackendSrv: () => backendSrvMock,
  locationService: {
    push: () => jest.fn(),
  },
}));

/**
 * DataSourceList
 */
describe('DataSourceList', () => {
  const FILLS = {
    success: '#DC382D',
    error: '#A7A7A7',
  };
  const TITLES = {
    success: 'Working as expected',
    error: `Can't retrieve a list of commands. Check that user has permissions to see a list of all commands.`,
  };

  beforeEach(() => {
    Object.values(backendSrvMock).forEach((mock) => mock.mockClear());
  });

  it('If datasources.length=0 should show no items message', () => {
    const wrapper = shallow(<DataSourceList dataSources={[]} />);
    const testedComponent = wrapper.findWhere((node) => node.is(Alert));
    expect(testedComponent.exists()).toBeTruthy();
  });

  /**
   * Item
   */
  describe('Item', () => {
    const getItem = (wrapper: ShallowComponent): ShallowWrapper =>
      wrapper.findWhere((node) => node.hasClass('card-item-wrapper')).first();

    /**
     * RMFCube
     */
    describe('RMFCube', () => {
      it('Should render', () => {
        const dataSources = [
          {
            commands: [''],
            jsonData: {},
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(RMFCube));
        expect(testedComponent.prop('fill')).toEqual(FILLS.success);
        expect(testedComponent.prop('title')).toEqual(TITLES.success);
      });
    });

    /**
     * Name
     */
    describe('Name', () => {
      it('Should render name', () => {
        const dataSources = [
          {
            commands: [],
            jsonData: {},
            name: 'hello',
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.hasClass('card-item-name'));
        expect(testedComponent.text()).toEqual(dataSources[0].name);
      });
    });

    /**
     * Url
     */
    describe('Url', () => {
      it('Should render url', () => {
        const dataSources = [
          {
            commands: [],
            jsonData: {},
            url: 'hello',
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.hasClass('card-item-sub-name'));
        expect(testedComponent.text()).toEqual(dataSources[0].url);
      });
    });

    /**
     * Title
     */
    describe('Title', () => {
      it('If there are not any commands should show title', () => {
        const dataSources = [
          {
            jsonData: {},
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.hasClass('card-item-type'));
        expect(testedComponent.exists()).toBeFalsy();
      });

      it('If there are some commands should hide title', () => {
        const dataSources = [
          {
            commands: [''],
            jsonData: {},
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.hasClass('card-item-type'));
        expect(testedComponent.exists()).not.toBeTruthy();
      });
    });

    /**
     * MultiLayerSecurity
     */
    describe('MultiLayerSecurity', () => {
      it('if jsonData.tlsAuth=true should be shown', () => {
        const dataSources = [
          {
            commands: [],
            jsonData: {
              tlsAuth: true,
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(MultiLayerSecurity));
        expect(testedComponent.exists()).toBeFalsy();
      });

      it('if jsonData.acl=true should be shown', () => {
        const dataSources = [
          {
            commands: ['get'],
            jsonData: {
              acl: true,
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(MultiLayerSecurity));
        expect(testedComponent.exists()).toBeFalsy();
      });

      it('if jsonData.acl=false and jsonData.tlsAuth=false should not be shown', () => {
        const dataSources = [
          {
            commands: ['get'],
            jsonData: {
              tlsAuth: false,
              acl: false,
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(MultiLayerSecurity));
        expect(testedComponent.exists()).not.toBeTruthy();
      });
    });

    /**
     * HighAvailability
     */
    describe('HighAvailability ', () => {
      it('if jsonData.client matches with "cluster" should be shown', () => {
        const dataSources = [
          {
            commands: [],
            jsonData: {
              client: 'cluster',
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(HighAvailability));
        expect(testedComponent.exists()).toBeTruthy();
        expect(testedComponent.prop('fill')).toEqual(FILLS.success);
      });

      it('if jsonData.client matches with "sentinel" should be shown', () => {
        const dataSources = [
          {
            commands: ['get'],
            jsonData: {
              client: 'sentinel',
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(HighAvailability));
        expect(testedComponent.exists()).toBeTruthy();
        expect(testedComponent.prop('fill')).toEqual(FILLS.success);
      });

      it('if jsonData.client does not match with cluster|sentinel should not be shown', () => {
        const dataSources = [
          {
            commands: ['get'],
            jsonData: {
              client: 'manual',
            },
          },
        ];
        const wrapper = shallow<typeof DataSourceList>(<DataSourceList dataSources={dataSources as any} />);
        const item = getItem(wrapper);
        const testedComponent = item.findWhere((node) => node.is(HighAvailability));
        expect(testedComponent.exists()).not.toBeTruthy();
      });
    });
  });

  /**
   * addNewDataSource
   */
  describe('addNewDataSource', () => {
    it('Should add new datasource and redirect on edit page', (done) => {
      const dataSources = [
        {
          id: 1,
          name: 'RMF Data Source',
          commands: [],
          jsonData: {},
        },
      ];
      const wrapper = shallow<ShallowComponent>(<DataSourceList dataSources={dataSources as any} />);
      const addDataSourceButton = wrapper.findWhere(
        (node) => node.name() === 'Button' && node.text() === 'Add RMF Data Source'
      );
      backendSrvMock.post.mockImplementationOnce(() => Promise.resolve({ datasource: { uid: 123 } }));
      addDataSourceButton.simulate('click');
      setImmediate(() => {
        expect(backendSrvMock.post).toHaveBeenCalledWith('/api/datasources', {
          name: 'IBM RMF for z/OS',
          type: DataSourceType.RMFTYPE,
          access: 'proxy',
        });
        done();
      });
    });

    it('Should calc new name', (done) => {
      const dataSources = [
        {
          id: 1,
          name: 'RMF',
          commands: [],
          jsonData: {},
        },
        {
          id: 2,
          name: 'RMF-1',
          commands: [],
          jsonData: {},
        },
      ];
      const wrapper = shallow<ShallowComponent>(<DataSourceList dataSources={dataSources as any} />);
      const addDataSourceButton = wrapper.findWhere(
        (node) => node.name() === 'Button' && node.text() === 'Add RMF Data Source'
      );
      backendSrvMock.post.mockImplementationOnce(() => Promise.resolve({ datasource: { uid: 123 } }));
      addDataSourceButton.simulate('click');
      setImmediate(() => {
        expect(backendSrvMock.post).toHaveBeenCalledWith('/api/datasources', {
          name: 'IBM RMF for z/OS',
          type: DataSourceType.RMFTYPE,
          access: 'proxy',
        });
        done();
      });
    });
  });

  afterAll(() => {
    jest.resetAllMocks();
  });
});

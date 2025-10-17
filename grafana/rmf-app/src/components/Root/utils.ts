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
import { FalconStatus } from './types';

const FOLDERS_API = '/api/folders';
const DASHBOARDS_API = '/api/dashboards/db';
const RMF_DATASOURCE_EXPR = "${datasource}";
const SYSPLEX_NAME = "sysplex";
const LPAR = "LPAR";
const OMEGAMON_DS = "OmegamonDs";
const OMEG_PLEX_LPAR_MVSSYS_EXPR = "${" + SYSPLEX_NAME + "}:${" + LPAR + "}:MVSSYS";
const JOB_NAME_FIELD = "Job Name";
const ASID_FIELD = "ASID (dec)";

// Don't use `getBackendSrv` to avoid built-in notifications: they won't work as we need them.

export async function findFolder(id: string): Promise<string | undefined> {
  const res = await fetch(`${FOLDERS_API}/${id}`);
  if (res.ok) {
    const data = await res.json();
    if (data.parents) {
      const parentPath = data.parents.map((parent: { title: string }) => parent.title).join(' / ');
      return `${parentPath} / ${data.title}`;
    }
    return data.title;
  }
  return undefined;
}

async function createFolder(uid: string, title: string) {
  const res = await fetch(`${FOLDERS_API}`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      uid: uid,
      title: title,
    }),
  });
  if (!res.ok) {
    throw new Error(`failed to create folder: ${res.status} - ${await res.text()}`);
  }
}

export async function deleteFolder(uid: string) {
  const res = await fetch(`${FOLDERS_API}/${uid}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  });
  if (!res.ok) {
    throw new Error(`failed to delete folder: ${res.status} - ${await res.text()}`);
  }
}

export async function findDashboard(query: string, tags : string[]) : Promise<string | undefined>  {
  let url = "/api/search?query=" + query;
  tags.forEach(tag => {
    url += "&tag=" + tag;
  });
  const res = await fetch(url);
  if (res.ok) {
    const data = await res.json();
    if (data.length === 1) {
      return data[0].url;
    }
  }
  return undefined;
}

async function createDashboard(folderUid: string, dashboard: object, falcon: FalconStatus) {
  if (falcon.enabled && falconJobRelated(dashboard)) {
    falconJobUpdate(falcon, dashboard);
  }
  if (falcon.enabled && falconSystemRelated(dashboard)) {
    falconSystemUpdate(falcon, dashboard);
  }
  // Don't use `getBackendSrv` to avoid notifications
  const res = await fetch(DASHBOARDS_API, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ dashboard: dashboard, folderUid: folderUid }),
  });
  if (!res.ok) {
    throw new Error(`failed to create dashboard ${res.status} - ${await res.text()}`);
  }
}

export async function installDashboards(folderUid: string, defaultFolderName: string, dashboards: object[], falcon: FalconStatus) {
  const folderPath = await findFolder(folderUid);
  if (folderPath === undefined) {
    await createFolder(folderUid, defaultFolderName);
  }
  await Promise.all(dashboards.map((dashboard) => createDashboard(folderUid, dashboard, falcon)));
}

export const FALCON_JOB_RELATED = [
  "DELAY - All kinds of delays",
  "OPD - OMVS Process Data",
  "PROC - Processor Delays",
  "PROCU - Processor Usage",
  "STOR - Storage Delay",
  "STORC - Common Storag",
  "STORCR - Common Storage Remaining",
  "STORF - Storage Frames Overview",
  "STORM - Storage Memory Objects",
  "USAGE - Job Oriented Usage",
];

export const FALCON_SYSTEM_RELATED = [
  ...FALCON_JOB_RELATED,
  "CHANNEL - I/O Channels",
  "CPC - Central Processor Complex",
  "DELAY - All kinds of delays",
  "DEV - Device delays",
  "DEVR - Device Resource Delays",
  "DSND - Dataset Delays",
  "EADM - Extended Asynchronous Data Mover",
  "ENCLAVE - Enclaves",
  "ENQ - Enqueue Delays",
  "HSM - Hierarchical Storage Manager Delays",
  "IOQ - I/O Queuing",
  "JES - Job Entry Subsystem Delays",
  "LOCKSP - Spin Lock Report",
  "LOCKSU - Suspend Lock Report",
  "OPD - OMVS Process Data",
  "PCIE - PCIE Activity",
  "PROC - Processor Delays",
  "PROCU - Processor Usage",
  "STOR - Storage Delays",
  "STORC - Common Storage",
  "STORCR - Common Storage Remaining",
  "STORF - Storage Frames Overview",
  "STORM - Storage Memory Objects",
  "STORR - Storage Resource Delays",
  "STORS - Storage by WLM class",
  "SYSINFO - System Info by WLM class",
  "USAGE - Job Oriented Usage",
  "Overall Image Activity",
  "Overall Image Activity (Timeline)",
];

function falconJobRelated(dashboard: any): boolean {
  return FALCON_JOB_RELATED.indexOf(dashboard.title) > -1;
}

function falconSystemRelated(dashboard: any): boolean {
  return FALCON_SYSTEM_RELATED.indexOf(dashboard.title) > -1;
}

function addVars(dashboard: any) {
  var sysplex_name_exist: boolean = false;
  var omegamon_ds_exist: boolean = false;
  dashboard.templating.list.forEach((t: any) => {
    if (t.name === LPAR) {
      t.query = "systems";
      t.definition = t.query;
    }
    if (t.name === SYSPLEX_NAME) {
      sysplex_name_exist = true;
    }
    if (t.name === OMEGAMON_DS) {
      omegamon_ds_exist = true;
    }
  })

  if (!sysplex_name_exist) {
    dashboard.templating.list.push({
      name: SYSPLEX_NAME,
      datasource: {
        type: "ibm-rmf-datasource",
        uid: RMF_DATASOURCE_EXPR
      },
      type: "query",
      query: "sysplex",
      hide: 1,
      includeAll: false,
      multi: false,
      allowCustomValue: false,
    })
  }

  if (!omegamon_ds_exist) {
    dashboard.templating.list.push({
      name: OMEGAMON_DS,
      label: "OMEGAMON Data source",
      description: "OMEGAMON Data source (optional)",
      datasource: {
        type: "ibm-rmf-datasource",
        uid: RMF_DATASOURCE_EXPR
      },
      type: "query",
      query: "OmegamonDs",
      hide: 2,
      includeAll: false,
      multi: false,
      allowCustomValue: false,
    })
  }

}

function falconJobUpdate(falcon: FalconStatus, dashboard: any) {
  addVars(dashboard);

  const JOB_NAME_EXPR = "${__data.fields[\"" + JOB_NAME_FIELD + "\"]}";
  const ASID_EXPR = "${__data.fields[\"" + ASID_FIELD + "\"]}";
  let baseUrl = falcon.asDashboard;
  if (baseUrl.indexOf("?") > -1) {
    baseUrl += "&";
  } else {
    baseUrl += "?";
  }

  dashboard.panels.forEach((p: any) => {
    p.fieldConfig.overrides = [];
    p.fieldConfig.overrides.push({
      matcher: {
        id: "byName",
        options: JOB_NAME_FIELD,
      },
      properties: [
        {
          id: "links",
          value : [
            {
              targetBlank: true,
              title: "OMEGAMON details for " + JOB_NAME_EXPR,
              url: baseUrl + "var-dataSource=${" + OMEGAMON_DS + "}"
                + "&var-managedSystem=" + OMEG_PLEX_LPAR_MVSSYS_EXPR
                + "&var-_addressSpaceName=" + JOB_NAME_EXPR
                + "&var-_ASID=" + ASID_EXPR
                + "&var-_iAddressSpaceName=" + JOB_NAME_EXPR
                + "&var-_iASID=" + ASID_EXPR
                + "&from=$__from"
                + "&to=$__to",
            }
          ]
        }
      ]
    });
  });
}

function falconSystemUpdate(falcon: FalconStatus, dashboard: any) {
  addVars(dashboard);
  let baseUrl = falcon.sysDashboard;
  if (baseUrl.indexOf("?") > -1) {
    baseUrl += "&";
  } else {
    baseUrl += "?";
  }

  if (!dashboard.links) {
    dashboard.links = [];
  }

  dashboard.links.push({
    type: "link",
    icon: "dashboard",
    keepTime: true,
    targetBlank: true,
    title: "z/OS Enterprise Overview",
    url: baseUrl + "var-dataSource=${" + OMEGAMON_DS + "}"
      + "&var-managedSystem=" + OMEG_PLEX_LPAR_MVSSYS_EXPR,
  });
}

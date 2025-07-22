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
const FOLDERS_API = '/api/folders';
const DASHBOARDS_API = '/api/dashboards/db';

// Don't use `getBackendSrv` to avoid built-in notifications: they won't work as we need them.

export async function findFolder(id: string): Promise<string | undefined> {
  const res = await fetch(`${FOLDERS_API}/${id}`);
  if (res.ok) {
    const data = await res.json();
    if (data.parents) {
      const parentPath = data.parents.map((parent) => parent.title).join(' / ');
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

async function createDashboard(folderUid: string, dashboard: object) {
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

export async function installDashboards(folderUid: string, defaultFolderName: string, dashboards: object[]) {
  const folderPath = await findFolder(folderUid);
  if (folderPath === undefined) {
    await createFolder(folderUid, defaultFolderName);
  }
  await Promise.all(dashboards.map((dashboard) => createDashboard(folderUid, dashboard)));
}

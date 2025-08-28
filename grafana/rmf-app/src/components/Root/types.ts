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

export enum OperCode {
  None,
  Install,
  Reset,
  Delete,
}

export enum OperStatus {
  None,
  Done,
  InProgress,
  Error,
}

export interface Operation {
  code: OperCode;
  status: OperStatus;
}

export interface FolderStatus {
  folderPath: string;
  installed: boolean;
  operation: Operation;
}

export interface FalconStatus {
  enabled: boolean;
  asDashboard: string
  sysDashboard: string
}
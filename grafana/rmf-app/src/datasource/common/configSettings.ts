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
export class ConfigSettings {
  static UrlSettings = {
    ROOT_URL: `/gpm/root.xml`,
    METRICS_URL: `/gpm/include/metrics.xml`,
    HEADERS_URL: `/gpm/include/reptrans.xsl`,
    PERFORM_URL: `/gpm/perform.xml?resource=%22`,
    RMFM3_URL: `/gpm/rmfm3.xml?report=`,
    CONTAINED_URL: `/gpm/contained.xml?resource=%22`,
    LISTMETRICS_URL: `/gpm/listmetrics.xml?resource=%22`,
    INDEXXML_URL: `/gpm/index.xml`,
    IMAGE_URL: `/gpm/include/`,
  };
}

export const OMEGAMON_DS_TYPE_NAME = "omegamon-datasource";
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
import { Parser } from './parser';
import 'jest';

describe('Test Report', () => {
  test('Should retuen formated outcome', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CHANNEL {Name=APPC_STR1}');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('report=CHANNEL&resource=,APPC_STR1,MVS_IMAGE');
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser('SYSPLEX.REPORT.CHANNEL');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('report=CHANNEL&resource=,,SYSPLEX');
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser('CPC.REPORT.CHANNEL');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('report=CHANNEL&resource=,,CPC');
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CPC {ulq=TEST12,Name=APPC_STR1}');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('report=CPC&resource=TEST12,APPC_STR1,MVS_IMAGE');
  });

  test('Should retuen mismatched resource name', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toBe("<Error 'reportName'>");
  });

  test('Should retuen mismatched resource type', () => {
    const parser = new Parser('MVS_IAGE.REPORT.CHANNEL {ulq=TEST12,Name=APPC_STR1}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "line 1, col 0: mismatched input 'MVS_IAGE' expecting Res_Type"
    );
  });

  test('Should check if a name is mismatched input.', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CHANNEL {ame=APPC_STR1}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "line 1, col 26: mismatched input 'ame' expecting {String_ULQ, String_NAME, YYYYMMDDhhmmss_RANGE, String_FILTER, String_WORKSCOPE}"
    );
  });
});

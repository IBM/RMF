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
  test('Should return formated outcome', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CHANNEL {Name=APPC_STR1}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe('report=CHANNEL&resource=,APPC_STR1,MVS_IMAGE');
  });

  test('Should return formated outcome', () => {
    const parser = new Parser('SYSPLEX.REPORT.CHANNEL');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe('report=CHANNEL&resource=,,SYSPLEX');
  });

  test('Should return formated outcome', () => {
    const parser = new Parser('CPC.REPORT.CHANNEL');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe('report=CHANNEL&resource=,,CPC');
  });

  test('Should return formated outcome', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CPC {ulq=TEST12,Name=APPC_STR1}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe('report=CPC&resource=TEST12,APPC_STR1,MVS_IMAGE');
  });

  test('Should return mismatched resource name', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain("line 1, col 17: missing {REPORT, WORKSCOPE, RANGE, ULQ, NAME, FILTER, PAT, LB, UB, HI, LO, ORD, ORD_OPTION, RES_TYPE, INTEGER, IDENTIFIER, STRING_UNQUOTED}");
  });

  test('Should return mismatched resource type', () => {
    const parser = new Parser('MVS_IAGE.REPORT.CHANNEL {ulq=TEST12,Name=APPC_STR1}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain(
      "line 1, col 0: mismatched input 'MVS_IAGE' expecting {RES_TYPE, WS}"
    );
  });

  test('Should check if a name is mismatched input.', () => {
    const parser = new Parser('MVS_IMAGE.REPORT.CHANNEL {ame=APPC_STR1}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain(
      "line 1, col 26: mismatched input 'ame' expecting {WORKSCOPE, ULQ, NAME, FILTER, WS}"
    );
  });
});

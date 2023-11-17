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

describe('Test Range Qualifier', () => {
  test('Should retuen formated outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP/+) by MVS image {range=20090725100000,20090725100000}');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe(
      'id=% CPU utilization (CP/+) by MVS image&resource=,,SYSPLEX&range=20090725100000,20090725100000'
    );
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP#/) by MVS image {ulq=test12,name=APPC_STR1,range=20090725100000,20090725100000}'
    );
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe(
      'id=% CPU utilization (CP#/) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&range=20090725100000,20090725100000'
    );
  });

  test('Should check if a dot is missing.', () => {
    const parser = new Parser('SYSPLEX,% CPU utilization (CP) by MVS image {range=20090725100000}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 7: no viable alternative at input 'SYSPLEX,'");
  });

  test('Should check if a comma is mismatched input.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {range=20090725100000}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 65: mismatched input '}' expecting ','");
  });

  test('Should check if a range is mismatched input.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,rang}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "line 1, col 69: mismatched input 'rang' expecting {String_ULQ, String_NAME, YYYYMMDDhhmmss_RANGE, String_FILTER, String_WORKSCOPE}"
    );
  });
});

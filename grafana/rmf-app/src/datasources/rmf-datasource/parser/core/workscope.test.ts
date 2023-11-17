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

describe('Test Workscope Qualifier', () => {
  test("Should retuen formated workscope type 'G' outcome", () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOPE=,,G}');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&WORKSCOPE=,,G'
    );
  });
  test("Should retuen formated workscope 'U,N,G' outcome", () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=ORD=nD,WORKSCOPE=U,N,G}'
    );
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=ORD=nD&WORKSCOPE=U,N,G'
    );
  });
  test('Should check if a workscope type parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOPE=U,N,}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "line 1, col 83: missing {'G', 'W', 'S', 'P', 'R', 'J'} at '}'"
    );
  });
  test('Should check if a workscope ULQ parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOPE=}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 79: mismatched input '}' expecting ','");
  });
  test('Should check if a workscope name parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOPE=,}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 80: mismatched input '}' expecting ','");
  });
  test('Should check if a workscope inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOP}');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "line 1, col 69: mismatched input 'WORKSCOP' expecting {String_ULQ, String_NAME, YYYYMMDDhhmmss_RANGE, String_FILTER, String_WORKSCOPE}"
    );
  });
});

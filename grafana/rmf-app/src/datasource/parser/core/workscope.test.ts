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
  test("Should return formated workscope type 'G' outcome", () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOPE=,,G}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&workscope=,,G'
    );
  });
  test("Should return formated workscope 'U,N,G' outcome", () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=ORD=nD,workscope=U,N,G}'
    );
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=ORD=nD&workscope=U,N,G'
    );
  });
  test('Should check if a workscope type parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,workscope=U,N,}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain(
      "line 1, col 85: missing WORKSCOPE_TYPE at '}'"
    );
  });
  test('Should check if a workscope ULQ parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,workscope=}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain(" expecting {REPORT, WORKSCOPE, RANGE, ULQ, NAME, FILTER, PAT, LB, UB, HI, LO, ORD, ORD_OPTION, RES_TYPE, INTEGER, IDENTIFIER, STRING_UNQUOTED, STRING_QUOTED, ','}");
  });
  test('Should check if a workscope name parameter inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,workscope=,}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain("expecting {REPORT, WORKSCOPE, RANGE, ULQ, NAME, FILTER, PAT, LB, UB, HI, LO, ORD, ORD_OPTION, RES_TYPE, INTEGER, IDENTIFIER, STRING_UNQUOTED, STRING_QUOTED, ','}");
  });
  test('Should check if a workscope inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,WORKSCOP}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain(
      "line 1, col 71: mismatched input 'WORKSCOP' expecting {WORKSCOPE, ULQ, NAME, FILTER, WS}"
    );
  });
});

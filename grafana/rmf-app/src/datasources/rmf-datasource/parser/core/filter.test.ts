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

describe('Test Filter Qualifier', () => {
  test('Should return formated filter PAT outcome', () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=Pat=test12*}'
    );
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=Pat=test12*'
    );
  });
  test('Should return formated filter UB outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {filter=ub=22,ulq=test12,name=APPC_STR1}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=ub=22'
    );
  });
  test('Should return formated filter LB outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=lb=33}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=lb=33'
    );
  });
  test('Should return formated multiple qualification outcome', () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=lb=33,filter=ord=na}'
    );
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=lb=33&filter=ord=na'
    );
  });
  test('Should return formated filter HI outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=HI=33}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=HI=33'
    );
  });
  test('Should return formated filter LO outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=lO=33}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=lO=33'
    );
  });
  test('Should return formated filter ORD outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=ORD=nD}');
    expect(parser.parse().errorFound).toBeFalsy();
    expect(parser.parse().query).toBe(
      'id=% CPU utilization (CP) by MVS image&resource=TEST12,APPC_STR1,SYSPLEX&filter=ORD=nD'
    );
  });
  test('Should check if a filter has mismatched multiple inputs.', () => {
    const parser = new Parser(
      'SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=ORD=nD=LO=33}'
    );
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain("line 1, col 84: mismatched input '=' expecting {',', '}', WS}");
  });
  test('Should check if a filter has mismatched inputs.', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image {ulq=test12,name=APPC_STR1,filter=LO=}');
    expect(parser.parse().errorFound).toBeTruthy();
    expect(parser.parse().errorMessage).toContain("line 1, col 81: missing INTEGER at '}'");
  });
});

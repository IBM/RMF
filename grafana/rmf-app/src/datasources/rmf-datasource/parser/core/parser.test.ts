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

describe('Test ResourceType and Metrics', () => {
  test('Should parse with message', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image');
    expect(parser.parse()).toBe('SYSPLEX.% CPU utilization (CP) by MVS image');
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser('SYSPLEX.% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('id=% CPU utilization (CP) by MVS image&resource=,,SYSPLEX');
  });

  test('Should retuen formated outcome', () => {
    const parser = new Parser('CPC.% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeFalsy();
    expect(parser.constructTree().query).toBe('id=% CPU utilization (CP) by MVS image&resource=,,CPC');
  });

  test('Should retuen mismatched Resource type parameter', () => {
    const parser = new Parser('TEST1.% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 0: mismatched input 'TEST1' expecting Res_Type");
  });

  test('Should check if mismatched Resource type missing.', () => {
    const parser = new Parser('.');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 0: mismatched input '.' expecting Res_Type");
  });

  test('Should check if a resource type is missing.', () => {
    const parser = new Parser('.% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 0: mismatched input '.' expecting Res_Type");
  });

  test('Should check if a resource type syntex is wrong.', () => {
    const parser = new Parser('@.% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("null line 1, col 0: token recognition error at: '@'");
  });

  test('Should check if a resource type & metrics syntex is wrong.', () => {
    const parser = new Parser('@.@');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "null line 1, col 0: token recognition error at: '@'null line 1, col 2: token recognition error at: '@'"
    );
  });

  test('Should check if a metric is missing.', () => {
    const parser = new Parser('SYSPLEX.');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("<Error 'metrics | report'>");
  });

  test('Should check if a metric is wrong syntex.', () => {
    const parser = new Parser('SYSPLEX.@');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain(
      "<Error 'metrics | report'>null line 1, col 8: token recognition error at: '@'"
    );
  });

  test('Should check if a dot is missing.', () => {
    const parser = new Parser('SYSPLEX,% CPU utilization (CP) by MVS image');
    expect(parser.constructTree().errorFound).toBeTruthy();
    expect(parser.constructTree().errorMessage).toContain("line 1, col 7: no viable alternative at input 'SYSPLEX,'");
  });
});

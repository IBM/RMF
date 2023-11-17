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
import { GrammarResult } from './type';
import { ErrorListener, Recognizer } from 'antlr4';

export class CustomErrorListener extends ErrorListener<any> {
  private lexerGrammarResult: GrammarResult;

  constructor() {
    super();
    this.lexerGrammarResult = {
      errorFound: false,
      errorMessage: '',
      query: '',
    };
  }

  syntaxError(
    recognizer: Recognizer<any>,
    offendingSymbol: any,
    line: number,
    column: number,
    msg: string,
    e: any
  ) {
    this.lexerGrammarResult.errorFound = true;
    this.lexerGrammarResult.errorMessage += `${offendingSymbol} line ${line}, col ${column}: ${msg}`;
  }

  reportAmbiguity(
    recognizer: Recognizer<any>,
    dfa: any,
    startIndex: number,
    stopIndex: number,
    exact: any,
    ambigAlts: any,
    configs: any
  ) {
    // this.lexerGrammarResult.errorFound = true;
  }

  reportAttemptingFullContext(
    recognizer: Recognizer<any>,
    dfa: any,
    startIndex: number,
    stopIndex: number,
    conflictingAlts: any,
    configs: any
  ) {
    // this.lexerGrammarResult.errorFound = true;
  }

  reportContextSensitivity(
    recognizer: Recognizer<any>,
    dfa: any,
    startIndex: number,
    stopIndex: number,
    conflictingAlts: any,
    configs: any
  ) {
    // this.lexerGrammarResult.errorFound = true;
  }

  getResult = () => {
    return this.lexerGrammarResult;
  };
}

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
import { InputStream, CommonTokenStream, ParseTreeWalker } from 'antlr4';
import RMFQueryLexer from '../lib/RMFQueryLexer';
import RMFQueryParser from '../lib/RMFQueryParser';
import { CustomErrorListener } from './customErrorListener';
import { CustomListener } from './customListener';
import { GrammarResult } from './type';

export class Parser {
  private queryLineText: string;
  lexerGrammarResult: GrammarResult;
  parserGrammarResult: GrammarResult;

  constructor(queryLine: string) {
    this.queryLineText = queryLine;

    this.lexerGrammarResult = {
      errorFound: false,
      errorMessage: '',
      query: '',
    };

    this.parserGrammarResult = {
      errorFound: false,
      errorMessage: '',
      query: '',
    };
  }

  parse(): string {
    return this.queryLineText;
  }

  constructTree(): GrammarResult {
    const lexerCustomErrorListener = new CustomErrorListener();
    const chars = new InputStream(this.queryLineText);

    const lexer = new RMFQueryLexer(chars);
    lexer.removeErrorListeners();
    lexer.addErrorListener(lexerCustomErrorListener);

    const tokens = new CommonTokenStream(lexer);
    const parser = new RMFQueryParser(tokens);
    parser.removeErrorListeners();
    parser.addErrorListener(lexerCustomErrorListener);
    parser.buildParseTrees = true;

    const tree = parser.rmfquery();
    const listener = new CustomListener();
    ParseTreeWalker.DEFAULT.walk(listener, tree);

    this.lexerGrammarResult = lexerCustomErrorListener.getResult();
    this.parserGrammarResult = listener.getTable();

    if (this.lexerGrammarResult.errorFound) {
      this.parserGrammarResult.errorFound = true;
      this.parserGrammarResult.errorMessage += this.lexerGrammarResult.errorMessage.trim();
    }

    return this.parserGrammarResult; // this function returns the start and stop indices.
  }
}

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
import { CharStream, CommonTokenStream, ParseTreeWalker } from 'antlr4';
import RMFQueryLexer from '../lib/RMFQueryLexer';
import RMFQueryParser, { QualifierContext, UlqContext } from '../lib/RMFQueryParser';
import { CustomErrorListener } from './customErrorListener';
import { GrammarResult } from './type';

const CONTEXT_STR = 'context';

export class Parser {
  private readonly queryLineText: string;

  constructor(queryLine: string) {
    this.queryLineText = queryLine;
  }

  parse(): GrammarResult {
    let lexerGrammarResult: { query: string; errorMessage: string; errorFound: boolean };
    let parserGrammarResult = {
      errorFound: false,
      errorMessage: '',
      query: '',
    };

    const lexerCustomErrorListener = new CustomErrorListener();
    const chars = new CharStream(this.queryLineText);
    const lexer = new RMFQueryLexer(chars);
    lexer.removeErrorListeners();
    lexer.addErrorListener(lexerCustomErrorListener);

    const tokens = new CommonTokenStream(lexer);
    const parser = new RMFQueryParser(tokens);
    parser.removeErrorListeners();
    parser.addErrorListener(lexerCustomErrorListener);
    parser.buildParseTrees = true;

    const tree = parser.query();

    const resType = (tree.RES_TYPE()?.getText() || '').trim().toUpperCase();
    const isReport = tree.REPORT() !== null;
    const identifier = (tree.identifier()?.getText() || '').trim();

    let qualifierValues = {};
    let filters: string[] = [];
    for (let qual of tree.qualifiers()?.qualifier_list() || []) {
      let q;
      if ((q = qual.name())) {
        qualifierValues['name'] = (q.string_()?.getText() || '').trim().toUpperCase();
      }
      if ((q = qual.ulq())) {
        qualifierValues['ulq'] = (q.string_()?.getText() || '').trim().toUpperCase();
      }
      if ((q = qual.workscope())) {
        qualifierValues['workscope'] = (q.workscopeValue()?.getText() || '').trim().toUpperCase();
      }
      if ((q = qual.filter())) {
        filters.push((q.filterValue()?.getText() || '').trim());
      }
    }

    let resource = `${qualifierValues['ulq'] || ''},${qualifierValues['name'] || ''},${resType}`;
    let workscope = qualifierValues['workscope'];
    let query = `${isReport ? 'report' : 'id'}=${identifier}&resource=${resource}`;
    if (filters) {
      query += `&filter=${filters.join(';')}`;
    }
    if (workscope) {
      query += `&workscope=${workscope}`;
    }
    parserGrammarResult.query = query;

    lexerGrammarResult = lexerCustomErrorListener.getResult();
    if (lexerGrammarResult.errorFound) {
      parserGrammarResult.errorFound = true;
      parserGrammarResult.errorMessage += lexerGrammarResult.errorMessage.trim();
    }
    return parserGrammarResult;
  }
}

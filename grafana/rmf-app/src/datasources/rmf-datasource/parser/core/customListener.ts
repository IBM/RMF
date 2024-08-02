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
import { ErrorNode } from 'antlr4';
import RMFQueryListener from '../lib/RMFQueryListener';
import {
  MetricsContext,
  Name_QualificationContext,
  RepnameContext,
  ReptypeContext,
  RestypeContext,
  Ulq_QualificationContext,
  Range_QualificationContext,
  Filter_QualificationContext,
  Workscope_QualificationContext,
  RmfqueryReportNameContext,
} from '../lib/RMFQueryParser';
import { GrammarResult } from './type';

export class CustomListener extends RMFQueryListener {
  tableSource = '';
  metricsTempl = 'id=$4&resource=$1,$2,$3';
  reportTempl = 'report=$4&resource=$1,$2,$3';
  private parserGrammarResult: GrammarResult;
  filters : string[] = [];

  constructor() {
    super();
    this.parserGrammarResult = {
      errorFound: false,
      errorMessage: '',
      query: '$1,$2,$3',
    };
  }

  exitReptype(ctx: ReptypeContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.reportTempl.replace('$3', resText.trim());
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error 'reportType'>";
      }
    }
  }

  exitRepname(ctx: RepnameContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$4', resText.trim());
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error 'reportName'>";
      }
    }
  }

  exitRmfqueryReportName(ctx: RmfqueryReportNameContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$4', resText.trim());
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error 'reportName'>";
      }
    }
  }

  exitRestype(ctx: RestypeContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.metricsTempl.replace('$3', resText.trim());
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error 'Res_Type'>";
      }
    }
  }

  exitMetrics(ctx: MetricsContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$4', resText.trim());
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error 'metrics | report'>";
      }
    }
  }

  exitUlq_Qualification(ctx: Ulq_QualificationContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.parserGrammarResult.query.replace(
          '$1',
          resText.toUpperCase().replace('ULQ=', '').trim()
        );
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error '{ULQ=<T>}'>";
      }
    }
  }

  exitName_Qualification(ctx: Name_QualificationContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query = this.parserGrammarResult.query.replace(
          '$2',
          resText.toUpperCase().replace('NAME=', '').trim()
        );
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error '{Name=<T>}'>";
      }
    }
  }

  exitRange_Qualification(ctx: Range_QualificationContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query += '&' + resText.trim();
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error '{Range=<T>.<T>}'>";
      }
    }
  }

  exitFilter_Qualification(ctx: Filter_QualificationContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        var filter = resText.trim();
        filter = filter.replace(new RegExp("filter=", "i"), "");
        this.filters.push(filter);
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error '{Filter=<T=<T>>}'>";
      }
    }
  }

  exitWorkscope_Qualification(ctx: Workscope_QualificationContext) {
    if (!this.parserGrammarResult.errorFound) {
      const resText = ctx.getText() ? ctx.getText() : '';
      if (resText !== undefined && resText.trim() !== '') {
        this.parserGrammarResult.query += '&' + resText.trim();
      } else {
        this.parserGrammarResult.errorFound = true;
        this.parserGrammarResult.errorMessage += "<Error '{workscope=<T>,<T>,<[G | W | S | P | R | J]>}'>";
      }
    }
  }

  visitErrorNode(node: ErrorNode) {
    this.parserGrammarResult.errorFound = true;
  }

  combinedFilters(): string {
    var result : string = "";
    this.filters.forEach((element) => {
      if (result.trim() !== '') {
        result = result + ";" + element;
      } else {
        result = result + element;
      }
    });
    return result;
  }

  getTable() {
    this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$1', '');
    this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$2', '');
    this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$3', '');
    this.parserGrammarResult.query = this.parserGrammarResult.query.replace('$4', '');
    if (this.filters.length != 0) {
      this.parserGrammarResult.query = this.parserGrammarResult.query + "&filter=" + this.combinedFilters();
    }
    return this.parserGrammarResult;
  }
}

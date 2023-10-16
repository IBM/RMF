import {ParserRuleContext, ErrorNode, ParseTreeListener, TerminalNode} from 'antlr4';

import {RmfqueryContext} from './RMFQueryParser';

import {RmfqueryReportContext} from './RMFQueryParser';

import {RmfqueryReportNameContext} from './RMFQueryParser';

import {RmfqueryMetricesContext} from './RMFQueryParser';

import {RmfqueryWithQulContext} from './RMFQueryParser';

import {RmfqueryWoutQulContext} from './RMFQueryParser';

import {QualificationsContext} from './RMFQueryParser';

import {MorequalificationContext} from './RMFQueryParser';

import {QualificationContext} from './RMFQueryParser';

import {Ulq_QualificationContext} from './RMFQueryParser';

import {Name_QualificationContext} from './RMFQueryParser';

import {Name_optionsContext} from './RMFQueryParser';

import {Only_stringContext} from './RMFQueryParser';

import {String_withdotContext} from './RMFQueryParser';

import {String_withdot_withnumberContext} from './RMFQueryParser';

import {Dot_withstringContext} from './RMFQueryParser';

import {Resource_withstringContext} from './RMFQueryParser';

import {Range_QualificationContext} from './RMFQueryParser';

import {Range_ValueContext} from './RMFQueryParser';

import {Filter_QualificationContext} from './RMFQueryParser';

import {Filter_optionsContext} from './RMFQueryParser';

import {Pat_filterContext} from './RMFQueryParser';

import {Lb_filterContext} from './RMFQueryParser';

import {Ub_filterContext} from './RMFQueryParser';

import {Hi_filterContext} from './RMFQueryParser';

import {Lo_filterContext} from './RMFQueryParser';

import {Ord_filterContext} from './RMFQueryParser';

import {Workscope_QualificationContext} from './RMFQueryParser';

import {Workscope_optionsContext} from './RMFQueryParser';

import {Workscope_ulqContext} from './RMFQueryParser';

import {Workscope_nameContext} from './RMFQueryParser';

import {Workscope_typeContext} from './RMFQueryParser';

import {ReptypenameContext} from './RMFQueryParser';

import {ReptypeContext} from './RMFQueryParser';

import {RepnameContext} from './RMFQueryParser';

import {RestypeContext} from './RMFQueryParser';

import {MetricsContext} from './RMFQueryParser';

export default RMFQueryListener;

export declare class RMFQueryListener implements ParseTreeListener {
    constructor();

    enterRmfquery(ctx: RmfqueryContext): void;

    exitRmfquery(ctx: RmfqueryContext): void;

    enterRmfqueryReport(ctx: RmfqueryReportContext): void;

    exitRmfqueryReport(ctx: RmfqueryReportContext): void;

    enterRmfqueryReportName(ctx: RmfqueryReportNameContext): void;

    exitRmfqueryReportName(ctx: RmfqueryReportNameContext): void;

    enterRmfqueryMetrices(ctx: RmfqueryMetricesContext): void;

    exitRmfqueryMetrices(ctx: RmfqueryMetricesContext): void;

    enterRmfqueryWithQul(ctx: RmfqueryWithQulContext): void;

    exitRmfqueryWithQul(ctx: RmfqueryWithQulContext): void;

    enterRmfqueryWoutQul(ctx: RmfqueryWoutQulContext): void;

    exitRmfqueryWoutQul(ctx: RmfqueryWoutQulContext): void;

    enterQualifications(ctx: QualificationsContext): void;

    exitQualifications(ctx: QualificationsContext): void;

    enterMorequalification(ctx: MorequalificationContext): void;

    exitMorequalification(ctx: MorequalificationContext): void;

    enterQualification(ctx: QualificationContext): void;

    exitQualification(ctx: QualificationContext): void;

    enterUlq_Qualification(ctx: Ulq_QualificationContext): void;

    exitUlq_Qualification(ctx: Ulq_QualificationContext): void;

    enterName_Qualification(ctx: Name_QualificationContext): void;

    exitName_Qualification(ctx: Name_QualificationContext): void;

    enterName_options(ctx: Name_optionsContext): void;

    exitName_options(ctx: Name_optionsContext): void;

    enterOnly_string(ctx: Only_stringContext): void;

    exitOnly_string(ctx: Only_stringContext): void;

    enterString_withdot(ctx: String_withdotContext): void;

    exitString_withdot(ctx: String_withdotContext): void;

    enterString_withdot_withnumber(ctx: String_withdot_withnumberContext): void;

    exitString_withdot_withnumber(ctx: String_withdot_withnumberContext): void;

    enterDot_withstring(ctx: Dot_withstringContext): void;

    exitDot_withstring(ctx: Dot_withstringContext): void;

    enterResource_withstring(ctx: Resource_withstringContext): void;

    exitResource_withstring(ctx: Resource_withstringContext): void;

    enterRange_Qualification(ctx: Range_QualificationContext): void;

    exitRange_Qualification(ctx: Range_QualificationContext): void;

    enterRange_Value(ctx: Range_ValueContext): void;

    exitRange_Value(ctx: Range_ValueContext): void;

    enterFilter_Qualification(ctx: Filter_QualificationContext): void;

    exitFilter_Qualification(ctx: Filter_QualificationContext): void;

    enterFilter_options(ctx: Filter_optionsContext): void;

    exitFilter_options(ctx: Filter_optionsContext): void;

    enterPat_filter(ctx: Pat_filterContext): void;

    exitPat_filter(ctx: Pat_filterContext): void;

    enterLb_filter(ctx: Lb_filterContext): void;

    exitLb_filter(ctx: Lb_filterContext): void;

    enterUb_filter(ctx: Ub_filterContext): void;

    exitUb_filter(ctx: Ub_filterContext): void;

    enterHi_filter(ctx: Hi_filterContext): void;

    exitHi_filter(ctx: Hi_filterContext): void;

    enterLo_filter(ctx: Lo_filterContext): void;

    exitLo_filter(ctx: Lo_filterContext): void;

    enterOrd_filter(ctx: Ord_filterContext): void;

    exitOrd_filter(ctx: Ord_filterContext): void;

    enterWorkscope_Qualification(ctx: Workscope_QualificationContext): void;

    exitWorkscope_Qualification(ctx: Workscope_QualificationContext): void;

    enterWorkscope_options(ctx: Workscope_optionsContext): void;

    exitWorkscope_options(ctx: Workscope_optionsContext): void;

    enterWorkscope_ulq(ctx: Workscope_ulqContext): void;

    exitWorkscope_ulq(ctx: Workscope_ulqContext): void;

    enterWorkscope_name(ctx: Workscope_nameContext): void;

    exitWorkscope_name(ctx: Workscope_nameContext): void;

    enterWorkscope_type(ctx: Workscope_typeContext): void;

    exitWorkscope_type(ctx: Workscope_typeContext): void;

    enterReptypename(ctx: ReptypenameContext): void;

    exitReptypename(ctx: ReptypenameContext): void;

    enterReptype(ctx: ReptypeContext): void;

    exitReptype(ctx: ReptypeContext): void;

    enterRepname(ctx: RepnameContext): void;

    exitRepname(ctx: RepnameContext): void;

    enterRestype(ctx: RestypeContext): void;

    exitRestype(ctx: RestypeContext): void;

    enterMetrics(ctx: MetricsContext): void;

    exitMetrics(ctx: MetricsContext): void;

    visitTerminal(node: TerminalNode): void;

    visitErrorNode(node: ErrorNode): void;

    enterEveryRule(node: ParserRuleContext): void;

    exitEveryRule(node: ParserRuleContext): void;
}

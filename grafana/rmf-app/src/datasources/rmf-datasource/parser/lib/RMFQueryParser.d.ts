import {CommonTokenStream, Parser, ParserRuleContext, TerminalNode} from 'antlr4';


export declare class RmfqueryContext extends ParserRuleContext {

    rmfqueryReport(): RmfqueryReportContext;

    rmfqueryMetrices(): RmfqueryMetricesContext;

}

export declare class RmfqueryReportContext extends ParserRuleContext {

    reptype(): ReptypeContext;

    String_REPORT(): TerminalNode;

    rmfqueryReportName(): RmfqueryReportNameContext;

    qualifications(): QualificationsContext;

}

export declare class RmfqueryReportNameContext extends ParserRuleContext {

    reptypename(): ReptypenameContext;

    repname(): RepnameContext;

}

export declare class RmfqueryMetricesContext extends ParserRuleContext {

    rmfqueryWoutQul(): RmfqueryWoutQulContext;

    rmfqueryWithQul(): RmfqueryWithQulContext;

}

export declare class RmfqueryWithQulContext extends ParserRuleContext {

    rmfqueryWoutQul(): RmfqueryWoutQulContext;

    qualifications(): QualificationsContext;

}

export declare class RmfqueryWoutQulContext extends ParserRuleContext {

    restype(): RestypeContext;

    DOT(): TerminalNode;

    metrics(): MetricsContext;

}

export declare class QualificationsContext extends ParserRuleContext {

    OPEN_FLR_BRAK(): TerminalNode;

    qualification(): QualificationContext;

    morequalification(): MorequalificationContext;

    CLOSE_FLR_BRAK(): TerminalNode;

}

export declare class MorequalificationContext extends ParserRuleContext {

}

export declare class QualificationContext extends ParserRuleContext {

    ulq_Qualification(): Ulq_QualificationContext;

    name_Qualification(): Name_QualificationContext;

    range_Qualification(): Range_QualificationContext;

    filter_Qualification(): Filter_QualificationContext;

    workscope_Qualification(): Workscope_QualificationContext;

}

export declare class Ulq_QualificationContext extends ParserRuleContext {

    String_ULQ(): TerminalNode;

    StringCharacters(): TerminalNode;

}

export declare class Name_QualificationContext extends ParserRuleContext {

    String_NAME(): TerminalNode;

    name_options(): Name_optionsContext;

}

export declare class Name_optionsContext extends ParserRuleContext {

    Number_Digits(): TerminalNode;

    only_string(): Only_stringContext;

    string_withdot(): String_withdotContext;

    string_withdot_withnumber(): String_withdot_withnumberContext;

}

export declare class Only_stringContext extends ParserRuleContext {

    resource_withstring(): Resource_withstringContext;

}

export declare class String_withdotContext extends ParserRuleContext {

    DOT(): TerminalNode;

    dot_withstring(): Dot_withstringContext;

}

export declare class String_withdot_withnumberContext extends ParserRuleContext {

    DOT(): TerminalNode;

    Number_Digits(): TerminalNode;

    resource_withstring(): Resource_withstringContext;

}

export declare class Dot_withstringContext extends ParserRuleContext {

}

export declare class Resource_withstringContext extends ParserRuleContext {

    StringCharacters(): TerminalNode;

    Res_Type(): TerminalNode;

    Number_Digits(): TerminalNode;

}

export declare class Range_QualificationContext extends ParserRuleContext {

    YYYYMMDDhhmmss_RANGE(): TerminalNode;

    range_Value(): Range_ValueContext;

}

export declare class Range_ValueContext extends ParserRuleContext {

    COMMA(): TerminalNode;

}

export declare class Filter_QualificationContext extends ParserRuleContext {

    String_FILTER(): TerminalNode;

    filter_options(): Filter_optionsContext;

}

export declare class Filter_optionsContext extends ParserRuleContext {

    pat_filter(): Pat_filterContext;

    lb_filter(): Lb_filterContext;

    ub_filter(): Ub_filterContext;

    hi_filter(): Hi_filterContext;

    lo_filter(): Lo_filterContext;

    ord_filter(): Ord_filterContext;

}

export declare class Pat_filterContext extends ParserRuleContext {

    String_PAT(): TerminalNode;

    StringCharacters(): TerminalNode;

}

export declare class Lb_filterContext extends ParserRuleContext {

    String_LB(): TerminalNode;

    Number_Digits(): TerminalNode;

}

export declare class Ub_filterContext extends ParserRuleContext {

    String_UB(): TerminalNode;

    Number_Digits(): TerminalNode;

}

export declare class Hi_filterContext extends ParserRuleContext {

    String_HI(): TerminalNode;

    Number_Digits(): TerminalNode;

}

export declare class Lo_filterContext extends ParserRuleContext {

    String_LO(): TerminalNode;

    Number_Digits(): TerminalNode;

}

export declare class Ord_filterContext extends ParserRuleContext {

    String_ORD(): TerminalNode;

    Ord_Options(): TerminalNode;

}

export declare class Workscope_QualificationContext extends ParserRuleContext {

    String_WORKSCOPE(): TerminalNode;

    workscope_options(): Workscope_optionsContext;

}

export declare class Workscope_optionsContext extends ParserRuleContext {

    workscope_ulq(): Workscope_ulqContext;

    workscope_name(): Workscope_nameContext;

    workscope_type(): Workscope_typeContext;

}

export declare class Workscope_ulqContext extends ParserRuleContext {

}

export declare class Workscope_nameContext extends ParserRuleContext {

}

export declare class Workscope_typeContext extends ParserRuleContext {

}

export declare class ReptypenameContext extends ParserRuleContext {

    Res_Type(): TerminalNode;

}

export declare class ReptypeContext extends ParserRuleContext {

    Res_Type(): TerminalNode;

}

export declare class RepnameContext extends ParserRuleContext {

    StringCharacters(): TerminalNode;

}

export declare class RestypeContext extends ParserRuleContext {

    Res_Type(): TerminalNode;

}

export declare class MetricsContext extends ParserRuleContext {

    StringCharacters(): TerminalNode;

}

export default RMFQueryParser;

export declare class RMFQueryParser extends Parser {
    readonly ruleNames: string[];
    readonly literalNames: string[];
    readonly symbolicNames: string[];

    constructor(input: CommonTokenStream);

    rmfquery(): RmfqueryContext;

    rmfqueryReport(): RmfqueryReportContext;

    rmfqueryReportName(): RmfqueryReportNameContext;

    rmfqueryMetrices(): RmfqueryMetricesContext;

    rmfqueryWithQul(): RmfqueryWithQulContext;

    rmfqueryWoutQul(): RmfqueryWoutQulContext;

    qualifications(): QualificationsContext;

    morequalification(): MorequalificationContext;

    qualification(): QualificationContext;

    ulq_Qualification(): Ulq_QualificationContext;

    name_Qualification(): Name_QualificationContext;

    name_options(): Name_optionsContext;

    only_string(): Only_stringContext;

    string_withdot(): String_withdotContext;

    string_withdot_withnumber(): String_withdot_withnumberContext;

    dot_withstring(): Dot_withstringContext;

    resource_withstring(): Resource_withstringContext;

    range_Qualification(): Range_QualificationContext;

    range_Value(): Range_ValueContext;

    filter_Qualification(): Filter_QualificationContext;

    filter_options(): Filter_optionsContext;

    pat_filter(): Pat_filterContext;

    lb_filter(): Lb_filterContext;

    ub_filter(): Ub_filterContext;

    hi_filter(): Hi_filterContext;

    lo_filter(): Lo_filterContext;

    ord_filter(): Ord_filterContext;

    workscope_Qualification(): Workscope_QualificationContext;

    workscope_options(): Workscope_optionsContext;

    workscope_ulq(): Workscope_ulqContext;

    workscope_name(): Workscope_nameContext;

    workscope_type(): Workscope_typeContext;

    reptypename(): ReptypenameContext;

    reptype(): ReptypeContext;

    repname(): RepnameContext;

    restype(): RestypeContext;

    metrics(): MetricsContext;

}

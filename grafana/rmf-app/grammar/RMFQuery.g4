grammar RMFQuery;

/* @header {
	package antlr;
} */

rmfquery
	: (rmfqueryReport | rmfqueryMetrices) EOF
	;

rmfqueryReport
	: ( reptype DOT String_REPORT DOT rmfqueryReportName qualifications)
	;

rmfqueryReportName
	:(reptypename | repname{1-100})
	;

rmfqueryMetrices
	: (rmfqueryWoutQul | rmfqueryWithQul)
	;

rmfqueryWithQul
	: ( rmfqueryWoutQul qualifications )
	;

rmfqueryWoutQul
	: ( restype DOT metrics{1-200} )
	;	

qualifications
	: ( OPEN_FLR_BRAK qualification morequalification CLOSE_FLR_BRAK )?
	;

morequalification
	: (COMMA qualification)*
	;

qualification
	: ( ulq_Qualification | name_Qualification | range_Qualification | filter_Qualification | workscope_Qualification )
	;

ulq_Qualification
	: String_ULQ '=' StringCharacters?
	;

String_ULQ
	: ([uU][lL][qQ])
	;

name_Qualification
	: String_NAME '=' name_options
	;

String_NAME
	: ([nN][aA][mM][eE])
	;

name_options
	: ( Number_Digits | only_string | string_withdot | string_withdot_withnumber)
	;

only_string
	: resource_withstring?
	;

string_withdot
	: ( resource_withstring? DOT resource_withstring?  dot_withstring )
	;

string_withdot_withnumber
	: ( resource_withstring? DOT Number_Digits )
	;

dot_withstring
	: ( DOT resource_withstring? )+
	;

resource_withstring
	: ( StringCharacters |  Res_Type | Number_Digits)
	;

range_Qualification
	: YYYYMMDDhhmmss_RANGE '=' range_Value
	;

YYYYMMDDhhmmss_RANGE
	: ([rR][aA][nN][gG][eE])
	;

range_Value
	: ( DateTime COMMA DateTime)
	;

filter_Qualification
	: String_FILTER '=' filter_options
	;

filter_options
	:  ( pat_filter | lb_filter | ub_filter | hi_filter | lo_filter | ord_filter )
	;

String_FILTER
	: ([fF][iI][lL][tT][eE][rR])
	;

pat_filter
	:  ( String_PAT '=' StringCharacters )
	;

lb_filter
	:  ( String_LB '=' Number_Digits )
	;

ub_filter
	:  ( String_UB '=' Number_Digits )
	;

hi_filter
	:  ( String_HI '=' Number_Digits )
	;

lo_filter
	:  ( String_LO '=' Number_Digits )
	;

ord_filter
	:  ( String_ORD '=' Ord_Options )
	;
	
String_PAT
	: ([pP][aA][tT])
	;
	
String_LB
	: ([lL][bB])
	;
	
String_UB
	: ([uU][bB])
	;
	
String_HI
	: ([hH][iI])
	;
	
String_LO
	: ([lL][oO])
	;
	
String_ORD
	: ([oO][rR][dD])
	;

String_REPORT
	: ([rR][eE][pP][oO][rR][tT])
	;

workscope_Qualification
	: String_WORKSCOPE '=' workscope_options
	;

workscope_options
	: ( workscope_ulq ',' workscope_name ',' workscope_type )
	;

String_WORKSCOPE
	: ([wW][oO][rR][kK][sS][cC][oO][pP][eE])
	;

workscope_ulq
	: StringCharacters*
	;

workscope_name
	: StringCharacters*
	;

workscope_type
	: ( 'G' | 'W' | 'S' | 'P' | 'R' | 'J' )
	;

reptypename: Res_Type
	;

reptype: Res_Type
	;

repname: StringCharacters?
	;

restype: Res_Type
	;
	
metrics
	:	StringCharacters?
	;

Ord_Options
	: (  ([nN][aA]) | ([nN][dD]) | ([vV][aA]) | ([vV][dD]) | ([nN][nN]) )
	;

Res_Type
	: ( ([sS][yY][sS][pP][lL][eE][xX]) 
	| ([mM][vV][sS][_][iI][mM][aA][gG][eE]) 
	| ([iI][/][oO][_][sS][uU][bB][sS][yY][sS][tT][eE][mM])
	| ([aA][lL][lL][_][sS][sS][iI][dD][sS])
	| ([aA][lL][lL][_][lL][cC][uU][sS])
	| ([aA][lL][lL][_][cC][hH][aA][nN][nN][eE][lL][sS])
	| ([aA][lL][lL][_][vV][oO][lL][uU][mM][eE][sS])
	| ([cC][rR][yY][pP][tT][oO])
	| ([pP][cC][iI][eE])
	| ([sS][cC][mM])
	| ([zZ][fF][sS])
	| ([aA][gG][gG][rR][eE][gG][aA][tT][eE])
	| ([pP][rR][oO][cC][eE][sS][sS][oO][rR])
	| ([sS][tT][oO][rR][aA][G][eE])
	| ([aA][uU][xX][iI][lL][iI][aA][rR][yY][_][sS][tT][oO][rR][aA][G][eE])
	| ([cC][eE][nN][tT][rR][aA][lL][_][sS][tT][oO][rR][aA][G][eE])
	| ([cC][sS][aA])
	| ([sS][qQ][aA])
	| ([eE][cC][sS][aA])
	| ([eE][nN][qQ][uU][eE][uU][eE])
	| ([oO][pP][eE][rR][aA][tT][oO][rR])
	| ([sS][wW][_][sS][uU][bB][sS][yY][sS][tT][eE][mM][sS])
	| ([jJ][eE][sS])
	| ([xX][cC][fF])
	| ([hH][sS][mM])
	| ([cC][pP][cC])
	| ([lL][pP][aA][rR])
	| ([cC][oO][uU][pP][lL][iI][nN][G][_][fF][aA][cC][iI][lL][iI][tT][yY])
	| ([cC][fF][_][sS][tT][rR][uU][cC][tT][uU][rR][eE])
	| ([wW][lL][mM][_][aA][cC][tT][iI][vV][eE][_][pP][oO][lL][iI][cC][yY])
	| ([aA][lL][lL][_][wW][lL][mM][_][wW][oO][rR][kK][lL][oO][aA][dD][sS])
	| ([wW][lL][mM][_][wW][oO][rR][kK][lL][oO][aA][dD])
	| ([wW][lL][mM][_][sS][eE][rR][vV][iI][cC][eE][_][cC][lL][aA][sS][sS])
	| ([wW][lL][mM][_][sS][cC][_][pP][eE][rR][iI][oO][dD])
	| ([aA][lL][lL][_][wW][lL][mM][_][rR][eE][pP][oO][rR][tT][_][cC][lL][aA][sS][sS][eE][sS])
	| ([wW][lL][mM][_][rR][eE][pP][oO][rR][tT][_][cC][lL][aA][sS][sS])
	| ([wW][lL][mM][_][rR][cC][_][pP][eE][rR][iI][oO][dD])
	| ([aA][lL][lL][_][wW][lL][mM][_][rR][eE][sS][oO][uU][rR][cC][eE][_][G][rR][oO][uU][pP][sS])
	| ([wW][lL][mM][_][rR][eE][sS][oO][uU][rR][cC][eE][_][G][rR][oO][uU][pP])
	| ([cC][hH][aA][nN][nN][eE][lL][_][pP][aA][tT][hH])
	| ([lL][oO][gG][iI][cC][aA][lL][_][cC][oO][nN][tT][rR][oO][lL][_][uU][nN][iI][tT])
	| ([sS][sS][iI][dD])
	| ([vV][oO][lL][uU][mM][eE])
	| ([cC][rR][yY][pP][tT][oO][_][cC][aA][rR][dD])
	| ([pP][cC][iI][eE][_][fF][uU][nN][cC][tT][iI][oO][nN])
	| ([sS][cC][mM][_][cC][aA][rR][dD])
	| ([eE][sS][qQ][aA])
	| ([fF][iI][lL][eE][sS][yY][sS][tT][eE][mM])
	)
	;

DateTime
	: DATETIME
	;

Number_Digits
	: ( NEWDIGIT )+
	;

StringCharacters
	: TEXT+
	;

DATETIME
	: ('1'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')('0'..'9')
	;

NEWDIGIT: ('0'..'9')
	;

TEXT
	: [a-zA-Z0-9_#$%* ()+/-]
	;

DOT
	: '.'
	;

COMMA
	: ','
	;

OPEN_FLR_BRAK
	: '{'
	;

CLOSE_FLR_BRAK
	: '}'
	;

WS: [ \n\t\r]+ -> skip;

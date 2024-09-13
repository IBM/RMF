/*
RMF Query grammar

Examples:
1.  query:      restype.metric
    example:    SYSPLEX.% delay for enqueue
    output:     ",TESTPLEX01,SYSPLEX"&id=8D1A20
    ----------------------------------------
2.  query:      restype.metric  {name=resname}
    example:    SYSPLEX.% delay for enqueue {name=TESTPLEX01}
    output:     ",TESTPLEX01,SYSPLEX"&id=8D1A20
    ----------------------------------------
3.  query:      restype.metric  {name=resname}
    example:    MVS_IMAGE.% delay {name=NAME123}
    output:     ",NAME123,MVS_IMAGE"&id=8D0160
    ----------------------------------------
4.  query:      restype.metric  {name=resname, ulq=resource ulq}
    example:    CF_STRUCTURE.sync service time {ulq=CF01, name=APPC_STR1}
    output:     "CF01,APPC_STR1,CF_STRUCTURE"&id=8D39A0
    ----------------------------------------
5.  query:      restype.metric  {ulq=resource ulq, name=resname}
    example:    WLM_WORKLOAD.wait time (ms) by WLM service class period {name=*,ulq=ULQTEST01}
    output:     "ULQTEST01,*,ALL_WLM_WORKLOADS"&id=8D5E30
    ----------------------------------------
6.  query:      restype.metric  {ulq=resource ulq, name=resname}
    example:    WLM_WORKLOAD.active time (ms) {name=ASCH,ulq=ULQTEST01}
    output:     "ULQTEST01,ASCH,WLM_WORKLOAD"&id=8D5DF0
    ----------------------------------------
*/
grammar RMFQuery;

rmfquery
	: (rmfqueryReport | rmfqueryMetrices) EOF
	;

rmfqueryReport
	: ( reptype DOT String_REPORT DOT rmfqueryReportName qualifications)
	;

rmfqueryReportName
	: (reptypename | repname{1-100})
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
	: ( Integer | only_string | string_withdot | string_withdot_withnumber)
	;

only_string
	: resource_withstring?
	;

string_withdot
	: ( resource_withstring? DOT resource_withstring?  dot_withstring )
	;

string_withdot_withnumber
	: ( resource_withstring? DOT Integer )
	;

dot_withstring
	: ( DOT resource_withstring? )+
	;

resource_withstring
	: ( StringCharacters |  Res_Type | Integer )
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
	: String_LB '=' ( Integer | Decimal )
	;

ub_filter
	: String_UB '=' ( Integer | Decimal)
	;

hi_filter
	:  ( String_HI '=' Integer )
	;

lo_filter
	:  ( String_LO '=' Integer )
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
	: workscope_ulq ',' workscope_name ',' workscope_type
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
	: N A
	| N D
	| V A
	| V D
	| N N
	;

Res_Type
	: S Y S P L E X
	| M V S '_' I M A G E
	| I '/' O '_' S U B S Y S T E M
	| A L L '_' S S I D S
	| A L L '_' L C U S
	| A L L '_' C H A N N E L S
	| A L L '_' V O L U M E S
	| C R Y P T O
	| P C I E
	| S C M
	| Z F S
	| A G G R E G A T E
	| P R O C C E S S O R
	| S T O R A G E
	| A U X I L I A R Y '_' S T O R A G E
	| C E N T R A L '_' S T O R A G E
	| C S A
	| S Q A
	| E C S A
	| E N Q U E U E
	| O P E R A T O R
	| S W '_' S U B S Y S T E M S
	| J E S
	| X C F
	| H S M
	| C P C
	| L P A R
	| C O U P L I N G '_' F A C I L I T Y
	| C F '_' S T R U C T U R E
	| W L M '_' A C T I V E '_' P O L I C Y
	| A L L '_' W L M '_' W O R K L O A D S
	| W L M '_' W O R K L O A D
	| W L M '_' S E R V I C E '_' C L A S S S
	| W L M '_' S C '_' P E R I O D
	| A L L '_' W L M '_' R E P O R T '_' C L A S S S E S
	| W L M '_' R E P O R T '_' C L A S S S
	| W L M '_' R C '_' P E R I O D
	| A L L '_' W L M '_' R E S O U R C E '_' G R O U P S
	| W L M '_' R E S O U R C E '_' G R O U P
	| C H A N N E L '_' P A T H
	| L O G I C A L '_' C O N T R O L '_' U N I T
	| S S I D
	| V O L U M E
	| C R Y P T O '_' C A R D
	| P C I E '_' F U N C T I O N
	| S C M '_' C A R D
	| E S Q A
	| F I L E S Y S T E M
	;

DateTime
	: DATETIME
	;

Integer
	: NEWDIGIT+
	;

Decimal
  : Integer DOT Integer
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

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

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


query: WS* RES_TYPE (DOT REPORT)? DOT identifier WS* qualifiers? WS* EOF;
// A workaround: some reports are also resource TYPES (e.g. CPC).
// In general, the problem is that we define keywords that are not distiguashable for antlr from
// string literals which we also support.
identifier: stringSpaced;
qualifiers: LBRACE WS* qualifier WS* (COMMA WS* qualifier WS*)* RBRACE;
qualifier: ulq | name | filter | workscope;
ulq: ULQ EQUAL string;
name: NAME EQUAL string;
filter: FILTER EQUAL filterValue;
filterValue: pat | lb | ub | hi | lo | ord;
pat: PAT EQUAL string;
lb: LB EQUAL number;
ub: UB EQUAL number;
hi: HI EQUAL INTEGER;
lo: LO EQUAL INTEGER;
ord: ORD EQUAL ORD_OPTION;
workscope: WORKSCOPE EQUAL workscopeValue;
workscopeValue: string? COMMA string? COMMA WORKSCOPE_TYPE;

// Another workaround: it won't work on token level.
number: INTEGER | DECIMAL;
stringUnquoted
  : IDENTIFIER | RES_TYPE | REPORT | WORKSCOPE | RANGE | ULQ | NAME | FILTER
  | PAT | LB | UB | HI | LO | ORD | ORD_OPTION | INTEGER | STRING_UNQUOTED;
stringSpaced: stringUnquoted (WS + stringUnquoted)*;
stringDotted: stringUnquoted (DOT stringUnquoted)*;
string: stringDotted | STRING_QUOTED;


REPORT: R E P O R T;
WORKSCOPE: W O R K S C O P E;
RANGE: R A N G E;
ULQ: U L Q;
NAME: N A M E;
FILTER: F I L T E R;
PAT: P A T;
LB: L B;
UB: U B;
HI: H I;
LO: L O;
ORD: O R D;
ORD_OPTION: N A | N D | V A | V D | N N;
RES_TYPE
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
WORKSCOPE_TYPE: G | W | S | P | R | J;
INTEGER: [0-9]+;
DECIMAL: INTEGER DOT INTEGER;
IDENTIFIER: [a-zA-Z]+;
STRING_UNQUOTED: STRING_ITEM_NO_QUOTE+;
STRING_QUOTED
  : SINGLE_QUOTE STRING_ITEM_SINGLE_QUOTE* SINGLE_QUOTE
  | DOUBLE_QOUTE STRING_ITEM_DOUBLE_QUOTE* DOUBLE_QOUTE
  ;
DOT: '.';
COMMA: ',';
LBRACE: '{';
RBRACE: '}';
EQUAL: '=';
WS: [ \n\t\r]+;


fragment SINGLE_QUOTE: '\'';
fragment DOUBLE_QOUTE: '"';
fragment STRING_ITEM_NO_QUOTE: ~[ .{}=,];
fragment STRING_ITEM_SINGLE_QUOTE: ~'\'';
fragment STRING_ITEM_DOUBLE_QUOTE: ~'"';
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

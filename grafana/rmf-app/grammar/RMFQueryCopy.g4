grammar RMFQuery

rmfquery
	: (RESTYPE '.' METRICS)|(RESTYPE '.' METRICS QUALIFICATIONS)
	;	
	
RESTYPE
	: ^[a-zA-Z]+[a-zA-Z_]+([a-zA-Z])$
	;
	
METRICS
	:^.*$
	;
	
QUALIFICATIONS
	: '{' (QUALIFICATION (',' QUALIFICATION)*)? '}'
	;
	
QUALIFICATION
	: ULQ | NAME | RANGE |FILTER | WORKSCOPE
	;
	
ULQ
	: ^([uU][lL][qQ])=[a-zA-Z0-9]
	;
	
NAME
	: ^.*$
	;

RANGE
	: DATETIME '.' DATETIME
	;

DATETIME
	: \b([0-9]{4})(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])(2[0-3]|[01][0-9])([0-5][0-9])([0-5][0-9])\b
	;

FILTER	
	: PAT | LB | UB | HI | LO | ORD
	;
	
PAT
	: ^([pP][aA][tT])=.*\S.*
	;
	
LB
	: ^([lL][bB])=[0-9]
	;

UB
	: ^([uU][bB])=[0-9]
	;

HI
	: ^([hH][iI])=[0-9]
	;

LO
	: ^([lL][oO])=[0-9]
	;

ORD	
	: ^([oO][rR][dD])=\b(NA|ND|VA|VD)\b
	;

WORKSCOPE
	: WORKSCOPEULQ ',' WORKSCOPENAME ',' WORKSCOPETYPE
	;
	
WORKSCOPEULQ
	: [a-zA-Z0-9]
	;
	
WORKSCOPEULQ
	: ^$|^.*$
	;
	
WORKSCOPETYPE
	: \b(G|W|S|P|R|J)\b
	;
	

	
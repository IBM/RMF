// Generated from ./grammar/RMFQuery.g4 by ANTLR 4.13.2
// noinspection ES6UnusedImports,JSUnusedGlobalSymbols,JSUnusedLocalSymbols

import {
	ATN,
	ATNDeserializer, DecisionState, DFA, FailedPredicateException,
	RecognitionException, NoViableAltException, BailErrorStrategy,
	Parser, ParserATNSimulator,
	RuleContext, ParserRuleContext, PredictionMode, PredictionContextCache,
	TerminalNode, RuleNode,
	Token, TokenStream,
	Interval, IntervalSet
} from 'antlr4';
import RMFQueryListener from "./RMFQueryListener.js";
// for running tests with parameters, TODO: discuss strategy for typed parameters in CI
// eslint-disable-next-line no-unused-vars
type int = number;

export default class RMFQueryParser extends Parser {
	public static readonly REPORT = 1;
	public static readonly REPORT_CAPTION = 2;
	public static readonly REPORT_BANNER = 3;
	public static readonly WORKSCOPE = 4;
	public static readonly RANGE = 5;
	public static readonly ULQ = 6;
	public static readonly NAME = 7;
	public static readonly FILTER = 8;
	public static readonly PAT = 9;
	public static readonly LB = 10;
	public static readonly UB = 11;
	public static readonly HI = 12;
	public static readonly LO = 13;
	public static readonly ORD = 14;
	public static readonly ORD_OPTION = 15;
	public static readonly RES_TYPE = 16;
	public static readonly WORKSCOPE_TYPE = 17;
	public static readonly INTEGER = 18;
	public static readonly DECIMAL = 19;
	public static readonly IDENTIFIER = 20;
	public static readonly STRING_UNQUOTED = 21;
	public static readonly STRING_QUOTED = 22;
	public static readonly DOT = 23;
	public static readonly SEMI = 24;
	public static readonly COMMA = 25;
	public static readonly LBRACE = 26;
	public static readonly RBRACE = 27;
	public static readonly EQUAL = 28;
	public static readonly WS = 29;
	public static override readonly EOF = Token.EOF;
	public static readonly RULE_query = 0;
	public static readonly RULE_identifier = 1;
	public static readonly RULE_qualifiers = 2;
	public static readonly RULE_qualifier = 3;
	public static readonly RULE_ulq = 4;
	public static readonly RULE_name = 5;
	public static readonly RULE_filter = 6;
	public static readonly RULE_filterValue = 7;
	public static readonly RULE_filterItem = 8;
	public static readonly RULE_pat = 9;
	public static readonly RULE_lb = 10;
	public static readonly RULE_ub = 11;
	public static readonly RULE_hi = 12;
	public static readonly RULE_lo = 13;
	public static readonly RULE_ord = 14;
	public static readonly RULE_workscope = 15;
	public static readonly RULE_workscopeValue = 16;
	public static readonly RULE_number = 17;
	public static readonly RULE_stringUnquoted = 18;
	public static readonly RULE_stringSpaced = 19;
	public static readonly RULE_stringDotted = 20;
	public static readonly RULE_string = 21;
	public static readonly literalNames: (string | null)[] = [ null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, null, 
                                                            null, "'.'", 
                                                            "';'", "','", 
                                                            "'{'", "'}'", 
                                                            "'='" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "REPORT", 
                                                             "REPORT_CAPTION", 
                                                             "REPORT_BANNER", 
                                                             "WORKSCOPE", 
                                                             "RANGE", "ULQ", 
                                                             "NAME", "FILTER", 
                                                             "PAT", "LB", 
                                                             "UB", "HI", 
                                                             "LO", "ORD", 
                                                             "ORD_OPTION", 
                                                             "RES_TYPE", 
                                                             "WORKSCOPE_TYPE", 
                                                             "INTEGER", 
                                                             "DECIMAL", 
                                                             "IDENTIFIER", 
                                                             "STRING_UNQUOTED", 
                                                             "STRING_QUOTED", 
                                                             "DOT", "SEMI", 
                                                             "COMMA", "LBRACE", 
                                                             "RBRACE", "EQUAL", 
                                                             "WS" ];
	// tslint:disable:no-trailing-whitespace
	public static readonly ruleNames: string[] = [
		"query", "identifier", "qualifiers", "qualifier", "ulq", "name", "filter", 
		"filterValue", "filterItem", "pat", "lb", "ub", "hi", "lo", "ord", "workscope", 
		"workscopeValue", "number", "stringUnquoted", "stringSpaced", "stringDotted", 
		"string",
	];
	public get grammarFileName(): string { return "RMFQuery.g4"; }
	public get literalNames(): (string | null)[] { return RMFQueryParser.literalNames; }
	public get symbolicNames(): (string | null)[] { return RMFQueryParser.symbolicNames; }
	public get ruleNames(): string[] { return RMFQueryParser.ruleNames; }
	public get serializedATN(): number[] { return RMFQueryParser._serializedATN; }

	protected createFailedPredicateException(predicate?: string, message?: string): FailedPredicateException {
		return new FailedPredicateException(this, predicate, message);
	}

	constructor(input: TokenStream) {
		super(input);
		this._interp = new ParserATNSimulator(this, RMFQueryParser._ATN, RMFQueryParser.DecisionsToDFA, new PredictionContextCache());
	}
	// @RuleVersion(0)
	public query(): QueryContext {
		let localctx: QueryContext = new QueryContext(this, this._ctx, this.state);
		this.enterRule(localctx, 0, RMFQueryParser.RULE_query);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 47;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===29) {
				{
				{
				this.state = 44;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 49;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 50;
			this.match(RMFQueryParser.RES_TYPE);
			this.state = 57;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 1, this._ctx) ) {
			case 1:
				{
				{
				this.state = 51;
				this.match(RMFQueryParser.DOT);
				this.state = 52;
				this.match(RMFQueryParser.REPORT);
				}
				}
				break;
			case 2:
				{
				{
				this.state = 53;
				this.match(RMFQueryParser.DOT);
				this.state = 54;
				this.match(RMFQueryParser.REPORT_CAPTION);
				}
				}
				break;
			case 3:
				{
				{
				this.state = 55;
				this.match(RMFQueryParser.DOT);
				this.state = 56;
				this.match(RMFQueryParser.REPORT_BANNER);
				}
				}
				break;
			}
			this.state = 59;
			this.match(RMFQueryParser.DOT);
			this.state = 60;
			this.identifier();
			this.state = 64;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 2, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 61;
					this.match(RMFQueryParser.WS);
					}
					}
				}
				this.state = 66;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 2, this._ctx);
			}
			this.state = 68;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===26) {
				{
				this.state = 67;
				this.qualifiers();
				}
			}

			this.state = 73;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===29) {
				{
				{
				this.state = 70;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 75;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 76;
			this.match(RMFQueryParser.EOF);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public identifier(): IdentifierContext {
		let localctx: IdentifierContext = new IdentifierContext(this, this._ctx, this.state);
		this.enterRule(localctx, 2, RMFQueryParser.RULE_identifier);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 78;
			this.stringSpaced();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public qualifiers(): QualifiersContext {
		let localctx: QualifiersContext = new QualifiersContext(this, this._ctx, this.state);
		this.enterRule(localctx, 4, RMFQueryParser.RULE_qualifiers);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 80;
			this.match(RMFQueryParser.LBRACE);
			this.state = 84;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===29) {
				{
				{
				this.state = 81;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 86;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 87;
			this.qualifier();
			this.state = 91;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===29) {
				{
				{
				this.state = 88;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 93;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 110;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===25) {
				{
				{
				this.state = 94;
				this.match(RMFQueryParser.COMMA);
				this.state = 98;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				while (_la===29) {
					{
					{
					this.state = 95;
					this.match(RMFQueryParser.WS);
					}
					}
					this.state = 100;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				}
				this.state = 101;
				this.qualifier();
				this.state = 105;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				while (_la===29) {
					{
					{
					this.state = 102;
					this.match(RMFQueryParser.WS);
					}
					}
					this.state = 107;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				}
				}
				}
				this.state = 112;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 113;
			this.match(RMFQueryParser.RBRACE);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public qualifier(): QualifierContext {
		let localctx: QualifierContext = new QualifierContext(this, this._ctx, this.state);
		this.enterRule(localctx, 6, RMFQueryParser.RULE_qualifier);
		try {
			this.state = 119;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 6:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 115;
				this.ulq();
				}
				break;
			case 7:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 116;
				this.name();
				}
				break;
			case 8:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 117;
				this.filter();
				}
				break;
			case 4:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 118;
				this.workscope();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public ulq(): UlqContext {
		let localctx: UlqContext = new UlqContext(this, this._ctx, this.state);
		this.enterRule(localctx, 8, RMFQueryParser.RULE_ulq);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 121;
			this.match(RMFQueryParser.ULQ);
			this.state = 122;
			this.match(RMFQueryParser.EQUAL);
			this.state = 123;
			this.string_();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public name(): NameContext {
		let localctx: NameContext = new NameContext(this, this._ctx, this.state);
		this.enterRule(localctx, 10, RMFQueryParser.RULE_name);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 125;
			this.match(RMFQueryParser.NAME);
			this.state = 126;
			this.match(RMFQueryParser.EQUAL);
			this.state = 127;
			this.string_();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public filter(): FilterContext {
		let localctx: FilterContext = new FilterContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, RMFQueryParser.RULE_filter);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 129;
			this.match(RMFQueryParser.FILTER);
			this.state = 130;
			this.match(RMFQueryParser.EQUAL);
			this.state = 131;
			this.filterValue();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public filterValue(): FilterValueContext {
		let localctx: FilterValueContext = new FilterValueContext(this, this._ctx, this.state);
		this.enterRule(localctx, 14, RMFQueryParser.RULE_filterValue);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 133;
			this.filterItem();
			this.state = 138;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===24) {
				{
				{
				this.state = 134;
				this.match(RMFQueryParser.SEMI);
				this.state = 135;
				this.filterItem();
				}
				}
				this.state = 140;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public filterItem(): FilterItemContext {
		let localctx: FilterItemContext = new FilterItemContext(this, this._ctx, this.state);
		this.enterRule(localctx, 16, RMFQueryParser.RULE_filterItem);
		try {
			this.state = 147;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 9:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 141;
				this.pat();
				}
				break;
			case 10:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 142;
				this.lb();
				}
				break;
			case 11:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 143;
				this.ub();
				}
				break;
			case 12:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 144;
				this.hi();
				}
				break;
			case 13:
				this.enterOuterAlt(localctx, 5);
				{
				this.state = 145;
				this.lo();
				}
				break;
			case 14:
				this.enterOuterAlt(localctx, 6);
				{
				this.state = 146;
				this.ord();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public pat(): PatContext {
		let localctx: PatContext = new PatContext(this, this._ctx, this.state);
		this.enterRule(localctx, 18, RMFQueryParser.RULE_pat);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 149;
			this.match(RMFQueryParser.PAT);
			this.state = 150;
			this.match(RMFQueryParser.EQUAL);
			this.state = 151;
			this.string_();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public lb(): LbContext {
		let localctx: LbContext = new LbContext(this, this._ctx, this.state);
		this.enterRule(localctx, 20, RMFQueryParser.RULE_lb);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 153;
			this.match(RMFQueryParser.LB);
			this.state = 154;
			this.match(RMFQueryParser.EQUAL);
			this.state = 155;
			this.number_();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public ub(): UbContext {
		let localctx: UbContext = new UbContext(this, this._ctx, this.state);
		this.enterRule(localctx, 22, RMFQueryParser.RULE_ub);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 157;
			this.match(RMFQueryParser.UB);
			this.state = 158;
			this.match(RMFQueryParser.EQUAL);
			this.state = 159;
			this.number_();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public hi(): HiContext {
		let localctx: HiContext = new HiContext(this, this._ctx, this.state);
		this.enterRule(localctx, 24, RMFQueryParser.RULE_hi);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 161;
			this.match(RMFQueryParser.HI);
			this.state = 162;
			this.match(RMFQueryParser.EQUAL);
			this.state = 163;
			this.match(RMFQueryParser.INTEGER);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public lo(): LoContext {
		let localctx: LoContext = new LoContext(this, this._ctx, this.state);
		this.enterRule(localctx, 26, RMFQueryParser.RULE_lo);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 165;
			this.match(RMFQueryParser.LO);
			this.state = 166;
			this.match(RMFQueryParser.EQUAL);
			this.state = 167;
			this.match(RMFQueryParser.INTEGER);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public ord(): OrdContext {
		let localctx: OrdContext = new OrdContext(this, this._ctx, this.state);
		this.enterRule(localctx, 28, RMFQueryParser.RULE_ord);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 169;
			this.match(RMFQueryParser.ORD);
			this.state = 170;
			this.match(RMFQueryParser.EQUAL);
			this.state = 171;
			this.match(RMFQueryParser.ORD_OPTION);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public workscope(): WorkscopeContext {
		let localctx: WorkscopeContext = new WorkscopeContext(this, this._ctx, this.state);
		this.enterRule(localctx, 30, RMFQueryParser.RULE_workscope);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 173;
			this.match(RMFQueryParser.WORKSCOPE);
			this.state = 174;
			this.match(RMFQueryParser.EQUAL);
			this.state = 175;
			this.workscopeValue();
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public workscopeValue(): WorkscopeValueContext {
		let localctx: WorkscopeValueContext = new WorkscopeValueContext(this, this._ctx, this.state);
		this.enterRule(localctx, 32, RMFQueryParser.RULE_workscopeValue);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 178;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 7733246) !== 0)) {
				{
				this.state = 177;
				this.string_();
				}
			}

			this.state = 180;
			this.match(RMFQueryParser.COMMA);
			this.state = 182;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 7733246) !== 0)) {
				{
				this.state = 181;
				this.string_();
				}
			}

			this.state = 184;
			this.match(RMFQueryParser.COMMA);
			this.state = 185;
			this.match(RMFQueryParser.WORKSCOPE_TYPE);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public number_(): NumberContext {
		let localctx: NumberContext = new NumberContext(this, this._ctx, this.state);
		this.enterRule(localctx, 34, RMFQueryParser.RULE_number);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 187;
			_la = this._input.LA(1);
			if(!(_la===18 || _la===19)) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public stringUnquoted(): StringUnquotedContext {
		let localctx: StringUnquotedContext = new StringUnquotedContext(this, this._ctx, this.state);
		this.enterRule(localctx, 36, RMFQueryParser.RULE_stringUnquoted);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 189;
			_la = this._input.LA(1);
			if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 3538942) !== 0))) {
			this._errHandler.recoverInline(this);
			}
			else {
				this._errHandler.reportMatch(this);
			    this.consume();
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public stringSpaced(): StringSpacedContext {
		let localctx: StringSpacedContext = new StringSpacedContext(this, this._ctx, this.state);
		this.enterRule(localctx, 38, RMFQueryParser.RULE_stringSpaced);
		let _la: number;
		try {
			let _alt: number;
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 191;
			this.stringUnquoted();
			this.state = 200;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 16, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 193;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					do {
						{
						{
						this.state = 192;
						this.match(RMFQueryParser.WS);
						}
						}
						this.state = 195;
						this._errHandler.sync(this);
						_la = this._input.LA(1);
					} while (_la===29);
					this.state = 197;
					this.stringUnquoted();
					}
					}
				}
				this.state = 202;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 16, this._ctx);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public stringDotted(): StringDottedContext {
		let localctx: StringDottedContext = new StringDottedContext(this, this._ctx, this.state);
		this.enterRule(localctx, 40, RMFQueryParser.RULE_stringDotted);
		let _la: number;
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 203;
			this.stringUnquoted();
			this.state = 208;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===23) {
				{
				{
				this.state = 204;
				this.match(RMFQueryParser.DOT);
				this.state = 205;
				this.stringUnquoted();
				}
				}
				this.state = 210;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}
	// @RuleVersion(0)
	public string_(): StringContext {
		let localctx: StringContext = new StringContext(this, this._ctx, this.state);
		this.enterRule(localctx, 42, RMFQueryParser.RULE_string);
		try {
			this.state = 213;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 1:
			case 2:
			case 3:
			case 4:
			case 5:
			case 6:
			case 7:
			case 8:
			case 9:
			case 10:
			case 11:
			case 12:
			case 13:
			case 14:
			case 15:
			case 16:
			case 18:
			case 20:
			case 21:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 211;
				this.stringDotted();
				}
				break;
			case 22:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 212;
				this.match(RMFQueryParser.STRING_QUOTED);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (re) {
			if (re instanceof RecognitionException) {
				localctx.exception = re;
				this._errHandler.reportError(this, re);
				this._errHandler.recover(this, re);
			} else {
				throw re;
			}
		}
		finally {
			this.exitRule();
		}
		return localctx;
	}

	public static readonly _serializedATN: number[] = [4,1,29,216,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,1,0,5,0,46,8,0,10,0,12,0,49,
	9,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,3,0,58,8,0,1,0,1,0,1,0,5,0,63,8,0,10,0,
	12,0,66,9,0,1,0,3,0,69,8,0,1,0,5,0,72,8,0,10,0,12,0,75,9,0,1,0,1,0,1,1,
	1,1,1,2,1,2,5,2,83,8,2,10,2,12,2,86,9,2,1,2,1,2,5,2,90,8,2,10,2,12,2,93,
	9,2,1,2,1,2,5,2,97,8,2,10,2,12,2,100,9,2,1,2,1,2,5,2,104,8,2,10,2,12,2,
	107,9,2,5,2,109,8,2,10,2,12,2,112,9,2,1,2,1,2,1,3,1,3,1,3,1,3,3,3,120,8,
	3,1,4,1,4,1,4,1,4,1,5,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,7,1,7,1,7,5,7,137,8,
	7,10,7,12,7,140,9,7,1,8,1,8,1,8,1,8,1,8,1,8,3,8,148,8,8,1,9,1,9,1,9,1,9,
	1,10,1,10,1,10,1,10,1,11,1,11,1,11,1,11,1,12,1,12,1,12,1,12,1,13,1,13,1,
	13,1,13,1,14,1,14,1,14,1,14,1,15,1,15,1,15,1,15,1,16,3,16,179,8,16,1,16,
	1,16,3,16,183,8,16,1,16,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,4,19,194,
	8,19,11,19,12,19,195,1,19,5,19,199,8,19,10,19,12,19,202,9,19,1,20,1,20,
	1,20,5,20,207,8,20,10,20,12,20,210,9,20,1,21,1,21,3,21,214,8,21,1,21,0,
	0,22,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,0,2,1,
	0,18,19,3,0,1,16,18,18,20,21,220,0,47,1,0,0,0,2,78,1,0,0,0,4,80,1,0,0,0,
	6,119,1,0,0,0,8,121,1,0,0,0,10,125,1,0,0,0,12,129,1,0,0,0,14,133,1,0,0,
	0,16,147,1,0,0,0,18,149,1,0,0,0,20,153,1,0,0,0,22,157,1,0,0,0,24,161,1,
	0,0,0,26,165,1,0,0,0,28,169,1,0,0,0,30,173,1,0,0,0,32,178,1,0,0,0,34,187,
	1,0,0,0,36,189,1,0,0,0,38,191,1,0,0,0,40,203,1,0,0,0,42,213,1,0,0,0,44,
	46,5,29,0,0,45,44,1,0,0,0,46,49,1,0,0,0,47,45,1,0,0,0,47,48,1,0,0,0,48,
	50,1,0,0,0,49,47,1,0,0,0,50,57,5,16,0,0,51,52,5,23,0,0,52,58,5,1,0,0,53,
	54,5,23,0,0,54,58,5,2,0,0,55,56,5,23,0,0,56,58,5,3,0,0,57,51,1,0,0,0,57,
	53,1,0,0,0,57,55,1,0,0,0,57,58,1,0,0,0,58,59,1,0,0,0,59,60,5,23,0,0,60,
	64,3,2,1,0,61,63,5,29,0,0,62,61,1,0,0,0,63,66,1,0,0,0,64,62,1,0,0,0,64,
	65,1,0,0,0,65,68,1,0,0,0,66,64,1,0,0,0,67,69,3,4,2,0,68,67,1,0,0,0,68,69,
	1,0,0,0,69,73,1,0,0,0,70,72,5,29,0,0,71,70,1,0,0,0,72,75,1,0,0,0,73,71,
	1,0,0,0,73,74,1,0,0,0,74,76,1,0,0,0,75,73,1,0,0,0,76,77,5,0,0,1,77,1,1,
	0,0,0,78,79,3,38,19,0,79,3,1,0,0,0,80,84,5,26,0,0,81,83,5,29,0,0,82,81,
	1,0,0,0,83,86,1,0,0,0,84,82,1,0,0,0,84,85,1,0,0,0,85,87,1,0,0,0,86,84,1,
	0,0,0,87,91,3,6,3,0,88,90,5,29,0,0,89,88,1,0,0,0,90,93,1,0,0,0,91,89,1,
	0,0,0,91,92,1,0,0,0,92,110,1,0,0,0,93,91,1,0,0,0,94,98,5,25,0,0,95,97,5,
	29,0,0,96,95,1,0,0,0,97,100,1,0,0,0,98,96,1,0,0,0,98,99,1,0,0,0,99,101,
	1,0,0,0,100,98,1,0,0,0,101,105,3,6,3,0,102,104,5,29,0,0,103,102,1,0,0,0,
	104,107,1,0,0,0,105,103,1,0,0,0,105,106,1,0,0,0,106,109,1,0,0,0,107,105,
	1,0,0,0,108,94,1,0,0,0,109,112,1,0,0,0,110,108,1,0,0,0,110,111,1,0,0,0,
	111,113,1,0,0,0,112,110,1,0,0,0,113,114,5,27,0,0,114,5,1,0,0,0,115,120,
	3,8,4,0,116,120,3,10,5,0,117,120,3,12,6,0,118,120,3,30,15,0,119,115,1,0,
	0,0,119,116,1,0,0,0,119,117,1,0,0,0,119,118,1,0,0,0,120,7,1,0,0,0,121,122,
	5,6,0,0,122,123,5,28,0,0,123,124,3,42,21,0,124,9,1,0,0,0,125,126,5,7,0,
	0,126,127,5,28,0,0,127,128,3,42,21,0,128,11,1,0,0,0,129,130,5,8,0,0,130,
	131,5,28,0,0,131,132,3,14,7,0,132,13,1,0,0,0,133,138,3,16,8,0,134,135,5,
	24,0,0,135,137,3,16,8,0,136,134,1,0,0,0,137,140,1,0,0,0,138,136,1,0,0,0,
	138,139,1,0,0,0,139,15,1,0,0,0,140,138,1,0,0,0,141,148,3,18,9,0,142,148,
	3,20,10,0,143,148,3,22,11,0,144,148,3,24,12,0,145,148,3,26,13,0,146,148,
	3,28,14,0,147,141,1,0,0,0,147,142,1,0,0,0,147,143,1,0,0,0,147,144,1,0,0,
	0,147,145,1,0,0,0,147,146,1,0,0,0,148,17,1,0,0,0,149,150,5,9,0,0,150,151,
	5,28,0,0,151,152,3,42,21,0,152,19,1,0,0,0,153,154,5,10,0,0,154,155,5,28,
	0,0,155,156,3,34,17,0,156,21,1,0,0,0,157,158,5,11,0,0,158,159,5,28,0,0,
	159,160,3,34,17,0,160,23,1,0,0,0,161,162,5,12,0,0,162,163,5,28,0,0,163,
	164,5,18,0,0,164,25,1,0,0,0,165,166,5,13,0,0,166,167,5,28,0,0,167,168,5,
	18,0,0,168,27,1,0,0,0,169,170,5,14,0,0,170,171,5,28,0,0,171,172,5,15,0,
	0,172,29,1,0,0,0,173,174,5,4,0,0,174,175,5,28,0,0,175,176,3,32,16,0,176,
	31,1,0,0,0,177,179,3,42,21,0,178,177,1,0,0,0,178,179,1,0,0,0,179,180,1,
	0,0,0,180,182,5,25,0,0,181,183,3,42,21,0,182,181,1,0,0,0,182,183,1,0,0,
	0,183,184,1,0,0,0,184,185,5,25,0,0,185,186,5,17,0,0,186,33,1,0,0,0,187,
	188,7,0,0,0,188,35,1,0,0,0,189,190,7,1,0,0,190,37,1,0,0,0,191,200,3,36,
	18,0,192,194,5,29,0,0,193,192,1,0,0,0,194,195,1,0,0,0,195,193,1,0,0,0,195,
	196,1,0,0,0,196,197,1,0,0,0,197,199,3,36,18,0,198,193,1,0,0,0,199,202,1,
	0,0,0,200,198,1,0,0,0,200,201,1,0,0,0,201,39,1,0,0,0,202,200,1,0,0,0,203,
	208,3,36,18,0,204,205,5,23,0,0,205,207,3,36,18,0,206,204,1,0,0,0,207,210,
	1,0,0,0,208,206,1,0,0,0,208,209,1,0,0,0,209,41,1,0,0,0,210,208,1,0,0,0,
	211,214,3,40,20,0,212,214,5,22,0,0,213,211,1,0,0,0,213,212,1,0,0,0,214,
	43,1,0,0,0,19,47,57,64,68,73,84,91,98,105,110,119,138,147,178,182,195,200,
	208,213];

	private static __ATN: ATN;
	public static get _ATN(): ATN {
		if (!RMFQueryParser.__ATN) {
			RMFQueryParser.__ATN = new ATNDeserializer().deserialize(RMFQueryParser._serializedATN);
		}

		return RMFQueryParser.__ATN;
	}


	static DecisionsToDFA = RMFQueryParser._ATN.decisionToState.map( (ds: DecisionState, index: number) => new DFA(ds, index) );

}

export class QueryContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public RES_TYPE(): TerminalNode {
		return this.getToken(RMFQueryParser.RES_TYPE, 0);
	}
	public DOT_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.DOT);
	}
	public DOT(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.DOT, i);
	}
	public identifier(): IdentifierContext {
		return this.getTypedRuleContext(IdentifierContext, 0) as IdentifierContext;
	}
	public EOF(): TerminalNode {
		return this.getToken(RMFQueryParser.EOF, 0);
	}
	public WS_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.WS);
	}
	public WS(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.WS, i);
	}
	public qualifiers(): QualifiersContext {
		return this.getTypedRuleContext(QualifiersContext, 0) as QualifiersContext;
	}
	public REPORT(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT, 0);
	}
	public REPORT_CAPTION(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT_CAPTION, 0);
	}
	public REPORT_BANNER(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT_BANNER, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_query;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterQuery) {
	 		listener.enterQuery(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitQuery) {
	 		listener.exitQuery(this);
		}
	}
}


export class IdentifierContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public stringSpaced(): StringSpacedContext {
		return this.getTypedRuleContext(StringSpacedContext, 0) as StringSpacedContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_identifier;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterIdentifier) {
	 		listener.enterIdentifier(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitIdentifier) {
	 		listener.exitIdentifier(this);
		}
	}
}


export class QualifiersContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LBRACE(): TerminalNode {
		return this.getToken(RMFQueryParser.LBRACE, 0);
	}
	public qualifier_list(): QualifierContext[] {
		return this.getTypedRuleContexts(QualifierContext) as QualifierContext[];
	}
	public qualifier(i: number): QualifierContext {
		return this.getTypedRuleContext(QualifierContext, i) as QualifierContext;
	}
	public RBRACE(): TerminalNode {
		return this.getToken(RMFQueryParser.RBRACE, 0);
	}
	public WS_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.WS);
	}
	public WS(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.WS, i);
	}
	public COMMA_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.COMMA);
	}
	public COMMA(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.COMMA, i);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_qualifiers;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterQualifiers) {
	 		listener.enterQualifiers(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitQualifiers) {
	 		listener.exitQualifiers(this);
		}
	}
}


export class QualifierContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public ulq(): UlqContext {
		return this.getTypedRuleContext(UlqContext, 0) as UlqContext;
	}
	public name(): NameContext {
		return this.getTypedRuleContext(NameContext, 0) as NameContext;
	}
	public filter(): FilterContext {
		return this.getTypedRuleContext(FilterContext, 0) as FilterContext;
	}
	public workscope(): WorkscopeContext {
		return this.getTypedRuleContext(WorkscopeContext, 0) as WorkscopeContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_qualifier;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterQualifier) {
	 		listener.enterQualifier(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitQualifier) {
	 		listener.exitQualifier(this);
		}
	}
}


export class UlqContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public ULQ(): TerminalNode {
		return this.getToken(RMFQueryParser.ULQ, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public string_(): StringContext {
		return this.getTypedRuleContext(StringContext, 0) as StringContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_ulq;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterUlq) {
	 		listener.enterUlq(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitUlq) {
	 		listener.exitUlq(this);
		}
	}
}


export class NameContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public NAME(): TerminalNode {
		return this.getToken(RMFQueryParser.NAME, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public string_(): StringContext {
		return this.getTypedRuleContext(StringContext, 0) as StringContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_name;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterName) {
	 		listener.enterName(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitName) {
	 		listener.exitName(this);
		}
	}
}


export class FilterContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public FILTER(): TerminalNode {
		return this.getToken(RMFQueryParser.FILTER, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public filterValue(): FilterValueContext {
		return this.getTypedRuleContext(FilterValueContext, 0) as FilterValueContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_filter;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterFilter) {
	 		listener.enterFilter(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitFilter) {
	 		listener.exitFilter(this);
		}
	}
}


export class FilterValueContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public filterItem_list(): FilterItemContext[] {
		return this.getTypedRuleContexts(FilterItemContext) as FilterItemContext[];
	}
	public filterItem(i: number): FilterItemContext {
		return this.getTypedRuleContext(FilterItemContext, i) as FilterItemContext;
	}
	public SEMI_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.SEMI);
	}
	public SEMI(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.SEMI, i);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_filterValue;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterFilterValue) {
	 		listener.enterFilterValue(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitFilterValue) {
	 		listener.exitFilterValue(this);
		}
	}
}


export class FilterItemContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public pat(): PatContext {
		return this.getTypedRuleContext(PatContext, 0) as PatContext;
	}
	public lb(): LbContext {
		return this.getTypedRuleContext(LbContext, 0) as LbContext;
	}
	public ub(): UbContext {
		return this.getTypedRuleContext(UbContext, 0) as UbContext;
	}
	public hi(): HiContext {
		return this.getTypedRuleContext(HiContext, 0) as HiContext;
	}
	public lo(): LoContext {
		return this.getTypedRuleContext(LoContext, 0) as LoContext;
	}
	public ord(): OrdContext {
		return this.getTypedRuleContext(OrdContext, 0) as OrdContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_filterItem;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterFilterItem) {
	 		listener.enterFilterItem(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitFilterItem) {
	 		listener.exitFilterItem(this);
		}
	}
}


export class PatContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public PAT(): TerminalNode {
		return this.getToken(RMFQueryParser.PAT, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public string_(): StringContext {
		return this.getTypedRuleContext(StringContext, 0) as StringContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_pat;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterPat) {
	 		listener.enterPat(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitPat) {
	 		listener.exitPat(this);
		}
	}
}


export class LbContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LB(): TerminalNode {
		return this.getToken(RMFQueryParser.LB, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public number_(): NumberContext {
		return this.getTypedRuleContext(NumberContext, 0) as NumberContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_lb;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterLb) {
	 		listener.enterLb(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitLb) {
	 		listener.exitLb(this);
		}
	}
}


export class UbContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public UB(): TerminalNode {
		return this.getToken(RMFQueryParser.UB, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public number_(): NumberContext {
		return this.getTypedRuleContext(NumberContext, 0) as NumberContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_ub;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterUb) {
	 		listener.enterUb(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitUb) {
	 		listener.exitUb(this);
		}
	}
}


export class HiContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public HI(): TerminalNode {
		return this.getToken(RMFQueryParser.HI, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public INTEGER(): TerminalNode {
		return this.getToken(RMFQueryParser.INTEGER, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_hi;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterHi) {
	 		listener.enterHi(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitHi) {
	 		listener.exitHi(this);
		}
	}
}


export class LoContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public LO(): TerminalNode {
		return this.getToken(RMFQueryParser.LO, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public INTEGER(): TerminalNode {
		return this.getToken(RMFQueryParser.INTEGER, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_lo;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterLo) {
	 		listener.enterLo(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitLo) {
	 		listener.exitLo(this);
		}
	}
}


export class OrdContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public ORD(): TerminalNode {
		return this.getToken(RMFQueryParser.ORD, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public ORD_OPTION(): TerminalNode {
		return this.getToken(RMFQueryParser.ORD_OPTION, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_ord;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterOrd) {
	 		listener.enterOrd(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitOrd) {
	 		listener.exitOrd(this);
		}
	}
}


export class WorkscopeContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public WORKSCOPE(): TerminalNode {
		return this.getToken(RMFQueryParser.WORKSCOPE, 0);
	}
	public EQUAL(): TerminalNode {
		return this.getToken(RMFQueryParser.EQUAL, 0);
	}
	public workscopeValue(): WorkscopeValueContext {
		return this.getTypedRuleContext(WorkscopeValueContext, 0) as WorkscopeValueContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_workscope;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterWorkscope) {
	 		listener.enterWorkscope(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitWorkscope) {
	 		listener.exitWorkscope(this);
		}
	}
}


export class WorkscopeValueContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public COMMA_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.COMMA);
	}
	public COMMA(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.COMMA, i);
	}
	public WORKSCOPE_TYPE(): TerminalNode {
		return this.getToken(RMFQueryParser.WORKSCOPE_TYPE, 0);
	}
	public string__list(): StringContext[] {
		return this.getTypedRuleContexts(StringContext) as StringContext[];
	}
	public string_(i: number): StringContext {
		return this.getTypedRuleContext(StringContext, i) as StringContext;
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_workscopeValue;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterWorkscopeValue) {
	 		listener.enterWorkscopeValue(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitWorkscopeValue) {
	 		listener.exitWorkscopeValue(this);
		}
	}
}


export class NumberContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public INTEGER(): TerminalNode {
		return this.getToken(RMFQueryParser.INTEGER, 0);
	}
	public DECIMAL(): TerminalNode {
		return this.getToken(RMFQueryParser.DECIMAL, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_number;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterNumber) {
	 		listener.enterNumber(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitNumber) {
	 		listener.exitNumber(this);
		}
	}
}


export class StringUnquotedContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public IDENTIFIER(): TerminalNode {
		return this.getToken(RMFQueryParser.IDENTIFIER, 0);
	}
	public RES_TYPE(): TerminalNode {
		return this.getToken(RMFQueryParser.RES_TYPE, 0);
	}
	public REPORT(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT, 0);
	}
	public REPORT_CAPTION(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT_CAPTION, 0);
	}
	public REPORT_BANNER(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT_BANNER, 0);
	}
	public WORKSCOPE(): TerminalNode {
		return this.getToken(RMFQueryParser.WORKSCOPE, 0);
	}
	public RANGE(): TerminalNode {
		return this.getToken(RMFQueryParser.RANGE, 0);
	}
	public ULQ(): TerminalNode {
		return this.getToken(RMFQueryParser.ULQ, 0);
	}
	public NAME(): TerminalNode {
		return this.getToken(RMFQueryParser.NAME, 0);
	}
	public FILTER(): TerminalNode {
		return this.getToken(RMFQueryParser.FILTER, 0);
	}
	public PAT(): TerminalNode {
		return this.getToken(RMFQueryParser.PAT, 0);
	}
	public LB(): TerminalNode {
		return this.getToken(RMFQueryParser.LB, 0);
	}
	public UB(): TerminalNode {
		return this.getToken(RMFQueryParser.UB, 0);
	}
	public HI(): TerminalNode {
		return this.getToken(RMFQueryParser.HI, 0);
	}
	public LO(): TerminalNode {
		return this.getToken(RMFQueryParser.LO, 0);
	}
	public ORD(): TerminalNode {
		return this.getToken(RMFQueryParser.ORD, 0);
	}
	public ORD_OPTION(): TerminalNode {
		return this.getToken(RMFQueryParser.ORD_OPTION, 0);
	}
	public INTEGER(): TerminalNode {
		return this.getToken(RMFQueryParser.INTEGER, 0);
	}
	public STRING_UNQUOTED(): TerminalNode {
		return this.getToken(RMFQueryParser.STRING_UNQUOTED, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_stringUnquoted;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterStringUnquoted) {
	 		listener.enterStringUnquoted(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitStringUnquoted) {
	 		listener.exitStringUnquoted(this);
		}
	}
}


export class StringSpacedContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public stringUnquoted_list(): StringUnquotedContext[] {
		return this.getTypedRuleContexts(StringUnquotedContext) as StringUnquotedContext[];
	}
	public stringUnquoted(i: number): StringUnquotedContext {
		return this.getTypedRuleContext(StringUnquotedContext, i) as StringUnquotedContext;
	}
	public WS_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.WS);
	}
	public WS(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.WS, i);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_stringSpaced;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterStringSpaced) {
	 		listener.enterStringSpaced(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitStringSpaced) {
	 		listener.exitStringSpaced(this);
		}
	}
}


export class StringDottedContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public stringUnquoted_list(): StringUnquotedContext[] {
		return this.getTypedRuleContexts(StringUnquotedContext) as StringUnquotedContext[];
	}
	public stringUnquoted(i: number): StringUnquotedContext {
		return this.getTypedRuleContext(StringUnquotedContext, i) as StringUnquotedContext;
	}
	public DOT_list(): TerminalNode[] {
	    	return this.getTokens(RMFQueryParser.DOT);
	}
	public DOT(i: number): TerminalNode {
		return this.getToken(RMFQueryParser.DOT, i);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_stringDotted;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterStringDotted) {
	 		listener.enterStringDotted(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitStringDotted) {
	 		listener.exitStringDotted(this);
		}
	}
}


export class StringContext extends ParserRuleContext {
	constructor(parser?: RMFQueryParser, parent?: ParserRuleContext, invokingState?: number) {
		super(parent, invokingState);
    	this.parser = parser;
	}
	public stringDotted(): StringDottedContext {
		return this.getTypedRuleContext(StringDottedContext, 0) as StringDottedContext;
	}
	public STRING_QUOTED(): TerminalNode {
		return this.getToken(RMFQueryParser.STRING_QUOTED, 0);
	}
    public get ruleIndex(): number {
    	return RMFQueryParser.RULE_string;
	}
	public enterRule(listener: RMFQueryListener): void {
	    if(listener.enterString) {
	 		listener.enterString(this);
		}
	}
	public exitRule(listener: RMFQueryListener): void {
	    if(listener.exitString) {
	 		listener.exitString(this);
		}
	}
}

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
	public static readonly WORKSCOPE = 2;
	public static readonly RANGE = 3;
	public static readonly ULQ = 4;
	public static readonly NAME = 5;
	public static readonly FILTER = 6;
	public static readonly PAT = 7;
	public static readonly LB = 8;
	public static readonly UB = 9;
	public static readonly HI = 10;
	public static readonly LO = 11;
	public static readonly ORD = 12;
	public static readonly ORD_OPTION = 13;
	public static readonly RES_TYPE = 14;
	public static readonly WORKSCOPE_TYPE = 15;
	public static readonly INTEGER = 16;
	public static readonly DECIMAL = 17;
	public static readonly IDENTIFIER = 18;
	public static readonly STRING_UNQUOTED = 19;
	public static readonly STRING_QUOTED = 20;
	public static readonly DOT = 21;
	public static readonly SEMI = 22;
	public static readonly COMMA = 23;
	public static readonly LBRACE = 24;
	public static readonly RBRACE = 25;
	public static readonly EQUAL = 26;
	public static readonly WS = 27;
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
                                                            null, "'.'", 
                                                            "';'", "','", 
                                                            "'{'", "'}'", 
                                                            "'='" ];
	public static readonly symbolicNames: (string | null)[] = [ null, "REPORT", 
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
			while (_la===27) {
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
			this.state = 53;
			this._errHandler.sync(this);
			switch ( this._interp.adaptivePredict(this._input, 1, this._ctx) ) {
			case 1:
				{
				this.state = 51;
				this.match(RMFQueryParser.DOT);
				this.state = 52;
				this.match(RMFQueryParser.REPORT);
				}
				break;
			}
			this.state = 55;
			this.match(RMFQueryParser.DOT);
			this.state = 56;
			this.identifier();
			this.state = 60;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 2, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 57;
					this.match(RMFQueryParser.WS);
					}
					}
				}
				this.state = 62;
				this._errHandler.sync(this);
				_alt = this._interp.adaptivePredict(this._input, 2, this._ctx);
			}
			this.state = 64;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if (_la===24) {
				{
				this.state = 63;
				this.qualifiers();
				}
			}

			this.state = 69;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===27) {
				{
				{
				this.state = 66;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 71;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 72;
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
			this.state = 74;
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
			this.state = 76;
			this.match(RMFQueryParser.LBRACE);
			this.state = 80;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===27) {
				{
				{
				this.state = 77;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 82;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 83;
			this.qualifier();
			this.state = 87;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===27) {
				{
				{
				this.state = 84;
				this.match(RMFQueryParser.WS);
				}
				}
				this.state = 89;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 106;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===23) {
				{
				{
				this.state = 90;
				this.match(RMFQueryParser.COMMA);
				this.state = 94;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				while (_la===27) {
					{
					{
					this.state = 91;
					this.match(RMFQueryParser.WS);
					}
					}
					this.state = 96;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				}
				this.state = 97;
				this.qualifier();
				this.state = 101;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
				while (_la===27) {
					{
					{
					this.state = 98;
					this.match(RMFQueryParser.WS);
					}
					}
					this.state = 103;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
				}
				}
				}
				this.state = 108;
				this._errHandler.sync(this);
				_la = this._input.LA(1);
			}
			this.state = 109;
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
			this.state = 115;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 4:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 111;
				this.ulq();
				}
				break;
			case 5:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 112;
				this.name();
				}
				break;
			case 6:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 113;
				this.filter();
				}
				break;
			case 2:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 114;
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
			this.state = 117;
			this.match(RMFQueryParser.ULQ);
			this.state = 118;
			this.match(RMFQueryParser.EQUAL);
			this.state = 119;
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
			this.state = 121;
			this.match(RMFQueryParser.NAME);
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
	public filter(): FilterContext {
		let localctx: FilterContext = new FilterContext(this, this._ctx, this.state);
		this.enterRule(localctx, 12, RMFQueryParser.RULE_filter);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 125;
			this.match(RMFQueryParser.FILTER);
			this.state = 126;
			this.match(RMFQueryParser.EQUAL);
			this.state = 127;
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
			this.state = 129;
			this.filterItem();
			this.state = 134;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===22) {
				{
				{
				this.state = 130;
				this.match(RMFQueryParser.SEMI);
				this.state = 131;
				this.filterItem();
				}
				}
				this.state = 136;
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
			this.state = 143;
			this._errHandler.sync(this);
			switch (this._input.LA(1)) {
			case 7:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 137;
				this.pat();
				}
				break;
			case 8:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 138;
				this.lb();
				}
				break;
			case 9:
				this.enterOuterAlt(localctx, 3);
				{
				this.state = 139;
				this.ub();
				}
				break;
			case 10:
				this.enterOuterAlt(localctx, 4);
				{
				this.state = 140;
				this.hi();
				}
				break;
			case 11:
				this.enterOuterAlt(localctx, 5);
				{
				this.state = 141;
				this.lo();
				}
				break;
			case 12:
				this.enterOuterAlt(localctx, 6);
				{
				this.state = 142;
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
			this.state = 145;
			this.match(RMFQueryParser.PAT);
			this.state = 146;
			this.match(RMFQueryParser.EQUAL);
			this.state = 147;
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
			this.state = 149;
			this.match(RMFQueryParser.LB);
			this.state = 150;
			this.match(RMFQueryParser.EQUAL);
			this.state = 151;
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
			this.state = 153;
			this.match(RMFQueryParser.UB);
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
	public hi(): HiContext {
		let localctx: HiContext = new HiContext(this, this._ctx, this.state);
		this.enterRule(localctx, 24, RMFQueryParser.RULE_hi);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 157;
			this.match(RMFQueryParser.HI);
			this.state = 158;
			this.match(RMFQueryParser.EQUAL);
			this.state = 159;
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
			this.state = 161;
			this.match(RMFQueryParser.LO);
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
	public ord(): OrdContext {
		let localctx: OrdContext = new OrdContext(this, this._ctx, this.state);
		this.enterRule(localctx, 28, RMFQueryParser.RULE_ord);
		try {
			this.enterOuterAlt(localctx, 1);
			{
			this.state = 165;
			this.match(RMFQueryParser.ORD);
			this.state = 166;
			this.match(RMFQueryParser.EQUAL);
			this.state = 167;
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
			this.state = 169;
			this.match(RMFQueryParser.WORKSCOPE);
			this.state = 170;
			this.match(RMFQueryParser.EQUAL);
			this.state = 171;
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
			this.state = 174;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 1933310) !== 0)) {
				{
				this.state = 173;
				this.string_();
				}
			}

			this.state = 176;
			this.match(RMFQueryParser.COMMA);
			this.state = 178;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			if ((((_la) & ~0x1F) === 0 && ((1 << _la) & 1933310) !== 0)) {
				{
				this.state = 177;
				this.string_();
				}
			}

			this.state = 180;
			this.match(RMFQueryParser.COMMA);
			this.state = 181;
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
			this.state = 183;
			_la = this._input.LA(1);
			if(!(_la===16 || _la===17)) {
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
			this.state = 185;
			_la = this._input.LA(1);
			if(!((((_la) & ~0x1F) === 0 && ((1 << _la) & 884734) !== 0))) {
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
			this.state = 187;
			this.stringUnquoted();
			this.state = 196;
			this._errHandler.sync(this);
			_alt = this._interp.adaptivePredict(this._input, 16, this._ctx);
			while (_alt !== 2 && _alt !== ATN.INVALID_ALT_NUMBER) {
				if (_alt === 1) {
					{
					{
					this.state = 189;
					this._errHandler.sync(this);
					_la = this._input.LA(1);
					do {
						{
						{
						this.state = 188;
						this.match(RMFQueryParser.WS);
						}
						}
						this.state = 191;
						this._errHandler.sync(this);
						_la = this._input.LA(1);
					} while (_la===27);
					this.state = 193;
					this.stringUnquoted();
					}
					}
				}
				this.state = 198;
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
			this.state = 199;
			this.stringUnquoted();
			this.state = 204;
			this._errHandler.sync(this);
			_la = this._input.LA(1);
			while (_la===21) {
				{
				{
				this.state = 200;
				this.match(RMFQueryParser.DOT);
				this.state = 201;
				this.stringUnquoted();
				}
				}
				this.state = 206;
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
			this.state = 209;
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
			case 16:
			case 18:
			case 19:
				this.enterOuterAlt(localctx, 1);
				{
				this.state = 207;
				this.stringDotted();
				}
				break;
			case 20:
				this.enterOuterAlt(localctx, 2);
				{
				this.state = 208;
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

	public static readonly _serializedATN: number[] = [4,1,27,212,2,0,7,0,2,
	1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,
	10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,
	7,17,2,18,7,18,2,19,7,19,2,20,7,20,2,21,7,21,1,0,5,0,46,8,0,10,0,12,0,49,
	9,0,1,0,1,0,1,0,3,0,54,8,0,1,0,1,0,1,0,5,0,59,8,0,10,0,12,0,62,9,0,1,0,
	3,0,65,8,0,1,0,5,0,68,8,0,10,0,12,0,71,9,0,1,0,1,0,1,1,1,1,1,2,1,2,5,2,
	79,8,2,10,2,12,2,82,9,2,1,2,1,2,5,2,86,8,2,10,2,12,2,89,9,2,1,2,1,2,5,2,
	93,8,2,10,2,12,2,96,9,2,1,2,1,2,5,2,100,8,2,10,2,12,2,103,9,2,5,2,105,8,
	2,10,2,12,2,108,9,2,1,2,1,2,1,3,1,3,1,3,1,3,3,3,116,8,3,1,4,1,4,1,4,1,4,
	1,5,1,5,1,5,1,5,1,6,1,6,1,6,1,6,1,7,1,7,1,7,5,7,133,8,7,10,7,12,7,136,9,
	7,1,8,1,8,1,8,1,8,1,8,1,8,3,8,144,8,8,1,9,1,9,1,9,1,9,1,10,1,10,1,10,1,
	10,1,11,1,11,1,11,1,11,1,12,1,12,1,12,1,12,1,13,1,13,1,13,1,13,1,14,1,14,
	1,14,1,14,1,15,1,15,1,15,1,15,1,16,3,16,175,8,16,1,16,1,16,3,16,179,8,16,
	1,16,1,16,1,16,1,17,1,17,1,18,1,18,1,19,1,19,4,19,190,8,19,11,19,12,19,
	191,1,19,5,19,195,8,19,10,19,12,19,198,9,19,1,20,1,20,1,20,5,20,203,8,20,
	10,20,12,20,206,9,20,1,21,1,21,3,21,210,8,21,1,21,0,0,22,0,2,4,6,8,10,12,
	14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,0,2,1,0,16,17,3,0,1,14,16,
	16,18,19,214,0,47,1,0,0,0,2,74,1,0,0,0,4,76,1,0,0,0,6,115,1,0,0,0,8,117,
	1,0,0,0,10,121,1,0,0,0,12,125,1,0,0,0,14,129,1,0,0,0,16,143,1,0,0,0,18,
	145,1,0,0,0,20,149,1,0,0,0,22,153,1,0,0,0,24,157,1,0,0,0,26,161,1,0,0,0,
	28,165,1,0,0,0,30,169,1,0,0,0,32,174,1,0,0,0,34,183,1,0,0,0,36,185,1,0,
	0,0,38,187,1,0,0,0,40,199,1,0,0,0,42,209,1,0,0,0,44,46,5,27,0,0,45,44,1,
	0,0,0,46,49,1,0,0,0,47,45,1,0,0,0,47,48,1,0,0,0,48,50,1,0,0,0,49,47,1,0,
	0,0,50,53,5,14,0,0,51,52,5,21,0,0,52,54,5,1,0,0,53,51,1,0,0,0,53,54,1,0,
	0,0,54,55,1,0,0,0,55,56,5,21,0,0,56,60,3,2,1,0,57,59,5,27,0,0,58,57,1,0,
	0,0,59,62,1,0,0,0,60,58,1,0,0,0,60,61,1,0,0,0,61,64,1,0,0,0,62,60,1,0,0,
	0,63,65,3,4,2,0,64,63,1,0,0,0,64,65,1,0,0,0,65,69,1,0,0,0,66,68,5,27,0,
	0,67,66,1,0,0,0,68,71,1,0,0,0,69,67,1,0,0,0,69,70,1,0,0,0,70,72,1,0,0,0,
	71,69,1,0,0,0,72,73,5,0,0,1,73,1,1,0,0,0,74,75,3,38,19,0,75,3,1,0,0,0,76,
	80,5,24,0,0,77,79,5,27,0,0,78,77,1,0,0,0,79,82,1,0,0,0,80,78,1,0,0,0,80,
	81,1,0,0,0,81,83,1,0,0,0,82,80,1,0,0,0,83,87,3,6,3,0,84,86,5,27,0,0,85,
	84,1,0,0,0,86,89,1,0,0,0,87,85,1,0,0,0,87,88,1,0,0,0,88,106,1,0,0,0,89,
	87,1,0,0,0,90,94,5,23,0,0,91,93,5,27,0,0,92,91,1,0,0,0,93,96,1,0,0,0,94,
	92,1,0,0,0,94,95,1,0,0,0,95,97,1,0,0,0,96,94,1,0,0,0,97,101,3,6,3,0,98,
	100,5,27,0,0,99,98,1,0,0,0,100,103,1,0,0,0,101,99,1,0,0,0,101,102,1,0,0,
	0,102,105,1,0,0,0,103,101,1,0,0,0,104,90,1,0,0,0,105,108,1,0,0,0,106,104,
	1,0,0,0,106,107,1,0,0,0,107,109,1,0,0,0,108,106,1,0,0,0,109,110,5,25,0,
	0,110,5,1,0,0,0,111,116,3,8,4,0,112,116,3,10,5,0,113,116,3,12,6,0,114,116,
	3,30,15,0,115,111,1,0,0,0,115,112,1,0,0,0,115,113,1,0,0,0,115,114,1,0,0,
	0,116,7,1,0,0,0,117,118,5,4,0,0,118,119,5,26,0,0,119,120,3,42,21,0,120,
	9,1,0,0,0,121,122,5,5,0,0,122,123,5,26,0,0,123,124,3,42,21,0,124,11,1,0,
	0,0,125,126,5,6,0,0,126,127,5,26,0,0,127,128,3,14,7,0,128,13,1,0,0,0,129,
	134,3,16,8,0,130,131,5,22,0,0,131,133,3,16,8,0,132,130,1,0,0,0,133,136,
	1,0,0,0,134,132,1,0,0,0,134,135,1,0,0,0,135,15,1,0,0,0,136,134,1,0,0,0,
	137,144,3,18,9,0,138,144,3,20,10,0,139,144,3,22,11,0,140,144,3,24,12,0,
	141,144,3,26,13,0,142,144,3,28,14,0,143,137,1,0,0,0,143,138,1,0,0,0,143,
	139,1,0,0,0,143,140,1,0,0,0,143,141,1,0,0,0,143,142,1,0,0,0,144,17,1,0,
	0,0,145,146,5,7,0,0,146,147,5,26,0,0,147,148,3,42,21,0,148,19,1,0,0,0,149,
	150,5,8,0,0,150,151,5,26,0,0,151,152,3,34,17,0,152,21,1,0,0,0,153,154,5,
	9,0,0,154,155,5,26,0,0,155,156,3,34,17,0,156,23,1,0,0,0,157,158,5,10,0,
	0,158,159,5,26,0,0,159,160,5,16,0,0,160,25,1,0,0,0,161,162,5,11,0,0,162,
	163,5,26,0,0,163,164,5,16,0,0,164,27,1,0,0,0,165,166,5,12,0,0,166,167,5,
	26,0,0,167,168,5,13,0,0,168,29,1,0,0,0,169,170,5,2,0,0,170,171,5,26,0,0,
	171,172,3,32,16,0,172,31,1,0,0,0,173,175,3,42,21,0,174,173,1,0,0,0,174,
	175,1,0,0,0,175,176,1,0,0,0,176,178,5,23,0,0,177,179,3,42,21,0,178,177,
	1,0,0,0,178,179,1,0,0,0,179,180,1,0,0,0,180,181,5,23,0,0,181,182,5,15,0,
	0,182,33,1,0,0,0,183,184,7,0,0,0,184,35,1,0,0,0,185,186,7,1,0,0,186,37,
	1,0,0,0,187,196,3,36,18,0,188,190,5,27,0,0,189,188,1,0,0,0,190,191,1,0,
	0,0,191,189,1,0,0,0,191,192,1,0,0,0,192,193,1,0,0,0,193,195,3,36,18,0,194,
	189,1,0,0,0,195,198,1,0,0,0,196,194,1,0,0,0,196,197,1,0,0,0,197,39,1,0,
	0,0,198,196,1,0,0,0,199,204,3,36,18,0,200,201,5,21,0,0,201,203,3,36,18,
	0,202,200,1,0,0,0,203,206,1,0,0,0,204,202,1,0,0,0,204,205,1,0,0,0,205,41,
	1,0,0,0,206,204,1,0,0,0,207,210,3,40,20,0,208,210,5,20,0,0,209,207,1,0,
	0,0,209,208,1,0,0,0,210,43,1,0,0,0,19,47,53,60,64,69,80,87,94,101,106,115,
	134,143,174,178,191,196,204,209];

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
	public REPORT(): TerminalNode {
		return this.getToken(RMFQueryParser.REPORT, 0);
	}
	public qualifiers(): QualifiersContext {
		return this.getTypedRuleContext(QualifiersContext, 0) as QualifiersContext;
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

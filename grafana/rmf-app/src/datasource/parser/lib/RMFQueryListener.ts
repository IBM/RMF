// Generated from ./grammar/RMFQuery.g4 by ANTLR 4.13.2

import {ParseTreeListener} from "antlr4";


import { QueryContext } from "./RMFQueryParser.js";
import { IdentifierContext } from "./RMFQueryParser.js";
import { QualifiersContext } from "./RMFQueryParser.js";
import { QualifierContext } from "./RMFQueryParser.js";
import { UlqContext } from "./RMFQueryParser.js";
import { NameContext } from "./RMFQueryParser.js";
import { FilterContext } from "./RMFQueryParser.js";
import { FilterValueContext } from "./RMFQueryParser.js";
import { FilterItemContext } from "./RMFQueryParser.js";
import { PatContext } from "./RMFQueryParser.js";
import { LbContext } from "./RMFQueryParser.js";
import { UbContext } from "./RMFQueryParser.js";
import { HiContext } from "./RMFQueryParser.js";
import { LoContext } from "./RMFQueryParser.js";
import { OrdContext } from "./RMFQueryParser.js";
import { WorkscopeContext } from "./RMFQueryParser.js";
import { WorkscopeValueContext } from "./RMFQueryParser.js";
import { NumberContext } from "./RMFQueryParser.js";
import { StringUnquotedContext } from "./RMFQueryParser.js";
import { StringSpacedContext } from "./RMFQueryParser.js";
import { StringDottedContext } from "./RMFQueryParser.js";
import { StringContext } from "./RMFQueryParser.js";


/**
 * This interface defines a complete listener for a parse tree produced by
 * `RMFQueryParser`.
 */
export default class RMFQueryListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by `RMFQueryParser.query`.
	 * @param ctx the parse tree
	 */
	enterQuery?: (ctx: QueryContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.query`.
	 * @param ctx the parse tree
	 */
	exitQuery?: (ctx: QueryContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.identifier`.
	 * @param ctx the parse tree
	 */
	enterIdentifier?: (ctx: IdentifierContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.identifier`.
	 * @param ctx the parse tree
	 */
	exitIdentifier?: (ctx: IdentifierContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.qualifiers`.
	 * @param ctx the parse tree
	 */
	enterQualifiers?: (ctx: QualifiersContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.qualifiers`.
	 * @param ctx the parse tree
	 */
	exitQualifiers?: (ctx: QualifiersContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.qualifier`.
	 * @param ctx the parse tree
	 */
	enterQualifier?: (ctx: QualifierContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.qualifier`.
	 * @param ctx the parse tree
	 */
	exitQualifier?: (ctx: QualifierContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.ulq`.
	 * @param ctx the parse tree
	 */
	enterUlq?: (ctx: UlqContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.ulq`.
	 * @param ctx the parse tree
	 */
	exitUlq?: (ctx: UlqContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.name`.
	 * @param ctx the parse tree
	 */
	enterName?: (ctx: NameContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.name`.
	 * @param ctx the parse tree
	 */
	exitName?: (ctx: NameContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.filter`.
	 * @param ctx the parse tree
	 */
	enterFilter?: (ctx: FilterContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.filter`.
	 * @param ctx the parse tree
	 */
	exitFilter?: (ctx: FilterContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.filterValue`.
	 * @param ctx the parse tree
	 */
	enterFilterValue?: (ctx: FilterValueContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.filterValue`.
	 * @param ctx the parse tree
	 */
	exitFilterValue?: (ctx: FilterValueContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.filterItem`.
	 * @param ctx the parse tree
	 */
	enterFilterItem?: (ctx: FilterItemContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.filterItem`.
	 * @param ctx the parse tree
	 */
	exitFilterItem?: (ctx: FilterItemContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.pat`.
	 * @param ctx the parse tree
	 */
	enterPat?: (ctx: PatContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.pat`.
	 * @param ctx the parse tree
	 */
	exitPat?: (ctx: PatContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.lb`.
	 * @param ctx the parse tree
	 */
	enterLb?: (ctx: LbContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.lb`.
	 * @param ctx the parse tree
	 */
	exitLb?: (ctx: LbContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.ub`.
	 * @param ctx the parse tree
	 */
	enterUb?: (ctx: UbContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.ub`.
	 * @param ctx the parse tree
	 */
	exitUb?: (ctx: UbContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.hi`.
	 * @param ctx the parse tree
	 */
	enterHi?: (ctx: HiContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.hi`.
	 * @param ctx the parse tree
	 */
	exitHi?: (ctx: HiContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.lo`.
	 * @param ctx the parse tree
	 */
	enterLo?: (ctx: LoContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.lo`.
	 * @param ctx the parse tree
	 */
	exitLo?: (ctx: LoContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.ord`.
	 * @param ctx the parse tree
	 */
	enterOrd?: (ctx: OrdContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.ord`.
	 * @param ctx the parse tree
	 */
	exitOrd?: (ctx: OrdContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.workscope`.
	 * @param ctx the parse tree
	 */
	enterWorkscope?: (ctx: WorkscopeContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.workscope`.
	 * @param ctx the parse tree
	 */
	exitWorkscope?: (ctx: WorkscopeContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.workscopeValue`.
	 * @param ctx the parse tree
	 */
	enterWorkscopeValue?: (ctx: WorkscopeValueContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.workscopeValue`.
	 * @param ctx the parse tree
	 */
	exitWorkscopeValue?: (ctx: WorkscopeValueContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.number`.
	 * @param ctx the parse tree
	 */
	enterNumber?: (ctx: NumberContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.number`.
	 * @param ctx the parse tree
	 */
	exitNumber?: (ctx: NumberContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.stringUnquoted`.
	 * @param ctx the parse tree
	 */
	enterStringUnquoted?: (ctx: StringUnquotedContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.stringUnquoted`.
	 * @param ctx the parse tree
	 */
	exitStringUnquoted?: (ctx: StringUnquotedContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.stringSpaced`.
	 * @param ctx the parse tree
	 */
	enterStringSpaced?: (ctx: StringSpacedContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.stringSpaced`.
	 * @param ctx the parse tree
	 */
	exitStringSpaced?: (ctx: StringSpacedContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.stringDotted`.
	 * @param ctx the parse tree
	 */
	enterStringDotted?: (ctx: StringDottedContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.stringDotted`.
	 * @param ctx the parse tree
	 */
	exitStringDotted?: (ctx: StringDottedContext) => void;
	/**
	 * Enter a parse tree produced by `RMFQueryParser.string`.
	 * @param ctx the parse tree
	 */
	enterString?: (ctx: StringContext) => void;
	/**
	 * Exit a parse tree produced by `RMFQueryParser.string`.
	 * @param ctx the parse tree
	 */
	exitString?: (ctx: StringContext) => void;
}


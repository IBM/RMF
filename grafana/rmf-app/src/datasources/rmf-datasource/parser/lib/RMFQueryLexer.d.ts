import {InputStream, Lexer} from 'antlr4';

export default  RMFQueryLexer;

export declare class RMFQueryLexer extends Lexer {
    readonly channelNames: string[];
    readonly modeNames: string[];
    readonly literalNames: string[];
    readonly symbolicNames: string[];
    readonly ruleNames: string[];
    readonly grammarFileName: string;

    constructor(input: InputStream);
}

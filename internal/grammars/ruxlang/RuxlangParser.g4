/*
 The MIT Licence
 Copyright (c) 2016 Jeremy Wilkins
 All rights reserved.

 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included in all
 copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 SOFTWARE.
*/
/**
 * A Rux language parser grammar for ANTLR 4 derived from the Rux Language Specification
 * https://ruxlang.org/
 *
 */
parser grammar RuxlangParser;

options {
    tokenVocab = RuxlangLexer;
}

// Must be at least one top level declaration, the rest is optional.
sourceFile
    : ( packageClause )? ( importDecl )* ( topLevelDecl )+
    ;

// If the package clause is omitted, and the package name isn't defined
// in any other file in this folder or in a separate project file, the
// containing folder names (dot separated) leading from the project to the
// path the file is contained in will be used as the package name, but if
// there is no project file the name of the containing folder will become
// the package name. 
packageClause
    : STMTEND* PACKAGE IDENTIFIER STMTEND
    ;

importDecl
    : STMTEND* IMPORT ( importSpec | ( BLOCK_START ( importSpec )+ BLOCK_END ) )
    ;

importSpec
    : packageIdent
    | ( ( DOT | IDENTIFIER ) ASSIGN packageIdent )
    ;

packageIdent
    : ( EXTENDED_IDENTIFIER | QUALIFIED_IDENTIFIER | IDENTIFIER )
    ;

binaryOp
    : ( OR | AND | EQ | NEQ | LT | LTE | GT | GTE | ADD | SUB | BIT_OR | CARET | MUL | DIV | MOD | SHL | SHR | AMP | BIT_CLEAR )
    ;

// A top level declaration supports running statements at the top level
// although this is discouraged by linters except for rxsh syntax usage
topLevelDecl
    : declaration
    | functionDecl
    | methodDecl
    | topLevelStatement
    ;

declaration
    : constDecl
    | typeDecl
    | varDecl
    ;

constDecl
    : STMTEND* CONST ( constSpec | BLOCK_START ( constSpec )+ BLOCK_END )
    ;

constSpec
    : identifierList ( typeId? ASSIGN expressionList )? STMTEND
    ;

identifierList
    : IDENTIFIER ( COMMA IDENTIFIER )*
    ;

expressionList
    : expression ( COMMA expression )*
    ;

typeDecl
    : TYPE ( typeSpec | BLOCK_START ( typeSpec )+ BLOCK_END )
    ;

typeSpec
    : IDENTIFIER typeId
    ;


// Function declarations

functionDecl
    : STMTEND* FUNC IDENTIFIER ( function | signature )
    ;

function
    : signature block
    ;

methodDecl
    : FUNC receiver IDENTIFIER ( function | signature )
    ;

receiver
    : parameters
    ;

varDecl
    : VAR ( varSpec | BLOCK_START ( varSpec )+ BLOCK_END )
    ;

varSpec
    : identifierList ( typeId ( ASSIGN expressionList )? | ASSIGN expressionList )
    ;


block
    : BLOCK_START statementList BLOCK_END
    ;

statementList
    : ( ( statement STMTEND? )* )
    ;

topLevelStatement
    : labeledStmt
    | simpleStmt
    | runExpr
    | gotoStmt
    | ifStmt
    | switchStmt
    | selectStmt
    | forStmt
    | deferStmt
	;

statement
    : declaration
    | labeledStmt
    | simpleStmt
    | returnStmt
    | yieldStmt
    | breakStmt
    | continueStmt
    | gotoStmt
    | fallthroughStmt
    | block
    | ifStmt
    | switchStmt
    | selectStmt
    | forStmt
    | deferStmt
	;

simpleStmt
    : expressionStmt
    | sendStmt
    | incDecStmt
    | assignment
    | shortVarDecl
    ;

expressionStmt
    : expression
    ;

sendStmt
    : expression SEND expression
    ;

incDecStmt
    : expression ( INC | DEC )
    ;

assignment
    : expressionList ( ASSIGN | ASSIGN_WITH_OP ) expressionList
    ;

shortVarDecl
    : identifierList ASSIGN_NEW expressionList
    ;

labeledStmt
    : IDENTIFIER COLON statement
    ;

returnStmt
    : RETURN expressionList?
    ;

yieldStmt
    : YIELD expressionList?
    ;

breakStmt
    : BREAK IDENTIFIER?
    ;

continueStmt
    : CONTINUE IDENTIFIER?
    ;

gotoStmt
    : GOTO IDENTIFIER
    ;

fallthroughStmt
    : FALLTHROUGH
    ;

deferStmt
    : DEFER expression
    ;

ifStmt
    : IF expression STMTEND? block ( ELSE ( ifStmt | block ) )?
    ;

switchStmt
    : exprSwitchStmt | typeSwitchStmt
    ;

exprSwitchStmt
    : SWITCH expression? BLOCK_START exprCaseClause* BLOCK_END
    ;

exprCaseClause
    : exprSwitchCase COLON statementList
    ;

exprSwitchCase
    : CASE expressionList | DEFAULT
    ;

typeSwitchStmt
    : SWITCH typeSwitchGuard BLOCK_START typeCaseClause* BLOCK_END
    ;
typeSwitchGuard
    : ( IDENTIFIER ASSIGN_NEW )? primaryExpr TYPE_UNWRAP
    ;
typeCaseClause
    : typeSwitchCase COLON statementList
    ;
typeSwitchCase
    : CASE typeList | DEFAULT
    ;
typeList
    : typeId ( COMMA typeId )*
    ;

selectStmt
    : SELECT BLOCK_START commClause* BLOCK_END
    ;
commClause
    : commCase COLON statementList
    ;
commCase
    : CASE ( sendStmt | recvStmt ) | DEFAULT
    ;
recvStmt
    : ( expressionList ASSIGN | identifierList ASSIGN_NEW )? expression
    ;

forStmt
    : FOR ( expression | inClause )? STMTEND? block
    ;

inClause
    : ( expressionList | identifierList )? IN expression
    ;

runExpr
    : RUN expression
    ;

typeId
    : typeName
    | typeLit
    | PAREN_START typeId PAREN_END
    ;

typeName
    : IDENTIFIER
    | QUALIFIED_IDENTIFIER
    ;

typeLit
    : arrayType
    | structType
    | referenceType
    | functionType
    | interfaceType
    | sliceType
    | mapType
    | pipeType
    ;


arrayType
    : BRACKET_START arrayLength BRACKET_END elementType
    ;

arrayLength
    : expression
    ;

elementType
    : typeId
    ;

referenceType
    : AMP typeId
    ;

interfaceType
    : INTERFACE BLOCK_START ( methodSpec )* BLOCK_END
    ;

sliceType
    : BRACKET_START BRACKET_END elementType
    ;

mapType
    : MAP LT typeId COMMA elementType GT
    ;

pipeType
    : ( PIPE | PIPEIN | PIPEOUT ) LT elementType GT
    ;

methodSpec
    : IDENTIFIER signature
    | typeName
    ;


functionType
    : FUNC signature
    ;

signature
    : parameters result?
    ;

result
    : parameters
    | typeId
    ;

parameters
    : PAREN_START ( parameterList )? PAREN_END
    ;

parameterList
    : parameterDecl ( COMMA STMTEND? parameterDecl )* STMTEND?
    ;

parameterDecl
    : identifierList? ARRAY_EXPAND? typeId
    ;

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Operands

operand
    : literal
    | operandName
    | methodExpr
    | sequenceExpr
    | runExpr
    | PAREN_START expression PAREN_END
    ;

literal
    : basicLit
    | compositeLit
    | functionLit
    ;

basicLit
    : INT_LIT
    | FLOAT_LIT
    | IMAGINARY_LIT
    | RUNE_LIT
    | STRING_LIT
    ;

sequenceExpr
    : ( literal | operandName ) SEQ_SEP ( literal | operandName )?
    ;

operandName
    : IDENTIFIER
    | QUALIFIED_IDENTIFIER
    ;

compositeLit
    : literalType literalValue
    ;

literalType
    : structType
    | arrayType
    | BRACKET_START ARRAY_EXPAND BRACKET_END elementType
    | sliceType
    | mapType
    | typeName
    ;

literalValue
    : BLOCK_START ( elementList COMMA? )? BLOCK_END
    ;

elementList
    : keyedElement (COMMA keyedElement)*
    ;

keyedElement
    : (key COLON)? element
    ;

key
    : IDENTIFIER
    | expression
    | literalValue
    ;

element
    : expression
    | literalValue
    ;

structType
    : STRUCT BLOCK_START ( fieldDecl )* BLOCK_END
    ;

fieldDecl
    : (identifierList typeId | anonymousField) STRING_LIT?
    ;

anonymousField
    : AMP? typeName
    ;

functionLit
    : FUNC function
    ;

primaryExpr
    : operand
    | conversion
    | primaryExpr selector
    | primaryExpr index
    | primaryExpr slice
    | primaryExpr typeAssertion
	| primaryExpr arguments
    ;

selector
    : DOT IDENTIFIER
    ;

index
    : BRACKET_START expression BRACKET_END
    ;

slice
    : BRACKET_START (( expression? COLON expression? ) | ( expression? COLON expression COLON expression )) BRACKET_END
    ;

typeAssertion
    : DOT PAREN_START typeId PAREN_END
    ;

arguments
    : PAREN_START ( ( expressionList | typeId ( COMMA expressionList )? ) ARRAY_EXPAND? COMMA? )? PAREN_END
    ;

methodExpr
    : receiverType DOT IDENTIFIER
    ;

receiverType
    : typeName
    | PAREN_START AMP typeName PAREN_END
    | PAREN_START receiverType PAREN_END
    ;

expression
    : unaryExpr
    | expression binaryOp expression
    ;

unaryExpr
    : primaryExpr
    | UNARY_OP unaryExpr
    ;

conversion
    : typeId PAREN_START expression PAREN_END
    ;

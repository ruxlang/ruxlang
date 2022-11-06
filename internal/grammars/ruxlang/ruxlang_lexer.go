// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package ruxlang

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RuxlangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ruxlanglexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ruxlanglexerLexerInit() {
	staticData := &ruxlanglexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'package'", "'import'", "'const'", "'map'", "'pipe'", "'pipein'",
		"'pipeout'", "'interface'", "'struct'", "'type'", "'func'", "'var'",
		"'return'", "'yield'", "'break'", "'continue'", "'goto'", "'fallthrough'",
		"'defer'", "'if'", "'else'", "'switch'", "'select'", "'case'", "'default'",
		"'for'", "'in'", "'run'", "", "", "", "", "", "", "", "", "", "':='",
		"'='", "'&'", "':'", "','", "'<-'", "'++'", "'--'", "'...'", "'..'",
		"'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'||'", "'&&'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'|'", "'^'", "'<<'", "'>>'", "'&^'", "'!'",
		"", "", "", "", "'.'",
	}
	staticData.symbolicNames = []string{
		"", "PACKAGE", "IMPORT", "CONST", "MAP", "PIPE", "PIPEIN", "PIPEOUT",
		"INTERFACE", "STRUCT", "TYPE", "FUNC", "VAR", "RETURN", "YIELD", "BREAK",
		"CONTINUE", "GOTO", "FALLTHROUGH", "DEFER", "IF", "ELSE", "SWITCH",
		"SELECT", "CASE", "DEFAULT", "FOR", "IN", "RUN", "TYPE_UNWRAP", "PAREN_START",
		"PAREN_END", "BRACKET_START", "BRACKET_END", "STMTEND", "BLOCK_START",
		"BLOCK_END", "ASSIGN_WITH_OP", "ASSIGN_NEW", "ASSIGN", "AMP", "COLON",
		"COMMA", "SEND", "INC", "DEC", "ARRAY_EXPAND", "SEQ_SEP", "LT", "GT",
		"LTE", "GTE", "EQ", "NEQ", "OR", "AND", "ADD", "SUB", "MUL", "DIV",
		"MOD", "BIT_OR", "CARET", "SHL", "SHR", "BIT_CLEAR", "NOT", "UNARY_OP",
		"INT_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "DOT", "QUALIFIED_IDENTIFIER",
		"EXTENDED_IDENTIFIER", "IDENTIFIER", "RUNE_LIT", "STRING_LIT", "LITTLE_U_VALUE",
		"BIG_U_VALUE", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"PACKAGE", "IMPORT", "CONST", "MAP", "PIPE", "PIPEIN", "PIPEOUT", "INTERFACE",
		"STRUCT", "TYPE", "FUNC", "VAR", "RETURN", "YIELD", "BREAK", "CONTINUE",
		"GOTO", "FALLTHROUGH", "DEFER", "IF", "ELSE", "SWITCH", "SELECT", "CASE",
		"DEFAULT", "FOR", "IN", "RUN", "TYPE_UNWRAP", "PAREN_START", "PAREN_END",
		"BRACKET_START", "BRACKET_END", "STMTEND", "BLOCK_START", "BLOCK_END",
		"NEWLINE", "ASSIGN_WITH_OP", "ASSIGN_NEW", "ASSIGN", "AMP", "COLON",
		"COMMA", "SEND", "INC", "DEC", "ARRAY_EXPAND", "SEQ_SEP", "LT", "GT",
		"LTE", "GTE", "EQ", "NEQ", "OR", "AND", "ADD", "SUB", "MUL", "DIV",
		"MOD", "BIT_OR", "CARET", "SHL", "SHR", "BIT_CLEAR", "NOT", "UNARY_OP",
		"INT_LIT", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMALS",
		"EXPONENT", "IMAGINARY_LIT", "DOT", "QUALIFIED_IDENTIFIER", "EXTENDED_IDENTIFIER",
		"IDENTIFIER", "RUNE_LIT", "STRING_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"STRING_CHAR", "RAW_STRING_CHAR", "UNICODE_VALUE", "LITTLE_U_VALUE",
		"BIG_U_VALUE", "BYTE_VALUE", "OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", "ESCAPED_CHAR",
		"LETTER", "DECIMAL_DIGIT", "OCTAL_DIGIT", "HEX_DIGIT", "UNICODE_CHAR",
		"UNICODE_DIGIT", "UNICODE_LETTER", "WS", "COMMENT", "LINE_COMMENT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 81, 768, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		2, 89, 7, 89, 2, 90, 7, 90, 2, 91, 7, 91, 2, 92, 7, 92, 2, 93, 7, 93, 2,
		94, 7, 94, 2, 95, 7, 95, 2, 96, 7, 96, 2, 97, 7, 97, 2, 98, 7, 98, 2, 99,
		7, 99, 2, 100, 7, 100, 2, 101, 7, 101, 2, 102, 7, 102, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 29, 3, 29, 391, 8, 29, 1, 29, 1, 29, 3, 29, 395,
		8, 29, 1, 30, 3, 30, 398, 8, 30, 1, 30, 1, 30, 3, 30, 402, 8, 30, 1, 31,
		3, 31, 405, 8, 31, 1, 31, 1, 31, 3, 31, 409, 8, 31, 1, 32, 3, 32, 412,
		8, 32, 1, 32, 1, 32, 3, 32, 416, 8, 32, 1, 33, 4, 33, 419, 8, 33, 11, 33,
		12, 33, 420, 1, 33, 3, 33, 424, 8, 33, 1, 34, 3, 34, 427, 8, 34, 1, 34,
		1, 34, 3, 34, 431, 8, 34, 1, 35, 3, 35, 434, 8, 35, 1, 35, 1, 35, 3, 35,
		438, 8, 35, 1, 36, 3, 36, 441, 8, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 453, 8, 37, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42,
		1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1,
		46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1,
		53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57,
		1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1,
		63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1, 66, 1, 66,
		1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 537, 8, 67, 1, 68, 1,
		68, 1, 68, 3, 68, 542, 8, 68, 1, 69, 1, 69, 5, 69, 546, 8, 69, 10, 69,
		12, 69, 549, 9, 69, 1, 70, 1, 70, 5, 70, 553, 8, 70, 10, 70, 12, 70, 556,
		9, 70, 1, 71, 1, 71, 1, 71, 4, 71, 561, 8, 71, 11, 71, 12, 71, 562, 1,
		72, 1, 72, 1, 72, 1, 72, 3, 72, 569, 8, 72, 1, 72, 3, 72, 572, 8, 72, 1,
		72, 1, 72, 1, 72, 3, 72, 577, 8, 72, 3, 72, 579, 8, 72, 1, 73, 4, 73, 582,
		8, 73, 11, 73, 12, 73, 583, 1, 74, 1, 74, 3, 74, 588, 8, 74, 1, 74, 1,
		74, 1, 75, 1, 75, 3, 75, 594, 8, 75, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77,
		1, 77, 1, 77, 1, 77, 1, 78, 1, 78, 1, 78, 1, 78, 1, 78, 1, 78, 4, 78, 610,
		8, 78, 11, 78, 12, 78, 611, 1, 79, 1, 79, 1, 79, 5, 79, 617, 8, 79, 10,
		79, 12, 79, 620, 9, 79, 1, 80, 1, 80, 1, 80, 3, 80, 625, 8, 80, 1, 80,
		1, 80, 1, 81, 1, 81, 3, 81, 631, 8, 81, 1, 82, 1, 82, 1, 82, 5, 82, 636,
		8, 82, 10, 82, 12, 82, 639, 9, 82, 1, 82, 1, 82, 1, 83, 1, 83, 1, 83, 5,
		83, 646, 8, 83, 10, 83, 12, 83, 649, 9, 83, 1, 83, 1, 83, 1, 84, 1, 84,
		1, 84, 1, 84, 3, 84, 657, 8, 84, 1, 85, 1, 85, 1, 85, 1, 85, 3, 85, 663,
		8, 85, 1, 86, 1, 86, 1, 86, 1, 86, 3, 86, 669, 8, 86, 1, 87, 1, 87, 1,
		87, 1, 87, 1, 87, 1, 87, 1, 87, 1, 87, 1, 88, 1, 88, 1, 88, 1, 88, 1, 88,
		1, 88, 1, 88, 1, 88, 1, 88, 1, 88, 1, 88, 1, 88, 1, 89, 1, 89, 3, 89, 693,
		8, 89, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 91, 1, 91, 1, 91, 1, 91, 1,
		91, 1, 92, 1, 92, 1, 92, 1, 93, 1, 93, 3, 93, 710, 8, 93, 1, 94, 1, 94,
		1, 95, 1, 95, 1, 96, 1, 96, 1, 97, 1, 97, 1, 98, 3, 98, 721, 8, 98, 1,
		99, 3, 99, 724, 8, 99, 1, 100, 4, 100, 727, 8, 100, 11, 100, 12, 100, 728,
		1, 100, 1, 100, 1, 101, 1, 101, 1, 101, 1, 101, 5, 101, 737, 8, 101, 10,
		101, 12, 101, 740, 9, 101, 1, 101, 1, 101, 1, 101, 1, 101, 1, 101, 1, 102,
		1, 102, 1, 102, 1, 102, 5, 102, 751, 8, 102, 10, 102, 12, 102, 754, 9,
		102, 1, 102, 1, 102, 1, 102, 5, 102, 759, 8, 102, 10, 102, 12, 102, 762,
		9, 102, 1, 102, 3, 102, 765, 8, 102, 1, 102, 1, 102, 1, 738, 0, 103, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 0, 75, 37,
		77, 38, 79, 39, 81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 91, 45, 93, 46,
		95, 47, 97, 48, 99, 49, 101, 50, 103, 51, 105, 52, 107, 53, 109, 54, 111,
		55, 113, 56, 115, 57, 117, 58, 119, 59, 121, 60, 123, 61, 125, 62, 127,
		63, 129, 64, 131, 65, 133, 66, 135, 67, 137, 68, 139, 0, 141, 0, 143, 0,
		145, 69, 147, 0, 149, 0, 151, 70, 153, 71, 155, 72, 157, 73, 159, 74, 161,
		75, 163, 76, 165, 0, 167, 0, 169, 0, 171, 0, 173, 0, 175, 77, 177, 78,
		179, 0, 181, 0, 183, 0, 185, 0, 187, 0, 189, 0, 191, 0, 193, 0, 195, 0,
		197, 0, 199, 0, 201, 79, 203, 80, 205, 81, 1, 0, 17, 1, 0, 13, 13, 1, 0,
		10, 10, 6, 0, 37, 37, 42, 43, 45, 45, 47, 47, 94, 94, 124, 124, 1, 0, 49,
		57, 2, 0, 88, 88, 120, 120, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45,
		2, 0, 34, 34, 92, 92, 1, 0, 96, 96, 9, 0, 34, 34, 39, 39, 92, 92, 97, 98,
		102, 102, 110, 110, 114, 114, 116, 116, 118, 118, 1, 0, 48, 57, 1, 0, 48,
		55, 3, 0, 48, 57, 65, 70, 97, 102, 20, 0, 48, 57, 1632, 1641, 1776, 1785,
		2406, 2415, 2534, 2543, 2662, 2671, 2790, 2799, 2918, 2927, 3047, 3055,
		3174, 3183, 3302, 3311, 3430, 3439, 3664, 3673, 3792, 3801, 3872, 3881,
		4160, 4169, 4969, 4977, 6112, 6121, 6160, 6169, 65296, 65305, 258, 0, 65,
		90, 97, 122, 170, 170, 181, 181, 186, 186, 192, 214, 216, 246, 248, 543,
		546, 563, 592, 685, 688, 696, 699, 705, 720, 721, 736, 740, 750, 750, 890,
		890, 902, 902, 904, 906, 908, 908, 910, 929, 931, 974, 976, 983, 986, 1011,
		1024, 1153, 1164, 1220, 1223, 1224, 1227, 1228, 1232, 1269, 1272, 1273,
		1329, 1366, 1369, 1369, 1377, 1415, 1488, 1514, 1520, 1522, 1569, 1594,
		1600, 1610, 1649, 1747, 1749, 1749, 1765, 1766, 1786, 1788, 1808, 1808,
		1810, 1836, 1920, 1957, 2309, 2361, 2365, 2365, 2384, 2384, 2392, 2401,
		2437, 2444, 2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489,
		2524, 2525, 2527, 2529, 2544, 2545, 2565, 2570, 2575, 2576, 2579, 2600,
		2602, 2608, 2610, 2611, 2613, 2614, 2616, 2617, 2649, 2652, 2654, 2654,
		2674, 2676, 2693, 2699, 2701, 2701, 2703, 2705, 2707, 2728, 2730, 2736,
		2738, 2739, 2741, 2745, 2749, 2749, 2768, 2768, 2784, 2784, 2821, 2828,
		2831, 2832, 2835, 2856, 2858, 2864, 2866, 2867, 2870, 2873, 2877, 2877,
		2908, 2909, 2911, 2913, 2949, 2954, 2958, 2960, 2962, 2965, 2969, 2970,
		2972, 2972, 2974, 2975, 2979, 2980, 2984, 2986, 2990, 2997, 2999, 3001,
		3077, 3084, 3086, 3088, 3090, 3112, 3114, 3123, 3125, 3129, 3168, 3169,
		3205, 3212, 3214, 3216, 3218, 3240, 3242, 3251, 3253, 3257, 3294, 3294,
		3296, 3297, 3333, 3340, 3342, 3344, 3346, 3368, 3370, 3385, 3424, 3425,
		3461, 3478, 3482, 3505, 3507, 3515, 3517, 3517, 3520, 3526, 3585, 3632,
		3634, 3635, 3648, 3654, 3713, 3714, 3716, 3716, 3719, 3720, 3722, 3722,
		3725, 3725, 3732, 3735, 3737, 3743, 3745, 3747, 3749, 3749, 3751, 3751,
		3754, 3755, 3757, 3760, 3762, 3763, 3773, 3780, 3782, 3782, 3804, 3805,
		3840, 3840, 3904, 3946, 3976, 3979, 4096, 4129, 4131, 4135, 4137, 4138,
		4176, 4181, 4256, 4293, 4304, 4342, 4352, 4441, 4447, 4514, 4520, 4601,
		4608, 4614, 4616, 4678, 4680, 4680, 4682, 4685, 4688, 4694, 4696, 4696,
		4698, 4701, 4704, 4742, 4744, 4744, 4746, 4749, 4752, 4782, 4784, 4784,
		4786, 4789, 4792, 4798, 4800, 4800, 4802, 4805, 4808, 4814, 4816, 4822,
		4824, 4846, 4848, 4878, 4880, 4880, 4882, 4885, 4888, 4894, 4896, 4934,
		4936, 4954, 5024, 5108, 5121, 5750, 5761, 5786, 5792, 5866, 6016, 6067,
		6176, 6263, 6272, 6312, 7680, 7835, 7840, 7929, 7936, 7957, 7960, 7965,
		7968, 8005, 8008, 8013, 8016, 8023, 8025, 8025, 8027, 8027, 8029, 8029,
		8031, 8061, 8064, 8116, 8118, 8124, 8126, 8126, 8130, 8132, 8134, 8140,
		8144, 8147, 8150, 8155, 8160, 8172, 8178, 8180, 8182, 8188, 8319, 8319,
		8450, 8450, 8455, 8455, 8458, 8467, 8469, 8469, 8473, 8477, 8484, 8484,
		8486, 8486, 8488, 8488, 8490, 8493, 8495, 8497, 8499, 8505, 8544, 8579,
		12293, 12295, 12321, 12329, 12337, 12341, 12344, 12346, 12353, 12436, 12445,
		12446, 12449, 12538, 12540, 12542, 12549, 12588, 12593, 12686, 12704, 12727,
		13312, 13312, 19893, 19893, 19968, 19968, 40869, 40869, 40960, 42124, 44032,
		44032, 55203, 55203, 63744, 64045, 64256, 64262, 64275, 64279, 64285, 64285,
		64287, 64296, 64298, 64310, 64312, 64316, 64318, 64318, 64320, 64321, 64323,
		64324, 64326, 64433, 64467, 64829, 64848, 64911, 64914, 64967, 65008, 65019,
		65136, 65138, 65140, 65140, 65142, 65276, 65313, 65338, 65345, 65370, 65382,
		65470, 65474, 65479, 65482, 65487, 65490, 65495, 65498, 65500, 3, 0, 9,
		9, 12, 12, 32, 32, 2, 0, 10, 10, 13, 13, 806, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0,
		0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0,
		0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0,
		0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105,
		1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0,
		0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1,
		0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0,
		127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0,
		0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 151,
		1, 0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0,
		0, 159, 1, 0, 0, 0, 0, 161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0, 0, 175, 1,
		0, 0, 0, 0, 177, 1, 0, 0, 0, 0, 201, 1, 0, 0, 0, 0, 203, 1, 0, 0, 0, 0,
		205, 1, 0, 0, 0, 1, 207, 1, 0, 0, 0, 3, 215, 1, 0, 0, 0, 5, 222, 1, 0,
		0, 0, 7, 228, 1, 0, 0, 0, 9, 232, 1, 0, 0, 0, 11, 237, 1, 0, 0, 0, 13,
		244, 1, 0, 0, 0, 15, 252, 1, 0, 0, 0, 17, 262, 1, 0, 0, 0, 19, 269, 1,
		0, 0, 0, 21, 274, 1, 0, 0, 0, 23, 279, 1, 0, 0, 0, 25, 283, 1, 0, 0, 0,
		27, 290, 1, 0, 0, 0, 29, 296, 1, 0, 0, 0, 31, 302, 1, 0, 0, 0, 33, 311,
		1, 0, 0, 0, 35, 316, 1, 0, 0, 0, 37, 328, 1, 0, 0, 0, 39, 334, 1, 0, 0,
		0, 41, 337, 1, 0, 0, 0, 43, 342, 1, 0, 0, 0, 45, 349, 1, 0, 0, 0, 47, 356,
		1, 0, 0, 0, 49, 361, 1, 0, 0, 0, 51, 369, 1, 0, 0, 0, 53, 373, 1, 0, 0,
		0, 55, 376, 1, 0, 0, 0, 57, 380, 1, 0, 0, 0, 59, 390, 1, 0, 0, 0, 61, 397,
		1, 0, 0, 0, 63, 404, 1, 0, 0, 0, 65, 411, 1, 0, 0, 0, 67, 423, 1, 0, 0,
		0, 69, 426, 1, 0, 0, 0, 71, 433, 1, 0, 0, 0, 73, 440, 1, 0, 0, 0, 75, 452,
		1, 0, 0, 0, 77, 456, 1, 0, 0, 0, 79, 459, 1, 0, 0, 0, 81, 461, 1, 0, 0,
		0, 83, 463, 1, 0, 0, 0, 85, 465, 1, 0, 0, 0, 87, 467, 1, 0, 0, 0, 89, 470,
		1, 0, 0, 0, 91, 473, 1, 0, 0, 0, 93, 476, 1, 0, 0, 0, 95, 480, 1, 0, 0,
		0, 97, 483, 1, 0, 0, 0, 99, 485, 1, 0, 0, 0, 101, 487, 1, 0, 0, 0, 103,
		490, 1, 0, 0, 0, 105, 493, 1, 0, 0, 0, 107, 496, 1, 0, 0, 0, 109, 499,
		1, 0, 0, 0, 111, 502, 1, 0, 0, 0, 113, 505, 1, 0, 0, 0, 115, 507, 1, 0,
		0, 0, 117, 509, 1, 0, 0, 0, 119, 511, 1, 0, 0, 0, 121, 513, 1, 0, 0, 0,
		123, 515, 1, 0, 0, 0, 125, 517, 1, 0, 0, 0, 127, 519, 1, 0, 0, 0, 129,
		522, 1, 0, 0, 0, 131, 525, 1, 0, 0, 0, 133, 528, 1, 0, 0, 0, 135, 536,
		1, 0, 0, 0, 137, 541, 1, 0, 0, 0, 139, 543, 1, 0, 0, 0, 141, 550, 1, 0,
		0, 0, 143, 557, 1, 0, 0, 0, 145, 578, 1, 0, 0, 0, 147, 581, 1, 0, 0, 0,
		149, 585, 1, 0, 0, 0, 151, 593, 1, 0, 0, 0, 153, 597, 1, 0, 0, 0, 155,
		599, 1, 0, 0, 0, 157, 603, 1, 0, 0, 0, 159, 613, 1, 0, 0, 0, 161, 621,
		1, 0, 0, 0, 163, 630, 1, 0, 0, 0, 165, 632, 1, 0, 0, 0, 167, 642, 1, 0,
		0, 0, 169, 656, 1, 0, 0, 0, 171, 662, 1, 0, 0, 0, 173, 668, 1, 0, 0, 0,
		175, 670, 1, 0, 0, 0, 177, 678, 1, 0, 0, 0, 179, 692, 1, 0, 0, 0, 181,
		694, 1, 0, 0, 0, 183, 699, 1, 0, 0, 0, 185, 704, 1, 0, 0, 0, 187, 709,
		1, 0, 0, 0, 189, 711, 1, 0, 0, 0, 191, 713, 1, 0, 0, 0, 193, 715, 1, 0,
		0, 0, 195, 717, 1, 0, 0, 0, 197, 720, 1, 0, 0, 0, 199, 723, 1, 0, 0, 0,
		201, 726, 1, 0, 0, 0, 203, 732, 1, 0, 0, 0, 205, 764, 1, 0, 0, 0, 207,
		208, 5, 112, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 99, 0, 0, 210, 211,
		5, 107, 0, 0, 211, 212, 5, 97, 0, 0, 212, 213, 5, 103, 0, 0, 213, 214,
		5, 101, 0, 0, 214, 2, 1, 0, 0, 0, 215, 216, 5, 105, 0, 0, 216, 217, 5,
		109, 0, 0, 217, 218, 5, 112, 0, 0, 218, 219, 5, 111, 0, 0, 219, 220, 5,
		114, 0, 0, 220, 221, 5, 116, 0, 0, 221, 4, 1, 0, 0, 0, 222, 223, 5, 99,
		0, 0, 223, 224, 5, 111, 0, 0, 224, 225, 5, 110, 0, 0, 225, 226, 5, 115,
		0, 0, 226, 227, 5, 116, 0, 0, 227, 6, 1, 0, 0, 0, 228, 229, 5, 109, 0,
		0, 229, 230, 5, 97, 0, 0, 230, 231, 5, 112, 0, 0, 231, 8, 1, 0, 0, 0, 232,
		233, 5, 112, 0, 0, 233, 234, 5, 105, 0, 0, 234, 235, 5, 112, 0, 0, 235,
		236, 5, 101, 0, 0, 236, 10, 1, 0, 0, 0, 237, 238, 5, 112, 0, 0, 238, 239,
		5, 105, 0, 0, 239, 240, 5, 112, 0, 0, 240, 241, 5, 101, 0, 0, 241, 242,
		5, 105, 0, 0, 242, 243, 5, 110, 0, 0, 243, 12, 1, 0, 0, 0, 244, 245, 5,
		112, 0, 0, 245, 246, 5, 105, 0, 0, 246, 247, 5, 112, 0, 0, 247, 248, 5,
		101, 0, 0, 248, 249, 5, 111, 0, 0, 249, 250, 5, 117, 0, 0, 250, 251, 5,
		116, 0, 0, 251, 14, 1, 0, 0, 0, 252, 253, 5, 105, 0, 0, 253, 254, 5, 110,
		0, 0, 254, 255, 5, 116, 0, 0, 255, 256, 5, 101, 0, 0, 256, 257, 5, 114,
		0, 0, 257, 258, 5, 102, 0, 0, 258, 259, 5, 97, 0, 0, 259, 260, 5, 99, 0,
		0, 260, 261, 5, 101, 0, 0, 261, 16, 1, 0, 0, 0, 262, 263, 5, 115, 0, 0,
		263, 264, 5, 116, 0, 0, 264, 265, 5, 114, 0, 0, 265, 266, 5, 117, 0, 0,
		266, 267, 5, 99, 0, 0, 267, 268, 5, 116, 0, 0, 268, 18, 1, 0, 0, 0, 269,
		270, 5, 116, 0, 0, 270, 271, 5, 121, 0, 0, 271, 272, 5, 112, 0, 0, 272,
		273, 5, 101, 0, 0, 273, 20, 1, 0, 0, 0, 274, 275, 5, 102, 0, 0, 275, 276,
		5, 117, 0, 0, 276, 277, 5, 110, 0, 0, 277, 278, 5, 99, 0, 0, 278, 22, 1,
		0, 0, 0, 279, 280, 5, 118, 0, 0, 280, 281, 5, 97, 0, 0, 281, 282, 5, 114,
		0, 0, 282, 24, 1, 0, 0, 0, 283, 284, 5, 114, 0, 0, 284, 285, 5, 101, 0,
		0, 285, 286, 5, 116, 0, 0, 286, 287, 5, 117, 0, 0, 287, 288, 5, 114, 0,
		0, 288, 289, 5, 110, 0, 0, 289, 26, 1, 0, 0, 0, 290, 291, 5, 121, 0, 0,
		291, 292, 5, 105, 0, 0, 292, 293, 5, 101, 0, 0, 293, 294, 5, 108, 0, 0,
		294, 295, 5, 100, 0, 0, 295, 28, 1, 0, 0, 0, 296, 297, 5, 98, 0, 0, 297,
		298, 5, 114, 0, 0, 298, 299, 5, 101, 0, 0, 299, 300, 5, 97, 0, 0, 300,
		301, 5, 107, 0, 0, 301, 30, 1, 0, 0, 0, 302, 303, 5, 99, 0, 0, 303, 304,
		5, 111, 0, 0, 304, 305, 5, 110, 0, 0, 305, 306, 5, 116, 0, 0, 306, 307,
		5, 105, 0, 0, 307, 308, 5, 110, 0, 0, 308, 309, 5, 117, 0, 0, 309, 310,
		5, 101, 0, 0, 310, 32, 1, 0, 0, 0, 311, 312, 5, 103, 0, 0, 312, 313, 5,
		111, 0, 0, 313, 314, 5, 116, 0, 0, 314, 315, 5, 111, 0, 0, 315, 34, 1,
		0, 0, 0, 316, 317, 5, 102, 0, 0, 317, 318, 5, 97, 0, 0, 318, 319, 5, 108,
		0, 0, 319, 320, 5, 108, 0, 0, 320, 321, 5, 116, 0, 0, 321, 322, 5, 104,
		0, 0, 322, 323, 5, 114, 0, 0, 323, 324, 5, 111, 0, 0, 324, 325, 5, 117,
		0, 0, 325, 326, 5, 103, 0, 0, 326, 327, 5, 104, 0, 0, 327, 36, 1, 0, 0,
		0, 328, 329, 5, 100, 0, 0, 329, 330, 5, 101, 0, 0, 330, 331, 5, 102, 0,
		0, 331, 332, 5, 101, 0, 0, 332, 333, 5, 114, 0, 0, 333, 38, 1, 0, 0, 0,
		334, 335, 5, 105, 0, 0, 335, 336, 5, 102, 0, 0, 336, 40, 1, 0, 0, 0, 337,
		338, 5, 101, 0, 0, 338, 339, 5, 108, 0, 0, 339, 340, 5, 115, 0, 0, 340,
		341, 5, 101, 0, 0, 341, 42, 1, 0, 0, 0, 342, 343, 5, 115, 0, 0, 343, 344,
		5, 119, 0, 0, 344, 345, 5, 105, 0, 0, 345, 346, 5, 116, 0, 0, 346, 347,
		5, 99, 0, 0, 347, 348, 5, 104, 0, 0, 348, 44, 1, 0, 0, 0, 349, 350, 5,
		115, 0, 0, 350, 351, 5, 101, 0, 0, 351, 352, 5, 108, 0, 0, 352, 353, 5,
		101, 0, 0, 353, 354, 5, 99, 0, 0, 354, 355, 5, 116, 0, 0, 355, 46, 1, 0,
		0, 0, 356, 357, 5, 99, 0, 0, 357, 358, 5, 97, 0, 0, 358, 359, 5, 115, 0,
		0, 359, 360, 5, 101, 0, 0, 360, 48, 1, 0, 0, 0, 361, 362, 5, 100, 0, 0,
		362, 363, 5, 101, 0, 0, 363, 364, 5, 102, 0, 0, 364, 365, 5, 97, 0, 0,
		365, 366, 5, 117, 0, 0, 366, 367, 5, 108, 0, 0, 367, 368, 5, 116, 0, 0,
		368, 50, 1, 0, 0, 0, 369, 370, 5, 102, 0, 0, 370, 371, 5, 111, 0, 0, 371,
		372, 5, 114, 0, 0, 372, 52, 1, 0, 0, 0, 373, 374, 5, 105, 0, 0, 374, 375,
		5, 110, 0, 0, 375, 54, 1, 0, 0, 0, 376, 377, 5, 114, 0, 0, 377, 378, 5,
		117, 0, 0, 378, 379, 5, 110, 0, 0, 379, 56, 1, 0, 0, 0, 380, 381, 3, 153,
		76, 0, 381, 382, 3, 59, 29, 0, 382, 383, 5, 116, 0, 0, 383, 384, 5, 121,
		0, 0, 384, 385, 5, 112, 0, 0, 385, 386, 5, 101, 0, 0, 386, 387, 1, 0, 0,
		0, 387, 388, 3, 61, 30, 0, 388, 58, 1, 0, 0, 0, 389, 391, 3, 67, 33, 0,
		390, 389, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392,
		394, 5, 40, 0, 0, 393, 395, 3, 67, 33, 0, 394, 393, 1, 0, 0, 0, 394, 395,
		1, 0, 0, 0, 395, 60, 1, 0, 0, 0, 396, 398, 3, 67, 33, 0, 397, 396, 1, 0,
		0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 401, 5, 41, 0, 0,
		400, 402, 3, 67, 33, 0, 401, 400, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402,
		62, 1, 0, 0, 0, 403, 405, 3, 67, 33, 0, 404, 403, 1, 0, 0, 0, 404, 405,
		1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 408, 5, 91, 0, 0, 407, 409, 3, 67,
		33, 0, 408, 407, 1, 0, 0, 0, 408, 409, 1, 0, 0, 0, 409, 64, 1, 0, 0, 0,
		410, 412, 3, 67, 33, 0, 411, 410, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412,
		413, 1, 0, 0, 0, 413, 415, 5, 93, 0, 0, 414, 416, 3, 67, 33, 0, 415, 414,
		1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 66, 1, 0, 0, 0, 417, 419, 3, 73,
		36, 0, 418, 417, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0,
		420, 421, 1, 0, 0, 0, 421, 424, 1, 0, 0, 0, 422, 424, 5, 0, 0, 1, 423,
		418, 1, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424, 68, 1, 0, 0, 0, 425, 427, 3,
		67, 33, 0, 426, 425, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 428, 1, 0,
		0, 0, 428, 430, 5, 123, 0, 0, 429, 431, 3, 67, 33, 0, 430, 429, 1, 0, 0,
		0, 430, 431, 1, 0, 0, 0, 431, 70, 1, 0, 0, 0, 432, 434, 3, 67, 33, 0, 433,
		432, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 437,
		5, 125, 0, 0, 436, 438, 3, 67, 33, 0, 437, 436, 1, 0, 0, 0, 437, 438, 1,
		0, 0, 0, 438, 72, 1, 0, 0, 0, 439, 441, 7, 0, 0, 0, 440, 439, 1, 0, 0,
		0, 440, 441, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 7, 1, 0, 0, 443,
		74, 1, 0, 0, 0, 444, 453, 7, 2, 0, 0, 445, 446, 5, 60, 0, 0, 446, 453,
		5, 60, 0, 0, 447, 448, 5, 62, 0, 0, 448, 453, 5, 62, 0, 0, 449, 453, 5,
		38, 0, 0, 450, 451, 5, 38, 0, 0, 451, 453, 5, 94, 0, 0, 452, 444, 1, 0,
		0, 0, 452, 445, 1, 0, 0, 0, 452, 447, 1, 0, 0, 0, 452, 449, 1, 0, 0, 0,
		452, 450, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 455, 5, 61, 0, 0, 455,
		76, 1, 0, 0, 0, 456, 457, 5, 58, 0, 0, 457, 458, 5, 61, 0, 0, 458, 78,
		1, 0, 0, 0, 459, 460, 5, 61, 0, 0, 460, 80, 1, 0, 0, 0, 461, 462, 5, 38,
		0, 0, 462, 82, 1, 0, 0, 0, 463, 464, 5, 58, 0, 0, 464, 84, 1, 0, 0, 0,
		465, 466, 5, 44, 0, 0, 466, 86, 1, 0, 0, 0, 467, 468, 5, 60, 0, 0, 468,
		469, 5, 45, 0, 0, 469, 88, 1, 0, 0, 0, 470, 471, 5, 43, 0, 0, 471, 472,
		5, 43, 0, 0, 472, 90, 1, 0, 0, 0, 473, 474, 5, 45, 0, 0, 474, 475, 5, 45,
		0, 0, 475, 92, 1, 0, 0, 0, 476, 477, 5, 46, 0, 0, 477, 478, 5, 46, 0, 0,
		478, 479, 5, 46, 0, 0, 479, 94, 1, 0, 0, 0, 480, 481, 5, 46, 0, 0, 481,
		482, 5, 46, 0, 0, 482, 96, 1, 0, 0, 0, 483, 484, 5, 60, 0, 0, 484, 98,
		1, 0, 0, 0, 485, 486, 5, 62, 0, 0, 486, 100, 1, 0, 0, 0, 487, 488, 5, 60,
		0, 0, 488, 489, 5, 61, 0, 0, 489, 102, 1, 0, 0, 0, 490, 491, 5, 62, 0,
		0, 491, 492, 5, 61, 0, 0, 492, 104, 1, 0, 0, 0, 493, 494, 5, 61, 0, 0,
		494, 495, 5, 61, 0, 0, 495, 106, 1, 0, 0, 0, 496, 497, 5, 33, 0, 0, 497,
		498, 5, 61, 0, 0, 498, 108, 1, 0, 0, 0, 499, 500, 5, 124, 0, 0, 500, 501,
		5, 124, 0, 0, 501, 110, 1, 0, 0, 0, 502, 503, 5, 38, 0, 0, 503, 504, 5,
		38, 0, 0, 504, 112, 1, 0, 0, 0, 505, 506, 5, 43, 0, 0, 506, 114, 1, 0,
		0, 0, 507, 508, 5, 45, 0, 0, 508, 116, 1, 0, 0, 0, 509, 510, 5, 42, 0,
		0, 510, 118, 1, 0, 0, 0, 511, 512, 5, 47, 0, 0, 512, 120, 1, 0, 0, 0, 513,
		514, 5, 37, 0, 0, 514, 122, 1, 0, 0, 0, 515, 516, 5, 124, 0, 0, 516, 124,
		1, 0, 0, 0, 517, 518, 5, 94, 0, 0, 518, 126, 1, 0, 0, 0, 519, 520, 5, 60,
		0, 0, 520, 521, 5, 60, 0, 0, 521, 128, 1, 0, 0, 0, 522, 523, 5, 62, 0,
		0, 523, 524, 5, 62, 0, 0, 524, 130, 1, 0, 0, 0, 525, 526, 5, 38, 0, 0,
		526, 527, 5, 94, 0, 0, 527, 132, 1, 0, 0, 0, 528, 529, 5, 33, 0, 0, 529,
		134, 1, 0, 0, 0, 530, 537, 3, 113, 56, 0, 531, 537, 3, 115, 57, 0, 532,
		537, 3, 133, 66, 0, 533, 537, 3, 125, 62, 0, 534, 537, 3, 81, 40, 0, 535,
		537, 3, 87, 43, 0, 536, 530, 1, 0, 0, 0, 536, 531, 1, 0, 0, 0, 536, 532,
		1, 0, 0, 0, 536, 533, 1, 0, 0, 0, 536, 534, 1, 0, 0, 0, 536, 535, 1, 0,
		0, 0, 537, 136, 1, 0, 0, 0, 538, 542, 3, 139, 69, 0, 539, 542, 3, 141,
		70, 0, 540, 542, 3, 143, 71, 0, 541, 538, 1, 0, 0, 0, 541, 539, 1, 0, 0,
		0, 541, 540, 1, 0, 0, 0, 542, 138, 1, 0, 0, 0, 543, 547, 7, 3, 0, 0, 544,
		546, 3, 189, 94, 0, 545, 544, 1, 0, 0, 0, 546, 549, 1, 0, 0, 0, 547, 545,
		1, 0, 0, 0, 547, 548, 1, 0, 0, 0, 548, 140, 1, 0, 0, 0, 549, 547, 1, 0,
		0, 0, 550, 554, 5, 48, 0, 0, 551, 553, 3, 191, 95, 0, 552, 551, 1, 0, 0,
		0, 553, 556, 1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555,
		142, 1, 0, 0, 0, 556, 554, 1, 0, 0, 0, 557, 558, 5, 48, 0, 0, 558, 560,
		7, 4, 0, 0, 559, 561, 3, 193, 96, 0, 560, 559, 1, 0, 0, 0, 561, 562, 1,
		0, 0, 0, 562, 560, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 144, 1, 0, 0,
		0, 564, 568, 3, 147, 73, 0, 565, 566, 3, 153, 76, 0, 566, 567, 3, 147,
		73, 0, 567, 569, 1, 0, 0, 0, 568, 565, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0,
		569, 571, 1, 0, 0, 0, 570, 572, 3, 149, 74, 0, 571, 570, 1, 0, 0, 0, 571,
		572, 1, 0, 0, 0, 572, 579, 1, 0, 0, 0, 573, 574, 3, 153, 76, 0, 574, 576,
		3, 147, 73, 0, 575, 577, 3, 149, 74, 0, 576, 575, 1, 0, 0, 0, 576, 577,
		1, 0, 0, 0, 577, 579, 1, 0, 0, 0, 578, 564, 1, 0, 0, 0, 578, 573, 1, 0,
		0, 0, 579, 146, 1, 0, 0, 0, 580, 582, 3, 189, 94, 0, 581, 580, 1, 0, 0,
		0, 582, 583, 1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 583, 584, 1, 0, 0, 0, 584,
		148, 1, 0, 0, 0, 585, 587, 7, 5, 0, 0, 586, 588, 7, 6, 0, 0, 587, 586,
		1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588, 589, 1, 0, 0, 0, 589, 590, 3, 147,
		73, 0, 590, 150, 1, 0, 0, 0, 591, 594, 3, 147, 73, 0, 592, 594, 3, 145,
		72, 0, 593, 591, 1, 0, 0, 0, 593, 592, 1, 0, 0, 0, 594, 595, 1, 0, 0, 0,
		595, 596, 5, 105, 0, 0, 596, 152, 1, 0, 0, 0, 597, 598, 5, 46, 0, 0, 598,
		154, 1, 0, 0, 0, 599, 600, 3, 159, 79, 0, 600, 601, 3, 153, 76, 0, 601,
		602, 3, 159, 79, 0, 602, 156, 1, 0, 0, 0, 603, 604, 3, 159, 79, 0, 604,
		605, 3, 153, 76, 0, 605, 609, 3, 159, 79, 0, 606, 607, 3, 153, 76, 0, 607,
		608, 3, 159, 79, 0, 608, 610, 1, 0, 0, 0, 609, 606, 1, 0, 0, 0, 610, 611,
		1, 0, 0, 0, 611, 609, 1, 0, 0, 0, 611, 612, 1, 0, 0, 0, 612, 158, 1, 0,
		0, 0, 613, 618, 3, 199, 99, 0, 614, 617, 3, 187, 93, 0, 615, 617, 3, 197,
		98, 0, 616, 614, 1, 0, 0, 0, 616, 615, 1, 0, 0, 0, 617, 620, 1, 0, 0, 0,
		618, 616, 1, 0, 0, 0, 618, 619, 1, 0, 0, 0, 619, 160, 1, 0, 0, 0, 620,
		618, 1, 0, 0, 0, 621, 624, 5, 39, 0, 0, 622, 625, 3, 173, 86, 0, 623, 625,
		3, 179, 89, 0, 624, 622, 1, 0, 0, 0, 624, 623, 1, 0, 0, 0, 625, 626, 1,
		0, 0, 0, 626, 627, 5, 39, 0, 0, 627, 162, 1, 0, 0, 0, 628, 631, 3, 165,
		82, 0, 629, 631, 3, 167, 83, 0, 630, 628, 1, 0, 0, 0, 630, 629, 1, 0, 0,
		0, 631, 164, 1, 0, 0, 0, 632, 637, 5, 96, 0, 0, 633, 636, 3, 171, 85, 0,
		634, 636, 3, 73, 36, 0, 635, 633, 1, 0, 0, 0, 635, 634, 1, 0, 0, 0, 636,
		639, 1, 0, 0, 0, 637, 635, 1, 0, 0, 0, 637, 638, 1, 0, 0, 0, 638, 640,
		1, 0, 0, 0, 639, 637, 1, 0, 0, 0, 640, 641, 5, 96, 0, 0, 641, 166, 1, 0,
		0, 0, 642, 647, 5, 34, 0, 0, 643, 646, 3, 169, 84, 0, 644, 646, 3, 179,
		89, 0, 645, 643, 1, 0, 0, 0, 645, 644, 1, 0, 0, 0, 646, 649, 1, 0, 0, 0,
		647, 645, 1, 0, 0, 0, 647, 648, 1, 0, 0, 0, 648, 650, 1, 0, 0, 0, 649,
		647, 1, 0, 0, 0, 650, 651, 5, 34, 0, 0, 651, 168, 1, 0, 0, 0, 652, 657,
		8, 7, 0, 0, 653, 657, 3, 175, 87, 0, 654, 657, 3, 177, 88, 0, 655, 657,
		3, 185, 92, 0, 656, 652, 1, 0, 0, 0, 656, 653, 1, 0, 0, 0, 656, 654, 1,
		0, 0, 0, 656, 655, 1, 0, 0, 0, 657, 170, 1, 0, 0, 0, 658, 663, 8, 8, 0,
		0, 659, 663, 3, 175, 87, 0, 660, 663, 3, 177, 88, 0, 661, 663, 3, 185,
		92, 0, 662, 658, 1, 0, 0, 0, 662, 659, 1, 0, 0, 0, 662, 660, 1, 0, 0, 0,
		662, 661, 1, 0, 0, 0, 663, 172, 1, 0, 0, 0, 664, 669, 3, 195, 97, 0, 665,
		669, 3, 175, 87, 0, 666, 669, 3, 177, 88, 0, 667, 669, 3, 185, 92, 0, 668,
		664, 1, 0, 0, 0, 668, 665, 1, 0, 0, 0, 668, 666, 1, 0, 0, 0, 668, 667,
		1, 0, 0, 0, 669, 174, 1, 0, 0, 0, 670, 671, 5, 92, 0, 0, 671, 672, 5, 117,
		0, 0, 672, 673, 1, 0, 0, 0, 673, 674, 3, 193, 96, 0, 674, 675, 3, 193,
		96, 0, 675, 676, 3, 193, 96, 0, 676, 677, 3, 193, 96, 0, 677, 176, 1, 0,
		0, 0, 678, 679, 5, 92, 0, 0, 679, 680, 5, 85, 0, 0, 680, 681, 1, 0, 0,
		0, 681, 682, 3, 193, 96, 0, 682, 683, 3, 193, 96, 0, 683, 684, 3, 193,
		96, 0, 684, 685, 3, 193, 96, 0, 685, 686, 3, 193, 96, 0, 686, 687, 3, 193,
		96, 0, 687, 688, 3, 193, 96, 0, 688, 689, 3, 193, 96, 0, 689, 178, 1, 0,
		0, 0, 690, 693, 3, 181, 90, 0, 691, 693, 3, 183, 91, 0, 692, 690, 1, 0,
		0, 0, 692, 691, 1, 0, 0, 0, 693, 180, 1, 0, 0, 0, 694, 695, 5, 92, 0, 0,
		695, 696, 3, 191, 95, 0, 696, 697, 3, 191, 95, 0, 697, 698, 3, 191, 95,
		0, 698, 182, 1, 0, 0, 0, 699, 700, 5, 92, 0, 0, 700, 701, 5, 120, 0, 0,
		701, 702, 3, 193, 96, 0, 702, 703, 3, 193, 96, 0, 703, 184, 1, 0, 0, 0,
		704, 705, 5, 92, 0, 0, 705, 706, 7, 9, 0, 0, 706, 186, 1, 0, 0, 0, 707,
		710, 3, 199, 99, 0, 708, 710, 5, 95, 0, 0, 709, 707, 1, 0, 0, 0, 709, 708,
		1, 0, 0, 0, 710, 188, 1, 0, 0, 0, 711, 712, 7, 10, 0, 0, 712, 190, 1, 0,
		0, 0, 713, 714, 7, 11, 0, 0, 714, 192, 1, 0, 0, 0, 715, 716, 7, 12, 0,
		0, 716, 194, 1, 0, 0, 0, 717, 718, 8, 1, 0, 0, 718, 196, 1, 0, 0, 0, 719,
		721, 7, 13, 0, 0, 720, 719, 1, 0, 0, 0, 721, 198, 1, 0, 0, 0, 722, 724,
		7, 14, 0, 0, 723, 722, 1, 0, 0, 0, 724, 200, 1, 0, 0, 0, 725, 727, 7, 15,
		0, 0, 726, 725, 1, 0, 0, 0, 727, 728, 1, 0, 0, 0, 728, 726, 1, 0, 0, 0,
		728, 729, 1, 0, 0, 0, 729, 730, 1, 0, 0, 0, 730, 731, 6, 100, 0, 0, 731,
		202, 1, 0, 0, 0, 732, 733, 5, 47, 0, 0, 733, 734, 5, 42, 0, 0, 734, 738,
		1, 0, 0, 0, 735, 737, 9, 0, 0, 0, 736, 735, 1, 0, 0, 0, 737, 740, 1, 0,
		0, 0, 738, 739, 1, 0, 0, 0, 738, 736, 1, 0, 0, 0, 739, 741, 1, 0, 0, 0,
		740, 738, 1, 0, 0, 0, 741, 742, 5, 42, 0, 0, 742, 743, 5, 47, 0, 0, 743,
		744, 1, 0, 0, 0, 744, 745, 6, 101, 0, 0, 745, 204, 1, 0, 0, 0, 746, 747,
		5, 47, 0, 0, 747, 748, 5, 47, 0, 0, 748, 752, 1, 0, 0, 0, 749, 751, 8,
		16, 0, 0, 750, 749, 1, 0, 0, 0, 751, 754, 1, 0, 0, 0, 752, 750, 1, 0, 0,
		0, 752, 753, 1, 0, 0, 0, 753, 755, 1, 0, 0, 0, 754, 752, 1, 0, 0, 0, 755,
		765, 3, 67, 33, 0, 756, 760, 5, 35, 0, 0, 757, 759, 8, 16, 0, 0, 758, 757,
		1, 0, 0, 0, 759, 762, 1, 0, 0, 0, 760, 758, 1, 0, 0, 0, 760, 761, 1, 0,
		0, 0, 761, 763, 1, 0, 0, 0, 762, 760, 1, 0, 0, 0, 763, 765, 3, 67, 33,
		0, 764, 746, 1, 0, 0, 0, 764, 756, 1, 0, 0, 0, 765, 766, 1, 0, 0, 0, 766,
		767, 6, 102, 0, 0, 767, 206, 1, 0, 0, 0, 50, 0, 390, 394, 397, 401, 404,
		408, 411, 415, 420, 423, 426, 430, 433, 437, 440, 452, 536, 541, 547, 554,
		562, 568, 571, 576, 578, 583, 587, 593, 611, 616, 618, 624, 630, 635, 637,
		645, 647, 656, 662, 668, 692, 709, 720, 723, 728, 738, 752, 760, 764, 1,
		6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RuxlangLexerInit initializes any static state used to implement RuxlangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRuxlangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuxlangLexerInit() {
	staticData := &ruxlanglexerLexerStaticData
	staticData.once.Do(ruxlanglexerLexerInit)
}

// NewRuxlangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRuxlangLexer(input antlr.CharStream) *RuxlangLexer {
	RuxlangLexerInit()
	l := new(RuxlangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ruxlanglexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "RuxlangLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RuxlangLexer tokens.
const (
	RuxlangLexerPACKAGE              = 1
	RuxlangLexerIMPORT               = 2
	RuxlangLexerCONST                = 3
	RuxlangLexerMAP                  = 4
	RuxlangLexerPIPE                 = 5
	RuxlangLexerPIPEIN               = 6
	RuxlangLexerPIPEOUT              = 7
	RuxlangLexerINTERFACE            = 8
	RuxlangLexerSTRUCT               = 9
	RuxlangLexerTYPE                 = 10
	RuxlangLexerFUNC                 = 11
	RuxlangLexerVAR                  = 12
	RuxlangLexerRETURN               = 13
	RuxlangLexerYIELD                = 14
	RuxlangLexerBREAK                = 15
	RuxlangLexerCONTINUE             = 16
	RuxlangLexerGOTO                 = 17
	RuxlangLexerFALLTHROUGH          = 18
	RuxlangLexerDEFER                = 19
	RuxlangLexerIF                   = 20
	RuxlangLexerELSE                 = 21
	RuxlangLexerSWITCH               = 22
	RuxlangLexerSELECT               = 23
	RuxlangLexerCASE                 = 24
	RuxlangLexerDEFAULT              = 25
	RuxlangLexerFOR                  = 26
	RuxlangLexerIN                   = 27
	RuxlangLexerRUN                  = 28
	RuxlangLexerTYPE_UNWRAP          = 29
	RuxlangLexerPAREN_START          = 30
	RuxlangLexerPAREN_END            = 31
	RuxlangLexerBRACKET_START        = 32
	RuxlangLexerBRACKET_END          = 33
	RuxlangLexerSTMTEND              = 34
	RuxlangLexerBLOCK_START          = 35
	RuxlangLexerBLOCK_END            = 36
	RuxlangLexerASSIGN_WITH_OP       = 37
	RuxlangLexerASSIGN_NEW           = 38
	RuxlangLexerASSIGN               = 39
	RuxlangLexerAMP                  = 40
	RuxlangLexerCOLON                = 41
	RuxlangLexerCOMMA                = 42
	RuxlangLexerSEND                 = 43
	RuxlangLexerINC                  = 44
	RuxlangLexerDEC                  = 45
	RuxlangLexerARRAY_EXPAND         = 46
	RuxlangLexerSEQ_SEP              = 47
	RuxlangLexerLT                   = 48
	RuxlangLexerGT                   = 49
	RuxlangLexerLTE                  = 50
	RuxlangLexerGTE                  = 51
	RuxlangLexerEQ                   = 52
	RuxlangLexerNEQ                  = 53
	RuxlangLexerOR                   = 54
	RuxlangLexerAND                  = 55
	RuxlangLexerADD                  = 56
	RuxlangLexerSUB                  = 57
	RuxlangLexerMUL                  = 58
	RuxlangLexerDIV                  = 59
	RuxlangLexerMOD                  = 60
	RuxlangLexerBIT_OR               = 61
	RuxlangLexerCARET                = 62
	RuxlangLexerSHL                  = 63
	RuxlangLexerSHR                  = 64
	RuxlangLexerBIT_CLEAR            = 65
	RuxlangLexerNOT                  = 66
	RuxlangLexerUNARY_OP             = 67
	RuxlangLexerINT_LIT              = 68
	RuxlangLexerFLOAT_LIT            = 69
	RuxlangLexerIMAGINARY_LIT        = 70
	RuxlangLexerDOT                  = 71
	RuxlangLexerQUALIFIED_IDENTIFIER = 72
	RuxlangLexerEXTENDED_IDENTIFIER  = 73
	RuxlangLexerIDENTIFIER           = 74
	RuxlangLexerRUNE_LIT             = 75
	RuxlangLexerSTRING_LIT           = 76
	RuxlangLexerLITTLE_U_VALUE       = 77
	RuxlangLexerBIG_U_VALUE          = 78
	RuxlangLexerWS                   = 79
	RuxlangLexerCOMMENT              = 80
	RuxlangLexerLINE_COMMENT         = 81
)

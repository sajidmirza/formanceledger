// Code generated from NumScript.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NumScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var numscriptlexerLexerStaticData struct {
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

func numscriptlexerLexerInit() {
	staticData := &numscriptlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'*'", "'allowing overdraft up to'", "'allowing unbounded overdraft'",
		"','", "", "", "", "", "'vars'", "'meta'", "'set_tx_meta'", "'set_account_meta'",
		"'print'", "'fail'", "'send'", "'source'", "'from'", "'max'", "'destination'",
		"'to'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'", "'{'",
		"'}'", "'='", "'account'", "'asset'", "'number'", "'monetary'", "'portion'",
		"'string'", "", "", "'remaining'", "'kept'", "'balance'", "", "'%'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
		"VARS", "META", "SET_TX_META", "SET_ACCOUNT_META", "PRINT", "FAIL",
		"SEND", "SOURCE", "FROM", "MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD",
		"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
		"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
		"TY_STRING", "STRING", "PORTION", "REMAINING", "KEPT", "BALANCE", "NUMBER",
		"PERCENT", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT",
		"LINE_COMMENT", "VARS", "META", "SET_TX_META", "SET_ACCOUNT_META", "PRINT",
		"FAIL", "SEND", "SOURCE", "FROM", "MAX", "DESTINATION", "TO", "ALLOCATE",
		"OP_ADD", "OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE",
		"RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY",
		"TY_PORTION", "TY_STRING", "STRING", "PORTION", "REMAINING", "KEPT",
		"BALANCE", "NUMBER", "PERCENT", "VARIABLE_NAME", "ACCOUNT", "ASSET",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 450, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 4, 4, 153, 8, 4,
		11, 4, 12, 4, 154, 1, 5, 4, 5, 158, 8, 5, 11, 5, 12, 5, 159, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 169, 8, 6, 10, 6, 12, 6, 172, 9, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 183, 8, 7,
		10, 7, 12, 7, 186, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 5, 36, 352,
		8, 36, 10, 36, 12, 36, 355, 9, 36, 1, 36, 1, 36, 1, 37, 4, 37, 360, 8,
		37, 11, 37, 12, 37, 361, 1, 37, 3, 37, 365, 8, 37, 1, 37, 1, 37, 3, 37,
		369, 8, 37, 1, 37, 4, 37, 372, 8, 37, 11, 37, 12, 37, 373, 1, 37, 4, 37,
		377, 8, 37, 11, 37, 12, 37, 378, 1, 37, 1, 37, 4, 37, 383, 8, 37, 11, 37,
		12, 37, 384, 3, 37, 387, 8, 37, 1, 37, 3, 37, 390, 8, 37, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 41, 4, 41, 416, 8, 41, 11, 41, 12, 41, 417, 1, 42, 1, 42, 1, 43, 1,
		43, 4, 43, 424, 8, 43, 11, 43, 12, 43, 425, 1, 43, 5, 43, 429, 8, 43, 10,
		43, 12, 43, 432, 9, 43, 1, 44, 1, 44, 4, 44, 436, 8, 44, 11, 44, 12, 44,
		437, 1, 44, 5, 44, 441, 8, 44, 10, 44, 12, 44, 444, 9, 44, 1, 45, 4, 45,
		447, 8, 45, 11, 45, 12, 45, 448, 2, 170, 184, 0, 46, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40,
		81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 1, 0, 10, 2, 0, 10, 10,
		13, 13, 2, 0, 9, 9, 32, 32, 6, 0, 32, 32, 45, 45, 48, 57, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 1, 0, 32, 32, 2, 0, 95, 95, 97, 122, 3, 0, 48, 57,
		95, 95, 97, 122, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 58, 65, 90, 95,
		95, 97, 122, 2, 0, 47, 57, 65, 90, 469, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 1, 93, 1, 0, 0, 0, 3, 95, 1, 0, 0, 0,
		5, 120, 1, 0, 0, 0, 7, 149, 1, 0, 0, 0, 9, 152, 1, 0, 0, 0, 11, 157, 1,
		0, 0, 0, 13, 163, 1, 0, 0, 0, 15, 178, 1, 0, 0, 0, 17, 191, 1, 0, 0, 0,
		19, 196, 1, 0, 0, 0, 21, 201, 1, 0, 0, 0, 23, 213, 1, 0, 0, 0, 25, 230,
		1, 0, 0, 0, 27, 236, 1, 0, 0, 0, 29, 241, 1, 0, 0, 0, 31, 246, 1, 0, 0,
		0, 33, 253, 1, 0, 0, 0, 35, 258, 1, 0, 0, 0, 37, 262, 1, 0, 0, 0, 39, 274,
		1, 0, 0, 0, 41, 277, 1, 0, 0, 0, 43, 286, 1, 0, 0, 0, 45, 288, 1, 0, 0,
		0, 47, 290, 1, 0, 0, 0, 49, 292, 1, 0, 0, 0, 51, 294, 1, 0, 0, 0, 53, 296,
		1, 0, 0, 0, 55, 298, 1, 0, 0, 0, 57, 300, 1, 0, 0, 0, 59, 302, 1, 0, 0,
		0, 61, 304, 1, 0, 0, 0, 63, 312, 1, 0, 0, 0, 65, 318, 1, 0, 0, 0, 67, 325,
		1, 0, 0, 0, 69, 334, 1, 0, 0, 0, 71, 342, 1, 0, 0, 0, 73, 349, 1, 0, 0,
		0, 75, 389, 1, 0, 0, 0, 77, 391, 1, 0, 0, 0, 79, 401, 1, 0, 0, 0, 81, 406,
		1, 0, 0, 0, 83, 415, 1, 0, 0, 0, 85, 419, 1, 0, 0, 0, 87, 421, 1, 0, 0,
		0, 89, 433, 1, 0, 0, 0, 91, 446, 1, 0, 0, 0, 93, 94, 5, 42, 0, 0, 94, 2,
		1, 0, 0, 0, 95, 96, 5, 97, 0, 0, 96, 97, 5, 108, 0, 0, 97, 98, 5, 108,
		0, 0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 119, 0, 0, 100, 101, 5, 105, 0,
		0, 101, 102, 5, 110, 0, 0, 102, 103, 5, 103, 0, 0, 103, 104, 5, 32, 0,
		0, 104, 105, 5, 111, 0, 0, 105, 106, 5, 118, 0, 0, 106, 107, 5, 101, 0,
		0, 107, 108, 5, 114, 0, 0, 108, 109, 5, 100, 0, 0, 109, 110, 5, 114, 0,
		0, 110, 111, 5, 97, 0, 0, 111, 112, 5, 102, 0, 0, 112, 113, 5, 116, 0,
		0, 113, 114, 5, 32, 0, 0, 114, 115, 5, 117, 0, 0, 115, 116, 5, 112, 0,
		0, 116, 117, 5, 32, 0, 0, 117, 118, 5, 116, 0, 0, 118, 119, 5, 111, 0,
		0, 119, 4, 1, 0, 0, 0, 120, 121, 5, 97, 0, 0, 121, 122, 5, 108, 0, 0, 122,
		123, 5, 108, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125, 5, 119, 0, 0, 125,
		126, 5, 105, 0, 0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 103, 0, 0, 128,
		129, 5, 32, 0, 0, 129, 130, 5, 117, 0, 0, 130, 131, 5, 110, 0, 0, 131,
		132, 5, 98, 0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 117, 0, 0, 134,
		135, 5, 110, 0, 0, 135, 136, 5, 100, 0, 0, 136, 137, 5, 101, 0, 0, 137,
		138, 5, 100, 0, 0, 138, 139, 5, 32, 0, 0, 139, 140, 5, 111, 0, 0, 140,
		141, 5, 118, 0, 0, 141, 142, 5, 101, 0, 0, 142, 143, 5, 114, 0, 0, 143,
		144, 5, 100, 0, 0, 144, 145, 5, 114, 0, 0, 145, 146, 5, 97, 0, 0, 146,
		147, 5, 102, 0, 0, 147, 148, 5, 116, 0, 0, 148, 6, 1, 0, 0, 0, 149, 150,
		5, 44, 0, 0, 150, 8, 1, 0, 0, 0, 151, 153, 7, 0, 0, 0, 152, 151, 1, 0,
		0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 10, 1, 0, 0, 0, 156, 158, 7, 1, 0, 0, 157, 156, 1, 0, 0, 0, 158, 159,
		1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0,
		0, 0, 161, 162, 6, 5, 0, 0, 162, 12, 1, 0, 0, 0, 163, 164, 5, 47, 0, 0,
		164, 165, 5, 42, 0, 0, 165, 170, 1, 0, 0, 0, 166, 169, 3, 13, 6, 0, 167,
		169, 9, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 172,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 173, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 42, 0, 0, 174, 175, 5, 47, 0,
		0, 175, 176, 1, 0, 0, 0, 176, 177, 6, 6, 0, 0, 177, 14, 1, 0, 0, 0, 178,
		179, 5, 47, 0, 0, 179, 180, 5, 47, 0, 0, 180, 184, 1, 0, 0, 0, 181, 183,
		9, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 185, 1, 0,
		0, 0, 184, 182, 1, 0, 0, 0, 185, 187, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0,
		187, 188, 3, 9, 4, 0, 188, 189, 1, 0, 0, 0, 189, 190, 6, 7, 0, 0, 190,
		16, 1, 0, 0, 0, 191, 192, 5, 118, 0, 0, 192, 193, 5, 97, 0, 0, 193, 194,
		5, 114, 0, 0, 194, 195, 5, 115, 0, 0, 195, 18, 1, 0, 0, 0, 196, 197, 5,
		109, 0, 0, 197, 198, 5, 101, 0, 0, 198, 199, 5, 116, 0, 0, 199, 200, 5,
		97, 0, 0, 200, 20, 1, 0, 0, 0, 201, 202, 5, 115, 0, 0, 202, 203, 5, 101,
		0, 0, 203, 204, 5, 116, 0, 0, 204, 205, 5, 95, 0, 0, 205, 206, 5, 116,
		0, 0, 206, 207, 5, 120, 0, 0, 207, 208, 5, 95, 0, 0, 208, 209, 5, 109,
		0, 0, 209, 210, 5, 101, 0, 0, 210, 211, 5, 116, 0, 0, 211, 212, 5, 97,
		0, 0, 212, 22, 1, 0, 0, 0, 213, 214, 5, 115, 0, 0, 214, 215, 5, 101, 0,
		0, 215, 216, 5, 116, 0, 0, 216, 217, 5, 95, 0, 0, 217, 218, 5, 97, 0, 0,
		218, 219, 5, 99, 0, 0, 219, 220, 5, 99, 0, 0, 220, 221, 5, 111, 0, 0, 221,
		222, 5, 117, 0, 0, 222, 223, 5, 110, 0, 0, 223, 224, 5, 116, 0, 0, 224,
		225, 5, 95, 0, 0, 225, 226, 5, 109, 0, 0, 226, 227, 5, 101, 0, 0, 227,
		228, 5, 116, 0, 0, 228, 229, 5, 97, 0, 0, 229, 24, 1, 0, 0, 0, 230, 231,
		5, 112, 0, 0, 231, 232, 5, 114, 0, 0, 232, 233, 5, 105, 0, 0, 233, 234,
		5, 110, 0, 0, 234, 235, 5, 116, 0, 0, 235, 26, 1, 0, 0, 0, 236, 237, 5,
		102, 0, 0, 237, 238, 5, 97, 0, 0, 238, 239, 5, 105, 0, 0, 239, 240, 5,
		108, 0, 0, 240, 28, 1, 0, 0, 0, 241, 242, 5, 115, 0, 0, 242, 243, 5, 101,
		0, 0, 243, 244, 5, 110, 0, 0, 244, 245, 5, 100, 0, 0, 245, 30, 1, 0, 0,
		0, 246, 247, 5, 115, 0, 0, 247, 248, 5, 111, 0, 0, 248, 249, 5, 117, 0,
		0, 249, 250, 5, 114, 0, 0, 250, 251, 5, 99, 0, 0, 251, 252, 5, 101, 0,
		0, 252, 32, 1, 0, 0, 0, 253, 254, 5, 102, 0, 0, 254, 255, 5, 114, 0, 0,
		255, 256, 5, 111, 0, 0, 256, 257, 5, 109, 0, 0, 257, 34, 1, 0, 0, 0, 258,
		259, 5, 109, 0, 0, 259, 260, 5, 97, 0, 0, 260, 261, 5, 120, 0, 0, 261,
		36, 1, 0, 0, 0, 262, 263, 5, 100, 0, 0, 263, 264, 5, 101, 0, 0, 264, 265,
		5, 115, 0, 0, 265, 266, 5, 116, 0, 0, 266, 267, 5, 105, 0, 0, 267, 268,
		5, 110, 0, 0, 268, 269, 5, 97, 0, 0, 269, 270, 5, 116, 0, 0, 270, 271,
		5, 105, 0, 0, 271, 272, 5, 111, 0, 0, 272, 273, 5, 110, 0, 0, 273, 38,
		1, 0, 0, 0, 274, 275, 5, 116, 0, 0, 275, 276, 5, 111, 0, 0, 276, 40, 1,
		0, 0, 0, 277, 278, 5, 97, 0, 0, 278, 279, 5, 108, 0, 0, 279, 280, 5, 108,
		0, 0, 280, 281, 5, 111, 0, 0, 281, 282, 5, 99, 0, 0, 282, 283, 5, 97, 0,
		0, 283, 284, 5, 116, 0, 0, 284, 285, 5, 101, 0, 0, 285, 42, 1, 0, 0, 0,
		286, 287, 5, 43, 0, 0, 287, 44, 1, 0, 0, 0, 288, 289, 5, 45, 0, 0, 289,
		46, 1, 0, 0, 0, 290, 291, 5, 40, 0, 0, 291, 48, 1, 0, 0, 0, 292, 293, 5,
		41, 0, 0, 293, 50, 1, 0, 0, 0, 294, 295, 5, 91, 0, 0, 295, 52, 1, 0, 0,
		0, 296, 297, 5, 93, 0, 0, 297, 54, 1, 0, 0, 0, 298, 299, 5, 123, 0, 0,
		299, 56, 1, 0, 0, 0, 300, 301, 5, 125, 0, 0, 301, 58, 1, 0, 0, 0, 302,
		303, 5, 61, 0, 0, 303, 60, 1, 0, 0, 0, 304, 305, 5, 97, 0, 0, 305, 306,
		5, 99, 0, 0, 306, 307, 5, 99, 0, 0, 307, 308, 5, 111, 0, 0, 308, 309, 5,
		117, 0, 0, 309, 310, 5, 110, 0, 0, 310, 311, 5, 116, 0, 0, 311, 62, 1,
		0, 0, 0, 312, 313, 5, 97, 0, 0, 313, 314, 5, 115, 0, 0, 314, 315, 5, 115,
		0, 0, 315, 316, 5, 101, 0, 0, 316, 317, 5, 116, 0, 0, 317, 64, 1, 0, 0,
		0, 318, 319, 5, 110, 0, 0, 319, 320, 5, 117, 0, 0, 320, 321, 5, 109, 0,
		0, 321, 322, 5, 98, 0, 0, 322, 323, 5, 101, 0, 0, 323, 324, 5, 114, 0,
		0, 324, 66, 1, 0, 0, 0, 325, 326, 5, 109, 0, 0, 326, 327, 5, 111, 0, 0,
		327, 328, 5, 110, 0, 0, 328, 329, 5, 101, 0, 0, 329, 330, 5, 116, 0, 0,
		330, 331, 5, 97, 0, 0, 331, 332, 5, 114, 0, 0, 332, 333, 5, 121, 0, 0,
		333, 68, 1, 0, 0, 0, 334, 335, 5, 112, 0, 0, 335, 336, 5, 111, 0, 0, 336,
		337, 5, 114, 0, 0, 337, 338, 5, 116, 0, 0, 338, 339, 5, 105, 0, 0, 339,
		340, 5, 111, 0, 0, 340, 341, 5, 110, 0, 0, 341, 70, 1, 0, 0, 0, 342, 343,
		5, 115, 0, 0, 343, 344, 5, 116, 0, 0, 344, 345, 5, 114, 0, 0, 345, 346,
		5, 105, 0, 0, 346, 347, 5, 110, 0, 0, 347, 348, 5, 103, 0, 0, 348, 72,
		1, 0, 0, 0, 349, 353, 5, 34, 0, 0, 350, 352, 7, 2, 0, 0, 351, 350, 1, 0,
		0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0,
		354, 356, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356, 357, 5, 34, 0, 0, 357,
		74, 1, 0, 0, 0, 358, 360, 7, 3, 0, 0, 359, 358, 1, 0, 0, 0, 360, 361, 1,
		0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 364, 1, 0, 0,
		0, 363, 365, 7, 4, 0, 0, 364, 363, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365,
		366, 1, 0, 0, 0, 366, 368, 5, 47, 0, 0, 367, 369, 7, 4, 0, 0, 368, 367,
		1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 371, 1, 0, 0, 0, 370, 372, 7, 3,
		0, 0, 371, 370, 1, 0, 0, 0, 372, 373, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0,
		373, 374, 1, 0, 0, 0, 374, 390, 1, 0, 0, 0, 375, 377, 7, 3, 0, 0, 376,
		375, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 376, 1, 0, 0, 0, 378, 379,
		1, 0, 0, 0, 379, 386, 1, 0, 0, 0, 380, 382, 5, 46, 0, 0, 381, 383, 7, 3,
		0, 0, 382, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0,
		384, 385, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386, 380, 1, 0, 0, 0, 386,
		387, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 390, 5, 37, 0, 0, 389, 359,
		1, 0, 0, 0, 389, 376, 1, 0, 0, 0, 390, 76, 1, 0, 0, 0, 391, 392, 5, 114,
		0, 0, 392, 393, 5, 101, 0, 0, 393, 394, 5, 109, 0, 0, 394, 395, 5, 97,
		0, 0, 395, 396, 5, 105, 0, 0, 396, 397, 5, 110, 0, 0, 397, 398, 5, 105,
		0, 0, 398, 399, 5, 110, 0, 0, 399, 400, 5, 103, 0, 0, 400, 78, 1, 0, 0,
		0, 401, 402, 5, 107, 0, 0, 402, 403, 5, 101, 0, 0, 403, 404, 5, 112, 0,
		0, 404, 405, 5, 116, 0, 0, 405, 80, 1, 0, 0, 0, 406, 407, 5, 98, 0, 0,
		407, 408, 5, 97, 0, 0, 408, 409, 5, 108, 0, 0, 409, 410, 5, 97, 0, 0, 410,
		411, 5, 110, 0, 0, 411, 412, 5, 99, 0, 0, 412, 413, 5, 101, 0, 0, 413,
		82, 1, 0, 0, 0, 414, 416, 7, 3, 0, 0, 415, 414, 1, 0, 0, 0, 416, 417, 1,
		0, 0, 0, 417, 415, 1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 84, 1, 0, 0,
		0, 419, 420, 5, 37, 0, 0, 420, 86, 1, 0, 0, 0, 421, 423, 5, 36, 0, 0, 422,
		424, 7, 5, 0, 0, 423, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 423,
		1, 0, 0, 0, 425, 426, 1, 0, 0, 0, 426, 430, 1, 0, 0, 0, 427, 429, 7, 6,
		0, 0, 428, 427, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0,
		430, 431, 1, 0, 0, 0, 431, 88, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 435,
		5, 64, 0, 0, 434, 436, 7, 7, 0, 0, 435, 434, 1, 0, 0, 0, 436, 437, 1, 0,
		0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 1, 0, 0, 0, 438, 442, 1, 0, 0, 0,
		439, 441, 7, 8, 0, 0, 440, 439, 1, 0, 0, 0, 441, 444, 1, 0, 0, 0, 442,
		440, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 90, 1, 0, 0, 0, 444, 442, 1,
		0, 0, 0, 445, 447, 7, 9, 0, 0, 446, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0,
		0, 448, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 92, 1, 0, 0, 0, 21,
		0, 154, 159, 168, 170, 184, 353, 361, 364, 368, 373, 378, 384, 386, 389,
		417, 425, 430, 437, 442, 448, 1, 6, 0, 0,
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

// NumScriptLexerInit initializes any static state used to implement NumScriptLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNumScriptLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumScriptLexerInit() {
	staticData := &numscriptlexerLexerStaticData
	staticData.once.Do(numscriptlexerLexerInit)
}

// NewNumScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNumScriptLexer(input antlr.CharStream) *NumScriptLexer {
	NumScriptLexerInit()
	l := new(NumScriptLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &numscriptlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "NumScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumScriptLexer tokens.
const (
	NumScriptLexerT__0              = 1
	NumScriptLexerT__1              = 2
	NumScriptLexerT__2              = 3
	NumScriptLexerT__3              = 4
	NumScriptLexerNEWLINE           = 5
	NumScriptLexerWHITESPACE        = 6
	NumScriptLexerMULTILINE_COMMENT = 7
	NumScriptLexerLINE_COMMENT      = 8
	NumScriptLexerVARS              = 9
	NumScriptLexerMETA              = 10
	NumScriptLexerSET_TX_META       = 11
	NumScriptLexerSET_ACCOUNT_META  = 12
	NumScriptLexerPRINT             = 13
	NumScriptLexerFAIL              = 14
	NumScriptLexerSEND              = 15
	NumScriptLexerSOURCE            = 16
	NumScriptLexerFROM              = 17
	NumScriptLexerMAX               = 18
	NumScriptLexerDESTINATION       = 19
	NumScriptLexerTO                = 20
	NumScriptLexerALLOCATE          = 21
	NumScriptLexerOP_ADD            = 22
	NumScriptLexerOP_SUB            = 23
	NumScriptLexerLPAREN            = 24
	NumScriptLexerRPAREN            = 25
	NumScriptLexerLBRACK            = 26
	NumScriptLexerRBRACK            = 27
	NumScriptLexerLBRACE            = 28
	NumScriptLexerRBRACE            = 29
	NumScriptLexerEQ                = 30
	NumScriptLexerTY_ACCOUNT        = 31
	NumScriptLexerTY_ASSET          = 32
	NumScriptLexerTY_NUMBER         = 33
	NumScriptLexerTY_MONETARY       = 34
	NumScriptLexerTY_PORTION        = 35
	NumScriptLexerTY_STRING         = 36
	NumScriptLexerSTRING            = 37
	NumScriptLexerPORTION           = 38
	NumScriptLexerREMAINING         = 39
	NumScriptLexerKEPT              = 40
	NumScriptLexerBALANCE           = 41
	NumScriptLexerNUMBER            = 42
	NumScriptLexerPERCENT           = 43
	NumScriptLexerVARIABLE_NAME     = 44
	NumScriptLexerACCOUNT           = 45
	NumScriptLexerASSET             = 46
)

// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser
import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type LustreLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var LustreLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func lustrelexerLexerInit() {
  staticData := &LustreLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'type'", "'='", "';'", "'const'", "':'", "'node'", "'('", "')'", 
    "'returns'", "'var'", "'let'", "'tel'", "'function'", "','", "'struct'", 
    "'{'", "'}'", "'enum'", "'int'", "'subrange'", "'['", "']'", "'of'", 
    "'bool'", "'real'", "'-'", "'--%PROPERTY'", "'--%REALIZABLE'", "'--%IVC'", 
    "'--%MAIN'", "'assert'", "'floor'", "'condact'", "'.'", "':='", "'pre'", 
    "'not'", "'*'", "'/'", "'div'", "'mod'", "'+'", "'<'", "'<='", "'>'", 
    "'>='", "'<>'", "'and'", "'or'", "'xor'", "'=>'", "'->'", "'if'", "'then'", 
    "'else'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "REAL", "BOOL", "INT", "ID", "WS", "SL_COMMENT", 
    "ML_COMMENT", "ERROR",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
    "T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
    "T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
    "T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48", 
    "T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "REAL", "BOOL", 
    "INT", "ID", "WS", "SL_COMMENT", "ML_COMMENT", "ERROR",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 63, 435, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
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
	62, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 
	6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 
	9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 
	1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 
	13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 
	1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 
	19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 
	1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 
	24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 
	1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 
	27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 
	1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 
	29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 
	1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 
	32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 
	1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 
	38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 
	1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 
	45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 
	1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 
	52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 
	1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 
	56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 378, 8, 56, 1, 57, 4, 57, 381, 8, 
	57, 11, 57, 12, 57, 382, 1, 58, 1, 58, 5, 58, 387, 8, 58, 10, 58, 12, 58, 
	390, 9, 58, 1, 59, 4, 59, 393, 8, 59, 11, 59, 12, 59, 394, 1, 59, 1, 59, 
	1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 404, 8, 60, 10, 60, 12, 60, 407, 
	9, 60, 1, 60, 3, 60, 410, 8, 60, 1, 60, 3, 60, 413, 8, 60, 1, 60, 3, 60, 
	416, 8, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 5, 61, 424, 8, 61, 
	10, 61, 12, 61, 427, 9, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 
	62, 1, 425, 0, 63, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 
	9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 
	18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 
	27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 
	36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 
	45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 
	107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 
	123, 62, 125, 63, 1, 0, 6, 1, 0, 48, 57, 5, 0, 33, 33, 65, 90, 95, 95, 
	97, 122, 126, 126, 6, 0, 33, 33, 48, 57, 65, 90, 95, 95, 97, 122, 126, 
	126, 3, 0, 9, 10, 12, 13, 32, 32, 3, 0, 10, 10, 13, 13, 37, 37, 2, 0, 10, 
	10, 13, 13, 443, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 
	0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 
	0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 
	0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 
	0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 
	1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 
	45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 
	0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 
	0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 
	0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 
	0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 
	1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 
	91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 
	0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 
	0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 
	1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 
	0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 1, 127, 1, 
	0, 0, 0, 3, 132, 1, 0, 0, 0, 5, 134, 1, 0, 0, 0, 7, 136, 1, 0, 0, 0, 9, 
	142, 1, 0, 0, 0, 11, 144, 1, 0, 0, 0, 13, 149, 1, 0, 0, 0, 15, 151, 1, 
	0, 0, 0, 17, 153, 1, 0, 0, 0, 19, 161, 1, 0, 0, 0, 21, 165, 1, 0, 0, 0, 
	23, 169, 1, 0, 0, 0, 25, 173, 1, 0, 0, 0, 27, 182, 1, 0, 0, 0, 29, 184, 
	1, 0, 0, 0, 31, 191, 1, 0, 0, 0, 33, 193, 1, 0, 0, 0, 35, 195, 1, 0, 0, 
	0, 37, 200, 1, 0, 0, 0, 39, 204, 1, 0, 0, 0, 41, 213, 1, 0, 0, 0, 43, 215, 
	1, 0, 0, 0, 45, 217, 1, 0, 0, 0, 47, 220, 1, 0, 0, 0, 49, 225, 1, 0, 0, 
	0, 51, 230, 1, 0, 0, 0, 53, 232, 1, 0, 0, 0, 55, 244, 1, 0, 0, 0, 57, 258, 
	1, 0, 0, 0, 59, 265, 1, 0, 0, 0, 61, 273, 1, 0, 0, 0, 63, 280, 1, 0, 0, 
	0, 65, 286, 1, 0, 0, 0, 67, 294, 1, 0, 0, 0, 69, 296, 1, 0, 0, 0, 71, 299, 
	1, 0, 0, 0, 73, 303, 1, 0, 0, 0, 75, 307, 1, 0, 0, 0, 77, 309, 1, 0, 0, 
	0, 79, 311, 1, 0, 0, 0, 81, 315, 1, 0, 0, 0, 83, 319, 1, 0, 0, 0, 85, 321, 
	1, 0, 0, 0, 87, 323, 1, 0, 0, 0, 89, 326, 1, 0, 0, 0, 91, 328, 1, 0, 0, 
	0, 93, 331, 1, 0, 0, 0, 95, 334, 1, 0, 0, 0, 97, 338, 1, 0, 0, 0, 99, 341, 
	1, 0, 0, 0, 101, 345, 1, 0, 0, 0, 103, 348, 1, 0, 0, 0, 105, 351, 1, 0, 
	0, 0, 107, 354, 1, 0, 0, 0, 109, 359, 1, 0, 0, 0, 111, 364, 1, 0, 0, 0, 
	113, 377, 1, 0, 0, 0, 115, 380, 1, 0, 0, 0, 117, 384, 1, 0, 0, 0, 119, 
	392, 1, 0, 0, 0, 121, 398, 1, 0, 0, 0, 123, 419, 1, 0, 0, 0, 125, 433, 
	1, 0, 0, 0, 127, 128, 5, 116, 0, 0, 128, 129, 5, 121, 0, 0, 129, 130, 5, 
	112, 0, 0, 130, 131, 5, 101, 0, 0, 131, 2, 1, 0, 0, 0, 132, 133, 5, 61, 
	0, 0, 133, 4, 1, 0, 0, 0, 134, 135, 5, 59, 0, 0, 135, 6, 1, 0, 0, 0, 136, 
	137, 5, 99, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139, 5, 110, 0, 0, 139, 
	140, 5, 115, 0, 0, 140, 141, 5, 116, 0, 0, 141, 8, 1, 0, 0, 0, 142, 143, 
	5, 58, 0, 0, 143, 10, 1, 0, 0, 0, 144, 145, 5, 110, 0, 0, 145, 146, 5, 
	111, 0, 0, 146, 147, 5, 100, 0, 0, 147, 148, 5, 101, 0, 0, 148, 12, 1, 
	0, 0, 0, 149, 150, 5, 40, 0, 0, 150, 14, 1, 0, 0, 0, 151, 152, 5, 41, 0, 
	0, 152, 16, 1, 0, 0, 0, 153, 154, 5, 114, 0, 0, 154, 155, 5, 101, 0, 0, 
	155, 156, 5, 116, 0, 0, 156, 157, 5, 117, 0, 0, 157, 158, 5, 114, 0, 0, 
	158, 159, 5, 110, 0, 0, 159, 160, 5, 115, 0, 0, 160, 18, 1, 0, 0, 0, 161, 
	162, 5, 118, 0, 0, 162, 163, 5, 97, 0, 0, 163, 164, 5, 114, 0, 0, 164, 
	20, 1, 0, 0, 0, 165, 166, 5, 108, 0, 0, 166, 167, 5, 101, 0, 0, 167, 168, 
	5, 116, 0, 0, 168, 22, 1, 0, 0, 0, 169, 170, 5, 116, 0, 0, 170, 171, 5, 
	101, 0, 0, 171, 172, 5, 108, 0, 0, 172, 24, 1, 0, 0, 0, 173, 174, 5, 102, 
	0, 0, 174, 175, 5, 117, 0, 0, 175, 176, 5, 110, 0, 0, 176, 177, 5, 99, 
	0, 0, 177, 178, 5, 116, 0, 0, 178, 179, 5, 105, 0, 0, 179, 180, 5, 111, 
	0, 0, 180, 181, 5, 110, 0, 0, 181, 26, 1, 0, 0, 0, 182, 183, 5, 44, 0, 
	0, 183, 28, 1, 0, 0, 0, 184, 185, 5, 115, 0, 0, 185, 186, 5, 116, 0, 0, 
	186, 187, 5, 114, 0, 0, 187, 188, 5, 117, 0, 0, 188, 189, 5, 99, 0, 0, 
	189, 190, 5, 116, 0, 0, 190, 30, 1, 0, 0, 0, 191, 192, 5, 123, 0, 0, 192, 
	32, 1, 0, 0, 0, 193, 194, 5, 125, 0, 0, 194, 34, 1, 0, 0, 0, 195, 196, 
	5, 101, 0, 0, 196, 197, 5, 110, 0, 0, 197, 198, 5, 117, 0, 0, 198, 199, 
	5, 109, 0, 0, 199, 36, 1, 0, 0, 0, 200, 201, 5, 105, 0, 0, 201, 202, 5, 
	110, 0, 0, 202, 203, 5, 116, 0, 0, 203, 38, 1, 0, 0, 0, 204, 205, 5, 115, 
	0, 0, 205, 206, 5, 117, 0, 0, 206, 207, 5, 98, 0, 0, 207, 208, 5, 114, 
	0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 110, 0, 0, 210, 211, 5, 103, 
	0, 0, 211, 212, 5, 101, 0, 0, 212, 40, 1, 0, 0, 0, 213, 214, 5, 91, 0, 
	0, 214, 42, 1, 0, 0, 0, 215, 216, 5, 93, 0, 0, 216, 44, 1, 0, 0, 0, 217, 
	218, 5, 111, 0, 0, 218, 219, 5, 102, 0, 0, 219, 46, 1, 0, 0, 0, 220, 221, 
	5, 98, 0, 0, 221, 222, 5, 111, 0, 0, 222, 223, 5, 111, 0, 0, 223, 224, 
	5, 108, 0, 0, 224, 48, 1, 0, 0, 0, 225, 226, 5, 114, 0, 0, 226, 227, 5, 
	101, 0, 0, 227, 228, 5, 97, 0, 0, 228, 229, 5, 108, 0, 0, 229, 50, 1, 0, 
	0, 0, 230, 231, 5, 45, 0, 0, 231, 52, 1, 0, 0, 0, 232, 233, 5, 45, 0, 0, 
	233, 234, 5, 45, 0, 0, 234, 235, 5, 37, 0, 0, 235, 236, 5, 80, 0, 0, 236, 
	237, 5, 82, 0, 0, 237, 238, 5, 79, 0, 0, 238, 239, 5, 80, 0, 0, 239, 240, 
	5, 69, 0, 0, 240, 241, 5, 82, 0, 0, 241, 242, 5, 84, 0, 0, 242, 243, 5, 
	89, 0, 0, 243, 54, 1, 0, 0, 0, 244, 245, 5, 45, 0, 0, 245, 246, 5, 45, 
	0, 0, 246, 247, 5, 37, 0, 0, 247, 248, 5, 82, 0, 0, 248, 249, 5, 69, 0, 
	0, 249, 250, 5, 65, 0, 0, 250, 251, 5, 76, 0, 0, 251, 252, 5, 73, 0, 0, 
	252, 253, 5, 90, 0, 0, 253, 254, 5, 65, 0, 0, 254, 255, 5, 66, 0, 0, 255, 
	256, 5, 76, 0, 0, 256, 257, 5, 69, 0, 0, 257, 56, 1, 0, 0, 0, 258, 259, 
	5, 45, 0, 0, 259, 260, 5, 45, 0, 0, 260, 261, 5, 37, 0, 0, 261, 262, 5, 
	73, 0, 0, 262, 263, 5, 86, 0, 0, 263, 264, 5, 67, 0, 0, 264, 58, 1, 0, 
	0, 0, 265, 266, 5, 45, 0, 0, 266, 267, 5, 45, 0, 0, 267, 268, 5, 37, 0, 
	0, 268, 269, 5, 77, 0, 0, 269, 270, 5, 65, 0, 0, 270, 271, 5, 73, 0, 0, 
	271, 272, 5, 78, 0, 0, 272, 60, 1, 0, 0, 0, 273, 274, 5, 97, 0, 0, 274, 
	275, 5, 115, 0, 0, 275, 276, 5, 115, 0, 0, 276, 277, 5, 101, 0, 0, 277, 
	278, 5, 114, 0, 0, 278, 279, 5, 116, 0, 0, 279, 62, 1, 0, 0, 0, 280, 281, 
	5, 102, 0, 0, 281, 282, 5, 108, 0, 0, 282, 283, 5, 111, 0, 0, 283, 284, 
	5, 111, 0, 0, 284, 285, 5, 114, 0, 0, 285, 64, 1, 0, 0, 0, 286, 287, 5, 
	99, 0, 0, 287, 288, 5, 111, 0, 0, 288, 289, 5, 110, 0, 0, 289, 290, 5, 
	100, 0, 0, 290, 291, 5, 97, 0, 0, 291, 292, 5, 99, 0, 0, 292, 293, 5, 116, 
	0, 0, 293, 66, 1, 0, 0, 0, 294, 295, 5, 46, 0, 0, 295, 68, 1, 0, 0, 0, 
	296, 297, 5, 58, 0, 0, 297, 298, 5, 61, 0, 0, 298, 70, 1, 0, 0, 0, 299, 
	300, 5, 112, 0, 0, 300, 301, 5, 114, 0, 0, 301, 302, 5, 101, 0, 0, 302, 
	72, 1, 0, 0, 0, 303, 304, 5, 110, 0, 0, 304, 305, 5, 111, 0, 0, 305, 306, 
	5, 116, 0, 0, 306, 74, 1, 0, 0, 0, 307, 308, 5, 42, 0, 0, 308, 76, 1, 0, 
	0, 0, 309, 310, 5, 47, 0, 0, 310, 78, 1, 0, 0, 0, 311, 312, 5, 100, 0, 
	0, 312, 313, 5, 105, 0, 0, 313, 314, 5, 118, 0, 0, 314, 80, 1, 0, 0, 0, 
	315, 316, 5, 109, 0, 0, 316, 317, 5, 111, 0, 0, 317, 318, 5, 100, 0, 0, 
	318, 82, 1, 0, 0, 0, 319, 320, 5, 43, 0, 0, 320, 84, 1, 0, 0, 0, 321, 322, 
	5, 60, 0, 0, 322, 86, 1, 0, 0, 0, 323, 324, 5, 60, 0, 0, 324, 325, 5, 61, 
	0, 0, 325, 88, 1, 0, 0, 0, 326, 327, 5, 62, 0, 0, 327, 90, 1, 0, 0, 0, 
	328, 329, 5, 62, 0, 0, 329, 330, 5, 61, 0, 0, 330, 92, 1, 0, 0, 0, 331, 
	332, 5, 60, 0, 0, 332, 333, 5, 62, 0, 0, 333, 94, 1, 0, 0, 0, 334, 335, 
	5, 97, 0, 0, 335, 336, 5, 110, 0, 0, 336, 337, 5, 100, 0, 0, 337, 96, 1, 
	0, 0, 0, 338, 339, 5, 111, 0, 0, 339, 340, 5, 114, 0, 0, 340, 98, 1, 0, 
	0, 0, 341, 342, 5, 120, 0, 0, 342, 343, 5, 111, 0, 0, 343, 344, 5, 114, 
	0, 0, 344, 100, 1, 0, 0, 0, 345, 346, 5, 61, 0, 0, 346, 347, 5, 62, 0, 
	0, 347, 102, 1, 0, 0, 0, 348, 349, 5, 45, 0, 0, 349, 350, 5, 62, 0, 0, 
	350, 104, 1, 0, 0, 0, 351, 352, 5, 105, 0, 0, 352, 353, 5, 102, 0, 0, 353, 
	106, 1, 0, 0, 0, 354, 355, 5, 116, 0, 0, 355, 356, 5, 104, 0, 0, 356, 357, 
	5, 101, 0, 0, 357, 358, 5, 110, 0, 0, 358, 108, 1, 0, 0, 0, 359, 360, 5, 
	101, 0, 0, 360, 361, 5, 108, 0, 0, 361, 362, 5, 115, 0, 0, 362, 363, 5, 
	101, 0, 0, 363, 110, 1, 0, 0, 0, 364, 365, 3, 115, 57, 0, 365, 366, 5, 
	46, 0, 0, 366, 367, 3, 115, 57, 0, 367, 112, 1, 0, 0, 0, 368, 369, 5, 116, 
	0, 0, 369, 370, 5, 114, 0, 0, 370, 371, 5, 117, 0, 0, 371, 378, 5, 101, 
	0, 0, 372, 373, 5, 102, 0, 0, 373, 374, 5, 97, 0, 0, 374, 375, 5, 108, 
	0, 0, 375, 376, 5, 115, 0, 0, 376, 378, 5, 101, 0, 0, 377, 368, 1, 0, 0, 
	0, 377, 372, 1, 0, 0, 0, 378, 114, 1, 0, 0, 0, 379, 381, 7, 0, 0, 0, 380, 
	379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 
	1, 0, 0, 0, 383, 116, 1, 0, 0, 0, 384, 388, 7, 1, 0, 0, 385, 387, 7, 2, 
	0, 0, 386, 385, 1, 0, 0, 0, 387, 390, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 
	388, 389, 1, 0, 0, 0, 389, 118, 1, 0, 0, 0, 390, 388, 1, 0, 0, 0, 391, 
	393, 7, 3, 0, 0, 392, 391, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394, 392, 
	1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 6, 59, 
	0, 0, 397, 120, 1, 0, 0, 0, 398, 399, 5, 45, 0, 0, 399, 400, 5, 45, 0, 
	0, 400, 409, 1, 0, 0, 0, 401, 405, 8, 4, 0, 0, 402, 404, 8, 5, 0, 0, 403, 
	402, 1, 0, 0, 0, 404, 407, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 406, 
	1, 0, 0, 0, 406, 410, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 408, 410, 1, 0, 
	0, 0, 409, 401, 1, 0, 0, 0, 409, 408, 1, 0, 0, 0, 410, 415, 1, 0, 0, 0, 
	411, 413, 5, 13, 0, 0, 412, 411, 1, 0, 0, 0, 412, 413, 1, 0, 0, 0, 413, 
	414, 1, 0, 0, 0, 414, 416, 5, 10, 0, 0, 415, 412, 1, 0, 0, 0, 415, 416, 
	1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 418, 6, 60, 0, 0, 418, 122, 1, 0, 
	0, 0, 419, 420, 5, 40, 0, 0, 420, 421, 5, 42, 0, 0, 421, 425, 1, 0, 0, 
	0, 422, 424, 9, 0, 0, 0, 423, 422, 1, 0, 0, 0, 424, 427, 1, 0, 0, 0, 425, 
	426, 1, 0, 0, 0, 425, 423, 1, 0, 0, 0, 426, 428, 1, 0, 0, 0, 427, 425, 
	1, 0, 0, 0, 428, 429, 5, 42, 0, 0, 429, 430, 5, 41, 0, 0, 430, 431, 1, 
	0, 0, 0, 431, 432, 6, 61, 0, 0, 432, 124, 1, 0, 0, 0, 433, 434, 9, 0, 0, 
	0, 434, 126, 1, 0, 0, 0, 10, 0, 377, 382, 388, 394, 405, 409, 412, 415, 
	425, 1, 6, 0, 0,
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

// LustreLexerInit initializes any static state used to implement LustreLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLustreLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LustreLexerInit() {
  staticData := &LustreLexerLexerStaticData
  staticData.once.Do(lustrelexerLexerInit)
}

// NewLustreLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLustreLexer(input antlr.CharStream) *LustreLexer {
  LustreLexerInit()
	l := new(LustreLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &LustreLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Lustre.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LustreLexer tokens.
const (
	LustreLexerT__0 = 1
	LustreLexerT__1 = 2
	LustreLexerT__2 = 3
	LustreLexerT__3 = 4
	LustreLexerT__4 = 5
	LustreLexerT__5 = 6
	LustreLexerT__6 = 7
	LustreLexerT__7 = 8
	LustreLexerT__8 = 9
	LustreLexerT__9 = 10
	LustreLexerT__10 = 11
	LustreLexerT__11 = 12
	LustreLexerT__12 = 13
	LustreLexerT__13 = 14
	LustreLexerT__14 = 15
	LustreLexerT__15 = 16
	LustreLexerT__16 = 17
	LustreLexerT__17 = 18
	LustreLexerT__18 = 19
	LustreLexerT__19 = 20
	LustreLexerT__20 = 21
	LustreLexerT__21 = 22
	LustreLexerT__22 = 23
	LustreLexerT__23 = 24
	LustreLexerT__24 = 25
	LustreLexerT__25 = 26
	LustreLexerT__26 = 27
	LustreLexerT__27 = 28
	LustreLexerT__28 = 29
	LustreLexerT__29 = 30
	LustreLexerT__30 = 31
	LustreLexerT__31 = 32
	LustreLexerT__32 = 33
	LustreLexerT__33 = 34
	LustreLexerT__34 = 35
	LustreLexerT__35 = 36
	LustreLexerT__36 = 37
	LustreLexerT__37 = 38
	LustreLexerT__38 = 39
	LustreLexerT__39 = 40
	LustreLexerT__40 = 41
	LustreLexerT__41 = 42
	LustreLexerT__42 = 43
	LustreLexerT__43 = 44
	LustreLexerT__44 = 45
	LustreLexerT__45 = 46
	LustreLexerT__46 = 47
	LustreLexerT__47 = 48
	LustreLexerT__48 = 49
	LustreLexerT__49 = 50
	LustreLexerT__50 = 51
	LustreLexerT__51 = 52
	LustreLexerT__52 = 53
	LustreLexerT__53 = 54
	LustreLexerT__54 = 55
	LustreLexerREAL = 56
	LustreLexerBOOL = 57
	LustreLexerINT = 58
	LustreLexerID = 59
	LustreLexerWS = 60
	LustreLexerSL_COMMENT = 61
	LustreLexerML_COMMENT = 62
	LustreLexerERROR = 63
)

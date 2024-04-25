// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lustre
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type LustreParser struct {
	*antlr.BaseParser
}

var LustreParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func lustreParserInit() {
  staticData := &LustreParserStaticData
  staticData.LiteralNames = []string{
    "", "'type'", "'='", "';'", "'const'", "':'", "'node'", "'('", "')'", 
    "'returns'", "'var'", "'let'", "'tel'", "'function'", "','", "'struct'", 
    "'{'", "'}'", "'enum'", "'int'", "'subrange'", "'['", "']'", "'of'", 
    "'bool'", "'real'", "'-'", "'automaton'", "'initial'", "'final'", "'state'", 
    "'unless'", "'until'", "'synchro'", "'fork'", "'if'", "'elsif'", "'else'", 
    "'restart'", "'resume'", "'do'", "'emit'", "'clock'", "'default'", "'last'", 
    "'--%PROPERTY'", "'--%REALIZABLE'", "'--%IVC'", "'--%MAIN'", "'assert'", 
    "'_'", "'..'", "'activate'", "'then'", "'when'", "'match'", "'|'", "'floor'", 
    "'condact'", "'.'", "':='", "'pre'", "'not'", "'*'", "'/'", "'div'", 
    "'mod'", "'+'", "'<'", "'<='", "'>'", "'>='", "'<>'", "'and'", "'or'", 
    "'xor'", "'=>'", "'->'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "REAL", "BOOL", "INT", "ID", 
    "WS", "SL_COMMENT", "ML_COMMENT", "ERROR",
  }
  staticData.RuleNames = []string{
    "program", "typedef", "constant", "node", "function", "varDeclList", 
    "varDeclGroup", "topLevelType", "type", "bound", "stateMachine", "stateDecl", 
    "transition", "arrow", "fork", "elsifFork", "elseFork", "target", "actions", 
    "dataDef", "scope", "localBlock", "eqs", "varDecl", "varID", "defaultDecl", 
    "lastDecl", "property", "realizabilityInputs", "ivc", "main", "assertion", 
    "equation", "simpleEquation", "lhs", "lhsID", "controlBlock", "emission", 
    "emissionBody", "return", "returnVar", "clockedBlock", "ifBlock", "matchBlock", 
    "pattern", "expr", "eID",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 85, 707, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7, 
	31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36, 
	2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2, 
	42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 1, 0, 
	1, 0, 1, 0, 1, 0, 5, 0, 99, 8, 0, 10, 0, 12, 0, 102, 9, 0, 1, 0, 1, 0, 
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 116, 
	8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 126, 8, 3, 
	1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 132, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 
	3, 3, 139, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 
	149, 8, 3, 10, 3, 12, 3, 152, 9, 3, 1, 3, 1, 3, 3, 3, 156, 8, 3, 1, 4, 
	1, 4, 1, 4, 1, 4, 3, 4, 162, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 168, 8, 
	4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 176, 8, 5, 10, 5, 12, 5, 179, 
	9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 184, 8, 6, 10, 6, 12, 6, 187, 9, 6, 1, 6, 
	1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 5, 7, 203, 8, 7, 10, 7, 12, 7, 206, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 1, 7, 5, 7, 215, 8, 7, 10, 7, 12, 7, 218, 9, 7, 1, 7, 3, 7, 
	221, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 237, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 
	243, 8, 8, 10, 8, 12, 8, 246, 9, 8, 1, 9, 3, 9, 249, 8, 9, 1, 9, 1, 9, 
	1, 10, 1, 10, 1, 10, 5, 10, 256, 8, 10, 10, 10, 12, 10, 259, 9, 10, 1, 
	11, 3, 11, 262, 8, 11, 1, 11, 3, 11, 265, 8, 11, 1, 11, 1, 11, 1, 11, 1, 
	11, 1, 11, 1, 11, 4, 11, 273, 8, 11, 11, 11, 12, 11, 274, 3, 11, 277, 8, 
	11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 284, 8, 11, 10, 11, 12, 11, 
	287, 9, 11, 1, 11, 1, 11, 3, 11, 291, 8, 11, 1, 11, 1, 11, 3, 11, 295, 
	8, 11, 3, 11, 297, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 3, 13, 304, 
	8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 313, 8, 
	14, 10, 14, 12, 14, 316, 9, 14, 1, 14, 3, 14, 319, 8, 14, 3, 14, 321, 8, 
	14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 
	1, 17, 3, 17, 334, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 341, 
	8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 347, 8, 18, 10, 18, 12, 18, 350, 
	9, 18, 1, 18, 1, 18, 3, 18, 354, 8, 18, 1, 19, 1, 19, 3, 19, 358, 8, 19, 
	1, 20, 3, 20, 361, 8, 20, 1, 20, 3, 20, 364, 8, 20, 1, 21, 1, 21, 1, 21, 
	1, 21, 5, 21, 370, 8, 21, 10, 21, 12, 21, 373, 9, 21, 1, 22, 1, 22, 1, 
	22, 1, 22, 5, 22, 379, 8, 22, 10, 22, 12, 22, 382, 9, 22, 1, 22, 1, 22, 
	1, 23, 1, 23, 1, 23, 5, 23, 389, 8, 23, 10, 23, 12, 23, 392, 9, 23, 1, 
	23, 1, 23, 1, 23, 3, 23, 397, 8, 23, 1, 23, 3, 23, 400, 8, 23, 1, 24, 3, 
	24, 403, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 
	1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 5, 
	28, 423, 8, 28, 10, 28, 12, 28, 426, 9, 28, 3, 28, 428, 8, 28, 1, 28, 1, 
	28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 436, 8, 29, 10, 29, 12, 29, 439, 
	9, 29, 3, 29, 441, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 3, 30, 447, 8, 30, 
	1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 458, 
	8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 
	34, 469, 8, 34, 10, 34, 12, 34, 472, 9, 34, 3, 34, 474, 8, 34, 1, 35, 1, 
	35, 3, 35, 478, 8, 35, 1, 36, 1, 36, 3, 36, 482, 8, 36, 1, 37, 1, 37, 1, 
	37, 1, 38, 1, 38, 1, 38, 3, 38, 490, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 
	5, 38, 496, 8, 38, 10, 38, 12, 38, 499, 9, 38, 1, 38, 1, 38, 1, 38, 3, 
	38, 504, 8, 38, 3, 38, 506, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 
	40, 1, 40, 5, 40, 515, 8, 40, 10, 40, 12, 40, 518, 9, 40, 1, 40, 1, 40, 
	3, 40, 522, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 528, 8, 41, 1, 42, 
	1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 535, 8, 42, 1, 42, 1, 42, 1, 42, 3, 
	42, 540, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 
	4, 43, 550, 8, 43, 11, 43, 12, 43, 551, 1, 44, 1, 44, 1, 45, 1, 45, 1, 
	45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 
	1, 45, 1, 45, 5, 45, 571, 8, 45, 10, 45, 12, 45, 574, 9, 45, 3, 45, 576, 
	8, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 4, 45, 585, 8, 
	45, 11, 45, 12, 45, 586, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 
	1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 
	45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 613, 8, 45, 10, 45, 
	12, 45, 616, 9, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 5, 45, 624, 
	8, 45, 10, 45, 12, 45, 627, 9, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 
	45, 5, 45, 635, 8, 45, 10, 45, 12, 45, 638, 9, 45, 1, 45, 1, 45, 3, 45, 
	642, 8, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 
	45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 
	1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 
	45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 
	1, 45, 1, 45, 1, 45, 5, 45, 687, 8, 45, 10, 45, 12, 45, 690, 9, 45, 1, 
	46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 
	702, 8, 46, 10, 46, 12, 46, 705, 9, 46, 1, 46, 0, 3, 16, 90, 92, 47, 0, 
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 
	76, 78, 80, 82, 84, 86, 88, 90, 92, 0, 6, 1, 0, 78, 81, 2, 0, 25, 25, 57, 
	57, 1, 0, 63, 66, 2, 0, 26, 26, 67, 67, 2, 0, 2, 2, 68, 72, 1, 0, 74, 75, 
	768, 0, 100, 1, 0, 0, 0, 2, 105, 1, 0, 0, 0, 4, 111, 1, 0, 0, 0, 6, 121, 
	1, 0, 0, 0, 8, 157, 1, 0, 0, 0, 10, 172, 1, 0, 0, 0, 12, 180, 1, 0, 0, 
	0, 14, 220, 1, 0, 0, 0, 16, 236, 1, 0, 0, 0, 18, 248, 1, 0, 0, 0, 20, 252, 
	1, 0, 0, 0, 22, 261, 1, 0, 0, 0, 24, 298, 1, 0, 0, 0, 26, 303, 1, 0, 0, 
	0, 28, 320, 1, 0, 0, 0, 30, 322, 1, 0, 0, 0, 32, 326, 1, 0, 0, 0, 34, 333, 
	1, 0, 0, 0, 36, 353, 1, 0, 0, 0, 38, 357, 1, 0, 0, 0, 40, 360, 1, 0, 0, 
	0, 42, 365, 1, 0, 0, 0, 44, 374, 1, 0, 0, 0, 46, 385, 1, 0, 0, 0, 48, 402, 
	1, 0, 0, 0, 50, 406, 1, 0, 0, 0, 52, 410, 1, 0, 0, 0, 54, 414, 1, 0, 0, 
	0, 56, 418, 1, 0, 0, 0, 58, 431, 1, 0, 0, 0, 60, 444, 1, 0, 0, 0, 62, 448, 
	1, 0, 0, 0, 64, 457, 1, 0, 0, 0, 66, 459, 1, 0, 0, 0, 68, 473, 1, 0, 0, 
	0, 70, 477, 1, 0, 0, 0, 72, 481, 1, 0, 0, 0, 74, 483, 1, 0, 0, 0, 76, 505, 
	1, 0, 0, 0, 78, 507, 1, 0, 0, 0, 80, 516, 1, 0, 0, 0, 82, 523, 1, 0, 0, 
	0, 84, 529, 1, 0, 0, 0, 86, 541, 1, 0, 0, 0, 88, 553, 1, 0, 0, 0, 90, 641, 
	1, 0, 0, 0, 92, 691, 1, 0, 0, 0, 94, 99, 3, 2, 1, 0, 95, 99, 3, 4, 2, 0, 
	96, 99, 3, 6, 3, 0, 97, 99, 3, 8, 4, 0, 98, 94, 1, 0, 0, 0, 98, 95, 1, 
	0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 97, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 
	98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 100, 1, 
	0, 0, 0, 103, 104, 5, 0, 0, 1, 104, 1, 1, 0, 0, 0, 105, 106, 5, 1, 0, 0, 
	106, 107, 5, 81, 0, 0, 107, 108, 5, 2, 0, 0, 108, 109, 3, 14, 7, 0, 109, 
	110, 5, 3, 0, 0, 110, 3, 1, 0, 0, 0, 111, 112, 5, 4, 0, 0, 112, 115, 5, 
	81, 0, 0, 113, 114, 5, 5, 0, 0, 114, 116, 3, 16, 8, 0, 115, 113, 1, 0, 
	0, 0, 115, 116, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 5, 2, 0, 0, 
	118, 119, 3, 90, 45, 0, 119, 120, 5, 3, 0, 0, 120, 5, 1, 0, 0, 0, 121, 
	122, 5, 6, 0, 0, 122, 123, 5, 81, 0, 0, 123, 125, 5, 7, 0, 0, 124, 126, 
	3, 10, 5, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 
	0, 0, 127, 128, 5, 8, 0, 0, 128, 129, 5, 9, 0, 0, 129, 131, 5, 7, 0, 0, 
	130, 132, 3, 10, 5, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 
	133, 1, 0, 0, 0, 133, 138, 5, 8, 0, 0, 134, 135, 5, 10, 0, 0, 135, 136, 
	3, 10, 5, 0, 136, 137, 5, 3, 0, 0, 137, 139, 1, 0, 0, 0, 138, 134, 1, 0, 
	0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 150, 5, 11, 0, 0, 
	141, 149, 3, 64, 32, 0, 142, 149, 3, 54, 27, 0, 143, 149, 3, 62, 31, 0, 
	144, 149, 3, 60, 30, 0, 145, 149, 3, 56, 28, 0, 146, 149, 3, 58, 29, 0, 
	147, 149, 3, 20, 10, 0, 148, 141, 1, 0, 0, 0, 148, 142, 1, 0, 0, 0, 148, 
	143, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148, 146, 
	1, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 
	0, 0, 150, 151, 1, 0, 0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 
	153, 155, 5, 12, 0, 0, 154, 156, 5, 3, 0, 0, 155, 154, 1, 0, 0, 0, 155, 
	156, 1, 0, 0, 0, 156, 7, 1, 0, 0, 0, 157, 158, 5, 13, 0, 0, 158, 159, 3, 
	92, 46, 0, 159, 161, 5, 7, 0, 0, 160, 162, 3, 10, 5, 0, 161, 160, 1, 0, 
	0, 0, 161, 162, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 8, 0, 0, 
	164, 165, 5, 9, 0, 0, 165, 167, 5, 7, 0, 0, 166, 168, 3, 10, 5, 0, 167, 
	166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 
	5, 8, 0, 0, 170, 171, 5, 3, 0, 0, 171, 9, 1, 0, 0, 0, 172, 177, 3, 12, 
	6, 0, 173, 174, 5, 3, 0, 0, 174, 176, 3, 12, 6, 0, 175, 173, 1, 0, 0, 0, 
	176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 
	11, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 185, 3, 92, 46, 0, 181, 182, 
	5, 14, 0, 0, 182, 184, 3, 92, 46, 0, 183, 181, 1, 0, 0, 0, 184, 187, 1, 
	0, 0, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 188, 1, 0, 0, 
	0, 187, 185, 1, 0, 0, 0, 188, 189, 5, 5, 0, 0, 189, 190, 3, 16, 8, 0, 190, 
	13, 1, 0, 0, 0, 191, 221, 3, 16, 8, 0, 192, 193, 5, 15, 0, 0, 193, 194, 
	5, 16, 0, 0, 194, 195, 5, 81, 0, 0, 195, 196, 5, 5, 0, 0, 196, 197, 3, 
	16, 8, 0, 197, 204, 1, 0, 0, 0, 198, 199, 5, 3, 0, 0, 199, 200, 5, 81, 
	0, 0, 200, 201, 5, 5, 0, 0, 201, 203, 3, 16, 8, 0, 202, 198, 1, 0, 0, 0, 
	203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 
	207, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 5, 17, 0, 0, 208, 221, 
	1, 0, 0, 0, 209, 210, 5, 18, 0, 0, 210, 211, 5, 16, 0, 0, 211, 216, 5, 
	81, 0, 0, 212, 213, 5, 14, 0, 0, 213, 215, 5, 81, 0, 0, 214, 212, 1, 0, 
	0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 
	217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 221, 5, 17, 0, 0, 220, 
	191, 1, 0, 0, 0, 220, 192, 1, 0, 0, 0, 220, 209, 1, 0, 0, 0, 221, 15, 1, 
	0, 0, 0, 222, 223, 6, 8, -1, 0, 223, 237, 5, 19, 0, 0, 224, 225, 5, 20, 
	0, 0, 225, 226, 5, 21, 0, 0, 226, 227, 3, 18, 9, 0, 227, 228, 5, 14, 0, 
	0, 228, 229, 3, 18, 9, 0, 229, 230, 5, 22, 0, 0, 230, 231, 5, 23, 0, 0, 
	231, 232, 5, 19, 0, 0, 232, 237, 1, 0, 0, 0, 233, 237, 5, 24, 0, 0, 234, 
	237, 5, 25, 0, 0, 235, 237, 5, 81, 0, 0, 236, 222, 1, 0, 0, 0, 236, 224, 
	1, 0, 0, 0, 236, 233, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 235, 1, 0, 
	0, 0, 237, 244, 1, 0, 0, 0, 238, 239, 10, 2, 0, 0, 239, 240, 5, 21, 0, 
	0, 240, 241, 5, 80, 0, 0, 241, 243, 5, 22, 0, 0, 242, 238, 1, 0, 0, 0, 
	243, 246, 1, 0, 0, 0, 244, 242, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 
	17, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 247, 249, 5, 26, 0, 0, 248, 247, 
	1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 80, 
	0, 0, 251, 19, 1, 0, 0, 0, 252, 253, 5, 27, 0, 0, 253, 257, 5, 81, 0, 0, 
	254, 256, 3, 22, 11, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 
	255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 21, 1, 0, 0, 0, 259, 257, 1, 
	0, 0, 0, 260, 262, 5, 28, 0, 0, 261, 260, 1, 0, 0, 0, 261, 262, 1, 0, 0, 
	0, 262, 264, 1, 0, 0, 0, 263, 265, 5, 29, 0, 0, 264, 263, 1, 0, 0, 0, 264, 
	265, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 5, 30, 0, 0, 267, 276, 
	5, 81, 0, 0, 268, 272, 5, 31, 0, 0, 269, 270, 3, 24, 12, 0, 270, 271, 5, 
	3, 0, 0, 271, 273, 1, 0, 0, 0, 272, 269, 1, 0, 0, 0, 273, 274, 1, 0, 0, 
	0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 277, 1, 0, 0, 0, 276, 
	268, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 296, 
	3, 38, 19, 0, 279, 285, 5, 32, 0, 0, 280, 281, 3, 24, 12, 0, 281, 282, 
	5, 3, 0, 0, 282, 284, 1, 0, 0, 0, 283, 280, 1, 0, 0, 0, 284, 287, 1, 0, 
	0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 294, 1, 0, 0, 0, 
	287, 285, 1, 0, 0, 0, 288, 290, 5, 33, 0, 0, 289, 291, 3, 36, 18, 0, 290, 
	289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293, 
	5, 34, 0, 0, 293, 295, 5, 3, 0, 0, 294, 288, 1, 0, 0, 0, 294, 295, 1, 0, 
	0, 0, 295, 297, 1, 0, 0, 0, 296, 279, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 
	297, 23, 1, 0, 0, 0, 298, 299, 5, 35, 0, 0, 299, 300, 3, 90, 45, 0, 300, 
	301, 3, 26, 13, 0, 301, 25, 1, 0, 0, 0, 302, 304, 3, 36, 18, 0, 303, 302, 
	1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 3, 28, 
	14, 0, 306, 27, 1, 0, 0, 0, 307, 321, 3, 34, 17, 0, 308, 309, 5, 35, 0, 
	0, 309, 310, 3, 90, 45, 0, 310, 314, 3, 26, 13, 0, 311, 313, 3, 30, 15, 
	0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 
	315, 1, 0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 319, 
	3, 32, 16, 0, 318, 317, 1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 321, 1, 
	0, 0, 0, 320, 307, 1, 0, 0, 0, 320, 308, 1, 0, 0, 0, 321, 29, 1, 0, 0, 
	0, 322, 323, 5, 36, 0, 0, 323, 324, 3, 90, 45, 0, 324, 325, 3, 26, 13, 
	0, 325, 31, 1, 0, 0, 0, 326, 327, 5, 37, 0, 0, 327, 328, 3, 26, 13, 0, 
	328, 33, 1, 0, 0, 0, 329, 330, 5, 38, 0, 0, 330, 334, 5, 81, 0, 0, 331, 
	332, 5, 39, 0, 0, 332, 334, 5, 81, 0, 0, 333, 329, 1, 0, 0, 0, 333, 331, 
	1, 0, 0, 0, 334, 35, 1, 0, 0, 0, 335, 336, 5, 40, 0, 0, 336, 354, 3, 38, 
	19, 0, 337, 338, 5, 40, 0, 0, 338, 340, 5, 16, 0, 0, 339, 341, 5, 41, 0, 
	0, 340, 339, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 342, 1, 0, 0, 0, 342, 
	348, 3, 76, 38, 0, 343, 344, 5, 3, 0, 0, 344, 345, 5, 41, 0, 0, 345, 347, 
	3, 76, 38, 0, 346, 343, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348, 346, 1, 
	0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 351, 1, 0, 0, 0, 350, 348, 1, 0, 0, 
	0, 351, 352, 5, 17, 0, 0, 352, 354, 1, 0, 0, 0, 353, 335, 1, 0, 0, 0, 353, 
	337, 1, 0, 0, 0, 354, 37, 1, 0, 0, 0, 355, 358, 3, 64, 32, 0, 356, 358, 
	3, 40, 20, 0, 357, 355, 1, 0, 0, 0, 357, 356, 1, 0, 0, 0, 358, 39, 1, 0, 
	0, 0, 359, 361, 3, 42, 21, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1, 0, 0, 
	0, 361, 363, 1, 0, 0, 0, 362, 364, 3, 44, 22, 0, 363, 362, 1, 0, 0, 0, 
	363, 364, 1, 0, 0, 0, 364, 41, 1, 0, 0, 0, 365, 371, 5, 10, 0, 0, 366, 
	367, 3, 46, 23, 0, 367, 368, 5, 3, 0, 0, 368, 370, 1, 0, 0, 0, 369, 366, 
	1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0, 
	0, 0, 372, 43, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 374, 380, 5, 11, 0, 0, 
	375, 376, 3, 64, 32, 0, 376, 377, 5, 3, 0, 0, 377, 379, 1, 0, 0, 0, 378, 
	375, 1, 0, 0, 0, 379, 382, 1, 0, 0, 0, 380, 378, 1, 0, 0, 0, 380, 381, 
	1, 0, 0, 0, 381, 383, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 383, 384, 5, 12, 
	0, 0, 384, 45, 1, 0, 0, 0, 385, 390, 3, 48, 24, 0, 386, 387, 5, 14, 0, 
	0, 387, 389, 3, 48, 24, 0, 388, 386, 1, 0, 0, 0, 389, 392, 1, 0, 0, 0, 
	390, 388, 1, 0, 0, 0, 390, 391, 1, 0, 0, 0, 391, 393, 1, 0, 0, 0, 392, 
	390, 1, 0, 0, 0, 393, 394, 5, 5, 0, 0, 394, 396, 3, 16, 8, 0, 395, 397, 
	3, 50, 25, 0, 396, 395, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 399, 1, 
	0, 0, 0, 398, 400, 3, 52, 26, 0, 399, 398, 1, 0, 0, 0, 399, 400, 1, 0, 
	0, 0, 400, 47, 1, 0, 0, 0, 401, 403, 5, 42, 0, 0, 402, 401, 1, 0, 0, 0, 
	402, 403, 1, 0, 0, 0, 403, 404, 1, 0, 0, 0, 404, 405, 5, 81, 0, 0, 405, 
	49, 1, 0, 0, 0, 406, 407, 5, 43, 0, 0, 407, 408, 5, 2, 0, 0, 408, 409, 
	3, 90, 45, 0, 409, 51, 1, 0, 0, 0, 410, 411, 5, 44, 0, 0, 411, 412, 5, 
	2, 0, 0, 412, 413, 3, 90, 45, 0, 413, 53, 1, 0, 0, 0, 414, 415, 5, 45, 
	0, 0, 415, 416, 3, 92, 46, 0, 416, 417, 5, 3, 0, 0, 417, 55, 1, 0, 0, 0, 
	418, 427, 5, 46, 0, 0, 419, 424, 5, 81, 0, 0, 420, 421, 5, 14, 0, 0, 421, 
	423, 5, 81, 0, 0, 422, 420, 1, 0, 0, 0, 423, 426, 1, 0, 0, 0, 424, 422, 
	1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 428, 1, 0, 0, 0, 426, 424, 1, 0, 
	0, 0, 427, 419, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 
	429, 430, 5, 3, 0, 0, 430, 57, 1, 0, 0, 0, 431, 440, 5, 47, 0, 0, 432, 
	437, 3, 92, 46, 0, 433, 434, 5, 14, 0, 0, 434, 436, 3, 92, 46, 0, 435, 
	433, 1, 0, 0, 0, 436, 439, 1, 0, 0, 0, 437, 435, 1, 0, 0, 0, 437, 438, 
	1, 0, 0, 0, 438, 441, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 440, 432, 1, 0, 
	0, 0, 440, 441, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 443, 5, 3, 0, 0, 
	443, 59, 1, 0, 0, 0, 444, 446, 5, 48, 0, 0, 445, 447, 5, 3, 0, 0, 446, 
	445, 1, 0, 0, 0, 446, 447, 1, 0, 0, 0, 447, 61, 1, 0, 0, 0, 448, 449, 5, 
	49, 0, 0, 449, 450, 3, 90, 45, 0, 450, 451, 5, 3, 0, 0, 451, 63, 1, 0, 
	0, 0, 452, 458, 3, 66, 33, 0, 453, 458, 3, 74, 37, 0, 454, 455, 3, 72, 
	36, 0, 455, 456, 3, 78, 39, 0, 456, 458, 1, 0, 0, 0, 457, 452, 1, 0, 0, 
	0, 457, 453, 1, 0, 0, 0, 457, 454, 1, 0, 0, 0, 458, 65, 1, 0, 0, 0, 459, 
	460, 3, 68, 34, 0, 460, 461, 5, 2, 0, 0, 461, 462, 3, 90, 45, 0, 462, 67, 
	1, 0, 0, 0, 463, 464, 5, 7, 0, 0, 464, 474, 5, 8, 0, 0, 465, 470, 3, 70, 
	35, 0, 466, 467, 5, 14, 0, 0, 467, 469, 3, 70, 35, 0, 468, 466, 1, 0, 0, 
	0, 469, 472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 
	474, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 473, 463, 1, 0, 0, 0, 473, 465, 
	1, 0, 0, 0, 474, 69, 1, 0, 0, 0, 475, 478, 3, 92, 46, 0, 476, 478, 5, 50, 
	0, 0, 477, 475, 1, 0, 0, 0, 477, 476, 1, 0, 0, 0, 478, 71, 1, 0, 0, 0, 
	479, 482, 3, 20, 10, 0, 480, 482, 3, 82, 41, 0, 481, 479, 1, 0, 0, 0, 481, 
	480, 1, 0, 0, 0, 482, 73, 1, 0, 0, 0, 483, 484, 5, 41, 0, 0, 484, 485, 
	3, 76, 38, 0, 485, 75, 1, 0, 0, 0, 486, 489, 3, 92, 46, 0, 487, 488, 5, 
	35, 0, 0, 488, 490, 3, 90, 45, 0, 489, 487, 1, 0, 0, 0, 489, 490, 1, 0, 
	0, 0, 490, 506, 1, 0, 0, 0, 491, 492, 5, 7, 0, 0, 492, 497, 3, 92, 46, 
	0, 493, 494, 5, 14, 0, 0, 494, 496, 3, 92, 46, 0, 495, 493, 1, 0, 0, 0, 
	496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0, 0, 0, 498, 
	500, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 500, 503, 5, 8, 0, 0, 501, 502, 
	5, 35, 0, 0, 502, 504, 3, 90, 45, 0, 503, 501, 1, 0, 0, 0, 503, 504, 1, 
	0, 0, 0, 504, 506, 1, 0, 0, 0, 505, 486, 1, 0, 0, 0, 505, 491, 1, 0, 0, 
	0, 506, 77, 1, 0, 0, 0, 507, 508, 5, 9, 0, 0, 508, 509, 3, 80, 40, 0, 509, 
	510, 5, 3, 0, 0, 510, 79, 1, 0, 0, 0, 511, 512, 3, 92, 46, 0, 512, 513, 
	5, 14, 0, 0, 513, 515, 1, 0, 0, 0, 514, 511, 1, 0, 0, 0, 515, 518, 1, 0, 
	0, 0, 516, 514, 1, 0, 0, 0, 516, 517, 1, 0, 0, 0, 517, 521, 1, 0, 0, 0, 
	518, 516, 1, 0, 0, 0, 519, 522, 3, 92, 46, 0, 520, 522, 5, 51, 0, 0, 521, 
	519, 1, 0, 0, 0, 521, 520, 1, 0, 0, 0, 522, 81, 1, 0, 0, 0, 523, 524, 5, 
	52, 0, 0, 524, 527, 3, 92, 46, 0, 525, 528, 3, 84, 42, 0, 526, 528, 3, 
	86, 43, 0, 527, 525, 1, 0, 0, 0, 527, 526, 1, 0, 0, 0, 528, 83, 1, 0, 0, 
	0, 529, 530, 5, 35, 0, 0, 530, 531, 3, 90, 45, 0, 531, 534, 5, 53, 0, 0, 
	532, 535, 3, 38, 19, 0, 533, 535, 3, 84, 42, 0, 534, 532, 1, 0, 0, 0, 534, 
	533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 539, 5, 37, 0, 0, 537, 540, 
	3, 38, 19, 0, 538, 540, 3, 84, 42, 0, 539, 537, 1, 0, 0, 0, 539, 538, 1, 
	0, 0, 0, 540, 85, 1, 0, 0, 0, 541, 542, 5, 54, 0, 0, 542, 543, 3, 90, 45, 
	0, 543, 549, 5, 55, 0, 0, 544, 545, 5, 56, 0, 0, 545, 546, 3, 88, 44, 0, 
	546, 547, 5, 5, 0, 0, 547, 548, 3, 38, 19, 0, 548, 550, 1, 0, 0, 0, 549, 
	544, 1, 0, 0, 0, 550, 551, 1, 0, 0, 0, 551, 549, 1, 0, 0, 0, 551, 552, 
	1, 0, 0, 0, 552, 87, 1, 0, 0, 0, 553, 554, 7, 0, 0, 0, 554, 89, 1, 0, 0, 
	0, 555, 556, 6, 45, -1, 0, 556, 642, 5, 81, 0, 0, 557, 642, 5, 80, 0, 0, 
	558, 642, 5, 78, 0, 0, 559, 642, 5, 79, 0, 0, 560, 561, 7, 1, 0, 0, 561, 
	562, 5, 7, 0, 0, 562, 563, 3, 90, 45, 0, 563, 564, 5, 8, 0, 0, 564, 642, 
	1, 0, 0, 0, 565, 566, 3, 92, 46, 0, 566, 575, 5, 7, 0, 0, 567, 572, 3, 
	90, 45, 0, 568, 569, 5, 14, 0, 0, 569, 571, 3, 90, 45, 0, 570, 568, 1, 
	0, 0, 0, 571, 574, 1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 572, 573, 1, 0, 0, 
	0, 573, 576, 1, 0, 0, 0, 574, 572, 1, 0, 0, 0, 575, 567, 1, 0, 0, 0, 575, 
	576, 1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 577, 578, 5, 8, 0, 0, 578, 642, 
	1, 0, 0, 0, 579, 580, 5, 58, 0, 0, 580, 581, 5, 7, 0, 0, 581, 584, 3, 90, 
	45, 0, 582, 583, 5, 14, 0, 0, 583, 585, 3, 90, 45, 0, 584, 582, 1, 0, 0, 
	0, 585, 586, 1, 0, 0, 0, 586, 584, 1, 0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 
	588, 1, 0, 0, 0, 588, 589, 5, 8, 0, 0, 589, 642, 1, 0, 0, 0, 590, 591, 
	5, 61, 0, 0, 591, 642, 3, 90, 45, 14, 592, 593, 5, 62, 0, 0, 593, 642, 
	3, 90, 45, 13, 594, 595, 5, 26, 0, 0, 595, 642, 3, 90, 45, 12, 596, 597, 
	5, 35, 0, 0, 597, 598, 3, 90, 45, 0, 598, 599, 5, 53, 0, 0, 599, 600, 3, 
	90, 45, 0, 600, 601, 5, 37, 0, 0, 601, 602, 3, 90, 45, 4, 602, 642, 1, 
	0, 0, 0, 603, 604, 5, 81, 0, 0, 604, 605, 5, 16, 0, 0, 605, 606, 5, 81, 
	0, 0, 606, 607, 5, 2, 0, 0, 607, 614, 3, 90, 45, 0, 608, 609, 5, 3, 0, 
	0, 609, 610, 5, 81, 0, 0, 610, 611, 5, 2, 0, 0, 611, 613, 3, 90, 45, 0, 
	612, 608, 1, 0, 0, 0, 613, 616, 1, 0, 0, 0, 614, 612, 1, 0, 0, 0, 614, 
	615, 1, 0, 0, 0, 615, 617, 1, 0, 0, 0, 616, 614, 1, 0, 0, 0, 617, 618, 
	5, 17, 0, 0, 618, 642, 1, 0, 0, 0, 619, 620, 5, 21, 0, 0, 620, 625, 3, 
	90, 45, 0, 621, 622, 5, 14, 0, 0, 622, 624, 3, 90, 45, 0, 623, 621, 1, 
	0, 0, 0, 624, 627, 1, 0, 0, 0, 625, 623, 1, 0, 0, 0, 625, 626, 1, 0, 0, 
	0, 626, 628, 1, 0, 0, 0, 627, 625, 1, 0, 0, 0, 628, 629, 5, 22, 0, 0, 629, 
	642, 1, 0, 0, 0, 630, 631, 5, 7, 0, 0, 631, 636, 3, 90, 45, 0, 632, 633, 
	5, 14, 0, 0, 633, 635, 3, 90, 45, 0, 634, 632, 1, 0, 0, 0, 635, 638, 1, 
	0, 0, 0, 636, 634, 1, 0, 0, 0, 636, 637, 1, 0, 0, 0, 637, 639, 1, 0, 0, 
	0, 638, 636, 1, 0, 0, 0, 639, 640, 5, 8, 0, 0, 640, 642, 1, 0, 0, 0, 641, 
	555, 1, 0, 0, 0, 641, 557, 1, 0, 0, 0, 641, 558, 1, 0, 0, 0, 641, 559, 
	1, 0, 0, 0, 641, 560, 1, 0, 0, 0, 641, 565, 1, 0, 0, 0, 641, 579, 1, 0, 
	0, 0, 641, 590, 1, 0, 0, 0, 641, 592, 1, 0, 0, 0, 641, 594, 1, 0, 0, 0, 
	641, 596, 1, 0, 0, 0, 641, 603, 1, 0, 0, 0, 641, 619, 1, 0, 0, 0, 641, 
	630, 1, 0, 0, 0, 642, 688, 1, 0, 0, 0, 643, 644, 10, 11, 0, 0, 644, 645, 
	7, 2, 0, 0, 645, 687, 3, 90, 45, 12, 646, 647, 10, 10, 0, 0, 647, 648, 
	7, 3, 0, 0, 648, 687, 3, 90, 45, 11, 649, 650, 10, 9, 0, 0, 650, 651, 7, 
	4, 0, 0, 651, 687, 3, 90, 45, 10, 652, 653, 10, 8, 0, 0, 653, 654, 5, 73, 
	0, 0, 654, 687, 3, 90, 45, 9, 655, 656, 10, 7, 0, 0, 656, 657, 7, 5, 0, 
	0, 657, 687, 3, 90, 45, 8, 658, 659, 10, 6, 0, 0, 659, 660, 5, 76, 0, 0, 
	660, 687, 3, 90, 45, 6, 661, 662, 10, 5, 0, 0, 662, 663, 5, 77, 0, 0, 663, 
	687, 3, 90, 45, 5, 664, 665, 10, 18, 0, 0, 665, 666, 5, 59, 0, 0, 666, 
	687, 5, 81, 0, 0, 667, 668, 10, 17, 0, 0, 668, 669, 5, 16, 0, 0, 669, 670, 
	5, 81, 0, 0, 670, 671, 5, 60, 0, 0, 671, 672, 3, 90, 45, 0, 672, 673, 5, 
	17, 0, 0, 673, 687, 1, 0, 0, 0, 674, 675, 10, 16, 0, 0, 675, 676, 5, 21, 
	0, 0, 676, 677, 3, 90, 45, 0, 677, 678, 5, 22, 0, 0, 678, 687, 1, 0, 0, 
	0, 679, 680, 10, 15, 0, 0, 680, 681, 5, 21, 0, 0, 681, 682, 3, 90, 45, 
	0, 682, 683, 5, 60, 0, 0, 683, 684, 3, 90, 45, 0, 684, 685, 5, 22, 0, 0, 
	685, 687, 1, 0, 0, 0, 686, 643, 1, 0, 0, 0, 686, 646, 1, 0, 0, 0, 686, 
	649, 1, 0, 0, 0, 686, 652, 1, 0, 0, 0, 686, 655, 1, 0, 0, 0, 686, 658, 
	1, 0, 0, 0, 686, 661, 1, 0, 0, 0, 686, 664, 1, 0, 0, 0, 686, 667, 1, 0, 
	0, 0, 686, 674, 1, 0, 0, 0, 686, 679, 1, 0, 0, 0, 687, 690, 1, 0, 0, 0, 
	688, 686, 1, 0, 0, 0, 688, 689, 1, 0, 0, 0, 689, 91, 1, 0, 0, 0, 690, 688, 
	1, 0, 0, 0, 691, 692, 6, 46, -1, 0, 692, 693, 5, 81, 0, 0, 693, 703, 1, 
	0, 0, 0, 694, 695, 10, 2, 0, 0, 695, 696, 5, 21, 0, 0, 696, 697, 5, 80, 
	0, 0, 697, 702, 5, 22, 0, 0, 698, 699, 10, 1, 0, 0, 699, 700, 5, 59, 0, 
	0, 700, 702, 5, 81, 0, 0, 701, 694, 1, 0, 0, 0, 701, 698, 1, 0, 0, 0, 702, 
	705, 1, 0, 0, 0, 703, 701, 1, 0, 0, 0, 703, 704, 1, 0, 0, 0, 704, 93, 1, 
	0, 0, 0, 705, 703, 1, 0, 0, 0, 76, 98, 100, 115, 125, 131, 138, 148, 150, 
	155, 161, 167, 177, 185, 204, 216, 220, 236, 244, 248, 257, 261, 264, 274, 
	276, 285, 290, 294, 296, 303, 314, 318, 320, 333, 340, 348, 353, 357, 360, 
	363, 371, 380, 390, 396, 399, 402, 424, 427, 437, 440, 446, 457, 470, 473, 
	477, 481, 489, 497, 503, 505, 516, 521, 527, 534, 539, 551, 572, 575, 586, 
	614, 625, 636, 641, 686, 688, 701, 703,
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

// LustreParserInit initializes any static state used to implement LustreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLustreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LustreParserInit() {
  staticData := &LustreParserStaticData
  staticData.once.Do(lustreParserInit)
}

// NewLustreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLustreParser(input antlr.TokenStream) *LustreParser {
	LustreParserInit()
	this := new(LustreParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &LustreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lustre.g4"

	return this
}


// LustreParser tokens.
const (
	LustreParserEOF = antlr.TokenEOF
	LustreParserT__0 = 1
	LustreParserT__1 = 2
	LustreParserT__2 = 3
	LustreParserT__3 = 4
	LustreParserT__4 = 5
	LustreParserT__5 = 6
	LustreParserT__6 = 7
	LustreParserT__7 = 8
	LustreParserT__8 = 9
	LustreParserT__9 = 10
	LustreParserT__10 = 11
	LustreParserT__11 = 12
	LustreParserT__12 = 13
	LustreParserT__13 = 14
	LustreParserT__14 = 15
	LustreParserT__15 = 16
	LustreParserT__16 = 17
	LustreParserT__17 = 18
	LustreParserT__18 = 19
	LustreParserT__19 = 20
	LustreParserT__20 = 21
	LustreParserT__21 = 22
	LustreParserT__22 = 23
	LustreParserT__23 = 24
	LustreParserT__24 = 25
	LustreParserT__25 = 26
	LustreParserT__26 = 27
	LustreParserT__27 = 28
	LustreParserT__28 = 29
	LustreParserT__29 = 30
	LustreParserT__30 = 31
	LustreParserT__31 = 32
	LustreParserT__32 = 33
	LustreParserT__33 = 34
	LustreParserT__34 = 35
	LustreParserT__35 = 36
	LustreParserT__36 = 37
	LustreParserT__37 = 38
	LustreParserT__38 = 39
	LustreParserT__39 = 40
	LustreParserT__40 = 41
	LustreParserT__41 = 42
	LustreParserT__42 = 43
	LustreParserT__43 = 44
	LustreParserT__44 = 45
	LustreParserT__45 = 46
	LustreParserT__46 = 47
	LustreParserT__47 = 48
	LustreParserT__48 = 49
	LustreParserT__49 = 50
	LustreParserT__50 = 51
	LustreParserT__51 = 52
	LustreParserT__52 = 53
	LustreParserT__53 = 54
	LustreParserT__54 = 55
	LustreParserT__55 = 56
	LustreParserT__56 = 57
	LustreParserT__57 = 58
	LustreParserT__58 = 59
	LustreParserT__59 = 60
	LustreParserT__60 = 61
	LustreParserT__61 = 62
	LustreParserT__62 = 63
	LustreParserT__63 = 64
	LustreParserT__64 = 65
	LustreParserT__65 = 66
	LustreParserT__66 = 67
	LustreParserT__67 = 68
	LustreParserT__68 = 69
	LustreParserT__69 = 70
	LustreParserT__70 = 71
	LustreParserT__71 = 72
	LustreParserT__72 = 73
	LustreParserT__73 = 74
	LustreParserT__74 = 75
	LustreParserT__75 = 76
	LustreParserT__76 = 77
	LustreParserREAL = 78
	LustreParserBOOL = 79
	LustreParserINT = 80
	LustreParserID = 81
	LustreParserWS = 82
	LustreParserSL_COMMENT = 83
	LustreParserML_COMMENT = 84
	LustreParserERROR = 85
)

// LustreParser rules.
const (
	LustreParserRULE_program = 0
	LustreParserRULE_typedef = 1
	LustreParserRULE_constant = 2
	LustreParserRULE_node = 3
	LustreParserRULE_function = 4
	LustreParserRULE_varDeclList = 5
	LustreParserRULE_varDeclGroup = 6
	LustreParserRULE_topLevelType = 7
	LustreParserRULE_type = 8
	LustreParserRULE_bound = 9
	LustreParserRULE_stateMachine = 10
	LustreParserRULE_stateDecl = 11
	LustreParserRULE_transition = 12
	LustreParserRULE_arrow = 13
	LustreParserRULE_fork = 14
	LustreParserRULE_elsifFork = 15
	LustreParserRULE_elseFork = 16
	LustreParserRULE_target = 17
	LustreParserRULE_actions = 18
	LustreParserRULE_dataDef = 19
	LustreParserRULE_scope = 20
	LustreParserRULE_localBlock = 21
	LustreParserRULE_eqs = 22
	LustreParserRULE_varDecl = 23
	LustreParserRULE_varID = 24
	LustreParserRULE_defaultDecl = 25
	LustreParserRULE_lastDecl = 26
	LustreParserRULE_property = 27
	LustreParserRULE_realizabilityInputs = 28
	LustreParserRULE_ivc = 29
	LustreParserRULE_main = 30
	LustreParserRULE_assertion = 31
	LustreParserRULE_equation = 32
	LustreParserRULE_simpleEquation = 33
	LustreParserRULE_lhs = 34
	LustreParserRULE_lhsID = 35
	LustreParserRULE_controlBlock = 36
	LustreParserRULE_emission = 37
	LustreParserRULE_emissionBody = 38
	LustreParserRULE_return = 39
	LustreParserRULE_returnVar = 40
	LustreParserRULE_clockedBlock = 41
	LustreParserRULE_ifBlock = 42
	LustreParserRULE_matchBlock = 43
	LustreParserRULE_pattern = 44
	LustreParserRULE_expr = 45
	LustreParserRULE_eID = 46
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTypedef() []ITypedefContext
	Typedef(i int) ITypedefContext
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	AllNode() []INodeContext
	Node(i int) INodeContext
	AllFunction() []IFunctionContext
	Function(i int) IFunctionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(LustreParserEOF, 0)
}

func (s *ProgramContext) AllTypedef() []ITypedefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypedefContext); ok {
			len++
		}
	}

	tst := make([]ITypedefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypedefContext); ok {
			tst[i] = t.(ITypedefContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Typedef(i int) ITypedefContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypedefContext)
}

func (s *ProgramContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ProgramContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Node(i int) INodeContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *ProgramContext) AllFunction() []IFunctionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionContext); ok {
			tst[i] = t.(IFunctionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Function(i int) IFunctionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LustreParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8274) != 0) {
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LustreParserT__0:
			{
				p.SetState(94)
				p.Typedef()
			}


		case LustreParserT__3:
			{
				p.SetState(95)
				p.Constant()
			}


		case LustreParserT__5:
			{
				p.SetState(96)
				p.Node()
			}


		case LustreParserT__12:
			{
				p.SetState(97)
				p.Function()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(LustreParserEOF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	TopLevelType() ITopLevelTypeContext

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_typedef
	return p
}

func InitEmptyTypedefContext(p *TypedefContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_typedef
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_typedef

	return p
}

func (s *TypedefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *TypedefContext) TopLevelType() ITopLevelTypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelTypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelTypeContext)
}

func (s *TypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypedefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterTypedef(s)
	}
}

func (s *TypedefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitTypedef(s)
	}
}

func (s *TypedefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitTypedef(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Typedef() (localctx ITypedefContext) {
	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LustreParserRULE_typedef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(LustreParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(108)
		p.TopLevelType()
	}
	{
		p.SetState(109)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext
	Type_() ITypeContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *ConstantContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstantContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LustreParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(LustreParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__4 {
		{
			p.SetState(113)
			p.Match(LustreParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(114)
			p.type_(0)
		}

	}
	{
		p.SetState(117)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(118)
		p.expr(0)
	}
	{
		p.SetState(119)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInput returns the input rule contexts.
	GetInput() IVarDeclListContext

	// GetOutput returns the output rule contexts.
	GetOutput() IVarDeclListContext

	// GetLocal returns the local rule contexts.
	GetLocal() IVarDeclListContext


	// SetInput sets the input rule contexts.
	SetInput(IVarDeclListContext)

	// SetOutput sets the output rule contexts.
	SetOutput(IVarDeclListContext)

	// SetLocal sets the local rule contexts.
	SetLocal(IVarDeclListContext)


	// Getter signatures
	ID() antlr.TerminalNode
	AllEquation() []IEquationContext
	Equation(i int) IEquationContext
	AllProperty() []IPropertyContext
	Property(i int) IPropertyContext
	AllAssertion() []IAssertionContext
	Assertion(i int) IAssertionContext
	AllMain() []IMainContext
	Main(i int) IMainContext
	AllRealizabilityInputs() []IRealizabilityInputsContext
	RealizabilityInputs(i int) IRealizabilityInputsContext
	AllIvc() []IIvcContext
	Ivc(i int) IIvcContext
	AllStateMachine() []IStateMachineContext
	StateMachine(i int) IStateMachineContext
	AllVarDeclList() []IVarDeclListContext
	VarDeclList(i int) IVarDeclListContext

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	input IVarDeclListContext 
	output IVarDeclListContext 
	local IVarDeclListContext 
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_node
	return p
}

func InitEmptyNodeContext(p *NodeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_node
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) GetInput() IVarDeclListContext { return s.input }

func (s *NodeContext) GetOutput() IVarDeclListContext { return s.output }

func (s *NodeContext) GetLocal() IVarDeclListContext { return s.local }


func (s *NodeContext) SetInput(v IVarDeclListContext) { s.input = v }

func (s *NodeContext) SetOutput(v IVarDeclListContext) { s.output = v }

func (s *NodeContext) SetLocal(v IVarDeclListContext) { s.local = v }


func (s *NodeContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *NodeContext) AllEquation() []IEquationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEquationContext); ok {
			len++
		}
	}

	tst := make([]IEquationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEquationContext); ok {
			tst[i] = t.(IEquationContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Equation(i int) IEquationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEquationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *NodeContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *NodeContext) AllAssertion() []IAssertionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssertionContext); ok {
			len++
		}
	}

	tst := make([]IAssertionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssertionContext); ok {
			tst[i] = t.(IAssertionContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Assertion(i int) IAssertionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssertionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *NodeContext) AllMain() []IMainContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMainContext); ok {
			len++
		}
	}

	tst := make([]IMainContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMainContext); ok {
			tst[i] = t.(IMainContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Main(i int) IMainContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *NodeContext) AllRealizabilityInputs() []IRealizabilityInputsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRealizabilityInputsContext); ok {
			len++
		}
	}

	tst := make([]IRealizabilityInputsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRealizabilityInputsContext); ok {
			tst[i] = t.(IRealizabilityInputsContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) RealizabilityInputs(i int) IRealizabilityInputsContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealizabilityInputsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealizabilityInputsContext)
}

func (s *NodeContext) AllIvc() []IIvcContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIvcContext); ok {
			len++
		}
	}

	tst := make([]IIvcContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIvcContext); ok {
			tst[i] = t.(IIvcContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Ivc(i int) IIvcContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIvcContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIvcContext)
}

func (s *NodeContext) AllStateMachine() []IStateMachineContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateMachineContext); ok {
			len++
		}
	}

	tst := make([]IStateMachineContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateMachineContext); ok {
			tst[i] = t.(IStateMachineContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) StateMachine(i int) IStateMachineContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateMachineContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateMachineContext)
}

func (s *NodeContext) AllVarDeclList() []IVarDeclListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclListContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclListContext); ok {
			tst[i] = t.(IVarDeclListContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) VarDeclList(i int) IVarDeclListContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclListContext)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitNode(s)
	}
}

func (s *NodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitNode(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LustreParserRULE_node)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(LustreParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(124)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).input = _x
		}

	}
	{
		p.SetState(127)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(LustreParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(130)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).output = _x
		}

	}
	{
		p.SetState(133)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__9 {
		{
			p.SetState(134)
			p.Match(LustreParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(135)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).local = _x
		}
		{
			p.SetState(136)
			p.Match(LustreParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(140)
		p.Match(LustreParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 6722414226440320) != 0) || _la == LustreParserID {
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(141)
				p.Equation()
			}


		case 2:
			{
				p.SetState(142)
				p.Property()
			}


		case 3:
			{
				p.SetState(143)
				p.Assertion()
			}


		case 4:
			{
				p.SetState(144)
				p.Main()
			}


		case 5:
			{
				p.SetState(145)
				p.RealizabilityInputs()
			}


		case 6:
			{
				p.SetState(146)
				p.Ivc()
			}


		case 7:
			{
				p.SetState(147)
				p.StateMachine()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(153)
		p.Match(LustreParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__2 {
		{
			p.SetState(154)
			p.Match(LustreParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInput returns the input rule contexts.
	GetInput() IVarDeclListContext

	// GetOutput returns the output rule contexts.
	GetOutput() IVarDeclListContext


	// SetInput sets the input rule contexts.
	SetInput(IVarDeclListContext)

	// SetOutput sets the output rule contexts.
	SetOutput(IVarDeclListContext)


	// Getter signatures
	EID() IEIDContext
	AllVarDeclList() []IVarDeclListContext
	VarDeclList(i int) IVarDeclListContext

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	input IVarDeclListContext 
	output IVarDeclListContext 
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) GetInput() IVarDeclListContext { return s.input }

func (s *FunctionContext) GetOutput() IVarDeclListContext { return s.output }


func (s *FunctionContext) SetInput(v IVarDeclListContext) { s.input = v }

func (s *FunctionContext) SetOutput(v IVarDeclListContext) { s.output = v }


func (s *FunctionContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *FunctionContext) AllVarDeclList() []IVarDeclListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclListContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclListContext); ok {
			tst[i] = t.(IVarDeclListContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) VarDeclList(i int) IVarDeclListContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclListContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LustreParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(LustreParserT__12)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(158)
		p.eID(0)
	}
	{
		p.SetState(159)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(160)

			var _x = p.VarDeclList()


			localctx.(*FunctionContext).input = _x
		}

	}
	{
		p.SetState(163)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(LustreParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(166)

			var _x = p.VarDeclList()


			localctx.(*FunctionContext).output = _x
		}

	}
	{
		p.SetState(169)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVarDeclListContext is an interface to support dynamic dispatch.
type IVarDeclListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarDeclGroup() []IVarDeclGroupContext
	VarDeclGroup(i int) IVarDeclGroupContext

	// IsVarDeclListContext differentiates from other interfaces.
	IsVarDeclListContext()
}

type VarDeclListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclListContext() *VarDeclListContext {
	var p = new(VarDeclListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDeclList
	return p
}

func InitEmptyVarDeclListContext(p *VarDeclListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDeclList
}

func (*VarDeclListContext) IsVarDeclListContext() {}

func NewVarDeclListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclListContext {
	var p = new(VarDeclListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_varDeclList

	return p
}

func (s *VarDeclListContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclListContext) AllVarDeclGroup() []IVarDeclGroupContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclGroupContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclGroupContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclGroupContext); ok {
			tst[i] = t.(IVarDeclGroupContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclListContext) VarDeclGroup(i int) IVarDeclGroupContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclGroupContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclGroupContext)
}

func (s *VarDeclListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarDeclListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterVarDeclList(s)
	}
}

func (s *VarDeclListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitVarDeclList(s)
	}
}

func (s *VarDeclListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitVarDeclList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) VarDeclList() (localctx IVarDeclListContext) {
	localctx = NewVarDeclListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LustreParserRULE_varDeclList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.VarDeclGroup()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(173)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(174)
				p.VarDeclGroup()
			}


		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVarDeclGroupContext is an interface to support dynamic dispatch.
type IVarDeclGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEID() []IEIDContext
	EID(i int) IEIDContext
	Type_() ITypeContext

	// IsVarDeclGroupContext differentiates from other interfaces.
	IsVarDeclGroupContext()
}

type VarDeclGroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclGroupContext() *VarDeclGroupContext {
	var p = new(VarDeclGroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDeclGroup
	return p
}

func InitEmptyVarDeclGroupContext(p *VarDeclGroupContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDeclGroup
}

func (*VarDeclGroupContext) IsVarDeclGroupContext() {}

func NewVarDeclGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclGroupContext {
	var p = new(VarDeclGroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_varDeclGroup

	return p
}

func (s *VarDeclGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclGroupContext) AllEID() []IEIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEIDContext); ok {
			len++
		}
	}

	tst := make([]IEIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEIDContext); ok {
			tst[i] = t.(IEIDContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclGroupContext) EID(i int) IEIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *VarDeclGroupContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarDeclGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterVarDeclGroup(s)
	}
}

func (s *VarDeclGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitVarDeclGroup(s)
	}
}

func (s *VarDeclGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitVarDeclGroup(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) VarDeclGroup() (localctx IVarDeclGroupContext) {
	localctx = NewVarDeclGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LustreParserRULE_varDeclGroup)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.eID(0)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == LustreParserT__13 {
		{
			p.SetState(181)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(182)
			p.eID(0)
		}


		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(LustreParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(189)
		p.type_(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITopLevelTypeContext is an interface to support dynamic dispatch.
type ITopLevelTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTopLevelTypeContext differentiates from other interfaces.
	IsTopLevelTypeContext()
}

type TopLevelTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelTypeContext() *TopLevelTypeContext {
	var p = new(TopLevelTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_topLevelType
	return p
}

func InitEmptyTopLevelTypeContext(p *TopLevelTypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_topLevelType
}

func (*TopLevelTypeContext) IsTopLevelTypeContext() {}

func NewTopLevelTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelTypeContext {
	var p = new(TopLevelTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_topLevelType

	return p
}

func (s *TopLevelTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelTypeContext) CopyAll(ctx *TopLevelTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TopLevelTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type RecordTypeContext struct {
	TopLevelTypeContext
}

func NewRecordTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordTypeContext {
	var p = new(RecordTypeContext)

	InitEmptyTopLevelTypeContext(&p.TopLevelTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelTypeContext))

	return p
}

func (s *RecordTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LustreParserID)
}

func (s *RecordTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LustreParserID, i)
}

func (s *RecordTypeContext) AllType_() []ITypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeContext); ok {
			len++
		}
	}

	tst := make([]ITypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeContext); ok {
			tst[i] = t.(ITypeContext)
			i++
		}
	}

	return tst
}

func (s *RecordTypeContext) Type_(i int) ITypeContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}


func (s *RecordTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRecordType(s)
	}
}

func (s *RecordTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRecordType(s)
	}
}

func (s *RecordTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRecordType(s)

	default:
		return t.VisitChildren(s)
	}
}


type EnumTypeContext struct {
	TopLevelTypeContext
}

func NewEnumTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EnumTypeContext {
	var p = new(EnumTypeContext)

	InitEmptyTopLevelTypeContext(&p.TopLevelTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelTypeContext))

	return p
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LustreParserID)
}

func (s *EnumTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LustreParserID, i)
}


func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (s *EnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}


type PlainTypeContext struct {
	TopLevelTypeContext
}

func NewPlainTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainTypeContext {
	var p = new(PlainTypeContext)

	InitEmptyTopLevelTypeContext(&p.TopLevelTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TopLevelTypeContext))

	return p
}

func (s *PlainTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}


func (s *PlainTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterPlainType(s)
	}
}

func (s *PlainTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitPlainType(s)
	}
}

func (s *PlainTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitPlainType(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LustreParser) TopLevelType() (localctx ITopLevelTypeContext) {
	localctx = NewTopLevelTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LustreParserRULE_topLevelType)
	var _la int

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__18, LustreParserT__19, LustreParserT__23, LustreParserT__24, LustreParserID:
		localctx = NewPlainTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.type_(0)
		}


	case LustreParserT__14:
		localctx = NewRecordTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(LustreParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		{
			p.SetState(194)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Match(LustreParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(196)
			p.type_(0)
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__2 {
			{
				p.SetState(198)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(199)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(200)
				p.Match(LustreParserT__4)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(201)
				p.type_(0)
			}


			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(207)
			p.Match(LustreParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__17:
		localctx = NewEnumTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Match(LustreParserT__17)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(210)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(212)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(213)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(LustreParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type ArrayTypeContext struct {
	TypeContext
}

func NewArrayTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayTypeContext {
	var p = new(ArrayTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *ArrayTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ArrayTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(LustreParserINT, 0)
}


func (s *ArrayTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrayType(s)
	}
}

func (s *ArrayTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrayType(s)
	}
}

func (s *ArrayTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrayType(s)

	default:
		return t.VisitChildren(s)
	}
}


type RealTypeContext struct {
	TypeContext
}

func NewRealTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RealTypeContext {
	var p = new(RealTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *RealTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRealType(s)
	}
}

func (s *RealTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRealType(s)
	}
}

func (s *RealTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRealType(s)

	default:
		return t.VisitChildren(s)
	}
}


type SubrangeTypeContext struct {
	TypeContext
}

func NewSubrangeTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubrangeTypeContext {
	var p = new(SubrangeTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *SubrangeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubrangeTypeContext) AllBound() []IBoundContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoundContext); ok {
			len++
		}
	}

	tst := make([]IBoundContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoundContext); ok {
			tst[i] = t.(IBoundContext)
			i++
		}
	}

	return tst
}

func (s *SubrangeTypeContext) Bound(i int) IBoundContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoundContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoundContext)
}


func (s *SubrangeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterSubrangeType(s)
	}
}

func (s *SubrangeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitSubrangeType(s)
	}
}

func (s *SubrangeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitSubrangeType(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntTypeContext struct {
	TypeContext
}

func NewIntTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntTypeContext {
	var p = new(IntTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *IntTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIntType(s)
	}
}

func (s *IntTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIntType(s)
	}
}

func (s *IntTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIntType(s)

	default:
		return t.VisitChildren(s)
	}
}


type UserTypeContext struct {
	TypeContext
}

func NewUserTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserTypeContext {
	var p = new(UserTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *UserTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *UserTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterUserType(s)
	}
}

func (s *UserTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitUserType(s)
	}
}

func (s *UserTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitUserType(s)

	default:
		return t.VisitChildren(s)
	}
}


type BoolTypeContext struct {
	TypeContext
}

func NewBoolTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolTypeContext {
	var p = new(BoolTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *BoolTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterBoolType(s)
	}
}

func (s *BoolTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitBoolType(s)
	}
}

func (s *BoolTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitBoolType(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LustreParser) Type_() (localctx ITypeContext) {
	return p.type_(0)
}

func (p *LustreParser) type_(_p int) (localctx ITypeContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTypeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, LustreParserRULE_type, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__18:
		localctx = NewIntTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(223)
			p.Match(LustreParserT__18)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__19:
		localctx = NewSubrangeTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(224)
			p.Match(LustreParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(LustreParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Bound()
		}
		{
			p.SetState(227)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Bound()
		}
		{
			p.SetState(229)
			p.Match(LustreParserT__21)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Match(LustreParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(LustreParserT__18)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__23:
		localctx = NewBoolTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(233)
			p.Match(LustreParserT__23)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__24:
		localctx = NewRealTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(LustreParserT__24)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserID:
		localctx = NewUserTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArrayTypeContext(p, NewTypeContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_type)
			p.SetState(238)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(239)
				p.Match(LustreParserT__20)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(240)
				p.Match(LustreParserINT)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(241)
				p.Match(LustreParserT__21)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IBoundContext is an interface to support dynamic dispatch.
type IBoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsBoundContext differentiates from other interfaces.
	IsBoundContext()
}

type BoundContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundContext() *BoundContext {
	var p = new(BoundContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_bound
	return p
}

func InitEmptyBoundContext(p *BoundContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_bound
}

func (*BoundContext) IsBoundContext() {}

func NewBoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundContext {
	var p = new(BoundContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_bound

	return p
}

func (s *BoundContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundContext) INT() antlr.TerminalNode {
	return s.GetToken(LustreParserINT, 0)
}

func (s *BoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterBound(s)
	}
}

func (s *BoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitBound(s)
	}
}

func (s *BoundContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitBound(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Bound() (localctx IBoundContext) {
	localctx = NewBoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LustreParserRULE_bound)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__25 {
		{
			p.SetState(247)
			p.Match(LustreParserT__25)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(250)
		p.Match(LustreParserINT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStateMachineContext is an interface to support dynamic dispatch.
type IStateMachineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllStateDecl() []IStateDeclContext
	StateDecl(i int) IStateDeclContext

	// IsStateMachineContext differentiates from other interfaces.
	IsStateMachineContext()
}

type StateMachineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateMachineContext() *StateMachineContext {
	var p = new(StateMachineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_stateMachine
	return p
}

func InitEmptyStateMachineContext(p *StateMachineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_stateMachine
}

func (*StateMachineContext) IsStateMachineContext() {}

func NewStateMachineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateMachineContext {
	var p = new(StateMachineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_stateMachine

	return p
}

func (s *StateMachineContext) GetParser() antlr.Parser { return s.parser }

func (s *StateMachineContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *StateMachineContext) AllStateDecl() []IStateDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateDeclContext); ok {
			len++
		}
	}

	tst := make([]IStateDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateDeclContext); ok {
			tst[i] = t.(IStateDeclContext)
			i++
		}
	}

	return tst
}

func (s *StateMachineContext) StateDecl(i int) IStateDeclContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateDeclContext)
}

func (s *StateMachineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateMachineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StateMachineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterStateMachine(s)
	}
}

func (s *StateMachineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitStateMachine(s)
	}
}

func (s *StateMachineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitStateMachine(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) StateMachine() (localctx IStateMachineContext) {
	localctx = NewStateMachineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LustreParserRULE_stateMachine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(LustreParserT__26)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(253)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1879048192) != 0) {
		{
			p.SetState(254)
			p.StateDecl()
		}


		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IStateDeclContext is an interface to support dynamic dispatch.
type IStateDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	DataDef() IDataDefContext
	AllTransition() []ITransitionContext
	Transition(i int) ITransitionContext
	Actions() IActionsContext

	// IsStateDeclContext differentiates from other interfaces.
	IsStateDeclContext()
}

type StateDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateDeclContext() *StateDeclContext {
	var p = new(StateDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_stateDecl
	return p
}

func InitEmptyStateDeclContext(p *StateDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_stateDecl
}

func (*StateDeclContext) IsStateDeclContext() {}

func NewStateDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateDeclContext {
	var p = new(StateDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_stateDecl

	return p
}

func (s *StateDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StateDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *StateDeclContext) DataDef() IDataDefContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDefContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
}

func (s *StateDeclContext) AllTransition() []ITransitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITransitionContext); ok {
			len++
		}
	}

	tst := make([]ITransitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITransitionContext); ok {
			tst[i] = t.(ITransitionContext)
			i++
		}
	}

	return tst
}

func (s *StateDeclContext) Transition(i int) ITransitionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITransitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITransitionContext)
}

func (s *StateDeclContext) Actions() IActionsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *StateDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StateDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterStateDecl(s)
	}
}

func (s *StateDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitStateDecl(s)
	}
}

func (s *StateDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitStateDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) StateDecl() (localctx IStateDeclContext) {
	localctx = NewStateDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LustreParserRULE_stateDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__27 {
		{
			p.SetState(260)
			p.Match(LustreParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__28 {
		{
			p.SetState(263)
			p.Match(LustreParserT__28)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(266)
		p.Match(LustreParserT__29)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__30 {
		{
			p.SetState(268)
			p.Match(LustreParserT__30)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LustreParserT__34 {
			{
				p.SetState(269)
				p.Transition()
			}
			{
				p.SetState(270)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(278)
		p.DataDef()
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__31 {
		{
			p.SetState(279)
			p.Match(LustreParserT__31)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__34 {
			{
				p.SetState(280)
				p.Transition()
			}
			{
				p.SetState(281)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(287)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == LustreParserT__32 {
			{
				p.SetState(288)
				p.Match(LustreParserT__32)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == LustreParserT__39 {
				{
					p.SetState(289)
					p.Actions()
				}

			}
			{
				p.SetState(292)
				p.Match(LustreParserT__33)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(293)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITransitionContext is an interface to support dynamic dispatch.
type ITransitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Arrow() IArrowContext

	// IsTransitionContext differentiates from other interfaces.
	IsTransitionContext()
}

type TransitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransitionContext() *TransitionContext {
	var p = new(TransitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_transition
	return p
}

func InitEmptyTransitionContext(p *TransitionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_transition
}

func (*TransitionContext) IsTransitionContext() {}

func NewTransitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransitionContext {
	var p = new(TransitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_transition

	return p
}

func (s *TransitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TransitionContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TransitionContext) Arrow() IArrowContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrowContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrowContext)
}

func (s *TransitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TransitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterTransition(s)
	}
}

func (s *TransitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitTransition(s)
	}
}

func (s *TransitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitTransition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Transition() (localctx ITransitionContext) {
	localctx = NewTransitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LustreParserRULE_transition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(LustreParserT__34)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(299)
		p.expr(0)
	}
	{
		p.SetState(300)
		p.Arrow()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IArrowContext is an interface to support dynamic dispatch.
type IArrowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Fork() IForkContext
	Actions() IActionsContext

	// IsArrowContext differentiates from other interfaces.
	IsArrowContext()
}

type ArrowContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowContext() *ArrowContext {
	var p = new(ArrowContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_arrow
	return p
}

func InitEmptyArrowContext(p *ArrowContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_arrow
}

func (*ArrowContext) IsArrowContext() {}

func NewArrowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowContext {
	var p = new(ArrowContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_arrow

	return p
}

func (s *ArrowContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowContext) Fork() IForkContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForkContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForkContext)
}

func (s *ArrowContext) Actions() IActionsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionsContext)
}

func (s *ArrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrow(s)
	}
}

func (s *ArrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrow(s)
	}
}

func (s *ArrowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrow(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Arrow() (localctx IArrowContext) {
	localctx = NewArrowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LustreParserRULE_arrow)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__39 {
		{
			p.SetState(302)
			p.Actions()
		}

	}
	{
		p.SetState(305)
		p.Fork()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IForkContext is an interface to support dynamic dispatch.
type IForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Target() ITargetContext
	Expr() IExprContext
	Arrow() IArrowContext
	AllElsifFork() []IElsifForkContext
	ElsifFork(i int) IElsifForkContext
	ElseFork() IElseForkContext

	// IsForkContext differentiates from other interfaces.
	IsForkContext()
}

type ForkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForkContext() *ForkContext {
	var p = new(ForkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_fork
	return p
}

func InitEmptyForkContext(p *ForkContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_fork
}

func (*ForkContext) IsForkContext() {}

func NewForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForkContext {
	var p = new(ForkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_fork

	return p
}

func (s *ForkContext) GetParser() antlr.Parser { return s.parser }

func (s *ForkContext) Target() ITargetContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *ForkContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForkContext) Arrow() IArrowContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrowContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrowContext)
}

func (s *ForkContext) AllElsifFork() []IElsifForkContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElsifForkContext); ok {
			len++
		}
	}

	tst := make([]IElsifForkContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElsifForkContext); ok {
			tst[i] = t.(IElsifForkContext)
			i++
		}
	}

	return tst
}

func (s *ForkContext) ElsifFork(i int) IElsifForkContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElsifForkContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElsifForkContext)
}

func (s *ForkContext) ElseFork() IElseForkContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseForkContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseForkContext)
}

func (s *ForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ForkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterFork(s)
	}
}

func (s *ForkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitFork(s)
	}
}

func (s *ForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitFork(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Fork() (localctx IForkContext) {
	localctx = NewForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LustreParserRULE_fork)
	var _alt int

	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__37, LustreParserT__38:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Target()
		}


	case LustreParserT__34:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Match(LustreParserT__34)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(309)
			p.expr(0)
		}
		{
			p.SetState(310)
			p.Arrow()
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(311)
					p.ElsifFork()
				}


			}
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(317)
				p.ElseFork()
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IElsifForkContext is an interface to support dynamic dispatch.
type IElsifForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Arrow() IArrowContext

	// IsElsifForkContext differentiates from other interfaces.
	IsElsifForkContext()
}

type ElsifForkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElsifForkContext() *ElsifForkContext {
	var p = new(ElsifForkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_elsifFork
	return p
}

func InitEmptyElsifForkContext(p *ElsifForkContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_elsifFork
}

func (*ElsifForkContext) IsElsifForkContext() {}

func NewElsifForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsifForkContext {
	var p = new(ElsifForkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_elsifFork

	return p
}

func (s *ElsifForkContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsifForkContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ElsifForkContext) Arrow() IArrowContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrowContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrowContext)
}

func (s *ElsifForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsifForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElsifForkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterElsifFork(s)
	}
}

func (s *ElsifForkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitElsifFork(s)
	}
}

func (s *ElsifForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitElsifFork(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) ElsifFork() (localctx IElsifForkContext) {
	localctx = NewElsifForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LustreParserRULE_elsifFork)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(LustreParserT__35)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(323)
		p.expr(0)
	}
	{
		p.SetState(324)
		p.Arrow()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IElseForkContext is an interface to support dynamic dispatch.
type IElseForkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Arrow() IArrowContext

	// IsElseForkContext differentiates from other interfaces.
	IsElseForkContext()
}

type ElseForkContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseForkContext() *ElseForkContext {
	var p = new(ElseForkContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_elseFork
	return p
}

func InitEmptyElseForkContext(p *ElseForkContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_elseFork
}

func (*ElseForkContext) IsElseForkContext() {}

func NewElseForkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseForkContext {
	var p = new(ElseForkContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_elseFork

	return p
}

func (s *ElseForkContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseForkContext) Arrow() IArrowContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrowContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrowContext)
}

func (s *ElseForkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseForkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ElseForkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterElseFork(s)
	}
}

func (s *ElseForkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitElseFork(s)
	}
}

func (s *ElseForkContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitElseFork(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) ElseFork() (localctx IElseForkContext) {
	localctx = NewElseForkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LustreParserRULE_elseFork)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(LustreParserT__36)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(327)
		p.Arrow()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_target
	return p
}

func InitEmptyTargetContext(p *TargetContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_target
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (s *TargetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitTarget(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LustreParserRULE_target)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__37:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.Match(LustreParserT__37)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__38:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(LustreParserT__38)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IActionsContext is an interface to support dynamic dispatch.
type IActionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DataDef() IDataDefContext
	AllEmissionBody() []IEmissionBodyContext
	EmissionBody(i int) IEmissionBodyContext

	// IsActionsContext differentiates from other interfaces.
	IsActionsContext()
}

type ActionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionsContext() *ActionsContext {
	var p = new(ActionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_actions
	return p
}

func InitEmptyActionsContext(p *ActionsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_actions
}

func (*ActionsContext) IsActionsContext() {}

func NewActionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionsContext {
	var p = new(ActionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_actions

	return p
}

func (s *ActionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionsContext) DataDef() IDataDefContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDefContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
}

func (s *ActionsContext) AllEmissionBody() []IEmissionBodyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmissionBodyContext); ok {
			len++
		}
	}

	tst := make([]IEmissionBodyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmissionBodyContext); ok {
			tst[i] = t.(IEmissionBodyContext)
			i++
		}
	}

	return tst
}

func (s *ActionsContext) EmissionBody(i int) IEmissionBodyContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmissionBodyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmissionBodyContext)
}

func (s *ActionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ActionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterActions(s)
	}
}

func (s *ActionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitActions(s)
	}
}

func (s *ActionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitActions(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Actions() (localctx IActionsContext) {
	localctx = NewActionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LustreParserRULE_actions)
	var _la int

	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.Match(LustreParserT__39)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(336)
			p.DataDef()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.Match(LustreParserT__39)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == LustreParserT__40 {
			{
				p.SetState(339)
				p.Match(LustreParserT__40)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(342)
			p.EmissionBody()
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__2 {
			{
				p.SetState(343)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(344)
				p.Match(LustreParserT__40)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(345)
				p.EmissionBody()
			}


			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(351)
			p.Match(LustreParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDataDefContext is an interface to support dynamic dispatch.
type IDataDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Equation() IEquationContext
	Scope() IScopeContext

	// IsDataDefContext differentiates from other interfaces.
	IsDataDefContext()
}

type DataDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataDefContext() *DataDefContext {
	var p = new(DataDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_dataDef
	return p
}

func InitEmptyDataDefContext(p *DataDefContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_dataDef
}

func (*DataDefContext) IsDataDefContext() {}

func NewDataDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataDefContext {
	var p = new(DataDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_dataDef

	return p
}

func (s *DataDefContext) GetParser() antlr.Parser { return s.parser }

func (s *DataDefContext) Equation() IEquationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEquationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *DataDefContext) Scope() IScopeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *DataDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DataDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterDataDef(s)
	}
}

func (s *DataDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitDataDef(s)
	}
}

func (s *DataDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitDataDef(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) DataDef() (localctx IDataDefContext) {
	localctx = NewDataDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LustreParserRULE_dataDef)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.Equation()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Scope()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LocalBlock() ILocalBlockContext
	Eqs() IEqsContext

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_scope
	return p
}

func InitEmptyScopeContext(p *ScopeContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_scope
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) LocalBlock() ILocalBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocalBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocalBlockContext)
}

func (s *ScopeContext) Eqs() IEqsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqsContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitScope(s)
	}
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Scope() (localctx IScopeContext) {
	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LustreParserRULE_scope)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__9 {
		{
			p.SetState(359)
			p.LocalBlock()
		}

	}
	p.SetState(363)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__10 {
		{
			p.SetState(362)
			p.Eqs()
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILocalBlockContext is an interface to support dynamic dispatch.
type ILocalBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

	// IsLocalBlockContext differentiates from other interfaces.
	IsLocalBlockContext()
}

type LocalBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalBlockContext() *LocalBlockContext {
	var p = new(LocalBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_localBlock
	return p
}

func InitEmptyLocalBlockContext(p *LocalBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_localBlock
}

func (*LocalBlockContext) IsLocalBlockContext() {}

func NewLocalBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalBlockContext {
	var p = new(LocalBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_localBlock

	return p
}

func (s *LocalBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalBlockContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *LocalBlockContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *LocalBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LocalBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterLocalBlock(s)
	}
}

func (s *LocalBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitLocalBlock(s)
	}
}

func (s *LocalBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitLocalBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) LocalBlock() (localctx ILocalBlockContext) {
	localctx = NewLocalBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LustreParserRULE_localBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Match(LustreParserT__9)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(366)
				p.VarDecl()
			}
			{
				p.SetState(367)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEqsContext is an interface to support dynamic dispatch.
type IEqsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEquation() []IEquationContext
	Equation(i int) IEquationContext

	// IsEqsContext differentiates from other interfaces.
	IsEqsContext()
}

type EqsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqsContext() *EqsContext {
	var p = new(EqsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_eqs
	return p
}

func InitEmptyEqsContext(p *EqsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_eqs
}

func (*EqsContext) IsEqsContext() {}

func NewEqsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqsContext {
	var p = new(EqsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_eqs

	return p
}

func (s *EqsContext) GetParser() antlr.Parser { return s.parser }

func (s *EqsContext) AllEquation() []IEquationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEquationContext); ok {
			len++
		}
	}

	tst := make([]IEquationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEquationContext); ok {
			tst[i] = t.(IEquationContext)
			i++
		}
	}

	return tst
}

func (s *EqsContext) Equation(i int) IEquationContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEquationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *EqsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EqsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterEqs(s)
	}
}

func (s *EqsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitEqs(s)
	}
}

func (s *EqsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitEqs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Eqs() (localctx IEqsContext) {
	localctx = NewEqsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LustreParserRULE_eqs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(LustreParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 5631698691686528) != 0) || _la == LustreParserID {
		{
			p.SetState(375)
			p.Equation()
		}
		{
			p.SetState(376)
			p.Match(LustreParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(383)
		p.Match(LustreParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVarID() []IVarIDContext
	VarID(i int) IVarIDContext
	Type_() ITypeContext
	DefaultDecl() IDefaultDeclContext
	LastDecl() ILastDeclContext

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) AllVarID() []IVarIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarIDContext); ok {
			len++
		}
	}

	tst := make([]IVarIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarIDContext); ok {
			tst[i] = t.(IVarIDContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarID(i int) IVarIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarIDContext)
}

func (s *VarDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclContext) DefaultDecl() IDefaultDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultDeclContext)
}

func (s *VarDeclContext) LastDecl() ILastDeclContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILastDeclContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILastDeclContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LustreParserRULE_varDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.VarID()
	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == LustreParserT__13 {
		{
			p.SetState(386)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(387)
			p.VarID()
		}


		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(393)
		p.Match(LustreParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(394)
		p.type_(0)
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__42 {
		{
			p.SetState(395)
			p.DefaultDecl()
		}

	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__43 {
		{
			p.SetState(398)
			p.LastDecl()
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IVarIDContext is an interface to support dynamic dispatch.
type IVarIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVarIDContext differentiates from other interfaces.
	IsVarIDContext()
}

type VarIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarIDContext() *VarIDContext {
	var p = new(VarIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varID
	return p
}

func InitEmptyVarIDContext(p *VarIDContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_varID
}

func (*VarIDContext) IsVarIDContext() {}

func NewVarIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarIDContext {
	var p = new(VarIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_varID

	return p
}

func (s *VarIDContext) GetParser() antlr.Parser { return s.parser }

func (s *VarIDContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *VarIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VarIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterVarID(s)
	}
}

func (s *VarIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitVarID(s)
	}
}

func (s *VarIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitVarID(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) VarID() (localctx IVarIDContext) {
	localctx = NewVarIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LustreParserRULE_varID)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__41 {
		{
			p.SetState(401)
			p.Match(LustreParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(404)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IDefaultDeclContext is an interface to support dynamic dispatch.
type IDefaultDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsDefaultDeclContext differentiates from other interfaces.
	IsDefaultDeclContext()
}

type DefaultDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultDeclContext() *DefaultDeclContext {
	var p = new(DefaultDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_defaultDecl
	return p
}

func InitEmptyDefaultDeclContext(p *DefaultDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_defaultDecl
}

func (*DefaultDeclContext) IsDefaultDeclContext() {}

func NewDefaultDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultDeclContext {
	var p = new(DefaultDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_defaultDecl

	return p
}

func (s *DefaultDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultDeclContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefaultDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DefaultDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterDefaultDecl(s)
	}
}

func (s *DefaultDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitDefaultDecl(s)
	}
}

func (s *DefaultDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitDefaultDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) DefaultDecl() (localctx IDefaultDeclContext) {
	localctx = NewDefaultDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LustreParserRULE_defaultDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(LustreParserT__42)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(408)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILastDeclContext is an interface to support dynamic dispatch.
type ILastDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsLastDeclContext differentiates from other interfaces.
	IsLastDeclContext()
}

type LastDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastDeclContext() *LastDeclContext {
	var p = new(LastDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lastDecl
	return p
}

func InitEmptyLastDeclContext(p *LastDeclContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lastDecl
}

func (*LastDeclContext) IsLastDeclContext() {}

func NewLastDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastDeclContext {
	var p = new(LastDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_lastDecl

	return p
}

func (s *LastDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LastDeclContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LastDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LastDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterLastDecl(s)
	}
}

func (s *LastDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitLastDecl(s)
	}
}

func (s *LastDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitLastDecl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) LastDecl() (localctx ILastDeclContext) {
	localctx = NewLastDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LustreParserRULE_lastDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(LustreParserT__43)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(411)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(412)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EID() IEIDContext

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_property
	return p
}

func InitEmptyPropertyContext(p *PropertyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_property
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LustreParserRULE_property)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(LustreParserT__44)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(415)
		p.eID(0)
	}
	{
		p.SetState(416)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IRealizabilityInputsContext is an interface to support dynamic dispatch.
type IRealizabilityInputsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsRealizabilityInputsContext differentiates from other interfaces.
	IsRealizabilityInputsContext()
}

type RealizabilityInputsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealizabilityInputsContext() *RealizabilityInputsContext {
	var p = new(RealizabilityInputsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_realizabilityInputs
	return p
}

func InitEmptyRealizabilityInputsContext(p *RealizabilityInputsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_realizabilityInputs
}

func (*RealizabilityInputsContext) IsRealizabilityInputsContext() {}

func NewRealizabilityInputsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealizabilityInputsContext {
	var p = new(RealizabilityInputsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_realizabilityInputs

	return p
}

func (s *RealizabilityInputsContext) GetParser() antlr.Parser { return s.parser }

func (s *RealizabilityInputsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LustreParserID)
}

func (s *RealizabilityInputsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LustreParserID, i)
}

func (s *RealizabilityInputsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealizabilityInputsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *RealizabilityInputsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRealizabilityInputs(s)
	}
}

func (s *RealizabilityInputsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRealizabilityInputs(s)
	}
}

func (s *RealizabilityInputsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRealizabilityInputs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) RealizabilityInputs() (localctx IRealizabilityInputsContext) {
	localctx = NewRealizabilityInputsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LustreParserRULE_realizabilityInputs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(LustreParserT__45)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(419)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(420)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(421)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(426)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(429)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IIvcContext is an interface to support dynamic dispatch.
type IIvcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEID() []IEIDContext
	EID(i int) IEIDContext

	// IsIvcContext differentiates from other interfaces.
	IsIvcContext()
}

type IvcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIvcContext() *IvcContext {
	var p = new(IvcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_ivc
	return p
}

func InitEmptyIvcContext(p *IvcContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_ivc
}

func (*IvcContext) IsIvcContext() {}

func NewIvcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IvcContext {
	var p = new(IvcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_ivc

	return p
}

func (s *IvcContext) GetParser() antlr.Parser { return s.parser }

func (s *IvcContext) AllEID() []IEIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEIDContext); ok {
			len++
		}
	}

	tst := make([]IEIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEIDContext); ok {
			tst[i] = t.(IEIDContext)
			i++
		}
	}

	return tst
}

func (s *IvcContext) EID(i int) IEIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *IvcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IvcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IvcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIvc(s)
	}
}

func (s *IvcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIvc(s)
	}
}

func (s *IvcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIvc(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Ivc() (localctx IIvcContext) {
	localctx = NewIvcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LustreParserRULE_ivc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(431)
		p.Match(LustreParserT__46)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(432)
			p.eID(0)
		}
		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(433)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(434)
				p.eID(0)
			}


			p.SetState(439)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(442)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_main
	return p
}

func InitEmptyMainContext(p *MainContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_main
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }
func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitMain(s)
	}
}

func (s *MainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitMain(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LustreParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(LustreParserT__47)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__2 {
		{
			p.SetState(445)
			p.Match(LustreParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAssertionContext is an interface to support dynamic dispatch.
type IAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsAssertionContext differentiates from other interfaces.
	IsAssertionContext()
}

type AssertionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertionContext() *AssertionContext {
	var p = new(AssertionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_assertion
	return p
}

func InitEmptyAssertionContext(p *AssertionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_assertion
}

func (*AssertionContext) IsAssertionContext() {}

func NewAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionContext {
	var p = new(AssertionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_assertion

	return p
}

func (s *AssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterAssertion(s)
	}
}

func (s *AssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitAssertion(s)
	}
}

func (s *AssertionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitAssertion(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Assertion() (localctx IAssertionContext) {
	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LustreParserRULE_assertion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(448)
		p.Match(LustreParserT__48)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(449)
		p.expr(0)
	}
	{
		p.SetState(450)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleEquation() ISimpleEquationContext
	Emission() IEmissionContext
	ControlBlock() IControlBlockContext
	Return_() IReturnContext

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_equation
	return p
}

func InitEmptyEquationContext(p *EquationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_equation
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) SimpleEquation() ISimpleEquationContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleEquationContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleEquationContext)
}

func (s *EquationContext) Emission() IEmissionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmissionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmissionContext)
}

func (s *EquationContext) ControlBlock() IControlBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControlBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControlBlockContext)
}

func (s *EquationContext) Return_() IReturnContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnContext)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (s *EquationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitEquation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LustreParserRULE_equation)
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__6, LustreParserT__49, LustreParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.SimpleEquation()
		}


	case LustreParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.Emission()
		}


	case LustreParserT__26, LustreParserT__51:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(454)
			p.ControlBlock()
		}
		{
			p.SetState(455)
			p.Return_()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISimpleEquationContext is an interface to support dynamic dispatch.
type ISimpleEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lhs() ILhsContext
	Expr() IExprContext

	// IsSimpleEquationContext differentiates from other interfaces.
	IsSimpleEquationContext()
}

type SimpleEquationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleEquationContext() *SimpleEquationContext {
	var p = new(SimpleEquationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_simpleEquation
	return p
}

func InitEmptySimpleEquationContext(p *SimpleEquationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_simpleEquation
}

func (*SimpleEquationContext) IsSimpleEquationContext() {}

func NewSimpleEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleEquationContext {
	var p = new(SimpleEquationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_simpleEquation

	return p
}

func (s *SimpleEquationContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleEquationContext) Lhs() ILhsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsContext)
}

func (s *SimpleEquationContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SimpleEquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleEquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleEquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterSimpleEquation(s)
	}
}

func (s *SimpleEquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitSimpleEquation(s)
	}
}

func (s *SimpleEquationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitSimpleEquation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) SimpleEquation() (localctx ISimpleEquationContext) {
	localctx = NewSimpleEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LustreParserRULE_simpleEquation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(459)
		p.Lhs()
	}
	{
		p.SetState(460)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(461)
		p.expr(0)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLhsID() []ILhsIDContext
	LhsID(i int) ILhsIDContext

	// IsLhsContext differentiates from other interfaces.
	IsLhsContext()
}

type LhsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsContext() *LhsContext {
	var p = new(LhsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lhs
	return p
}

func InitEmptyLhsContext(p *LhsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lhs
}

func (*LhsContext) IsLhsContext() {}

func NewLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsContext {
	var p = new(LhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_lhs

	return p
}

func (s *LhsContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsContext) AllLhsID() []ILhsIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILhsIDContext); ok {
			len++
		}
	}

	tst := make([]ILhsIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILhsIDContext); ok {
			tst[i] = t.(ILhsIDContext)
			i++
		}
	}

	return tst
}

func (s *LhsContext) LhsID(i int) ILhsIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhsIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILhsIDContext)
}

func (s *LhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterLhs(s)
	}
}

func (s *LhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitLhs(s)
	}
}

func (s *LhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitLhs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LustreParserRULE_lhs)
	var _la int

	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(464)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case LustreParserT__49, LustreParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.LhsID()
		}
		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(466)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(467)
				p.LhsID()
			}


			p.SetState(472)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILhsIDContext is an interface to support dynamic dispatch.
type ILhsIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EID() IEIDContext

	// IsLhsIDContext differentiates from other interfaces.
	IsLhsIDContext()
}

type LhsIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhsIDContext() *LhsIDContext {
	var p = new(LhsIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lhsID
	return p
}

func InitEmptyLhsIDContext(p *LhsIDContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_lhsID
}

func (*LhsIDContext) IsLhsIDContext() {}

func NewLhsIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsIDContext {
	var p = new(LhsIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_lhsID

	return p
}

func (s *LhsIDContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsIDContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *LhsIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LhsIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterLhsID(s)
	}
}

func (s *LhsIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitLhsID(s)
	}
}

func (s *LhsIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitLhsID(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) LhsID() (localctx ILhsIDContext) {
	localctx = NewLhsIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LustreParserRULE_lhsID)
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.eID(0)
		}


	case LustreParserT__49:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Match(LustreParserT__49)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IControlBlockContext is an interface to support dynamic dispatch.
type IControlBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StateMachine() IStateMachineContext
	ClockedBlock() IClockedBlockContext

	// IsControlBlockContext differentiates from other interfaces.
	IsControlBlockContext()
}

type ControlBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlBlockContext() *ControlBlockContext {
	var p = new(ControlBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_controlBlock
	return p
}

func InitEmptyControlBlockContext(p *ControlBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_controlBlock
}

func (*ControlBlockContext) IsControlBlockContext() {}

func NewControlBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlBlockContext {
	var p = new(ControlBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_controlBlock

	return p
}

func (s *ControlBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlBlockContext) StateMachine() IStateMachineContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateMachineContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateMachineContext)
}

func (s *ControlBlockContext) ClockedBlock() IClockedBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClockedBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClockedBlockContext)
}

func (s *ControlBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ControlBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterControlBlock(s)
	}
}

func (s *ControlBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitControlBlock(s)
	}
}

func (s *ControlBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitControlBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) ControlBlock() (localctx IControlBlockContext) {
	localctx = NewControlBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LustreParserRULE_controlBlock)
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(479)
			p.StateMachine()
		}


	case LustreParserT__51:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(480)
			p.ClockedBlock()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEmissionContext is an interface to support dynamic dispatch.
type IEmissionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EmissionBody() IEmissionBodyContext

	// IsEmissionContext differentiates from other interfaces.
	IsEmissionContext()
}

type EmissionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmissionContext() *EmissionContext {
	var p = new(EmissionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_emission
	return p
}

func InitEmptyEmissionContext(p *EmissionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_emission
}

func (*EmissionContext) IsEmissionContext() {}

func NewEmissionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmissionContext {
	var p = new(EmissionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_emission

	return p
}

func (s *EmissionContext) GetParser() antlr.Parser { return s.parser }

func (s *EmissionContext) EmissionBody() IEmissionBodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmissionBodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmissionBodyContext)
}

func (s *EmissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EmissionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterEmission(s)
	}
}

func (s *EmissionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitEmission(s)
	}
}

func (s *EmissionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitEmission(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Emission() (localctx IEmissionContext) {
	localctx = NewEmissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, LustreParserRULE_emission)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(LustreParserT__40)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(484)
		p.EmissionBody()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEmissionBodyContext is an interface to support dynamic dispatch.
type IEmissionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEID() []IEIDContext
	EID(i int) IEIDContext
	Expr() IExprContext

	// IsEmissionBodyContext differentiates from other interfaces.
	IsEmissionBodyContext()
}

type EmissionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmissionBodyContext() *EmissionBodyContext {
	var p = new(EmissionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_emissionBody
	return p
}

func InitEmptyEmissionBodyContext(p *EmissionBodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_emissionBody
}

func (*EmissionBodyContext) IsEmissionBodyContext() {}

func NewEmissionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmissionBodyContext {
	var p = new(EmissionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_emissionBody

	return p
}

func (s *EmissionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EmissionBodyContext) AllEID() []IEIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEIDContext); ok {
			len++
		}
	}

	tst := make([]IEIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEIDContext); ok {
			tst[i] = t.(IEIDContext)
			i++
		}
	}

	return tst
}

func (s *EmissionBodyContext) EID(i int) IEIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *EmissionBodyContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EmissionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmissionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EmissionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterEmissionBody(s)
	}
}

func (s *EmissionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitEmissionBody(s)
	}
}

func (s *EmissionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitEmissionBody(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) EmissionBody() (localctx IEmissionBodyContext) {
	localctx = NewEmissionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LustreParserRULE_emissionBody)
	var _la int

	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.eID(0)
		}
		p.SetState(489)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(487)
				p.Match(LustreParserT__34)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(488)
				p.expr(0)
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


	case LustreParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(491)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(492)
			p.eID(0)
		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(493)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(494)
				p.eID(0)
			}


			p.SetState(499)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(500)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(501)
				p.Match(LustreParserT__34)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(502)
				p.expr(0)
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IReturnContext is an interface to support dynamic dispatch.
type IReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReturnVar() IReturnVarContext

	// IsReturnContext differentiates from other interfaces.
	IsReturnContext()
}

type ReturnContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnContext() *ReturnContext {
	var p = new(ReturnContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_return
	return p
}

func InitEmptyReturnContext(p *ReturnContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_return
}

func (*ReturnContext) IsReturnContext() {}

func NewReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnContext {
	var p = new(ReturnContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_return

	return p
}

func (s *ReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnContext) ReturnVar() IReturnVarContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnVarContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnVarContext)
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Return_() (localctx IReturnContext) {
	localctx = NewReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, LustreParserRULE_return)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(LustreParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(508)
		p.ReturnVar()
	}
	{
		p.SetState(509)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IReturnVarContext is an interface to support dynamic dispatch.
type IReturnVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEID() []IEIDContext
	EID(i int) IEIDContext

	// IsReturnVarContext differentiates from other interfaces.
	IsReturnVarContext()
}

type ReturnVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnVarContext() *ReturnVarContext {
	var p = new(ReturnVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_returnVar
	return p
}

func InitEmptyReturnVarContext(p *ReturnVarContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_returnVar
}

func (*ReturnVarContext) IsReturnVarContext() {}

func NewReturnVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnVarContext {
	var p = new(ReturnVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_returnVar

	return p
}

func (s *ReturnVarContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnVarContext) AllEID() []IEIDContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEIDContext); ok {
			len++
		}
	}

	tst := make([]IEIDContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEIDContext); ok {
			tst[i] = t.(IEIDContext)
			i++
		}
	}

	return tst
}

func (s *ReturnVarContext) EID(i int) IEIDContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *ReturnVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ReturnVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterReturnVar(s)
	}
}

func (s *ReturnVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitReturnVar(s)
	}
}

func (s *ReturnVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitReturnVar(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) ReturnVar() (localctx IReturnVarContext) {
	localctx = NewReturnVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, LustreParserRULE_returnVar)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(511)
				p.eID(0)
			}
			{
				p.SetState(512)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserID:
		{
			p.SetState(519)
			p.eID(0)
		}


	case LustreParserT__50:
		{
			p.SetState(520)
			p.Match(LustreParserT__50)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IClockedBlockContext is an interface to support dynamic dispatch.
type IClockedBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EID() IEIDContext
	IfBlock() IIfBlockContext
	MatchBlock() IMatchBlockContext

	// IsClockedBlockContext differentiates from other interfaces.
	IsClockedBlockContext()
}

type ClockedBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClockedBlockContext() *ClockedBlockContext {
	var p = new(ClockedBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_clockedBlock
	return p
}

func InitEmptyClockedBlockContext(p *ClockedBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_clockedBlock
}

func (*ClockedBlockContext) IsClockedBlockContext() {}

func NewClockedBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClockedBlockContext {
	var p = new(ClockedBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_clockedBlock

	return p
}

func (s *ClockedBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ClockedBlockContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *ClockedBlockContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *ClockedBlockContext) MatchBlock() IMatchBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchBlockContext)
}

func (s *ClockedBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClockedBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ClockedBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterClockedBlock(s)
	}
}

func (s *ClockedBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitClockedBlock(s)
	}
}

func (s *ClockedBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitClockedBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) ClockedBlock() (localctx IClockedBlockContext) {
	localctx = NewClockedBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, LustreParserRULE_clockedBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(523)
		p.Match(LustreParserT__51)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(524)
		p.eID(0)
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__34:
		{
			p.SetState(525)
			p.IfBlock()
		}


	case LustreParserT__53:
		{
			p.SetState(526)
			p.MatchBlock()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllDataDef() []IDataDefContext
	DataDef(i int) IDataDefContext
	AllIfBlock() []IIfBlockContext
	IfBlock(i int) IIfBlockContext

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfBlockContext) AllDataDef() []IDataDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataDefContext); ok {
			len++
		}
	}

	tst := make([]IDataDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataDefContext); ok {
			tst[i] = t.(IDataDefContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) DataDef(i int) IDataDefContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
}

func (s *IfBlockContext) AllIfBlock() []IIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockContext); ok {
			tst[i] = t.(IIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockContext) IfBlock(i int) IIfBlockContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (s *IfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, LustreParserRULE_ifBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		p.Match(LustreParserT__34)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(530)
		p.expr(0)
	}
	{
		p.SetState(531)
		p.Match(LustreParserT__52)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__6, LustreParserT__9, LustreParserT__10, LustreParserT__26, LustreParserT__36, LustreParserT__40, LustreParserT__49, LustreParserT__51, LustreParserID:
		{
			p.SetState(532)
			p.DataDef()
		}


	case LustreParserT__34:
		{
			p.SetState(533)
			p.IfBlock()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(536)
		p.Match(LustreParserT__36)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__6, LustreParserT__8, LustreParserT__9, LustreParserT__10, LustreParserT__26, LustreParserT__36, LustreParserT__40, LustreParserT__49, LustreParserT__51, LustreParserID:
		{
			p.SetState(537)
			p.DataDef()
		}


	case LustreParserT__34:
		{
			p.SetState(538)
			p.IfBlock()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IMatchBlockContext is an interface to support dynamic dispatch.
type IMatchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	AllPattern() []IPatternContext
	Pattern(i int) IPatternContext
	AllDataDef() []IDataDefContext
	DataDef(i int) IDataDefContext

	// IsMatchBlockContext differentiates from other interfaces.
	IsMatchBlockContext()
}

type MatchBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchBlockContext() *MatchBlockContext {
	var p = new(MatchBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_matchBlock
	return p
}

func InitEmptyMatchBlockContext(p *MatchBlockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_matchBlock
}

func (*MatchBlockContext) IsMatchBlockContext() {}

func NewMatchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchBlockContext {
	var p = new(MatchBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_matchBlock

	return p
}

func (s *MatchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchBlockContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MatchBlockContext) AllPattern() []IPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPatternContext); ok {
			len++
		}
	}

	tst := make([]IPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPatternContext); ok {
			tst[i] = t.(IPatternContext)
			i++
		}
	}

	return tst
}

func (s *MatchBlockContext) Pattern(i int) IPatternContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *MatchBlockContext) AllDataDef() []IDataDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataDefContext); ok {
			len++
		}
	}

	tst := make([]IDataDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataDefContext); ok {
			tst[i] = t.(IDataDefContext)
			i++
		}
	}

	return tst
}

func (s *MatchBlockContext) DataDef(i int) IDataDefContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDataDefContext)
}

func (s *MatchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MatchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterMatchBlock(s)
	}
}

func (s *MatchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitMatchBlock(s)
	}
}

func (s *MatchBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitMatchBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) MatchBlock() (localctx IMatchBlockContext) {
	localctx = NewMatchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, LustreParserRULE_matchBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.Match(LustreParserT__53)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(542)
		p.expr(0)
	}
	{
		p.SetState(543)
		p.Match(LustreParserT__54)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == LustreParserT__55 {
		{
			p.SetState(544)
			p.Match(LustreParserT__55)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Pattern()
		}
		{
			p.SetState(546)
			p.Match(LustreParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(547)
			p.DataDef()
		}


		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	REAL() antlr.TerminalNode
	BOOL() antlr.TerminalNode

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}

func (s *PatternContext) INT() antlr.TerminalNode {
	return s.GetToken(LustreParserINT, 0)
}

func (s *PatternContext) REAL() antlr.TerminalNode {
	return s.GetToken(LustreParserREAL, 0)
}

func (s *PatternContext) BOOL() antlr.TerminalNode {
	return s.GetToken(LustreParserBOOL, 0)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (s *PatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitPattern(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LustreParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, LustreParserRULE_pattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 78)) & ^0x3f) == 0 && ((int64(1) << (_la - 78)) & 15) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type RecordExprContext struct {
	ExprContext
}

func NewRecordExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordExprContext {
	var p = new(RecordExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RecordExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LustreParserID)
}

func (s *RecordExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LustreParserID, i)
}

func (s *RecordExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RecordExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *RecordExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRecordExpr(s)
	}
}

func (s *RecordExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRecordExpr(s)
	}
}

func (s *RecordExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRecordExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(LustreParserINT, 0)
}


func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayExprContext struct {
	ExprContext
}

func NewArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExprContext {
	var p = new(ArrayExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrayExpr(s)
	}
}

func (s *ArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrayExpr(s)
	}
}

func (s *ArrayExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrayExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type CastExprContext struct {
	ExprContext
	op antlr.Token
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext {
	var p = new(CastExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}


func (s *CastExprContext) GetOp() antlr.Token { return s.op }


func (s *CastExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CastExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *CastExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterCastExpr(s)
	}
}

func (s *CastExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitCastExpr(s)
	}
}

func (s *CastExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitCastExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type RealExprContext struct {
	ExprContext
}

func NewRealExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RealExprContext {
	var p = new(RealExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RealExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealExprContext) REAL() antlr.TerminalNode {
	return s.GetToken(LustreParserREAL, 0)
}


func (s *RealExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRealExpr(s)
	}
}

func (s *RealExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRealExpr(s)
	}
}

func (s *RealExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRealExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type IfThenElseExprContext struct {
	ExprContext
}

func NewIfThenElseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfThenElseExprContext {
	var p = new(IfThenElseExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfThenElseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfThenElseExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IfThenElseExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *IfThenElseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIfThenElseExpr(s)
	}
}

func (s *IfThenElseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIfThenElseExpr(s)
	}
}

func (s *IfThenElseExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIfThenElseExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type BinaryExprContext struct {
	ExprContext
	op antlr.Token
}

func NewBinaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExprContext {
	var p = new(BinaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}


func (s *BinaryExprContext) GetOp() antlr.Token { return s.op }


func (s *BinaryExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *BinaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterBinaryExpr(s)
	}
}

func (s *BinaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitBinaryExpr(s)
	}
}

func (s *BinaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitBinaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type PreExprContext struct {
	ExprContext
}

func NewPreExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreExprContext {
	var p = new(PreExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PreExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *PreExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterPreExpr(s)
	}
}

func (s *PreExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitPreExpr(s)
	}
}

func (s *PreExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitPreExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type RecordAccessExprContext struct {
	ExprContext
}

func NewRecordAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordAccessExprContext {
	var p = new(RecordAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RecordAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordAccessExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RecordAccessExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *RecordAccessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRecordAccessExpr(s)
	}
}

func (s *RecordAccessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRecordAccessExpr(s)
	}
}

func (s *RecordAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRecordAccessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type NegateExprContext struct {
	ExprContext
}

func NewNegateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegateExprContext {
	var p = new(NegateExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NegateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegateExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *NegateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterNegateExpr(s)
	}
}

func (s *NegateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitNegateExpr(s)
	}
}

func (s *NegateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitNegateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type CondactExprContext struct {
	ExprContext
}

func NewCondactExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CondactExprContext {
	var p = new(CondactExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CondactExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondactExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CondactExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *CondactExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterCondactExpr(s)
	}
}

func (s *CondactExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitCondactExpr(s)
	}
}

func (s *CondactExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitCondactExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type NotExprContext struct {
	ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayAccessExprContext struct {
	ExprContext
}

func NewArrayAccessExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessExprContext {
	var p = new(ArrayAccessExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayAccessExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayAccessExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ArrayAccessExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrayAccessExpr(s)
	}
}

func (s *ArrayAccessExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrayAccessExpr(s)
	}
}

func (s *ArrayAccessExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrayAccessExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayUpdateExprContext struct {
	ExprContext
}

func NewArrayUpdateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayUpdateExprContext {
	var p = new(ArrayUpdateExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrayUpdateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayUpdateExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayUpdateExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *ArrayUpdateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrayUpdateExpr(s)
	}
}

func (s *ArrayUpdateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrayUpdateExpr(s)
	}
}

func (s *ArrayUpdateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrayUpdateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(LustreParserBOOL, 0)
}


func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type CallExprContext struct {
	ExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *CallExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type TupleExprContext struct {
	ExprContext
}

func NewTupleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleExprContext {
	var p = new(TupleExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TupleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *TupleExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *TupleExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterTupleExpr(s)
	}
}

func (s *TupleExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitTupleExpr(s)
	}
}

func (s *TupleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitTupleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type RecordUpdateExprContext struct {
	ExprContext
}

func NewRecordUpdateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordUpdateExprContext {
	var p = new(RecordUpdateExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *RecordUpdateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordUpdateExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *RecordUpdateExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RecordUpdateExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *RecordUpdateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRecordUpdateExpr(s)
	}
}

func (s *RecordUpdateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRecordUpdateExpr(s)
	}
}

func (s *RecordUpdateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRecordUpdateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LustreParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LustreParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 90
	p.EnterRecursionRule(localctx, 90, LustreParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(641)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(556)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(557)
			p.Match(LustreParserINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		localctx = NewRealExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(558)
			p.Match(LustreParserREAL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(559)
			p.Match(LustreParserBOOL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewCastExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(560)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CastExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == LustreParserT__24 || _la == LustreParserT__56) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CastExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(561)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(562)
			p.expr(0)
		}
		{
			p.SetState(563)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(565)
			p.eID(0)
		}
		{
			p.SetState(566)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 7349874626331148416) != 0) || ((int64((_la - 78)) & ^0x3f) == 0 && ((int64(1) << (_la - 78)) & 15) != 0) {
			{
				p.SetState(567)
				p.expr(0)
			}
			p.SetState(572)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == LustreParserT__13 {
				{
					p.SetState(568)
					p.Match(LustreParserT__13)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(569)
					p.expr(0)
				}


				p.SetState(574)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(577)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 7:
		localctx = NewCondactExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(579)
			p.Match(LustreParserT__57)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(580)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(581)
			p.expr(0)
		}
		p.SetState(584)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LustreParserT__13 {
			{
				p.SetState(582)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(583)
				p.expr(0)
			}


			p.SetState(586)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(588)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 8:
		localctx = NewPreExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(590)
			p.Match(LustreParserT__60)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(591)
			p.expr(14)
		}


	case 9:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(592)
			p.Match(LustreParserT__61)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(593)
			p.expr(13)
		}


	case 10:
		localctx = NewNegateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(594)
			p.Match(LustreParserT__25)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(595)
			p.expr(12)
		}


	case 11:
		localctx = NewIfThenElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(596)
			p.Match(LustreParserT__34)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(597)
			p.expr(0)
		}
		{
			p.SetState(598)
			p.Match(LustreParserT__52)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(599)
			p.expr(0)
		}
		{
			p.SetState(600)
			p.Match(LustreParserT__36)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(601)
			p.expr(4)
		}


	case 12:
		localctx = NewRecordExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(603)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(605)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(606)
			p.Match(LustreParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(607)
			p.expr(0)
		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__2 {
			{
				p.SetState(608)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(609)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(610)
				p.Match(LustreParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(611)
				p.expr(0)
			}


			p.SetState(616)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(617)
			p.Match(LustreParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 13:
		localctx = NewArrayExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(619)
			p.Match(LustreParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(620)
			p.expr(0)
		}
		p.SetState(625)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(621)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(622)
				p.expr(0)
			}


			p.SetState(627)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(628)
			p.Match(LustreParserT__21)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 14:
		localctx = NewTupleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(630)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(631)
			p.expr(0)
		}
		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(632)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(633)
				p.expr(0)
			}


			p.SetState(638)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(639)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(688)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(686)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 72, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(643)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(644)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((int64((_la - 63)) & ^0x3f) == 0 && ((int64(1) << (_la - 63)) & 15) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(645)
					p.expr(12)
				}


			case 2:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(646)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(647)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LustreParserT__25 || _la == LustreParserT__66) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(648)
					p.expr(11)
				}


			case 3:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(649)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(650)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LustreParserT__1 || ((int64((_la - 68)) & ^0x3f) == 0 && ((int64(1) << (_la - 68)) & 31) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(651)
					p.expr(10)
				}


			case 4:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(652)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(653)

					var _m = p.Match(LustreParserT__72)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(654)
					p.expr(9)
				}


			case 5:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(655)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(656)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LustreParserT__73 || _la == LustreParserT__74) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(657)
					p.expr(8)
				}


			case 6:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(658)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(659)

					var _m = p.Match(LustreParserT__75)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(660)
					p.expr(6)
				}


			case 7:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(661)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(662)

					var _m = p.Match(LustreParserT__76)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(663)
					p.expr(5)
				}


			case 8:
				localctx = NewRecordAccessExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(664)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(665)
					p.Match(LustreParserT__58)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(666)
					p.Match(LustreParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 9:
				localctx = NewRecordUpdateExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(667)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(668)
					p.Match(LustreParserT__15)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(669)
					p.Match(LustreParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(670)
					p.Match(LustreParserT__59)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(671)
					p.expr(0)
				}
				{
					p.SetState(672)
					p.Match(LustreParserT__16)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 10:
				localctx = NewArrayAccessExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(674)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(675)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(676)
					p.expr(0)
				}
				{
					p.SetState(677)
					p.Match(LustreParserT__21)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 11:
				localctx = NewArrayUpdateExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(679)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(680)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(681)
					p.expr(0)
				}
				{
					p.SetState(682)
					p.Match(LustreParserT__59)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(683)
					p.expr(0)
				}
				{
					p.SetState(684)
					p.Match(LustreParserT__21)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(690)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IEIDContext is an interface to support dynamic dispatch.
type IEIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEIDContext differentiates from other interfaces.
	IsEIDContext()
}

type EIDContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEIDContext() *EIDContext {
	var p = new(EIDContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_eID
	return p
}

func InitEmptyEIDContext(p *EIDContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LustreParserRULE_eID
}

func (*EIDContext) IsEIDContext() {}

func NewEIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EIDContext {
	var p = new(EIDContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LustreParserRULE_eID

	return p
}

func (s *EIDContext) GetParser() antlr.Parser { return s.parser }

func (s *EIDContext) CopyAll(ctx *EIDContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type BaseEIDContext struct {
	EIDContext
}

func NewBaseEIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseEIDContext {
	var p = new(BaseEIDContext)

	InitEmptyEIDContext(&p.EIDContext)
	p.parser = parser
	p.CopyAll(ctx.(*EIDContext))

	return p
}

func (s *BaseEIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseEIDContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *BaseEIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterBaseEID(s)
	}
}

func (s *BaseEIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitBaseEID(s)
	}
}

func (s *BaseEIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitBaseEID(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayEIDContext struct {
	EIDContext
}

func NewArrayEIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayEIDContext {
	var p = new(ArrayEIDContext)

	InitEmptyEIDContext(&p.EIDContext)
	p.parser = parser
	p.CopyAll(ctx.(*EIDContext))

	return p
}

func (s *ArrayEIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayEIDContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *ArrayEIDContext) INT() antlr.TerminalNode {
	return s.GetToken(LustreParserINT, 0)
}


func (s *ArrayEIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterArrayEID(s)
	}
}

func (s *ArrayEIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitArrayEID(s)
	}
}

func (s *ArrayEIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitArrayEID(s)

	default:
		return t.VisitChildren(s)
	}
}


type RecordEIDContext struct {
	EIDContext
}

func NewRecordEIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecordEIDContext {
	var p = new(RecordEIDContext)

	InitEmptyEIDContext(&p.EIDContext)
	p.parser = parser
	p.CopyAll(ctx.(*EIDContext))

	return p
}

func (s *RecordEIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecordEIDContext) EID() IEIDContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEIDContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEIDContext)
}

func (s *RecordEIDContext) ID() antlr.TerminalNode {
	return s.GetToken(LustreParserID, 0)
}


func (s *RecordEIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.EnterRecordEID(s)
	}
}

func (s *RecordEIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LustreListener); ok {
		listenerT.ExitRecordEID(s)
	}
}

func (s *RecordEIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LustreVisitor:
		return t.VisitRecordEID(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LustreParser) EID() (localctx IEIDContext) {
	return p.eID(0)
}

func (p *LustreParser) eID(_p int) (localctx IEIDContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEIDContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEIDContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 92
	p.EnterRecursionRule(localctx, 92, LustreParserRULE_eID, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewBaseEIDContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(692)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(703)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(701)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 74, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArrayEIDContext(p, NewEIDContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_eID)
				p.SetState(694)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(695)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(696)
					p.Match(LustreParserINT)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(697)
					p.Match(LustreParserT__21)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 2:
				localctx = NewRecordEIDContext(p, NewEIDContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_eID)
				p.SetState(698)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(699)
					p.Match(LustreParserT__58)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(700)
					p.Match(LustreParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(705)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 75, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


func (p *LustreParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
			var t *TypeContext = nil
			if localctx != nil { t = localctx.(*TypeContext) }
			return p.Type__Sempred(t, predIndex)

	case 45:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)

	case 46:
			var t *EIDContext = nil
			if localctx != nil { t = localctx.(*EIDContext) }
			return p.EID_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LustreParser) Type__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LustreParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
			return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
			return p.Precpred(p.GetParserRuleContext(), 16)

	case 11:
			return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LustreParser) EID_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}


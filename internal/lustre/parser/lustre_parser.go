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
    "program", "typedef", "constant", "node", "function", "varDeclList", 
    "varDeclGroup", "topLevelType", "type", "bound", "property", "realizabilityInputs", 
    "ivc", "main", "assertion", "equation", "lhs", "expr", "eID",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 63, 406, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 
	43, 8, 0, 10, 0, 12, 0, 46, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 60, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 
	1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 76, 
	8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 84, 8, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 93, 8, 3, 10, 3, 12, 3, 96, 9, 3, 1, 3, 
	1, 3, 3, 3, 100, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 106, 8, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 3, 4, 112, 8, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 
	5, 120, 8, 5, 10, 5, 12, 5, 123, 9, 5, 1, 6, 1, 6, 1, 6, 5, 6, 128, 8, 
	6, 10, 6, 12, 6, 131, 9, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 
	7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 147, 8, 7, 10, 7, 12, 7, 150, 
	9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 159, 8, 7, 10, 7, 
	12, 7, 162, 9, 7, 1, 7, 3, 7, 165, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 181, 8, 8, 
	1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 187, 8, 8, 10, 8, 12, 8, 190, 9, 8, 1, 9, 
	3, 9, 193, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 
	1, 11, 1, 11, 5, 11, 205, 8, 11, 10, 11, 12, 11, 208, 9, 11, 3, 11, 210, 
	8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 218, 8, 12, 10, 
	12, 12, 12, 221, 9, 12, 3, 12, 223, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 
	3, 13, 229, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 
	15, 238, 8, 15, 1, 15, 3, 15, 241, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 
	16, 1, 16, 1, 16, 5, 16, 250, 8, 16, 10, 16, 12, 16, 253, 9, 16, 1, 17, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 5, 17, 270, 8, 17, 10, 17, 12, 17, 273, 9, 17, 
	3, 17, 275, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 4, 
	17, 284, 8, 17, 11, 17, 12, 17, 285, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 312, 
	8, 17, 10, 17, 12, 17, 315, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 
	17, 5, 17, 323, 8, 17, 10, 17, 12, 17, 326, 9, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 1, 17, 1, 17, 5, 17, 334, 8, 17, 10, 17, 12, 17, 337, 9, 17, 1, 
	17, 1, 17, 3, 17, 341, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 386, 8, 17, 10, 17, 12, 17, 
	389, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 
	18, 1, 18, 5, 18, 401, 8, 18, 10, 18, 12, 18, 404, 9, 18, 1, 18, 0, 3, 
	16, 34, 36, 19, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 
	30, 32, 34, 36, 0, 5, 2, 0, 25, 25, 32, 32, 1, 0, 38, 41, 2, 0, 26, 26, 
	42, 42, 2, 0, 2, 2, 43, 47, 1, 0, 49, 50, 455, 0, 44, 1, 0, 0, 0, 2, 49, 
	1, 0, 0, 0, 4, 55, 1, 0, 0, 0, 6, 65, 1, 0, 0, 0, 8, 101, 1, 0, 0, 0, 10, 
	116, 1, 0, 0, 0, 12, 124, 1, 0, 0, 0, 14, 164, 1, 0, 0, 0, 16, 180, 1, 
	0, 0, 0, 18, 192, 1, 0, 0, 0, 20, 196, 1, 0, 0, 0, 22, 200, 1, 0, 0, 0, 
	24, 213, 1, 0, 0, 0, 26, 226, 1, 0, 0, 0, 28, 230, 1, 0, 0, 0, 30, 240, 
	1, 0, 0, 0, 32, 246, 1, 0, 0, 0, 34, 340, 1, 0, 0, 0, 36, 390, 1, 0, 0, 
	0, 38, 43, 3, 2, 1, 0, 39, 43, 3, 4, 2, 0, 40, 43, 3, 6, 3, 0, 41, 43, 
	3, 8, 4, 0, 42, 38, 1, 0, 0, 0, 42, 39, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 
	42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 
	0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 0, 0, 1, 48, 
	1, 1, 0, 0, 0, 49, 50, 5, 1, 0, 0, 50, 51, 5, 59, 0, 0, 51, 52, 5, 2, 0, 
	0, 52, 53, 3, 14, 7, 0, 53, 54, 5, 3, 0, 0, 54, 3, 1, 0, 0, 0, 55, 56, 
	5, 4, 0, 0, 56, 59, 5, 59, 0, 0, 57, 58, 5, 5, 0, 0, 58, 60, 3, 16, 8, 
	0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 
	5, 2, 0, 0, 62, 63, 3, 34, 17, 0, 63, 64, 5, 3, 0, 0, 64, 5, 1, 0, 0, 0, 
	65, 66, 5, 6, 0, 0, 66, 67, 5, 59, 0, 0, 67, 69, 5, 7, 0, 0, 68, 70, 3, 
	10, 5, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 
	72, 5, 8, 0, 0, 72, 73, 5, 9, 0, 0, 73, 75, 5, 7, 0, 0, 74, 76, 3, 10, 
	5, 0, 75, 74, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 78, 
	5, 8, 0, 0, 78, 83, 5, 3, 0, 0, 79, 80, 5, 10, 0, 0, 80, 81, 3, 10, 5, 
	0, 81, 82, 5, 3, 0, 0, 82, 84, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0, 83, 84, 
	1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 94, 5, 11, 0, 0, 86, 93, 3, 30, 15, 
	0, 87, 93, 3, 20, 10, 0, 88, 93, 3, 28, 14, 0, 89, 93, 3, 26, 13, 0, 90, 
	93, 3, 22, 11, 0, 91, 93, 3, 24, 12, 0, 92, 86, 1, 0, 0, 0, 92, 87, 1, 
	0, 0, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 
	91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 
	0, 95, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 5, 12, 0, 0, 98, 100, 
	5, 3, 0, 0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 7, 1, 0, 0, 0, 
	101, 102, 5, 13, 0, 0, 102, 103, 3, 36, 18, 0, 103, 105, 5, 7, 0, 0, 104, 
	106, 3, 10, 5, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 
	1, 0, 0, 0, 107, 108, 5, 8, 0, 0, 108, 109, 5, 9, 0, 0, 109, 111, 5, 7, 
	0, 0, 110, 112, 3, 10, 5, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 
	112, 113, 1, 0, 0, 0, 113, 114, 5, 8, 0, 0, 114, 115, 5, 3, 0, 0, 115, 
	9, 1, 0, 0, 0, 116, 121, 3, 12, 6, 0, 117, 118, 5, 3, 0, 0, 118, 120, 3, 
	12, 6, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0, 0, 
	0, 121, 122, 1, 0, 0, 0, 122, 11, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124, 
	129, 3, 36, 18, 0, 125, 126, 5, 14, 0, 0, 126, 128, 3, 36, 18, 0, 127, 
	125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 129, 130, 
	1, 0, 0, 0, 130, 132, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 133, 5, 5, 
	0, 0, 133, 134, 3, 16, 8, 0, 134, 13, 1, 0, 0, 0, 135, 165, 3, 16, 8, 0, 
	136, 137, 5, 15, 0, 0, 137, 138, 5, 16, 0, 0, 138, 139, 5, 59, 0, 0, 139, 
	140, 5, 5, 0, 0, 140, 141, 3, 16, 8, 0, 141, 148, 1, 0, 0, 0, 142, 143, 
	5, 3, 0, 0, 143, 144, 5, 59, 0, 0, 144, 145, 5, 5, 0, 0, 145, 147, 3, 16, 
	8, 0, 146, 142, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 
	148, 149, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 
	152, 5, 17, 0, 0, 152, 165, 1, 0, 0, 0, 153, 154, 5, 18, 0, 0, 154, 155, 
	5, 16, 0, 0, 155, 160, 5, 59, 0, 0, 156, 157, 5, 14, 0, 0, 157, 159, 5, 
	59, 0, 0, 158, 156, 1, 0, 0, 0, 159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 
	0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 
	165, 5, 17, 0, 0, 164, 135, 1, 0, 0, 0, 164, 136, 1, 0, 0, 0, 164, 153, 
	1, 0, 0, 0, 165, 15, 1, 0, 0, 0, 166, 167, 6, 8, -1, 0, 167, 181, 5, 19, 
	0, 0, 168, 169, 5, 20, 0, 0, 169, 170, 5, 21, 0, 0, 170, 171, 3, 18, 9, 
	0, 171, 172, 5, 14, 0, 0, 172, 173, 3, 18, 9, 0, 173, 174, 5, 22, 0, 0, 
	174, 175, 5, 23, 0, 0, 175, 176, 5, 19, 0, 0, 176, 181, 1, 0, 0, 0, 177, 
	181, 5, 24, 0, 0, 178, 181, 5, 25, 0, 0, 179, 181, 5, 59, 0, 0, 180, 166, 
	1, 0, 0, 0, 180, 168, 1, 0, 0, 0, 180, 177, 1, 0, 0, 0, 180, 178, 1, 0, 
	0, 0, 180, 179, 1, 0, 0, 0, 181, 188, 1, 0, 0, 0, 182, 183, 10, 2, 0, 0, 
	183, 184, 5, 21, 0, 0, 184, 185, 5, 58, 0, 0, 185, 187, 5, 22, 0, 0, 186, 
	182, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 
	1, 0, 0, 0, 189, 17, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191, 193, 5, 26, 
	0, 0, 192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 
	194, 195, 5, 58, 0, 0, 195, 19, 1, 0, 0, 0, 196, 197, 5, 27, 0, 0, 197, 
	198, 3, 36, 18, 0, 198, 199, 5, 3, 0, 0, 199, 21, 1, 0, 0, 0, 200, 209, 
	5, 28, 0, 0, 201, 206, 5, 59, 0, 0, 202, 203, 5, 14, 0, 0, 203, 205, 5, 
	59, 0, 0, 204, 202, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206, 204, 1, 0, 0, 
	0, 206, 207, 1, 0, 0, 0, 207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 209, 
	201, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 
	5, 3, 0, 0, 212, 23, 1, 0, 0, 0, 213, 222, 5, 29, 0, 0, 214, 219, 3, 36, 
	18, 0, 215, 216, 5, 14, 0, 0, 216, 218, 3, 36, 18, 0, 217, 215, 1, 0, 0, 
	0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 
	223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 214, 1, 0, 0, 0, 222, 223, 
	1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224, 225, 5, 3, 0, 0, 225, 25, 1, 0, 
	0, 0, 226, 228, 5, 30, 0, 0, 227, 229, 5, 3, 0, 0, 228, 227, 1, 0, 0, 0, 
	228, 229, 1, 0, 0, 0, 229, 27, 1, 0, 0, 0, 230, 231, 5, 31, 0, 0, 231, 
	232, 3, 34, 17, 0, 232, 233, 5, 3, 0, 0, 233, 29, 1, 0, 0, 0, 234, 241, 
	3, 32, 16, 0, 235, 237, 5, 7, 0, 0, 236, 238, 3, 32, 16, 0, 237, 236, 1, 
	0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 5, 8, 0, 
	0, 240, 234, 1, 0, 0, 0, 240, 235, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 
	243, 5, 2, 0, 0, 243, 244, 3, 34, 17, 0, 244, 245, 5, 3, 0, 0, 245, 31, 
	1, 0, 0, 0, 246, 251, 3, 36, 18, 0, 247, 248, 5, 14, 0, 0, 248, 250, 3, 
	36, 18, 0, 249, 247, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 
	0, 0, 251, 252, 1, 0, 0, 0, 252, 33, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 
	254, 255, 6, 17, -1, 0, 255, 341, 5, 59, 0, 0, 256, 341, 5, 58, 0, 0, 257, 
	341, 5, 56, 0, 0, 258, 341, 5, 57, 0, 0, 259, 260, 7, 0, 0, 0, 260, 261, 
	5, 7, 0, 0, 261, 262, 3, 34, 17, 0, 262, 263, 5, 8, 0, 0, 263, 341, 1, 
	0, 0, 0, 264, 265, 3, 36, 18, 0, 265, 274, 5, 7, 0, 0, 266, 271, 3, 34, 
	17, 0, 267, 268, 5, 14, 0, 0, 268, 270, 3, 34, 17, 0, 269, 267, 1, 0, 0, 
	0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 
	275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 266, 1, 0, 0, 0, 274, 275, 
	1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 5, 8, 0, 0, 277, 341, 1, 0, 
	0, 0, 278, 279, 5, 33, 0, 0, 279, 280, 5, 7, 0, 0, 280, 283, 3, 34, 17, 
	0, 281, 282, 5, 14, 0, 0, 282, 284, 3, 34, 17, 0, 283, 281, 1, 0, 0, 0, 
	284, 285, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 
	287, 1, 0, 0, 0, 287, 288, 5, 8, 0, 0, 288, 341, 1, 0, 0, 0, 289, 290, 
	5, 36, 0, 0, 290, 341, 3, 34, 17, 14, 291, 292, 5, 37, 0, 0, 292, 341, 
	3, 34, 17, 13, 293, 294, 5, 26, 0, 0, 294, 341, 3, 34, 17, 12, 295, 296, 
	5, 53, 0, 0, 296, 297, 3, 34, 17, 0, 297, 298, 5, 54, 0, 0, 298, 299, 3, 
	34, 17, 0, 299, 300, 5, 55, 0, 0, 300, 301, 3, 34, 17, 4, 301, 341, 1, 
	0, 0, 0, 302, 303, 5, 59, 0, 0, 303, 304, 5, 16, 0, 0, 304, 305, 5, 59, 
	0, 0, 305, 306, 5, 2, 0, 0, 306, 313, 3, 34, 17, 0, 307, 308, 5, 3, 0, 
	0, 308, 309, 5, 59, 0, 0, 309, 310, 5, 2, 0, 0, 310, 312, 3, 34, 17, 0, 
	311, 307, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 
	314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 317, 
	5, 17, 0, 0, 317, 341, 1, 0, 0, 0, 318, 319, 5, 21, 0, 0, 319, 324, 3, 
	34, 17, 0, 320, 321, 5, 14, 0, 0, 321, 323, 3, 34, 17, 0, 322, 320, 1, 
	0, 0, 0, 323, 326, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 
	0, 325, 327, 1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 327, 328, 5, 22, 0, 0, 328, 
	341, 1, 0, 0, 0, 329, 330, 5, 7, 0, 0, 330, 335, 3, 34, 17, 0, 331, 332, 
	5, 14, 0, 0, 332, 334, 3, 34, 17, 0, 333, 331, 1, 0, 0, 0, 334, 337, 1, 
	0, 0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 338, 1, 0, 0, 
	0, 337, 335, 1, 0, 0, 0, 338, 339, 5, 8, 0, 0, 339, 341, 1, 0, 0, 0, 340, 
	254, 1, 0, 0, 0, 340, 256, 1, 0, 0, 0, 340, 257, 1, 0, 0, 0, 340, 258, 
	1, 0, 0, 0, 340, 259, 1, 0, 0, 0, 340, 264, 1, 0, 0, 0, 340, 278, 1, 0, 
	0, 0, 340, 289, 1, 0, 0, 0, 340, 291, 1, 0, 0, 0, 340, 293, 1, 0, 0, 0, 
	340, 295, 1, 0, 0, 0, 340, 302, 1, 0, 0, 0, 340, 318, 1, 0, 0, 0, 340, 
	329, 1, 0, 0, 0, 341, 387, 1, 0, 0, 0, 342, 343, 10, 11, 0, 0, 343, 344, 
	7, 1, 0, 0, 344, 386, 3, 34, 17, 12, 345, 346, 10, 10, 0, 0, 346, 347, 
	7, 2, 0, 0, 347, 386, 3, 34, 17, 11, 348, 349, 10, 9, 0, 0, 349, 350, 7, 
	3, 0, 0, 350, 386, 3, 34, 17, 10, 351, 352, 10, 8, 0, 0, 352, 353, 5, 48, 
	0, 0, 353, 386, 3, 34, 17, 9, 354, 355, 10, 7, 0, 0, 355, 356, 7, 4, 0, 
	0, 356, 386, 3, 34, 17, 8, 357, 358, 10, 6, 0, 0, 358, 359, 5, 51, 0, 0, 
	359, 386, 3, 34, 17, 6, 360, 361, 10, 5, 0, 0, 361, 362, 5, 52, 0, 0, 362, 
	386, 3, 34, 17, 5, 363, 364, 10, 18, 0, 0, 364, 365, 5, 34, 0, 0, 365, 
	386, 5, 59, 0, 0, 366, 367, 10, 17, 0, 0, 367, 368, 5, 16, 0, 0, 368, 369, 
	5, 59, 0, 0, 369, 370, 5, 35, 0, 0, 370, 371, 3, 34, 17, 0, 371, 372, 5, 
	17, 0, 0, 372, 386, 1, 0, 0, 0, 373, 374, 10, 16, 0, 0, 374, 375, 5, 21, 
	0, 0, 375, 376, 3, 34, 17, 0, 376, 377, 5, 22, 0, 0, 377, 386, 1, 0, 0, 
	0, 378, 379, 10, 15, 0, 0, 379, 380, 5, 21, 0, 0, 380, 381, 3, 34, 17, 
	0, 381, 382, 5, 35, 0, 0, 382, 383, 3, 34, 17, 0, 383, 384, 5, 22, 0, 0, 
	384, 386, 1, 0, 0, 0, 385, 342, 1, 0, 0, 0, 385, 345, 1, 0, 0, 0, 385, 
	348, 1, 0, 0, 0, 385, 351, 1, 0, 0, 0, 385, 354, 1, 0, 0, 0, 385, 357, 
	1, 0, 0, 0, 385, 360, 1, 0, 0, 0, 385, 363, 1, 0, 0, 0, 385, 366, 1, 0, 
	0, 0, 385, 373, 1, 0, 0, 0, 385, 378, 1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 
	387, 385, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 35, 1, 0, 0, 0, 389, 387, 
	1, 0, 0, 0, 390, 391, 6, 18, -1, 0, 391, 392, 5, 59, 0, 0, 392, 402, 1, 
	0, 0, 0, 393, 394, 10, 2, 0, 0, 394, 395, 5, 21, 0, 0, 395, 396, 5, 58, 
	0, 0, 396, 401, 5, 22, 0, 0, 397, 398, 10, 1, 0, 0, 398, 399, 5, 34, 0, 
	0, 399, 401, 5, 59, 0, 0, 400, 393, 1, 0, 0, 0, 400, 397, 1, 0, 0, 0, 401, 
	404, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 37, 1, 
	0, 0, 0, 404, 402, 1, 0, 0, 0, 38, 42, 44, 59, 69, 75, 83, 92, 94, 99, 
	105, 111, 121, 129, 148, 160, 164, 180, 188, 192, 206, 209, 219, 222, 228, 
	237, 240, 251, 271, 274, 285, 313, 324, 335, 340, 385, 387, 400, 402,
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
	LustreParserREAL = 56
	LustreParserBOOL = 57
	LustreParserINT = 58
	LustreParserID = 59
	LustreParserWS = 60
	LustreParserSL_COMMENT = 61
	LustreParserML_COMMENT = 62
	LustreParserERROR = 63
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
	LustreParserRULE_property = 10
	LustreParserRULE_realizabilityInputs = 11
	LustreParserRULE_ivc = 12
	LustreParserRULE_main = 13
	LustreParserRULE_assertion = 14
	LustreParserRULE_equation = 15
	LustreParserRULE_lhs = 16
	LustreParserRULE_expr = 17
	LustreParserRULE_eID = 18
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8274) != 0) {
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LustreParserT__0:
			{
				p.SetState(38)
				p.Typedef()
			}


		case LustreParserT__3:
			{
				p.SetState(39)
				p.Constant()
			}


		case LustreParserT__5:
			{
				p.SetState(40)
				p.Node()
			}


		case LustreParserT__12:
			{
				p.SetState(41)
				p.Function()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
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
		p.SetState(49)
		p.Match(LustreParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(52)
		p.TopLevelType()
	}
	{
		p.SetState(53)
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
		p.SetState(55)
		p.Match(LustreParserT__3)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__4 {
		{
			p.SetState(57)
			p.Match(LustreParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(58)
			p.type_(0)
		}

	}
	{
		p.SetState(61)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(62)
		p.expr(0)
	}
	{
		p.SetState(63)
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
		p.SetState(65)
		p.Match(LustreParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(68)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).input = _x
		}

	}
	{
		p.SetState(71)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(LustreParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(74)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).output = _x
		}

	}
	{
		p.SetState(77)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Match(LustreParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__9 {
		{
			p.SetState(79)
			p.Match(LustreParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(80)

			var _x = p.VarDeclList()


			localctx.(*NodeContext).local = _x
		}
		{
			p.SetState(81)
			p.Match(LustreParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(85)
		p.Match(LustreParserT__10)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 576460756464173184) != 0) {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case LustreParserT__6, LustreParserID:
			{
				p.SetState(86)
				p.Equation()
			}


		case LustreParserT__26:
			{
				p.SetState(87)
				p.Property()
			}


		case LustreParserT__30:
			{
				p.SetState(88)
				p.Assertion()
			}


		case LustreParserT__29:
			{
				p.SetState(89)
				p.Main()
			}


		case LustreParserT__27:
			{
				p.SetState(90)
				p.RealizabilityInputs()
			}


		case LustreParserT__28:
			{
				p.SetState(91)
				p.Ivc()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		p.Match(LustreParserT__11)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__2 {
		{
			p.SetState(98)
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
		p.SetState(101)
		p.Match(LustreParserT__12)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(102)
		p.eID(0)
	}
	{
		p.SetState(103)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(104)

			var _x = p.VarDeclList()


			localctx.(*FunctionContext).input = _x
		}

	}
	{
		p.SetState(107)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(LustreParserT__8)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(LustreParserT__6)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(110)

			var _x = p.VarDeclList()


			localctx.(*FunctionContext).output = _x
		}

	}
	{
		p.SetState(113)
		p.Match(LustreParserT__7)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(114)
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
		p.SetState(116)
		p.VarDeclGroup()
	}
	p.SetState(121)
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
				p.SetState(117)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(118)
				p.VarDeclGroup()
			}


		}
		p.SetState(123)
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
		p.SetState(124)
		p.eID(0)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == LustreParserT__13 {
		{
			p.SetState(125)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(126)
			p.eID(0)
		}


		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(132)
		p.Match(LustreParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(133)
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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserT__18, LustreParserT__19, LustreParserT__23, LustreParserT__24, LustreParserID:
		localctx = NewPlainTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.type_(0)
		}


	case LustreParserT__14:
		localctx = NewRecordTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Match(LustreParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

		{
			p.SetState(138)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(LustreParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(140)
			p.type_(0)
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__2 {
			{
				p.SetState(142)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(144)
				p.Match(LustreParserT__4)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(145)
				p.type_(0)
			}


			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(151)
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
			p.SetState(153)
			p.Match(LustreParserT__17)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(156)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(157)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
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
	p.SetState(180)
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
			p.SetState(167)
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
			p.SetState(168)
			p.Match(LustreParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(LustreParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Bound()
		}
		{
			p.SetState(171)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Bound()
		}
		{
			p.SetState(173)
			p.Match(LustreParserT__21)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(LustreParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(175)
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
			p.SetState(177)
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
			p.SetState(178)
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
			p.SetState(179)
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
	p.SetState(188)
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
			p.SetState(182)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(183)
				p.Match(LustreParserT__20)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(184)
				p.Match(LustreParserINT)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(185)
				p.Match(LustreParserT__21)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


		}
		p.SetState(190)
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
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__25 {
		{
			p.SetState(191)
			p.Match(LustreParserT__25)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(194)
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
	p.EnterRule(localctx, 20, LustreParserRULE_property)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(LustreParserT__26)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(197)
		p.eID(0)
	}
	{
		p.SetState(198)
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
	p.EnterRule(localctx, 22, LustreParserRULE_realizabilityInputs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Match(LustreParserT__27)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(201)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(202)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(203)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(211)
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
	p.EnterRule(localctx, 24, LustreParserRULE_ivc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(LustreParserT__28)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserID {
		{
			p.SetState(214)
			p.eID(0)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(215)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(216)
				p.eID(0)
			}


			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(224)
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
	p.EnterRule(localctx, 26, LustreParserRULE_main)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(LustreParserT__29)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == LustreParserT__2 {
		{
			p.SetState(227)
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
	p.EnterRule(localctx, 28, LustreParserRULE_assertion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(LustreParserT__30)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(231)
		p.expr(0)
	}
	{
		p.SetState(232)
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
	Expr() IExprContext
	Lhs() ILhsContext

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

func (s *EquationContext) Expr() IExprContext {
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

func (s *EquationContext) Lhs() ILhsContext {
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
	p.EnterRule(localctx, 30, LustreParserRULE_equation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LustreParserID:
		{
			p.SetState(234)
			p.Lhs()
		}


	case LustreParserT__6:
		{
			p.SetState(235)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == LustreParserID {
			{
				p.SetState(236)
				p.Lhs()
			}

		}
		{
			p.SetState(239)
			p.Match(LustreParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(242)
		p.Match(LustreParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(243)
		p.expr(0)
	}
	{
		p.SetState(244)
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


// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEID() []IEIDContext
	EID(i int) IEIDContext

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

func (s *LhsContext) AllEID() []IEIDContext {
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

func (s *LhsContext) EID(i int) IEIDContext {
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
	p.EnterRule(localctx, 32, LustreParserRULE_lhs)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.eID(0)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == LustreParserT__13 {
		{
			p.SetState(247)
			p.Match(LustreParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(248)
			p.eID(0)
		}


		p.SetState(253)
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, LustreParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(255)
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
			p.SetState(256)
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
			p.SetState(257)
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
			p.SetState(258)
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
			p.SetState(259)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CastExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == LustreParserT__24 || _la == LustreParserT__31) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CastExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(260)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(261)
			p.expr(0)
		}
		{
			p.SetState(262)
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
			p.SetState(264)
			p.eID(0)
		}
		{
			p.SetState(265)
			p.Match(LustreParserT__6)
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


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1089871328969752704) != 0) {
			{
				p.SetState(266)
				p.expr(0)
			}
			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == LustreParserT__13 {
				{
					p.SetState(267)
					p.Match(LustreParserT__13)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(268)
					p.expr(0)
				}


				p.SetState(273)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(276)
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
			p.SetState(278)
			p.Match(LustreParserT__32)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(280)
			p.expr(0)
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == LustreParserT__13 {
			{
				p.SetState(281)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(282)
				p.expr(0)
			}


			p.SetState(285)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(287)
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
			p.SetState(289)
			p.Match(LustreParserT__35)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(290)
			p.expr(14)
		}


	case 9:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(291)
			p.Match(LustreParserT__36)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(292)
			p.expr(13)
		}


	case 10:
		localctx = NewNegateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(293)
			p.Match(LustreParserT__25)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(294)
			p.expr(12)
		}


	case 11:
		localctx = NewIfThenElseExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(295)
			p.Match(LustreParserT__52)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(296)
			p.expr(0)
		}
		{
			p.SetState(297)
			p.Match(LustreParserT__53)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(298)
			p.expr(0)
		}
		{
			p.SetState(299)
			p.Match(LustreParserT__54)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(300)
			p.expr(4)
		}


	case 12:
		localctx = NewRecordExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(302)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(LustreParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(304)
			p.Match(LustreParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(LustreParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(306)
			p.expr(0)
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__2 {
			{
				p.SetState(307)
				p.Match(LustreParserT__2)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(308)
				p.Match(LustreParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(309)
				p.Match(LustreParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(310)
				p.expr(0)
			}


			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(316)
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
			p.SetState(318)
			p.Match(LustreParserT__20)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(319)
			p.expr(0)
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(320)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(321)
				p.expr(0)
			}


			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(327)
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
			p.SetState(329)
			p.Match(LustreParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(330)
			p.expr(0)
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == LustreParserT__13 {
			{
				p.SetState(331)
				p.Match(LustreParserT__13)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(332)
				p.expr(0)
			}


			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(338)
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
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(385)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(343)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4123168604160) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(344)
					p.expr(12)
				}


			case 2:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(346)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LustreParserT__25 || _la == LustreParserT__41) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(347)
					p.expr(11)
				}


			case 3:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(348)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(349)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 272678883688452) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(350)
					p.expr(10)
				}


			case 4:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(352)

					var _m = p.Match(LustreParserT__47)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(353)
					p.expr(9)
				}


			case 5:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(355)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LustreParserT__48 || _la == LustreParserT__49) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(356)
					p.expr(8)
				}


			case 6:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(357)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(358)

					var _m = p.Match(LustreParserT__50)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(359)
					p.expr(6)
				}


			case 7:
				localctx = NewBinaryExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(361)

					var _m = p.Match(LustreParserT__51)

					localctx.(*BinaryExprContext).op = _m
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(362)
					p.expr(5)
				}


			case 8:
				localctx = NewRecordAccessExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(363)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(364)
					p.Match(LustreParserT__33)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(365)
					p.Match(LustreParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 9:
				localctx = NewRecordUpdateExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(366)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(367)
					p.Match(LustreParserT__15)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(368)
					p.Match(LustreParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(369)
					p.Match(LustreParserT__34)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(370)
					p.expr(0)
				}
				{
					p.SetState(371)
					p.Match(LustreParserT__16)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 10:
				localctx = NewArrayAccessExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(373)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(374)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(375)
					p.expr(0)
				}
				{
					p.SetState(376)
					p.Match(LustreParserT__21)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 11:
				localctx = NewArrayUpdateExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_expr)
				p.SetState(378)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(379)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(380)
					p.expr(0)
				}
				{
					p.SetState(381)
					p.Match(LustreParserT__34)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(382)
					p.expr(0)
				}
				{
					p.SetState(383)
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
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, LustreParserRULE_eID, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewBaseEIDContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(391)
		p.Match(LustreParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewArrayEIDContext(p, NewEIDContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_eID)
				p.SetState(393)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(394)
					p.Match(LustreParserT__20)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(395)
					p.Match(LustreParserINT)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(396)
					p.Match(LustreParserT__21)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 2:
				localctx = NewRecordEIDContext(p, NewEIDContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LustreParserRULE_eID)
				p.SetState(397)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(398)
					p.Match(LustreParserT__33)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(399)
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
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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

	case 17:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)

	case 18:
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


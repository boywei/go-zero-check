// Code generated from /Users/wei/GoProjects/go-zero-check/internal/synlong/Synlong.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Synlong

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


type SynlongParser struct {
	*antlr.BaseParser
}

var SynlongParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func synlongParserInit() {
  staticData := &SynlongParserStaticData
  staticData.LiteralNames = []string{
    "", "'type'", "';'", "'='", "'enum'", "'{'", "'}'", "'bool'", "'char'", 
    "'byte'", "'ubyte'", "'short'", "'ushort'", "'int'", "'uint'", "'long'", 
    "'ulong'", "'float'", "'real'", "','", "'^'", "':'", "'const'", "'['", 
    "']'", "'function'", "'node'", "'('", "')'", "'returns'", "'let'", "'tel'", 
    "'var'", "'when'", "'not'", "'last'", "'automaton'", "'initial'", "'final'", 
    "'state'", "'unless'", "'until'", "'if'", "'resume'", "'restart'", "'''", 
    "'.'", "'pre'", "'->'", "'fby'", "'merge'", "'#'", "'..'", "'default'", 
    "'with'", "'then'", "'else'", "'case'", "'of'", "'|'", "'-'", "'true'", 
    "'false'", "'_'", "'<<'", "'>>'", "'mapw'", "'mapwi'", "'foldw'", "'foldwi'", 
    "'(make'", "'(flatten'", "'+$'", "'-$'", "'not$'", "'short$'", "'int$'", 
    "'float$'", "'real$'", "'$+$'", "'$-$'", "'$*$'", "'$/$'", "'$mod$'", 
    "'$div$'", "'$=$'", "'$<>$'", "'$<$'", "'$>$'", "'$<=$'", "'$>=$'", 
    "'$and$'", "'$or$'", "'$xor$'", "'map'", "'fold'", "'mapi'", "'foldi'", 
    "'mapfold'", "'+'", "'*'", "'/'", "'mod'", "'div'", "'<>'", "'<'", "'>'", 
    "'<='", "'>='", "'and'", "'or'", "'xor'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "", "", "", "", "", "", "", "", "", "COMMENT", "ID", "CHAR", "INTEGER", 
    "UINT", "FLOAT", "REAL", "USHORT", "SHORT", "WS",
  }
  staticData.RuleNames = []string{
    "program", "decls", "type_block", "type_decl", "type_def", "type_expr", 
    "field_decl", "typevar", "const_block", "const_decl", "const_expr", 
    "const_list", "const_label_expr", "user_op_decl", "op_kind", "params", 
    "returns_clause", "op_body", "local_block", "var_decls", "var_id", "when_decl", 
    "clock_expr", "last_decl", "equation", "lhs", "lhs_id", "return_statement", 
    "returns_var", "state_machine", "state_decl", "data_def", "transition", 
    "expr", "simple_expr", "tempo_expr", "bool_expr", "array_expr", "struct_expr", 
    "label_expr", "mixed_constructor", "label_or_index", "index", "switch_expr", 
    "case_expr", "pattern", "apply_expr", "prefix_operator", "prefix_unary_operator", 
    "prefix_binary_operator", "iterator", "list", "unary_arith_op", "bin_arith_op", 
    "bin_relation_op", "bin_bool_op", "atom",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 121, 744, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7, 
	31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36, 
	2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2, 
	42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47, 
	7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7, 
	52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 1, 0, 5, 0, 
	116, 8, 0, 10, 0, 12, 0, 119, 9, 0, 1, 1, 1, 1, 1, 1, 3, 1, 124, 8, 1, 
	1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 130, 8, 2, 10, 2, 12, 2, 133, 9, 2, 1, 3, 
	1, 3, 1, 3, 3, 3, 138, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 145, 8, 
	4, 10, 4, 12, 4, 148, 9, 4, 1, 4, 3, 4, 151, 8, 4, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 
	5, 1, 5, 1, 5, 5, 5, 171, 8, 5, 10, 5, 12, 5, 174, 9, 5, 1, 5, 1, 5, 3, 
	5, 178, 8, 5, 1, 5, 1, 5, 1, 5, 5, 5, 183, 8, 5, 10, 5, 12, 5, 186, 9, 
	5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 198, 
	8, 8, 10, 8, 12, 8, 201, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 208, 
	8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 224, 8, 10, 10, 10, 12, 10, 227, 
	9, 10, 1, 10, 1, 10, 3, 10, 231, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 
	10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 245, 8, 10, 
	10, 10, 12, 10, 248, 9, 10, 1, 11, 1, 11, 1, 11, 5, 11, 253, 8, 11, 10, 
	11, 12, 11, 256, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 
	1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 274, 
	8, 15, 10, 15, 12, 15, 277, 9, 15, 3, 15, 279, 8, 15, 1, 15, 1, 15, 1, 
	16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 288, 8, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 5, 17, 294, 8, 17, 10, 17, 12, 17, 297, 9, 17, 1, 17, 1, 17, 3, 
	17, 301, 8, 17, 3, 17, 303, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 309, 
	8, 18, 10, 18, 12, 18, 312, 9, 18, 1, 19, 1, 19, 5, 19, 316, 8, 19, 10, 
	19, 12, 19, 319, 9, 19, 1, 19, 1, 19, 1, 19, 3, 19, 324, 8, 19, 1, 19, 
	3, 19, 327, 8, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 
	22, 3, 22, 337, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 
	1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 350, 8, 24, 1, 25, 1, 25, 1, 25, 1, 
	25, 5, 25, 356, 8, 25, 10, 25, 12, 25, 359, 9, 25, 3, 25, 361, 8, 25, 1, 
	26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 4, 28, 369, 8, 28, 11, 28, 12, 28, 
	370, 1, 29, 1, 29, 3, 29, 375, 8, 29, 1, 29, 4, 29, 378, 8, 29, 11, 29, 
	12, 29, 379, 1, 30, 3, 30, 383, 8, 30, 1, 30, 3, 30, 386, 8, 30, 1, 30, 
	1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 394, 8, 30, 10, 30, 12, 30, 397, 
	9, 30, 3, 30, 399, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 406, 
	8, 30, 10, 30, 12, 30, 409, 9, 30, 3, 30, 411, 8, 30, 1, 31, 1, 31, 1, 
	31, 1, 31, 3, 31, 417, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 423, 8, 
	31, 10, 31, 12, 31, 426, 9, 31, 1, 31, 3, 31, 429, 8, 31, 1, 32, 1, 32, 
	1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 
	33, 1, 33, 1, 33, 1, 33, 3, 33, 447, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 
	1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 460, 8, 34, 1, 
	34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 
	1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 482, 
	8, 34, 10, 34, 12, 34, 485, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 
	35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 
	1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 
	35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 518, 8, 35, 1, 36, 1, 36, 1, 36, 
	1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 
	37, 1, 37, 1, 37, 4, 37, 536, 8, 37, 11, 37, 12, 37, 537, 1, 37, 1, 37, 
	1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 
	37, 552, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 557, 8, 38, 10, 38, 12, 38, 
	560, 9, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 
	40, 1, 40, 4, 40, 572, 8, 40, 11, 40, 12, 40, 573, 1, 40, 1, 40, 1, 40, 
	1, 40, 1, 41, 1, 41, 1, 41, 3, 41, 583, 8, 41, 1, 42, 1, 42, 1, 42, 1, 
	42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 
	1, 43, 1, 43, 4, 43, 601, 8, 43, 11, 43, 12, 43, 602, 1, 43, 1, 43, 3, 
	43, 607, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 
	3, 45, 617, 8, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 623, 8, 45, 1, 46, 
	1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 
	46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 
	1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 
	46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 
	1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 
	46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 
	1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 
	46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 706, 8, 46, 1, 47, 
	1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 717, 8, 
	47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 5, 51, 727, 
	8, 51, 10, 51, 12, 51, 730, 9, 51, 3, 51, 732, 8, 51, 1, 52, 1, 52, 1, 
	53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 0, 3, 10, 20, 
	68, 57, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 
	70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 
	106, 108, 110, 112, 0, 10, 1, 0, 25, 26, 1, 0, 43, 44, 1, 0, 72, 78, 1, 
	0, 79, 93, 1, 0, 94, 98, 3, 0, 34, 34, 60, 60, 99, 99, 2, 0, 60, 60, 99, 
	103, 2, 0, 3, 3, 104, 108, 1, 0, 109, 111, 2, 0, 61, 62, 114, 120, 792, 
	0, 117, 1, 0, 0, 0, 2, 123, 1, 0, 0, 0, 4, 125, 1, 0, 0, 0, 6, 134, 1, 
	0, 0, 0, 8, 150, 1, 0, 0, 0, 10, 177, 1, 0, 0, 0, 12, 187, 1, 0, 0, 0, 
	14, 191, 1, 0, 0, 0, 16, 193, 1, 0, 0, 0, 18, 202, 1, 0, 0, 0, 20, 230, 
	1, 0, 0, 0, 22, 249, 1, 0, 0, 0, 24, 257, 1, 0, 0, 0, 26, 261, 1, 0, 0, 
	0, 28, 267, 1, 0, 0, 0, 30, 269, 1, 0, 0, 0, 32, 282, 1, 0, 0, 0, 34, 302, 
	1, 0, 0, 0, 36, 304, 1, 0, 0, 0, 38, 313, 1, 0, 0, 0, 40, 328, 1, 0, 0, 
	0, 42, 330, 1, 0, 0, 0, 44, 336, 1, 0, 0, 0, 46, 338, 1, 0, 0, 0, 48, 349, 
	1, 0, 0, 0, 50, 360, 1, 0, 0, 0, 52, 362, 1, 0, 0, 0, 54, 364, 1, 0, 0, 
	0, 56, 368, 1, 0, 0, 0, 58, 372, 1, 0, 0, 0, 60, 382, 1, 0, 0, 0, 62, 428, 
	1, 0, 0, 0, 64, 430, 1, 0, 0, 0, 66, 446, 1, 0, 0, 0, 68, 459, 1, 0, 0, 
	0, 70, 517, 1, 0, 0, 0, 72, 519, 1, 0, 0, 0, 74, 551, 1, 0, 0, 0, 76, 553, 
	1, 0, 0, 0, 78, 563, 1, 0, 0, 0, 80, 567, 1, 0, 0, 0, 82, 582, 1, 0, 0, 
	0, 84, 584, 1, 0, 0, 0, 86, 606, 1, 0, 0, 0, 88, 608, 1, 0, 0, 0, 90, 622, 
	1, 0, 0, 0, 92, 705, 1, 0, 0, 0, 94, 716, 1, 0, 0, 0, 96, 718, 1, 0, 0, 
	0, 98, 720, 1, 0, 0, 0, 100, 722, 1, 0, 0, 0, 102, 731, 1, 0, 0, 0, 104, 
	733, 1, 0, 0, 0, 106, 735, 1, 0, 0, 0, 108, 737, 1, 0, 0, 0, 110, 739, 
	1, 0, 0, 0, 112, 741, 1, 0, 0, 0, 114, 116, 3, 2, 1, 0, 115, 114, 1, 0, 
	0, 0, 116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 
	118, 1, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 124, 3, 4, 2, 0, 121, 124, 
	3, 16, 8, 0, 122, 124, 3, 26, 13, 0, 123, 120, 1, 0, 0, 0, 123, 121, 1, 
	0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 3, 1, 0, 0, 0, 125, 131, 5, 1, 0, 0, 
	126, 127, 3, 6, 3, 0, 127, 128, 5, 2, 0, 0, 128, 130, 1, 0, 0, 0, 129, 
	126, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 
	1, 0, 0, 0, 132, 5, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 137, 5, 113, 
	0, 0, 135, 136, 5, 3, 0, 0, 136, 138, 3, 8, 4, 0, 137, 135, 1, 0, 0, 0, 
	137, 138, 1, 0, 0, 0, 138, 7, 1, 0, 0, 0, 139, 151, 3, 10, 5, 0, 140, 141, 
	5, 4, 0, 0, 141, 142, 5, 5, 0, 0, 142, 146, 5, 113, 0, 0, 143, 145, 5, 
	113, 0, 0, 144, 143, 1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 
	0, 0, 146, 147, 1, 0, 0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 
	149, 151, 5, 6, 0, 0, 150, 139, 1, 0, 0, 0, 150, 140, 1, 0, 0, 0, 151, 
	9, 1, 0, 0, 0, 152, 153, 6, 5, -1, 0, 153, 178, 5, 7, 0, 0, 154, 178, 5, 
	8, 0, 0, 155, 178, 5, 9, 0, 0, 156, 178, 5, 10, 0, 0, 157, 178, 5, 11, 
	0, 0, 158, 178, 5, 12, 0, 0, 159, 178, 5, 13, 0, 0, 160, 178, 5, 14, 0, 
	0, 161, 178, 5, 15, 0, 0, 162, 178, 5, 16, 0, 0, 163, 178, 5, 17, 0, 0, 
	164, 178, 5, 18, 0, 0, 165, 178, 3, 14, 7, 0, 166, 167, 5, 5, 0, 0, 167, 
	172, 3, 12, 6, 0, 168, 169, 5, 19, 0, 0, 169, 171, 3, 12, 6, 0, 170, 168, 
	1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 
	0, 0, 173, 175, 1, 0, 0, 0, 174, 172, 1, 0, 0, 0, 175, 176, 5, 6, 0, 0, 
	176, 178, 1, 0, 0, 0, 177, 152, 1, 0, 0, 0, 177, 154, 1, 0, 0, 0, 177, 
	155, 1, 0, 0, 0, 177, 156, 1, 0, 0, 0, 177, 157, 1, 0, 0, 0, 177, 158, 
	1, 0, 0, 0, 177, 159, 1, 0, 0, 0, 177, 160, 1, 0, 0, 0, 177, 161, 1, 0, 
	0, 0, 177, 162, 1, 0, 0, 0, 177, 163, 1, 0, 0, 0, 177, 164, 1, 0, 0, 0, 
	177, 165, 1, 0, 0, 0, 177, 166, 1, 0, 0, 0, 178, 184, 1, 0, 0, 0, 179, 
	180, 10, 1, 0, 0, 180, 181, 5, 20, 0, 0, 181, 183, 3, 20, 10, 0, 182, 179, 
	1, 0, 0, 0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 
	0, 0, 185, 11, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 188, 5, 113, 0, 0, 
	188, 189, 5, 21, 0, 0, 189, 190, 3, 10, 5, 0, 190, 13, 1, 0, 0, 0, 191, 
	192, 5, 113, 0, 0, 192, 15, 1, 0, 0, 0, 193, 199, 5, 22, 0, 0, 194, 195, 
	3, 18, 9, 0, 195, 196, 5, 2, 0, 0, 196, 198, 1, 0, 0, 0, 197, 194, 1, 0, 
	0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 
	200, 17, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 5, 113, 0, 0, 203, 
	204, 5, 21, 0, 0, 204, 207, 3, 10, 5, 0, 205, 206, 5, 3, 0, 0, 206, 208, 
	3, 20, 10, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 19, 1, 0, 
	0, 0, 209, 210, 6, 10, -1, 0, 210, 231, 5, 113, 0, 0, 211, 231, 3, 112, 
	56, 0, 212, 213, 3, 104, 52, 0, 213, 214, 3, 20, 10, 6, 214, 231, 1, 0, 
	0, 0, 215, 216, 5, 23, 0, 0, 216, 217, 3, 22, 11, 0, 217, 218, 5, 24, 0, 
	0, 218, 231, 1, 0, 0, 0, 219, 220, 5, 5, 0, 0, 220, 225, 3, 24, 12, 0, 
	221, 222, 5, 19, 0, 0, 222, 224, 3, 24, 12, 0, 223, 221, 1, 0, 0, 0, 224, 
	227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 228, 
	1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 229, 5, 6, 0, 0, 229, 231, 1, 0, 
	0, 0, 230, 209, 1, 0, 0, 0, 230, 211, 1, 0, 0, 0, 230, 212, 1, 0, 0, 0, 
	230, 215, 1, 0, 0, 0, 230, 219, 1, 0, 0, 0, 231, 246, 1, 0, 0, 0, 232, 
	233, 10, 5, 0, 0, 233, 234, 3, 106, 53, 0, 234, 235, 3, 20, 10, 6, 235, 
	245, 1, 0, 0, 0, 236, 237, 10, 4, 0, 0, 237, 238, 3, 110, 55, 0, 238, 239, 
	3, 20, 10, 5, 239, 245, 1, 0, 0, 0, 240, 241, 10, 3, 0, 0, 241, 242, 3, 
	108, 54, 0, 242, 243, 3, 20, 10, 4, 243, 245, 1, 0, 0, 0, 244, 232, 1, 
	0, 0, 0, 244, 236, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 245, 248, 1, 0, 0, 
	0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 21, 1, 0, 0, 0, 248, 
	246, 1, 0, 0, 0, 249, 254, 3, 20, 10, 0, 250, 251, 5, 19, 0, 0, 251, 253, 
	3, 20, 10, 0, 252, 250, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 
	0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 23, 1, 0, 0, 0, 256, 254, 1, 0, 0, 
	0, 257, 258, 5, 113, 0, 0, 258, 259, 5, 21, 0, 0, 259, 260, 3, 20, 10, 
	0, 260, 25, 1, 0, 0, 0, 261, 262, 3, 28, 14, 0, 262, 263, 5, 113, 0, 0, 
	263, 264, 3, 30, 15, 0, 264, 265, 3, 32, 16, 0, 265, 266, 3, 34, 17, 0, 
	266, 27, 1, 0, 0, 0, 267, 268, 7, 0, 0, 0, 268, 29, 1, 0, 0, 0, 269, 278, 
	5, 27, 0, 0, 270, 275, 3, 38, 19, 0, 271, 272, 5, 2, 0, 0, 272, 274, 3, 
	38, 19, 0, 273, 271, 1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 
	0, 0, 275, 276, 1, 0, 0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 
	278, 270, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 
	281, 5, 28, 0, 0, 281, 31, 1, 0, 0, 0, 282, 283, 5, 29, 0, 0, 283, 284, 
	3, 30, 15, 0, 284, 33, 1, 0, 0, 0, 285, 303, 5, 2, 0, 0, 286, 288, 3, 36, 
	18, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 
	289, 295, 5, 30, 0, 0, 290, 291, 3, 48, 24, 0, 291, 292, 5, 2, 0, 0, 292, 
	294, 1, 0, 0, 0, 293, 290, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 
	1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 298, 1, 0, 0, 0, 297, 295, 1, 0, 
	0, 0, 298, 300, 5, 31, 0, 0, 299, 301, 5, 2, 0, 0, 300, 299, 1, 0, 0, 0, 
	300, 301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 285, 1, 0, 0, 0, 302, 
	287, 1, 0, 0, 0, 303, 35, 1, 0, 0, 0, 304, 310, 5, 32, 0, 0, 305, 306, 
	3, 38, 19, 0, 306, 307, 5, 2, 0, 0, 307, 309, 1, 0, 0, 0, 308, 305, 1, 
	0, 0, 0, 309, 312, 1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 
	0, 311, 37, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 317, 3, 40, 20, 0, 314, 
	316, 3, 40, 20, 0, 315, 314, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 
	1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319, 317, 1, 0, 
	0, 0, 320, 321, 5, 21, 0, 0, 321, 323, 3, 10, 5, 0, 322, 324, 3, 42, 21, 
	0, 323, 322, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 326, 1, 0, 0, 0, 325, 
	327, 3, 46, 23, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 39, 
	1, 0, 0, 0, 328, 329, 5, 113, 0, 0, 329, 41, 1, 0, 0, 0, 330, 331, 5, 33, 
	0, 0, 331, 332, 3, 44, 22, 0, 332, 43, 1, 0, 0, 0, 333, 337, 5, 113, 0, 
	0, 334, 335, 5, 34, 0, 0, 335, 337, 5, 113, 0, 0, 336, 333, 1, 0, 0, 0, 
	336, 334, 1, 0, 0, 0, 337, 45, 1, 0, 0, 0, 338, 339, 5, 35, 0, 0, 339, 
	340, 5, 3, 0, 0, 340, 341, 3, 20, 10, 0, 341, 47, 1, 0, 0, 0, 342, 343, 
	3, 50, 25, 0, 343, 344, 5, 3, 0, 0, 344, 345, 3, 66, 33, 0, 345, 350, 1, 
	0, 0, 0, 346, 347, 3, 58, 29, 0, 347, 348, 3, 54, 27, 0, 348, 350, 1, 0, 
	0, 0, 349, 342, 1, 0, 0, 0, 349, 346, 1, 0, 0, 0, 350, 49, 1, 0, 0, 0, 
	351, 352, 5, 27, 0, 0, 352, 361, 5, 28, 0, 0, 353, 357, 3, 52, 26, 0, 354, 
	356, 3, 52, 26, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357, 355, 
	1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 
	0, 0, 360, 351, 1, 0, 0, 0, 360, 353, 1, 0, 0, 0, 361, 51, 1, 0, 0, 0, 
	362, 363, 5, 113, 0, 0, 363, 53, 1, 0, 0, 0, 364, 365, 5, 29, 0, 0, 365, 
	366, 3, 56, 28, 0, 366, 55, 1, 0, 0, 0, 367, 369, 5, 113, 0, 0, 368, 367, 
	1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 
	0, 0, 371, 57, 1, 0, 0, 0, 372, 374, 5, 36, 0, 0, 373, 375, 5, 113, 0, 
	0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 377, 1, 0, 0, 0, 376, 
	378, 3, 60, 30, 0, 377, 376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 377, 
	1, 0, 0, 0, 379, 380, 1, 0, 0, 0, 380, 59, 1, 0, 0, 0, 381, 383, 5, 37, 
	0, 0, 382, 381, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 385, 1, 0, 0, 0, 
	384, 386, 5, 38, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 
	387, 1, 0, 0, 0, 387, 388, 5, 39, 0, 0, 388, 398, 5, 113, 0, 0, 389, 390, 
	5, 40, 0, 0, 390, 395, 3, 64, 32, 0, 391, 392, 5, 2, 0, 0, 392, 394, 3, 
	64, 32, 0, 393, 391, 1, 0, 0, 0, 394, 397, 1, 0, 0, 0, 395, 393, 1, 0, 
	0, 0, 395, 396, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 
	398, 389, 1, 0, 0, 0, 398, 399, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 
	410, 3, 62, 31, 0, 401, 402, 5, 41, 0, 0, 402, 407, 3, 64, 32, 0, 403, 
	404, 5, 2, 0, 0, 404, 406, 3, 64, 32, 0, 405, 403, 1, 0, 0, 0, 406, 409, 
	1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 411, 1, 0, 
	0, 0, 409, 407, 1, 0, 0, 0, 410, 401, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 
	411, 61, 1, 0, 0, 0, 412, 413, 3, 48, 24, 0, 413, 414, 5, 2, 0, 0, 414, 
	429, 1, 0, 0, 0, 415, 417, 3, 36, 18, 0, 416, 415, 1, 0, 0, 0, 416, 417, 
	1, 0, 0, 0, 417, 418, 1, 0, 0, 0, 418, 424, 5, 30, 0, 0, 419, 420, 3, 48, 
	24, 0, 420, 421, 5, 2, 0, 0, 421, 423, 1, 0, 0, 0, 422, 419, 1, 0, 0, 0, 
	423, 426, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 
	427, 1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 427, 429, 5, 31, 0, 0, 428, 412, 
	1, 0, 0, 0, 428, 416, 1, 0, 0, 0, 429, 63, 1, 0, 0, 0, 430, 431, 5, 42, 
	0, 0, 431, 432, 3, 66, 33, 0, 432, 433, 7, 1, 0, 0, 433, 434, 5, 113, 0, 
	0, 434, 65, 1, 0, 0, 0, 435, 447, 3, 68, 34, 0, 436, 437, 5, 35, 0, 0, 
	437, 438, 5, 45, 0, 0, 438, 447, 5, 113, 0, 0, 439, 447, 3, 70, 35, 0, 
	440, 447, 3, 72, 36, 0, 441, 447, 3, 74, 37, 0, 442, 447, 3, 76, 38, 0, 
	443, 447, 3, 80, 40, 0, 444, 447, 3, 86, 43, 0, 445, 447, 3, 92, 46, 0, 
	446, 435, 1, 0, 0, 0, 446, 436, 1, 0, 0, 0, 446, 439, 1, 0, 0, 0, 446, 
	440, 1, 0, 0, 0, 446, 441, 1, 0, 0, 0, 446, 442, 1, 0, 0, 0, 446, 443, 
	1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 446, 445, 1, 0, 0, 0, 447, 67, 1, 0, 
	0, 0, 448, 449, 6, 34, -1, 0, 449, 460, 5, 113, 0, 0, 450, 460, 3, 112, 
	56, 0, 451, 452, 3, 104, 52, 0, 452, 453, 3, 68, 34, 5, 453, 460, 1, 0, 
	0, 0, 454, 455, 5, 27, 0, 0, 455, 456, 3, 10, 5, 0, 456, 457, 3, 68, 34, 
	0, 457, 458, 5, 28, 0, 0, 458, 460, 1, 0, 0, 0, 459, 448, 1, 0, 0, 0, 459, 
	450, 1, 0, 0, 0, 459, 451, 1, 0, 0, 0, 459, 454, 1, 0, 0, 0, 460, 483, 
	1, 0, 0, 0, 461, 462, 10, 4, 0, 0, 462, 463, 3, 106, 53, 0, 463, 464, 3, 
	68, 34, 5, 464, 482, 1, 0, 0, 0, 465, 466, 10, 3, 0, 0, 466, 467, 3, 110, 
	55, 0, 467, 468, 3, 68, 34, 4, 468, 482, 1, 0, 0, 0, 469, 470, 10, 2, 0, 
	0, 470, 471, 3, 108, 54, 0, 471, 472, 3, 68, 34, 3, 472, 482, 1, 0, 0, 
	0, 473, 474, 10, 7, 0, 0, 474, 475, 5, 23, 0, 0, 475, 476, 3, 20, 10, 0, 
	476, 477, 5, 24, 0, 0, 477, 482, 1, 0, 0, 0, 478, 479, 10, 6, 0, 0, 479, 
	480, 5, 46, 0, 0, 480, 482, 5, 113, 0, 0, 481, 461, 1, 0, 0, 0, 481, 465, 
	1, 0, 0, 0, 481, 469, 1, 0, 0, 0, 481, 473, 1, 0, 0, 0, 481, 478, 1, 0, 
	0, 0, 482, 485, 1, 0, 0, 0, 483, 481, 1, 0, 0, 0, 483, 484, 1, 0, 0, 0, 
	484, 69, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0, 486, 487, 5, 47, 0, 0, 487, 
	518, 3, 68, 34, 0, 488, 489, 3, 68, 34, 0, 489, 490, 5, 48, 0, 0, 490, 
	491, 3, 68, 34, 0, 491, 518, 1, 0, 0, 0, 492, 493, 5, 49, 0, 0, 493, 494, 
	5, 27, 0, 0, 494, 495, 3, 68, 34, 0, 495, 496, 5, 2, 0, 0, 496, 497, 3, 
	20, 10, 0, 497, 498, 5, 2, 0, 0, 498, 499, 3, 68, 34, 0, 499, 500, 5, 28, 
	0, 0, 500, 518, 1, 0, 0, 0, 501, 502, 3, 68, 34, 0, 502, 503, 5, 49, 0, 
	0, 503, 504, 3, 68, 34, 0, 504, 518, 1, 0, 0, 0, 505, 506, 3, 68, 34, 0, 
	506, 507, 5, 33, 0, 0, 507, 508, 3, 44, 22, 0, 508, 518, 1, 0, 0, 0, 509, 
	510, 5, 50, 0, 0, 510, 511, 5, 113, 0, 0, 511, 512, 5, 27, 0, 0, 512, 513, 
	3, 68, 34, 0, 513, 514, 5, 19, 0, 0, 514, 515, 3, 68, 34, 0, 515, 516, 
	5, 28, 0, 0, 516, 518, 1, 0, 0, 0, 517, 486, 1, 0, 0, 0, 517, 488, 1, 0, 
	0, 0, 517, 492, 1, 0, 0, 0, 517, 501, 1, 0, 0, 0, 517, 505, 1, 0, 0, 0, 
	517, 509, 1, 0, 0, 0, 518, 71, 1, 0, 0, 0, 519, 520, 5, 51, 0, 0, 520, 
	521, 5, 27, 0, 0, 521, 522, 3, 102, 51, 0, 522, 523, 5, 28, 0, 0, 523, 
	73, 1, 0, 0, 0, 524, 525, 3, 68, 34, 0, 525, 526, 5, 23, 0, 0, 526, 527, 
	5, 115, 0, 0, 527, 528, 5, 52, 0, 0, 528, 529, 5, 115, 0, 0, 529, 530, 
	5, 24, 0, 0, 530, 552, 1, 0, 0, 0, 531, 532, 5, 27, 0, 0, 532, 533, 3, 
	68, 34, 0, 533, 535, 5, 46, 0, 0, 534, 536, 3, 84, 42, 0, 535, 534, 1, 
	0, 0, 0, 536, 537, 1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 537, 538, 1, 0, 0, 
	0, 538, 539, 1, 0, 0, 0, 539, 540, 5, 53, 0, 0, 540, 541, 3, 68, 34, 0, 
	541, 542, 5, 28, 0, 0, 542, 552, 1, 0, 0, 0, 543, 544, 3, 68, 34, 0, 544, 
	545, 5, 20, 0, 0, 545, 546, 3, 20, 10, 0, 546, 552, 1, 0, 0, 0, 547, 548, 
	5, 23, 0, 0, 548, 549, 3, 102, 51, 0, 549, 550, 5, 24, 0, 0, 550, 552, 
	1, 0, 0, 0, 551, 524, 1, 0, 0, 0, 551, 531, 1, 0, 0, 0, 551, 543, 1, 0, 
	0, 0, 551, 547, 1, 0, 0, 0, 552, 75, 1, 0, 0, 0, 553, 554, 5, 5, 0, 0, 
	554, 558, 3, 78, 39, 0, 555, 557, 3, 78, 39, 0, 556, 555, 1, 0, 0, 0, 557, 
	560, 1, 0, 0, 0, 558, 556, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 561, 
	1, 0, 0, 0, 560, 558, 1, 0, 0, 0, 561, 562, 5, 6, 0, 0, 562, 77, 1, 0, 
	0, 0, 563, 564, 5, 113, 0, 0, 564, 565, 5, 21, 0, 0, 565, 566, 3, 68, 34, 
	0, 566, 79, 1, 0, 0, 0, 567, 568, 5, 27, 0, 0, 568, 569, 5, 113, 0, 0, 
	569, 571, 5, 54, 0, 0, 570, 572, 3, 82, 41, 0, 571, 570, 1, 0, 0, 0, 572, 
	573, 1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 575, 
	1, 0, 0, 0, 575, 576, 5, 3, 0, 0, 576, 577, 3, 68, 34, 0, 577, 578, 5, 
	28, 0, 0, 578, 81, 1, 0, 0, 0, 579, 580, 5, 46, 0, 0, 580, 583, 5, 113, 
	0, 0, 581, 583, 3, 84, 42, 0, 582, 579, 1, 0, 0, 0, 582, 581, 1, 0, 0, 
	0, 583, 83, 1, 0, 0, 0, 584, 585, 5, 23, 0, 0, 585, 586, 3, 68, 34, 0, 
	586, 587, 5, 24, 0, 0, 587, 85, 1, 0, 0, 0, 588, 589, 5, 42, 0, 0, 589, 
	590, 3, 68, 34, 0, 590, 591, 5, 55, 0, 0, 591, 592, 3, 68, 34, 0, 592, 
	593, 5, 56, 0, 0, 593, 594, 3, 68, 34, 0, 594, 607, 1, 0, 0, 0, 595, 596, 
	5, 27, 0, 0, 596, 597, 5, 57, 0, 0, 597, 598, 3, 68, 34, 0, 598, 600, 5, 
	58, 0, 0, 599, 601, 3, 88, 44, 0, 600, 599, 1, 0, 0, 0, 601, 602, 1, 0, 
	0, 0, 602, 600, 1, 0, 0, 0, 602, 603, 1, 0, 0, 0, 603, 604, 1, 0, 0, 0, 
	604, 605, 5, 28, 0, 0, 605, 607, 1, 0, 0, 0, 606, 588, 1, 0, 0, 0, 606, 
	595, 1, 0, 0, 0, 607, 87, 1, 0, 0, 0, 608, 609, 5, 59, 0, 0, 609, 610, 
	3, 90, 45, 0, 610, 611, 5, 21, 0, 0, 611, 612, 3, 68, 34, 0, 612, 89, 1, 
	0, 0, 0, 613, 623, 5, 113, 0, 0, 614, 623, 5, 114, 0, 0, 615, 617, 5, 60, 
	0, 0, 616, 615, 1, 0, 0, 0, 616, 617, 1, 0, 0, 0, 617, 618, 1, 0, 0, 0, 
	618, 623, 5, 115, 0, 0, 619, 623, 5, 61, 0, 0, 620, 623, 5, 62, 0, 0, 621, 
	623, 5, 63, 0, 0, 622, 613, 1, 0, 0, 0, 622, 614, 1, 0, 0, 0, 622, 616, 
	1, 0, 0, 0, 622, 619, 1, 0, 0, 0, 622, 620, 1, 0, 0, 0, 622, 621, 1, 0, 
	0, 0, 623, 91, 1, 0, 0, 0, 624, 625, 3, 94, 47, 0, 625, 626, 5, 27, 0, 
	0, 626, 627, 3, 102, 51, 0, 627, 628, 5, 28, 0, 0, 628, 706, 1, 0, 0, 0, 
	629, 630, 5, 27, 0, 0, 630, 631, 3, 100, 50, 0, 631, 632, 5, 64, 0, 0, 
	632, 633, 3, 94, 47, 0, 633, 634, 5, 2, 0, 0, 634, 635, 3, 20, 10, 0, 635, 
	636, 5, 65, 0, 0, 636, 637, 5, 28, 0, 0, 637, 638, 5, 27, 0, 0, 638, 639, 
	3, 102, 51, 0, 639, 640, 5, 28, 0, 0, 640, 706, 1, 0, 0, 0, 641, 642, 5, 
	27, 0, 0, 642, 643, 5, 66, 0, 0, 643, 644, 5, 64, 0, 0, 644, 645, 3, 94, 
	47, 0, 645, 646, 5, 2, 0, 0, 646, 647, 3, 20, 10, 0, 647, 648, 5, 65, 0, 
	0, 648, 649, 5, 42, 0, 0, 649, 650, 3, 68, 34, 0, 650, 651, 5, 53, 0, 0, 
	651, 652, 5, 27, 0, 0, 652, 653, 3, 102, 51, 0, 653, 654, 5, 28, 0, 0, 
	654, 655, 5, 28, 0, 0, 655, 656, 5, 27, 0, 0, 656, 657, 3, 102, 51, 0, 
	657, 658, 5, 28, 0, 0, 658, 706, 1, 0, 0, 0, 659, 660, 5, 27, 0, 0, 660, 
	661, 5, 67, 0, 0, 661, 662, 5, 64, 0, 0, 662, 663, 3, 94, 47, 0, 663, 664, 
	5, 2, 0, 0, 664, 665, 3, 20, 10, 0, 665, 666, 5, 65, 0, 0, 666, 667, 5, 
	42, 0, 0, 667, 668, 3, 68, 34, 0, 668, 669, 5, 53, 0, 0, 669, 670, 5, 27, 
	0, 0, 670, 671, 3, 102, 51, 0, 671, 672, 5, 28, 0, 0, 672, 673, 5, 28, 
	0, 0, 673, 674, 5, 27, 0, 0, 674, 675, 3, 102, 51, 0, 675, 676, 5, 28, 
	0, 0, 676, 706, 1, 0, 0, 0, 677, 678, 5, 27, 0, 0, 678, 679, 5, 68, 0, 
	0, 679, 680, 5, 64, 0, 0, 680, 681, 3, 94, 47, 0, 681, 682, 5, 2, 0, 0, 
	682, 683, 3, 20, 10, 0, 683, 684, 5, 65, 0, 0, 684, 685, 5, 42, 0, 0, 685, 
	686, 3, 68, 34, 0, 686, 687, 5, 28, 0, 0, 687, 688, 5, 27, 0, 0, 688, 689, 
	3, 102, 51, 0, 689, 690, 5, 28, 0, 0, 690, 706, 1, 0, 0, 0, 691, 692, 5, 
	27, 0, 0, 692, 693, 5, 69, 0, 0, 693, 694, 5, 64, 0, 0, 694, 695, 3, 94, 
	47, 0, 695, 696, 5, 2, 0, 0, 696, 697, 3, 20, 10, 0, 697, 698, 5, 65, 0, 
	0, 698, 699, 5, 42, 0, 0, 699, 700, 3, 68, 34, 0, 700, 701, 5, 28, 0, 0, 
	701, 702, 5, 27, 0, 0, 702, 703, 3, 102, 51, 0, 703, 704, 5, 28, 0, 0, 
	704, 706, 1, 0, 0, 0, 705, 624, 1, 0, 0, 0, 705, 629, 1, 0, 0, 0, 705, 
	641, 1, 0, 0, 0, 705, 659, 1, 0, 0, 0, 705, 677, 1, 0, 0, 0, 705, 691, 
	1, 0, 0, 0, 706, 93, 1, 0, 0, 0, 707, 717, 5, 113, 0, 0, 708, 717, 3, 96, 
	48, 0, 709, 717, 3, 98, 49, 0, 710, 711, 5, 70, 0, 0, 711, 712, 5, 113, 
	0, 0, 712, 717, 5, 28, 0, 0, 713, 714, 5, 71, 0, 0, 714, 715, 5, 113, 0, 
	0, 715, 717, 5, 28, 0, 0, 716, 707, 1, 0, 0, 0, 716, 708, 1, 0, 0, 0, 716, 
	709, 1, 0, 0, 0, 716, 710, 1, 0, 0, 0, 716, 713, 1, 0, 0, 0, 717, 95, 1, 
	0, 0, 0, 718, 719, 7, 2, 0, 0, 719, 97, 1, 0, 0, 0, 720, 721, 7, 3, 0, 
	0, 721, 99, 1, 0, 0, 0, 722, 723, 7, 4, 0, 0, 723, 101, 1, 0, 0, 0, 724, 
	728, 3, 68, 34, 0, 725, 727, 3, 68, 34, 0, 726, 725, 1, 0, 0, 0, 727, 730, 
	1, 0, 0, 0, 728, 726, 1, 0, 0, 0, 728, 729, 1, 0, 0, 0, 729, 732, 1, 0, 
	0, 0, 730, 728, 1, 0, 0, 0, 731, 724, 1, 0, 0, 0, 731, 732, 1, 0, 0, 0, 
	732, 103, 1, 0, 0, 0, 733, 734, 7, 5, 0, 0, 734, 105, 1, 0, 0, 0, 735, 
	736, 7, 6, 0, 0, 736, 107, 1, 0, 0, 0, 737, 738, 7, 7, 0, 0, 738, 109, 
	1, 0, 0, 0, 739, 740, 7, 8, 0, 0, 740, 111, 1, 0, 0, 0, 741, 742, 7, 9, 
	0, 0, 742, 113, 1, 0, 0, 0, 60, 117, 123, 131, 137, 146, 150, 172, 177, 
	184, 199, 207, 225, 230, 244, 246, 254, 275, 278, 287, 295, 300, 302, 310, 
	317, 323, 326, 336, 349, 357, 360, 370, 374, 379, 382, 385, 395, 398, 407, 
	410, 416, 424, 428, 446, 459, 481, 483, 517, 537, 551, 558, 573, 582, 602, 
	606, 616, 622, 705, 716, 728, 731,
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

// SynlongParserInit initializes any static state used to implement SynlongParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSynlongParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SynlongParserInit() {
  staticData := &SynlongParserStaticData
  staticData.once.Do(synlongParserInit)
}

// NewSynlongParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSynlongParser(input antlr.TokenStream) *SynlongParser {
	SynlongParserInit()
	this := new(SynlongParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &SynlongParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Synlong.g4"

	return this
}


// SynlongParser tokens.
const (
	SynlongParserEOF = antlr.TokenEOF
	SynlongParserT__0 = 1
	SynlongParserT__1 = 2
	SynlongParserT__2 = 3
	SynlongParserT__3 = 4
	SynlongParserT__4 = 5
	SynlongParserT__5 = 6
	SynlongParserT__6 = 7
	SynlongParserT__7 = 8
	SynlongParserT__8 = 9
	SynlongParserT__9 = 10
	SynlongParserT__10 = 11
	SynlongParserT__11 = 12
	SynlongParserT__12 = 13
	SynlongParserT__13 = 14
	SynlongParserT__14 = 15
	SynlongParserT__15 = 16
	SynlongParserT__16 = 17
	SynlongParserT__17 = 18
	SynlongParserT__18 = 19
	SynlongParserT__19 = 20
	SynlongParserT__20 = 21
	SynlongParserT__21 = 22
	SynlongParserT__22 = 23
	SynlongParserT__23 = 24
	SynlongParserT__24 = 25
	SynlongParserT__25 = 26
	SynlongParserT__26 = 27
	SynlongParserT__27 = 28
	SynlongParserT__28 = 29
	SynlongParserT__29 = 30
	SynlongParserT__30 = 31
	SynlongParserT__31 = 32
	SynlongParserT__32 = 33
	SynlongParserT__33 = 34
	SynlongParserT__34 = 35
	SynlongParserT__35 = 36
	SynlongParserT__36 = 37
	SynlongParserT__37 = 38
	SynlongParserT__38 = 39
	SynlongParserT__39 = 40
	SynlongParserT__40 = 41
	SynlongParserT__41 = 42
	SynlongParserT__42 = 43
	SynlongParserT__43 = 44
	SynlongParserT__44 = 45
	SynlongParserT__45 = 46
	SynlongParserT__46 = 47
	SynlongParserT__47 = 48
	SynlongParserT__48 = 49
	SynlongParserT__49 = 50
	SynlongParserT__50 = 51
	SynlongParserT__51 = 52
	SynlongParserT__52 = 53
	SynlongParserT__53 = 54
	SynlongParserT__54 = 55
	SynlongParserT__55 = 56
	SynlongParserT__56 = 57
	SynlongParserT__57 = 58
	SynlongParserT__58 = 59
	SynlongParserT__59 = 60
	SynlongParserT__60 = 61
	SynlongParserT__61 = 62
	SynlongParserT__62 = 63
	SynlongParserT__63 = 64
	SynlongParserT__64 = 65
	SynlongParserT__65 = 66
	SynlongParserT__66 = 67
	SynlongParserT__67 = 68
	SynlongParserT__68 = 69
	SynlongParserT__69 = 70
	SynlongParserT__70 = 71
	SynlongParserT__71 = 72
	SynlongParserT__72 = 73
	SynlongParserT__73 = 74
	SynlongParserT__74 = 75
	SynlongParserT__75 = 76
	SynlongParserT__76 = 77
	SynlongParserT__77 = 78
	SynlongParserT__78 = 79
	SynlongParserT__79 = 80
	SynlongParserT__80 = 81
	SynlongParserT__81 = 82
	SynlongParserT__82 = 83
	SynlongParserT__83 = 84
	SynlongParserT__84 = 85
	SynlongParserT__85 = 86
	SynlongParserT__86 = 87
	SynlongParserT__87 = 88
	SynlongParserT__88 = 89
	SynlongParserT__89 = 90
	SynlongParserT__90 = 91
	SynlongParserT__91 = 92
	SynlongParserT__92 = 93
	SynlongParserT__93 = 94
	SynlongParserT__94 = 95
	SynlongParserT__95 = 96
	SynlongParserT__96 = 97
	SynlongParserT__97 = 98
	SynlongParserT__98 = 99
	SynlongParserT__99 = 100
	SynlongParserT__100 = 101
	SynlongParserT__101 = 102
	SynlongParserT__102 = 103
	SynlongParserT__103 = 104
	SynlongParserT__104 = 105
	SynlongParserT__105 = 106
	SynlongParserT__106 = 107
	SynlongParserT__107 = 108
	SynlongParserT__108 = 109
	SynlongParserT__109 = 110
	SynlongParserT__110 = 111
	SynlongParserCOMMENT = 112
	SynlongParserID = 113
	SynlongParserCHAR = 114
	SynlongParserINTEGER = 115
	SynlongParserUINT = 116
	SynlongParserFLOAT = 117
	SynlongParserREAL = 118
	SynlongParserUSHORT = 119
	SynlongParserSHORT = 120
	SynlongParserWS = 121
)

// SynlongParser rules.
const (
	SynlongParserRULE_program = 0
	SynlongParserRULE_decls = 1
	SynlongParserRULE_type_block = 2
	SynlongParserRULE_type_decl = 3
	SynlongParserRULE_type_def = 4
	SynlongParserRULE_type_expr = 5
	SynlongParserRULE_field_decl = 6
	SynlongParserRULE_typevar = 7
	SynlongParserRULE_const_block = 8
	SynlongParserRULE_const_decl = 9
	SynlongParserRULE_const_expr = 10
	SynlongParserRULE_const_list = 11
	SynlongParserRULE_const_label_expr = 12
	SynlongParserRULE_user_op_decl = 13
	SynlongParserRULE_op_kind = 14
	SynlongParserRULE_params = 15
	SynlongParserRULE_returns_clause = 16
	SynlongParserRULE_op_body = 17
	SynlongParserRULE_local_block = 18
	SynlongParserRULE_var_decls = 19
	SynlongParserRULE_var_id = 20
	SynlongParserRULE_when_decl = 21
	SynlongParserRULE_clock_expr = 22
	SynlongParserRULE_last_decl = 23
	SynlongParserRULE_equation = 24
	SynlongParserRULE_lhs = 25
	SynlongParserRULE_lhs_id = 26
	SynlongParserRULE_return_statement = 27
	SynlongParserRULE_returns_var = 28
	SynlongParserRULE_state_machine = 29
	SynlongParserRULE_state_decl = 30
	SynlongParserRULE_data_def = 31
	SynlongParserRULE_transition = 32
	SynlongParserRULE_expr = 33
	SynlongParserRULE_simple_expr = 34
	SynlongParserRULE_tempo_expr = 35
	SynlongParserRULE_bool_expr = 36
	SynlongParserRULE_array_expr = 37
	SynlongParserRULE_struct_expr = 38
	SynlongParserRULE_label_expr = 39
	SynlongParserRULE_mixed_constructor = 40
	SynlongParserRULE_label_or_index = 41
	SynlongParserRULE_index = 42
	SynlongParserRULE_switch_expr = 43
	SynlongParserRULE_case_expr = 44
	SynlongParserRULE_pattern = 45
	SynlongParserRULE_apply_expr = 46
	SynlongParserRULE_prefix_operator = 47
	SynlongParserRULE_prefix_unary_operator = 48
	SynlongParserRULE_prefix_binary_operator = 49
	SynlongParserRULE_iterator = 50
	SynlongParserRULE_list = 51
	SynlongParserRULE_unary_arith_op = 52
	SynlongParserRULE_bin_arith_op = 53
	SynlongParserRULE_bin_relation_op = 54
	SynlongParserRULE_bin_bool_op = 55
	SynlongParserRULE_atom = 56
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDecls() []IDeclsContext
	Decls(i int) IDeclsContext

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
	p.RuleIndex = SynlongParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDecls() []IDeclsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclsContext); ok {
			len++
		}
	}

	tst := make([]IDeclsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclsContext); ok {
			tst[i] = t.(IDeclsContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Decls(i int) IDeclsContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclsContext); ok {
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

	return t.(IDeclsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SynlongParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 104857602) != 0) {
		{
			p.SetState(114)
			p.Decls()
		}


		p.SetState(119)
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


// IDeclsContext is an interface to support dynamic dispatch.
type IDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_block() IType_blockContext
	Const_block() IConst_blockContext
	User_op_decl() IUser_op_declContext

	// IsDeclsContext differentiates from other interfaces.
	IsDeclsContext()
}

type DeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclsContext() *DeclsContext {
	var p = new(DeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_decls
	return p
}

func InitEmptyDeclsContext(p *DeclsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_decls
}

func (*DeclsContext) IsDeclsContext() {}

func NewDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclsContext {
	var p = new(DeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_decls

	return p
}

func (s *DeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclsContext) Type_block() IType_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_blockContext)
}

func (s *DeclsContext) Const_block() IConst_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_blockContext)
}

func (s *DeclsContext) User_op_decl() IUser_op_declContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUser_op_declContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUser_op_declContext)
}

func (s *DeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterDecls(s)
	}
}

func (s *DeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitDecls(s)
	}
}

func (s *DeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitDecls(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Decls() (localctx IDeclsContext) {
	localctx = NewDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SynlongParserRULE_decls)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Type_block()
		}


	case SynlongParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.Const_block()
		}


	case SynlongParserT__24, SynlongParserT__25:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)
			p.User_op_decl()
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


// IType_blockContext is an interface to support dynamic dispatch.
type IType_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllType_decl() []IType_declContext
	Type_decl(i int) IType_declContext

	// IsType_blockContext differentiates from other interfaces.
	IsType_blockContext()
}

type Type_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_blockContext() *Type_blockContext {
	var p = new(Type_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_block
	return p
}

func InitEmptyType_blockContext(p *Type_blockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_block
}

func (*Type_blockContext) IsType_blockContext() {}

func NewType_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_blockContext {
	var p = new(Type_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_type_block

	return p
}

func (s *Type_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_blockContext) AllType_decl() []IType_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_declContext); ok {
			len++
		}
	}

	tst := make([]IType_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_declContext); ok {
			tst[i] = t.(IType_declContext)
			i++
		}
	}

	return tst
}

func (s *Type_blockContext) Type_decl(i int) IType_declContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_declContext); ok {
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

	return t.(IType_declContext)
}

func (s *Type_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Type_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterType_block(s)
	}
}

func (s *Type_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitType_block(s)
	}
}

func (s *Type_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitType_block(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Type_block() (localctx IType_blockContext) {
	localctx = NewType_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SynlongParserRULE_type_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(SynlongParserT__0)
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


	for _la == SynlongParserID {
		{
			p.SetState(126)
			p.Type_decl()
		}
		{
			p.SetState(127)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(133)
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


// IType_declContext is an interface to support dynamic dispatch.
type IType_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_def() IType_defContext

	// IsType_declContext differentiates from other interfaces.
	IsType_declContext()
}

type Type_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_declContext() *Type_declContext {
	var p = new(Type_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_decl
	return p
}

func InitEmptyType_declContext(p *Type_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_decl
}

func (*Type_declContext) IsType_declContext() {}

func NewType_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declContext {
	var p = new(Type_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_type_decl

	return p
}

func (s *Type_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Type_declContext) Type_def() IType_defContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_defContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_defContext)
}

func (s *Type_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Type_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterType_decl(s)
	}
}

func (s *Type_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitType_decl(s)
	}
}

func (s *Type_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitType_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Type_decl() (localctx IType_declContext) {
	localctx = NewType_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SynlongParserRULE_type_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__2 {
		{
			p.SetState(135)
			p.Match(SynlongParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Type_def()
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


// IType_defContext is an interface to support dynamic dispatch.
type IType_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_expr() IType_exprContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsType_defContext differentiates from other interfaces.
	IsType_defContext()
}

type Type_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_defContext() *Type_defContext {
	var p = new(Type_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_def
	return p
}

func InitEmptyType_defContext(p *Type_defContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_def
}

func (*Type_defContext) IsType_defContext() {}

func NewType_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_defContext {
	var p = new(Type_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_type_def

	return p
}

func (s *Type_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_defContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Type_defContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SynlongParserID)
}

func (s *Type_defContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SynlongParserID, i)
}

func (s *Type_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Type_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterType_def(s)
	}
}

func (s *Type_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitType_def(s)
	}
}

func (s *Type_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitType_def(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Type_def() (localctx IType_defContext) {
	localctx = NewType_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SynlongParserRULE_type_def)
	var _la int

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__4, SynlongParserT__6, SynlongParserT__7, SynlongParserT__8, SynlongParserT__9, SynlongParserT__10, SynlongParserT__11, SynlongParserT__12, SynlongParserT__13, SynlongParserT__14, SynlongParserT__15, SynlongParserT__16, SynlongParserT__17, SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.type_expr(0)
		}


	case SynlongParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.Match(SynlongParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Match(SynlongParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserID {
			{
				p.SetState(143)
				p.Match(SynlongParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(149)
			p.Match(SynlongParserT__5)
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


// IType_exprContext is an interface to support dynamic dispatch.
type IType_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Typevar() ITypevarContext
	AllField_decl() []IField_declContext
	Field_decl(i int) IField_declContext
	Type_expr() IType_exprContext
	Const_expr() IConst_exprContext

	// IsType_exprContext differentiates from other interfaces.
	IsType_exprContext()
}

type Type_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_exprContext() *Type_exprContext {
	var p = new(Type_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_expr
	return p
}

func InitEmptyType_exprContext(p *Type_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_type_expr
}

func (*Type_exprContext) IsType_exprContext() {}

func NewType_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_exprContext {
	var p = new(Type_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_type_expr

	return p
}

func (s *Type_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_exprContext) Typevar() ITypevarContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypevarContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypevarContext)
}

func (s *Type_exprContext) AllField_decl() []IField_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IField_declContext); ok {
			len++
		}
	}

	tst := make([]IField_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IField_declContext); ok {
			tst[i] = t.(IField_declContext)
			i++
		}
	}

	return tst
}

func (s *Type_exprContext) Field_decl(i int) IField_declContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField_declContext); ok {
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

	return t.(IField_declContext)
}

func (s *Type_exprContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Type_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Type_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Type_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterType_expr(s)
	}
}

func (s *Type_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitType_expr(s)
	}
}

func (s *Type_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitType_expr(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SynlongParser) Type_expr() (localctx IType_exprContext) {
	return p.type_expr(0)
}

func (p *SynlongParser) type_expr(_p int) (localctx IType_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewType_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IType_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SynlongParserRULE_type_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__6:
		{
			p.SetState(153)
			p.Match(SynlongParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__7:
		{
			p.SetState(154)
			p.Match(SynlongParserT__7)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__8:
		{
			p.SetState(155)
			p.Match(SynlongParserT__8)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__9:
		{
			p.SetState(156)
			p.Match(SynlongParserT__9)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__10:
		{
			p.SetState(157)
			p.Match(SynlongParserT__10)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__11:
		{
			p.SetState(158)
			p.Match(SynlongParserT__11)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__12:
		{
			p.SetState(159)
			p.Match(SynlongParserT__12)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__13:
		{
			p.SetState(160)
			p.Match(SynlongParserT__13)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__14:
		{
			p.SetState(161)
			p.Match(SynlongParserT__14)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__15:
		{
			p.SetState(162)
			p.Match(SynlongParserT__15)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__16:
		{
			p.SetState(163)
			p.Match(SynlongParserT__16)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__17:
		{
			p.SetState(164)
			p.Match(SynlongParserT__17)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserID:
		{
			p.SetState(165)
			p.Typevar()
		}


	case SynlongParserT__4:
		{
			p.SetState(166)
			p.Match(SynlongParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Field_decl()
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__18 {
			{
				p.SetState(168)
				p.Match(SynlongParserT__18)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(169)
				p.Field_decl()
			}


			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(175)
			p.Match(SynlongParserT__5)
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
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewType_exprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_type_expr)
			p.SetState(179)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(180)
				p.Match(SynlongParserT__19)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(181)
				p.const_expr(0)
			}


		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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


// IField_declContext is an interface to support dynamic dispatch.
type IField_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_expr() IType_exprContext

	// IsField_declContext differentiates from other interfaces.
	IsField_declContext()
}

type Field_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_declContext() *Field_declContext {
	var p = new(Field_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_field_decl
	return p
}

func InitEmptyField_declContext(p *Field_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_field_decl
}

func (*Field_declContext) IsField_declContext() {}

func NewField_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_declContext {
	var p = new(Field_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_field_decl

	return p
}

func (s *Field_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_declContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Field_declContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Field_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Field_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterField_decl(s)
	}
}

func (s *Field_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitField_decl(s)
	}
}

func (s *Field_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitField_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Field_decl() (localctx IField_declContext) {
	localctx = NewField_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SynlongParserRULE_field_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(188)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(189)
		p.type_expr(0)
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


// ITypevarContext is an interface to support dynamic dispatch.
type ITypevarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsTypevarContext differentiates from other interfaces.
	IsTypevarContext()
}

type TypevarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypevarContext() *TypevarContext {
	var p = new(TypevarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_typevar
	return p
}

func InitEmptyTypevarContext(p *TypevarContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_typevar
}

func (*TypevarContext) IsTypevarContext() {}

func NewTypevarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypevarContext {
	var p = new(TypevarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_typevar

	return p
}

func (s *TypevarContext) GetParser() antlr.Parser { return s.parser }

func (s *TypevarContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *TypevarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypevarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypevarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterTypevar(s)
	}
}

func (s *TypevarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitTypevar(s)
	}
}

func (s *TypevarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitTypevar(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Typevar() (localctx ITypevarContext) {
	localctx = NewTypevarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SynlongParserRULE_typevar)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(SynlongParserID)
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


// IConst_blockContext is an interface to support dynamic dispatch.
type IConst_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_decl() []IConst_declContext
	Const_decl(i int) IConst_declContext

	// IsConst_blockContext differentiates from other interfaces.
	IsConst_blockContext()
}

type Const_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_blockContext() *Const_blockContext {
	var p = new(Const_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_block
	return p
}

func InitEmptyConst_blockContext(p *Const_blockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_block
}

func (*Const_blockContext) IsConst_blockContext() {}

func NewConst_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_blockContext {
	var p = new(Const_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_const_block

	return p
}

func (s *Const_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_blockContext) AllConst_decl() []IConst_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_declContext); ok {
			len++
		}
	}

	tst := make([]IConst_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_declContext); ok {
			tst[i] = t.(IConst_declContext)
			i++
		}
	}

	return tst
}

func (s *Const_blockContext) Const_decl(i int) IConst_declContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_declContext); ok {
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

	return t.(IConst_declContext)
}

func (s *Const_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Const_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterConst_block(s)
	}
}

func (s *Const_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitConst_block(s)
	}
}

func (s *Const_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitConst_block(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Const_block() (localctx IConst_blockContext) {
	localctx = NewConst_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SynlongParserRULE_const_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(SynlongParserT__21)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == SynlongParserID {
		{
			p.SetState(194)
			p.Const_decl()
		}
		{
			p.SetState(195)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(201)
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


// IConst_declContext is an interface to support dynamic dispatch.
type IConst_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Type_expr() IType_exprContext
	Const_expr() IConst_exprContext

	// IsConst_declContext differentiates from other interfaces.
	IsConst_declContext()
}

type Const_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_declContext() *Const_declContext {
	var p = new(Const_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_decl
	return p
}

func InitEmptyConst_declContext(p *Const_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_decl
}

func (*Const_declContext) IsConst_declContext() {}

func NewConst_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_declContext {
	var p = new(Const_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_const_decl

	return p
}

func (s *Const_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_declContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Const_declContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Const_declContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Const_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Const_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterConst_decl(s)
	}
}

func (s *Const_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitConst_decl(s)
	}
}

func (s *Const_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitConst_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Const_decl() (localctx IConst_declContext) {
	localctx = NewConst_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SynlongParserRULE_const_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(204)
		p.type_expr(0)
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__2 {
		{
			p.SetState(205)
			p.Match(SynlongParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(206)
			p.const_expr(0)
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


// IConst_exprContext is an interface to support dynamic dispatch.
type IConst_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Atom() IAtomContext
	Unary_arith_op() IUnary_arith_opContext
	AllConst_expr() []IConst_exprContext
	Const_expr(i int) IConst_exprContext
	Const_list() IConst_listContext
	AllConst_label_expr() []IConst_label_exprContext
	Const_label_expr(i int) IConst_label_exprContext
	Bin_arith_op() IBin_arith_opContext
	Bin_bool_op() IBin_bool_opContext
	Bin_relation_op() IBin_relation_opContext

	// IsConst_exprContext differentiates from other interfaces.
	IsConst_exprContext()
}

type Const_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_exprContext() *Const_exprContext {
	var p = new(Const_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_expr
	return p
}

func InitEmptyConst_exprContext(p *Const_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_expr
}

func (*Const_exprContext) IsConst_exprContext() {}

func NewConst_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_exprContext {
	var p = new(Const_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_const_expr

	return p
}

func (s *Const_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Const_exprContext) Atom() IAtomContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Const_exprContext) Unary_arith_op() IUnary_arith_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_arith_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_arith_opContext)
}

func (s *Const_exprContext) AllConst_expr() []IConst_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_exprContext); ok {
			len++
		}
	}

	tst := make([]IConst_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_exprContext); ok {
			tst[i] = t.(IConst_exprContext)
			i++
		}
	}

	return tst
}

func (s *Const_exprContext) Const_expr(i int) IConst_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
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

	return t.(IConst_exprContext)
}

func (s *Const_exprContext) Const_list() IConst_listContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_listContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_listContext)
}

func (s *Const_exprContext) AllConst_label_expr() []IConst_label_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_label_exprContext); ok {
			len++
		}
	}

	tst := make([]IConst_label_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_label_exprContext); ok {
			tst[i] = t.(IConst_label_exprContext)
			i++
		}
	}

	return tst
}

func (s *Const_exprContext) Const_label_expr(i int) IConst_label_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_label_exprContext); ok {
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

	return t.(IConst_label_exprContext)
}

func (s *Const_exprContext) Bin_arith_op() IBin_arith_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_arith_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_arith_opContext)
}

func (s *Const_exprContext) Bin_bool_op() IBin_bool_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_bool_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_bool_opContext)
}

func (s *Const_exprContext) Bin_relation_op() IBin_relation_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_relation_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_relation_opContext)
}

func (s *Const_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Const_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterConst_expr(s)
	}
}

func (s *Const_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitConst_expr(s)
	}
}

func (s *Const_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitConst_expr(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SynlongParser) Const_expr() (localctx IConst_exprContext) {
	return p.const_expr(0)
}

func (p *SynlongParser) const_expr(_p int) (localctx IConst_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewConst_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConst_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, SynlongParserRULE_const_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserID:
		{
			p.SetState(210)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__60, SynlongParserT__61, SynlongParserCHAR, SynlongParserINTEGER, SynlongParserUINT, SynlongParserFLOAT, SynlongParserREAL, SynlongParserUSHORT, SynlongParserSHORT:
		{
			p.SetState(211)
			p.Atom()
		}


	case SynlongParserT__33, SynlongParserT__59, SynlongParserT__98:
		{
			p.SetState(212)
			p.Unary_arith_op()
		}
		{
			p.SetState(213)
			p.const_expr(6)
		}


	case SynlongParserT__22:
		{
			p.SetState(215)
			p.Match(SynlongParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Const_list()
		}
		{
			p.SetState(217)
			p.Match(SynlongParserT__23)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__4:
		{
			p.SetState(219)
			p.Match(SynlongParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(220)
			p.Const_label_expr()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__18 {
			{
				p.SetState(221)
				p.Match(SynlongParserT__18)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(222)
				p.Const_label_expr()
			}


			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(228)
			p.Match(SynlongParserT__5)
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
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(244)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewConst_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_const_expr)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(233)
					p.Bin_arith_op()
				}
				{
					p.SetState(234)
					p.const_expr(6)
				}


			case 2:
				localctx = NewConst_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_const_expr)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(237)
					p.Bin_bool_op()
				}
				{
					p.SetState(238)
					p.const_expr(5)
				}


			case 3:
				localctx = NewConst_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_const_expr)
				p.SetState(240)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(241)
					p.Bin_relation_op()
				}
				{
					p.SetState(242)
					p.const_expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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


// IConst_listContext is an interface to support dynamic dispatch.
type IConst_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConst_expr() []IConst_exprContext
	Const_expr(i int) IConst_exprContext

	// IsConst_listContext differentiates from other interfaces.
	IsConst_listContext()
}

type Const_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_listContext() *Const_listContext {
	var p = new(Const_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_list
	return p
}

func InitEmptyConst_listContext(p *Const_listContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_list
}

func (*Const_listContext) IsConst_listContext() {}

func NewConst_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_listContext {
	var p = new(Const_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_const_list

	return p
}

func (s *Const_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_listContext) AllConst_expr() []IConst_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConst_exprContext); ok {
			len++
		}
	}

	tst := make([]IConst_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConst_exprContext); ok {
			tst[i] = t.(IConst_exprContext)
			i++
		}
	}

	return tst
}

func (s *Const_listContext) Const_expr(i int) IConst_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
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

	return t.(IConst_exprContext)
}

func (s *Const_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Const_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterConst_list(s)
	}
}

func (s *Const_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitConst_list(s)
	}
}

func (s *Const_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitConst_list(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Const_list() (localctx IConst_listContext) {
	localctx = NewConst_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SynlongParserRULE_const_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.const_expr(0)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == SynlongParserT__18 {
		{
			p.SetState(250)
			p.Match(SynlongParserT__18)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(251)
			p.const_expr(0)
		}


		p.SetState(256)
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


// IConst_label_exprContext is an interface to support dynamic dispatch.
type IConst_label_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Const_expr() IConst_exprContext

	// IsConst_label_exprContext differentiates from other interfaces.
	IsConst_label_exprContext()
}

type Const_label_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_label_exprContext() *Const_label_exprContext {
	var p = new(Const_label_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_label_expr
	return p
}

func InitEmptyConst_label_exprContext(p *Const_label_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_const_label_expr
}

func (*Const_label_exprContext) IsConst_label_exprContext() {}

func NewConst_label_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_label_exprContext {
	var p = new(Const_label_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_const_label_expr

	return p
}

func (s *Const_label_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_label_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Const_label_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Const_label_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_label_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Const_label_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterConst_label_expr(s)
	}
}

func (s *Const_label_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitConst_label_expr(s)
	}
}

func (s *Const_label_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitConst_label_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Const_label_expr() (localctx IConst_label_exprContext) {
	localctx = NewConst_label_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SynlongParserRULE_const_label_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(259)
		p.const_expr(0)
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


// IUser_op_declContext is an interface to support dynamic dispatch.
type IUser_op_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Op_kind() IOp_kindContext
	ID() antlr.TerminalNode
	Params() IParamsContext
	Returns_clause() IReturns_clauseContext
	Op_body() IOp_bodyContext

	// IsUser_op_declContext differentiates from other interfaces.
	IsUser_op_declContext()
}

type User_op_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUser_op_declContext() *User_op_declContext {
	var p = new(User_op_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_user_op_decl
	return p
}

func InitEmptyUser_op_declContext(p *User_op_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_user_op_decl
}

func (*User_op_declContext) IsUser_op_declContext() {}

func NewUser_op_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *User_op_declContext {
	var p = new(User_op_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_user_op_decl

	return p
}

func (s *User_op_declContext) GetParser() antlr.Parser { return s.parser }

func (s *User_op_declContext) Op_kind() IOp_kindContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_kindContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_kindContext)
}

func (s *User_op_declContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *User_op_declContext) Params() IParamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *User_op_declContext) Returns_clause() IReturns_clauseContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturns_clauseContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturns_clauseContext)
}

func (s *User_op_declContext) Op_body() IOp_bodyContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_bodyContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_bodyContext)
}

func (s *User_op_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *User_op_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *User_op_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterUser_op_decl(s)
	}
}

func (s *User_op_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitUser_op_decl(s)
	}
}

func (s *User_op_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitUser_op_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) User_op_decl() (localctx IUser_op_declContext) {
	localctx = NewUser_op_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SynlongParserRULE_user_op_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Op_kind()
	}
	{
		p.SetState(262)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Params()
	}
	{
		p.SetState(264)
		p.Returns_clause()
	}
	{
		p.SetState(265)
		p.Op_body()
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


// IOp_kindContext is an interface to support dynamic dispatch.
type IOp_kindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOp_kindContext differentiates from other interfaces.
	IsOp_kindContext()
}

type Op_kindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_kindContext() *Op_kindContext {
	var p = new(Op_kindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_op_kind
	return p
}

func InitEmptyOp_kindContext(p *Op_kindContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_op_kind
}

func (*Op_kindContext) IsOp_kindContext() {}

func NewOp_kindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_kindContext {
	var p = new(Op_kindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_op_kind

	return p
}

func (s *Op_kindContext) GetParser() antlr.Parser { return s.parser }
func (s *Op_kindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_kindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Op_kindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterOp_kind(s)
	}
}

func (s *Op_kindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitOp_kind(s)
	}
}

func (s *Op_kindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitOp_kind(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Op_kind() (localctx IOp_kindContext) {
	localctx = NewOp_kindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SynlongParserRULE_op_kind)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SynlongParserT__24 || _la == SynlongParserT__25) {
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


// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_decls() []IVar_declsContext
	Var_decls(i int) IVar_declsContext

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllVar_decls() []IVar_declsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declsContext); ok {
			len++
		}
	}

	tst := make([]IVar_declsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declsContext); ok {
			tst[i] = t.(IVar_declsContext)
			i++
		}
	}

	return tst
}

func (s *ParamsContext) Var_decls(i int) IVar_declsContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declsContext); ok {
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

	return t.(IVar_declsContext)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitParams(s)
	}
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SynlongParserRULE_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(SynlongParserT__26)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserID {
		{
			p.SetState(270)
			p.Var_decls()
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__1 {
			{
				p.SetState(271)
				p.Match(SynlongParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(272)
				p.Var_decls()
			}


			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(280)
		p.Match(SynlongParserT__27)
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


// IReturns_clauseContext is an interface to support dynamic dispatch.
type IReturns_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Params() IParamsContext

	// IsReturns_clauseContext differentiates from other interfaces.
	IsReturns_clauseContext()
}

type Returns_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturns_clauseContext() *Returns_clauseContext {
	var p = new(Returns_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_returns_clause
	return p
}

func InitEmptyReturns_clauseContext(p *Returns_clauseContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_returns_clause
}

func (*Returns_clauseContext) IsReturns_clauseContext() {}

func NewReturns_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returns_clauseContext {
	var p = new(Returns_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_returns_clause

	return p
}

func (s *Returns_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Returns_clauseContext) Params() IParamsContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *Returns_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Returns_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Returns_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterReturns_clause(s)
	}
}

func (s *Returns_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitReturns_clause(s)
	}
}

func (s *Returns_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitReturns_clause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Returns_clause() (localctx IReturns_clauseContext) {
	localctx = NewReturns_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SynlongParserRULE_returns_clause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(SynlongParserT__28)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(283)
		p.Params()
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


// IOp_bodyContext is an interface to support dynamic dispatch.
type IOp_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Local_block() ILocal_blockContext
	AllEquation() []IEquationContext
	Equation(i int) IEquationContext

	// IsOp_bodyContext differentiates from other interfaces.
	IsOp_bodyContext()
}

type Op_bodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_bodyContext() *Op_bodyContext {
	var p = new(Op_bodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_op_body
	return p
}

func InitEmptyOp_bodyContext(p *Op_bodyContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_op_body
}

func (*Op_bodyContext) IsOp_bodyContext() {}

func NewOp_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_bodyContext {
	var p = new(Op_bodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_op_body

	return p
}

func (s *Op_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_bodyContext) Local_block() ILocal_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocal_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocal_blockContext)
}

func (s *Op_bodyContext) AllEquation() []IEquationContext {
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

func (s *Op_bodyContext) Equation(i int) IEquationContext {
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

func (s *Op_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Op_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterOp_body(s)
	}
}

func (s *Op_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitOp_body(s)
	}
}

func (s *Op_bodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitOp_body(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Op_body() (localctx IOp_bodyContext) {
	localctx = NewOp_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SynlongParserRULE_op_body)
	var _la int

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__29, SynlongParserT__31:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == SynlongParserT__31 {
			{
				p.SetState(286)
				p.Local_block()
			}

		}
		{
			p.SetState(289)
			p.Match(SynlongParserT__29)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__26 || _la == SynlongParserT__35 || _la == SynlongParserID {
			{
				p.SetState(290)
				p.Equation()
			}
			{
				p.SetState(291)
				p.Match(SynlongParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(298)
			p.Match(SynlongParserT__30)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == SynlongParserT__1 {
			{
				p.SetState(299)
				p.Match(SynlongParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
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


// ILocal_blockContext is an interface to support dynamic dispatch.
type ILocal_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_decls() []IVar_declsContext
	Var_decls(i int) IVar_declsContext

	// IsLocal_blockContext differentiates from other interfaces.
	IsLocal_blockContext()
}

type Local_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocal_blockContext() *Local_blockContext {
	var p = new(Local_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_local_block
	return p
}

func InitEmptyLocal_blockContext(p *Local_blockContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_local_block
}

func (*Local_blockContext) IsLocal_blockContext() {}

func NewLocal_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Local_blockContext {
	var p = new(Local_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_local_block

	return p
}

func (s *Local_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Local_blockContext) AllVar_decls() []IVar_declsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_declsContext); ok {
			len++
		}
	}

	tst := make([]IVar_declsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_declsContext); ok {
			tst[i] = t.(IVar_declsContext)
			i++
		}
	}

	return tst
}

func (s *Local_blockContext) Var_decls(i int) IVar_declsContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_declsContext); ok {
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

	return t.(IVar_declsContext)
}

func (s *Local_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Local_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Local_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLocal_block(s)
	}
}

func (s *Local_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLocal_block(s)
	}
}

func (s *Local_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLocal_block(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Local_block() (localctx ILocal_blockContext) {
	localctx = NewLocal_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SynlongParserRULE_local_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(SynlongParserT__31)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == SynlongParserID {
		{
			p.SetState(305)
			p.Var_decls()
		}
		{
			p.SetState(306)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(312)
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


// IVar_declsContext is an interface to support dynamic dispatch.
type IVar_declsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVar_id() []IVar_idContext
	Var_id(i int) IVar_idContext
	Type_expr() IType_exprContext
	When_decl() IWhen_declContext
	Last_decl() ILast_declContext

	// IsVar_declsContext differentiates from other interfaces.
	IsVar_declsContext()
}

type Var_declsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_declsContext() *Var_declsContext {
	var p = new(Var_declsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_var_decls
	return p
}

func InitEmptyVar_declsContext(p *Var_declsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_var_decls
}

func (*Var_declsContext) IsVar_declsContext() {}

func NewVar_declsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_declsContext {
	var p = new(Var_declsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_var_decls

	return p
}

func (s *Var_declsContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_declsContext) AllVar_id() []IVar_idContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_idContext); ok {
			len++
		}
	}

	tst := make([]IVar_idContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_idContext); ok {
			tst[i] = t.(IVar_idContext)
			i++
		}
	}

	return tst
}

func (s *Var_declsContext) Var_id(i int) IVar_idContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_idContext); ok {
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

	return t.(IVar_idContext)
}

func (s *Var_declsContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Var_declsContext) When_decl() IWhen_declContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhen_declContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhen_declContext)
}

func (s *Var_declsContext) Last_decl() ILast_declContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILast_declContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILast_declContext)
}

func (s *Var_declsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_declsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Var_declsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterVar_decls(s)
	}
}

func (s *Var_declsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitVar_decls(s)
	}
}

func (s *Var_declsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitVar_decls(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Var_decls() (localctx IVar_declsContext) {
	localctx = NewVar_declsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SynlongParserRULE_var_decls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Var_id()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == SynlongParserID {
		{
			p.SetState(314)
			p.Var_id()
		}


		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(320)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(321)
		p.type_expr(0)
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__32 {
		{
			p.SetState(322)
			p.When_decl()
		}

	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__34 {
		{
			p.SetState(325)
			p.Last_decl()
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


// IVar_idContext is an interface to support dynamic dispatch.
type IVar_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVar_idContext differentiates from other interfaces.
	IsVar_idContext()
}

type Var_idContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_idContext() *Var_idContext {
	var p = new(Var_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_var_id
	return p
}

func InitEmptyVar_idContext(p *Var_idContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_var_id
}

func (*Var_idContext) IsVar_idContext() {}

func NewVar_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_idContext {
	var p = new(Var_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_var_id

	return p
}

func (s *Var_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_idContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Var_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Var_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterVar_id(s)
	}
}

func (s *Var_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitVar_id(s)
	}
}

func (s *Var_idContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitVar_id(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Var_id() (localctx IVar_idContext) {
	localctx = NewVar_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SynlongParserRULE_var_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(SynlongParserID)
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


// IWhen_declContext is an interface to support dynamic dispatch.
type IWhen_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Clock_expr() IClock_exprContext

	// IsWhen_declContext differentiates from other interfaces.
	IsWhen_declContext()
}

type When_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhen_declContext() *When_declContext {
	var p = new(When_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_when_decl
	return p
}

func InitEmptyWhen_declContext(p *When_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_when_decl
}

func (*When_declContext) IsWhen_declContext() {}

func NewWhen_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_declContext {
	var p = new(When_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_when_decl

	return p
}

func (s *When_declContext) GetParser() antlr.Parser { return s.parser }

func (s *When_declContext) Clock_expr() IClock_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClock_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClock_exprContext)
}

func (s *When_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *When_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterWhen_decl(s)
	}
}

func (s *When_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitWhen_decl(s)
	}
}

func (s *When_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitWhen_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) When_decl() (localctx IWhen_declContext) {
	localctx = NewWhen_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SynlongParserRULE_when_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.Match(SynlongParserT__32)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(331)
		p.Clock_expr()
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


// IClock_exprContext is an interface to support dynamic dispatch.
type IClock_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsClock_exprContext differentiates from other interfaces.
	IsClock_exprContext()
}

type Clock_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClock_exprContext() *Clock_exprContext {
	var p = new(Clock_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_clock_expr
	return p
}

func InitEmptyClock_exprContext(p *Clock_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_clock_expr
}

func (*Clock_exprContext) IsClock_exprContext() {}

func NewClock_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Clock_exprContext {
	var p = new(Clock_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_clock_expr

	return p
}

func (s *Clock_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Clock_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Clock_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Clock_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Clock_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterClock_expr(s)
	}
}

func (s *Clock_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitClock_expr(s)
	}
}

func (s *Clock_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitClock_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Clock_expr() (localctx IClock_exprContext) {
	localctx = NewClock_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SynlongParserRULE_clock_expr)
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(SynlongParserT__33)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(SynlongParserID)
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


// ILast_declContext is an interface to support dynamic dispatch.
type ILast_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const_expr() IConst_exprContext

	// IsLast_declContext differentiates from other interfaces.
	IsLast_declContext()
}

type Last_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLast_declContext() *Last_declContext {
	var p = new(Last_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_last_decl
	return p
}

func InitEmptyLast_declContext(p *Last_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_last_decl
}

func (*Last_declContext) IsLast_declContext() {}

func NewLast_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Last_declContext {
	var p = new(Last_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_last_decl

	return p
}

func (s *Last_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Last_declContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Last_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Last_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Last_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLast_decl(s)
	}
}

func (s *Last_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLast_decl(s)
	}
}

func (s *Last_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLast_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Last_decl() (localctx ILast_declContext) {
	localctx = NewLast_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SynlongParserRULE_last_decl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(SynlongParserT__34)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(339)
		p.Match(SynlongParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(340)
		p.const_expr(0)
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
	Lhs() ILhsContext
	Expr() IExprContext
	State_machine() IState_machineContext
	Return_statement() IReturn_statementContext

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
	p.RuleIndex = SynlongParserRULE_equation
	return p
}

func InitEmptyEquationContext(p *EquationContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_equation
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

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

func (s *EquationContext) State_machine() IState_machineContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IState_machineContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IState_machineContext)
}

func (s *EquationContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (s *EquationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitEquation(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Equation() (localctx IEquationContext) {
	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SynlongParserRULE_equation)
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__26, SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(342)
			p.Lhs()
		}
		{
			p.SetState(343)
			p.Match(SynlongParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Expr()
		}


	case SynlongParserT__35:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(346)
			p.State_machine()
		}
		{
			p.SetState(347)
			p.Return_statement()
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


// ILhsContext is an interface to support dynamic dispatch.
type ILhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLhs_id() []ILhs_idContext
	Lhs_id(i int) ILhs_idContext

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
	p.RuleIndex = SynlongParserRULE_lhs
	return p
}

func InitEmptyLhsContext(p *LhsContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_lhs
}

func (*LhsContext) IsLhsContext() {}

func NewLhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LhsContext {
	var p = new(LhsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_lhs

	return p
}

func (s *LhsContext) GetParser() antlr.Parser { return s.parser }

func (s *LhsContext) AllLhs_id() []ILhs_idContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILhs_idContext); ok {
			len++
		}
	}

	tst := make([]ILhs_idContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILhs_idContext); ok {
			tst[i] = t.(ILhs_idContext)
			i++
		}
	}

	return tst
}

func (s *LhsContext) Lhs_id(i int) ILhs_idContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILhs_idContext); ok {
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

	return t.(ILhs_idContext)
}

func (s *LhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLhs(s)
	}
}

func (s *LhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLhs(s)
	}
}

func (s *LhsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLhs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Lhs() (localctx ILhsContext) {
	localctx = NewLhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SynlongParserRULE_lhs)
	var _la int

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__26:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(353)
			p.Lhs_id()
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserID {
			{
				p.SetState(354)
				p.Lhs_id()
			}


			p.SetState(359)
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


// ILhs_idContext is an interface to support dynamic dispatch.
type ILhs_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsLhs_idContext differentiates from other interfaces.
	IsLhs_idContext()
}

type Lhs_idContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLhs_idContext() *Lhs_idContext {
	var p = new(Lhs_idContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_lhs_id
	return p
}

func InitEmptyLhs_idContext(p *Lhs_idContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_lhs_id
}

func (*Lhs_idContext) IsLhs_idContext() {}

func NewLhs_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lhs_idContext {
	var p = new(Lhs_idContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_lhs_id

	return p
}

func (s *Lhs_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Lhs_idContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Lhs_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lhs_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Lhs_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLhs_id(s)
	}
}

func (s *Lhs_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLhs_id(s)
	}
}

func (s *Lhs_idContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLhs_id(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Lhs_id() (localctx ILhs_idContext) {
	localctx = NewLhs_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SynlongParserRULE_lhs_id)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(SynlongParserID)
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


// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Returns_var() IReturns_varContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) Returns_var() IReturns_varContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturns_varContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturns_varContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SynlongParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(SynlongParserT__28)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(365)
		p.Returns_var()
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


// IReturns_varContext is an interface to support dynamic dispatch.
type IReturns_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsReturns_varContext differentiates from other interfaces.
	IsReturns_varContext()
}

type Returns_varContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturns_varContext() *Returns_varContext {
	var p = new(Returns_varContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_returns_var
	return p
}

func InitEmptyReturns_varContext(p *Returns_varContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_returns_var
}

func (*Returns_varContext) IsReturns_varContext() {}

func NewReturns_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Returns_varContext {
	var p = new(Returns_varContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_returns_var

	return p
}

func (s *Returns_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Returns_varContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SynlongParserID)
}

func (s *Returns_varContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SynlongParserID, i)
}

func (s *Returns_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Returns_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Returns_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterReturns_var(s)
	}
}

func (s *Returns_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitReturns_var(s)
	}
}

func (s *Returns_varContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitReturns_var(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Returns_var() (localctx IReturns_varContext) {
	localctx = NewReturns_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SynlongParserRULE_returns_var)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == SynlongParserID {
		{
			p.SetState(367)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(370)
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


// IState_machineContext is an interface to support dynamic dispatch.
type IState_machineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	AllState_decl() []IState_declContext
	State_decl(i int) IState_declContext

	// IsState_machineContext differentiates from other interfaces.
	IsState_machineContext()
}

type State_machineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyState_machineContext() *State_machineContext {
	var p = new(State_machineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_state_machine
	return p
}

func InitEmptyState_machineContext(p *State_machineContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_state_machine
}

func (*State_machineContext) IsState_machineContext() {}

func NewState_machineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *State_machineContext {
	var p = new(State_machineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_state_machine

	return p
}

func (s *State_machineContext) GetParser() antlr.Parser { return s.parser }

func (s *State_machineContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *State_machineContext) AllState_decl() []IState_declContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IState_declContext); ok {
			len++
		}
	}

	tst := make([]IState_declContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IState_declContext); ok {
			tst[i] = t.(IState_declContext)
			i++
		}
	}

	return tst
}

func (s *State_machineContext) State_decl(i int) IState_declContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IState_declContext); ok {
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

	return t.(IState_declContext)
}

func (s *State_machineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_machineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *State_machineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterState_machine(s)
	}
}

func (s *State_machineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitState_machine(s)
	}
}

func (s *State_machineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitState_machine(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) State_machine() (localctx IState_machineContext) {
	localctx = NewState_machineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SynlongParserRULE_state_machine)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Match(SynlongParserT__35)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserID {
		{
			p.SetState(373)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 962072674304) != 0) {
		{
			p.SetState(376)
			p.State_decl()
		}


		p.SetState(379)
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


// IState_declContext is an interface to support dynamic dispatch.
type IState_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Data_def() IData_defContext
	AllTransition() []ITransitionContext
	Transition(i int) ITransitionContext

	// IsState_declContext differentiates from other interfaces.
	IsState_declContext()
}

type State_declContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyState_declContext() *State_declContext {
	var p = new(State_declContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_state_decl
	return p
}

func InitEmptyState_declContext(p *State_declContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_state_decl
}

func (*State_declContext) IsState_declContext() {}

func NewState_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *State_declContext {
	var p = new(State_declContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_state_decl

	return p
}

func (s *State_declContext) GetParser() antlr.Parser { return s.parser }

func (s *State_declContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *State_declContext) Data_def() IData_defContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_defContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_defContext)
}

func (s *State_declContext) AllTransition() []ITransitionContext {
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

func (s *State_declContext) Transition(i int) ITransitionContext {
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

func (s *State_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *State_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterState_decl(s)
	}
}

func (s *State_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitState_decl(s)
	}
}

func (s *State_declContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitState_decl(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) State_decl() (localctx IState_declContext) {
	localctx = NewState_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SynlongParserRULE_state_decl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__36 {
		{
			p.SetState(381)
			p.Match(SynlongParserT__36)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__37 {
		{
			p.SetState(384)
			p.Match(SynlongParserT__37)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	}
	{
		p.SetState(387)
		p.Match(SynlongParserT__38)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(388)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__39 {
		{
			p.SetState(389)
			p.Match(SynlongParserT__39)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Transition()
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__1 {
			{
				p.SetState(391)
				p.Match(SynlongParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(392)
				p.Transition()
			}


			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(400)
		p.Data_def()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == SynlongParserT__40 {
		{
			p.SetState(401)
			p.Match(SynlongParserT__40)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(402)
			p.Transition()
		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == SynlongParserT__1 {
			{
				p.SetState(403)
				p.Match(SynlongParserT__1)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(404)
				p.Transition()
			}


			p.SetState(409)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
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


// IData_defContext is an interface to support dynamic dispatch.
type IData_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEquation() []IEquationContext
	Equation(i int) IEquationContext
	Local_block() ILocal_blockContext

	// IsData_defContext differentiates from other interfaces.
	IsData_defContext()
}

type Data_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_defContext() *Data_defContext {
	var p = new(Data_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_data_def
	return p
}

func InitEmptyData_defContext(p *Data_defContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_data_def
}

func (*Data_defContext) IsData_defContext() {}

func NewData_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_defContext {
	var p = new(Data_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_data_def

	return p
}

func (s *Data_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_defContext) AllEquation() []IEquationContext {
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

func (s *Data_defContext) Equation(i int) IEquationContext {
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

func (s *Data_defContext) Local_block() ILocal_blockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocal_blockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocal_blockContext)
}

func (s *Data_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Data_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterData_def(s)
	}
}

func (s *Data_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitData_def(s)
	}
}

func (s *Data_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitData_def(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Data_def() (localctx IData_defContext) {
	localctx = NewData_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SynlongParserRULE_data_def)
	var _la int

	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__26, SynlongParserT__35, SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Equation()
		}
		{
			p.SetState(413)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__29, SynlongParserT__31:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == SynlongParserT__31 {
			{
				p.SetState(415)
				p.Local_block()
			}

		}
		{
			p.SetState(418)
			p.Match(SynlongParserT__29)
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


		for _la == SynlongParserT__26 || _la == SynlongParserT__35 || _la == SynlongParserID {
			{
				p.SetState(419)
				p.Equation()
			}
			{
				p.SetState(420)
				p.Match(SynlongParserT__1)
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
		{
			p.SetState(427)
			p.Match(SynlongParserT__30)
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


// ITransitionContext is an interface to support dynamic dispatch.
type ITransitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	ID() antlr.TerminalNode

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
	p.RuleIndex = SynlongParserRULE_transition
	return p
}

func InitEmptyTransitionContext(p *TransitionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_transition
}

func (*TransitionContext) IsTransitionContext() {}

func NewTransitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransitionContext {
	var p = new(TransitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_transition

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

func (s *TransitionContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *TransitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TransitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterTransition(s)
	}
}

func (s *TransitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitTransition(s)
	}
}

func (s *TransitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitTransition(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Transition() (localctx ITransitionContext) {
	localctx = NewTransitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SynlongParserRULE_transition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(SynlongParserT__41)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Expr()
	}
	{
		p.SetState(432)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SynlongParserT__42 || _la == SynlongParserT__43) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(433)
		p.Match(SynlongParserID)
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


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_expr() ISimple_exprContext
	ID() antlr.TerminalNode
	Tempo_expr() ITempo_exprContext
	Bool_expr() IBool_exprContext
	Array_expr() IArray_exprContext
	Struct_expr() IStruct_exprContext
	Mixed_constructor() IMixed_constructorContext
	Switch_expr() ISwitch_exprContext
	Apply_expr() IApply_exprContext

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
	p.RuleIndex = SynlongParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *ExprContext) Tempo_expr() ITempo_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITempo_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITempo_exprContext)
}

func (s *ExprContext) Bool_expr() IBool_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_exprContext)
}

func (s *ExprContext) Array_expr() IArray_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_exprContext)
}

func (s *ExprContext) Struct_expr() IStruct_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_exprContext)
}

func (s *ExprContext) Mixed_constructor() IMixed_constructorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMixed_constructorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMixed_constructorContext)
}

func (s *ExprContext) Switch_expr() ISwitch_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_exprContext)
}

func (s *ExprContext) Apply_expr() IApply_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApply_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApply_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SynlongParserRULE_expr)
	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.simple_expr(0)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.Match(SynlongParserT__34)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(437)
			p.Match(SynlongParserT__44)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(439)
			p.Tempo_expr()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(440)
			p.Bool_expr()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)
			p.Array_expr()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(442)
			p.Struct_expr()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(443)
			p.Mixed_constructor()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(444)
			p.Switch_expr()
		}


	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(445)
			p.Apply_expr()
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


// ISimple_exprContext is an interface to support dynamic dispatch.
type ISimple_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Atom() IAtomContext
	Unary_arith_op() IUnary_arith_opContext
	AllSimple_expr() []ISimple_exprContext
	Simple_expr(i int) ISimple_exprContext
	Type_expr() IType_exprContext
	Bin_arith_op() IBin_arith_opContext
	Bin_bool_op() IBin_bool_opContext
	Bin_relation_op() IBin_relation_opContext
	Const_expr() IConst_exprContext

	// IsSimple_exprContext differentiates from other interfaces.
	IsSimple_exprContext()
}

type Simple_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_exprContext() *Simple_exprContext {
	var p = new(Simple_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_simple_expr
	return p
}

func InitEmptySimple_exprContext(p *Simple_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_simple_expr
}

func (*Simple_exprContext) IsSimple_exprContext() {}

func NewSimple_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_exprContext {
	var p = new(Simple_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_simple_expr

	return p
}

func (s *Simple_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Simple_exprContext) Atom() IAtomContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Simple_exprContext) Unary_arith_op() IUnary_arith_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnary_arith_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnary_arith_opContext)
}

func (s *Simple_exprContext) AllSimple_expr() []ISimple_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_exprContext); ok {
			len++
		}
	}

	tst := make([]ISimple_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_exprContext); ok {
			tst[i] = t.(ISimple_exprContext)
			i++
		}
	}

	return tst
}

func (s *Simple_exprContext) Simple_expr(i int) ISimple_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
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

	return t.(ISimple_exprContext)
}

func (s *Simple_exprContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Simple_exprContext) Bin_arith_op() IBin_arith_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_arith_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_arith_opContext)
}

func (s *Simple_exprContext) Bin_bool_op() IBin_bool_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_bool_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_bool_opContext)
}

func (s *Simple_exprContext) Bin_relation_op() IBin_relation_opContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBin_relation_opContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBin_relation_opContext)
}

func (s *Simple_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Simple_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Simple_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterSimple_expr(s)
	}
}

func (s *Simple_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitSimple_expr(s)
	}
}

func (s *Simple_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitSimple_expr(s)

	default:
		return t.VisitChildren(s)
	}
}





func (p *SynlongParser) Simple_expr() (localctx ISimple_exprContext) {
	return p.simple_expr(0)
}

func (p *SynlongParser) simple_expr(_p int) (localctx ISimple_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSimple_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISimple_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 68
	p.EnterRecursionRule(localctx, 68, SynlongParserRULE_simple_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserID:
		{
			p.SetState(449)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__60, SynlongParserT__61, SynlongParserCHAR, SynlongParserINTEGER, SynlongParserUINT, SynlongParserFLOAT, SynlongParserREAL, SynlongParserUSHORT, SynlongParserSHORT:
		{
			p.SetState(450)
			p.Atom()
		}


	case SynlongParserT__33, SynlongParserT__59, SynlongParserT__98:
		{
			p.SetState(451)
			p.Unary_arith_op()
		}
		{
			p.SetState(452)
			p.simple_expr(5)
		}


	case SynlongParserT__26:
		{
			p.SetState(454)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(455)
			p.type_expr(0)
		}
		{
			p.SetState(456)
			p.simple_expr(0)
		}
		{
			p.SetState(457)
			p.Match(SynlongParserT__27)
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
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(481)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSimple_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_simple_expr)
				p.SetState(461)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(462)
					p.Bin_arith_op()
				}
				{
					p.SetState(463)
					p.simple_expr(5)
				}


			case 2:
				localctx = NewSimple_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_simple_expr)
				p.SetState(465)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(466)
					p.Bin_bool_op()
				}
				{
					p.SetState(467)
					p.simple_expr(4)
				}


			case 3:
				localctx = NewSimple_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_simple_expr)
				p.SetState(469)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(470)
					p.Bin_relation_op()
				}
				{
					p.SetState(471)
					p.simple_expr(3)
				}


			case 4:
				localctx = NewSimple_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_simple_expr)
				p.SetState(473)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(474)
					p.Match(SynlongParserT__22)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(475)
					p.const_expr(0)
				}
				{
					p.SetState(476)
					p.Match(SynlongParserT__23)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case 5:
				localctx = NewSimple_exprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SynlongParserRULE_simple_expr)
				p.SetState(478)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(479)
					p.Match(SynlongParserT__45)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(480)
					p.Match(SynlongParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
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


// ITempo_exprContext is an interface to support dynamic dispatch.
type ITempo_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimple_expr() []ISimple_exprContext
	Simple_expr(i int) ISimple_exprContext
	Const_expr() IConst_exprContext
	Clock_expr() IClock_exprContext
	ID() antlr.TerminalNode

	// IsTempo_exprContext differentiates from other interfaces.
	IsTempo_exprContext()
}

type Tempo_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempo_exprContext() *Tempo_exprContext {
	var p = new(Tempo_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_tempo_expr
	return p
}

func InitEmptyTempo_exprContext(p *Tempo_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_tempo_expr
}

func (*Tempo_exprContext) IsTempo_exprContext() {}

func NewTempo_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tempo_exprContext {
	var p = new(Tempo_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_tempo_expr

	return p
}

func (s *Tempo_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Tempo_exprContext) AllSimple_expr() []ISimple_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_exprContext); ok {
			len++
		}
	}

	tst := make([]ISimple_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_exprContext); ok {
			tst[i] = t.(ISimple_exprContext)
			i++
		}
	}

	return tst
}

func (s *Tempo_exprContext) Simple_expr(i int) ISimple_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
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

	return t.(ISimple_exprContext)
}

func (s *Tempo_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Tempo_exprContext) Clock_expr() IClock_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClock_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClock_exprContext)
}

func (s *Tempo_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Tempo_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tempo_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Tempo_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterTempo_expr(s)
	}
}

func (s *Tempo_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitTempo_expr(s)
	}
}

func (s *Tempo_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitTempo_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Tempo_expr() (localctx ITempo_exprContext) {
	localctx = NewTempo_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SynlongParserRULE_tempo_expr)
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(486)
			p.Match(SynlongParserT__46)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(487)
			p.simple_expr(0)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(488)
			p.simple_expr(0)
		}
		{
			p.SetState(489)
			p.Match(SynlongParserT__47)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(490)
			p.simple_expr(0)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(492)
			p.Match(SynlongParserT__48)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(494)
			p.simple_expr(0)
		}
		{
			p.SetState(495)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(496)
			p.const_expr(0)
		}
		{
			p.SetState(497)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(498)
			p.simple_expr(0)
		}
		{
			p.SetState(499)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(501)
			p.simple_expr(0)
		}
		{
			p.SetState(502)
			p.Match(SynlongParserT__48)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(503)
			p.simple_expr(0)
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(505)
			p.simple_expr(0)
		}
		{
			p.SetState(506)
			p.Match(SynlongParserT__32)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(507)
			p.Clock_expr()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(509)
			p.Match(SynlongParserT__49)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(510)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(511)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(512)
			p.simple_expr(0)
		}
		{
			p.SetState(513)
			p.Match(SynlongParserT__18)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(514)
			p.simple_expr(0)
		}
		{
			p.SetState(515)
			p.Match(SynlongParserT__27)
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


// IBool_exprContext is an interface to support dynamic dispatch.
type IBool_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	List() IListContext

	// IsBool_exprContext differentiates from other interfaces.
	IsBool_exprContext()
}

type Bool_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_exprContext() *Bool_exprContext {
	var p = new(Bool_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bool_expr
	return p
}

func InitEmptyBool_exprContext(p *Bool_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bool_expr
}

func (*Bool_exprContext) IsBool_exprContext() {}

func NewBool_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_exprContext {
	var p = new(Bool_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_bool_expr

	return p
}

func (s *Bool_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_exprContext) List() IListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *Bool_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Bool_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterBool_expr(s)
	}
}

func (s *Bool_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitBool_expr(s)
	}
}

func (s *Bool_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitBool_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Bool_expr() (localctx IBool_exprContext) {
	localctx = NewBool_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SynlongParserRULE_bool_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(519)
		p.Match(SynlongParserT__50)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(520)
		p.Match(SynlongParserT__26)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(521)
		p.List()
	}
	{
		p.SetState(522)
		p.Match(SynlongParserT__27)
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


// IArray_exprContext is an interface to support dynamic dispatch.
type IArray_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimple_expr() []ISimple_exprContext
	Simple_expr(i int) ISimple_exprContext
	AllINTEGER() []antlr.TerminalNode
	INTEGER(i int) antlr.TerminalNode
	AllIndex() []IIndexContext
	Index(i int) IIndexContext
	Const_expr() IConst_exprContext
	List() IListContext

	// IsArray_exprContext differentiates from other interfaces.
	IsArray_exprContext()
}

type Array_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_exprContext() *Array_exprContext {
	var p = new(Array_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_array_expr
	return p
}

func InitEmptyArray_exprContext(p *Array_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_array_expr
}

func (*Array_exprContext) IsArray_exprContext() {}

func NewArray_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_exprContext {
	var p = new(Array_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_array_expr

	return p
}

func (s *Array_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_exprContext) AllSimple_expr() []ISimple_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_exprContext); ok {
			len++
		}
	}

	tst := make([]ISimple_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_exprContext); ok {
			tst[i] = t.(ISimple_exprContext)
			i++
		}
	}

	return tst
}

func (s *Array_exprContext) Simple_expr(i int) ISimple_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
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

	return t.(ISimple_exprContext)
}

func (s *Array_exprContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SynlongParserINTEGER)
}

func (s *Array_exprContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SynlongParserINTEGER, i)
}

func (s *Array_exprContext) AllIndex() []IIndexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexContext); ok {
			len++
		}
	}

	tst := make([]IIndexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexContext); ok {
			tst[i] = t.(IIndexContext)
			i++
		}
	}

	return tst
}

func (s *Array_exprContext) Index(i int) IIndexContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
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

	return t.(IIndexContext)
}

func (s *Array_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Array_exprContext) List() IListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *Array_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Array_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterArray_expr(s)
	}
}

func (s *Array_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitArray_expr(s)
	}
}

func (s *Array_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitArray_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Array_expr() (localctx IArray_exprContext) {
	localctx = NewArray_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SynlongParserRULE_array_expr)
	var _la int

	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(524)
			p.simple_expr(0)
		}
		{
			p.SetState(525)
			p.Match(SynlongParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(526)
			p.Match(SynlongParserINTEGER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(527)
			p.Match(SynlongParserT__51)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Match(SynlongParserINTEGER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(529)
			p.Match(SynlongParserT__23)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(531)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(532)
			p.simple_expr(0)
		}
		{
			p.SetState(533)
			p.Match(SynlongParserT__45)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(535)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == SynlongParserT__22 {
			{
				p.SetState(534)
				p.Index()
			}


			p.SetState(537)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(539)
			p.Match(SynlongParserT__52)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(540)
			p.simple_expr(0)
		}
		{
			p.SetState(541)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(543)
			p.simple_expr(0)
		}
		{
			p.SetState(544)
			p.Match(SynlongParserT__19)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(545)
			p.const_expr(0)
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(547)
			p.Match(SynlongParserT__22)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(548)
			p.List()
		}
		{
			p.SetState(549)
			p.Match(SynlongParserT__23)
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


// IStruct_exprContext is an interface to support dynamic dispatch.
type IStruct_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabel_expr() []ILabel_exprContext
	Label_expr(i int) ILabel_exprContext

	// IsStruct_exprContext differentiates from other interfaces.
	IsStruct_exprContext()
}

type Struct_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_exprContext() *Struct_exprContext {
	var p = new(Struct_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_struct_expr
	return p
}

func InitEmptyStruct_exprContext(p *Struct_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_struct_expr
}

func (*Struct_exprContext) IsStruct_exprContext() {}

func NewStruct_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_exprContext {
	var p = new(Struct_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_struct_expr

	return p
}

func (s *Struct_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_exprContext) AllLabel_expr() []ILabel_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabel_exprContext); ok {
			len++
		}
	}

	tst := make([]ILabel_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabel_exprContext); ok {
			tst[i] = t.(ILabel_exprContext)
			i++
		}
	}

	return tst
}

func (s *Struct_exprContext) Label_expr(i int) ILabel_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabel_exprContext); ok {
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

	return t.(ILabel_exprContext)
}

func (s *Struct_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Struct_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterStruct_expr(s)
	}
}

func (s *Struct_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitStruct_expr(s)
	}
}

func (s *Struct_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitStruct_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Struct_expr() (localctx IStruct_exprContext) {
	localctx = NewStruct_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SynlongParserRULE_struct_expr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(553)
		p.Match(SynlongParserT__4)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(554)
		p.Label_expr()
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == SynlongParserID {
		{
			p.SetState(555)
			p.Label_expr()
		}


		p.SetState(560)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(561)
		p.Match(SynlongParserT__5)
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


// ILabel_exprContext is an interface to support dynamic dispatch.
type ILabel_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Simple_expr() ISimple_exprContext

	// IsLabel_exprContext differentiates from other interfaces.
	IsLabel_exprContext()
}

type Label_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabel_exprContext() *Label_exprContext {
	var p = new(Label_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_label_expr
	return p
}

func InitEmptyLabel_exprContext(p *Label_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_label_expr
}

func (*Label_exprContext) IsLabel_exprContext() {}

func NewLabel_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Label_exprContext {
	var p = new(Label_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_label_expr

	return p
}

func (s *Label_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Label_exprContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Label_exprContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *Label_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Label_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Label_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLabel_expr(s)
	}
}

func (s *Label_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLabel_expr(s)
	}
}

func (s *Label_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLabel_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Label_expr() (localctx ILabel_exprContext) {
	localctx = NewLabel_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SynlongParserRULE_label_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(564)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(565)
		p.simple_expr(0)
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


// IMixed_constructorContext is an interface to support dynamic dispatch.
type IMixed_constructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Simple_expr() ISimple_exprContext
	AllLabel_or_index() []ILabel_or_indexContext
	Label_or_index(i int) ILabel_or_indexContext

	// IsMixed_constructorContext differentiates from other interfaces.
	IsMixed_constructorContext()
}

type Mixed_constructorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixed_constructorContext() *Mixed_constructorContext {
	var p = new(Mixed_constructorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_mixed_constructor
	return p
}

func InitEmptyMixed_constructorContext(p *Mixed_constructorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_mixed_constructor
}

func (*Mixed_constructorContext) IsMixed_constructorContext() {}

func NewMixed_constructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mixed_constructorContext {
	var p = new(Mixed_constructorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_mixed_constructor

	return p
}

func (s *Mixed_constructorContext) GetParser() antlr.Parser { return s.parser }

func (s *Mixed_constructorContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Mixed_constructorContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *Mixed_constructorContext) AllLabel_or_index() []ILabel_or_indexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabel_or_indexContext); ok {
			len++
		}
	}

	tst := make([]ILabel_or_indexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabel_or_indexContext); ok {
			tst[i] = t.(ILabel_or_indexContext)
			i++
		}
	}

	return tst
}

func (s *Mixed_constructorContext) Label_or_index(i int) ILabel_or_indexContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabel_or_indexContext); ok {
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

	return t.(ILabel_or_indexContext)
}

func (s *Mixed_constructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mixed_constructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Mixed_constructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterMixed_constructor(s)
	}
}

func (s *Mixed_constructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitMixed_constructor(s)
	}
}

func (s *Mixed_constructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitMixed_constructor(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Mixed_constructor() (localctx IMixed_constructorContext) {
	localctx = NewMixed_constructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SynlongParserRULE_mixed_constructor)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(567)
		p.Match(SynlongParserT__26)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(568)
		p.Match(SynlongParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(569)
		p.Match(SynlongParserT__53)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(571)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == SynlongParserT__22 || _la == SynlongParserT__45 {
		{
			p.SetState(570)
			p.Label_or_index()
		}


		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(575)
		p.Match(SynlongParserT__2)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(576)
		p.simple_expr(0)
	}
	{
		p.SetState(577)
		p.Match(SynlongParserT__27)
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


// ILabel_or_indexContext is an interface to support dynamic dispatch.
type ILabel_or_indexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Index() IIndexContext

	// IsLabel_or_indexContext differentiates from other interfaces.
	IsLabel_or_indexContext()
}

type Label_or_indexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabel_or_indexContext() *Label_or_indexContext {
	var p = new(Label_or_indexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_label_or_index
	return p
}

func InitEmptyLabel_or_indexContext(p *Label_or_indexContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_label_or_index
}

func (*Label_or_indexContext) IsLabel_or_indexContext() {}

func NewLabel_or_indexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Label_or_indexContext {
	var p = new(Label_or_indexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_label_or_index

	return p
}

func (s *Label_or_indexContext) GetParser() antlr.Parser { return s.parser }

func (s *Label_or_indexContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Label_or_indexContext) Index() IIndexContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *Label_or_indexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Label_or_indexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Label_or_indexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterLabel_or_index(s)
	}
}

func (s *Label_or_indexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitLabel_or_index(s)
	}
}

func (s *Label_or_indexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitLabel_or_index(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Label_or_index() (localctx ILabel_or_indexContext) {
	localctx = NewLabel_or_indexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SynlongParserRULE_label_or_index)
	p.SetState(582)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__45:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(579)
			p.Match(SynlongParserT__45)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(580)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(581)
			p.Index()
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


// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_expr() ISimple_exprContext

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SynlongParserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.Match(SynlongParserT__22)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(585)
		p.simple_expr(0)
	}
	{
		p.SetState(586)
		p.Match(SynlongParserT__23)
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


// ISwitch_exprContext is an interface to support dynamic dispatch.
type ISwitch_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimple_expr() []ISimple_exprContext
	Simple_expr(i int) ISimple_exprContext
	AllCase_expr() []ICase_exprContext
	Case_expr(i int) ICase_exprContext

	// IsSwitch_exprContext differentiates from other interfaces.
	IsSwitch_exprContext()
}

type Switch_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_exprContext() *Switch_exprContext {
	var p = new(Switch_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_switch_expr
	return p
}

func InitEmptySwitch_exprContext(p *Switch_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_switch_expr
}

func (*Switch_exprContext) IsSwitch_exprContext() {}

func NewSwitch_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_exprContext {
	var p = new(Switch_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_switch_expr

	return p
}

func (s *Switch_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_exprContext) AllSimple_expr() []ISimple_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_exprContext); ok {
			len++
		}
	}

	tst := make([]ISimple_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_exprContext); ok {
			tst[i] = t.(ISimple_exprContext)
			i++
		}
	}

	return tst
}

func (s *Switch_exprContext) Simple_expr(i int) ISimple_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
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

	return t.(ISimple_exprContext)
}

func (s *Switch_exprContext) AllCase_expr() []ICase_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICase_exprContext); ok {
			len++
		}
	}

	tst := make([]ICase_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICase_exprContext); ok {
			tst[i] = t.(ICase_exprContext)
			i++
		}
	}

	return tst
}

func (s *Switch_exprContext) Case_expr(i int) ICase_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_exprContext); ok {
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

	return t.(ICase_exprContext)
}

func (s *Switch_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Switch_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterSwitch_expr(s)
	}
}

func (s *Switch_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitSwitch_expr(s)
	}
}

func (s *Switch_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitSwitch_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Switch_expr() (localctx ISwitch_exprContext) {
	localctx = NewSwitch_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SynlongParserRULE_switch_expr)
	var _la int

	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserT__41:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(588)
			p.Match(SynlongParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(589)
			p.simple_expr(0)
		}
		{
			p.SetState(590)
			p.Match(SynlongParserT__54)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(591)
			p.simple_expr(0)
		}
		{
			p.SetState(592)
			p.Match(SynlongParserT__55)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(593)
			p.simple_expr(0)
		}


	case SynlongParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(596)
			p.Match(SynlongParserT__56)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(597)
			p.simple_expr(0)
		}
		{
			p.SetState(598)
			p.Match(SynlongParserT__57)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(600)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == SynlongParserT__58 {
			{
				p.SetState(599)
				p.Case_expr()
			}


			p.SetState(602)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(604)
			p.Match(SynlongParserT__27)
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


// ICase_exprContext is an interface to support dynamic dispatch.
type ICase_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pattern() IPatternContext
	Simple_expr() ISimple_exprContext

	// IsCase_exprContext differentiates from other interfaces.
	IsCase_exprContext()
}

type Case_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_exprContext() *Case_exprContext {
	var p = new(Case_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_case_expr
	return p
}

func InitEmptyCase_exprContext(p *Case_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_case_expr
}

func (*Case_exprContext) IsCase_exprContext() {}

func NewCase_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_exprContext {
	var p = new(Case_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_case_expr

	return p
}

func (s *Case_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_exprContext) Pattern() IPatternContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatternContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *Case_exprContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *Case_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Case_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterCase_expr(s)
	}
}

func (s *Case_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitCase_expr(s)
	}
}

func (s *Case_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitCase_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Case_expr() (localctx ICase_exprContext) {
	localctx = NewCase_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SynlongParserRULE_case_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		p.Match(SynlongParserT__58)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(609)
		p.Pattern()
	}
	{
		p.SetState(610)
		p.Match(SynlongParserT__20)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(611)
		p.simple_expr(0)
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
	CHAR() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

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
	p.RuleIndex = SynlongParserRULE_pattern
	return p
}

func InitEmptyPatternContext(p *PatternContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_pattern
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *PatternContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SynlongParserCHAR, 0)
}

func (s *PatternContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SynlongParserINTEGER, 0)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (s *PatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitPattern(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Pattern() (localctx IPatternContext) {
	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SynlongParserRULE_pattern)
	var _la int

	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(613)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(614)
			p.Match(SynlongParserCHAR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__59, SynlongParserINTEGER:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(616)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == SynlongParserT__59 {
			{
				p.SetState(615)
				p.Match(SynlongParserT__59)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}
		{
			p.SetState(618)
			p.Match(SynlongParserINTEGER)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__60:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(619)
			p.Match(SynlongParserT__60)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__61:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(620)
			p.Match(SynlongParserT__61)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__62:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(621)
			p.Match(SynlongParserT__62)
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


// IApply_exprContext is an interface to support dynamic dispatch.
type IApply_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prefix_operator() IPrefix_operatorContext
	AllList() []IListContext
	List(i int) IListContext
	Iterator() IIteratorContext
	Const_expr() IConst_exprContext
	Simple_expr() ISimple_exprContext

	// IsApply_exprContext differentiates from other interfaces.
	IsApply_exprContext()
}

type Apply_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApply_exprContext() *Apply_exprContext {
	var p = new(Apply_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_apply_expr
	return p
}

func InitEmptyApply_exprContext(p *Apply_exprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_apply_expr
}

func (*Apply_exprContext) IsApply_exprContext() {}

func NewApply_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Apply_exprContext {
	var p = new(Apply_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_apply_expr

	return p
}

func (s *Apply_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Apply_exprContext) Prefix_operator() IPrefix_operatorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_operatorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_operatorContext)
}

func (s *Apply_exprContext) AllList() []IListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListContext); ok {
			len++
		}
	}

	tst := make([]IListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListContext); ok {
			tst[i] = t.(IListContext)
			i++
		}
	}

	return tst
}

func (s *Apply_exprContext) List(i int) IListContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
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

	return t.(IListContext)
}

func (s *Apply_exprContext) Iterator() IIteratorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIteratorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIteratorContext)
}

func (s *Apply_exprContext) Const_expr() IConst_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConst_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConst_exprContext)
}

func (s *Apply_exprContext) Simple_expr() ISimple_exprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_exprContext)
}

func (s *Apply_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Apply_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Apply_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterApply_expr(s)
	}
}

func (s *Apply_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitApply_expr(s)
	}
}

func (s *Apply_exprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitApply_expr(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Apply_expr() (localctx IApply_exprContext) {
	localctx = NewApply_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SynlongParserRULE_apply_expr)
	p.SetState(705)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(624)
			p.Prefix_operator()
		}
		{
			p.SetState(625)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(626)
			p.List()
		}
		{
			p.SetState(627)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(629)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Iterator()
		}
		{
			p.SetState(631)
			p.Match(SynlongParserT__63)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(632)
			p.Prefix_operator()
		}
		{
			p.SetState(633)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(634)
			p.const_expr(0)
		}
		{
			p.SetState(635)
			p.Match(SynlongParserT__64)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(636)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(637)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(638)
			p.List()
		}
		{
			p.SetState(639)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(641)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(642)
			p.Match(SynlongParserT__65)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(643)
			p.Match(SynlongParserT__63)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(644)
			p.Prefix_operator()
		}
		{
			p.SetState(645)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(646)
			p.const_expr(0)
		}
		{
			p.SetState(647)
			p.Match(SynlongParserT__64)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(648)
			p.Match(SynlongParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(649)
			p.simple_expr(0)
		}
		{
			p.SetState(650)
			p.Match(SynlongParserT__52)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(651)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(652)
			p.List()
		}
		{
			p.SetState(653)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(654)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(655)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(656)
			p.List()
		}
		{
			p.SetState(657)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(659)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(660)
			p.Match(SynlongParserT__66)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(661)
			p.Match(SynlongParserT__63)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(662)
			p.Prefix_operator()
		}
		{
			p.SetState(663)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(664)
			p.const_expr(0)
		}
		{
			p.SetState(665)
			p.Match(SynlongParserT__64)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(666)
			p.Match(SynlongParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(667)
			p.simple_expr(0)
		}
		{
			p.SetState(668)
			p.Match(SynlongParserT__52)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(669)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(670)
			p.List()
		}
		{
			p.SetState(671)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(672)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(673)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(674)
			p.List()
		}
		{
			p.SetState(675)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(677)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(678)
			p.Match(SynlongParserT__67)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(679)
			p.Match(SynlongParserT__63)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(680)
			p.Prefix_operator()
		}
		{
			p.SetState(681)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(682)
			p.const_expr(0)
		}
		{
			p.SetState(683)
			p.Match(SynlongParserT__64)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(684)
			p.Match(SynlongParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(685)
			p.simple_expr(0)
		}
		{
			p.SetState(686)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(687)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(688)
			p.List()
		}
		{
			p.SetState(689)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(691)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(692)
			p.Match(SynlongParserT__68)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(693)
			p.Match(SynlongParserT__63)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(694)
			p.Prefix_operator()
		}
		{
			p.SetState(695)
			p.Match(SynlongParserT__1)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(696)
			p.const_expr(0)
		}
		{
			p.SetState(697)
			p.Match(SynlongParserT__64)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(698)
			p.Match(SynlongParserT__41)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(699)
			p.simple_expr(0)
		}
		{
			p.SetState(700)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(701)
			p.Match(SynlongParserT__26)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(702)
			p.List()
		}
		{
			p.SetState(703)
			p.Match(SynlongParserT__27)
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


// IPrefix_operatorContext is an interface to support dynamic dispatch.
type IPrefix_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Prefix_unary_operator() IPrefix_unary_operatorContext
	Prefix_binary_operator() IPrefix_binary_operatorContext

	// IsPrefix_operatorContext differentiates from other interfaces.
	IsPrefix_operatorContext()
}

type Prefix_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_operatorContext() *Prefix_operatorContext {
	var p = new(Prefix_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_operator
	return p
}

func InitEmptyPrefix_operatorContext(p *Prefix_operatorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_operator
}

func (*Prefix_operatorContext) IsPrefix_operatorContext() {}

func NewPrefix_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_operatorContext {
	var p = new(Prefix_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_prefix_operator

	return p
}

func (s *Prefix_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Prefix_operatorContext) ID() antlr.TerminalNode {
	return s.GetToken(SynlongParserID, 0)
}

func (s *Prefix_operatorContext) Prefix_unary_operator() IPrefix_unary_operatorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_unary_operatorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_unary_operatorContext)
}

func (s *Prefix_operatorContext) Prefix_binary_operator() IPrefix_binary_operatorContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefix_binary_operatorContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefix_binary_operatorContext)
}

func (s *Prefix_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Prefix_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterPrefix_operator(s)
	}
}

func (s *Prefix_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitPrefix_operator(s)
	}
}

func (s *Prefix_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitPrefix_operator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Prefix_operator() (localctx IPrefix_operatorContext) {
	localctx = NewPrefix_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SynlongParserRULE_prefix_operator)
	p.SetState(716)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SynlongParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(707)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__71, SynlongParserT__72, SynlongParserT__73, SynlongParserT__74, SynlongParserT__75, SynlongParserT__76, SynlongParserT__77:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(708)
			p.Prefix_unary_operator()
		}


	case SynlongParserT__78, SynlongParserT__79, SynlongParserT__80, SynlongParserT__81, SynlongParserT__82, SynlongParserT__83, SynlongParserT__84, SynlongParserT__85, SynlongParserT__86, SynlongParserT__87, SynlongParserT__88, SynlongParserT__89, SynlongParserT__90, SynlongParserT__91, SynlongParserT__92:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(709)
			p.Prefix_binary_operator()
		}


	case SynlongParserT__69:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(710)
			p.Match(SynlongParserT__69)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(711)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(712)
			p.Match(SynlongParserT__27)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case SynlongParserT__70:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(713)
			p.Match(SynlongParserT__70)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(714)
			p.Match(SynlongParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(715)
			p.Match(SynlongParserT__27)
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


// IPrefix_unary_operatorContext is an interface to support dynamic dispatch.
type IPrefix_unary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrefix_unary_operatorContext differentiates from other interfaces.
	IsPrefix_unary_operatorContext()
}

type Prefix_unary_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_unary_operatorContext() *Prefix_unary_operatorContext {
	var p = new(Prefix_unary_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_unary_operator
	return p
}

func InitEmptyPrefix_unary_operatorContext(p *Prefix_unary_operatorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_unary_operator
}

func (*Prefix_unary_operatorContext) IsPrefix_unary_operatorContext() {}

func NewPrefix_unary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_unary_operatorContext {
	var p = new(Prefix_unary_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_prefix_unary_operator

	return p
}

func (s *Prefix_unary_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Prefix_unary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_unary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Prefix_unary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterPrefix_unary_operator(s)
	}
}

func (s *Prefix_unary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitPrefix_unary_operator(s)
	}
}

func (s *Prefix_unary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitPrefix_unary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Prefix_unary_operator() (localctx IPrefix_unary_operatorContext) {
	localctx = NewPrefix_unary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SynlongParserRULE_prefix_unary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(718)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 127) != 0)) {
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


// IPrefix_binary_operatorContext is an interface to support dynamic dispatch.
type IPrefix_binary_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrefix_binary_operatorContext differentiates from other interfaces.
	IsPrefix_binary_operatorContext()
}

type Prefix_binary_operatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefix_binary_operatorContext() *Prefix_binary_operatorContext {
	var p = new(Prefix_binary_operatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_binary_operator
	return p
}

func InitEmptyPrefix_binary_operatorContext(p *Prefix_binary_operatorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_prefix_binary_operator
}

func (*Prefix_binary_operatorContext) IsPrefix_binary_operatorContext() {}

func NewPrefix_binary_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prefix_binary_operatorContext {
	var p = new(Prefix_binary_operatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_prefix_binary_operator

	return p
}

func (s *Prefix_binary_operatorContext) GetParser() antlr.Parser { return s.parser }
func (s *Prefix_binary_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prefix_binary_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Prefix_binary_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterPrefix_binary_operator(s)
	}
}

func (s *Prefix_binary_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitPrefix_binary_operator(s)
	}
}

func (s *Prefix_binary_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitPrefix_binary_operator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Prefix_binary_operator() (localctx IPrefix_binary_operatorContext) {
	localctx = NewPrefix_binary_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SynlongParserRULE_prefix_binary_operator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(720)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 79)) & ^0x3f) == 0 && ((int64(1) << (_la - 79)) & 32767) != 0)) {
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


// IIteratorContext is an interface to support dynamic dispatch.
type IIteratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIteratorContext differentiates from other interfaces.
	IsIteratorContext()
}

type IteratorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIteratorContext() *IteratorContext {
	var p = new(IteratorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_iterator
	return p
}

func InitEmptyIteratorContext(p *IteratorContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_iterator
}

func (*IteratorContext) IsIteratorContext() {}

func NewIteratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IteratorContext {
	var p = new(IteratorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_iterator

	return p
}

func (s *IteratorContext) GetParser() antlr.Parser { return s.parser }
func (s *IteratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IteratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IteratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterIterator(s)
	}
}

func (s *IteratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitIterator(s)
	}
}

func (s *IteratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitIterator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Iterator() (localctx IIteratorContext) {
	localctx = NewIteratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SynlongParserRULE_iterator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(722)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 94)) & ^0x3f) == 0 && ((int64(1) << (_la - 94)) & 31) != 0)) {
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


// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimple_expr() []ISimple_exprContext
	Simple_expr(i int) ISimple_exprContext

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllSimple_expr() []ISimple_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_exprContext); ok {
			len++
		}
	}

	tst := make([]ISimple_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_exprContext); ok {
			tst[i] = t.(ISimple_exprContext)
			i++
		}
	}

	return tst
}

func (s *ListContext) Simple_expr(i int) ISimple_exprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_exprContext); ok {
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

	return t.(ISimple_exprContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SynlongParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(731)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8070450549562015744) != 0) || ((int64((_la - 99)) & ^0x3f) == 0 && ((int64(1) << (_la - 99)) & 4177921) != 0) {
		{
			p.SetState(724)
			p.simple_expr(0)
		}
		p.SetState(728)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 8070450549562015744) != 0) || ((int64((_la - 99)) & ^0x3f) == 0 && ((int64(1) << (_la - 99)) & 4177921) != 0) {
			{
				p.SetState(725)
				p.simple_expr(0)
			}


			p.SetState(730)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
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


// IUnary_arith_opContext is an interface to support dynamic dispatch.
type IUnary_arith_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnary_arith_opContext differentiates from other interfaces.
	IsUnary_arith_opContext()
}

type Unary_arith_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_arith_opContext() *Unary_arith_opContext {
	var p = new(Unary_arith_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_unary_arith_op
	return p
}

func InitEmptyUnary_arith_opContext(p *Unary_arith_opContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_unary_arith_op
}

func (*Unary_arith_opContext) IsUnary_arith_opContext() {}

func NewUnary_arith_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_arith_opContext {
	var p = new(Unary_arith_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_unary_arith_op

	return p
}

func (s *Unary_arith_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Unary_arith_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_arith_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Unary_arith_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterUnary_arith_op(s)
	}
}

func (s *Unary_arith_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitUnary_arith_op(s)
	}
}

func (s *Unary_arith_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitUnary_arith_op(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Unary_arith_op() (localctx IUnary_arith_opContext) {
	localctx = NewUnary_arith_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SynlongParserRULE_unary_arith_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(733)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SynlongParserT__33 || _la == SynlongParserT__59 || _la == SynlongParserT__98) {
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


// IBin_arith_opContext is an interface to support dynamic dispatch.
type IBin_arith_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBin_arith_opContext differentiates from other interfaces.
	IsBin_arith_opContext()
}

type Bin_arith_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_arith_opContext() *Bin_arith_opContext {
	var p = new(Bin_arith_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_arith_op
	return p
}

func InitEmptyBin_arith_opContext(p *Bin_arith_opContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_arith_op
}

func (*Bin_arith_opContext) IsBin_arith_opContext() {}

func NewBin_arith_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_arith_opContext {
	var p = new(Bin_arith_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_bin_arith_op

	return p
}

func (s *Bin_arith_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Bin_arith_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_arith_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Bin_arith_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterBin_arith_op(s)
	}
}

func (s *Bin_arith_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitBin_arith_op(s)
	}
}

func (s *Bin_arith_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitBin_arith_op(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Bin_arith_op() (localctx IBin_arith_opContext) {
	localctx = NewBin_arith_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SynlongParserRULE_bin_arith_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(735)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 60)) & ^0x3f) == 0 && ((int64(1) << (_la - 60)) & 17042430230529) != 0)) {
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


// IBin_relation_opContext is an interface to support dynamic dispatch.
type IBin_relation_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBin_relation_opContext differentiates from other interfaces.
	IsBin_relation_opContext()
}

type Bin_relation_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_relation_opContext() *Bin_relation_opContext {
	var p = new(Bin_relation_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_relation_op
	return p
}

func InitEmptyBin_relation_opContext(p *Bin_relation_opContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_relation_op
}

func (*Bin_relation_opContext) IsBin_relation_opContext() {}

func NewBin_relation_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_relation_opContext {
	var p = new(Bin_relation_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_bin_relation_op

	return p
}

func (s *Bin_relation_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Bin_relation_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_relation_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Bin_relation_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterBin_relation_op(s)
	}
}

func (s *Bin_relation_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitBin_relation_op(s)
	}
}

func (s *Bin_relation_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitBin_relation_op(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Bin_relation_op() (localctx IBin_relation_opContext) {
	localctx = NewBin_relation_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SynlongParserRULE_bin_relation_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(737)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SynlongParserT__2 || ((int64((_la - 104)) & ^0x3f) == 0 && ((int64(1) << (_la - 104)) & 31) != 0)) {
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


// IBin_bool_opContext is an interface to support dynamic dispatch.
type IBin_bool_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBin_bool_opContext differentiates from other interfaces.
	IsBin_bool_opContext()
}

type Bin_bool_opContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBin_bool_opContext() *Bin_bool_opContext {
	var p = new(Bin_bool_opContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_bool_op
	return p
}

func InitEmptyBin_bool_opContext(p *Bin_bool_opContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_bin_bool_op
}

func (*Bin_bool_opContext) IsBin_bool_opContext() {}

func NewBin_bool_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bin_bool_opContext {
	var p = new(Bin_bool_opContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_bin_bool_op

	return p
}

func (s *Bin_bool_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Bin_bool_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bin_bool_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Bin_bool_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterBin_bool_op(s)
	}
}

func (s *Bin_bool_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitBin_bool_op(s)
	}
}

func (s *Bin_bool_opContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitBin_bool_op(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Bin_bool_op() (localctx IBin_bool_opContext) {
	localctx = NewBin_bool_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SynlongParserRULE_bin_bool_op)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(739)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 109)) & ^0x3f) == 0 && ((int64(1) << (_la - 109)) & 7) != 0)) {
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


// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CHAR() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	UINT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	REAL() antlr.TerminalNode
	USHORT() antlr.TerminalNode
	SHORT() antlr.TerminalNode

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SynlongParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SynlongParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CHAR() antlr.TerminalNode {
	return s.GetToken(SynlongParserCHAR, 0)
}

func (s *AtomContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SynlongParserINTEGER, 0)
}

func (s *AtomContext) UINT() antlr.TerminalNode {
	return s.GetToken(SynlongParserUINT, 0)
}

func (s *AtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SynlongParserFLOAT, 0)
}

func (s *AtomContext) REAL() antlr.TerminalNode {
	return s.GetToken(SynlongParserREAL, 0)
}

func (s *AtomContext) USHORT() antlr.TerminalNode {
	return s.GetToken(SynlongParserUSHORT, 0)
}

func (s *AtomContext) SHORT() antlr.TerminalNode {
	return s.GetToken(SynlongParserSHORT, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SynlongListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SynlongVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *SynlongParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SynlongParserRULE_atom)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(741)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 61)) & ^0x3f) == 0 && ((int64(1) << (_la - 61)) & 1143914305352105987) != 0)) {
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


func (p *SynlongParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
			var t *Type_exprContext = nil
			if localctx != nil { t = localctx.(*Type_exprContext) }
			return p.Type_expr_Sempred(t, predIndex)

	case 10:
			var t *Const_exprContext = nil
			if localctx != nil { t = localctx.(*Const_exprContext) }
			return p.Const_expr_Sempred(t, predIndex)

	case 34:
			var t *Simple_exprContext = nil
			if localctx != nil { t = localctx.(*Simple_exprContext) }
			return p.Simple_expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SynlongParser) Type_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SynlongParser) Const_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SynlongParser) Simple_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}


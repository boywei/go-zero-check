// Code generated from /Users/wei/GoProjects/go-zero-check/internal/synlong/Synlong.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Synlong

import "github.com/antlr4-go/antlr/v4"

// BaseSynlongListener is a complete listener for a parse tree produced by SynlongParser.
type BaseSynlongListener struct{}

var _ SynlongListener = &BaseSynlongListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSynlongListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSynlongListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSynlongListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSynlongListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseSynlongListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseSynlongListener) ExitProgram(ctx *ProgramContext) {}

// EnterDecls is called when production decls is entered.
func (s *BaseSynlongListener) EnterDecls(ctx *DeclsContext) {}

// ExitDecls is called when production decls is exited.
func (s *BaseSynlongListener) ExitDecls(ctx *DeclsContext) {}

// EnterType_block is called when production type_block is entered.
func (s *BaseSynlongListener) EnterType_block(ctx *Type_blockContext) {}

// ExitType_block is called when production type_block is exited.
func (s *BaseSynlongListener) ExitType_block(ctx *Type_blockContext) {}

// EnterType_decl is called when production type_decl is entered.
func (s *BaseSynlongListener) EnterType_decl(ctx *Type_declContext) {}

// ExitType_decl is called when production type_decl is exited.
func (s *BaseSynlongListener) ExitType_decl(ctx *Type_declContext) {}

// EnterType_def is called when production type_def is entered.
func (s *BaseSynlongListener) EnterType_def(ctx *Type_defContext) {}

// ExitType_def is called when production type_def is exited.
func (s *BaseSynlongListener) ExitType_def(ctx *Type_defContext) {}

// EnterType_expr is called when production type_expr is entered.
func (s *BaseSynlongListener) EnterType_expr(ctx *Type_exprContext) {}

// ExitType_expr is called when production type_expr is exited.
func (s *BaseSynlongListener) ExitType_expr(ctx *Type_exprContext) {}

// EnterField_decl is called when production field_decl is entered.
func (s *BaseSynlongListener) EnterField_decl(ctx *Field_declContext) {}

// ExitField_decl is called when production field_decl is exited.
func (s *BaseSynlongListener) ExitField_decl(ctx *Field_declContext) {}

// EnterTypevar is called when production typevar is entered.
func (s *BaseSynlongListener) EnterTypevar(ctx *TypevarContext) {}

// ExitTypevar is called when production typevar is exited.
func (s *BaseSynlongListener) ExitTypevar(ctx *TypevarContext) {}

// EnterConst_block is called when production const_block is entered.
func (s *BaseSynlongListener) EnterConst_block(ctx *Const_blockContext) {}

// ExitConst_block is called when production const_block is exited.
func (s *BaseSynlongListener) ExitConst_block(ctx *Const_blockContext) {}

// EnterConst_decl is called when production const_decl is entered.
func (s *BaseSynlongListener) EnterConst_decl(ctx *Const_declContext) {}

// ExitConst_decl is called when production const_decl is exited.
func (s *BaseSynlongListener) ExitConst_decl(ctx *Const_declContext) {}

// EnterConst_expr is called when production const_expr is entered.
func (s *BaseSynlongListener) EnterConst_expr(ctx *Const_exprContext) {}

// ExitConst_expr is called when production const_expr is exited.
func (s *BaseSynlongListener) ExitConst_expr(ctx *Const_exprContext) {}

// EnterConst_list is called when production const_list is entered.
func (s *BaseSynlongListener) EnterConst_list(ctx *Const_listContext) {}

// ExitConst_list is called when production const_list is exited.
func (s *BaseSynlongListener) ExitConst_list(ctx *Const_listContext) {}

// EnterConst_label_expr is called when production const_label_expr is entered.
func (s *BaseSynlongListener) EnterConst_label_expr(ctx *Const_label_exprContext) {}

// ExitConst_label_expr is called when production const_label_expr is exited.
func (s *BaseSynlongListener) ExitConst_label_expr(ctx *Const_label_exprContext) {}

// EnterUser_op_decl is called when production user_op_decl is entered.
func (s *BaseSynlongListener) EnterUser_op_decl(ctx *User_op_declContext) {}

// ExitUser_op_decl is called when production user_op_decl is exited.
func (s *BaseSynlongListener) ExitUser_op_decl(ctx *User_op_declContext) {}

// EnterOp_kind is called when production op_kind is entered.
func (s *BaseSynlongListener) EnterOp_kind(ctx *Op_kindContext) {}

// ExitOp_kind is called when production op_kind is exited.
func (s *BaseSynlongListener) ExitOp_kind(ctx *Op_kindContext) {}

// EnterParams is called when production params is entered.
func (s *BaseSynlongListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseSynlongListener) ExitParams(ctx *ParamsContext) {}

// EnterReturns_clause is called when production returns_clause is entered.
func (s *BaseSynlongListener) EnterReturns_clause(ctx *Returns_clauseContext) {}

// ExitReturns_clause is called when production returns_clause is exited.
func (s *BaseSynlongListener) ExitReturns_clause(ctx *Returns_clauseContext) {}

// EnterOp_body is called when production op_body is entered.
func (s *BaseSynlongListener) EnterOp_body(ctx *Op_bodyContext) {}

// ExitOp_body is called when production op_body is exited.
func (s *BaseSynlongListener) ExitOp_body(ctx *Op_bodyContext) {}

// EnterLocal_block is called when production local_block is entered.
func (s *BaseSynlongListener) EnterLocal_block(ctx *Local_blockContext) {}

// ExitLocal_block is called when production local_block is exited.
func (s *BaseSynlongListener) ExitLocal_block(ctx *Local_blockContext) {}

// EnterVar_decls is called when production var_decls is entered.
func (s *BaseSynlongListener) EnterVar_decls(ctx *Var_declsContext) {}

// ExitVar_decls is called when production var_decls is exited.
func (s *BaseSynlongListener) ExitVar_decls(ctx *Var_declsContext) {}

// EnterVar_id is called when production var_id is entered.
func (s *BaseSynlongListener) EnterVar_id(ctx *Var_idContext) {}

// ExitVar_id is called when production var_id is exited.
func (s *BaseSynlongListener) ExitVar_id(ctx *Var_idContext) {}

// EnterWhen_decl is called when production when_decl is entered.
func (s *BaseSynlongListener) EnterWhen_decl(ctx *When_declContext) {}

// ExitWhen_decl is called when production when_decl is exited.
func (s *BaseSynlongListener) ExitWhen_decl(ctx *When_declContext) {}

// EnterClock_expr is called when production clock_expr is entered.
func (s *BaseSynlongListener) EnterClock_expr(ctx *Clock_exprContext) {}

// ExitClock_expr is called when production clock_expr is exited.
func (s *BaseSynlongListener) ExitClock_expr(ctx *Clock_exprContext) {}

// EnterLast_decl is called when production last_decl is entered.
func (s *BaseSynlongListener) EnterLast_decl(ctx *Last_declContext) {}

// ExitLast_decl is called when production last_decl is exited.
func (s *BaseSynlongListener) ExitLast_decl(ctx *Last_declContext) {}

// EnterEquation is called when production equation is entered.
func (s *BaseSynlongListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BaseSynlongListener) ExitEquation(ctx *EquationContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BaseSynlongListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BaseSynlongListener) ExitLhs(ctx *LhsContext) {}

// EnterLhs_id is called when production lhs_id is entered.
func (s *BaseSynlongListener) EnterLhs_id(ctx *Lhs_idContext) {}

// ExitLhs_id is called when production lhs_id is exited.
func (s *BaseSynlongListener) ExitLhs_id(ctx *Lhs_idContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseSynlongListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseSynlongListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterReturns_var is called when production returns_var is entered.
func (s *BaseSynlongListener) EnterReturns_var(ctx *Returns_varContext) {}

// ExitReturns_var is called when production returns_var is exited.
func (s *BaseSynlongListener) ExitReturns_var(ctx *Returns_varContext) {}

// EnterState_machine is called when production state_machine is entered.
func (s *BaseSynlongListener) EnterState_machine(ctx *State_machineContext) {}

// ExitState_machine is called when production state_machine is exited.
func (s *BaseSynlongListener) ExitState_machine(ctx *State_machineContext) {}

// EnterState_decl is called when production state_decl is entered.
func (s *BaseSynlongListener) EnterState_decl(ctx *State_declContext) {}

// ExitState_decl is called when production state_decl is exited.
func (s *BaseSynlongListener) ExitState_decl(ctx *State_declContext) {}

// EnterData_def is called when production data_def is entered.
func (s *BaseSynlongListener) EnterData_def(ctx *Data_defContext) {}

// ExitData_def is called when production data_def is exited.
func (s *BaseSynlongListener) ExitData_def(ctx *Data_defContext) {}

// EnterTransition is called when production transition is entered.
func (s *BaseSynlongListener) EnterTransition(ctx *TransitionContext) {}

// ExitTransition is called when production transition is exited.
func (s *BaseSynlongListener) ExitTransition(ctx *TransitionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSynlongListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSynlongListener) ExitExpr(ctx *ExprContext) {}

// EnterSimple_expr is called when production simple_expr is entered.
func (s *BaseSynlongListener) EnterSimple_expr(ctx *Simple_exprContext) {}

// ExitSimple_expr is called when production simple_expr is exited.
func (s *BaseSynlongListener) ExitSimple_expr(ctx *Simple_exprContext) {}

// EnterTempo_expr is called when production tempo_expr is entered.
func (s *BaseSynlongListener) EnterTempo_expr(ctx *Tempo_exprContext) {}

// ExitTempo_expr is called when production tempo_expr is exited.
func (s *BaseSynlongListener) ExitTempo_expr(ctx *Tempo_exprContext) {}

// EnterBool_expr is called when production bool_expr is entered.
func (s *BaseSynlongListener) EnterBool_expr(ctx *Bool_exprContext) {}

// ExitBool_expr is called when production bool_expr is exited.
func (s *BaseSynlongListener) ExitBool_expr(ctx *Bool_exprContext) {}

// EnterArray_expr is called when production array_expr is entered.
func (s *BaseSynlongListener) EnterArray_expr(ctx *Array_exprContext) {}

// ExitArray_expr is called when production array_expr is exited.
func (s *BaseSynlongListener) ExitArray_expr(ctx *Array_exprContext) {}

// EnterStruct_expr is called when production struct_expr is entered.
func (s *BaseSynlongListener) EnterStruct_expr(ctx *Struct_exprContext) {}

// ExitStruct_expr is called when production struct_expr is exited.
func (s *BaseSynlongListener) ExitStruct_expr(ctx *Struct_exprContext) {}

// EnterLabel_expr is called when production label_expr is entered.
func (s *BaseSynlongListener) EnterLabel_expr(ctx *Label_exprContext) {}

// ExitLabel_expr is called when production label_expr is exited.
func (s *BaseSynlongListener) ExitLabel_expr(ctx *Label_exprContext) {}

// EnterMixed_constructor is called when production mixed_constructor is entered.
func (s *BaseSynlongListener) EnterMixed_constructor(ctx *Mixed_constructorContext) {}

// ExitMixed_constructor is called when production mixed_constructor is exited.
func (s *BaseSynlongListener) ExitMixed_constructor(ctx *Mixed_constructorContext) {}

// EnterLabel_or_index is called when production label_or_index is entered.
func (s *BaseSynlongListener) EnterLabel_or_index(ctx *Label_or_indexContext) {}

// ExitLabel_or_index is called when production label_or_index is exited.
func (s *BaseSynlongListener) ExitLabel_or_index(ctx *Label_or_indexContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSynlongListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSynlongListener) ExitIndex(ctx *IndexContext) {}

// EnterSwitch_expr is called when production switch_expr is entered.
func (s *BaseSynlongListener) EnterSwitch_expr(ctx *Switch_exprContext) {}

// ExitSwitch_expr is called when production switch_expr is exited.
func (s *BaseSynlongListener) ExitSwitch_expr(ctx *Switch_exprContext) {}

// EnterCase_expr is called when production case_expr is entered.
func (s *BaseSynlongListener) EnterCase_expr(ctx *Case_exprContext) {}

// ExitCase_expr is called when production case_expr is exited.
func (s *BaseSynlongListener) ExitCase_expr(ctx *Case_exprContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSynlongListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSynlongListener) ExitPattern(ctx *PatternContext) {}

// EnterApply_expr is called when production apply_expr is entered.
func (s *BaseSynlongListener) EnterApply_expr(ctx *Apply_exprContext) {}

// ExitApply_expr is called when production apply_expr is exited.
func (s *BaseSynlongListener) ExitApply_expr(ctx *Apply_exprContext) {}

// EnterPrefix_operator is called when production prefix_operator is entered.
func (s *BaseSynlongListener) EnterPrefix_operator(ctx *Prefix_operatorContext) {}

// ExitPrefix_operator is called when production prefix_operator is exited.
func (s *BaseSynlongListener) ExitPrefix_operator(ctx *Prefix_operatorContext) {}

// EnterPrefix_unary_operator is called when production prefix_unary_operator is entered.
func (s *BaseSynlongListener) EnterPrefix_unary_operator(ctx *Prefix_unary_operatorContext) {}

// ExitPrefix_unary_operator is called when production prefix_unary_operator is exited.
func (s *BaseSynlongListener) ExitPrefix_unary_operator(ctx *Prefix_unary_operatorContext) {}

// EnterPrefix_binary_operator is called when production prefix_binary_operator is entered.
func (s *BaseSynlongListener) EnterPrefix_binary_operator(ctx *Prefix_binary_operatorContext) {}

// ExitPrefix_binary_operator is called when production prefix_binary_operator is exited.
func (s *BaseSynlongListener) ExitPrefix_binary_operator(ctx *Prefix_binary_operatorContext) {}

// EnterIterator is called when production iterator is entered.
func (s *BaseSynlongListener) EnterIterator(ctx *IteratorContext) {}

// ExitIterator is called when production iterator is exited.
func (s *BaseSynlongListener) ExitIterator(ctx *IteratorContext) {}

// EnterList is called when production list is entered.
func (s *BaseSynlongListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseSynlongListener) ExitList(ctx *ListContext) {}

// EnterUnary_arith_op is called when production unary_arith_op is entered.
func (s *BaseSynlongListener) EnterUnary_arith_op(ctx *Unary_arith_opContext) {}

// ExitUnary_arith_op is called when production unary_arith_op is exited.
func (s *BaseSynlongListener) ExitUnary_arith_op(ctx *Unary_arith_opContext) {}

// EnterBin_arith_op is called when production bin_arith_op is entered.
func (s *BaseSynlongListener) EnterBin_arith_op(ctx *Bin_arith_opContext) {}

// ExitBin_arith_op is called when production bin_arith_op is exited.
func (s *BaseSynlongListener) ExitBin_arith_op(ctx *Bin_arith_opContext) {}

// EnterBin_relation_op is called when production bin_relation_op is entered.
func (s *BaseSynlongListener) EnterBin_relation_op(ctx *Bin_relation_opContext) {}

// ExitBin_relation_op is called when production bin_relation_op is exited.
func (s *BaseSynlongListener) ExitBin_relation_op(ctx *Bin_relation_opContext) {}

// EnterBin_bool_op is called when production bin_bool_op is entered.
func (s *BaseSynlongListener) EnterBin_bool_op(ctx *Bin_bool_opContext) {}

// ExitBin_bool_op is called when production bin_bool_op is exited.
func (s *BaseSynlongListener) ExitBin_bool_op(ctx *Bin_bool_opContext) {}

// EnterAtom is called when production atom is entered.
func (s *BaseSynlongListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BaseSynlongListener) ExitAtom(ctx *AtomContext) {}

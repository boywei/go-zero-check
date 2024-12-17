// Code generated from /Users/wei/GoProjects/go-zero-check/internal/synlong/Synlong.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Synlong

import "github.com/antlr4-go/antlr/v4"


// SynlongListener is a complete listener for a parse tree produced by SynlongParser.
type SynlongListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterDecls is called when entering the decls production.
	EnterDecls(c *DeclsContext)

	// EnterType_block is called when entering the type_block production.
	EnterType_block(c *Type_blockContext)

	// EnterType_decl is called when entering the type_decl production.
	EnterType_decl(c *Type_declContext)

	// EnterType_def is called when entering the type_def production.
	EnterType_def(c *Type_defContext)

	// EnterType_expr is called when entering the type_expr production.
	EnterType_expr(c *Type_exprContext)

	// EnterField_decl is called when entering the field_decl production.
	EnterField_decl(c *Field_declContext)

	// EnterTypevar is called when entering the typevar production.
	EnterTypevar(c *TypevarContext)

	// EnterConst_block is called when entering the const_block production.
	EnterConst_block(c *Const_blockContext)

	// EnterConst_decl is called when entering the const_decl production.
	EnterConst_decl(c *Const_declContext)

	// EnterConst_expr is called when entering the const_expr production.
	EnterConst_expr(c *Const_exprContext)

	// EnterConst_list is called when entering the const_list production.
	EnterConst_list(c *Const_listContext)

	// EnterConst_label_expr is called when entering the const_label_expr production.
	EnterConst_label_expr(c *Const_label_exprContext)

	// EnterUser_op_decl is called when entering the user_op_decl production.
	EnterUser_op_decl(c *User_op_declContext)

	// EnterOp_kind is called when entering the op_kind production.
	EnterOp_kind(c *Op_kindContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterReturns_clause is called when entering the returns_clause production.
	EnterReturns_clause(c *Returns_clauseContext)

	// EnterOp_body is called when entering the op_body production.
	EnterOp_body(c *Op_bodyContext)

	// EnterLocal_block is called when entering the local_block production.
	EnterLocal_block(c *Local_blockContext)

	// EnterVar_decls is called when entering the var_decls production.
	EnterVar_decls(c *Var_declsContext)

	// EnterVar_id is called when entering the var_id production.
	EnterVar_id(c *Var_idContext)

	// EnterWhen_decl is called when entering the when_decl production.
	EnterWhen_decl(c *When_declContext)

	// EnterClock_expr is called when entering the clock_expr production.
	EnterClock_expr(c *Clock_exprContext)

	// EnterLast_decl is called when entering the last_decl production.
	EnterLast_decl(c *Last_declContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterLhs_id is called when entering the lhs_id production.
	EnterLhs_id(c *Lhs_idContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterReturns_var is called when entering the returns_var production.
	EnterReturns_var(c *Returns_varContext)

	// EnterState_machine is called when entering the state_machine production.
	EnterState_machine(c *State_machineContext)

	// EnterState_decl is called when entering the state_decl production.
	EnterState_decl(c *State_declContext)

	// EnterData_def is called when entering the data_def production.
	EnterData_def(c *Data_defContext)

	// EnterTransition is called when entering the transition production.
	EnterTransition(c *TransitionContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSimple_expr is called when entering the simple_expr production.
	EnterSimple_expr(c *Simple_exprContext)

	// EnterTempo_expr is called when entering the tempo_expr production.
	EnterTempo_expr(c *Tempo_exprContext)

	// EnterBool_expr is called when entering the bool_expr production.
	EnterBool_expr(c *Bool_exprContext)

	// EnterArray_expr is called when entering the array_expr production.
	EnterArray_expr(c *Array_exprContext)

	// EnterStruct_expr is called when entering the struct_expr production.
	EnterStruct_expr(c *Struct_exprContext)

	// EnterLabel_expr is called when entering the label_expr production.
	EnterLabel_expr(c *Label_exprContext)

	// EnterMixed_constructor is called when entering the mixed_constructor production.
	EnterMixed_constructor(c *Mixed_constructorContext)

	// EnterLabel_or_index is called when entering the label_or_index production.
	EnterLabel_or_index(c *Label_or_indexContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterSwitch_expr is called when entering the switch_expr production.
	EnterSwitch_expr(c *Switch_exprContext)

	// EnterCase_expr is called when entering the case_expr production.
	EnterCase_expr(c *Case_exprContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterApply_expr is called when entering the apply_expr production.
	EnterApply_expr(c *Apply_exprContext)

	// EnterPrefix_operator is called when entering the prefix_operator production.
	EnterPrefix_operator(c *Prefix_operatorContext)

	// EnterPrefix_unary_operator is called when entering the prefix_unary_operator production.
	EnterPrefix_unary_operator(c *Prefix_unary_operatorContext)

	// EnterPrefix_binary_operator is called when entering the prefix_binary_operator production.
	EnterPrefix_binary_operator(c *Prefix_binary_operatorContext)

	// EnterIterator is called when entering the iterator production.
	EnterIterator(c *IteratorContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterUnary_arith_op is called when entering the unary_arith_op production.
	EnterUnary_arith_op(c *Unary_arith_opContext)

	// EnterBin_arith_op is called when entering the bin_arith_op production.
	EnterBin_arith_op(c *Bin_arith_opContext)

	// EnterBin_relation_op is called when entering the bin_relation_op production.
	EnterBin_relation_op(c *Bin_relation_opContext)

	// EnterBin_bool_op is called when entering the bin_bool_op production.
	EnterBin_bool_op(c *Bin_bool_opContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitDecls is called when exiting the decls production.
	ExitDecls(c *DeclsContext)

	// ExitType_block is called when exiting the type_block production.
	ExitType_block(c *Type_blockContext)

	// ExitType_decl is called when exiting the type_decl production.
	ExitType_decl(c *Type_declContext)

	// ExitType_def is called when exiting the type_def production.
	ExitType_def(c *Type_defContext)

	// ExitType_expr is called when exiting the type_expr production.
	ExitType_expr(c *Type_exprContext)

	// ExitField_decl is called when exiting the field_decl production.
	ExitField_decl(c *Field_declContext)

	// ExitTypevar is called when exiting the typevar production.
	ExitTypevar(c *TypevarContext)

	// ExitConst_block is called when exiting the const_block production.
	ExitConst_block(c *Const_blockContext)

	// ExitConst_decl is called when exiting the const_decl production.
	ExitConst_decl(c *Const_declContext)

	// ExitConst_expr is called when exiting the const_expr production.
	ExitConst_expr(c *Const_exprContext)

	// ExitConst_list is called when exiting the const_list production.
	ExitConst_list(c *Const_listContext)

	// ExitConst_label_expr is called when exiting the const_label_expr production.
	ExitConst_label_expr(c *Const_label_exprContext)

	// ExitUser_op_decl is called when exiting the user_op_decl production.
	ExitUser_op_decl(c *User_op_declContext)

	// ExitOp_kind is called when exiting the op_kind production.
	ExitOp_kind(c *Op_kindContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitReturns_clause is called when exiting the returns_clause production.
	ExitReturns_clause(c *Returns_clauseContext)

	// ExitOp_body is called when exiting the op_body production.
	ExitOp_body(c *Op_bodyContext)

	// ExitLocal_block is called when exiting the local_block production.
	ExitLocal_block(c *Local_blockContext)

	// ExitVar_decls is called when exiting the var_decls production.
	ExitVar_decls(c *Var_declsContext)

	// ExitVar_id is called when exiting the var_id production.
	ExitVar_id(c *Var_idContext)

	// ExitWhen_decl is called when exiting the when_decl production.
	ExitWhen_decl(c *When_declContext)

	// ExitClock_expr is called when exiting the clock_expr production.
	ExitClock_expr(c *Clock_exprContext)

	// ExitLast_decl is called when exiting the last_decl production.
	ExitLast_decl(c *Last_declContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitLhs_id is called when exiting the lhs_id production.
	ExitLhs_id(c *Lhs_idContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitReturns_var is called when exiting the returns_var production.
	ExitReturns_var(c *Returns_varContext)

	// ExitState_machine is called when exiting the state_machine production.
	ExitState_machine(c *State_machineContext)

	// ExitState_decl is called when exiting the state_decl production.
	ExitState_decl(c *State_declContext)

	// ExitData_def is called when exiting the data_def production.
	ExitData_def(c *Data_defContext)

	// ExitTransition is called when exiting the transition production.
	ExitTransition(c *TransitionContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSimple_expr is called when exiting the simple_expr production.
	ExitSimple_expr(c *Simple_exprContext)

	// ExitTempo_expr is called when exiting the tempo_expr production.
	ExitTempo_expr(c *Tempo_exprContext)

	// ExitBool_expr is called when exiting the bool_expr production.
	ExitBool_expr(c *Bool_exprContext)

	// ExitArray_expr is called when exiting the array_expr production.
	ExitArray_expr(c *Array_exprContext)

	// ExitStruct_expr is called when exiting the struct_expr production.
	ExitStruct_expr(c *Struct_exprContext)

	// ExitLabel_expr is called when exiting the label_expr production.
	ExitLabel_expr(c *Label_exprContext)

	// ExitMixed_constructor is called when exiting the mixed_constructor production.
	ExitMixed_constructor(c *Mixed_constructorContext)

	// ExitLabel_or_index is called when exiting the label_or_index production.
	ExitLabel_or_index(c *Label_or_indexContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitSwitch_expr is called when exiting the switch_expr production.
	ExitSwitch_expr(c *Switch_exprContext)

	// ExitCase_expr is called when exiting the case_expr production.
	ExitCase_expr(c *Case_exprContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitApply_expr is called when exiting the apply_expr production.
	ExitApply_expr(c *Apply_exprContext)

	// ExitPrefix_operator is called when exiting the prefix_operator production.
	ExitPrefix_operator(c *Prefix_operatorContext)

	// ExitPrefix_unary_operator is called when exiting the prefix_unary_operator production.
	ExitPrefix_unary_operator(c *Prefix_unary_operatorContext)

	// ExitPrefix_binary_operator is called when exiting the prefix_binary_operator production.
	ExitPrefix_binary_operator(c *Prefix_binary_operatorContext)

	// ExitIterator is called when exiting the iterator production.
	ExitIterator(c *IteratorContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitUnary_arith_op is called when exiting the unary_arith_op production.
	ExitUnary_arith_op(c *Unary_arith_opContext)

	// ExitBin_arith_op is called when exiting the bin_arith_op production.
	ExitBin_arith_op(c *Bin_arith_opContext)

	// ExitBin_relation_op is called when exiting the bin_relation_op production.
	ExitBin_relation_op(c *Bin_relation_opContext)

	// ExitBin_bool_op is called when exiting the bin_bool_op production.
	ExitBin_bool_op(c *Bin_bool_opContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)
}

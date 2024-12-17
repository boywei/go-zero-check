// Code generated from /Users/wei/GoProjects/go-zero-check/internal/synlong/Synlong.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Synlong

import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by SynlongParser.
type SynlongVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SynlongParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by SynlongParser#decls.
	VisitDecls(ctx *DeclsContext) interface{}

	// Visit a parse tree produced by SynlongParser#type_block.
	VisitType_block(ctx *Type_blockContext) interface{}

	// Visit a parse tree produced by SynlongParser#type_decl.
	VisitType_decl(ctx *Type_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#type_def.
	VisitType_def(ctx *Type_defContext) interface{}

	// Visit a parse tree produced by SynlongParser#type_expr.
	VisitType_expr(ctx *Type_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#field_decl.
	VisitField_decl(ctx *Field_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#typevar.
	VisitTypevar(ctx *TypevarContext) interface{}

	// Visit a parse tree produced by SynlongParser#const_block.
	VisitConst_block(ctx *Const_blockContext) interface{}

	// Visit a parse tree produced by SynlongParser#const_decl.
	VisitConst_decl(ctx *Const_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#const_expr.
	VisitConst_expr(ctx *Const_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#const_list.
	VisitConst_list(ctx *Const_listContext) interface{}

	// Visit a parse tree produced by SynlongParser#const_label_expr.
	VisitConst_label_expr(ctx *Const_label_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#user_op_decl.
	VisitUser_op_decl(ctx *User_op_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#op_kind.
	VisitOp_kind(ctx *Op_kindContext) interface{}

	// Visit a parse tree produced by SynlongParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by SynlongParser#returns_clause.
	VisitReturns_clause(ctx *Returns_clauseContext) interface{}

	// Visit a parse tree produced by SynlongParser#op_body.
	VisitOp_body(ctx *Op_bodyContext) interface{}

	// Visit a parse tree produced by SynlongParser#local_block.
	VisitLocal_block(ctx *Local_blockContext) interface{}

	// Visit a parse tree produced by SynlongParser#var_decls.
	VisitVar_decls(ctx *Var_declsContext) interface{}

	// Visit a parse tree produced by SynlongParser#var_id.
	VisitVar_id(ctx *Var_idContext) interface{}

	// Visit a parse tree produced by SynlongParser#when_decl.
	VisitWhen_decl(ctx *When_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#clock_expr.
	VisitClock_expr(ctx *Clock_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#last_decl.
	VisitLast_decl(ctx *Last_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by SynlongParser#lhs.
	VisitLhs(ctx *LhsContext) interface{}

	// Visit a parse tree produced by SynlongParser#lhs_id.
	VisitLhs_id(ctx *Lhs_idContext) interface{}

	// Visit a parse tree produced by SynlongParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by SynlongParser#returns_var.
	VisitReturns_var(ctx *Returns_varContext) interface{}

	// Visit a parse tree produced by SynlongParser#state_machine.
	VisitState_machine(ctx *State_machineContext) interface{}

	// Visit a parse tree produced by SynlongParser#state_decl.
	VisitState_decl(ctx *State_declContext) interface{}

	// Visit a parse tree produced by SynlongParser#data_def.
	VisitData_def(ctx *Data_defContext) interface{}

	// Visit a parse tree produced by SynlongParser#transition.
	VisitTransition(ctx *TransitionContext) interface{}

	// Visit a parse tree produced by SynlongParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SynlongParser#simple_expr.
	VisitSimple_expr(ctx *Simple_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#tempo_expr.
	VisitTempo_expr(ctx *Tempo_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#bool_expr.
	VisitBool_expr(ctx *Bool_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#array_expr.
	VisitArray_expr(ctx *Array_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#struct_expr.
	VisitStruct_expr(ctx *Struct_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#label_expr.
	VisitLabel_expr(ctx *Label_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#mixed_constructor.
	VisitMixed_constructor(ctx *Mixed_constructorContext) interface{}

	// Visit a parse tree produced by SynlongParser#label_or_index.
	VisitLabel_or_index(ctx *Label_or_indexContext) interface{}

	// Visit a parse tree produced by SynlongParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by SynlongParser#switch_expr.
	VisitSwitch_expr(ctx *Switch_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#case_expr.
	VisitCase_expr(ctx *Case_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by SynlongParser#apply_expr.
	VisitApply_expr(ctx *Apply_exprContext) interface{}

	// Visit a parse tree produced by SynlongParser#prefix_operator.
	VisitPrefix_operator(ctx *Prefix_operatorContext) interface{}

	// Visit a parse tree produced by SynlongParser#prefix_unary_operator.
	VisitPrefix_unary_operator(ctx *Prefix_unary_operatorContext) interface{}

	// Visit a parse tree produced by SynlongParser#prefix_binary_operator.
	VisitPrefix_binary_operator(ctx *Prefix_binary_operatorContext) interface{}

	// Visit a parse tree produced by SynlongParser#iterator.
	VisitIterator(ctx *IteratorContext) interface{}

	// Visit a parse tree produced by SynlongParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by SynlongParser#unary_arith_op.
	VisitUnary_arith_op(ctx *Unary_arith_opContext) interface{}

	// Visit a parse tree produced by SynlongParser#bin_arith_op.
	VisitBin_arith_op(ctx *Bin_arith_opContext) interface{}

	// Visit a parse tree produced by SynlongParser#bin_relation_op.
	VisitBin_relation_op(ctx *Bin_relation_opContext) interface{}

	// Visit a parse tree produced by SynlongParser#bin_bool_op.
	VisitBin_bool_op(ctx *Bin_bool_opContext) interface{}

	// Visit a parse tree produced by SynlongParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

}
// Code generated from /Users/wei/GoProjects/go-zero-check/internal/synlong/Synlong.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Synlong

import "github.com/antlr4-go/antlr/v4"


type BaseSynlongVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSynlongVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitDecls(ctx *DeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitType_block(ctx *Type_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitType_decl(ctx *Type_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitType_def(ctx *Type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitType_expr(ctx *Type_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitField_decl(ctx *Field_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitTypevar(ctx *TypevarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitConst_block(ctx *Const_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitConst_decl(ctx *Const_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitConst_expr(ctx *Const_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitConst_list(ctx *Const_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitConst_label_expr(ctx *Const_label_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitUser_op_decl(ctx *User_op_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitOp_kind(ctx *Op_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitReturns_clause(ctx *Returns_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitOp_body(ctx *Op_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLocal_block(ctx *Local_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitVar_decls(ctx *Var_declsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitVar_id(ctx *Var_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitWhen_decl(ctx *When_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitClock_expr(ctx *Clock_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLast_decl(ctx *Last_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLhs(ctx *LhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLhs_id(ctx *Lhs_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitReturns_var(ctx *Returns_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitState_machine(ctx *State_machineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitState_decl(ctx *State_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitData_def(ctx *Data_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitTransition(ctx *TransitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitSimple_expr(ctx *Simple_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitTempo_expr(ctx *Tempo_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitBool_expr(ctx *Bool_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitArray_expr(ctx *Array_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitStruct_expr(ctx *Struct_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLabel_expr(ctx *Label_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitMixed_constructor(ctx *Mixed_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitLabel_or_index(ctx *Label_or_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitSwitch_expr(ctx *Switch_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitCase_expr(ctx *Case_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitApply_expr(ctx *Apply_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitPrefix_operator(ctx *Prefix_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitPrefix_unary_operator(ctx *Prefix_unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitPrefix_binary_operator(ctx *Prefix_binary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitIterator(ctx *IteratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitUnary_arith_op(ctx *Unary_arith_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitBin_arith_op(ctx *Bin_arith_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitBin_relation_op(ctx *Bin_relation_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitBin_bool_op(ctx *Bin_bool_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSynlongVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

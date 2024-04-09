// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lustre
import "github.com/antlr4-go/antlr/v4"


type BaseLustreVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLustreVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitTypedef(ctx *TypedefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitNode(ctx *NodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitVarDeclList(ctx *VarDeclListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitVarDeclGroup(ctx *VarDeclGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitPlainType(ctx *PlainTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitEnumType(ctx *EnumTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRealType(ctx *RealTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitSubrangeType(ctx *SubrangeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitUserType(ctx *UserTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitBound(ctx *BoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRealizabilityInputs(ctx *RealizabilityInputsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIvc(ctx *IvcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitMain(ctx *MainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitAssertion(ctx *AssertionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitEquation(ctx *EquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitLhs(ctx *LhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRecordExpr(ctx *RecordExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitCastExpr(ctx *CastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRealExpr(ctx *RealExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIfThenElseExpr(ctx *IfThenElseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitBinaryExpr(ctx *BinaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitPreExpr(ctx *PreExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRecordAccessExpr(ctx *RecordAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitNegateExpr(ctx *NegateExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitCondactExpr(ctx *CondactExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrayAccessExpr(ctx *ArrayAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrayUpdateExpr(ctx *ArrayUpdateExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRecordUpdateExpr(ctx *RecordUpdateExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitBaseEID(ctx *BaseEIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrayEID(ctx *ArrayEIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitRecordEID(ctx *RecordEIDContext) interface{} {
	return v.VisitChildren(ctx)
}

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

func (v *BaseLustreVisitor) VisitStateMachine(ctx *StateMachineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitStateDecl(ctx *StateDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitTransition(ctx *TransitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitArrow(ctx *ArrowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitFork(ctx *ForkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitElsifFork(ctx *ElsifForkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitElseFork(ctx *ElseForkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitTarget(ctx *TargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitActions(ctx *ActionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitDataDef(ctx *DataDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitLocalBlock(ctx *LocalBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitEqs(ctx *EqsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitVarDecl(ctx *VarDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitVarID(ctx *VarIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitDefaultDecl(ctx *DefaultDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitLastDecl(ctx *LastDeclContext) interface{} {
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

func (v *BaseLustreVisitor) VisitSimpleEquation(ctx *SimpleEquationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitLhs(ctx *LhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitLhsID(ctx *LhsIDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitControlBlock(ctx *ControlBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitEmission(ctx *EmissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitEmissionBody(ctx *EmissionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitReturn(ctx *ReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitReturnVar(ctx *ReturnVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitClockedBlock(ctx *ClockedBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitIfBlock(ctx *IfBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitMatchBlock(ctx *MatchBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLustreVisitor) VisitPattern(ctx *PatternContext) interface{} {
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

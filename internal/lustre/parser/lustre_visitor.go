// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lustre
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by LustreParser.
type LustreVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LustreParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LustreParser#typedef.
	VisitTypedef(ctx *TypedefContext) interface{}

	// Visit a parse tree produced by LustreParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by LustreParser#node.
	VisitNode(ctx *NodeContext) interface{}

	// Visit a parse tree produced by LustreParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by LustreParser#varDeclList.
	VisitVarDeclList(ctx *VarDeclListContext) interface{}

	// Visit a parse tree produced by LustreParser#varDeclGroup.
	VisitVarDeclGroup(ctx *VarDeclGroupContext) interface{}

	// Visit a parse tree produced by LustreParser#plainType.
	VisitPlainType(ctx *PlainTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#enumType.
	VisitEnumType(ctx *EnumTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#realType.
	VisitRealType(ctx *RealTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#subrangeType.
	VisitSubrangeType(ctx *SubrangeTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#intType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#userType.
	VisitUserType(ctx *UserTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#boolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by LustreParser#bound.
	VisitBound(ctx *BoundContext) interface{}

	// Visit a parse tree produced by LustreParser#stateMachine.
	VisitStateMachine(ctx *StateMachineContext) interface{}

	// Visit a parse tree produced by LustreParser#stateDecl.
	VisitStateDecl(ctx *StateDeclContext) interface{}

	// Visit a parse tree produced by LustreParser#transition.
	VisitTransition(ctx *TransitionContext) interface{}

	// Visit a parse tree produced by LustreParser#arrow.
	VisitArrow(ctx *ArrowContext) interface{}

	// Visit a parse tree produced by LustreParser#fork.
	VisitFork(ctx *ForkContext) interface{}

	// Visit a parse tree produced by LustreParser#elsifFork.
	VisitElsifFork(ctx *ElsifForkContext) interface{}

	// Visit a parse tree produced by LustreParser#elseFork.
	VisitElseFork(ctx *ElseForkContext) interface{}

	// Visit a parse tree produced by LustreParser#target.
	VisitTarget(ctx *TargetContext) interface{}

	// Visit a parse tree produced by LustreParser#actions.
	VisitActions(ctx *ActionsContext) interface{}

	// Visit a parse tree produced by LustreParser#dataDef.
	VisitDataDef(ctx *DataDefContext) interface{}

	// Visit a parse tree produced by LustreParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by LustreParser#localBlock.
	VisitLocalBlock(ctx *LocalBlockContext) interface{}

	// Visit a parse tree produced by LustreParser#eqs.
	VisitEqs(ctx *EqsContext) interface{}

	// Visit a parse tree produced by LustreParser#varDecl.
	VisitVarDecl(ctx *VarDeclContext) interface{}

	// Visit a parse tree produced by LustreParser#varID.
	VisitVarID(ctx *VarIDContext) interface{}

	// Visit a parse tree produced by LustreParser#defaultDecl.
	VisitDefaultDecl(ctx *DefaultDeclContext) interface{}

	// Visit a parse tree produced by LustreParser#lastDecl.
	VisitLastDecl(ctx *LastDeclContext) interface{}

	// Visit a parse tree produced by LustreParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by LustreParser#realizabilityInputs.
	VisitRealizabilityInputs(ctx *RealizabilityInputsContext) interface{}

	// Visit a parse tree produced by LustreParser#ivc.
	VisitIvc(ctx *IvcContext) interface{}

	// Visit a parse tree produced by LustreParser#main.
	VisitMain(ctx *MainContext) interface{}

	// Visit a parse tree produced by LustreParser#assertion.
	VisitAssertion(ctx *AssertionContext) interface{}

	// Visit a parse tree produced by LustreParser#equation.
	VisitEquation(ctx *EquationContext) interface{}

	// Visit a parse tree produced by LustreParser#simpleEquation.
	VisitSimpleEquation(ctx *SimpleEquationContext) interface{}

	// Visit a parse tree produced by LustreParser#lhs.
	VisitLhs(ctx *LhsContext) interface{}

	// Visit a parse tree produced by LustreParser#lhsID.
	VisitLhsID(ctx *LhsIDContext) interface{}

	// Visit a parse tree produced by LustreParser#controlBlock.
	VisitControlBlock(ctx *ControlBlockContext) interface{}

	// Visit a parse tree produced by LustreParser#emission.
	VisitEmission(ctx *EmissionContext) interface{}

	// Visit a parse tree produced by LustreParser#emissionBody.
	VisitEmissionBody(ctx *EmissionBodyContext) interface{}

	// Visit a parse tree produced by LustreParser#return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by LustreParser#returnVar.
	VisitReturnVar(ctx *ReturnVarContext) interface{}

	// Visit a parse tree produced by LustreParser#clockedBlock.
	VisitClockedBlock(ctx *ClockedBlockContext) interface{}

	// Visit a parse tree produced by LustreParser#ifBlock.
	VisitIfBlock(ctx *IfBlockContext) interface{}

	// Visit a parse tree produced by LustreParser#matchBlock.
	VisitMatchBlock(ctx *MatchBlockContext) interface{}

	// Visit a parse tree produced by LustreParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by LustreParser#recordExpr.
	VisitRecordExpr(ctx *RecordExprContext) interface{}

	// Visit a parse tree produced by LustreParser#intExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by LustreParser#arrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by LustreParser#castExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by LustreParser#realExpr.
	VisitRealExpr(ctx *RealExprContext) interface{}

	// Visit a parse tree produced by LustreParser#ifThenElseExpr.
	VisitIfThenElseExpr(ctx *IfThenElseExprContext) interface{}

	// Visit a parse tree produced by LustreParser#binaryExpr.
	VisitBinaryExpr(ctx *BinaryExprContext) interface{}

	// Visit a parse tree produced by LustreParser#preExpr.
	VisitPreExpr(ctx *PreExprContext) interface{}

	// Visit a parse tree produced by LustreParser#recordAccessExpr.
	VisitRecordAccessExpr(ctx *RecordAccessExprContext) interface{}

	// Visit a parse tree produced by LustreParser#negateExpr.
	VisitNegateExpr(ctx *NegateExprContext) interface{}

	// Visit a parse tree produced by LustreParser#condactExpr.
	VisitCondactExpr(ctx *CondactExprContext) interface{}

	// Visit a parse tree produced by LustreParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by LustreParser#arrayAccessExpr.
	VisitArrayAccessExpr(ctx *ArrayAccessExprContext) interface{}

	// Visit a parse tree produced by LustreParser#arrayUpdateExpr.
	VisitArrayUpdateExpr(ctx *ArrayUpdateExprContext) interface{}

	// Visit a parse tree produced by LustreParser#boolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by LustreParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by LustreParser#tupleExpr.
	VisitTupleExpr(ctx *TupleExprContext) interface{}

	// Visit a parse tree produced by LustreParser#recordUpdateExpr.
	VisitRecordUpdateExpr(ctx *RecordUpdateExprContext) interface{}

	// Visit a parse tree produced by LustreParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by LustreParser#baseEID.
	VisitBaseEID(ctx *BaseEIDContext) interface{}

	// Visit a parse tree produced by LustreParser#arrayEID.
	VisitArrayEID(ctx *ArrayEIDContext) interface{}

	// Visit a parse tree produced by LustreParser#recordEID.
	VisitRecordEID(ctx *RecordEIDContext) interface{}

}
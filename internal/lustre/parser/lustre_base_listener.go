// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lustre
import "github.com/antlr4-go/antlr/v4"

// BaseLustreListener is a complete listener for a parse tree produced by LustreParser.
type BaseLustreListener struct{}

var _ LustreListener = &BaseLustreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLustreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLustreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLustreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLustreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLustreListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLustreListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaseLustreListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaseLustreListener) ExitTypedef(ctx *TypedefContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseLustreListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseLustreListener) ExitConstant(ctx *ConstantContext) {}

// EnterNode is called when production node is entered.
func (s *BaseLustreListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseLustreListener) ExitNode(ctx *NodeContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseLustreListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseLustreListener) ExitFunction(ctx *FunctionContext) {}

// EnterVarDeclList is called when production varDeclList is entered.
func (s *BaseLustreListener) EnterVarDeclList(ctx *VarDeclListContext) {}

// ExitVarDeclList is called when production varDeclList is exited.
func (s *BaseLustreListener) ExitVarDeclList(ctx *VarDeclListContext) {}

// EnterVarDeclGroup is called when production varDeclGroup is entered.
func (s *BaseLustreListener) EnterVarDeclGroup(ctx *VarDeclGroupContext) {}

// ExitVarDeclGroup is called when production varDeclGroup is exited.
func (s *BaseLustreListener) ExitVarDeclGroup(ctx *VarDeclGroupContext) {}

// EnterPlainType is called when production plainType is entered.
func (s *BaseLustreListener) EnterPlainType(ctx *PlainTypeContext) {}

// ExitPlainType is called when production plainType is exited.
func (s *BaseLustreListener) ExitPlainType(ctx *PlainTypeContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BaseLustreListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BaseLustreListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BaseLustreListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BaseLustreListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseLustreListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseLustreListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterRealType is called when production realType is entered.
func (s *BaseLustreListener) EnterRealType(ctx *RealTypeContext) {}

// ExitRealType is called when production realType is exited.
func (s *BaseLustreListener) ExitRealType(ctx *RealTypeContext) {}

// EnterSubrangeType is called when production subrangeType is entered.
func (s *BaseLustreListener) EnterSubrangeType(ctx *SubrangeTypeContext) {}

// ExitSubrangeType is called when production subrangeType is exited.
func (s *BaseLustreListener) ExitSubrangeType(ctx *SubrangeTypeContext) {}

// EnterIntType is called when production intType is entered.
func (s *BaseLustreListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production intType is exited.
func (s *BaseLustreListener) ExitIntType(ctx *IntTypeContext) {}

// EnterUserType is called when production userType is entered.
func (s *BaseLustreListener) EnterUserType(ctx *UserTypeContext) {}

// ExitUserType is called when production userType is exited.
func (s *BaseLustreListener) ExitUserType(ctx *UserTypeContext) {}

// EnterBoolType is called when production boolType is entered.
func (s *BaseLustreListener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production boolType is exited.
func (s *BaseLustreListener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterBound is called when production bound is entered.
func (s *BaseLustreListener) EnterBound(ctx *BoundContext) {}

// ExitBound is called when production bound is exited.
func (s *BaseLustreListener) ExitBound(ctx *BoundContext) {}

// EnterStateMachine is called when production stateMachine is entered.
func (s *BaseLustreListener) EnterStateMachine(ctx *StateMachineContext) {}

// ExitStateMachine is called when production stateMachine is exited.
func (s *BaseLustreListener) ExitStateMachine(ctx *StateMachineContext) {}

// EnterStateDecl is called when production stateDecl is entered.
func (s *BaseLustreListener) EnterStateDecl(ctx *StateDeclContext) {}

// ExitStateDecl is called when production stateDecl is exited.
func (s *BaseLustreListener) ExitStateDecl(ctx *StateDeclContext) {}

// EnterTransition is called when production transition is entered.
func (s *BaseLustreListener) EnterTransition(ctx *TransitionContext) {}

// ExitTransition is called when production transition is exited.
func (s *BaseLustreListener) ExitTransition(ctx *TransitionContext) {}

// EnterArrow is called when production arrow is entered.
func (s *BaseLustreListener) EnterArrow(ctx *ArrowContext) {}

// ExitArrow is called when production arrow is exited.
func (s *BaseLustreListener) ExitArrow(ctx *ArrowContext) {}

// EnterFork is called when production fork is entered.
func (s *BaseLustreListener) EnterFork(ctx *ForkContext) {}

// ExitFork is called when production fork is exited.
func (s *BaseLustreListener) ExitFork(ctx *ForkContext) {}

// EnterElsifFork is called when production elsifFork is entered.
func (s *BaseLustreListener) EnterElsifFork(ctx *ElsifForkContext) {}

// ExitElsifFork is called when production elsifFork is exited.
func (s *BaseLustreListener) ExitElsifFork(ctx *ElsifForkContext) {}

// EnterElseFork is called when production elseFork is entered.
func (s *BaseLustreListener) EnterElseFork(ctx *ElseForkContext) {}

// ExitElseFork is called when production elseFork is exited.
func (s *BaseLustreListener) ExitElseFork(ctx *ElseForkContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseLustreListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseLustreListener) ExitTarget(ctx *TargetContext) {}

// EnterActions is called when production actions is entered.
func (s *BaseLustreListener) EnterActions(ctx *ActionsContext) {}

// ExitActions is called when production actions is exited.
func (s *BaseLustreListener) ExitActions(ctx *ActionsContext) {}

// EnterDataDef is called when production dataDef is entered.
func (s *BaseLustreListener) EnterDataDef(ctx *DataDefContext) {}

// ExitDataDef is called when production dataDef is exited.
func (s *BaseLustreListener) ExitDataDef(ctx *DataDefContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseLustreListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseLustreListener) ExitScope(ctx *ScopeContext) {}

// EnterLocalBlock is called when production localBlock is entered.
func (s *BaseLustreListener) EnterLocalBlock(ctx *LocalBlockContext) {}

// ExitLocalBlock is called when production localBlock is exited.
func (s *BaseLustreListener) ExitLocalBlock(ctx *LocalBlockContext) {}

// EnterEqs is called when production eqs is entered.
func (s *BaseLustreListener) EnterEqs(ctx *EqsContext) {}

// ExitEqs is called when production eqs is exited.
func (s *BaseLustreListener) ExitEqs(ctx *EqsContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseLustreListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseLustreListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarID is called when production varID is entered.
func (s *BaseLustreListener) EnterVarID(ctx *VarIDContext) {}

// ExitVarID is called when production varID is exited.
func (s *BaseLustreListener) ExitVarID(ctx *VarIDContext) {}

// EnterDefaultDecl is called when production defaultDecl is entered.
func (s *BaseLustreListener) EnterDefaultDecl(ctx *DefaultDeclContext) {}

// ExitDefaultDecl is called when production defaultDecl is exited.
func (s *BaseLustreListener) ExitDefaultDecl(ctx *DefaultDeclContext) {}

// EnterLastDecl is called when production lastDecl is entered.
func (s *BaseLustreListener) EnterLastDecl(ctx *LastDeclContext) {}

// ExitLastDecl is called when production lastDecl is exited.
func (s *BaseLustreListener) ExitLastDecl(ctx *LastDeclContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseLustreListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseLustreListener) ExitProperty(ctx *PropertyContext) {}

// EnterRealizabilityInputs is called when production realizabilityInputs is entered.
func (s *BaseLustreListener) EnterRealizabilityInputs(ctx *RealizabilityInputsContext) {}

// ExitRealizabilityInputs is called when production realizabilityInputs is exited.
func (s *BaseLustreListener) ExitRealizabilityInputs(ctx *RealizabilityInputsContext) {}

// EnterIvc is called when production ivc is entered.
func (s *BaseLustreListener) EnterIvc(ctx *IvcContext) {}

// ExitIvc is called when production ivc is exited.
func (s *BaseLustreListener) ExitIvc(ctx *IvcContext) {}

// EnterMain is called when production main is entered.
func (s *BaseLustreListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseLustreListener) ExitMain(ctx *MainContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseLustreListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseLustreListener) ExitAssertion(ctx *AssertionContext) {}

// EnterEquation is called when production equation is entered.
func (s *BaseLustreListener) EnterEquation(ctx *EquationContext) {}

// ExitEquation is called when production equation is exited.
func (s *BaseLustreListener) ExitEquation(ctx *EquationContext) {}

// EnterSimpleEquation is called when production simpleEquation is entered.
func (s *BaseLustreListener) EnterSimpleEquation(ctx *SimpleEquationContext) {}

// ExitSimpleEquation is called when production simpleEquation is exited.
func (s *BaseLustreListener) ExitSimpleEquation(ctx *SimpleEquationContext) {}

// EnterLhs is called when production lhs is entered.
func (s *BaseLustreListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BaseLustreListener) ExitLhs(ctx *LhsContext) {}

// EnterLhsID is called when production lhsID is entered.
func (s *BaseLustreListener) EnterLhsID(ctx *LhsIDContext) {}

// ExitLhsID is called when production lhsID is exited.
func (s *BaseLustreListener) ExitLhsID(ctx *LhsIDContext) {}

// EnterControlBlock is called when production controlBlock is entered.
func (s *BaseLustreListener) EnterControlBlock(ctx *ControlBlockContext) {}

// ExitControlBlock is called when production controlBlock is exited.
func (s *BaseLustreListener) ExitControlBlock(ctx *ControlBlockContext) {}

// EnterEmission is called when production emission is entered.
func (s *BaseLustreListener) EnterEmission(ctx *EmissionContext) {}

// ExitEmission is called when production emission is exited.
func (s *BaseLustreListener) ExitEmission(ctx *EmissionContext) {}

// EnterEmissionBody is called when production emissionBody is entered.
func (s *BaseLustreListener) EnterEmissionBody(ctx *EmissionBodyContext) {}

// ExitEmissionBody is called when production emissionBody is exited.
func (s *BaseLustreListener) ExitEmissionBody(ctx *EmissionBodyContext) {}

// EnterReturn is called when production return is entered.
func (s *BaseLustreListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production return is exited.
func (s *BaseLustreListener) ExitReturn(ctx *ReturnContext) {}

// EnterReturnVar is called when production returnVar is entered.
func (s *BaseLustreListener) EnterReturnVar(ctx *ReturnVarContext) {}

// ExitReturnVar is called when production returnVar is exited.
func (s *BaseLustreListener) ExitReturnVar(ctx *ReturnVarContext) {}

// EnterClockedBlock is called when production clockedBlock is entered.
func (s *BaseLustreListener) EnterClockedBlock(ctx *ClockedBlockContext) {}

// ExitClockedBlock is called when production clockedBlock is exited.
func (s *BaseLustreListener) ExitClockedBlock(ctx *ClockedBlockContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseLustreListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseLustreListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterMatchBlock is called when production matchBlock is entered.
func (s *BaseLustreListener) EnterMatchBlock(ctx *MatchBlockContext) {}

// ExitMatchBlock is called when production matchBlock is exited.
func (s *BaseLustreListener) ExitMatchBlock(ctx *MatchBlockContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseLustreListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseLustreListener) ExitPattern(ctx *PatternContext) {}

// EnterRecordExpr is called when production recordExpr is entered.
func (s *BaseLustreListener) EnterRecordExpr(ctx *RecordExprContext) {}

// ExitRecordExpr is called when production recordExpr is exited.
func (s *BaseLustreListener) ExitRecordExpr(ctx *RecordExprContext) {}

// EnterIntExpr is called when production intExpr is entered.
func (s *BaseLustreListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production intExpr is exited.
func (s *BaseLustreListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterArrayExpr is called when production arrayExpr is entered.
func (s *BaseLustreListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production arrayExpr is exited.
func (s *BaseLustreListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BaseLustreListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BaseLustreListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterRealExpr is called when production realExpr is entered.
func (s *BaseLustreListener) EnterRealExpr(ctx *RealExprContext) {}

// ExitRealExpr is called when production realExpr is exited.
func (s *BaseLustreListener) ExitRealExpr(ctx *RealExprContext) {}

// EnterIfThenElseExpr is called when production ifThenElseExpr is entered.
func (s *BaseLustreListener) EnterIfThenElseExpr(ctx *IfThenElseExprContext) {}

// ExitIfThenElseExpr is called when production ifThenElseExpr is exited.
func (s *BaseLustreListener) ExitIfThenElseExpr(ctx *IfThenElseExprContext) {}

// EnterBinaryExpr is called when production binaryExpr is entered.
func (s *BaseLustreListener) EnterBinaryExpr(ctx *BinaryExprContext) {}

// ExitBinaryExpr is called when production binaryExpr is exited.
func (s *BaseLustreListener) ExitBinaryExpr(ctx *BinaryExprContext) {}

// EnterPreExpr is called when production preExpr is entered.
func (s *BaseLustreListener) EnterPreExpr(ctx *PreExprContext) {}

// ExitPreExpr is called when production preExpr is exited.
func (s *BaseLustreListener) ExitPreExpr(ctx *PreExprContext) {}

// EnterRecordAccessExpr is called when production recordAccessExpr is entered.
func (s *BaseLustreListener) EnterRecordAccessExpr(ctx *RecordAccessExprContext) {}

// ExitRecordAccessExpr is called when production recordAccessExpr is exited.
func (s *BaseLustreListener) ExitRecordAccessExpr(ctx *RecordAccessExprContext) {}

// EnterNegateExpr is called when production negateExpr is entered.
func (s *BaseLustreListener) EnterNegateExpr(ctx *NegateExprContext) {}

// ExitNegateExpr is called when production negateExpr is exited.
func (s *BaseLustreListener) ExitNegateExpr(ctx *NegateExprContext) {}

// EnterCondactExpr is called when production condactExpr is entered.
func (s *BaseLustreListener) EnterCondactExpr(ctx *CondactExprContext) {}

// ExitCondactExpr is called when production condactExpr is exited.
func (s *BaseLustreListener) ExitCondactExpr(ctx *CondactExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseLustreListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseLustreListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterArrayAccessExpr is called when production arrayAccessExpr is entered.
func (s *BaseLustreListener) EnterArrayAccessExpr(ctx *ArrayAccessExprContext) {}

// ExitArrayAccessExpr is called when production arrayAccessExpr is exited.
func (s *BaseLustreListener) ExitArrayAccessExpr(ctx *ArrayAccessExprContext) {}

// EnterArrayUpdateExpr is called when production arrayUpdateExpr is entered.
func (s *BaseLustreListener) EnterArrayUpdateExpr(ctx *ArrayUpdateExprContext) {}

// ExitArrayUpdateExpr is called when production arrayUpdateExpr is exited.
func (s *BaseLustreListener) ExitArrayUpdateExpr(ctx *ArrayUpdateExprContext) {}

// EnterBoolExpr is called when production boolExpr is entered.
func (s *BaseLustreListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production boolExpr is exited.
func (s *BaseLustreListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseLustreListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseLustreListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterTupleExpr is called when production tupleExpr is entered.
func (s *BaseLustreListener) EnterTupleExpr(ctx *TupleExprContext) {}

// ExitTupleExpr is called when production tupleExpr is exited.
func (s *BaseLustreListener) ExitTupleExpr(ctx *TupleExprContext) {}

// EnterRecordUpdateExpr is called when production recordUpdateExpr is entered.
func (s *BaseLustreListener) EnterRecordUpdateExpr(ctx *RecordUpdateExprContext) {}

// ExitRecordUpdateExpr is called when production recordUpdateExpr is exited.
func (s *BaseLustreListener) ExitRecordUpdateExpr(ctx *RecordUpdateExprContext) {}

// EnterIdExpr is called when production idExpr is entered.
func (s *BaseLustreListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production idExpr is exited.
func (s *BaseLustreListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterBaseEID is called when production baseEID is entered.
func (s *BaseLustreListener) EnterBaseEID(ctx *BaseEIDContext) {}

// ExitBaseEID is called when production baseEID is exited.
func (s *BaseLustreListener) ExitBaseEID(ctx *BaseEIDContext) {}

// EnterArrayEID is called when production arrayEID is entered.
func (s *BaseLustreListener) EnterArrayEID(ctx *ArrayEIDContext) {}

// ExitArrayEID is called when production arrayEID is exited.
func (s *BaseLustreListener) ExitArrayEID(ctx *ArrayEIDContext) {}

// EnterRecordEID is called when production recordEID is entered.
func (s *BaseLustreListener) EnterRecordEID(ctx *RecordEIDContext) {}

// ExitRecordEID is called when production recordEID is exited.
func (s *BaseLustreListener) ExitRecordEID(ctx *RecordEIDContext) {}

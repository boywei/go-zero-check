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

// EnterLhs is called when production lhs is entered.
func (s *BaseLustreListener) EnterLhs(ctx *LhsContext) {}

// ExitLhs is called when production lhs is exited.
func (s *BaseLustreListener) ExitLhs(ctx *LhsContext) {}

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

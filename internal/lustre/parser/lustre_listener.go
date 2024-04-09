// Code generated from /Users/wei/GoProjects/go-zero-check/internal/lustre/Lustre.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lustre
import "github.com/antlr4-go/antlr/v4"


// LustreListener is a complete listener for a parse tree produced by LustreParser.
type LustreListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypedef is called when entering the typedef production.
	EnterTypedef(c *TypedefContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterVarDeclList is called when entering the varDeclList production.
	EnterVarDeclList(c *VarDeclListContext)

	// EnterVarDeclGroup is called when entering the varDeclGroup production.
	EnterVarDeclGroup(c *VarDeclGroupContext)

	// EnterPlainType is called when entering the plainType production.
	EnterPlainType(c *PlainTypeContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterRealType is called when entering the realType production.
	EnterRealType(c *RealTypeContext)

	// EnterSubrangeType is called when entering the subrangeType production.
	EnterSubrangeType(c *SubrangeTypeContext)

	// EnterIntType is called when entering the intType production.
	EnterIntType(c *IntTypeContext)

	// EnterUserType is called when entering the userType production.
	EnterUserType(c *UserTypeContext)

	// EnterBoolType is called when entering the boolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterBound is called when entering the bound production.
	EnterBound(c *BoundContext)

	// EnterProperty is called when entering the property production.
	EnterProperty(c *PropertyContext)

	// EnterRealizabilityInputs is called when entering the realizabilityInputs production.
	EnterRealizabilityInputs(c *RealizabilityInputsContext)

	// EnterIvc is called when entering the ivc production.
	EnterIvc(c *IvcContext)

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterEquation is called when entering the equation production.
	EnterEquation(c *EquationContext)

	// EnterLhs is called when entering the lhs production.
	EnterLhs(c *LhsContext)

	// EnterRecordExpr is called when entering the recordExpr production.
	EnterRecordExpr(c *RecordExprContext)

	// EnterIntExpr is called when entering the intExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterArrayExpr is called when entering the arrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterRealExpr is called when entering the realExpr production.
	EnterRealExpr(c *RealExprContext)

	// EnterIfThenElseExpr is called when entering the ifThenElseExpr production.
	EnterIfThenElseExpr(c *IfThenElseExprContext)

	// EnterBinaryExpr is called when entering the binaryExpr production.
	EnterBinaryExpr(c *BinaryExprContext)

	// EnterPreExpr is called when entering the preExpr production.
	EnterPreExpr(c *PreExprContext)

	// EnterRecordAccessExpr is called when entering the recordAccessExpr production.
	EnterRecordAccessExpr(c *RecordAccessExprContext)

	// EnterNegateExpr is called when entering the negateExpr production.
	EnterNegateExpr(c *NegateExprContext)

	// EnterCondactExpr is called when entering the condactExpr production.
	EnterCondactExpr(c *CondactExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterArrayAccessExpr is called when entering the arrayAccessExpr production.
	EnterArrayAccessExpr(c *ArrayAccessExprContext)

	// EnterArrayUpdateExpr is called when entering the arrayUpdateExpr production.
	EnterArrayUpdateExpr(c *ArrayUpdateExprContext)

	// EnterBoolExpr is called when entering the boolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterTupleExpr is called when entering the tupleExpr production.
	EnterTupleExpr(c *TupleExprContext)

	// EnterRecordUpdateExpr is called when entering the recordUpdateExpr production.
	EnterRecordUpdateExpr(c *RecordUpdateExprContext)

	// EnterIdExpr is called when entering the idExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterBaseEID is called when entering the baseEID production.
	EnterBaseEID(c *BaseEIDContext)

	// EnterArrayEID is called when entering the arrayEID production.
	EnterArrayEID(c *ArrayEIDContext)

	// EnterRecordEID is called when entering the recordEID production.
	EnterRecordEID(c *RecordEIDContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypedef is called when exiting the typedef production.
	ExitTypedef(c *TypedefContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitVarDeclList is called when exiting the varDeclList production.
	ExitVarDeclList(c *VarDeclListContext)

	// ExitVarDeclGroup is called when exiting the varDeclGroup production.
	ExitVarDeclGroup(c *VarDeclGroupContext)

	// ExitPlainType is called when exiting the plainType production.
	ExitPlainType(c *PlainTypeContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitRealType is called when exiting the realType production.
	ExitRealType(c *RealTypeContext)

	// ExitSubrangeType is called when exiting the subrangeType production.
	ExitSubrangeType(c *SubrangeTypeContext)

	// ExitIntType is called when exiting the intType production.
	ExitIntType(c *IntTypeContext)

	// ExitUserType is called when exiting the userType production.
	ExitUserType(c *UserTypeContext)

	// ExitBoolType is called when exiting the boolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitBound is called when exiting the bound production.
	ExitBound(c *BoundContext)

	// ExitProperty is called when exiting the property production.
	ExitProperty(c *PropertyContext)

	// ExitRealizabilityInputs is called when exiting the realizabilityInputs production.
	ExitRealizabilityInputs(c *RealizabilityInputsContext)

	// ExitIvc is called when exiting the ivc production.
	ExitIvc(c *IvcContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitEquation is called when exiting the equation production.
	ExitEquation(c *EquationContext)

	// ExitLhs is called when exiting the lhs production.
	ExitLhs(c *LhsContext)

	// ExitRecordExpr is called when exiting the recordExpr production.
	ExitRecordExpr(c *RecordExprContext)

	// ExitIntExpr is called when exiting the intExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitArrayExpr is called when exiting the arrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitRealExpr is called when exiting the realExpr production.
	ExitRealExpr(c *RealExprContext)

	// ExitIfThenElseExpr is called when exiting the ifThenElseExpr production.
	ExitIfThenElseExpr(c *IfThenElseExprContext)

	// ExitBinaryExpr is called when exiting the binaryExpr production.
	ExitBinaryExpr(c *BinaryExprContext)

	// ExitPreExpr is called when exiting the preExpr production.
	ExitPreExpr(c *PreExprContext)

	// ExitRecordAccessExpr is called when exiting the recordAccessExpr production.
	ExitRecordAccessExpr(c *RecordAccessExprContext)

	// ExitNegateExpr is called when exiting the negateExpr production.
	ExitNegateExpr(c *NegateExprContext)

	// ExitCondactExpr is called when exiting the condactExpr production.
	ExitCondactExpr(c *CondactExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitArrayAccessExpr is called when exiting the arrayAccessExpr production.
	ExitArrayAccessExpr(c *ArrayAccessExprContext)

	// ExitArrayUpdateExpr is called when exiting the arrayUpdateExpr production.
	ExitArrayUpdateExpr(c *ArrayUpdateExprContext)

	// ExitBoolExpr is called when exiting the boolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitTupleExpr is called when exiting the tupleExpr production.
	ExitTupleExpr(c *TupleExprContext)

	// ExitRecordUpdateExpr is called when exiting the recordUpdateExpr production.
	ExitRecordUpdateExpr(c *RecordUpdateExprContext)

	// ExitIdExpr is called when exiting the idExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitBaseEID is called when exiting the baseEID production.
	ExitBaseEID(c *BaseEIDContext)

	// ExitArrayEID is called when exiting the arrayEID production.
	ExitArrayEID(c *ArrayEIDContext)

	// ExitRecordEID is called when exiting the recordEID production.
	ExitRecordEID(c *RecordEIDContext)
}

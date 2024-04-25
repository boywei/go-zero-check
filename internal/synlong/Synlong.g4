grammar Synlong;

program: (typedef | constant | node | function)* EOF;

typedef: 'type' ID '=' topLevelType ';';

constant: 'const' ID (':' type)? '=' expr ';';

node:
  'node' ID '(' input=varDeclList? ')'
  'returns' '(' output=varDeclList? ')'
  ('var' local=varDeclList ';')?
  'let'
    (equation | property | assertion | main | realizabilityInputs | ivc | stateMachine)*
  'tel' ';'?
;

function:
  'function' eID '(' input=varDeclList? ')'
  'returns' '(' output=varDeclList? ')' ';'
;

varDeclList: varDeclGroup (';' varDeclGroup)*;

varDeclGroup: eID (',' eID)* ':' type;

topLevelType: type                                       # plainType
    | 'struct' '{' (ID ':' type) (';' ID ':' type)* '}'  # recordType
    | 'enum' '{' ID (',' ID)* '}'                        # enumType
    ;

type: 'int'                                              # intType
    | 'subrange' '[' bound ',' bound ']' 'of' 'int'      # subrangeType
    | 'bool'                                             # boolType
    | 'real'                                             # realType
    | type '[' INT ']'                                   # arrayType
    | ID                                                 # userType
    ;

bound: '-'? INT;

// 4.5.3 状态机
stateMachine: 'automaton' ID stateDecl*;

stateDecl:
  'initial'? 'final'? 'state' ID
  ('unless' (transition ';')+)?
  dataDef
  ('until' (transition ';')*
  ('synchro' actions? 'fork' ';')?)?
;

transition: 'if' expr arrow;

arrow: actions? fork;

fork: target
  | 'if' expr arrow elsifFork* elseFork?
;

elsifFork: 'elsif' expr arrow;

elseFork: 'else' arrow;

target: 'restart' ID
  | 'resume' ID
;

actions: 'do' dataDef
  | 'do' '{' 'emit'? emissionBody (';' 'emit' emissionBody)* '}'
;

dataDef: equation | scope;

scope: localBlock? eqs?;

localBlock: 'var' (varDecl ';')*;

eqs: 'let' (equation ';')* 'tel';

varDecl: varID (',' varID)* ':' type
//  whenDecl?
  defaultDecl?
  lastDecl?
;

varID: 'clock'? ID;

//whenDecl: 'when' clockExpr;

defaultDecl: 'default' '=' expr;

lastDecl: 'last' '=' expr;

property: '--%PROPERTY' eID ';';

realizabilityInputs: '--%REALIZABLE' (ID (',' ID)*)? ';';

ivc: '--%IVC' (eID (',' eID)*)? ';';

main: '--%MAIN' ';'?;

assertion: 'assert' expr ';';

// 4.5.1 等式
//equation: (lhs | '(' lhs? ')') '=' expr ';';
equation: simpleEquation
  | emission
  | controlBlock return
;

simpleEquation: lhs '=' expr;

lhs: '(' ')'
  | lhsID (',' lhsID)*
;

lhsID: eID | '_';

controlBlock: stateMachine | clockedBlock;

emission: 'emit' emissionBody;

emissionBody: eID ('if' expr)?
  | '(' eID (',' eID)* ')' ('if' expr)?
;

return: 'returns' returnVar ';';

returnVar: (eID ',')* (eID | '..');

// 4.5.2 条件块
clockedBlock: 'activate' eID (ifBlock | matchBlock);

ifBlock: 'if' expr 'then' (dataDef | ifBlock)
  'else' (dataDef | ifBlock)
;

matchBlock: 'when' expr 'match' ('|' pattern ':' dataDef)+;

pattern: ID
  | INT
  | REAL
  | BOOL
;

expr: ID                                                       # idExpr
    | INT                                                      # intExpr
    | REAL                                                     # realExpr
    | BOOL                                                     # boolExpr
    | op=('real' | 'floor') '(' expr ')'                       # castExpr
    | eID '(' (expr (',' expr)*)? ')'                          # callExpr
    | 'condact' '(' expr (',' expr)+ ')'                       # condactExpr
    | expr '.' ID                                              # recordAccessExpr
    | expr '{' ID ':=' expr '}'                                # recordUpdateExpr
    | expr '[' expr ']'                                        # arrayAccessExpr
    | expr '[' expr ':=' expr ']'                              # arrayUpdateExpr
    | 'pre' expr                                               # preExpr
    | 'not' expr                                               # notExpr
    | '-' expr                                                 # negateExpr
    | expr op=('*' | '/' | 'div' | 'mod') expr                 # binaryExpr
    | expr op=('+' | '-') expr                                 # binaryExpr
    | expr op=('<' | '<=' | '>' | '>=' | '=' | '<>') expr      # binaryExpr
    | expr op='and' expr                                       # binaryExpr
    | expr op=('or' | 'xor') expr                              # binaryExpr
    | <assoc=right> expr op='=>' expr                          # binaryExpr
    | <assoc=right> expr op='->' expr                          # binaryExpr
    | 'if' expr 'then' expr 'else' expr                        # ifThenElseExpr
    | ID '{' ID '=' expr (';' ID '=' expr)* '}'                # recordExpr
    | '[' expr (',' expr)* ']'                                 # arrayExpr
    | '(' expr (',' expr)* ')'                                 # tupleExpr
    ;

// eID used internally. Users should only use ID.
eID: ID                                                        # baseEID
   | eID '[' INT ']'                                           # arrayEID
   | eID '.' ID                                                # recordEID
   ;

REAL: INT '.' INT;

BOOL: 'true' | 'false';
INT: [0-9]+;

// ~ is used internally. Users should not use it.
ID: [a-zA-Z_~!][a-zA-Z_0-9~!]*;

WS: [ \t\n\r\f]+ -> skip;

SL_COMMENT: '--' (~[%\n\r] ~[\n\r]* | /* empty */) ('\r'? '\n')? -> skip;
ML_COMMENT: '(*' .*? '*)' -> skip;

ERROR: .;
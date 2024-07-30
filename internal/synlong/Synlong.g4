grammar Synlong;

// L2C程序由一个或多个声明组成
program         : decls+ ;

// 声明可以是类型块、常量块或用户操作声明
decls           : type_block
                | const_block
                | user_op_decl ;

// 类型块以“type”关键字开始，包括一个或多个类型声明
type_block      : 'type' '{' type_decl (';' type_decl)* '}' ;

// 类型声明可以是标识符或标识符等于类型定义
type_decl       : ID ('=' type_def)? ;

// 类型定义可以是类型表达式或枚举
type_def        : type_expr
                | 'enum' '{' ID (' ' ID)* '}' ;

// 类型表达式可以是基本类型、类型变量、结构体类型或带大小的数组类型
type_expr       : 'bool'
                | 'char'
                | 'byte'
                | 'ubyte'
                | 'short'
                | 'ushort'
                | 'int'
                | 'uint'
                | 'long'
                | 'ulong'
                | 'float'
                | 'real'
                | typevar
                | '{' field_decl (';' field_decl)* '}'
                | type_expr '^' const_expr ;

// 字段声明包括标识符和类型表达式
field_decl      : ID ':' type_expr ;

// 类型变量是标识符
typevar         : ID ;

// 常量块以“const”关键字开始，包括一个或多个常量声明
const_block     : 'const' '{' const_decl (';' const_decl)* '}' ;

// 常量声明包括标识符、类型表达式和可选的常量表达式
const_decl      : ID ':' type_expr ('=' const_expr)? ;

// 常量表达式可以是标识符、原子常量、单目运算符、双目运算符、列表或结构体
const_expr      : ID
                | atom
                | unary_arith_op const_expr
                | const_expr bin_arith_op const_expr
                | const_expr bin_bool_op const_expr
                | const_expr bin_relation_op const_expr
                | '[' const_list ']'
                | '{' const_label_expr (';' const_label_expr)* '}' ;

// 常量列表是一个或多个常量表达式
const_list      : const_expr (' ' const_expr)* ;

// 常量标签表达式包括标识符和常量表达式
const_label_expr: ID ':' const_expr ;

// 变量声明包括一个或多个变量标识符、类型表达式和可选的时钟和初始值声明
var_decls       : var_id (' ' var_id)* ':' type_expr (when_decl)? (last_decl)? ;

// 变量标识符是标识符
var_id          : ID ;

// 时钟声明以“when”关键字开始
when_decl       : 'when' clock_expr ;

// 时钟表达式是标识符或“not”关键字和标识符
clock_expr      : ID | 'not' ID ;

// 初始值声明以“last”关键字开始
last_decl       : 'last' '=' const_expr ;

// 用户操作声明包括操作种类、标识符、参数、返回类型和操作主体
user_op_decl    : op_kind ID params 'returns' params op_body ;

// 操作种类可以是“function”或“node”
op_kind         : 'function' | 'node' ;

// 参数是用括号括起来的一个或多个变量声明
params          : '(' (var_decls (';' var_decls)*)? ')' ;

// 操作主体可以是分号或局部变量块和方程块
op_body         : ';' | (local_block)? 'let' (equation (';' equation)*) 'tel' ';'? ;

// 局部变量块以“var”关键字开始，包括一个或多个变量声明
local_block     : 'var' '{' var_decls (';' var_decls)* '}' ;

// 方程包括左值和表达式或状态机和返回变量
equation        : lhs '=' expr
                | state_machine 'return' returns_var ;

// 左值可以是空括号或一个或多个左值标识符
lhs             : '(' ')' | lhs_id (' ' lhs_id)* ;

// 左值标识符是标识符
lhs_id          : ID ;

// 返回变量是一个或多个标识符
returns_var     : ID (' ' ID)* ('..')? ;

// 状态机以“automaton”关键字开始，包括一个或多个状态声明
state_machine   : 'automaton' (ID)? '{' state_decl+ '}' ;

// 状态声明可以是初始状态、最终状态、状态标识符和状态转换
state_decl      : ('initial')? ('final')? 'state' ID ('unless' '{' transition (';' transition)* '}' )? data_def ('until' '{' transition (';' transition)* '}' )? ;

// 数据定义可以是方程或局部变量块和方程块
data_def        : equation ';'
                | (local_block)? 'let' (equation (';' equation)*) 'tel' ;

// 状态转换以“if”关键字开始，包括表达式和目标状态
transition      : 'if' expr ('resume' | 'restart') ID ;

// 时钟块以“activate”关键字开始，包括“if”块或“match”块
clocked_block   : 'activate' (ID)? (if_block | match_block) ;

// “if”块以“if”关键字开始，包括表达式、数据定义或嵌套的“if”块
if_block        : 'if' expr 'then' (data_def | if_block) 'else' (data_def | if_block) ;

// “match”块以“when”关键字开始，包括“match”关键字、模式和数据定义
match_block     : 'when' expr 'match' ( '|' pattern ':' data_def )+ ;

// 表达式可以是简单表达式、时钟表达式、布尔表达式、数组表达式、结构体表达式或应用表达式
expr            : simple_expr
                | 'last' '\'' ID
                | tempo_expr
                | bool_expr
                | array_expr
                | struct_expr
                | mixed_constructor
                | switch_expr
                | apply_expr ;

// 列表是一个或多个简单表达式
list            : simple_expr (' ' simple_expr)* ;

// 时钟表达式可以是“pre”关键字和简单表达式或一系列其他时钟操作
tempo_expr      : 'pre' simple_expr
                | simple_expr '->' simple_expr
                | 'fby' '(' simple_expr ';' const_expr ';' simple_expr ')'
                | simple_expr 'fby' simple_expr
                | simple_expr 'when' clock_expr
                | 'merge' ID '(' simple_expr ')' '(' simple_expr ')' ;

// 布尔表达式以“#”关键字开始，包括列表
bool_expr       : '#' '(' list ')' ;

// 数组表达式包括简单表达式和范围或其他数组操作
array_expr      : simple_expr '[' INTEGER '..' INTEGER ']'
                | '(' simple_expr '.' (' ' index)+ 'default' simple_expr ')'
                | simple_expr '^' const_expr
                | '[' list ']' ;

// 结构体表达式以大括号括起来，包括一个或多个标签表达式
struct_expr     : '{' label_expr (' ' label_expr)* '}' ;

// 混合构造器表达式以“with”关键字开始，包括标签或索引和简单表达式
mixed_constructor : '(' ID 'with' (' ' label_or_index '+')* '=' simple_expr ')' ;

// 标签表达式包括标识符和简单表达式
label_expr      : ID ':' simple_expr ;

// 索引是用方括号括起来的简单表达式
index           : '[' simple_expr ']' ;

// 标签或索引可以是点符号和标识符或索引
label_or_index  : '.' ID | index ;

// 切换表达式以“if”关键字开始，包括简单表达式和分支
switch_expr     : 'if' simple_expr 'then' simple_expr 'else' simple_expr
                | '(' 'case' simple_expr 'of' ( case_expr )+ ')' ;

// 分支表达式以竖线和模式开始，包括简单表达式
case_expr       : '|' pattern ':' simple_expr ;

// 模式可以是标识符、字符、整数、布尔值或下划线
pattern         : ID
                | CHAR
                | '-'? INTEGER
                | 'true'
                | 'false'
                | '_' ;

// 应用表达式包括前缀运算符和列表或其他应用操作
apply_expr      : prefix_operator '(' list ')'
                | '(' iterator '<<' prefix_operator ';' const_expr '>>' ')' '(' list ')'
                | '(' 'mapw' '<<' prefix_operator ';' const_expr '>>' 'if' simple_expr 'default' '(' list ')' ')' '(' list ')'
                | '(' 'mapwi' '<<' prefix_operator ';' const_expr '>>' 'if' simple_expr 'default' '(' list ')' ')' '(' list ')'
                | '(' 'foldw' '<<' prefix_operator ';' const_expr '>>' 'if' simple_expr ')' '(' list ')'
                | '(' 'foldwi' '<<' prefix_operator ';' const_expr '>>' 'if' simple_expr ')' '(' list ')' ;

// 前缀运算符可以是标识符或其他前缀操作符
prefix_operator : ID
                | prefix_unary_operator
                | prefix_binary_operator
                | '(make' ID ')'
                | '(flatten' ID ')' ;

// 前缀单目运算符
prefix_unary_operator : '+$' | '-$' | 'not$' | 'short$' | 'int$' | 'float$' | 'real$' ;

// 前缀双目运算符
prefix_binary_operator: '$+$' | '$-$' | '$*$' | '$/$' | '$mod$' | '$div$'
                      | '$=$' | '$<>$' | '$<$' | '$>$' | '$<=$' | '$>=$'
                      | '$and$' | '$or$' | '$xor$' ;

// 迭代器可以是“map”、“fold”或其他迭代操作
iterator        : 'map' | 'fold' | 'mapi' | 'foldi' | 'mapfold' ;

// 简单表达式可以是标识符、原子、索引操作、字段访问或其他简单运算
simple_expr     : ID
                | atom
                | simple_expr '[' const_expr ']'
                | simple_expr '.' ID
                | unary_arith_op simple_expr
                | simple_expr bin_arith_op simple_expr
                | simple_expr bin_bool_op simple_expr
                | simple_expr bin_relation_op simple_expr
                | '(' type_expr simple_expr ')' ;

// 单目算术运算符可以是加号、减号或“not”关键字
unary_arith_op  : '-' | '+' | 'not' ;

// 双目算术运算符可以是加号、减号、乘号、除号、模运算或整除
bin_arith_op    : '+' | '-' | '*' | '/' | 'mod' | 'div' ;

// 关系运算符可以是等号、不等号、小于、大于、小于等于或大于等于
bin_relation_op : '=' | '<>' | '<' | '>' | '<=' | '>=' ;

// 布尔运算符可以是与、或或异或
bin_bool_op     : 'and' | 'or' | 'xor' ;

// 原子常量可以是真、假、字符、整数、无符号整数、浮点数、实数、无符号短整数或短整数
atom            : 'true'
                | 'false'
                | CHAR
                | INTEGER
                | UINT
                | FLOAT
                | REAL
                | USHORT
                | SHORT ;

// 数字和标识符的定义
DIGIT10         : [0-9] ;
EXPONENT        : [eE] [+-]? DIGIT10+ ;
FLOAT           : DIGIT10+ '.' DIGIT10* EXPONENT? 'f' ;
REAL            : DIGIT10+ '.' DIGIT10* EXPONENT? ;
INTEGER         : '0' | [1-9] DIGIT10* ;
UINT            : INTEGER 'u' ;
USHORT          : INTEGER 'us' ;
SHORT           : INTEGER 's' ;
LETTER          : [a-zA-Z] ;
CHAR            : LETTER | [+-] ;
ID              : LETTER (DIGIT10 | LETTER | '_')* ;

// 空白字符和注释的处理
WS              : [ \t\r\n\f]+ -> skip ;
SINGLELINE_COMMENT: '--' (~[%\n\r] ~[\n\r]* | /* empty */) ('\r'? '\n')? -> skip;
MULTILINE_COMMENT: '(*' .*? '*)' -> skip;

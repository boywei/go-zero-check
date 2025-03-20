package synlong

import (
	"fmt"
	gen "github.com/boywei/go-zero-check/internal/synlong/gen"

	"github.com/antlr4-go/antlr/v4"
)

type CustomSyntaxError struct {
	line, column int
	msg          string
}

func (c CustomSyntaxError) Error() string {
	//TODO implement me
	fmt.Println("CustomSyntaxError")
	return fmt.Sprintf("line %d:%d %s", c.line, c.column, c.msg)
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // embed default which ensures we fit the interface
	errors                      []error
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errors = append(c.errors, &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

// CheckSynLong 通过antlr4生成的文件来检验输入的文件是否符合规范
func CheckSynLong(file string) (string, error) {
	// 创建输入流
	inputStream := antlr.NewInputStream(file)

	// 创建词法分析器
	lexer := gen.NewSynlongLexer(inputStream)
	lexerErrors := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 创建解析器
	parser := gen.NewSynlongParser(stream)
	parserErrors := &CustomErrorListener{}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrors)

	// 解析输入
	parser.BuildParseTrees = true
	parseTree := parser.Program()

	// TODO:遍历
	//antlr.ParseTreeWalkerDefault.Walk(antlr.NewTraceListener(parser), parseTree)

	fmt.Println(parseTree.ToStringTree(nil, parser))

	// 检查是否有语法错误
	if parser.HasError() {
		return parser.GetError().GetMessage(), nil
	}

	return "成功识别语法!", nil
}

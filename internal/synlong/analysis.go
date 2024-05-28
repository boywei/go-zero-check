package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/boywei/go-zero-check/internal/lustre/parser"
	log "github.com/sirupsen/logrus"
)

type LustreListener struct {
	*parser.BaseLustreListener
	p *parser.LustreParser
	t antlr.Tree
}

func NewLustreListener(p *parser.LustreParser, t antlr.Tree) *LustreListener {
	return &LustreListener{
		p: p,
		t: t,
	}
}

func (this *LustreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := this.p.RuleNames[i]
	fmt.Printf("==> %s 《 %s 》\n", ruleName, ctx.GetText())
}

//func (this *LustreListener) EnterNode(ctx *parser.NodeContext) {
//	fmt.Println("Enter node:")
//	fmt.Println(ctx.GetText())
//}

func printLevelPrefix(ctx antlr.ParserRuleContext) {
	level := 0
	t := ctx.GetParent()
	for t != nil {
		level++
		fmt.Print(" ")
		t = t.GetParent()
	}
}

func main() {
	//input := "node Multiply(x: real; y: real) returns (z: real); let z = x * y; tel"
	//inputStream := antlr.NewInputStream(input)
	filePath := "internal/lustre/demo/test.lus"
	inputStream, err := antlr.NewFileStream(filePath)
	if err != nil {
		log.Error("open file error: ", err)
	}
	lexer := parser.NewLustreLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLustreParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// 添加一个监听器并开始解析
	tree := p.Program() // 解析整个程序(从program开始), 而不是node开始
	listener := NewLustreListener(p, tree)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
}

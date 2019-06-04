// 简易四则运算计算器
package calc

import (
	"github.com/lwlwilliam/go/compiling/calc/eval"
	lexer2 "github.com/lwlwilliam/go/compiling/calc/lexer"
	parser2 "github.com/lwlwilliam/go/compiling/calc/parser"
)

func Calc(input string) int64 {
	lexer := lexer2.NewLexer(input)
	parser := parser2.NewParser(lexer)

	exp := parser.ParseExpression(parser2.LOWEST)
	return eval.Eval(exp)
}
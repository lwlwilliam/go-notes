// 语法树
// 比如 1+1 这个是正确语法，1+ 就不是了，最终这些语法会解释成树状，如下所示
//   	  +
// 		/   \
// 		1   1
package ast

import (
	"bytes"
	"github.com/lwlwilliam/go/compiling/calc/lexer"
)

// 表达式
type Expression interface {
	String() string
}

// 数字表达式
type IntegerLiteralExpression struct {
	Token lexer.Token
	Value int64
}

// 数字表达式的字符串表达形式
func (il *IntegerLiteralExpression) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    lexer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    lexer.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" ")
	out.WriteString(ie.Operator)
	out.WriteString(" ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

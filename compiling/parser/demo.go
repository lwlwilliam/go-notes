package main

import (
	"go/token"
	"fmt"
	"go/ast"
	"go/parser"
)

func main() {
	expr := `a == 1 && b == 2`
	expr = `a == 1 && b == 2 && in_array(c, []int{1, 2, 3, 4})`

	fset := token.NewFileSet()
	exprAst, err := parser.ParseExpr(expr)
	if err != nil {
		fmt.Println(err)
		return
	}
	ast.Print(fset, exprAst)
}

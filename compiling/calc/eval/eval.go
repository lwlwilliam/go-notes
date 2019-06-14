package eval

import "github.com/lwlwilliam/go-notes/compiling/calc/ast"

func Eval(exp ast.Expression) int64 {
	switch node := exp.(type) {
	case *ast.IntegerLiteralExpression:
		return node.Value
	case *ast.PrefixExpression:
		rightV := Eval(node.Right)
		return evalPrefixExpression(node.Operator, rightV)
	case *ast.InfixExpression:
		leftV := Eval(node.Left)
		rightV := Eval(node.Right)
		return evalInfixExpression(leftV, node.Operator, rightV)
	}

	return 0
}

func evalPrefixExpression(operator string, right int64) int64 {
	if operator != "-" {
		return 0
	}
	return -right
}

func evalInfixExpression(left int64, operator string, right int64) int64 {
	switch operator {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		if right != 0 {
			return left / right
		} else {
			return 0
		}
	default:
		return 0
	}
}

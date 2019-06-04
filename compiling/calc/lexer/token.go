// token，词元
package lexer

// 四则运算的关键字
// +-*/() 以及数字
const (
	ILLEGAL = "ILLEGAL" // 非法字符
	EOF     = "EOF"
	INT     = "INT"

	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	LPAREN = "("
	RPAREN = ")"
)

// Token 的类型及字面量
// 示例：Token{Type: INT, Literal: "66"}
type Token struct {
	Type    string
	Literal string
}

func newToken(tokenType string, c byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(c),
	}
}

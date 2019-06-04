package lexer

import "testing"

func TestTokenizer(t *testing.T) {
	input := `(5 + -10 * 2 + 15 / 3) * 2`
	tests := []struct {
		expectedType    string
		expectedLiteral string
	}{
		{LPAREN, "("},
		{INT, "5"},
		{PLUS, "+"},
		{MINUS, "-"},
		{INT, "10"},
		{ASTERISK, "*"},
		{INT, "2"},
		{PLUS, "+"},
		{INT, "15"},
		{SLASH, "/"},
		{INT, "3"},
		{RPAREN, ")"},
		{ASTERISK, "*"},
		{INT, "2"},
	}

	l := NewLexer(input)

	for k, v := range tests {
		tok := l.NextToken()

		if tok.Type != v.expectedType {
			t.Fatalf("tests[%d] - type wrong. expected=%q, got=%q",
				k, v.expectedType, tok.Type)
		}

		if tok.Literal != v.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				k, v.expectedLiteral, tok.Literal)
		}
	}
}

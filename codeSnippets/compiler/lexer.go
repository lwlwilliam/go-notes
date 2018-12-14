// 参照 https://play.rust-lang.org/?gist=070c3b6b985098a306c62881d7f2f82c&version=stable&mode=debug&edition=2015 所写的一个词法分析器 demo
package complier

import (
	"unicode"
	"fmt"
)

type Token []string

func tokenize(s string) Token {
	tokens	:= Token{}
	chars	:= []byte(s)
	len		:= len(chars)
	i		:= 0

	for i < len {
		switch chars[i] {
		case '+':
			tokens = append(tokens, "Plus")
			i ++
		default:
			if unicode.IsDigit(rune(chars[i])) {
				num := []byte{chars[i]}
				i ++

				for i < len && unicode.IsDigit(rune(chars[i])) {
					num = append(num, chars[i])
					i ++
				}

				tokens = append(tokens, string(num))
			} else {
				i ++
			}
		}

		i ++
	}

	return tokens
}

func main() {
	tokens := tokenize("12 + 3")
	fmt.Println(tokens)

	tokens = tokenize("abc")
	fmt.Println(tokens)
}

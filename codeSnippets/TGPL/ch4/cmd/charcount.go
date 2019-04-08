// 用于统计输入中每个 Unicode 码点出现的次数。虽然 Unicode 全部码点的数量巨大，但是出现在特定文档
// 中的字符种类并没有多少，使用 map 可以用比较然的方式来跟踪那些出现过字符的次数。
// 编译后可以这样测试：cat charcount.go | ./charcount
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main()  {
	counts := make(map[rune]int)	// counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int	// count of lengths of UTF-8 encodings
	invalid := 0					// count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n , err := in.ReadRune()	// returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// 输入的是无效的 UTF-8 编码字符，unicode.ReplacementChar 表示无效字符，编码长度为 1
		if r == unicode.ReplacementChar && n == 1 {
			invalid ++
			continue
		}

		counts[r] ++
		utflen[n] ++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

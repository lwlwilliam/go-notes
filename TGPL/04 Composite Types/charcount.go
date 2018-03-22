// Charcount computes counts of Unicode characters.
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[rune]int)     // counts of Unicode characters
    var utflen [utf8.UTFMax + 1]int  // count of lengths of UTF-8 encodings
    invalid := 0                     // count of invalid UTF-8 characters

    in := bufio.NewReader(os.Stdin)

    for {
        // ReadRune 方法执行 UTF-8 解码并返回三个值：解码的 rune 字符的值，字符 UTF-8 编码后的长度，和一个错误值。
        r, n, err := in.ReadRune()  // returns rune, nbytes, error

        // 可预期的错误值只有对应的文件结尾的 io.EOF
        if err == io.EOF {
            break
        }

        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }

        // 如果输入的是无效的 UTF-8 编码的字符，返回的 unicode.ReplacementChar 表示无效字符，并且编码长度为 1
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

    fmt.Print("\nlen\tcount\n")

    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }

    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}
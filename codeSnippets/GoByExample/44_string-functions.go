/*
标准库的 strings 包提供了很多有用的字符串相关函数。
 */
package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("1. Contains  :", s.Contains("test", "es"))
	p("2. Count     :", s.Count("test", "t"))
	p("3. HasPrefix :", s.HasPrefix("test", "te"))
	p("4. HasSuffix :", s.HasSuffix("test", "st"))
	p("5. Index     :", s.Index("test", "e"))
	p("6. Join      :", s.Join([]string{"a", "b"}, "-"))
	p("7. Repeat    :", s.Repeat("a", 5))
	p("8. Replace   :", s.Replace("foo", "o", "0", -1))
	p("9. Split     :", s.Split("a-b-c-d-e", "-"))
	p("10. ToLower  :", s.ToLower("TEST"))
	p("11. ToUpper  :", s.ToUpper("test"))
	p()

	// 虽然不是 strings 的一部分，但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制
	p("12. Len      :", len("hello"))
	p("13. Char     :", "hello"[1])
	fmt.Printf("14. %c\n", 101)
}

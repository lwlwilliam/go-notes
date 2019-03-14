/*
Go 提供内置的正则表达式。
 */
package main

//import "bytes"
import "fmt"
import (
	"regexp"
	"bytes"
)

func main() {
	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("1.", match)

	// 上面是直接使用字符串，但是对于一些其他的正则任务，需要 Compile 一个优化的 Regexp 结构体
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("2.", r)

	// 这个结构体有很多方法。这是类似上面的匹配测试
	fmt.Println("3.", r.MatchString("peach"))

	// 查找匹配字符串
	fmt.Println("4.", r.FindString("peach punch"))

	// 查找一个次匹配的字符串的开始和结束索引
	fmt.Println("5.", r.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配的字符串（也就是括号里的内容）
	fmt.Println("6.", r.FindStringSubmatch("peach punch"))

	// 返回完全匹配和局部匹配字符串的开始和结束索引
	fmt.Println("7.", r.FindStringSubmatchIndex("peach punch"))

	// 带 All 的这个函数返回所有的匹配项
	fmt.Println("8.", r.FindAllString("peach punch pinch", -1))

	// 带 All 的这个函数返回所有匹配项的索引
	fmt.Println("9.", r.FindAllStringIndex("peach punch pinch", -1))

	// 返回所有完全匹配和局部匹配项
	fmt.Println("10.", r.FindAllStringSubmatch("peach punch pinch", -1))

	// 返回所有完全匹配和局部匹配项的索引
	fmt.Println("11.", r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 返回前两个完全匹配项
	fmt.Println("12.", r.FindAllString("peach punch pinch", 2))

	// 在上面的例子中，使用了字符串作为参数，并使用了如 MatchString 这样的方法。
	// 也可以提供 []byte 参数并将 String 从函数名中去掉
	fmt.Println("13.", r.Match([]byte("peach")))

	// 创建正则表达式常量时，可以使用 Compile 的变体 MustCompile。因为 Compile 返回两个值，不能用于常量
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("14.", r)

	// 替换匹配部分
	fmt.Println("15.", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	// Func 变量允许传递匹配内容到一个给定的函数中
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("16.", in, string(out))
}

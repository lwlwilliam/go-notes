// 非递归版本的 comma 函数，使用 bytes.Buffer 代替字符串链接操作
// TODO: 虽然实现了，但是感觉实现得不太好，用了两个循环
package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch3/comma"
)

func main()  {
	fmt.Println(comma.Comma2("abcdefghijklmnopqrstuvwxyz"))
}
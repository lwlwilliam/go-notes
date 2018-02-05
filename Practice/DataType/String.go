package main

import (
	"fmt"
)

func main() {
	var a = "hello \n world"
	var b = `hello \n world`

	fmt.Println(a)
	fmt.Println("------")
	fmt.Println(b)
	fmt.Println("\n")
	fmt.Println(len(a))  // 字符串长度
	fmt.Println(a[1])    // 获取单个字节
	fmt.Println(a + b)   // 字符串连接
}

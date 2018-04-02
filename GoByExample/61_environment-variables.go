/*
环境变量是一个在为 Unix 程序传递配置信息的普遍方式。
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// os.Setenv 来设置一个键值对
	os.Setenv("FOO", "1")

	// os.Getenv 获取一个键对应的值
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// os.Environ 来列出所有环境变量键值对，这个函数会返回一个 key=value 形式的字符串切片
	// 可以使用 strings.Split 来得到键和值
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
		//fmt.Println(e)
	}
}

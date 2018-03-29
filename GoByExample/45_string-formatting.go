/*
Go 在传统的 printf 中对字符串格式化提供了优异的支持
 */
package main

import (
	"fmt"
	"os"
)
//import "os"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// 打印 point 的一个实例
	fmt.Printf("%v\n", p)  // {1 2}

	// 如果值是一个结构体，%+v 输出将包括结构体的字段名
	fmt.Printf("%+v\n", p)  // {x:1 y:2}

	// 输入这个值的 Go 语法表示。例如，值的运行源代码片段
	fmt.Printf("%#v\n", p)  // main.point{x:1, y:2}

	// 值的类型
	fmt.Printf("%T\n", p)  // main.point

	// 格式化布尔值
	fmt.Printf("%t\n", true)  // true

	// 格式化整型数，%d 进行标准的十进制格式化
	fmt.Printf("%d\n", 123)  // 123

	// 输出二进制表示
	fmt.Printf("%b\n", 14)  // 1110

	// 输出给定整数的对应字符
	fmt.Printf("%c\n", 33)  // !

	// 提供十六进制编码
	fmt.Printf("%x\n", 456)  // 1c8

	// 最基本的十进制格式化
	fmt.Printf("%f\n", 78.9)  // 78.900000

	// 格式化为科学记数法
	fmt.Printf("%e\n",123400000.0)  // 1.234000e+08
	fmt.Printf("%E\n",123400000.0)  // 1.234000E+08

	// 基本的字符串输出
	fmt.Printf("%s\n", "\"string\"")  // "string"

	// 像 Go 源代码那样带双引号的输出
	fmt.Printf("%q\n", "\"string\"")  // "\"string"\"

	// 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示
	fmt.Printf("%x\n", "hex this")  // 6865782074686973

	// 输入一个指针的值，使用 %p
	fmt.Printf("%p\n", &p)  // 0xc04204a080

	fmt.Printf("|%6d|%6d|\n", 12, 345)			// |    12|   345|

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)		// |  1.20|  3.45|

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)	// |1.20  |3.45  |

	fmt.Printf("|%6s|%6s|\n", "foo", "b")		// |   foo|     b|

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")		// |foo   |b     |

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)  // a string

	fmt.Fprintf(os.Stderr, "an %s\n", "error")  // an error
}

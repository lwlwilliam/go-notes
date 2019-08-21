package main

import (
	"fmt"
)

type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
	var m MyInt = 2
	m = m * m	// 基本类型的运算符仍然适用

	fmt.Println(m.Positive())			// true
	fmt.Println(MyInt(-1).Positive())	// false

/* 取消注释后报错
	var n int
	n = int(m)	// 需要转换
	n = m		// 非法
*/
}

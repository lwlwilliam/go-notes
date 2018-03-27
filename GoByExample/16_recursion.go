/*
Go 支持递归
这是一个经典的阶乘示例
 */
package main

import "fmt"

// fact 函数在到达 face(0) 前一起调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	fmt.Println(fact(7))
}
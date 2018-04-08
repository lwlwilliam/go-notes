// 这个示例程序展示不是总能获取值的地址
package main

import "fmt"

// duration 是一个基于 int 类型的类型
type duration int

// 使用更可读的方式格式化 duration 值
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func (d duration) pretty2() string {
	return fmt.Sprintf("Duration: %d", d)
}

// main 是应用程序的入口
func main() {
//	var d duration = 42
//	fmt.Println(d.pretty())

	// 不能通过指针调用 duration(42) 的方法
	// 不能获取 duration(42) 的地址
	fmt.Println(duration(42).pretty())

	fmt.Println(duration(42).pretty2())
}

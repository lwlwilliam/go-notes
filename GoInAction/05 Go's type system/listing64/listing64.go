// 这个示例程序展示无法从另一个包里
// 访问未公开的标识符
package main

import (
	"fmt"
	"./counters"
)

// main 是应用程序的入口
func main() {
	// 创建一个未公开的类型的变量
	// 并将其初始化为 10
	/*
	不能引用未公开的名字 counters.alertCounter
	未定义: counters.alertCounter

	cannot refer to unexported name counters.alertCounter
	undefined: counters.alertCounter
	 */

	//counter := counters.alertCounter(10)
	//fmt.Printf("Counter: %d\n", counter)

	counter := counters.AlertCounter(10)
	fmt.Printf("Counter: %d\n", counter)

	counter2 := counters.New(10)
	fmt.Printf("Counter2: %d\t%T\n", counter2, counter2)
}

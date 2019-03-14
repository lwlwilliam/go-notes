/*
Go 协程在执行中来说是轻量级的线程

运行这个程序时，首先看到阻塞式调用的输入，然后是两个 Go 协程的交替输出，
这种交替的情况表示 Go 运行时是以异步的方式运行协程的
 */
package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i ++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 使用一般方式调用
	f("direct")

	// 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用
	go f("goroutine")

	// 可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在两个 Go 协程在独立的 Go 协程中异步的运行，所以需要等它们执行结束。这里的 Scanln 代码
	// 需要在程序退出前按下任意键结束
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

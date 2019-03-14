package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i ++ {
		c <- x
		x, y = y, x + y
	}

	//close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)

	// for ... range 语句可以重复地从 channel 接收值直到其关闭为止
	// 如果使用 for ... range 时，channel 没有被关闭将会导致死锁。
	// 并且只应在发送者 goroutine 中关闭 channel，永远不要在发送者中关闭，如果向关闭的 channel 中发送数据会导致 panic
	// 可以注释该语句测试一下
	for i := range c {
		fmt.Println(i)
	}

	// 当以上循环被注释后，ok 为 true。这说明了当 channel 关闭且无值可接收时（充要条件），ok 才会 false，并且 i 为 channel 传输的零值
	i, ok := <- c
	fmt.Println(i, ok)

	i, ok = <- c
	fmt.Println(i, ok)

	i, ok = <- c
	fmt.Println(i, ok)

	i, ok = <- c
	fmt.Println(i, ok)
}

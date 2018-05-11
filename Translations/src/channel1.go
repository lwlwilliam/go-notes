package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
		close(ch)
	}()

	fmt.Println(<- ch)	// 打印"Hello!"
	fmt.Println(<- ch)	// 打印零值""，无阻塞
	fmt.Println(<- ch)	// 再次打印""
	v, ok := <- ch		// v 为""，ok 为 false

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("The channel is empty.")
	}

	// 一直从 ch 中接收值直至其被关闭
	for v := range ch {
		fmt.Println("loop:", v)	// 不会被执行
	}
}

package main

import (
	"fmt"
	"time"
)

// Publish 在给定的时间向 stdout 打印文本。它会立即返回，不会产生阻塞。
func Publish(text string, delay time.Duration) {
	go func() {
//		for {
			time.Sleep(delay)
			fmt.Println("BREAKING NEWS:", text)
//		}
	}()  // 注意"()"，必须调用该匿名函数
}

func main() {
	Publish("A goroutine starts a new thread.", 5 * time.Second)
	fmt.Println("Let's hope the news will published before I leave.")

	// 等待 news 的发布
	time.Sleep(12 * time.Second)

	fmt.Println("Ten seconds later: I'm leaving now.")
}

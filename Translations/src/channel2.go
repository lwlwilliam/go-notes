package main

import (
	"fmt"
	"time"
)

// Publish 在给定时间向 stdout 打印文本，当文本被发布时，它会关闭等待的 channel
func Publish(text string, delay time.Duration) (wait <- chan struct {}) {
	ch := make(chan struct {})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}

func main() {
	wait := Publish("important news", 2 * time.Second)
	// 做一些其他事情

	if res, ok := <- wait; ok {	// 阻塞直至文本被发布
		fmt.Println(res)	
	} else {
		fmt.Println("The channel is empty.")
	}
}

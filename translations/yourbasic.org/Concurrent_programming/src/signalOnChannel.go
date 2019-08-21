package main

import (
	"fmt"
	"time"
)

// 在给定时间打印文本，完成之后，等待中的 channel 会关闭
func Publish(text string, delay time.Duration) (wait <- chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println("BREAKING NEWS:", text)
		close(ch)
	}()
	return ch
}

func main() {
	wait := Publish("Channels let goroutines communicate.", 5 * time.Second)
	fmt.Println("Waiting for news.")
	<- wait
	fmt.Println("Time to leave.")
}

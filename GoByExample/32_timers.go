/*
常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。
Go 内置的定时器和打点器特性让这些很容易实现。

第一个定时器将在程序开始后 2s 失效，但是第二个在它没失效之前就停止了
 */
package main

import (
	"time"
	"fmt"
)

func main() {
	// 定时器表示在未来某一时刻的独立事件。告诉定时器需要等待的时间，
	// 然后它将提供一个用于通知的通道。这里的定时器将等待 2 秒
	timer1 := time.NewTimer(time.Second * 2)

	// 直到这个定时器的通道 C 明确的发送了定时器失效的值之前，将一直阻塞
	<- timer1.C
	fmt.Println("Timer 1 expired")

	// 如果需要的仅仅是单纯的等待，需要使用 time.Sleep。定时器是有用的原因之一就是
	// 可以在定时器失效之前，取消这个定时器
	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println(<- timer2.C)
	}()

	//time.Sleep(time.Second * 2)

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	} else {
		fmt.Println("Timer 2 expired")
	}
}

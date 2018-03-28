/*
定时器是想要在未来某一时刻执行一次使用的
打点器则是想要在固定的时间间隔重复执行准备的。
 */
package main

import (
	"time"
	"fmt"
)

func main() {
	// 打点器和定时器的机制有点相似：一个通道用来发送数据。
	// 在通道上使用内置的 range 来迭代值每隔 500ms 发送一次的值
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 打点器可以和定时器一样被停止。一旦一个打点停止了，将不能再从它的通道中接收到值。
	// 这里在运行后 1600ms 停止这个打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
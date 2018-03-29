/*
速率限制是一个重要的控制服务资源利用和质量的途径。
Go 通过 Go 协程、通道和打点器优美的支持了速率限制
 */
package main

import (
	"time"
	"fmt"
)

func main() {
	// 基本的速率限制。假设想限制我们接收请求的处理，将这些请求发送给一个相同的通道
	// 往缓冲通道 requests 发送 5 个任务
	requests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		requests <- i
	}
	close(requests)

	// 这个 limiter 通道将每 1000ms 接收一个值。这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 1000)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，限制 1000ms 执行一次请求
	// 每 1000ms 接收一个请求，打印接收请求的时间（这里会阻塞）
	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}


	// 有时候想临时进行速率限制，并且不影响整体的速率控制，可以通过通道缓冲来实现。
	// 这个 burstyLimiter 通道用来进行 3 次临时的脉冲型速率限制
	burstyLimiter := make(chan time.Time, 3)

	// 往缓冲通道发送 3 个时间，把通道缓冲占满
	for i := 0; i < 3; i ++ {
		burstyLimiter <- time.Now()
	}

	// 开启 Go 协程，每 1000ms 往 burstyLimiter 发送一个当前时间
	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()

	// 往 burstyRequests 通道发送 5 个任务
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 读取 burstyRequests 和 burstyLimiter 通道任务，打印接收到请求的时间
	// 由于 burstyLimiter 通道 3 个缓冲被占满了，所以从该通道最初获取的 3 个值几乎是同时完成的，
	// 而另一个 Go 协程每隔 1000ms 往 burstyLimiter 发送一个数据，所以之后从 burstyLimiter 读取的时候也是隔 3000ms 一次
	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("burstyRequest", req, time.Now())
	}
}

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
	requests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i ++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

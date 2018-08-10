package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	AvailableMemory = 10 << 20  // 10 MB
	AverageMemoryPerRequest = 10 << 10  // 10 KB
	MAXREQS = AvailableMemory / AverageMemoryPerRequest
)

var sem = make(chan int, MAXREQS)

type Request struct {
	a, b int
	replyc chan int
}

func process(r *Request) {
	// Do something
	// May take a long time and use a lot of memory of CPU
	rand.Seed(time.Now().UnixNano())
	r.replyc <- rand.Intn(1000)
}

func handle(r *Request) {
	process(r)
	// 发送信号，表示已完成，空出一个缓冲位置
	<- sem
}

func Server(queue chan *Request) {
	// 当 channel 缓冲满了时，阻塞直到有可用缓冲位置处理请求

	// 这里是书上的 demo，但是感觉这里没有完善
//	sem <- 1
//	request := <- queue
//	go handle(request)


	// ##########################################################################
	// 这里是自己完善的
	for {
		select {
		case sem <- 1:
			request := <- queue
			go handle(request)
		case <- time.After(5):
			log.Println("等待超时")
		}
	}
}

func main() {
	// 以下是书上的 demo，应该是不完善的
//	queue := make(chan *Request)
//	go Server(queue)


	// ################################################################
	// 自己完善的

	queue := make(chan *Request)
	// 启动服务器监听请求
	go Server(queue)

	const N = 10
	var req [N]*Request
	// 模拟请求
	for i := 0; i < N; i ++ {
		req[i] = &Request{
			a: i,
			b: i + 1,
			replyc: make(chan int) }
		queue <- req[i]
	}

	// 响应请求
	for _, v := range req {
		fmt.Println(<- v.replyc)
	}
}

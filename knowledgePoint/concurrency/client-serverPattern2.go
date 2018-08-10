package main

import (
	"fmt"
	"time"
)

type Request struct {
	a, b int
	replyc chan int  // 请求中的响应 channel
}

type binOp func(a, b int) int

// 请求处理函数
func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		// 如果收到请求就直接处理请求，如果接收到停止信号，终止函数
		select {
		case req := <- service:
			go run(op, req)
		case <- quit:
			return
		}
	}
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	// 请求服务 channel
	service = make(chan *Request)
	// 停止服务信号 channel
	quit = make(chan bool)
	// 开启服务
	go server(op, service, quit)

	// 返回请求服务 channel 和 停止服务信号 channel 以便接收请求进行处理以及及时停止服务
	return service, quit
}

func main() {
	// 指定处理函数，获取请求及停止信号 channel
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 10
	var reqs [N]Request

	// 模拟 N 个请求
	for i := 0; i < N; i ++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)

		// 此处模拟请求处理时间
		time.Sleep(1 * time.Second)
		adder <- req  // 往 adder channel 中发送请求等待处理
	}

	// checks:
	for i := N - 1; i >= 0; i -- {
		if <- reqs[i].replyc != N + 2 * i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request", i, "is ok!")
		}
	}
	quit <- true
	fmt.Println("done")
}

// Client-server 应用是 goroutine 和 channel 应用最闪光的地方。
// Go 服务器通过 goroutine 响应客户端，因此每一个客户端请求都会启动一个 goroutine。
// 一个常用的技巧是客户端-服务器模型包含一个 channel，服务器就通过这个 channel 往
// 客户端发送响应。
package main

import "fmt"

// 请求结构体
type Request struct {
	a, b int  // 相当于请求参数吧
	replyc chan int  // Request 中的响应 channel
}

// 对请求参数的操作函数
type binOp func(a, b int) int

// 对请求进行处理，并发送到请求中的 replyc channel
func run (op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

// 服务处理程序，使用 binOp 类型的函数对请求进行处理
func server (op binOp, service chan *Request) {
	for {
		// 接收请求
		req := <- service
		// 为请求开启 goroutine
		go run(op, req)
	}
}

// 启动服务器，返回 Request channel 的地址
func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	// 启动服务程序
	go server(op, reqChan)
	return reqChan
}

func main() {
	// adder 是 chan *Request，请求通道啊
	// 开启服务器，并创建了请求 channel *Request
	adder := startServer(func(a, b int) int {
		return a + b
	})

	const N = 100
	// 创建数组保存请求数据
	var reqs [N]Request
	// 模拟 100 个请求
	for i := 0; i < N; i ++ {
		// 把请求内容保存到数组对应的 Request 中
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req  // 完成一次请求
	}

	for i := N - 1; i >= 0; i -- {
		if <- reqs[i].replyc != N + 2 * i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request", i, "is ok!")
		}
	}
	fmt.Println("done")
}

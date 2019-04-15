// Clock2 is a TCP server that periodically writes the time.
// 使用 netcat 工具（nc 命令）作为客户端请求 clock2
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main()  {
	// net.Listener 对象，监听一个网络端口上到来的连接，该例中用的是 TCP 的 localhost:8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 直接阻塞，直到一个新的连接被创建，然后返回一个 net.Conn 对象来表示这个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// 处理一个完整的客户端连接
		go handleConn(conn) // handle connection concurrently
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for {
		// 获取当前时刻，写到客户端
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

// echo 服务器
// 注册并监听一个端口，阻塞在 accept 操作，并等待客户端连接，accept 调用返回一个连接对象
// echo 服务非常简单，把客户端的请求数据写回到客户端，就像回声一样，直到某一方关闭连接
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on" + *host + ":" + *port)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			os.Exit(1)
		}

		fmt.Printf("Received message %s -> %s\n", conn.RemoteAddr(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		io.Copy(conn, conn)
	}
}

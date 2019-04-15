// listen、accept 从客户端过来的连接
// 对每一个连接，程序都会建立一个新的 handleConn 的 goroutine
package main

import (
	"net"
	"log"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch8/chat/core"
)

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go core.Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go core.HandleConn(conn)
	}
}

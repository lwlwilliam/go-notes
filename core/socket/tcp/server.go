// 与 client2.go 配合使用
// 服务器很忙，瞬间有大量 client 端连接尝试向 server 建立，server 端的 listen backlog 队列满，server accept 不及时
// ，这将导致 client 端 Dial 阻塞
package main

import (
	"net"
	"log"
	"time"
)

func main()  {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Println("error listen:", err)
		return
	}
	defer l.Close()
	log.Println("listen ok")

	var i int
	for {
		// 每 10 秒接收一个连接
		time.Sleep(time.Second * 10)
		conn, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		i ++
		log.Printf("%d: accept a new connection from %v\n", i, conn.RemoteAddr())
	}
}

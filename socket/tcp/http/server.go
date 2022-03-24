// socket 实现的简单 http 服务器
// 所有 TODO 已解决
package main

import (
	"log"
	"net"
)

// HTTP 响应报文，行分隔符为 \r\n
// TODO: 确定一下是否 200 报文要 Content-Length 头(DONE: 由于没有长度，客户端不知道什么时候应该结束，所以会一直读取数据，直到服务端关闭连接导致客户端出错，虽然客户端可以接收服务端发送的数据，但由于异常的关闭连接，所以一般会提示连接已被重置)
// var content = []byte("HTTP/1.1 200 OK\r\nContent-type:text/plain\r\nContent-Length: 12\r\n\r\nHello world!")
var content = []byte("HTTP/1.1 200 OK\r\nTransfer-Encoding: chunked\r\n\r\n0\r\n\r\n")
//var content = []byte("HTTP/1.1 302 Moved Temporarily\r\nLocation: http://example.com")

func handleConn(conn *net.Conn) {
	// TODO: 这里发现一定要把连接里的数据都读出来才能正常写入(DONE: 不影响)
	//var buf = make([]byte, 1) // buf 要确保所有数据都读出来了，最好循环读到 EOF
	//_, err := (*conn).Read(buf)
	//checkErr(err)

	n, err := (*conn).Write(content)
	defer (*conn).Close()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("write %d bytes to %s\n", n, (*conn).RemoteAddr())
}

func main() {
	addr := "localhost:10000"
	listener, err := net.Listen("tcp", addr)
	log.Printf("listening %s...", addr)
	checkErr(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		log.Printf("accept from %s\n", conn.RemoteAddr())

		if err != nil {
			continue
		}
		go handleConn(&conn)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

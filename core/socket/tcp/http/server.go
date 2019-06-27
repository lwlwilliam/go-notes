// socket 实现的简单 http 服务器
package main

import (
	"log"
	"net"
)

// HTTP 响应报文，行分隔符为 \r\n
var content = []byte("HTTP/1.1 200 OK\r\nContent-type:text/plain\r\n\r\nHello world!")

func handleConn(conn *net.Conn) {
	// TODO:这里发现一定要把连接里的数据都读出来才能正常写入
	var buf = make([]byte, 1) // buf 要确保所有数据都读出来了，最好循环读到 EOF
	_, err := (*conn).Read(buf)
	checkErr(err)

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
	log.Println("listening...")
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

// socket 实现的简单 http 服务器
package HTTP

import (
	"log"
	"net"
	"os"
)

// TODO:暂时不清楚这响应报文的准确格式，是 \r\n 呢？还是 \n 呢？又好像都不对
//var content = `HTTP/1.1 200 OK\r\nContent-type: text/html\r\n\r\nHello world!`

// 只能保证这个是可以的
var content = `HTTP/1.1 200 OK
Content-type: text/plain

Hello world!`

func handleConn(conn net.Conn) {
	// 这里发现一定要把连接里的数据都读出来才能正常写入
	var buf = make([]byte, 1024)  // buf 要确保所有数据都读出来了，最好循环读到 EOF
	_, err := conn.Read(buf)
	checkErr(err)

	conn.Write([]byte(content))
	conn.Close()
}

func main()  {
	addr := "localhost:10000"
	listener, err := net.Listen("tcp", addr)
	checkErr(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// 只关闭网络连接中写的部分
package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct {}{}
	}()
	mustCopy(conn, os.Stdin)

	if tcpconn, ok := conn.(*net.TCPConn); ok {
		log.Println("tcp connection close")
		tcpconn.CloseWrite()
	} else {
		conn.Close()
	}

	<- done
}

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

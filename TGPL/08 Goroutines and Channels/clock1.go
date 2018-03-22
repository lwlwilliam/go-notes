// Clock1 is a TCP server that periodically writes the time.
// listener 对象的 Appept 方法会直接阻塞，导致每次只能处理一个连接
package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)  // e.g., connection aborted
			continue
		}
		handleConn(conn)  // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

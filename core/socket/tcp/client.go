package main

import (
	"net"
	"time"
)

func main()  {
	// 阻塞 Dial
	//conn, err := net.Dial("tcp", "localhost:8888")

	// 带上超时机制的 Dial
	conn, err := net.DialTimeout("tcp", "localhost:8888", 30 * time.Second)

	if err != nil {
		// handle error
	}


	// read or write on conn
	conn.Write([]byte("Hello world!"))
}

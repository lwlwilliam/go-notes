package main

import (
	"log"
	"net"
	"time"
)

func main()  {
	addr := ":10000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		time.Sleep(5 * time.Second)
		conn.Close()
	}
}

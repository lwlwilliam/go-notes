package main

import (
	"log"
	"net"
	"os"
)

func handleConn(conn net.Conn)  {
	_, err := conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	checkErr(err)
	_, err = conn.Write([]byte("Content-Type: text/html\r\n"))
	checkErr(err)
	_, err = conn.Write([]byte("Content-Length: 4\r\n\r\n"))
	checkErr(err)
	_, err = conn.Write([]byte("Body"))
	checkErr(err)
	conn.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main()  {
	addr := ":10000"
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
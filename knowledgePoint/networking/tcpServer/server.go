package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server...")
	
	// listener:
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}

	// accept connnection from client:
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error acception", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		// read from client
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		// conn.Write([]byte("Received data: " + string(buf)))
		fmt.Printf("Received data: %v\n", string(buf))
	}
}
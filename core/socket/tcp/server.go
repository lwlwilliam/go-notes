// go tcpsocket server
package main

import (
	"net"
	"fmt"
)

func handleConn(c net.Conn)  {
	c.Write([]byte("Received"))
	defer c.Close()

	//for {
		// read from the connection
		// ... ...
		// write to the connection
		// ... ...
	//}
}

func main()  {
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

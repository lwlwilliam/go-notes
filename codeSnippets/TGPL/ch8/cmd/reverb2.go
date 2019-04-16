// 配合 netcat2 使用
package main

import (
	"net"
	"time"
	"fmt"
	"strings"
	"bufio"
	"log"
)

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn)  {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo (c, input.Text(), 1 * time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	log.Println("connection close")
	c.Close()
}

func main()  {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}


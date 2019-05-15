package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:6001"
	conn, err := net.Dial("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 10240)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))

	reader := bufio.NewReader(os.Stdin)
	var cmd, line string
	for {
		fmt.Print(">>> ")
		line, err = reader.ReadString('\n')
		fmt.Sscan(line, &cmd)
		if len(line) == 1 {
			continue
		}

		go sender(conn, line)
	}
}

func sender(conn net.Conn, line string) {
	n, err := conn.Write([]byte(line))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		n, err = conn.Read(buf)
		if err == io.EOF {
			conn.Close()
			break
		}
		fmt.Print(string(buf[:n]))
	}
	return
}

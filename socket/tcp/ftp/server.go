package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	service := "127.0.0.1:6001"
	listener, err := net.Listen("tcp", service)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	n, err := conn.Write([]byte("ftp server running..."))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("has written %d bytes to the client\n", n)

	reader := bufio.NewReader(conn)
	var cmd, file string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(line)
		fmt.Println(len(strings.Fields(line)))
		if len(line) == 0 {	// keep alive
			continue
		}
		cmd = strings.ToLower(strings.Fields(line)[0])
		if len(strings.Fields(line)) > 1 {
			file = strings.Fields(line)[1] // get the file
		}

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal("path error")
		}

		file = filepath.Join(pwd, file)
		fmt.Println(file)

		switch cmd {
		case "get":
			fd, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			content, err := ioutil.ReadAll(fd)
			if err != nil {
				log.Fatal(err)
			}
			fd.Close()

			n, err := conn.Write(content)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("has written %d bytes of the file to client", n)
		case "push":
			fmt.Println("upload...")
			n, err := conn.Write([]byte("upload..."))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("has written %d bytes to the client", n)
		case "exit":
			return
		default:
			fmt.Println("invalid command")
			n, err := conn.Write([]byte("invalid command"))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("has written %d bytes to client", n)
		}
	}
}

package main

import (
	"net"
	"log"
	"fmt"
	"bytes"
	"net/url"
	"strings"
	"io"
)

func main() {
	// 监听 tcp 8080 端口
	socket, err := net.Listen("tcp", ":8888")
	fatalError(err)

	// 接收客户端请求
	for {
		client, err := socket.Accept()
		fatalError(err)

		// 处理请求
		go requestHandler(client)
	}
}

func requestHandler(client net.Conn) {
	if client == nil {
		return
	}

	defer client.Close()

	var buf [1024]byte
	n, err := client.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}

	var method, host, address string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	} else {
		// host 不带端口，默认 80
		if strings.Index(hostPortURL.Host, ":") == -1 {
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	if method == "CONNECT" {
		fmt.Fprintf(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}

	// 进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}

func fatalError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
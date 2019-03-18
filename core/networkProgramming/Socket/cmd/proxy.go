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
	// 监听 tcp 60000 端口
	socket, err := net.Listen("tcp", ":60000")
	fatalError(err)

	// 接收客户端请求
	for {
		request, err := socket.Accept()
		fatalError(err)

		// 处理请求
		go requestHandler(request)
	}
}

func requestHandler(request net.Conn) {
	if request == nil {
		return
	}

	defer request.Close()

	// 获取请求内容
	var buf [1024]byte
	n, err := request.Read(buf[:])
	if err != nil {
		log.Println("Read request:", err)
		return
	}

	var method, host, address string

	// 获取请求方法和主机
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println("Parse host:", err)
		return
	}

	// 确定端口
	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	} else {
		if strings.Index(hostPortURL.Host, ":") == -1 {
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	// 到目标服务器的 socket
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("Create socket to target server:", err)
		return
	}

	// 记录请求 socket
	log.Println("Attempt to request:", address)

	// 创建隧道
	if method == "CONNECT" {
		fmt.Fprintf(request, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}

	// 进行转发
	go io.Copy(server, request)
	io.Copy(request, server)
}

func fatalError(err error) {
	if err != nil {
		log.Panic("Panic", err)
	}
}

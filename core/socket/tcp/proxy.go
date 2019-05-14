// 正向代理基本原理：
// 1. 代理程序以守护进程方式运行，接受客户端的连接，实际就是 tcp socket 编程；
// 2. 与客户端建立连接后，以多线程，多进程的方式（在 Go 中就是 goroutine）处理连接；
// 3. 每个连接都是一个独立的 goroutine。首先就是要读取客户端的请求内容，对请求内容进行解析，
//    在这里以每次 1024 字节对内容进行处理，默认客户端以 HTTP/HTTPS 协议与代理通信，所以按
//    协议规范进行解析即可；
// 4. 解析客户端内容后可以获取请求主机和端口号，代理通过主机和端口号可以建立与目标服务器的连接，
//    然后把客户端的请求内容写到连接，再把服务器返回的内容写回到与客户端的连接即可；
// 5. 以上就是代理的简单原理；
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
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
	var buf = make([]byte, 1024)
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

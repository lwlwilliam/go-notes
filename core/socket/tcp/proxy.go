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
	"strconv"
)

func main() {
	// 监听 tcp 60000 端口
	socket, err := net.Listen("tcp", "127.0.0.1:60000")
	fatalError(err)

	log.Println("Running...")

	// 接收客户端请求
	for {
		request, err := socket.Accept()
		fatalError(err)

		log.Println("Local address:", request.LocalAddr())
		log.Println("Client address:", request.RemoteAddr())

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
	buf := make([]byte, 1024)
	buff := bytes.NewBuffer(nil)
	requestLen := 0
	contentLen := 0

	for {
		n, err := request.Read(buf)
		if err != nil {
			if err == io.EOF {
				requestLen += n
				break
			}

			log.Println("Read request err:", err)
			continue
		}

		buff.Write(buf[:n])
		requestLen += n

		// 获取实体长度
		headers := bytes.Split(buff.Bytes(), []byte("\r\n"))
		for _, v := range headers {
			if bytes.Index(v, []byte("Content-Length:")) == 0 {
				contentLen, err = strconv.Atoi(fmt.Sprintf("%d", bytes.TrimSpace(bytes.Split(v, []byte(":"))[1])))
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		}

		// 确定实体长度
		if headerEnd := bytes.Index(buff.Bytes(), []byte("\r\n\r\n")); headerEnd > -1 {
			// 实体数据接收完成
			if contentLen == buff.Len() - headerEnd - 4 {
				break
			} else if contentLen < buff.Len() - headerEnd - 4 {
				log.Println("Bad request: Content-Length, ", contentLen, "; Transfer Length, ", buff.Len() - headerEnd)
				continue
			}
		}
	}

	log.Println(buff.Len(), requestLen)

	// 获取请求方法和主机
	var method, host, address string
	end := bytes.IndexByte(buff.Bytes()[:requestLen], '\n')
	if end < 0 {
		log.Println("Bad request:")
		log.Println(string(buff.Bytes()))
		return
	}
	fmt.Sscanf(string(buff.Bytes()[:end]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println("Parse url err:", err)
		return
	}

	// address = host + port
	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	} else {
		if strings.Index(hostPortURL.Host, ":") == -1 {
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	// 记录请求 socket
	log.Println("Address the client requests:", address)

	// 连接目标服务器
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("Create socket to target server err:", err)
		return
	}

	// 打印请求报文
	fmt.Printf("%s", buff.Bytes()[:requestLen])

	// 创建隧道
	if method == "CONNECT" {
		fmt.Fprintf(request, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buff.Bytes()[:requestLen])
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
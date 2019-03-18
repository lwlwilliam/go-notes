// socket 实现简单的 http 客户端
package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	socket := os.Args[1]

	// 相当于 socket() + bind() + listen()
	conn, err := net.Dial("tcp", socket)
	checkError(err)

	// 发送 HTTP 请求
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: " + socket + "\r\n\r\n"))
	checkError(err)

	// 获取 HTTP 响应
	result, err := resp(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

// 读取响应报文
func resp(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	//result := bytes.NewBuffer([]byte("Just a test\n"))
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		// 从连接中读取字符到 buf 中，再把 buf 写入缓冲 result 中
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])

		// 读取出错直接中断操作
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}

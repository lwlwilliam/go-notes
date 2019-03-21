// 模拟 http get 请求
package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	socket := os.Args[1]

	// 获取 tcp 地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", socket)
	checkError(err)

	// 建立 tcp 连接
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// 写入 http 请求报文
	_, err = conn.Write([]byte(
		"GET / HTTP/1.1\r\n" +
		"Host: " + socket + "\r\n\r\n"))
	checkError(err, "write to conn")

	// 读取 http 响应
	result, err := ioutil.ReadAll(conn)
	checkError(err, "read from conn")

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error, step string)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error(), "; step: ", step)
		os.Exit(1)
	}
}

// 模拟时间服务器
// 单任务
// 可以用 tcpClient.go 或 tcpClient2.go 来测试
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"log"
)

func main() {
	// 创建 tcp 地址
	service := ":10000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 监听 tcp 地址
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		daytime := time.Now().String()
		daytime = "HTTP/1.1 302 Moved Temporarily\r\nLocation: https://github.com\r\n\r\n"
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

// 模拟时间服务器
// 支持多并发
// 可以处理客户端实际请求内容
package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"strings"
)

func main() {
	service := ":8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))  // 2 分钟超时
	request := make([]byte, 128)  // 设置请求最大长度为 128 字节防止洪水攻击
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			fmt.Println("read_len = 0")
			break  // 连接已被客户端关闭
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			fmt.Println("timestamp")
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			fmt.Println("now:", daytime)
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128)  // 清除上次读取的内容
	}
}

func checkError(err error)  {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

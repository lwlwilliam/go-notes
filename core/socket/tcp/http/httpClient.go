// socket 实现 http 客户端
// 串行指定读取客户端返回内容大小，不推荐使用
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

func main() {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("访问的公网 IP 地址是：", conn.RemoteAddr().String(),
		"; IP 地址数据类型：", reflect.TypeOf(conn.RemoteAddr().String()))
	fmt.Println("客户端连接的地址及端口是：", conn.LocalAddr(),
		"; IP 地址数据类型：", reflect.TypeOf(conn.LocalAddr()))

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("发送数据长度：", n)

	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(string(buf[:n]))
	conn.Close()
}
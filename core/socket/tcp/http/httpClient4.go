// socket 实现 http 客户端
// io 读取 http 响应
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"reflect"
)

func main()  {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conn.RemoteAddr().String(),
		reflect.TypeOf(conn.RemoteAddr().String()))
	fmt.Println(conn.LocalAddr(),
		reflect.TypeOf(conn.LocalAddr()))

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("size:", n)

	io.Copy(os.Stdout, conn)
	conn.Close()
}

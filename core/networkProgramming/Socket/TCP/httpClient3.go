// socket 实现 http 客户端
// 按行读取
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			conn.Close()
		}
		fmt.Print(line)
	}
}

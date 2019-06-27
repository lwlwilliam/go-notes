// socket 实现 http 客户端
// 按照指定大小循环读取
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"bufio"
	"os"
	"bytes"
	"io/ioutil"
)

func main() {
	request(1)
}

func request(t uint) {
	addr := "www.baidu.com:80"
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n%s\n", conn.RemoteAddr().String(), reflect.TypeOf(conn.RemoteAddr().String()))
	fmt.Printf("%s\n%s\n", conn.LocalAddr(), reflect.TypeOf(conn.LocalAddr()))
	fmt.Println()

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	switch t {
	case 1:
		readChunk(&conn)
	case 2:
		readLine(&conn)
	case 3:
		ioCopy(&conn)
	case 4:
		buffer(&conn)
	case 5:
		readAll(&conn)
	default:
		readChunk(&conn)
	}
}

// 按块读取
func readChunk(conn *net.Conn) {
	buf := make([]byte, 100)
	for {
		n, err := (*conn).Read(buf)
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				fmt.Print(string(buf[:n]))
			}

			break
		}
		fmt.Print(string(buf[:n]))
	}
}

// 按行读取
func readLine(conn *net.Conn) {
	r := bufio.NewReader(*conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				fmt.Print(line)
			}
			break
		}
		fmt.Print(line)
	}
}

// io.copy
func ioCopy(conn *net.Conn) {
	io.Copy(os.Stdout, *conn)
}

// 写入 buffer
func buffer(conn *net.Conn) {
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		// 从连接中读取字符到 buf 中，再把 buf 写入缓冲 result 中
		n, err := (*conn).Read(buf[0:])
		result.Write(buf[0:n])

		// 读取出错直接中断操作
		if err != nil {
			log.Println(err)
			if err == io.EOF {
				break
			}
		}
	}

	result.WriteTo(os.Stdout)
}

// 读取所有
func readAll(conn *net.Conn) {
	result, err := ioutil.ReadAll(*conn)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("%s", result)
}

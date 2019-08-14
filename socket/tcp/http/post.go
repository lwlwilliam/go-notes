// post 模拟了浏览器的表单请求，分别是普通的 form 表单提交以及包含文件上传的表单提交
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
)

var (
	hostname *string
	port     *string
)

func main() {
	conn := dial()

	// request
	t := 2
	switch t {
	case 1:
		urlencoded(conn)
	case 2:
		multipart(conn)
	default:
		urlencoded(conn)
	}

	// response
	resp(conn)
}

func dial() *net.Conn {
	hostname = flag.String("h", "127.0.0.1", "hostname")
	port = flag.String("p", "80", "port")
	flag.Parse()

	socket := *hostname + ":" + *port
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}

	return &conn
}

func resp(conn *net.Conn) {
	buf := make([]byte, 100)
	for {
		n, err := (*conn).Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s", buf[:n])
			}

			log.Println(err)
			break
		}
		fmt.Printf("%s", buf[:n])
	}
}

// x-www-form-urlencoded 方式 post 提交
func urlencoded(conn *net.Conn)  {
	//x-www-form-urlencoded
	body := "name=William A&hobby=running"
	requestMsg := bytes.NewBuffer(nil)
	requestMsg.WriteString("POST /test/web/php/test.php HTTP/1.1\r\n")
	requestMsg.WriteString("Host: 127.0.0.1\r\n")
	requestMsg.WriteString("Content-Type: application/x-www-form-urlencoded\r\n")
	requestMsg.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	requestMsg.WriteString("\r\n")
	requestMsg.WriteString(body)

	_, err := (*conn).Write(requestMsg.Bytes())
	if err != nil {
		log.Fatalf("write: %v", err)
	}
}

// multipart/form-data 模拟浏览器以 <form action="" enctype="multipart/form-data"></form> 形式提交表单（上传文件）
func multipart(conn *net.Conn)  {
	// body 边界
	customizedBoundary := "customizedBoundary"

	// 准备要上传的文件
	fd, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	// 构造 body
	body := "--" + customizedBoundary + "\r\n" +
		"Content-Disposition: form-data; name=name\r\n\r\n" +
		"William A\r\n" +
		"--" + customizedBoundary + "\r\n" +
		"Content-Disposition: form-data; name=hobby\r\n\r\n" +
		"programming\r\n" +
		"--" + customizedBoundary + "\r\n" +
		"Content-Disposition: form-data; name=file; filename=\"test.png\"\r\n" +
		"Content-Type: image/png\r\n\r\n" +
		string(content) + "\r\n" +
		"--" + customizedBoundary

	requestMsg := bytes.NewBuffer(nil)
	requestMsg.WriteString("POST /test/web/php/test.php HTTP/1.1\r\n")
	requestMsg.WriteString("Host: 127.0.0.1\r\n")
	requestMsg.WriteString("Content-Type: multipart/form-data;boundary=" + customizedBoundary + "\r\n")
	requestMsg.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n")
	requestMsg.WriteString("\r\n")
	requestMsg.WriteString(body)

	fmt.Println("##################################")
	fmt.Println("request")
	fmt.Println("##################################")
	fmt.Printf("\n\n")
	fmt.Print(requestMsg.String())

	fmt.Printf("\n\n\n")
	fmt.Println("##################################")
	fmt.Println("response")
	fmt.Println("##################################")
	fmt.Printf("\n\n")

	_, err = (*conn).Write(requestMsg.Bytes())
	if err != nil {
		log.Fatalf("write: %v", err)
	}
}

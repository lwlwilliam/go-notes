package main

import (
	"flag"
	"net"
	"log"
	"bytes"
	"fmt"
	"strconv"
	"io"
	"os"
	"io/ioutil"
)

var (
	hostname *string
	port *string
)

func main() {
	hostname = flag.String("h", "127.0.0.1", "hostname")
	port = flag.String("p", "80", "port")
	flag.Parse()

	socket := *hostname + ":" + *port
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		log.Fatalf("dial: %v", err)
	}
	defer conn.Close()


	// x-www-form-urlencoded
	//body := "name=William A&hobby=running"
	//requestMsg := bytes.NewBuffer(nil)
	//requestMsg.WriteString("POST /test/web/php/test.php HTTP/1.1\r\n")
	//requestMsg.WriteString("Host: 127.0.0.1\r\n")
	//requestMsg.WriteString("Content-Type: application/x-www-form-urlencoded\r\n")
	//requestMsg.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n")
	//requestMsg.WriteString(body)


	// multipart/form-data 模拟浏览器以 <form action="" enctype="multipart/form-data"></form> 形式提交表单（上传文件）
	customizedBoundary := "customizedBoundary"
	fd, err := os.Open("test.png")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}
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
	requestMsg.WriteString("Content-Length: " + strconv.Itoa(len(body)) + "\r\n\r\n")
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

	_, err = conn.Write(requestMsg.Bytes())
	if err != nil {
		log.Fatalf("write: %v", err)
	}

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("%s", buf[:n])
			}

			log.Println(err, ":", n)
			break
		}
		fmt.Printf("%s", buf[:n])
	}
}

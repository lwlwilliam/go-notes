package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	service := "/tmp/echo.sock"
	conn, err := net.Dial("unix", service)
	checkErr(err)
	defer conn.Close()

	n, err := conn.Write([]byte("Hello world!"))
	checkErr(err)

	buf := make([]byte, 512)
	n, err = conn.Read(buf[0:n])
	checkErr(err)

	fmt.Println(string(buf))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
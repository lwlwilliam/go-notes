package main

import (
	"log"
	"net"
	"fmt"
)

func main()  {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", "localhost:8888")
	// 网络不可达或对方服务未启动
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println(string(buf))
	defer conn.Close()
	log.Println("dial ok")
}

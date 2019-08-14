// type TCPAddr struct {
//     IP	IP
//     Port	int
// }
//
// 创建 TCPAddr，会自动调用 DNS 解析程序
// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s addr:port\n", os.Args[0])
		os.Exit(1)
	}
	addr := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	fmt.Println(tcpAddr.IP)
	fmt.Println(tcpAddr.Port)
	os.Exit(0)
}

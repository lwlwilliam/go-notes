// 多播客户端
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("239.0.2.250")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port:0}
	dstAddr := &net.UDPAddr{IP:ip, Port:9999}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	conn.Write([]byte("Hello world!"))

	fmt.Printf("<%s>\n", conn.RemoteAddr())
}
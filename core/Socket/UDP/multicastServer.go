// 多播服务端
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "239.0.2.250:9999")
	if err != nil {
		fmt.Println(err)
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local: <%s>\n", listener.LocalAddr().String())

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s\n", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
	}
}

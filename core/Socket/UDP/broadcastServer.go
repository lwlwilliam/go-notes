package main

import (
	"fmt"
	"net"
)

func main() {
	//listener, err := net.ListenUDP("udp", &net.UDPAddr{IP:net.IPv4zero, Port:9999})

	// 这块代码跟以上注释代码等价
	udpAddr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	listener, err := net.ListenUDP("udp", udpAddr)


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

		_, err = listener.WriteToUDP([]byte("Hello client!"), remoteAddr)

		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}
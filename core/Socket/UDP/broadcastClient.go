package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ip := net.ParseIP("192.168.30.255")

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port:0}
	dstAddr := &net.UDPAddr{IP: ip, Port:9999}

	conn, err := net.ListenUDP("udp", srcAddr)
	if err != nil {
		fmt.Println(err)
	}

	n, err := conn.WriteToUDP([]byte("Hello client!"), dstAddr)
	if err != nil {
		fmt.Println(err)
	}

	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read %s from <%s>\n", data[:n], remoteAddr.String())

	b := make([]byte, 1)
	os.Stdin.Read(b)
}

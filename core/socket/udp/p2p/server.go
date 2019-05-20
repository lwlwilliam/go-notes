package main

import (
	"net"
	"log"
	"time"
	"fmt"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 9981})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("local: <%s> \n", listener.LocalAddr().String())

	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])

		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("punch hole, %s <--> %s\n", peers[0].String(), peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)

			log.Println("the exit of the transit server will not affect the communication between peers")
			return
		}
	}
}

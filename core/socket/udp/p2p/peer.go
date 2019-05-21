// p2p peers
package main

import (
	"net"
	"log"
	"fmt"
	"strings"
	"strconv"
	"time"
	"flag"
)

const HandShakeMsg = "udp hole message"

var (
	tag, ip string
)

func main() {
	tag = *flag.String("t", "mac", "the process tag")
	ip = *flag.String("i", "", "the process tag")
	flag.Parse()

	if ip == "" {
		log.Fatal("IP address can not be empty.")
	}

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9982}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(ip), Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = conn.Write([]byte("Hello, I'm new peer:" + tag)); err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()

	anotherPeer := parseAddr(string(data[:n]))
	fmt.Printf("local: %s server: %s another: %s\n", srcAddr, remoteAddr, anotherPeer.String())

	// start to punch hole
	bidirectionHole(srcAddr, &anotherPeer)
}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP: net.ParseIP(t[0]),
		Port:port,
	}
}

func bidirectionHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// send message to another peer (the nat device of another peer will discard the message, because of the invalid origin of the it)，
	// to punch a hole between the peers
	if _, err = conn.Write([]byte(HandShakeMsg)); err != nil {
		log.Println("send handshake:", err)
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			if _, err = conn.Write([]byte("from [" + tag + "]")); err != nil {
				log.Println("send msg fail", err)
			}
		}
	}()

	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("received: %s\n", data[:n])
		}
	}
}

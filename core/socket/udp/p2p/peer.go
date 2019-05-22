// p2p peers
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

const HandShakeMsg = "udp hole message"

var (
	tag, ip *string
)

func main() {
	tag = flag.String("t", "mac", "the process tag")
	ip = flag.String("i", "", "the server ip")
	flag.Parse()

	if *ip == "" {
		log.Fatal("IP address can not be empty.")
	}

	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9983}
	dstAddr := &net.UDPAddr{IP: net.ParseIP(*ip), Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = conn.Write([]byte("Hello, I'm new peer:" + *tag)); err != nil {
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
		IP:   net.ParseIP(t[0]),
		Port: port,
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

	log.Println("dial successfully:", conn.RemoteAddr())

	go func() {
		for {
			time.Sleep(10 * time.Second)
			log.Println("write at", conn.LocalAddr())
			if _, err = conn.Write([]byte("from [" + *tag + "]")); err != nil {
				log.Println("send msg fail", err)
			} else {
				log.Println("send msg successfully to:", conn.RemoteAddr())
			}
		}
	}()

	for {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		log.Println("read at", conn.LocalAddr())
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("received from %s: %s\n", conn.RemoteAddr(), data[:n])
		}
	}
}

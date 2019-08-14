// 实现 ping 功能
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
	"log"
)

// icmp 报文格式
type ICMP struct {
	Type        uint8
	Code        uint8
	CheckSum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func usage() {
	msg := `
Need to run as root!
Usage:
	goping host
	Example: ./goping www.baidu.com`

	fmt.Println(msg)
	os.Exit(0)
}

// seq 是序号，Sequence number
func getICMP(seq uint16) ICMP {
	// 这个就是 icmp 的报文内容了
	icmp := ICMP{
		Type:        8,
		Code:        0,
		CheckSum:    0,
		Identifier:  0,
		SequenceNum: seq,
	}

	log.Println("icmp before:", icmp)

	var buffer bytes.Buffer
	// 注意大小端问题
	binary.Write(&buffer, binary.BigEndian, icmp)
	icmp.CheckSum = CheckSum(buffer.Bytes())
	buffer.Reset()

	log.Println("icmp after:", icmp)

	return icmp
}

func sendICMPRequest(icmp ICMP, destAddr *net.IPAddr) error {
	conn, err := net.DialIP("ip4:icmp", nil, destAddr)
	if err != nil {
		log.Fatalf("Fail to connect to remote host: %s\n", err)
	}
	defer conn.Close()

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, icmp)

	if _, err := conn.Write(buffer.Bytes()); err != nil {
		return err
	}

	tStart := time.Now()

	conn.SetReadDeadline((time.Now().Add(time.Second * 2)))

	recv := make([]byte, 1024)
	receiveCnt, err := conn.Read(recv)

	if err != nil {
		return err
	}

	tEnd := time.Now()
	duration := tEnd.Sub(tStart).Nanoseconds() / 1e6

	fmt.Printf("%d bytes from %s: seq=%d time=%dms\n", receiveCnt, destAddr.String(), icmp.SequenceNum, duration)

	return err
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	host := os.Args[1]
	raddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		log.Fatalf("Fail to resolve %s, %s\n", host, err)
	}

	fmt.Printf("Ping %s (%s):\n\n", raddr.String(), host)

	for i := 1; i < 6; i++ {
		if err = sendICMPRequest(getICMP(uint16(i)), raddr); err != nil {
			log.Fatalf("Error: %s\n", err)
		}
		time.Sleep(2 * time.Second)
	}
}

// 检验和
func CheckSum(data []byte) uint16 {
	var (
		sum    uint32
		length int = len(data)
		index  int
	)
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 {
		sum += uint32(data[index])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
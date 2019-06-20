package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
	"time"
)

type DNSHeader struct {
	Id                                              uint16
	Flags                                           uint16
	Questions, AnswerRRs, AuthorityRRS, AdditionRRS uint16
}

func (header *DNSHeader) SetFlag(QR uint16, OperationCode uint16, AuthoritativeAnswer uint16, Truncation uint16, RecursionDesired uint16, RecursionAvailable uint16, ResponseCode uint16) {
	header.Flags = QR<<15 + OperationCode<<11 + AuthoritativeAnswer<<10 + Truncation<<9 + RecursionDesired<<8 + RecursionAvailable<<7 + ResponseCode
}

type DNSQueries struct {
	Type  uint16
	Class uint16
}

func ParseDomainName(domain string) []byte {
	var (
		buffer   bytes.Buffer
		segments = strings.Split(domain, ".")
	)

	// 把域名分成多段，先写入段长度，再写入段内容。如果循环
	// 例如：www.baidu.com 则写入 3 www 5 baidu 3 com。则 1 + 3 + 1 + 5 + 1 + 3 + 1 = 15 个字节
	for _, seg := range segments {
		binary.Write(&buffer, binary.BigEndian, byte(len(seg)))
		binary.Write(&buffer, binary.BigEndian, []byte(seg))
	}
	binary.Write(&buffer, binary.BigEndian, byte(0x00)) // 最后写入 NUL

	return buffer.Bytes()
}
func Send(DNSServer, domain string) ([]byte, int, time.Duration) {
	requestHeader := DNSHeader{
		Id:           0x0010,
		Questions:    1, // 请求报文只需要这个就可以了吧
		AnswerRRs:    0,
		AuthorityRRS: 0,
		AdditionRRS:  0,
	}
	requestHeader.SetFlag(0, 0, 0, 0, 1, 0, 0)

	requestQueries := DNSQueries{
		Type:  1,
		Class: 1,
	}

	var (
		conn   net.Conn
		err    error
		buffer bytes.Buffer
	)

	if conn, err = net.Dial("udp", DNSServer); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	defer conn.Close()

	domainB := ParseDomainName(domain)
	fmt.Println(len(domainB), domainB)
	binary.Write(&buffer, binary.BigEndian, requestHeader)
	binary.Write(&buffer, binary.BigEndian, domainB)
	binary.Write(&buffer, binary.BigEndian, requestQueries)

	buf := make([]byte, 1024)
	t1 := time.Now()
	if _, err := conn.Write(buffer.Bytes()); err != nil {
		fmt.Println(err.Error())
		return make([]byte, 0), 0, 0
	}
	length, err := conn.Read(buf)
	t := time.Now().Sub(t1)
	return buf, length, t
}
func main() {
	remsg, n, _ := Send("8.8.8.8:53", "www.baidu.com")
	fmt.Println("responseLen:  ", n)

	// Header
	fmt.Printf("ID:%#18x\n", remsg[:2])            // ID
	fmt.Printf("Flags:%#15x\n", remsg[2:4])        // Flags
	fmt.Printf("Questions:%#11x\n", remsg[4:6])    // Questions
	fmt.Printf("AnswerRRs:%#11x\n", remsg[6:8])    // AnswerRRs
	fmt.Printf("AuthorityRRs:%#8x\n", remsg[8:10]) // AuthorityRRs
	fmt.Printf("AdditionRRs:%#9x\n", remsg[10:12]) // AdditionRRs

	fmt.Println(remsg[12:27]) // Queries

	fmt.Println(remsg[27:]) // RRs
}

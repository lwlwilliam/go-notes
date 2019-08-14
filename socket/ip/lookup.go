// Host lookup
//
// 解析返回多个 IP 地址
// func LookupHost(name string) (addrs []string, err os.Error)
//
// 返回 canonical 主机
// func LookupCNAME(name string) (cname string, err os.Error)
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	addrC, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	fmt.Println("CNAME:", addrC)

	os.Exit(0)
}

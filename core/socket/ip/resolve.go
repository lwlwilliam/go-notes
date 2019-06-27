// ResolveIPAddr 解析 IP 地址
//
// type IPAddr {
//     IP IP
// }
//
// 对 IP 主机名执行 DNS 查询
// func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
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

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is", addr.String())

	fmt.Println(addr.Network())
	fmt.Println(addr.Zone)

	os.Exit(0)
}

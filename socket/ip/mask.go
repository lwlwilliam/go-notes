// IPmask
//
// type IPMask []byte
//
// func IPv4Mask(a, b, c, d byte) IPMask
//
// func (ip IP) DefaultMask() IPMask
//
// func (ip IP) Mask(mask IPMask) IP
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}

	mask := addr.DefaultMask()	// 默认掩码：255.255.255.0

	// 自定义掩码
	//mask = net.IPv4Mask(255, 255, 252, 0)
	//fmt.Println(mask)

	network := addr.Mask(mask)	// 地址和掩码进行"与运算"，得到网络地址
	ones, bits := mask.Size()	// 掩码中位为 1 的个数以及总位数

	fmt.Fprintf(os.Stdout, "Address is %s\n", addr.String())
	fmt.Fprintf(os.Stdout, "Default mask length is %d\n", bits)
	fmt.Fprintf(os.Stdout, "Leading ones count is %d\n", ones)
	fmt.Fprintf(os.Stdout, "Mask is (hex) %s\n", mask.String())
	fmt.Fprintf(os.Stdout, "Network is %s\n", network.String())

	os.Exit(0)
}

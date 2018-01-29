// 迫不得已，把题目改了。。。TODO: 重做
package main

import (
	"fmt"
	"strings"
)

//type IPAddr [4]byte
type IPAddr []string

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return strings.Join(ip, ".")
}

func main() {
	addrs := map[string]IPAddr{
		//"loopback": {127, 0, 0, 1},
		//"goobleDNS": {8, 8, 8, 8},
		"loopback": {"127", "0", "0", "1"},
		"goobleDNS": {"8", "8", "8", "8"},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

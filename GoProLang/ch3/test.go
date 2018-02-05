package main

import "fmt"

func main() {
	var u uint8 = 1
	fmt.Println(u << 1, u >> 1, u & 1, u | 1, u ^ 1)
}

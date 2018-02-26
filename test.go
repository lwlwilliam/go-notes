package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	var n = make(map[string]int)

	var o = map[string]int{}

	fmt.Println(m, n, o)
}

package main

import (
	"fmt"
	"./geometry"
)

func main() {
	var m = geometry.Point{6, 7}
	var n = geometry.Point{3, 3}
	fmt.Println(m.Distance(n))
}

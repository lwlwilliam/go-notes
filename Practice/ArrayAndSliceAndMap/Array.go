package main

import (
	"fmt"
)

func main() {
	var x [5]int
	x[0] = 2
	x[1] = 3
	x[2] = 3
	x[3] = 2
	x[4] = 12
	var sum int
	for _, elem := range x {
		sum += elem
	}
	fmt.Println(sum)

	var a = [3]int{1, 2, 3}
	for _, elem := range a {
		sum += elem
	}
	fmt.Println(sum)

	var b = [3]int{}
	b[0] = 1
	b[1] = 2
	b[2] = 3
	for _, elem := range b {
		sum += elem
	}
	fmt.Println(sum)

	var c = [...]int{1, 2, 3}
	for _, elem := range c {
		sum += elem
	}
	fmt.Println(sum)
}

package main

import (
	"fmt"
)

func main() {
	// const beef, two, c = "meat", 2, "veg"
	// const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

	const (
		Sunday, Monday, Tuesday, Wednesday = 0, 1, 2, 3
		Thursday, Friday, Saturday = 4, 5, 6
	)

	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	const (
		a = "test"
		b = iota + 10
		c
		d
		e = iota
		f
		g
	)

	const (
		A = iota
		B
		C
	)

	fmt.Println(a, b, c, d, e, f, g, A, B, C, Unknown, Female, Male, Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)


	
}

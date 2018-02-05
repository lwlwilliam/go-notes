package main

import (
	"fmt"
)

func main() {
	const x string = "hello world"
	const y = "hello world"

	fmt.Println(x)
	fmt.Println(y)

	var (
		a int     = 10
		b float64 = 32.45
		c bool    = true
	)

	const (
		Pi float64 = 3.14
		True bool  = true
	)

	fmt.Println(a, b, c)
	fmt.Println(Pi, True)
}

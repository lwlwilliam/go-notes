package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		println(f, "   ", 1)
		return float64(-f)
	}
	println(f)
	return float64(f)
}

func main() {
	println("math.Sqrt2 = ", math.Sqrt2)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

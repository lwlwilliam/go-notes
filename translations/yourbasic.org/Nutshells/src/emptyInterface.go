package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	var x interface{}

	x = 2.4
	fmt.Println(x)

	x = &Point{1, 2}
	fmt.Println(x)
}

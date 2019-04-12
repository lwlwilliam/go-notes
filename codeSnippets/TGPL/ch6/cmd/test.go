package main

import "fmt"

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	Point
}

func main()  {
	var c ColoredPoint
	c.X = 3
	fmt.Println(c)
}

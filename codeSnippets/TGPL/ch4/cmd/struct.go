package main

import "fmt"

type Point struct {
	X, Y int
}

func main()  {
	p := Point{Y:2}
	fmt.Println(p)

	pt := Point{1, 2}
	fmt.Println(pt)
}

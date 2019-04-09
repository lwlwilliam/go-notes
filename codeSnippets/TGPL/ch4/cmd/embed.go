package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
	int	// 非结构体匿名成员
}

func main()  {
	var w Wheel

	w = Wheel{Circle{Point{8, 8}, 5}, 20, 3 /* 非结构体匿名成员 */}

	w = Wheel{
		Circle: Circle{
			Point: Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
		int: 4,	// 非结构体匿名成员
	}

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}

package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func double_area(rect Rect) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

func main() {
	var rect = Rect{100, 200}
	fmt.Println(double_area(rect))
	fmt.Println("Width: ", rect.width, "Length: ", rect.length)
}

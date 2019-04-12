package main

import (
	"image/color"
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main()  {
	red := color.RGBA{255, 0, 0, 0}
	blue := color.RGBA{0, 0, 255, 0}
	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point))	// p and q now share the same Point
	q.Point = p.Point
	*(p.Point) = Point{8, 8}
	fmt.Println(*p.Point, *q.Point)
}

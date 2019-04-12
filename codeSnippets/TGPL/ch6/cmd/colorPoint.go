package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p *Point) Test() {
	fmt.Println(p)
}

type Conflict struct {
	X string
}

func (c *Conflict) Test() {
	fmt.Println(c)
}

type ColoredPoint struct {
	Point
	Conflict
	Color color.RGBA
}

func main()  {
	var c ColoredPoint
	c.Point.X = 3 // 如果没有嵌入 Conflict 的话，可以直接 color.X 进行赋值
	c.Y = 4
	fmt.Println(c)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, Conflict{"red"}, red}
	var q = ColoredPoint{Point{1, 1}, Conflict{"blue"}, blue}
	fmt.Println(p.Distance(q.Point))
	p.Conflict.Test() // 因为 Point 和 Conflict 都有 Test 方法，所以不能 p.Test() 调用
}

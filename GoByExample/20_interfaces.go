/*
接口是方法特征的命名集合
要在 Go 中实现一个接口，只需要实现接口中的所有方法
 */
package main

import (
	"math"
	"fmt"
)

// 这里是一个几何体的基本接口
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// rect 实现 geometry 接口
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

// circle 实现 geometry 接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量是接口类型，那么可以调用这个被命名的接口中的方法。
// 利用这个特性，它可以在任何 geometry 上
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// 结构体类型 circle 和 rect 都实现了 geometry 接口，所以可以用它们的实例
	// 作为 measure 的参数
	measure(r)
	measure(c)
}

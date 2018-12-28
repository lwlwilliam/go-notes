package main

import (
	"fmt"
	"strconv"
)

type Test interface{}

type MyString interface {
	String() string
}

type Temp int

type Point struct {
	x, y int
}

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " 摄氏度"
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
	var x MyString
	fmt.Printf("%v %T\n", x, x)

	x = Temp(27)
	fmt.Printf("%v %T\n", x, x)

	x = &Point{1, 6}
	fmt.Printf("%v %T\n", x, x)

	var y Test
	y.Haha()
}

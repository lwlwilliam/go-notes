package main

import (
	"strconv"
	"fmt"
)

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

	x = Temp(27)
	fmt.Println(x)

	x = &Point{1, 6}
	fmt.Println(x)
}

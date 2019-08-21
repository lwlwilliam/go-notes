package main

import (
	"fmt"
)

type MyString interface {
	String() string
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
	var x MyString
	fmt.Printf("%v, %T; %v, %T; %v\n", x, x, nil, nil, x == nil)

	x = (*Point)(nil)
	fmt.Printf("%v, %T; %v, %T; %v\n", x, x, nil, nil, x == nil)
}
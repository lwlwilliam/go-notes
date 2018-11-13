package main

import (
	"fmt"
)

type MyType struct {
	n int
}

func (p *MyType) Value() int { return p.n }

func main() {
	pm := new(MyType)
	pm = &MyType {
		n: 666,
	}

	fmt.Println(pm.Value())		// 0 (零值)
}

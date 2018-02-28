package main

import (
	"fmt"
)

func main() {
	_ = double(4)	
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

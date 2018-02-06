package main

import (
	"fmt"
)

func set_value(x_ptr *int) {
	*x_ptr = 100
}

func main() {
	x_ptr := new(int)
	set_value(x_ptr)

	// x_ptr 指向的地址
	fmt.Println(x_ptr)

	// x_ptr 本身的地址
	fmt.Println(&x_ptr)

	// x_ptr 指向的地址值
	fmt.Println(*x_ptr)
}

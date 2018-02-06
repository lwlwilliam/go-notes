package main

import (
	"fmt"
)

func not_change(x int) {
	x = 200
}

func change(x *int) {
	*x = 200
}

func main() {
	var x int = 100
	fmt.Println(x)
	not_change(x)
	fmt.Println(x)
	change(&x)
	fmt.Println(x)
}

package main

import (
	"fmt"
	haha "../test/subtest"
)

func main() {
	var a = 100
	var b = 200

	fmt.Println("Add demo: ", haha.Add(a, b))
	fmt.Println("Subtract demo: ", haha.Subtract(a, b))
	fmt.Println("Multiply demo: ", haha.Multiply(a, b))
	fmt.Println("Divide demo: ", haha.Divide(a, b))
}

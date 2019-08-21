package main

import (
	"fmt"
)

func main() {
	var stack []string

	stack = append(stack, "world!")  // Push
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		n := len(stack) - 1	// Top element
		fmt.Println(stack[n])

		stack = stack[:n]	// Pop
	}
}

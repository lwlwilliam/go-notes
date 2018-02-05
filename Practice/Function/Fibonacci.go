package main

import "fmt"

func fibonacci(n int) int {
	var retVal = 0
	if n == 1 {
		retVal = 1
	} else if n == 2 {
		retVal = 2
	} else {
		retVal = fibonacci(n - 2) + fibonacci(n - 1)
	}
	return retVal
}

func main() {
	fmt.Println(fibonacci(5))
}

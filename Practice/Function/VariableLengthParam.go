package main

import "fmt"

func sum(arr ...int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
}

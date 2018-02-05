package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}

	var sum = func(arr ...int) int {
		total_sum := 0
		for _, val := range arr {
			total_sum += val
		}
		return total_sum
	}

	fmt.Println(sum(arr...))
}

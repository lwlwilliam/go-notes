package main

import "fmt"

func slice_sum(arr []int) (int, float64) {
	sum := 0
	avg := 0.0
	for _, elem := range arr {
		sum += elem
	}
	avg = float64(sum) / float64(len(arr))
	return sum, avg
}

func main() {
	var arr = []int{3, 2, 3, 1, 6, 4, 8, 9}
	fmt.Println(slice_sum(arr))
}

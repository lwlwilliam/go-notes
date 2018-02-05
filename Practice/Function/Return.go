package main

import "fmt"

func slice_sum(arr []int) (sum int) {
	sum = 0
	for _, elem := range arr {
		sum += elem
	}
	return
}

func main() {
	var arr1 = []int{1, 3, 2, 3, 2}
	var arr2 = []int{3, 2, 3, 1, 6, 4, 8, 9}
	fmt.Println(slice_sum(arr1))
	fmt.Println(slice_sum(arr2))
}

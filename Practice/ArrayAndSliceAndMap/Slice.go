package main

import (
	"fmt"
)

func main() {
	var x = make([]float64, 5)
	fmt.Println("Capacity: ", cap(x), "Length: ", len(x))

	var y = make([]float64, 5, 10)
	fmt.Println("Capacity: ", cap(y), "Length: ", len(y))

	for i := 0; i < len(x); i ++ {
		x[i] = float64(i)
	}
	fmt.Println(x)

	for i := 0; i < len(y); i ++ {
		y[i] = float64(i)
	}
	fmt.Println(y)

	var arr = [5]int{1, 2, 3, 4, 5}
	var slice = arr[1:3]
	fmt.Println(slice)
	slice = append(slice,4, 5, 6)
	fmt.Println(slice)

	var slice2 = arr[3:]
	copy(slice, slice2)
	fmt.Println(slice)
}


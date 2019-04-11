package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main()  {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum([]int{1, 2, 3, 4, 5}...))
	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", test)
}

func test([]int)  {

}

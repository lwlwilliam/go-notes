/*
变参函数。可变参数函数。在调用时可以用任意数量的参数。
fmt.Println 是一个常见的变参函数
 */
package main

import "fmt"

// 这个函数接受任意数目的 int 作为参数
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// 函数使用常规的调用方式，传入独立的参数
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// 如果有一个含有多个值的 slice，想把它作为参数使用
	sum(nums...)
}

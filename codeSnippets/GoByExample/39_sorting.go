/*
Go 的 sort 包实现了内置和用户自定义数据类型的排序功能
 */
package main

import (
	"sort"
	"fmt"
)

func main() {
	// 排序方法是正对内置数据类型的；
	// 这是一个字符串的例子。排序会改变给定的序列，并不返回新值
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:  :", ints)

	//ints := []int{1, 2, 3}
	// 检查一个序列是否已经是排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted :", s)
}

/*
在 Go 中，数组是一个具有固定长度且编号的元素序列
 */
package main

import "fmt"

func main() {
	// 创建一个数组 a 来存放刚好 5 个 int。元素的类型和长度是数组类型的一部分。数组黑夜是零值的，对于 int 数组来说就是 0。
	var a [5]int
	fmt.Println("emp:", a)

	// custom
	a = [5]int{1, 2, 3, 4, 5}

	// 通过 array[index] = value 的语法来设置数组指定位置的值，用 array[index] 获取值。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 使用内置函数 len 返回数组长度，cap 返回数组容量。
	fmt.Println("len:", len(a))
	fmt.Println("cap:", cap(a))

	// 声明并初始化一个数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i ++ {
		for j := 0; j < 3;  j ++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

package main

import (
    "fmt"
)

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1

    // 容量足够，直接扩展 slice（依然在原有的底层数组之上），将新添加的 y 元素复制到新扩展的空间，并返回 slice。
	// 因此，输入的 x 和输出的 z 共享相同的底层数组。
    if zlen <= cap(x) {
        // There is room to grow. Extend the slice.
        z = x[:zlen]

	// 容量不够
    } else {
        // There is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
		// 容量不够，分配一个足够大的 slice 用于保存新的结果，先将输入的 x 复制到新的空间，然后添加 y 元素。
		// 结果 z 和输入的 x 引用的将是不同的底层数组。
        zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)  // a built-in function; see text
    }
	z[len(x)] = y
	return z
}

func main() {
    var x, y []int
    for i := 0; i < 10; i ++ {
        y = appendInt(x, i)
        fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
        x = y
    }
}

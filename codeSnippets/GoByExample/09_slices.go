package main

import "fmt"

func main() {
	// 与数组不同，slice 的类型仅由它所包含的元素决定（不需要元素的个数）。要创建一个长度非零的空 slice，需要使用
	// 内建方法 make。这里创建一个长度为 3 的 string 类型的 slice（初始化为零值）。
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 可以和数组一个设置和获取值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len 返回 slice 的长度
	fmt.Println("len:", len(s))

	// 内建的 append 返回一个包含了一个或多个新值的 slice。append 可能返回新的 slice，需要接受其返回值。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice 可以被 copy。将 s 复制给 c。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice 支持通过 slice[low:high] 语法进行"切片"操作。
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slice 可以组成多维数据结构。内部的 slice 长度可以不一致。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i ++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j ++ {
			twoD[i][j] = i + j
		}
	}

	// 注意，slice 和数组是不同的类型，但是通过 fmt.Println 打印结果类似。
	fmt.Println("2d:", twoD)
}

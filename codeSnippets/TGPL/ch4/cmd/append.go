package main

import (
	"fmt"
	append2 "github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch4/append"
)

func main()  {
	var x, y []int
	for i := 0; i < 10; i ++ {
		y = append2.AppendInt(x, i)

		fmt.Printf("%d cap = %d\t%v\n", i, cap(y), y)
		x = y
	}

	var m []int
	m = append2.AppendInt2(m, []int{1, 2, 3, 4, 5, 6, 7}...)
	m = append2.AppendInt2(m, []int{1, 2, 3, 4, 5, 6, 7}...)
	fmt.Println(m)
}

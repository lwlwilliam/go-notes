package main

import (
	"fmt"
	append2 "github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch4/append"
)

func main()  {
	arr := [2]int{1, 2}
	a := arr[0:1]
	a = append2.AppendInt(a, 1)
	fmt.Println(a)
}

package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch6/intList"
	"fmt"
)

func main()  {
	list2 := intList.IntList{
		Value: 2,
		Tail: nil,
	}

	list := intList.IntList{
		Value: 1,
		Tail: &list2,
	}

	fmt.Println(list.Sum())
}

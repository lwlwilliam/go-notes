package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch4/sliceStack"
)

func main() {
	s := sliceStack.Stack{}
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	a, err := s.Remove(0)
	if err == nil {
		fmt.Println(a)
	}
}

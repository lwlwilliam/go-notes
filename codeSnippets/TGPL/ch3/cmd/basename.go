package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch3/basename"
)

func main()  {
	fmt.Println(basename.Basename1("/home/go/src/test.go"))
	fmt.Println(basename.Basename2("/home/go/src/test.go"))
}

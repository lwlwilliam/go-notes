package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/fetch"
	"fmt"
	"os"
)

func main()  {
	fmt.Println(fetch.Fetch("https://github.com/"))
	fmt.Println(os.Getwd())
}
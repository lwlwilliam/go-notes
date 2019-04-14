package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch7/bytecounter"
)

func main()  {
	var b bytecounter.ByteCounter
	b.Write([]byte("Hello"))
	fmt.Println(b)

	b = 0 // reset the counter
	var name = "William"
	fmt.Fprintf(&b, "Hello, %s\n", name)
	fmt.Println(b)
}

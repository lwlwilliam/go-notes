package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch2/tempconv"
)

func main()  {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.BoilingC - tempconv.AbsoluteZeroC)
}

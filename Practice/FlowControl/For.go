package main

import (
	"fmt"
)

func main() {
	var i int = 1

	for ; i <= 5; i ++ {
		fmt.Println(i)
	}

	for i := 1; i <= 5; i ++ {
		fmt.Println(i)
	}

	for i > 0 {
		fmt.Println(i)
		i --
	}

	// 死循环
	//for {
	//	...
	//}
}

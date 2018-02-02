// 调用自定义包
package main

import (
	"fmt"
	"./tempconv"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))

	fmt.Println(tempconv.FreezingC)
	fmt.Println(tempconv.CToF(tempconv.FreezingC))

	fmt.Println(tempconv.BoilingC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}

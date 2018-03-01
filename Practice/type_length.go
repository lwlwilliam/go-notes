package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 12
	fmt.Println("length of a: ", unsafe.Sizeof(a))
	var b int = 12
	fmt.Println("length of b(int): ", unsafe.Sizeof(b))
	var c int8 = 12
	fmt.Println("length of c(int): ", unsafe.Sizeof(c))
	var d int16 = 12
	fmt.Println("length of d(int): ", unsafe.Sizeof(d))
	var e int32 = 12
	fmt.Println("length of e(int): ", unsafe.Sizeof(e))
	var f int64 = 12
	fmt.Println("length of f(int): ", unsafe.Sizeof(f))
}

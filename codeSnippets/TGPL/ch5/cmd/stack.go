package main

import (
	"runtime"
	"os"
	"fmt"
)

func main()  {
	defer printStack()
	f(3)
}

func f(x int)  {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	//if x > 1 {
		f(x - 1)
	//}
}

func printStack()  {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
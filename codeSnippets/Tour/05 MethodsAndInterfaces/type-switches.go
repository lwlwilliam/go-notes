package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v %T is %v.\n", v, v, v*2)
	case string:
		fmt.Printf("%q %T is %v bytes long.\n", v, v, len(v))
	default:
		fmt.Printf("I don't know about type %T.\n", v)
	}
}

func main() {
	do(21)
	do("Hello")
	do(true)
}

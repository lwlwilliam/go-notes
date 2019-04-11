package main

import "fmt"

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

func main()  {
	fmt.Println(triple(8))
}

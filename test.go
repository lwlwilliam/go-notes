package main

import "fmt"

func main() {
	s := "128"
	n := 0
	for i := 0; i < len(s); i ++ {
		n *= 10
		n += (int(s[i]) - int('0'))
	}

	fmt.Printf("%d\t%T\n", n, n)
}
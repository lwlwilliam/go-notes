package main

import (
	"fmt"
)

func main() {
	s := "hello, world"
	t := s
	s[4] = 'K'
	fmt.Println(s)
}

package main

import "fmt"

func first() {
	fmt.Println("first func run")
}

func second() {
	fmt.Println("second func run")
}

func main() {
	defer second()
	first()
}

package main

import "fmt"

func recoverPanic() (res interface{}) {
	defer func() {
		if p := recover(); p == nil {
			res = "error"
		} else {
			res = p
		}
	}()

	panic(nil)
}

func main()  {
	fmt.Println(recoverPanic())
}

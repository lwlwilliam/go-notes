package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan struct{})	// 禁用该 channel
	ch1 = nil
	ch2 := make(chan int)

	go func() {
		ch2 <- 1
	}()
		
	select {
	case <- ch1:
		fmt.Println("Received from ch1")	// 不会执行
	case <- ch2:
		fmt.Println("Received from ch2")
	}
}

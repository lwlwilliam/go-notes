package main

import (
	"fmt"
	"time"
)

func main() {
	AFP := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		AFP <- "Hello go!"
	}()

	select {
	case news := <- AFP:
		fmt.Println(news)
	case <- time.After(time.Second):
		fmt.Println("Time out: No news in one second.")
	/*
	default:
		fmt.Println("Out")
	*/
	}
}

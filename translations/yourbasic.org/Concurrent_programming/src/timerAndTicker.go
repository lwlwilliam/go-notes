package main

import (
	"fmt"
	"time"
)

func main() {
	AFP := make(chan int)

	// Timeout
	go func() {
		select {
		case news := <- AFP:
			fmt.Println(news)
		case <- time.After(time.Second):
			fmt.Println("No news in a second.")
		}
	}()

/*

	// time.NewTimer
	go func() {
		for alive := true; alive; {
			timer := time.NewTimer(time.Second)
			select {
			case news := <- AFP:
				timer.Stop()
				fmt.Println(news)
			case <- timer.C:
				alive = false
				fmt.Println("No news in a second. Service aborting.")
			}
		}
	}()
	

	// time.Tick
	/*
	go func() {
		for now := range time.Tick(time.Second) {
			fmt.Println(now, statusUpdate())
		}
	}()
	*/


/*
	go func() {

		timer := time.AfterFunc(time.Second, func() {
			fmt.Println("Foo run for more than a second.")
		})
		defer timer.Stop()

		// Do heavy work
	}()
	*/

	time.Sleep(10 * time.Second)
}

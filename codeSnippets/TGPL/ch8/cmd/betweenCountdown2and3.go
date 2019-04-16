package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	abort := make(chan bool, 1)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- true
	}()

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <- time.After(10 * time.Second):
		// Do nothing
	case <- abort:
		fmt.Println("Launch aborted!")
		return
	}
	fmt.Println("Launch")
}

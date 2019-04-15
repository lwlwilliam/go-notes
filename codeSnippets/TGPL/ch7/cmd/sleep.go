package main

import (
	"flag"
	"time"
	"fmt"
)

// ./sleep -period 3s
// ./sleep -period 1m
var period = flag.Duration("period", 1 * time.Second, "sleep period")

func main()  {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

	var t time.Duration
	t = time.Minute
	fmt.Println(t)
}

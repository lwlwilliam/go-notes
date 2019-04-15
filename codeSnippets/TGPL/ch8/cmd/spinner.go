package main

import (
	"fmt"
	"time"
)

func main()  {
	go spinner(100 * time.Millisecond)
	const n = 45
	start := time.Now()
	fibN := fib(n) // slow
	total := time.Since(start)
	fmt.Printf("\rFibonacci(%d) = %d; execution time: %v\n", n, fibN, total)
}

func spinner(delay time.Duration)  {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int  {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}

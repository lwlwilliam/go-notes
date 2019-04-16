package main

import "fmt"

func main()  {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for x := 0; x <= 100 ; x ++ {
			naturals <- x
		}
		close(naturals)
	}()

	// squarer
	go func() {
		for {
			x, ok := <- naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// printer (in main goroutine)
	for {
		res, ok := <- squares
		if !ok {
			break
		}
		fmt.Println(res)
	}
}

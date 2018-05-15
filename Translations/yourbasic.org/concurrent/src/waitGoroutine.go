package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var test = wg	// 这时候未使用，可以复制

	wg.Add(2)
	// var test  = wg  // WaitGroup 使用后不能复制

	go func() {
		fmt.Println("go A")
		wg.Done()
	}()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("go B")
		wg.Done()
	}()


	test.Add(1)
	go func() {
		fmt.Println("copy test")
		test.Done()
	}()

	wg.Wait()
	test.Wait()
}

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 666  // 死锁，超过缓冲区大小
	fmt.Println(<-ch)
	ch <- 3  // 不会死锁，因为缓冲区已经释放了一个空位
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

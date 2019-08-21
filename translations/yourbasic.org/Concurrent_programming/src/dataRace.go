package main

import "fmt"

func race() {
	wait := make(chan struct{})
	n := 0
	go func() {
		n ++	// 读取，递增，读入
		close(wait)
	}()
	n ++	// 访问冲突
	<- wait
	fmt.Println(n)	// 输出: <unspecified>
}

func sharingIsCaring() {
	ch := make(chan int)
	go func() {
		n := 0	// 局部变量只在一个 goroutine 中可见
		n ++
		ch <- n	// 数据离开 goroutine
	}()
	n := <- ch	// ...（译者注：数据）安全地到达另一个。（译者注：goroutine）
	n ++
	fmt.Println(n)	// 输出：2
}

func main() {
	race()
	sharingIsCaring()
}

package main

import "fmt"

// Generator 返回一个整数类型的 channel，传递 1, 2, 3 等整数
// 为了令以下 goroutine 停止，要关闭 channel
func Generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		// 循环往 channel 写入读取数据直到 channel 关闭
		for {
			select {
			case ch <- n:
				n ++
			case <- ch:
				return
			}
		}
	}()
	return ch
}

func main() {
	number := Generator()
	fmt.Println(<- number)
	fmt.Println(<- number)
	fmt.Println(<- number)
	fmt.Println(<- number)
	fmt.Println(<- number)
	fmt.Println(<- number)
	close(number)
	// ...
}

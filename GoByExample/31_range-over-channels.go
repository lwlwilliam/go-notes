/*
可以用 for 和 range 遍历从通道中取得的值
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 注意，要使用缓冲通道，否则会死锁
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range 迭代从 queue 中得到的每个值。因为在前面 close 了这个通道，这个迭代会在接收完 2 个值之后结束
	// 如果没有 close，将在这个循环中继续阻塞执行，等待接收第三个值
	//go func() {
		for elem := range queue {
			fmt.Println(elem)
		}
	//}()

	//time.Sleep(time.Second)
}

/*
通道是连接多个 Go 协程的管道。可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收

默认发送和接收操作阻塞的，直到发送方和接收方都准备完毕
这个特性允许我们不使用任何其他的同步操作
 */
package main

import "fmt"

func main() {
	// 使用 make(chan val-type) 创建一个新的通道。通道类型就是需要传递值的类型
	messages := make(chan string)

	// 使用 channel <- 语法发送一个新的值到通道中
	go func() {
		messages <- "ping"
	}()

	// 俣用 <- channel 语法从通道中接收一个值
	msg := <- messages
	fmt.Println(msg)
}

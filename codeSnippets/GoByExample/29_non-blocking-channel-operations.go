/*
常规的通过通道发送和接收数据是阻塞的。然而，可以使用带一个 default 子句的 select 来实现
非阻塞的发送、接收，甚至是非阻塞的多路 select
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// custom
	{
		go func() {
			messages <- "test"
		}()

		go func() {
			signals <- true
		}()

		time.Sleep(time.Second * 2)
	}

	// 非阻塞接收。如果 messages 中存在数据，select 将这个值带入 <- messages case 中；否则，直接到 default 分支中
	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
		default:
			fmt.Println("no message receive")
	}

	// 一个非阻塞发送的实现方法和上面一样
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
	}

	// 可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	// 这里试图在 messages 和 signals 上同时使用非阻塞的接受操作
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
		case sig := <- signals:
			fmt.Println("received signal", sig)
			default:
				fmt.Println("no activity")
	}
}

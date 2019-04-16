// 让程序支持在倒计时中，用户按下 return 键时直接中断发射流程
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	for count := 10; count > 0; {
		select {
		case <-tick:
			fmt.Println(count)
			count--
		case <-abort:
			goto Exit
		default:
			//log.Println("block")
		}
	}
Exit:
}

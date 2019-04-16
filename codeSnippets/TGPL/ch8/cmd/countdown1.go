// 模拟火箭发射的倒计时
// time.Tick 函数返回一个 channel，程序会周期性地像一个节拍器一样向这个 channel 发送事件
// 每一个事件的值是一个时间戳，不过更有意思的是其传送方式
package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown -- {
		fmt.Println(countdown)
		<- tick
	}
	fmt.Println("launch")
}

// ch 的 buffer 大小是 1，所以会交替的为空或为满，所以只有一个 case 可以进行下去，无论 i 是奇数或者偶
// 偶数，它都会打印 0 2 4 6 8
package main

import "fmt"

func main()  {
	ch := make(chan int, 1)
	for i := 0; i < 10; i ++ {
		select {
		case x := <- ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

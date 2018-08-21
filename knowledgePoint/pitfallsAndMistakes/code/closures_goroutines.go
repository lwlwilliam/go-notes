package main
import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本 A
	for ix := range values {
		func() {
			fmt.Print(ix, " ")
		}()  // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本 B：和 A 版本类似，但是通过调用闭包作为一个协程
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	time.Sleep(5e9)
	fmt.Println()
	
	// 版本 C：正确的处理方式
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	time.Sleep(5e9)
	fmt.Println()
	// 版本 D：输出值：
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(5e9)
	fmt.Println()
}

// 这个示例程序展示如何创建 goroutine 以及 goroutine 调度器的行为
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main 是所以 Go 程序的入口
func main() {
	// 分配 2 个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(runtime.NumCPU())

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count ++ {
			for char := 'a'; char < 'a' + 26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		// 显示字母表 3 次
		for count := 0; count < 3; count ++ {
			for char := 'A'; char < 'A' + 26; char ++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}

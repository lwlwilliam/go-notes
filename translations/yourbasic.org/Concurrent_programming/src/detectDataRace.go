package main

import "fmt"

func main() {
	i := 0
	go func() {
		i ++	// 写入
	}()
	fmt.Println(i)	// 并发读取
}

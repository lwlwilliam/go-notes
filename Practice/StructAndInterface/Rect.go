package main

import (
	"fmt"
)

type Rect struct {
	width, length float64
}

func main() {
	var rect Rect
	rect.width = 100
	rect.length = 200
	fmt.Println(rect.width * rect.length)


	// 初始化方式赋值
	var rect2 = Rect{width: 100, length: 200}
	fmt.Println(rect2.width * rect2.length)


	// 如果知道结构体成员定义的顺序，也可以不用 key:value 的方式赋值，直接按定义顺序赋值
	var rect3 = Rect{100, 200}
	fmt.Println(rect3.width * rect3.length)
}

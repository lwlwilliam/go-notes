package pkg

import "fmt"

// 不可访问
var i = 123
var f  = 12.3

// 可访问
var S = "Hello world"

// 可访问
type Stru struct {
	A string
	b string  // 不可访问
}

// 类型可访问
type T int

// 方法不可访问
func (t *T) test() {
	fmt.Println("customize method")
}

// 方法可访问
func (t *T) Haha() {
	fmt.Println("Haha")
}

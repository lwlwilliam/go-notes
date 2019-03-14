/*
Go 支持指针，允许在程序中通过引用传递值或者数据结构

将通过两个函数：zeroval 和 zeroptr 来比较指针和值类型的不同。
 */
package main

import "fmt"

// 得到一个 ival 形参的拷贝
func zeroval(ival int) {
	ival = 0
}

// 得到一个 int 指针。函数体内的 *iptr 解引用这个指针，从它的内存地址得到这个地址对应的当前值。
// 对一个解引用的指针赋值将会改变这个指针引用的真实地趣的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	// 不能改变 i 的值
	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针
	// 可以改变 i 的值
	zeroptr(&i)
	fmt.Println("zeroval:", i)

	fmt.Println("pointer:", &i)
}

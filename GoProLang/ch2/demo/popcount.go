/**
	用于返回一个数字中含二进制 1 bit 的个数。它使用 init 初始化函数来生成辅助表格 pc，
	pc 表格用于处理每个 8 bit 宽度的数字含二进制的 1 bi 的 bit 个数，这样的话在处理
	64 bit 宽度的数字时就没有必要循环 64 次，只需要 8 次查表就可以了。（这并不是最快
	的统计 1 bit 数目的算法，但是它可以方便演示 init 函数的用法，并且演示了如何预生在
	辅助表格，这是编程中常用的技术）。
 */
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (nuber of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 译注：对于 pc 这类需要复杂处理的初始化，可以通过将初始化逻辑包装为一个匿名函数处理，像下面这样：
/*
var pc [256]byte = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}
 */

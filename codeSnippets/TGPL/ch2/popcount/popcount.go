// TODO: 有空看一下啊。。。这操作看不懂
package popcount

import "fmt"

var pc [256]byte

func init()  {
	for i := range pc {
		// i/2 = 0, 0, 1, 1, 2, 2, ... 127, 127
		// i&1 = 0, 1, 0, 1, 0, 1, ... 0, 1
		pc[i] = pc[i/2] + byte(i&1)
	}
	fmt.Println(pc)
}

// 返回一个数字中含二进制 1 bit 的个数
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

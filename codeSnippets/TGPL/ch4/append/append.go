package append

// appendInt 函数，专门用于处理 []int 类型的 slice
//
func AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	// 容量足够
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]

	// 容量不足
	// zlen > cap(x)
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortize linear complexity.
		// 分配两位容量
		zcap := zlen
		// 以防 len(x) 为 0 的情况
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

package append

// appendInt 函数，专门用于处理 []int 类型的 slice
// 内置的 append 函数可能使用比 AppendInt 更复杂的内存扩展策略。因此，通常我们并不知道 append 调用是否导致了内存的重新分配，
// 因此我们也不能确认新的 slice 和原始的 slice 是否引用的是相同的底层数组空间。同样，我们不能确认在原先的 slice 上的操作是否
// 会影响到新的 slice。因此，通常是将 append 返回的结果直接赋值给输入的 slice 变量：
// runes = append(runes, r)
// 更新 slice 变量不仅对调用 append 函数是必要的，实际上对应任何可能导致长度、容量或底层数组变化的操作都是必要的。
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

// 可追加多个值
func AppendInt2(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	copy(z[len(x):], y)
	return z
}

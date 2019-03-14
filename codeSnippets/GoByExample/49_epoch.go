/*
Go 实现获取 Unix 时间的秒数，毫秒数，或者微秒数
 */
package main

import "fmt"
import "time"

func main() {
	now := time.Now()

	// 分别使用带 Unix 或者 UnixNano 的 time.Now 来获取从自 协调世界时 起到现在的秒数或纳秒数
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println("1.", now)

	// UnixMillis 是不存在的，所以要从纳秒转化一下
	millis := nanos / 1000000
	fmt.Println("2.", secs)
	fmt.Println("3.", millis)
	fmt.Println("4.", nanos)

	// 也可以将 协调世界时 起的整数秒或者纳秒转化到相应的时间
	fmt.Println("5.", time.Unix(secs, 0))
	fmt.Println("6.", time.Unix(0, nanos))
}

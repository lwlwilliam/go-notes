// 测量潜在低效的版本和使用了 strings.Join 的版本的时间差异
// TODO: 这个目前没有什么好办法测试，字符串没到一定长度都相差不大
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	fmt.Println(time.Now().UnixNano())
	fmt.Println()
	var s, sep string
	for i := 1; i < len(os.Args); i ++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(time.Now().UnixNano())
	fmt.Println()
	strings.Join(os.Args[1:], " ")
	fmt.Println(time.Now().UnixNano())
}

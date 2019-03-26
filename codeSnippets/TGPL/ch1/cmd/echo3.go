// 如果连接涉及的数据量很大，如 echo1 和 echo2 中用 += 连接字符串，每次循环迭代字符串 s 的内容都会更新
// s 原来的内容已经不再使用，需要在适当时机对它进行垃圾回收
// 以下是一种简单且高效的解决方案
package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

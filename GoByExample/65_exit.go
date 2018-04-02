/*
使用 os.Exit 来立即进行带给定状态的退出

注意，不像 C 语言等，Go 不使用在 main 中返回一个整数来指明退出状态。
如果想以非零状态退出，那么就要使用 os.Exit

如果使用 go run 来运行该程序，那么退出状态将会被 go 捕获并打印，
使用编译并执行一个二进制文件的格式，可以在终端中查看退出状态
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当使用 os.Exit 时 defer 将不会执行，所以这里的 fmt.Println 将永远不会被调用
	defer fmt.Println("!")

	// 退出并且退出状态为 3
	os.Exit(3)
}

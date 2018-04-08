// 这个示例程序展示如何使用 io.Reader 和 io.Writer 接口
// 写一个简单的 curl 程序
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init 在 main 函数之前调用
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

func main() {
	// 从 web 服务器得到响应
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 body 复制到 stdout
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

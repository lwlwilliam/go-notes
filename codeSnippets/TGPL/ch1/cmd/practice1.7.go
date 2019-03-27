// 使用 io.Copy 替换 fetch 中的 ioutil.ReadAll，避免申请一个缓冲区（变量 b）来存储
package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)

func main()  {
	count := os.Args[1:]

	for _, url := range count {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTTP GET error: %v\n", err)
			continue
		}

		//b, err := ioutil.ReadAll(resp.Body)
		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Copy error:", err)
		}
		fmt.Println()
		fmt.Printf("Fetch %d bytes.", n)

		// 打印响应头部
		//respHeader := resp.Header
		//for key, val := range respHeader {
		//	fmt.Printf("%s: %v\n", key, val)
		//}

		err = resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTTP body close error: %v\n", err)
		}
		fmt.Println()
	}
}

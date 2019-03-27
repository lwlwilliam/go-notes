// 如果输入的 url 参数没有 http:// 前缀的话，为这个 url 加上该前缀
package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

func main()  {
	count := os.Args[1:]

	for _, url := range count {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTTP GET error: %v\n", err)
			continue
		}

		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("Copy error:", err)
		}
		fmt.Println()
		fmt.Printf("Fetch %d bytes.", n)

		err = resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "HTTP body close error: %v\n", err)
		}
		fmt.Println()
	}
}

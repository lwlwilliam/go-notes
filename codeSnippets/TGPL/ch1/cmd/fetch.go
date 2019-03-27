// Fetch 获取对应的 URL，并将其源文本打印出来
// 灵感来源于 curl 工具，当然，curl 提供的功能更为复杂丰富
package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
)

func main()  {
	for _,  url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}
		//_, err = ioutil.ReadAll(resp.Body)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			continue
		}

		// 打印响应体
		fmt.Printf("%s", b)

		// 打印响应头
		for key, value := range resp.Header {
			if strings.ToLower(key) == "set-cookie" {
				fmt.Println()
				fmt.Println("You should set cookie by yourself.")
				fmt.Println()
			}
			fmt.Println(key, ":",  value)
		}
	}
}

// 简单的 web 服务器
package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
)

func main()  {
	// 将所有发送到 / 路径下的请求和 handle 函数关联起来，/ 开头的请求其实就是所有发送到当前站点上的请求
	http.HandleFunc("/", handler)
	// 将所有发送到 /test 路径下的请求和 testhandle 函数关联起来
	http.HandleFunc("/test", testhandler)
	// 服务监听 8000 端口
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

// handler 打印请求 url 的 path
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func testhandler(w http.ResponseWriter, r *http.Request)  {
	for key, val := range r.Header {
		fmt.Fprintf(w, "%s: %v\n", key, val)
	}
}

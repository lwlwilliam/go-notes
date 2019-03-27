// 微型 echo 和计数服务器
// 在 server1 的基础上对请求的次数进行计算；
// 对 url 的请求结果会包含各种 url 被访问的总次数，直接对 /count 这个 url 的访问要除外
//
// 根据请求 url 的不同会调用不同的函数
// 在这些代码的背后，服务器每一次接收请求处理时都会另起一个 goroutine，这样服务器就可以同一时间处理多个请求。
// 然而在并发情况下，假如真的有两个请求同一时刻去更新 count，那么这个值可能并不会被正确地增加。
// 这个程序可能会引发一个严重的 bug：竞态条件。
// 为了避免这个问题，我们必须保证每次修改变量的最多只能有一个 goroutine，这也就是代码里的 mu.Lock() 和 mu.Unlock() 调用将修改 count
// 的所有行为包在中间的目的
package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var mu sync.Mutex
var count int

func main()  {
	// 这里发现有个问题，浏览器请求 /count 的时候似乎 count ++ 都会生效，但是并不会产生 URL.Path 的输出
	// 原因是浏览器请求都会带上 /favicon.icon 这个请求，这个是个性化网站图标
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler 打印请求 url 的 path
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

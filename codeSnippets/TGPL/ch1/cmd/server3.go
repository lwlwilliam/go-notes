// 简单的 web 服务器
package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var count uint
var mu sync.Mutex

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler 打印请求 url 的 path
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count ++
	mu.Unlock()

	fmt.Fprintf(w, "%s, %s, %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}

package http

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string  {
	return fmt.Sprintf("$%.2f", d)
}

type Database map[string]dollars

// version 1
/*
func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
*/

// version 2
// handler 基于 URL 的路径部分(req.URL.path)来决定执行什么逻辑
// 虽然可以继续向 ServeHTTP 方法中添加 case，但在一个实际的应用中，将每个 case 中的逻辑定义到一个分开的方法或
// 函数中会很实用
func (db Database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

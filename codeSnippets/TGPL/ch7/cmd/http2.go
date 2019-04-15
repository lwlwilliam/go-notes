package main

import (
	"net/http"
	http2 "github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch7/http"
)

func main()  {
	db := http2.Database{"shoes": 11, "socks": 3}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.ListenAndServe("localhost:60000", nil)
}
package main

import (
	"log"
	"net/http"
	http2 "github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch7/http"
)

func main()  {
	db := http2.Database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
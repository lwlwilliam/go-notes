package main

import (
	"golang.org/x/net/html"
	"os"
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
	"net/http"
)

func main()  {
	url := "https://github.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v\n", err)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	links.Outline(nil, doc)
}

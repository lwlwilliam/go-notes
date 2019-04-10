package main

import (
	"fmt"
	"os"

	"net/http"

	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://github.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v\n", err)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range links.Visit(nil, doc) {
		if link != "" {
			fmt.Println(link)
		}
	}
}

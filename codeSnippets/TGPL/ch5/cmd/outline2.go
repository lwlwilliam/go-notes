package main

import (
	"net/http"
	"fmt"
	"os"
	"golang.org/x/net/html"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
)

func main()  {
	url := "https://github.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v\n", err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse: %v\n", err)
	}
	resp.Body.Close()

	links.ForEachNode(doc, links.StartElement, links.EndElement)
}
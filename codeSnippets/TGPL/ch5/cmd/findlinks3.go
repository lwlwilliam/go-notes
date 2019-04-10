package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
)

func main()  {
	url := []string{"https://github.com"}
	links.BreadthFirst(links.Crawl, url)
}

package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
	"golang.org/x/net/html"
	"net/http"
	"log"
	"os"
)

func main()  {
	url := "https://github.com"
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("http.Get: %v", err)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("html.Parse: %v", err)
		os.Exit(1)
	}
	fmt.Println(links.SoleTitle(doc))
}

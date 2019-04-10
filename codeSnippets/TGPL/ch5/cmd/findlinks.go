package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
	"fmt"
	"log"
)

func main()  {
	links, err := links.Extract("https://github.com")
	if err != nil {
		log.Println(err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
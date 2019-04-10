package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch5/links"
	"strings"
)

func main()  {
	url := "https://github.com"
	links, _ := links.FindLinks(url)
	for key, link := range links {
		if strings.HasPrefix(link, "https://") {
			fmt.Println(link)
		} else {
			fmt.Printf("%s%s\n", url, link)
		}

		if key == len(links) - 1 {
			fmt.Println(key)
		}
	}
}

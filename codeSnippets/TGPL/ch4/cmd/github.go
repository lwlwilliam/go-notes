package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch4/github"
	"log"
	"fmt"
)

func main()  {
	search := []string{"repo:golang/go is:open json decoder"}
	result, err := github.SearchIssues(search)

	//result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

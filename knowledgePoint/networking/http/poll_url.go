package main
import (
	"fmt"
	"net/http"
)

var urls = []string {
	"https://baidu.com/",
	"http://localhost",
	"https://golang.google.cn/",
}

func main() {
	// Execute an HTTP HEAD request for all urls
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println(url, "connection failed.")
			// fmt.Println("Error", url, err)
		}
		fmt.Print(url, ": ", resp.Status, "\n")
	}
}
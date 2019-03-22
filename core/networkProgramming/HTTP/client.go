package main

import (
	"fmt"
	"net/http"
)

func main()  {
	resp, _ := http.Get("http://localhost:8080")

	for headerName, headerValue := range resp.Header {
		fmt.Println(headerName, ":", headerValue[0])
	}
}

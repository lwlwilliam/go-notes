package main

import (
	"strings"
	"net/http"
	"log"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("a=b&c=d")
	resp, err := http.Post("http://example.com", "application/x-www-form-urlencoded", r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println(err)
				fmt.Printf("%s", buf[:n])
				break
			}

			log.Fatal(err)
		}
		fmt.Printf("%s", buf[:n])
	}
}
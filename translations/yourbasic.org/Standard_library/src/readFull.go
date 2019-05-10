package main

import (
	"strings"
	"io"
	"log"
	"fmt"
)

func main() {
	r := strings.NewReader("abcde")

	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", buf)

	if _, err := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}
}

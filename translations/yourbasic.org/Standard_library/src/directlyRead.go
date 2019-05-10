package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("abcde")

	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Printf("%s", buf[:n])
		if err == io.EOF {
			break
		}
	}
}

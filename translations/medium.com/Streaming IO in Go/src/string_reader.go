package main

import (
	"strings"
	"io"
	"fmt"
)

func main() {
	reader := strings.NewReader("Clear is better than clever")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}

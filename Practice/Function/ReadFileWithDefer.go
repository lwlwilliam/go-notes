package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fname := "README.md"
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil {
		os.Exit(1)
	}
	bReader := bufio.NewReader(f)
	for {
		line, ok := bReader.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Println(strings.Trim(line, "\r\n"))
	}
}

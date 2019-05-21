package main

import (
	"os"
	"log"
	"io"
	"fmt"
)

func main() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := make([]byte, 4)
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
}

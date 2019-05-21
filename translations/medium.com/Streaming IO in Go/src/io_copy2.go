package main

import (
	"os"
	"log"
	"io"
)

func main() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	}
}

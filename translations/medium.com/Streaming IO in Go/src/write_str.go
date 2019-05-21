package main

import (
	"os"
	"log"
	"io"
)

func main() {
	file, err := os.Create("./magic_msg.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		log.Fatal(err)
	}
}

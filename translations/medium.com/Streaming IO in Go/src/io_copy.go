package main

import (
	"bytes"
	"os"
	"log"
	"io"
	"fmt"
)

func main() {
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	file, err := os.Create("./proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := io.Copy(file, proverbs); err != nil {
		log.Fatal(err)
	}
	fmt.Println("file created")
}

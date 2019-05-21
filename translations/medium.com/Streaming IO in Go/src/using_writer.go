package main

import (
	"bytes"
	"log"
	"fmt"
)

func main() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))
		if err != nil {
			log.Fatal(err)
		}

		if n != len(p) {
			log.Fatal("failed to write data")
		}
	}

	fmt.Println(writer.String())
}

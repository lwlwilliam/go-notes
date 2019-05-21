package main

import (
	"os"
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

	file, err := os.Create("./proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, p := range proverbs {
		n, err := file.Write([]byte(p))
		if err != nil {
			log.Fatal(err)
		}

		if n != len(p) {
			log.Fatal(err)
		}
	}
	fmt.Println("file write done")
}

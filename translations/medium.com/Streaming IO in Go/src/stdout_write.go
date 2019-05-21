package main

import (
	"os"
	"log"
)

func main() {
	proverbs := []string {
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	for _, p := range proverbs {
		n, err := os.Stdout.Write([]byte(p))
		if err != nil {
			log.Fatal(err)
		}

		if n != len(p) {
			log.Fatal("failed to write data")
		}
	}
}

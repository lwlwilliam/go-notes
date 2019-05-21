package main

import (
	"os"
	"log"
	"bufio"
	"io"
	"fmt"
)

func main() {
	file, err := os.Open("./planets.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Print(line)
	}
}

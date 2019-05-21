package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	bytes, err := ioutil.ReadFile("./planets.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bytes)
}

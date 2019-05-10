package main

import (
	"strings"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	r := strings.NewReader("abcde")

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", buf)
}

package main

import (
	"fmt"
	"time"
	. "math/rand"
)

func main() {
	Seed(time.Now().UnixNano())
	fmt.Println(Intn(100))
}

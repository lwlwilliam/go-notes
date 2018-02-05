package main

import (
	"fmt"
)

func main() {
	var dog_age = 10

	if dog_age > 10 {
		fmt.Println("A git dog")
	} else if dog_age > 1 && dog_age <= 10 {
		fmt.Println("A small dog")
	} else {
		fmt.Println("A baby dog")
	}
}

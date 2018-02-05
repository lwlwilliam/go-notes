package main

import (
	"fmt"
)

func main() {
	var x = map[string]string{
		"A": "Apple",
		"B": "Banana",
		"O": "Orange",
		"P": "Pear",
	}

	for key, val := range x {
		fmt.Println("Key: ", key, "Val: ", val)
	}

	var y map[string]string
	y = make(map[string]string)

	y["A"] = "Apple"
	y["B"] = "Banana"
	y["O"] = "Orange"
	y["P"] = "Pear"

	for key, val := range y {
		fmt.Println("Key: ", key, "Val: ", val)
	}

	z := make(map[string]string)
	z["A"] = "Apple"
	z["B"] = "Banana"
	z["O"] = "Orange"
	z["P"] = "Pear"

	for key, val := range z {
		fmt.Println("Key: ", key, "Val: ", val)
	}

	fmt.Println("The length of z before delete: ", len(z))
	delete(z, "A")
	fmt.Println("The length of z after delete: ", len(z))
}

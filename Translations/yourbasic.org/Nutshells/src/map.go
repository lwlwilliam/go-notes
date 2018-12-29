package main

import "fmt"

func main() {
	m := make(map[string]float64)

	m["pi"] = 3.1416	// Add a new key-value pair.
	fmt.Println(m)		// map[pi:3.1416]

	v := m["pi"]	// v == 3.1416
	fmt.Println(v)
	v = m["pie"]	// v == 0 (zero value)
	fmt.Println(v)

	_, found := m["pi"]	// found == true
	fmt.Println(found)
	_, found = m["pie"]	// found == false
	fmt.Println(found)

	if x, found := m["pi"]; found {
		fmt.Println(x)
	}

	delete(m, "pi")		// Delete a key-value pair.
	fmt.Println(m)			// map[]
	fmt.Println(len(m))
}

package main

import "fmt"

type IntSet struct { /* ... */ }

func (*IntSet) String() string {
	return "string"
}

func main()  {
	//var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
	var i IntSet
	var _ = i.String()

	//var _ fmt.Stringer = i // compile error: IntSet lacks String method
	var _ fmt.Stringer = &i
}

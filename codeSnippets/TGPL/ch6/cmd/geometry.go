package main

import (
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch6/geometry"
	"fmt"
)

func main()  {
	p := geometry.Point{3, 3}
	q := geometry.Point{4, 4}
	r := geometry.Point{5, 5}
	s := geometry.Point{6, 6}
	t := geometry.Point{7, 7}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	path := geometry.Path{p, q, r, s, t}
	fmt.Println(path.Distance())
}

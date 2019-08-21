package main

import (
	stdRand "math/rand"
	"fmt"
	"time"
	cusRand "github.com/lwlwilliam/go/translations/yourbasic.org/Nutshells/src/math/rand"
)

func main() {
	m := cusRand.Int()
	stdRand.Seed(time.Now().UnixNano())
	n := stdRand.Intn(100)

	fmt.Println(m, n)
}

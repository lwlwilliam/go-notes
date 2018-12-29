package main

import (
	cusRand "github.com/lwlwilliam/Golang/Translations/yourbasic.org/Nutshells/src/math/rand"
	stdRand "math/rand"
	"fmt"
	"time"
)

func main() {
	m := cusRand.Int()
	stdRand.Seed(time.Now().UnixNano())
	n := stdRand.Intn(100)

	fmt.Println(m, n)
}

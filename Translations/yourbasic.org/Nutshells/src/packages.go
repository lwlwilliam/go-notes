package main

import (
	"fmt"
	"math/rand"
	"github.com/lwlwilliam/Golang/Translations/yourbasic.org/Object-oriented_programming/src/timer"
	"time"
)

func main() {
	n := rand.Intn(100)
	g := timer.StopWatch{}
	g.Start()
	time.Sleep(3 * time.Second)
	t := g.Total()
	fmt.Println(t, n)
}

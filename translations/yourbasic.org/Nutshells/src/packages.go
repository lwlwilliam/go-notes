package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/lwlwilliam/go-notes/translations/yourbasic.org/Object-oriented_programming/src/timer"
)

func main() {
	n := rand.Intn(100)
	g := timer.StopWatch{}
	g.Start()
	time.Sleep(3 * time.Second)
	t := g.Total()
	fmt.Println(t, n)
}

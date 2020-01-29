package main

import (
	"fmt"
	"time"
	"github.com/lwlwilliam/go-notes/translations/yourbasic.org/Object-oriented_programming/src/timer"
)

func main() {
	clock := new(timer.StopWatch)
	clock.Start()
	time.Sleep(time.Second)

	total := clock.Total()
	fmt.Println(total)
	/*
		if clock.running {	// 非法
			// ...
		}
	*/
}

package main

import (
	"fmt"
	"github.com/lwlwilliam/Golang/Translations/yourbasic.org/Object-oriented_programming/src/timer" //	"./timer"	// 不建议使用相对路径导入 package
	"time"
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

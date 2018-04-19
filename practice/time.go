package main

import (
	"fmt"
	"time"
)

func main()  {
	test := time.Now()
	fmt.Println(test.Format("2006-01-02 03:04:05 PM"))


	now, _:= time.Parse("2006-01-02 15:04:05", "2018-03-17 18:55:06")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-04-04 14:04:04", time.Local)
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"), t.Format("2006-01-02 03:04:05 PM"))
}

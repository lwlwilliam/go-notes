/*
Go 支持通过基于描述模板的时间格式化和解析
 */
package main

import (
	"fmt"
)
import "time"

func main() {
	p := fmt.Println

	// 按照 RFC3339 进行格式化的例子，使用对应模式常量
	t := time.Now()
	p("1.", t.Format(time.RFC3339))

	// 时间解析使用同 Format 相同的形式值
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	p("2.", t1)


	// Format 和 Parse 使用基于例子的形式来决定日期格式，一般只要使用 time 包中提供的模式常量就行了，
	// 但是也可以实现自定义模式。模式必须使用时间 Mon Jan 2 15:04:05 MST 2006 来指定给定时间/字符串的格式化/解析方式。
	// 时间一定要按照如下所示：2006 为年，15 为小时，Monday 代表星期几等
	p("*.", t.Format("2006-03-30 15:04:05"))
	p("3.", t.Format("3:04PM"))
	p("4.", t.Format("Mon Jan _2 15:04:05 2006"))
	p("5.", t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p("6.", t2)

	// 对于纯数字表示的时间，也可以使用标准的格式化字符串来表示时间值的组成
	fmt.Printf("7. %d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse 函数在输入的时间格式不正确时会返回一个错误
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p("8.", e)
}

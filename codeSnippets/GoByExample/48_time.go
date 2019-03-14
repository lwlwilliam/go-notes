/*
Go 对时间和时间段提供了大量的支持
 */
package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 当前时间
	now := time.Now()
	p("1.", now)
	//p(now.Year())

	// 通过提供年月日等信息，可以构建一个 time。时间总是关联着位置信息，例如时区
	then := time.Date(
		2009, 11, 17, 20, 34, 59, 651387237, time.UTC,
	)

	p("2.", then)

	// 提供出时间的各个组成部分
	p("3.", then.Year())
	p("4.", then.Month())
	p("5.", then.Day())
	p("6.", then.Hour())
	p("7.", then.Minute())
	p("8.", then.Second())
	p("9.", then.Nanosecond())
	p("10.", then.Location())

	// 输出星期一到星期日
	p("11.", then.Weekday())

	// 比较两个时间，判断时间先后，精确到秒
	p("12.", then.Before(now))
	p("13.", then.After(now))
	p("14.", then.Equal(now))

	// 方法 Sub 返回一个 Duration 来表示两个时间点的间隔时间
	diff := now.Sub(then)
	p("15.", diff)

	// 计算出不同单位下的时间长度值
	p("16.", diff.Hours())
	p("17.", diff.Minutes())
	p("18.", diff.Seconds())
	p("19.", diff.Nanoseconds())

	// 使用 Add 将时间后移一个时间间隔，或者使用一个 - 符号来将时间前移一个时间间隔
	p("20.", then.Add(diff))
	p("21.", then.Add(-diff))
}

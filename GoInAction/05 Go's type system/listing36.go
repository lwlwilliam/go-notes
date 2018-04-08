// 这个示例程序展示了 Go 语言里如何使用接口
package main

import (
	"fmt"
)

// notifier 是一个定义了通过类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email,
	)
}

// main 是应用程序的入口
func main() {
	// 创建一个 user 类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}

    // 注意，这里不能将 u（类型是 user）作为参数，因为 user 类型并没有实现 notifier
//	sendNotification(&u)
	u.notify()
}

// sendNotification 接受一个实现了 notifier 接口的值并发送通知
/*
func sendNotification(n notifier) {
	n.notify()
}
*/

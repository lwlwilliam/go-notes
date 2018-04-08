// 这个示例程序展示当内部类型和外部类型要实现同一个接口时的做法
package main

import "fmt"

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序时定义一个用户类型
type user struct {
	name string
	email string
}

// 通过 user 类型值的指针调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", 
		u.name, 
		u.email, 
	)
}

// admin 代表一个拥有权限的管理员用户
type admin struct {
	user
	level string
}

// 通过 admin 类型值的指针调用的方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", 
		a.name, 
		a.email, 
	)
}

// main 是应用程序的入口
func main() {
	// 创建一个 admin 用户
	ad := admin{
		user: user{
			name: "john smith", 
			email: "john@yahoo.com", 
		}, 
		level: "super", 
	}

	// 给 admin 用户发送一个通知
	// 接口的嵌入的内部类型实现并没有提升到外部类型
	sendNotification(&ad)

	// 可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法没有被提升
	ad.notify()
}

// sendNotification 接受一个实现了 nofifier 接口的值并发送通知
func sendNotification(n notifier) {
	n.notify()
}

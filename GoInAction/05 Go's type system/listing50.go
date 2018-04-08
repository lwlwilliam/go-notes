// 这个示例程序展示如何将一个类型嵌入另一个类型，以及内部类型和外部类型之间的关系
package main

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name string
	email string
}

// notify 实现了一个可以通过 user 类型值的指针调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", 
		u.name, 
		u.email, 
	)
}

// admin 代表一个拥有权限的管理员用户
type admin struct {
	user  // 嵌入类型
	level string
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

	// 可以直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法也被提升到外部类型
	ad.notify()
}

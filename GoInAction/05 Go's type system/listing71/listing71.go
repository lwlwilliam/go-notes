// 这个示例程序展示公开的结构类型中未公开的字段无法直接访问
package main

import (
	"fmt"
	"./entities"
)

// main 是应用程序的入口 

func main() {
	// 创建 entities 包中的 User 类型的值
	u := entities.User{
		Name: "Bill", 
		email: "bill@email.com",  // 结构字面量中结构 entities.User 的字段 email 未知，因此编译错误
	}

	fmt.Printf("User: %v\n", u)
}

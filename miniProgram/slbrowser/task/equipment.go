// 装备
package task
import (
	//"../tool"
)

// 装备类型
type Equipment struct {
	// 级别：白色为 1，蓝色为 2，绿色为 3，紫色为 4，金色为 5
	Level	int
	// 等级：越高级越好
	Grade	int
	// 主要是耐久度，格式为"100/100"，前面的整数为当前耐久度，后面的为当前最大耐久度
	Info	string
}


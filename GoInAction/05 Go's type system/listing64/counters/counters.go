// counters 包提供告警计数器的功能
package counters

// alertCounter 是一个未公开的类型
// 这个类型用于保存告警计数
type alertCounter int
type AlertCounter int

// New 创建并返回一个未公开的
// alertCounter 类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}

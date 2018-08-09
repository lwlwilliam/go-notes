// 惰性生成器
// 生成器是指当被调用时返回一个序列中下一个值的函数，例如：
//		generateInteger() => 0
//		generateInteger() => 1
//		generateInteger() => 2
//		...
// 生成器每次返回的是序列中下一个值而非整个序列；这种特性也称之为惰性求值：只在需要时进行求值，同时保留相关变量资源（内存和 CPU）：这是一项在需
// 要时对表达式进行求值的技术。例如，生成一个无限数量的偶数序列：要产生这样一个序列并且再一个一个地使用可能会很困难，而且内存会溢出！但是一个含
// 有通道和 go 协程的函数能轻易实现这个需求。
package main

import (
	"fmt"
)

type Any interface{}
type EvalFunc func(Any) (Any, Any)

func main() {
	// 公差为 2 的数列
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	// 以初始值为 0，返回一个可获取数列的函数
	even := BuildLazyIntEvaluator(evenFunc, 0)

	// 获取数列的前二十项
	for i := 0; i < 20; i ++ {
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

// 惰性生产器的工厂函数（这里应该放在一个工具包中实现）
// 这里只负责使用传入的函数对初始状态参数进行循环处理，处理完返回可获取处理结果的函数
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <- retValChan
	}
	go loopFunc()
	return retFunc
}

// 整数惰性产生器
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

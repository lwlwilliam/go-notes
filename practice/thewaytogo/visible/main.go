package main

import (
	"fmt"
	"./pkg"  // 不建议这么写
)

func main() {
	var stru = pkg.Stru {A:"haha"}  //, b:"yes"}  b 不可访问
	var t pkg.T = 3

	fmt.Println(pkg.S, stru)
	t.Haha()
	// t.test()  test 方法不可访问
}

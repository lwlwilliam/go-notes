/*
Go 的 math/rand 包提供了伪随机数生成器
 */
package main

import "time"
import "fmt"
import "math/rand"

func main() {
	// rand.Intn 返回一个随机的整数 n，0 ≤ n ≤ 100
	fmt.Println("first:")
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println()

	// rand.Float64 返回一个 64 位浮点数 f，0.0 ≤ f ≤ 1.0
	fmt.Println("second:")
	fmt.Println(rand.Float64())
	fmt.Println()
	fmt.Println()

	// 生成 5.0 ≤ f ≤ 10.0
	fmt.Println("third:")
	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	fmt.Println()

	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。要产生变化的序列，需要给定一个变化的种子。
	// 需要注意的是，如果出于加密目的，需要使用随机数的放，请使用 crypto/rand 包，此方法不够安全
	fmt.Println("fourth:")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	fmt.Println()

	// 如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列
	fmt.Println("fifth:")
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	fmt.Println()


	fmt.Println("sixth:")
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

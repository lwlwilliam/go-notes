/*
SHA1 散列经常用于生成二进制文件或者文本块的短标识。例如，git 版本控制系统大量使用 SHA1 来标识受版本控制的文件和目录。
这里是 Go 中如何进行 SHA1 散列计算的例子。

Go 在多个 crypto/* 包中实现了一系列散列函数

可以使用与以下相似的方式来计算其他形式的散列值。如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。

注意，如果需要密码学上的安全散列，需要小心的研究一下哈希强度
 */
package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	// 产生一个散列值的方式是 sha1.New()
	// sha1.Write(bytes)，然后 sha1.Sum([]byte{}).
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串，需要使用 []byte(s) 来强制转换成字节数组
	h.Write([]byte(s))

	// 这个用来得到最终的散列值的字符切片。Sum 的参数可以用来在现有的字符切片追加额外的字节切片：一般不需要
	bs := h.Sum(nil)

	fmt.Println(s)
	// 散列值以可读 16 进制格式输出
	fmt.Printf("%x\n", bs)
}
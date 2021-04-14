// 真正的 Go 字符串变量 Name 对应的大小是 16 字节，Name 变量并没有直接对应"gopher"字符串，而是对应 16 字节的 reflect.StringHeader 结构体：
// type reflect.StringHeader struct {
//     Data uintptr
//     Len  int
// }
// 从汇编角度看，Name前8字节对应底层真实字符串数据的指针，后8字节对应底层真实字符串的有效长度，这里是6字节。
// GLOBL ·NameData(SB),$8
DATA ·NameData(SB)/8,$"gopher"
GLOBL ·Name(SB),$16
DATA ·Name+0(SB)/8,$·NameData(SB)
DATA ·Name+8(SB)/8,$6
// 直到这里，以上汇编运行时会产生以下错误：
// pkg.NameData: missing Go type information for global symbol: size 8
// 意思就是 NameData 符号没有类型信息。其实 Go 汇编语言中定义的数据并没有所谓的类型，每个符号不过是对应一块内存而已，因此 NameData 也是没有类型的。但是 Go 语言是自带垃圾回收器的语言，而 Go 汇编在自动垃圾回收体系框架内。当 Go 语言的垃圾回收器在扫描到 NameData 变量时，无法知晓该变量内部是否包含指针，因此就出现了错误。错误的根本原因不是 NameData 没有类型，而是 NameData 没有标注是否含有指针信息。
// 将第7行注释，用第16、17行替代
#include "textflag.h"
GLOBL ·NameData(SB),NOPTR,$8

// 或者用 pkgStr2_amd64.s 的方法

// 当然也可以在 pkgStr.go 文件中直接声明 var NameData [8]byte

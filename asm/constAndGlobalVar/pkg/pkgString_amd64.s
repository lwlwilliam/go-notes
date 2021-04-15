// 从 Go 汇编语言角度看，字符串只是一种结构体，字符串头的结构体定义如下：
// type reflect.StringHeader struct {
//     Data uintptr
//     Len int
// }
GLOBL ·helloworld(SB),$16
#include "textflag.h"
GLOBL text(SB),NOPTR,$16

DATA text+0(SB)/8,$"Hello wo"
DATA text+8(SB)/8,$"rld\n"
// 注意，不要将 text 和 长度的顺序弄乱了
DATA ·helloworld+0(SB)/8,$text(SB)
DATA ·helloworld+8(SB)/8,$12

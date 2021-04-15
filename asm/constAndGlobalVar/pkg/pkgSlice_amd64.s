// 切片类型变量和字符串变量相似，只不过对应的是切片头结构体而已。切片头的结构体如下：
// type reflect.SliceHeader struct {
//     Data uintptr
//     Len int
//     Cap int
// }
GLOBL ·helloworldslice(SB),$24
#include "textflag.h"
GLOBL textslice(SB),NOPTR,$16

DATA textslice+0(SB)/16,$"Hello world"
DATA ·helloworldslice+0(SB)/8,$textslice(SB)
DATA ·helloworldslice+8(SB)/8,$11
DATA ·helloworldslice+16(SB)/8,$20

// Name2 内存为24字节，多出的8字节存放底层的"gopher2"字符串
// ·Name2 符号前16字节依然对应 reflect.StringHeader 结构体
// Data 部分对应 $·Name+16(SB)，表示数据的地址为Name符号往后偏移16字节位置；Len依然对应6字节
GLOBL ·Name2(SB),$24
DATA ·Name2+0(SB)/8,$·Name2+16(SB)
DATA ·Name2+8(SB)/8,$7
DATA ·Name2+16(SB)/8,$"gopher2"

// map/channel 等类型并没有公开的内部结构，它们只是一种未知类型的指针，无法直接初始化。
// 在汇编代码中只能为类似变量定义并进行零值初始化
GLOBL ·m(SB),$8 // var m map[string]int
DATA ·m+0(SB)/8,$0
GLOBL ·ch(SB),$8
DATA ·ch+0(SB)/8,$0

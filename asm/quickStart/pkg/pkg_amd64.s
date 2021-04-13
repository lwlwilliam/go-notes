// GLOBL 命令用于将符号导出，以供其他代码引用
GLOBL ·Id(SB),$8
// DATA 命令用于初始化包变量，语法：DATA symbol+offset(SB)/width, value
// symbol 为变量在汇编语言中对应的标识符，offset 是符号开始地址的偏移量，width 是要初始化内存的宽度大小，value 是要初始化的值。
// 其中当前包中 Go 语言定义的符号 symbol，在汇编代码中对应 ·symbol，其中点符号"·"为一个特殊的 unicode 符号
DATA ·Id+0(SB)/1,$0x37
DATA ·Id+1(SB)/1,$0x25
DATA ·Id+2(SB)/1,$0x00
DATA ·Id+3(SB)/1,$0x00
DATA ·Id+4(SB)/1,$0x00
DATA ·Id+5(SB)/1,$0x00
DATA ·Id+6(SB)/1,$0x00
DATA ·Id+7(SB)/1,$0x00

// TEXT ·main(SB),$16-0 用于定义 main() 函数，其中 $16-0 表示 main() 函数的帧大小是16字节（对应 string 头部结构体的大小，用于给 runtime.printstring() 函数传递参数），0 表示 main() 函数没有参数和返回值。main() 函数内部通过调用运行时内部的 runtime.printstring(SB)函数来打印字符串，然后调用 runtime.printnl() 打印换行符号。
TEXT ·main(SB),$16-0
    MOVQ ·helloworld+0(SB),AX; MOVQ AX,0(SP)
    MOVQ ·helloworld+8(SB),BX; MOVQ BX,8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET

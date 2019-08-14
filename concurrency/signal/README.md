### 信号

操作系统信号(signal)是 IPC 中唯一一种异步的通信方法，它的本质是用软件来模拟硬件的中断机制。信号用来通知某个进程有某个事件发生了。例如，在命令行终端按下某些快捷键，就会挂起或停止正在运行的程序。另外，通过 kill 命令杀死某个进程的操作也有信号的参与。

每一个信号都有一个以`SIG`为前缀的名字，例如 SIGINT、SIGQUIT 以及 SIGKILL 等。但是，在操作系统内部，这些信号都由正整数表示，这些正整数称为信号编号。在 Linux 或 Unix 的命令行终端下，可以使用`kill -l`命令来查看当前系统所支持的信号，如下所示。

```
 1) SIGHUP       2) SIGINT       3) SIGQUIT      4) SIGILL
 5) SIGTRAP      6) SIGABRT      7) SIGEMT       8) SIGFPE
 9) SIGKILL     10) SIGBUS      11) SIGSEGV     12) SIGSYS
13) SIGPIPE     14) SIGALRM     15) SIGTERM     16) SIGURG
17) SIGSTOP     18) SIGTSTP     19) SIGCONT     20) SIGCHLD
21) SIGTTIN     22) SIGTTOU     23) SIGIO       24) SIGXCPU
25) SIGXFSZ     26) SIGVTALRM   27) SIGPROF     28) SIGWINCH
29) SIGINFO     30) SIGUSR1     31) SIGUSR2
```

`kill -s signal_name pid`命令可以给 pid 所属进程发送信号。signal_name 既可以是数字，也可以是字符串，例如要发送中断信号，可以`kill -s 2 pid`，也可以`kill -s SIGINT
 pid`。
 
 在类 Unix 操作系统下有两种信号既不能自行处理，也不会被忽略，它们是 SIGKILL 和 SIGSTOP，对它们的响应只能是执行系统的默认操作。这种策略的最根本原因是：它们向系统的超级用户提供了使进程终止或停止的可靠方法。这种保障不论对于应用程序还是操作系统来说，都是非常有必要的。
 
 对于其他信号，除了能够自行处理它们之外，还可以在之后的任意时刻恢复它们的系统默认操作。

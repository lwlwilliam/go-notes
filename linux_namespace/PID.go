// PID Namespace 用于隔离进程 ID。同一个进程在不同的 PID Namespace 里可以拥有不同的 PID。这样就可以理解，在 docker container 里面，使用 ps -ef 经常会发现，在容器内，前台运行的那个进程 PID 是1，但是在容器外，使用 ps -ef 会发现同样的进程却有不同的 PID，这就是 PID Namespace 做的事情。
// 测试：
//  1. go run PID.go
//  2. 另开一个 shell
//  3. 在另一个 shell 查看 ps ajx
//  4. 在 go run PID.go 中创建的 sh 中执行 echo $$ 查看当前 PID，跟另一个 shell 中查看到的 PID 进行对比，可以发现 PID 不一样的
// 注意：不能在两个 shell 通过 ps ajx 来对比，因为 ps/top 等命令会使用 /proc 中的内容。
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
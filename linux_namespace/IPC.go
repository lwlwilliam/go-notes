// IPC Namespace 用来隔离 System V IPC 和 POSIX message queue。每一个 IPC Namespace 都有自己的 System V IPC 和 POSIX message queue。
// 测试方法：
//  1. ipcs 命令查看当前 ipc 信息
//  2. ipcmk -Q 创建 message queue
//  3. ipcs 查看最新 ipc 信息，现在应该可以看到一个 queue 了
//  4. go run IPC.go
//  5. 在新的 sh 中运行 ipcs 查看 ipc 信息，应该是查不到的，说明 ipc 被隔离了
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

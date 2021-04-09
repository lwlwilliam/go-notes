// Mount Namespace 用来隔离各个进程看到的挂载点视图。在不同 Namespace 的进程中，看到的文件系统层次是不一样的。在 Mount Namespace 中调用 mount() 和 umount() 仅仅只会影响到当前 Namespace，对全局的文件系统是没有影响的
// 1. go run Mount.go
// 2. ps ajx
// 3. ls /proc
// 4. mount -t proc proc /proc
// 5. ps ajx # 对比上面的结果
// 6. ls /proc # 对比上面的结果
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

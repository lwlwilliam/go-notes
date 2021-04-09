// UTS Namespace 主要用来隔离 nodename 和 domainname 两个系统标识。在 UTS Namespace 中，每个 Namespace 允许有自己的 hostname。
// 测试方法：
// 	1. go run UTS.go
// 	2. hostname 查看原始 hostname
//  3. hostname -b hello 修改 hostname
//  4. hostname 查看新 hostname
//  5. exit 退出
//  6. hostname 查看 hostname 是否已改变
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
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

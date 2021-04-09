// Network Namespace 用于隔离网络设备、IP 地址端口等网络栈的 Namespace。Network Namespace 可以让每个容器拥有自己独立的（虚拟的）网络设备，而且容器内的应用可以绑定到自己的端口，每个 Namespace 内的端口都不会互相冲突，在宿主机上搭建网桥后，就能很方便地实现容器之间的通信，而且不同容器上的应用可以使用相同的端口。
//
// 分别在宿主机和程序中运行 ifconfig 命令
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
	//}
	//// 设置用户和用户组为 1
	//cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(1), Gid: uint32(1)}

	// Linux 内核3.19以上修改了，需要用以下代替上面的代码
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 100,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 100,
				HostID:      0,
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(-1)
}

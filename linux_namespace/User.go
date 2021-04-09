// User Namespace 主要是隔离用户的用户组ID。也就是说，一个进程的 User ID 和 Group ID 在 User Namespace 内外可以是不同的。比较常用的是，在宿主机上以一个非 root 用户运行创建一个 User Namespace，然后在 User Namespace 里却映射成 root 用户。这意味着，这个进程在 User Namespace 中有 root 权限，但是在外面却没有 root 的权限。
//
// 分别在宿主机和程序中运行 id 命令，就可以看到是不同的用户
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
	//	Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
	//}
	//// 设置用户和用户组为 1
	//cmd.SysProcAttr.Credential = &syscall.Credential{Uid: uint32(1), Gid: uint32(1)}

	// Linux 内核3.19以上修改了，需要用以下代替上面的代码
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
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

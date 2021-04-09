// Linux Cgroups(Control Groups)提供了对一组进程及将来子进程的资源限制、控制和统计的能力，这些资源包括CPU、内存、存储、网络等。通过Cgroups，可以方便地限制某个进程的资源占用，并且可以实时地监控进程的监控和统计信息。
//
// cgroup 是对进程分组管理的一种机制，一个 cgroup 包含一组进程，并可以在这个 cgroup 上增加 Linux subsystem 的各种参数配置，将一组进程和一组subsystem的系统参数关联起来。
// subsystem 是一组资源控制的模块。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		fmt.Printf("%v", cmd.Process.Pid)

		os.Mkdir(path.Join(cgroupMemHierarchyMount, "testmemorylimit"), 0755)
		ioutil.WriteFile(path.Join(cgroupMemHierarchyMount, "testmemorylimit", "task"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait()
}

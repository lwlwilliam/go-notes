// 模拟 ps aux | grep redis 命令
package main

import (
	"os/exec"
	"bytes"
	"fmt"
)

func main() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "redis")

	var outputBuf1 bytes.Buffer
	// 输出到 buf1 中
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup %s\n", err)
		return
	}
	// 等待 cmd1 结束
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
		return
	}

	// 设置输入为 buf1
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	// 设置输出为 buf2
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The second command can not be startup: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Coundn't wait for the second command: %s\n", err)
		return
	}

	fmt.Printf("%s\n", outputBuf2.Bytes())
}
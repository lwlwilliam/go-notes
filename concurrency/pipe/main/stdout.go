package main

import (
	"os/exec"
	"fmt"
	"bufio"
	"bytes"
	"io"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")

	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Could't obtain the stdout pipe for command NO.0: %s\n", err)
		return
	}

	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command NO.0 can not be startup: %s\n", err)
		return
	}

	t := 1
	switch t {
	case 1:
		bufferReader(stdout0)
	default:
		buffer(stdout0)
	}
}

func buffer(stdout0 io.ReadCloser)  {
	var outputBuf0 bytes.Buffer
	for {
		tempOutput := make([]byte, 3)
		n, err := stdout0.Read(tempOutput)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
				return
			}
		}
		if n > 0 {
			outputBuf0.Write(tempOutput[:n])
		}
	}
	fmt.Printf("Output: %s\n", outputBuf0.String())
	fmt.Println(len(outputBuf0.String()))
}

func bufferReader(stdout0 io.ReadCloser) {
	outputBuf0 := bufio.NewReader(stdout0)
	output0, _, err := outputBuf0.ReadLine()
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
		return
	}
	fmt.Printf("%s\n", string(output0))
}
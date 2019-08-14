package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"sync"
)

func main() {
	sigRecv1 := make(chan os.Signal, 1)
	sigRecv2 := make(chan os.Signal, 1)
	// 希望自行处理的信号，当 sigs 为空切片时，相当于要自行处理所有信号
	sigs := []os.Signal{syscall.SIGINT, syscall.SIGABRT}
	sigs = []os.Signal{}
	// 当操作系统向当前进程发送指定信号时发出通知
	// 会向两个 channel 发送信号
	signal.Notify(sigRecv1, sigs...)
	signal.Notify(sigRecv2, sigs...)

	var wg sync.WaitGroup
	wg.Add(2)

	// 实际场景中，这样做比较危险，因为这相当于忽略了当前进程本该处理的信号
	go func() {
		count := 0
		for sig := range sigRecv1 {
			fmt.Printf("Received a signal from sigRecv1: %s\n", sig)
			count ++

			if count > 2 {
				signal.Stop(sigRecv1)
			}
		}

		wg.Done()
	}()

	go func() {
		for sig := range sigRecv2 {
			fmt.Printf("Received a signal from sigRecv2: %s\n", sig)
		}

		wg.Done()
	}()

	wg.Wait()
}

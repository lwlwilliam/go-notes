package main

import (
	"sync"
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

func main()  {
	// 要查询的目录
	roots := []string{"/Applications/MAMP/htdocs/notes"}
	// 开始并发查询目录
	var w sync.WaitGroup
	var fileSize = make(chan int64)
	var sema = make(chan struct{}, 20)
	for _, root := range roots {
		w.Add(1)
		go walkDir(root, &w, fileSize, sema)
	}

	// 等待目录查询结束，获取文件数和大小
	go func() {
		w.Wait()
		close(fileSize)
	}()

	// 打印目录所有文件总大小
	var nfiles, nbytes int64
Loop:
	for {
		select {
		case size, ok := <- fileSize:
			if !ok {
				break Loop
			}
			nfiles ++
			nbytes += size
		default:
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, w *sync.WaitGroup, fileSize chan <- int64, sema chan struct{})  {
	defer w.Done()
	for _, entry := range dirents(dir, sema) {
		if entry.IsDir() {
			w.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, w, fileSize, sema)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string, sema chan struct{}) []os.FileInfo  {
	sema <- struct{}{}
	defer func() {
		<- sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// ...determine roots.
	roots := []string{"/Applications/MAMP/htdocs/notes"}
	// Traverse each root of the file tree in parallel.
	fileSize := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSize)
	}

	go func() {
		n.Wait()
		close(fileSize)
	}()

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

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
// 返回目录项
func dirents(dir string) []os.FileInfo {
	// ioutil.ReadDir 函数会返回一个 os.FileInfo 类型的 slice，os.FileInfo 类型也是
	// os.Stat 这个函数的返回值。
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

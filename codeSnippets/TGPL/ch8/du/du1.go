package du

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func WalkDir(dir string, fileSizes chan <- int64)  {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			WalkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
// 返回目录项
func dirents(dir string) []os.FileInfo  {
	// ioutil.ReadDir 函数会返回一个 os.FileInfo 类型的 slice，os.FileInfo 类型也是
	// os.Stat 这个函数的返回值。
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

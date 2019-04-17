package main

import (
	"flag"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch8/du"
	"fmt"
	"os"
)

func main()  {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		fmt.Println(os.Getwd())
		roots = []string{"."}
	}

	roots = []string{"/Applications/MAMP/htdocs/notes"}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			du.WalkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles ++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
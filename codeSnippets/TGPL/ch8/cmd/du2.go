package main

import (
	"flag"
	"time"
	"fmt"
	"os"
	"github.com/lwlwilliam/Golang/codeSnippets/TGPL/ch8/du"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main()  {
	// ...start background goroutine...
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		wd, _ := os.Getwd()
		fmt.Println(wd)
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

	// Print the results periodically.
	var ticker = new(time.Ticker)
	if *verbose {
		ticker = time.NewTicker(10 * time.Millisecond)
	}
	var nfiles, nbytes int64
Loop:
	for {
		select {
		case size, ok := <- fileSizes:
			if !ok {
				break Loop
			}
			nfiles ++
			nbytes += size
		case <- ticker.C:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64)  {
	fmt.Printf("%d files %.1f KB\n", nfiles, float64(nbytes)/1e3)
}

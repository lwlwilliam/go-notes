// FIFO 又名命名管道
package main

import (
	"os"
	"log"
	"fmt"
	"sync"
)

func main() {
	rd, wt, err := os.Pipe()
	if err != nil {
		log.Fatalln(err)
	}


	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		n, err := wt.Write([]byte("Hello world!"))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Written %d bytes. [file-based pipe]\n", n)
		wg.Done()
	}()


	output := make([]byte, 100)
	n, err := rd.Read(output)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Read %d bytes. [file-based pipe]\n", n)

	wg.Wait()
}

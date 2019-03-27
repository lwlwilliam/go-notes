/// 同时获取所有 url
package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

func main()  {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<- ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch error: %s: %v", url, err)
		return
	}

	// ioutil.Discard 可以看作是垃圾桶，不用的东西就扔里面
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

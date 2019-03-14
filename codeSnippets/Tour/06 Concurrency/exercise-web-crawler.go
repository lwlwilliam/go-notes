package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 方法返回 URL body 以及一个保存页面中 URL 的 slice
	Fetch(url string) (body string, urls []string, err error)
}

type URLs struct {
	mux sync.Mutex
	data map[string]string
}

var urlData URLs

// 不重复保存内容，由于 map 是非并发安全的，所以要加锁
func (u URLs) Add(url string, body string) {
	u.mux.Lock()

	if _, ok := u.data[url]; !ok {
		u.data[url] = body
	}

	u.mux.Unlock()
}

func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		//fmt.Println(err)
		return
	}

	//fmt.Printf("found: %s %q\n", url, body)
	urlData.Add(url, body)

	// 并发获取
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth - 1, fetcher, wg)
	}

	return
}

func main() {
	var wg sync.WaitGroup

	urlData.data = make(map[string]string)

	wg.Add(1)
	Crawl("https://golang.org/", 2, fetcher, &wg)
	wg.Wait()

	for k, v := range urlData.data {
		fmt.Println(k, v)
	}
}

// fakeFetcher 是返回封装结果的 Fetcher
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// 模拟获取网页 body 及网页中的链接
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// 模拟数据
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
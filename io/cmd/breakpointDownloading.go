// 断点下载
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// 与 chunk2 作为对比，这个函数不会出现数据丢失的问题
func chunk1(resp *http.Response) (s []byte, err error) {
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	s = buf.Bytes()

	return s, nil
}

// 这个在 Mac 机有点问题（在 err 为 nil 的前提下，Read() 返回的 n 与 step 偶尔出现不等，也就是读取不完全，并且出现的字节范围是随机的，例如有时候会在 1001-2000 的范围读不全，有时候会在 2001-3000 等等），在 Windows 下则不存在这问题
// TODO: 现在问题又来了，在 Mac 机时，如果 download() 的 step 设置比较小的话就不会出问题（至少暂时没出问题）。step 为 300 时已经开始出问题了
func chunk2(resp *http.Response) (s []byte, err error) {
	ck := make([]byte, 1024)
	n, err := resp.Body.Read(ck[0:])

	return ck[0:n], err
}

func download(totalContentLength int, count *int, step int, url string, file *os.File)  {
	start, end := 0, step
	for start < totalContentLength - 1 {
		// 分片下载
		req, err := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

		fmt.Printf("%d-%d\n", start, end)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}

		// 读取块
		hasRead, err := chunk1(resp)
		if err != nil && err != io.EOF {
			resp.Body.Close()
			log.Fatal("read", err)
		}
		n := len(hasRead)

		// 读取字节与预期不符
		if n < step {
			fmt.Println(resp.Status)
			for k, v := range resp.Header {
				fmt.Printf("%s: %s\n", k, strings.Join(v, ""))
			}
			fmt.Println(n)
		}

		// 读取总字节
		*count += n
		_, err = file.Write(hasRead)
		if err != nil && err != io.EOF {
			resp.Body.Close()
			log.Fatal(err)
		}

		start, end = end + 1, end + step
		resp.Body.Close()
	}
}

func main() {
	// 通过 HEAD 获取服务器及文件信息
	url := "http://notes.id/Writing an interpreter in Go.pdf"
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	acceptRanges := resp.Header.Get("Accept-Ranges")
	contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if acceptRanges != "bytes" {
		log.Fatal("Can't accept ranges.")
	}
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range resp.Header {
		fmt.Printf("%s: %v\n", k, strings.Join(v, ""))
	}

	// 准备文件
	file, err := os.OpenFile("test.pdf", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 开始下载
	count, step := 0, 300
	download(contentLength, &count, step, url, file)

	fmt.Printf("contentLength: %d; read from remote: %d\n", contentLength, count)
	log.Println("reach end")
}

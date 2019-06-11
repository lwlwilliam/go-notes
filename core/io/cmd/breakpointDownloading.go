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

func buff(resp *http.Response) (s []byte, err error) {
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	s = buf.Bytes()

	return s, nil
}

// 这个在 Mac 机有点问题，在 Windows 没问题
func chunk(resp *http.Response) (s []byte, err error) {
	ck := make([]byte, 1024)
	n, err := resp.Body.Read(ck[0:])

	return ck[0:n], err
}

func download(totalContentLength int, count *int, url string, file *os.File)  {
	start, end := 0, 1000
	for start < totalContentLength - 1 {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

		fmt.Printf("%d-%d\n", start, end)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		hasRead, err := buff(resp)
		if err != nil && err != io.EOF {
			log.Fatal("read", err)
		}

		n := len(hasRead)
		if n < 1000 {
			fmt.Println(resp.Status)
			for k, v := range resp.Header {
				fmt.Printf("%s: %s\n", k, strings.Join(v, ""))
			}
			fmt.Println(n)
		}
		*count += n
		_, err = file.Write(hasRead)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		start, end = end + 1, end + 1000
		resp.Body.Close()
	}
}

func main() {
	// 通过 HEAD 获取服务器及文件信息
	url := "http://notes.id/Writing%20an%20interpreter%20in%20Go.pdf"
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

	// 准备文件
	file, err := os.OpenFile("test.pdf", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 开始下载
	count := 0
	download(contentLength, &count, url, file)

	fmt.Printf("contentLength: %d; read from remote: %d\n", contentLength, count)
	log.Println("reach end")
}

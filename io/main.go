package io

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	url := "http://test.id/Writing%20an%20interpreter%20in%20Go.pdf"
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

	file, err := os.OpenFile("test.pdf", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start, end, count := 0, 1000, 0
	//buf := bytes.NewBuffer(nil)
	chunk := make([]byte, 1024)
	for start < contentLength-1 {
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

		//io.Copy(buf, resp.Body)
		//n := buf.Len()
		contentLengthStr := resp.Header.Get("Content-Length")
		contentLength, _ := strconv.Atoi(contentLengthStr)
		n, err := resp.Body.Read(chunk[0:contentLength])
		if n < 1000 {
			fmt.Println(resp.Status)
			for k, v := range resp.Header {
				fmt.Printf("%s: %s\n", k, strings.Join(v, ""))
			}
			fmt.Println(n)
		}

		if err != nil && err != io.EOF {
			log.Fatal("read", err)
		}
		count += n
		//file.Write(buf.Bytes())
		//buf.Reset()
		file.Write(chunk[0:n])

		start, end = end + 1, end + 1000
		resp.Body.Close()
	}

	fmt.Printf("contentLength: %d; read from remote: %d\n", contentLength, count)
	log.Println("reach end")
}

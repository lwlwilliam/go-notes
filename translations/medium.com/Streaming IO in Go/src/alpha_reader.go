package main

import (
	"io"
	"fmt"
)

type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

// if r is an alpha then return it, otherwise return nul
func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur	// (the len of the string that unread) + 1
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		n ++
		a.cur ++
	}
	copy(p, buf)
	return n, nil
}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sum?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
}

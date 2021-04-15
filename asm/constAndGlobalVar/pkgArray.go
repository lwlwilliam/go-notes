package main

import "github.com/lwlwilliam/go-notes/asm/constAndGlobalVar/pkg"

func main() {
	for k := range pkg.InitNum() {
		println(pkg.InitNum()[k])
	}
}

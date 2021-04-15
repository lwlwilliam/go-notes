package main

import "github.com/lwlwilliam/go-notes/asm/constAndGlobalVar/pkg"

func main() {
	for k := range pkg.GetCount() {
		println(pkg.GetCount()[k])
	}

	println(pkg.GetCount2())
}

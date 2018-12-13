package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

func main() {
	//dir, _ := os.Getwd()
	//fmt.Println(dir)

	contents, err := goconfig.LoadConfigFile("testdata/conf.ini")
	fatalErr(err)

	key, err := contents.GetValue("Demoa", "key1a")
	fatalErr(err)

	fmt.Println(key)
}

func fatalErr(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}
}

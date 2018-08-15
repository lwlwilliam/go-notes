package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	res, err := http.Get("https://golang.google.cn")
	checkError(err)

	data, err := ioutil.ReadAll(res.Body)
	checkError(err)

	// fmt.Printf("Got: %q\n", string(data))
	fmt.Printf("Got: %v\n", string(data))
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Get: %v", err)
	}
}
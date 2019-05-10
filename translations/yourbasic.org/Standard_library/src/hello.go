package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

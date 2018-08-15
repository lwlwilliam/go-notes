package main
import (
	"fmt"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Println(req.URL.Path)

	w.Header().Set("Content-Type", "text/html")

	// fmt.Fprint and more so fmt.Fprintf are good functions to use to write to the http.ResponseWrite(which implements io.Writer).
	// fmt.Fprint(w, "<h1>Hello," + req.URL.Path[1:], "</h1>\n")
	fmt.Fprint(w, "Hello," + req.URL.Path[1:], "\n")
	fmt.Println(req.URL.RawQuery)
}

func main() {
	// this form can also be used: http.Handle("/", http.HandlerFunc(HFunc)), HandlerFunc is just a type name for the signature `type HandlerFunc func(ResponseWriter, *Request)`
	http.HandleFunc("/", HelloServer)
	// if need the security of https, use http.ListenAndServeTLS() instead of http.ListenAndServe()
	err := http.ListenAndServe("localhost:8080", nil)

	// the two lines above can be replaced by this line:
	// err := http.ListenAndServe("localhost:8080", http.HandlerFunc(HelloServer))

	if err != nil {
		log.Fatal("ListenerAndServe:", err.Error())
	}
}
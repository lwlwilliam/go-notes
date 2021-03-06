package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type Pxy struct {}

func NewProxy() *Pxy {
	return &Pxy{}
}

// ServeHTTP is the main handler for all requests.
func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n",
		req.Method,
		req.Host,
		req.RemoteAddr,
	)

	if req.Method != "CONNECT" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		rw.Write([]byte("This is a http tunnel proxy, only CONNECT method is allowed."))
		return
	}

	// Step 1
	host := req.URL.Host
	hij, ok := rw.(http.Hijacker)
	if !ok {
		panic("HTTP Server does not support hijacking")
	}

	client, _, err := hij.Hijack()
	if err != nil {
		return
	}

	// Step 2
	server, err := net.Dial("tcp", host)
	if err != nil {
		return
	}
	client.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

	// Step 3
	go io.Copy(server, client)
	io.Copy(client, server)
}

func main() {
	proxy := NewProxy()
	http.ListenAndServe("0.0.0.0:60000", proxy)
}
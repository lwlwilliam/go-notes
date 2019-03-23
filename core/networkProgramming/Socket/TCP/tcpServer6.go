package main

import (
	"log"
	"net"
	"os"
)

var content = `HTTP/1.1 200 OK
Date: Sat, 29 Jul 2017 06:18:23 GMT
Content-Type: text/html
Connection: Keep-Alive
Server: BWS/1.1
X-UA-Compatible: IE=Edge,chrome=1
BDPAGETYPE: 3
Set-Cookie: BDSVRTM=0; path=/

<html>
<body>
	<h1>Hello world</h1>
</body>
</html>
`

func handleConn(conn net.Conn) {
	conn.Write([]byte(content))
	conn.Close()
}

func main()  {
	addr := "0.0.0.0:10000"
	listener, err := net.Listen("tcp", addr)
	checkErr(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

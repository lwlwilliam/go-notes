package main

import (
	"flag"
	"net"
	"fmt"
	"log"
)

func handler(server net.Conn, transitionPort int) {
	buf := make([]byte, 1024)
	for {
		// read the request from the server
		n, err := server.Read(buf)
		log.Printf("read from server <%s>\n", server.RemoteAddr())
		if err != nil {
			log.Println("server read:", err)
			continue
		}

		// link to the intranet service
		local, err := net.Dial("tcp", fmt.Sprintf(":%d", transitionPort))
		log.Printf("dial to local <%s>\n", local.RemoteAddr())
		if err != nil {
			log.Println("local dial:", err)
			continue
		}

		// write the request of the server to the local
		log.Printf("write to local <%s>\n", local.RemoteAddr())
		data := buf[:n]
		n, err = local.Write(data)
		if err != nil {
			log.Println("local write:", err)
			continue
		}

		// read the local response
		log.Printf("read from local <%s>\n", server.RemoteAddr())
		n, err = local.Read(buf)
		local.Close()
		if err != nil {
			log.Println("local read:", err)
			continue
		}

		// write the local response to the server
		data = buf[:n]
		n, err = server.Write(data)
		if err != nil {
			log.Println("server write:", err)
			continue
		}
	}
}

func main() {
	host := flag.String("h", "127.0.0.1", "server address")
	transitionPort := flag.Int("tp", 60001, "the port keeps link to the server")
	exportPort := flag.Int("ep", 80, "the port provides service to the internet")
	flag.Parse()

	server, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *transitionPort))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connected to the server <%s> from <%s>", server.RemoteAddr(), server.LocalAddr())

	go handler(server, *exportPort)

	select {}
}

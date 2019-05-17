package main

import (
	"flag"
	"os"
	"net"
	"fmt"
	"log"
)

func handler(remote net.Conn, localPort int) {
	buf := make([]byte, 1024)
	for {
		n, err := remote.Read(buf)
		if err != nil {
			continue
		}

		data := buf[:n]

		local, err := net.Dial("tcp", fmt.Sprintf(":%d", localPort))
		if err != nil {
			continue
		}

		n, err = local.Write(data)
		if err != nil {
			continue
		}

		n, err = local.Read(buf)
		local.Close()
		if err != nil {
			continue
		}

		data = buf[:n]

		n, err = remote.Write(data)
		if err != nil {
			continue
		}
	}
}

func main() {
	host := flag.String("h", "127.0.0.1", "server address")
	remotePort := flag.Int("rp", 8888, "server port")
	localPort := flag.Int("lp", 80, "local port")
	flag.Parse()
	if flag.NFlag() != 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	remote, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *remotePort))
	if err != nil {
		log.Fatal(err)
	}

	go handler(remote, *localPort)

	select {}
}

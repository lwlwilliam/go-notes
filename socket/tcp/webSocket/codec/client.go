package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"os"
	"flag"
	"github.com/lwlwilliam/go-notes/socket/tcp/webSocket/codec/xmlcodec"
)

type Person struct {
	Name   string
	Emails []string
}

func main() {
	ws := flag.String("ws", "ws://localhost:12345", "websocket")
	flag.Parse()

	conn, err := websocket.Dial(*ws, "", "http://localhost")
	checkError(err)

	person := Person{Name: "Jan",
		Emails: []string{"ja@newmarch.name", "jan.newmarch@gmail.com"},
	}

	err = xmlcodec.XMLCodec.Send(conn, &person)
	if err != nil {
		fmt.Println("Couldn't send msg " + err.Error())
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

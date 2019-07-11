package main

import (
	"flag"
	"golang.org/x/net/websocket"
	"fmt"
	"os"
	"io"
)

func main() {
	ws := flag.String("ws", "ws://localhost:12345", "websocket address")
	flag.Parse()

	conn, err := websocket.Dial(*ws, "", "http://localhost")
	checkError(err)
	var msg string
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Could't receive msg" + err.Error())
			break
		}
		fmt.Println("Received from server:" + msg)

		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Could't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	}
}

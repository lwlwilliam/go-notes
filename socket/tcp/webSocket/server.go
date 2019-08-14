package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		// receive
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)


		// send
		msg := "Hello world"
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

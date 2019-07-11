package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/websocket"
	"os"
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

	person := Person{
		Name:   "Jan",
		Emails: []string{"ja@newmarch.name", "jan.newmarch@gmail.com"},
	}

	err = websocket.JSON.Send(conn, person)
	//err = websocket.JSON.Send(conn, "234")
	if err != nil {
		fmt.Println("Could't send msg", err)
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err)
		os.Exit(1)
	}
}

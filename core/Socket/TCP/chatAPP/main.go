package main

import (
	"net"
	"flag"
	"strings"
	"fmt"
)

// hold all of the available clients, received data, and potential incoming or terminating clients.
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// hold information about the socket connection and data to be sent
type Client struct {
	socket net.Conn
	data   chan []byte
}

func startServerMode() {

}

func startClientMode() {

}

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		startServerMode()
	} else {
		startClientMode()
	}
}

// This goroutine will run for the lifespan of the server.
func (manager *ClientManager) start() {
	for {
		select {
		// If it reads data from the register channel,
		// the connection will be stored and a status will be printed in the logs.
		case connection := <- manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		// If the unregister channel has data and that data which represents a connection,
		// exists in our managed clients map, then the data channel for that connection will be closed
		// and the connection will be removed from the list.
		case connection := <- manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		// If the broadcast channel has data it means we've received a message. This message should be sent to
		// every connection we're watching so this is done by looping through the available connections.
		case message := <- manager.broadcast:
			for connection := range manager.clients {
				select {
				case connection.data <- message:
				// If we can't send the message to a client, that client is closed and removed from the list
				// of managed clients.
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

// The function will be a goroutine, which will exist for every connection that is established.
// For as long as the goroutine is available, it will be receiving data from a particular client.
func (manager *ClientManager) receiver(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		// If there was an error, for example the connection broke, the client will be unregistered and
		// formally closed.
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		// If everything went well and the message received wasn't empty, it will be added to the broadcast
		// channel to be distributed to all clients by the manager.
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
			manager.broadcast <- message
		}
	}
}

// distribute messages.
func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		// If the client has data to be sent and there are no errors, that data will be
		case message, ok := <- client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}
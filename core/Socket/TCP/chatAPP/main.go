package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	PROTOCOL = "tcp"
	HOST     = "localhost"
	PORT     = "10000"
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

// listen for connections
func startServerMode() {
	fmt.Println("Starting server...")
	listener, error := net.Listen(PROTOCOL, HOST + ":" + PORT)
	if error != nil {
		fmt.Println(error)
	}

	// initialize manager and start the manager goroutine
	manager := ClientManager{
		clients: make(map[*Client]bool),
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()

	// do a continuous loop listening for connections
	for {
		// If a connection is accepted, it will be registered
		// and prepared for sending and receiving of data.
		connection, _ := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receiver(client)
		go manager.send(client)
	}
}

func startClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial(PROTOCOL, HOST + ":" + PORT)
	if error != nil {
		fmt.Println(error)
	}
	client := &Client{socket: connection}
	go client.receive()
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		connection.Write([]byte(strings.TrimRight(message, "\n")))
	}
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
		case connection := <-manager.register:
			manager.clients[connection] = true
			fmt.Println("Added new connection!")
		// If the unregister channel has data and that data which represents a connection,
		// exists in our managed clients map, then the data channel for that connection will be closed
		// and the connection will be removed from the list.
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				fmt.Println("A connection has terminated!")
			}
		// If the broadcast channel has data it means we've received a message. This message should be sent to
		// every connection we're watching so this is done by looping through the available connections.
		case message := <-manager.broadcast:
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
		// sent to the client in question. If there is an error and the loop breaks, the
		// connection to the socket will end.
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}

// the difference between client and server is that we aren't taking
// our received messages to a central processor.
func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}
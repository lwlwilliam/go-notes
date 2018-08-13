package main
import (
	"bufio"
	"net"
	"os"
	"fmt"
	"strings"
)

func main() {
	// open connection:
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	// read from stdin
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	// read inputted content until the carriage return
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\r\n")	// windows 换行

	// send info to server until Quit:
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// read the Q and quit
		if trimmedInput == "Q" {
			return
		}

		// send info to the server
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))

		// output := make([]byte, 512)
		// _, err := conn.Read(output)

		if err != nil {
			fmt.Println("Error reading", err.Error())
		}

		// fmt.Printf("Server received: %v", output)
	}
}

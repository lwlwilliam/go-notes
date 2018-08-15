// Simple multi-thread/multi-core TCP server.
package main
import (
	"syscall"
	"flag"
	"net"
	"fmt"
)

const maxRead = 100

func main() {
	// parse the parameters from the command line, if the amount of parameters is not 2 then panic and the program is terminated.
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	// concat the two parameters from the command line.
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))

	// start the server
	listener := initServer(hostAndPort)

	// loop to accept message from the client
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		// handle the message from the client
		go connectionHandler(conn)
	}
}

// start the server and listen to it
func initServer(hostAndPort string) *net.TCPListener {
	// resolve the tcp address and listen to it
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '" + hostAndPort + "'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to:", listener.Addr().String())
	return listener
}

// handle the connection with client
func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from:", connFrom)

	// say hello to the client first
	sayHello(conn)

	// then loop to read from the client
	for {
		var ibuf []byte = make([]byte, maxRead + 1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0  // to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		case syscall.EAGAIN:  // try again
			continue
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	println("Closed connection:", connFrom)
	checkError(err, "Close: ")
}

// write "Let's GO!" to the connection
func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote " + string(wrote) + " bytes.")
}

// handle the message read from the client
func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		// print the message
		print("<", length, ":")
		for i := 0; ; i ++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">\n")
	}
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error())  // terminate program
	}
}

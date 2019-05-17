package main

import (
	"flag"
	"fmt"
	"net"
	"log"
)

type transition struct {
	internet  *net.TCPListener
	intranet  *net.TCPListener
	messages  map[int]*message
	messageID int
}

type message struct {
	id          int
	internet    net.Conn
	intranet    net.Conn
	internetMsg chan []byte
	intranetMsg chan []byte
}

func New() *transition {
	return &transition{
		messages:  make(map[int]*message),
		messageID: 0,
	}
}

func (t *transition) listen(internetPort int, intranetPort int) error {
	// listen for the internet
	socket, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", internetPort))
	log.Println("resolve the internet addr...")
	if err != nil {
		return err
	}
	t.internet, err = net.ListenTCP("tcp", socket)
	log.Println("listen to the internet...")
	if err != nil {
		return err
	}

	// listen for the intranet
	socket, err = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", intranetPort))
	log.Println("resolve the intranet addr...")
	if err != nil {
		return err
	}
	t.intranet, err = net.ListenTCP("tcp", socket)
	log.Println("listen to the intranet...")
	if err != nil {
		return err
	}

	go t.accept()
	return nil
}

func (t *transition) close() {
	t.internet.Close()
	t.intranet.Close()

	for _, v := range t.messages {
		v.internet.Close()
		v.intranet.Close()
	}
}

func (t *transition) delChannel(id int) {
	msgs := t.messages
	delete(msgs, id)
	t.messages = msgs
}

func (t *transition) accept() {
	// keep the link of intranet
	intranet, err := t.intranet.Accept()
	if err != nil {
		log.Println("accept intranet error:", err)
		return
	}
	log.Printf("accept from intranet <%s>", intranet.RemoteAddr())

	// accept multi links from the internet
	for {
		conn, err := t.internet.Accept()
		if err != nil {
			log.Println("accept internet error:", err)
			continue
		}
		log.Printf("accept from internet <%s>", conn.RemoteAddr())

		msg := &message{
			id:          t.messageID,
			internet:    conn,
			intranet:    intranet,
			internetMsg: make(chan []byte),
			intranetMsg: make(chan []byte),
		}
		t.messageID++

		msgs := t.messages
		msgs[msg.id] = msg
		t.messages = msgs

		go t.writeToIntranet(msg)
		go t.writeToInternet(msg)

		go t.read(msg)
	}
}

// accept the messages from internet and write them to intranet
func (t *transition) writeToIntranet(msg *message) {
	defer func() {
		fmt.Println("internetMsg exit")
	}()

	for {
		select {
		case data, isClose := <-msg.internetMsg:
			if !isClose {
				return
			}

			_, err := msg.intranet.Write(data)
			if err != nil {
				return
			}
		}
	}
}

// accept the messages from intranet and write them to internet
func (t *transition) writeToInternet(msg *message) {
	defer func() {
		fmt.Println("intranetMsg exit")
	}()

	for {
		select {
		case data, isClose := <-msg.intranetMsg:
			if !isClose {
				return
			}

			_, err := msg.internet.Write(data)
			if err != nil {
				return
			}
		}
	}
}

func (t *transition) read(msg *message) {
	defer func() {
		close(msg.internetMsg)
		close(msg.intranetMsg)
		t.delChannel(msg.id)
		fmt.Println("read exit")
	}()

	buf := make([]byte, 1024)
	for {
		n, err := msg.internet.Read(buf)
		if err != nil {
			log.Println("internet read error:", err)
			return
		}

		log.Printf("%s", buf[:n])
		msg.internetMsg <- buf[:n]
		n, err = msg.intranet.Read(buf)
		if err != nil {
			log.Println("intranet read error:", err)
			return
		}

		log.Printf("%s", buf[:n])
		msg.intranetMsg <- buf[:n]
	}
}

func main() {
	internetPort := flag.Int("ter", 60000, "the port exports to the internet")
	intranetPort := flag.Int("tra", 60001, "the port links to the intranet")
	flag.Parse()

	t := New()

	err := t.listen(*internetPort, *intranetPort)
	if err != nil {
		log.Fatal(err)
	}

	select {}
}

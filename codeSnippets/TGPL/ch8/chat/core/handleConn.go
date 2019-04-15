// 为客户端创建一个消息发出 channel 并通过 entering channel 来通知客户端的到来。然后它会读取客户端发来的每一行文本，
// 并通过全局的消息 channel 来将这些文本发送出去，并会每条消息带上发送者的前缀来标明消息身份。当客户端发送完毕后，
// handleConn 会通过 leaving 这个 channel 来通知客户端的离开并关闭连接。
package core

import (
	"net"
	"bufio"
	"fmt"
)

func HandleConn(conn net.Conn)  {
	ch := make(chan string) // outgoing client message
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string)  {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}
// 内部的 clients 会记录当前建立连接的客户端集合，其记录的内容是每一个客户端的消息发出 channel 的"资格"信息
// broadcaster 监听来自全局的 entering 和 leaving 的 channel 来获知客户端的到来和离开事件。当其接收到
// 其中的一个事件时，会更新 clients 集合。当该事件是离开行为时，它会关闭客户端的消息发出 channel。broadcaster
// 也会监听全局的消息 channel，所有的客户端都会都会向这个 channel 中发送消息。当 broadcaster 接收到什么消息时，
// 就会将其广播至所有连接到服务端的客户端。
package core

type client chan <- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func Broadcaster()  {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <- messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			// 向所有客户端广播收到的消息
			for cli := range clients {
				cli <- msg
			}

		// 用户进入聊天室
		// 添加用户到 clients 中
		case cli := <- entering:
			clients[cli] = true

		// 用户离开聊天室
		// 从 clients 中删除该用户
		case cli := <- leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

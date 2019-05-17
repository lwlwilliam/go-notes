package main

import (
	"net"
	"fmt"
	"flag"
	"os"
)

type MidServer struct {
	// 客户端监听
	clientLis *net.TCPListener
	// 后端服务连接
	transferLis *net.TCPListener
	// 所有通道
	channels map[int]*Channel
	// 当前通道 id
	curChannelID int
}

type Channel struct {
	// 通道 id
	id int
	// 客户端连接
	client net.Conn
	// 后端服务连接
	transfer net.Conn
	// 客户端接收消息
	clientRecvMsg chan []byte
	// 后端服务发送消息
	transferSendMsg chan []byte
}

func New() *MidServer {
	return &MidServer{
		channels: make(map[int]*Channel),
		curChannelID: 0,
	}
}

// 启动服务
func (m *MidServer) Start(clientPort int, transferPort int) error {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", clientPort))
	if err != nil {
		return err
	}
	m.clientLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	addr, err = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", transferPort))
	if err != nil {
		return err
	}
	m.transferLis, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	go m.AcceptLoop()
	return nil
}

func (m *MidServer) Stop() {
	m.clientLis.Close()
	m.transferLis.Close()

	// 循环关闭通道连接
	for _, v := range m.channels {
		v.client.Close()
		v.transfer.Close()
	}
}

// 删除通道
func (m *MidServer) DelChannel(id int) {
	chs := m.channels
	delete(chs, id)
	m.channels = chs
}

// 处理连接
func (m *MidServer) AcceptLoop() {
	transfer, err := m.transferLis.Accept()
	if err != nil {
		return
	}

	for {
		client, err := m.clientLis.Accept()
		if err != nil {
			continue
		}

		// 创建通道
		ch := &Channel{
			id:m.curChannelID,
			client:client,
			transfer:transfer,
			clientRecvMsg:make(chan []byte),
			transferSendMsg:make(chan []byte),
		}
		m.curChannelID ++

		// 把通道放入 channels 中
		chs := m.channels
		chs[ch.id] = ch
		m.channels = chs

		// 启动一个 goroutine 处理客户端消息
		go m.ClientMsgLoop(ch)
		// 启动一个 goroutine 处理后端服务消息
		go m.TransferMsgLoop(ch)

		go m.MsgLoop(ch)
	}
}

// 处理客户端消息
func (m *MidServer) ClientMsgLoop(ch *Channel) {
	defer func() {
		fmt.Println("ClientMsgLoop exit")
	}()

	for {
		select {
		case data, isClose := <- ch.transferSendMsg:
			if !isClose {
				return
			}

			_, err := ch.client.Write(data)
			if err != nil {
				return
			}
		}
	}
}

func (m *MidServer) TransferMsgLoop(ch *Channel) {
	defer func() {
		fmt.Println("TransferMsgLoop exit")
	}()

	for {
		select {
		case data, isClose := <- ch.clientRecvMsg:
			if !isClose {
				return
			}

			_, err := ch.transfer.Write(data)
			if err != nil {
				return
			}
		}
	}
}

func (m *MidServer) MsgLoop(ch *Channel) {
	defer func() {
		// 关闭 channel, 好让 ClientMsgLoop 与 TransferMsgLoop 退出
		close(ch.clientRecvMsg)
		close(ch.transferSendMsg)
		m.DelChannel(ch.id)
		fmt.Println("MsgLoop exit")
	}()

	buf := make([]byte, 1024)
	for {
		n, err := ch.client.Read(buf)
		if err != nil {
			return
		}
		ch.clientRecvMsg <- buf[:n]
		n, err = ch.transfer.Read(buf)
		if err != nil {
			return
		}
		ch.transferSendMsg <- buf[:n]
	}
}

func main() {
	localPort := flag.Int("lp", 8080, "客户端访问端口")
	remotePort := flag.Int("rp", 8888, "服务访问端口")
	flag.Parse()
	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	ms := New()
	// 启动服务
	ms.Start(*localPort, *remotePort)

	select{}
}

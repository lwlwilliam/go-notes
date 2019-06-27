// 服务器的 tcp 连接数是有限制的
// 对方服务的 listen backlog 满
// TODO:理论上貌似是可以通过连接来填满服务器的 backlog 达到拒绝服务攻击的目的
package main

import (
	"net"
	"log"
)

func foreverConn()  {
	var sl []net.Conn
	// 串行地尝试建立连接，在 mac 下会一次性建立 128 个连接，然后后续每阻塞 10 s 才成功建立一条连接。TODO: 这个连接数量以及阻塞 10 s 才成功建立一条连接需要斟酌一下
	// 也就是说在 server 端 backlog 满时（未及时 accept），客户端将阻塞在 Dial 上，直到 server 端进行一次 accept。至于为何是 128，这与 darwin 下的默认设置有关：
	// $ sysctl -a | grep kern.ipc.somaxconn
	// kern.ipc.somaxconn: 128
	for i := 1; ; i ++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}
}

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", "localhost:80")

	// 等价于 net.Dial
	//raddr, err := net.ResolveTCPAddr("tcp", "localhost:80")
	//conn, err = net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Printf("%d: dial error: %s\n", i, err)
		return nil
	}

	log.Println(i, ":connect to server ok")
	return conn
}

func main()  {
	foreverConn()
}

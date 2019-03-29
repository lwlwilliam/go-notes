// 对方服务的 listen backlog 满
package main

import (
	"net"
	"log"
	"time"
)

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Printf("%d: dial error: %s\n", i, err)
		return nil
	}
	log.Println(i, ":connect to server ok")
	return conn
}

func main()  {
	var sl []net.Conn
	// 串行地尝试建立连接，在 mac 下会一次性建立 128 个连接，然后后续每阻塞 10 s 才成功建立一条连接。
	// 也就是说在 server 端 backuplog 满时（未及时 accept），客户端将阻塞在 Dial 上，直到 server 端
	// 进行一次 accept。至于为何是 128，这与 darwin 下的默认设置有关：
	// $ sysctl -a | grep kern.ipc.somaxconn
	// kern.ipc.somaxconn: 128
	for i := 1; i < 1000; i ++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}

	time.Sleep(time.Second * 10000)
}

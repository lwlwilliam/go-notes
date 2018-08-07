// select 语句用来从多个 send/receive channel 操作中进行选择。
// 该语句会一直阻塞直到 send/receive 中的某个操作已经就绪。如果多个操作都处于就绪状态，select 会
// 随机选择其中的一个。除了每一个 case 语句都是 channel 操作外，它与 switch 都很类似。

// 假设数据库复制并储存在不同的服务器，每台服务器的响应时间取决于自身的负荷及网络延迟。我们向所有服务
// 器发送请求，然后使用 select 语句对响应的 channel 进行等待。首先响应的服务器被 select 选中，其他的
// 响应则被忽略。因此，我们可以把同一个请求发送给多台服务器并把最快的响应返回给用户。
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- output2:
		fmt.Println(s2)
	case <- time.After(2 * time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
}
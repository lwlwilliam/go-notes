// 使用 channel 解决竞争状态
// 以下程序创建了一个缓冲量为 1 的 channel，并传递到 increment goroutine 中。该 channel 用于确保唯一的
// goroutine 访问递增 x 的临界资源。
// 由于该 channel 缓冲量为 1，其他 goroutine 尝试往 channel 中写入时会阻塞直至 channel 中的值被读取，也就
// 是在 x 递增之后。
package main
import (
	"fmt"
	"sync"
)

var x = 0
func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i ++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
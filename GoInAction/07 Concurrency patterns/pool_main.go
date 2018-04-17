// 这个示例程序展示如何使用 pool 包来共享一组模拟的数据库连接
package main

import (
	"log"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"./pool"
)

const (
	maxGoroutines   = 25  // 要使用的 goroutine 的数量
	pooledResources = 2   // 池中的资源的数量
)

// dbConnection 模拟了管理数据库连接的结构，当前版本只包含一个字段 ID，用来保存每个连接的唯一标识
type dbConnection struct {
	ID int32
}

// Close 实现了 io.Closer 接口，以便 dbConnection 可以被池管理。Close 用来完成任意资源的释放管理
// 这里只是报告了连接正在被关闭，并显示出要关闭连接的标识
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 用来给每个连接分配一个独一无二的 id
var idCounter int32

// createConnection 是一个工厂函数，
// 当需要一个新连接时，资源池会调用这个函数
// 这个函数给连接生成了一个唯一标识，显示连接正在被创建，并返回指向带有唯一标识的 dbConnection 类型值的指针
// 唯一标识是通过 atomic.AddInt32 函数生成的。这个函数可以安全地增加包级变量 idCounter 的值
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建用来管理连接的池，传入工厂函数和要管理的资源的数量，返回一个指向 Pool 值的指针，并检查可能的错误
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query ++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Shutdown Program.")
	p.Close()
}

// perforQueries 用来测试连接的资源池
func performQueries(query int, p *pool.Pool) {
	// 从池里请求一个连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将该连接放回池里
	defer p.Release(conn)

	// 用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
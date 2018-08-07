// 有缓冲的 channel 一个重要的应用就是 worker pool 的实现。
// 通常，worker pool 是任务线程的集合。一旦它们完成了分配的任务，就会令自身可重用。
// 以下是 worker pool 的核心功能：
// 		1.	创建 goroutine pool 监听缓冲 input channel 等待分配任务；
//		2.	添加任务到缓冲 input channel；
//		3.	在任务结束后，往缓冲 output channel 中写入结果；
//		4.	读取并打印从有缓冲的 output channel 中得到的结果；	
package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

// 创建 Job 和 Result 表示任务和结果
type Job struct {
	id			int
	randomno	int
}
type Result struct {
	job			Job
	sumofdigits	int
}

// 创建缓冲 jobs 及 results channel
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// 每个 job 休眠 1 秒，模拟现实任务
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

// 读取 jobs channel 中的所有数据，并把对应的结果写入 Result 并发送到 results channel 中，完成后标记 wg.Done()
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// 创建工作池，根据参数决定要创建的 job，每创建一个 job 都设置 wg.Add(1)。创建完 job 后，wg.Wait() 等待执行。
// 所有 job 完成后，关闭 results channel。
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i ++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	// 与 allocate 函数不同，results channel 需要在这里关闭，因为 worker 函数只是相当于 createWorkerPool 的工作子线程。
	close(results)
}

// 分配任务。根据参数决定要分配的任务数量。分配完任务后，关闭 jobs channel。
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i ++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

// 从 results channel 读取任务结果。读取结束后往 done channel 发送信号。
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()

	// 分配指定数据的任务
	noOfJobs := 100
	go allocate(noOfJobs)

	// 等待结果信号
	done := make(chan bool)
	go result(done)

	// 创建指定数量的 worker pool
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)

	// 结束信号
	<- done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
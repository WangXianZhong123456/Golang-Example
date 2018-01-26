// 这个示例程序展示了如何使用
//有缓冲的通道来处理一堆工作
// 携程池
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGroutinees = 4
	taskLoad         = 5
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {
	//创建缓冲区10的chan
	tasks := make(chan string, taskLoad)
	wg.Add(numberGroutinees)
	for gr := 1; gr <= numberGroutinees; gr++ {
		go worker(tasks, gr)
	}
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task1:%d", post)
	}
	time.Sleep(time.Duration(time.Second * 5))
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task2:%d", post)
	}
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	//defer fmt.Println("123")
	defer wg.Add(-1)

	for {
		// 等待分配工作
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d :Shutting Down\n", worker)
			return
		}
		fmt.Printf("Worker :%d----Started %s\n", worker, task)
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker %d :completed%s\n", worker, task)
	}
}

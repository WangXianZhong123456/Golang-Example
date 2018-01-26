package main

import "fmt"
import "math/rand"
import "sync"
import "time"

const (
	numberGroutiness = 4
	taskLoad         = 5
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGroutiness)
	for gr := 1; gr <= numberGroutiness; gr++ {
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
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d :Shuting Down\n", worker)
			return
		}

		fmt.Printf("Worker :%d----Started %s\n", worker, task)
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d:completed %s\n", worker, task)
	}
}

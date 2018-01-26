package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)
	wg.Add(1)
	// 第一个 接过了接力棒,阻塞状态，等待发令
	go Runner(baton)
	//发令，比赛开始，启动阻塞Runner携程
	baton <- 1
	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d is Running With Baton \n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the line,Ready to get the baton \n", newRunner)
		//		baton <- newRunner
	}

	if newRunner != 1 {
		go Runner(baton)
	}

	time.Sleep(3 * time.Second)

	if runner == 4 {
		fmt.Printf("Runner %d finished,Race Over", runner)
		wg.Done()
		return

	}

	fmt.Printf("Runner %d Exchange With Runner%d", runner, newRunner)
	baton <- newRunner

}

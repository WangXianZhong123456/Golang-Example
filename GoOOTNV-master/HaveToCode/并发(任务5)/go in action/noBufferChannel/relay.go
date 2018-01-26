// 这个例子展示如何用无缓冲通道来模拟接力赛
// 4个goroutine间的接力比赛
package main

import (
	"fmt"
	"sync"
	"time"
)

// 用来等待程序结束
var wg sync.WaitGroup

func main() {
	//创建一个无缓冲通道，控制携程的同步-----这是个接力棒啊，一根棒控制4个携程
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
	//等待接力棒
	runner := <-baton

	// 开始跑步啦
	fmt.Printf("Runner %d is Running With Baton\n", runner)
	// 创建下一个跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To the line,Ready to get the baton\n", newRunner)
		//go Runner(baton)
	}
	if newRunner != 1 {
		go Runner(baton)
	}

	//正在跑啊
	time.Sleep(time.Second * 2)
	if runner == 4 {
		fmt.Printf("Runner %d finished,Race Over", runner)
		wg.Done()
		return
	}
	// 交接接力棒
	fmt.Printf("Runner %d Exchange With Runner%d", runner, newRunner)
	baton <- newRunner
}

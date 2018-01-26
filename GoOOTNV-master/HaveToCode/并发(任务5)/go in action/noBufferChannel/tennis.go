// 本例子是go in action 中的棒球例子

//2个goroutine之间的网球比赛，无缓冲通道间共享count
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//create a no buffer channel
	court := make(chan int)
	// represet 2 goroutine
	wg.Add(2)

	go player("me", court)
	go player("you", court)
	// go player("3",court)
	time.Sleep(time.Second * 3)
	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed \n", name)
			close(court)
			return
		}
		fmt.Printf("player %s hit %d\n", name, ball)
		ball++

		court <- ball
	}
}

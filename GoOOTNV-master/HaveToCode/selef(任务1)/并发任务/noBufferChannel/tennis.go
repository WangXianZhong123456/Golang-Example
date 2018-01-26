package main

import (
	"fmt"
	"math/rand"
	"sync"
	//"time"
)

var wg sync.WaitGroup

func main() {
	court := make(chan int)
	wg.Add(2)

	go Player("me", court)
	go Player("you", court)
	court <- 1
	wg.Wait()
}

func Player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("player %s won \n", name)
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed \n", name)
			close(court)
			return
		}
		fmt.Printf("player %s hit %d \n", name, ball)
		ball++
		court <- ball
	}

}

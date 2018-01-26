package main

import "fmt"

//随机打印0 1 01 01 10 10 10，死循环
func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i, _ := <-ch
		fmt.Println("Value received:", i)
	}
}

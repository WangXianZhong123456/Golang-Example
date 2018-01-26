package main

import (
	"fmt"
	"time"
)

func main() {
	//w无缓冲通道用来控制主程序不提前退出

	//c有缓冲通道，送入一个数据不会阻塞，继续执行
	//c1:=make(chan int)  无缓冲

	//c2:=make(chan int,1)  有缓冲
	//无缓冲：不仅仅是向 c1 通道放 1，而是一直要等有别的携程 <-c1 接手了这个参数，
	//那么c1<-1才会继续下去，要不然就一直阻塞着。

	//有缓冲： c2<-1 则不会阻塞，因为缓冲大小是1(其实是缓冲大小为0)，
	//只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。
	c := make(chan int, 10) // 会先写入缓存,如果没有第11个写入的话，是不会去执行gofunc
	//w := make(chan bool)
	go func() {
		fmt.Println("time out2")
		for {
			select {
			case e, _ := <-c:
				//time.Sleep(3 * time.Second)
				fmt.Println(e)

				//c <- 99
			//3秒后写入超时
			case <-time.After(time.Second * 3):
				fmt.Println("time out")

			}
			//w <- true
		}
		//c <- 8

	}()

	for i := 0; i < 11; i++ {
		if i >= 10 {
			c <- i
		} else {
			c <- i // 会先写入缓存,如果没有第11个写入的话，是不会去执行gofunc
		}

		//<-w
		//<-c
	}
	//time.Sleep(1 * time.Second)

	//<-c
	//close(c)
	//c <-1 注释掉，引发time out
	/*c2 := make(chan bool, 0)
	go func() {
		fmt.Println("case1 Go Go Go")
		<-c2
		//close(c2)
		//c2 <- true
	}()
	c2 <- true*/
}

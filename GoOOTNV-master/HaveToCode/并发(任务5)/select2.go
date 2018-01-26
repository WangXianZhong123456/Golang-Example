package main

import (
	"fmt"
	"time"
)

func main() {
	// w无缓冲通道用来控制主程序不提前退出
	w := make(chan bool)
	// c有缓冲通道，送入一个数据不会阻塞，继续执行
	c := make(chan int, 2)
	go func() {

		select {
		case e, _ := <-c:
			fmt.Println(e)
		//3秒后写入超时
		case <-time.After(time.Second * 3):
			fmt.Println("time out")
		}

		w <- true
	}()
	c <- 9
	//c <-1 注释掉，引发time out
	<-w
}

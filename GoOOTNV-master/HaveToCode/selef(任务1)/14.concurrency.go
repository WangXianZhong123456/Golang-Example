package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*并发*/
func main() {
	printCase(1)
	go Go(1) /*如果不加sleep就不回输出*/
	time.Sleep(1 * time.Second)

	/*1、首先，永远是符号<-进行读取或者写入，譬如v,ok := <-c是读取，而c <- v是写入。
	  2、其次，读取时，如果没有ok，也是可以读取的。不过如果closed也是能读的，没有赋值而已；
		如果要知道是否closed得加ok，也就是除非chan永远不关闭，否则读取应该用v,ok := <-c而不是用v := <-c的方式。
	  3、再次，不能向closed的chan写入，所以一般写入时需要用一个信号的chan（一般buffer为1），
		来判断是否写入或者放弃，用select判断是写入成功了，还是正在关闭需要放弃写入。
	  4、最后，如果closed后，chan有数据，ok还是true的，直到chan没有数据了才false。*/

	printCase(2)
	c2 := make(chan bool, 11)
	go func() {
		Go(2)
		c2 <- true
		//close(c2)
		//c2 <- true
	}()
	<-c2

	printCase(3)
	c3 := make(chan bool)
	go func() {
		Go(3)
		c3 <- true
		close(c3) // 当close的时候要读取,否则不回调用
	}()
	//<-c3

	//fmt.Println("case 3, receive:", v)
	for v := range c3 { // 相当于 v, _ := <-c3 // c=10,ok=true，读取出来一个
		fmt.Println("case 3, receive:", v)
	}

	printCase(5)
	c5 := make(chan bool, 1)
	go func() {
		Go(5)
		<-c5
	}()
	c5 <- true

	/*select {
	case v, _ := <-c5:
		fmt.Println("case 5, receive:", v)
	default: // 没有可读的，走这个分支
	}*/
	// case 6: 使用并发, 在最后一个并发中往channel中写入值通知主线程结束,
	// BAD: 最后一个启动的并发不一定是最后结束, 多核cpu调度不是按启动顺序调度的
	printCase(6)
	runtime.GOMAXPROCS(runtime.NumCPU()) // 经测试,在mac中不需要设置这个参数也能使用多核
	c6 := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go2(c6, i)
	}
	<-c6

	// case 7: fix case 6, 所有启动的并发都往channel中写入执行结束通知, 主线程从channel中读取10次
	printCase(7)
	c7 := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go3(c7, i)
	}

	for i := 0; i < 10; i++ {
		//fmt.Println("case 7 i=", i) 你会发现执行的先后顺序不一样
		<-c7
	}

	// case 8: fix case 6, 通过sync.WaitGroup方法等待所有任务结束
	/*说说WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。
	  1、WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。
	  2、Add:添加或者减少等待goroutine的数量
	  3、Done:相当于Add(-1)
	  4、Wait:执行阻塞，直到所有的WaitGroup数量变成0*/
	printCase(8)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go4(&wg, i) // 注意这里传递的是指针,非值拷贝
	}
	wg.Wait()

	/*要注意的时，select和switch有点类似，但select的case后只能是IO操作。
	  上述程序的执行过程是这样的，程序进入select后，如果没有case中的channel可读，则阻塞，直到有channel可读；
	  如果仅有一个channel可读，则执行这个case；如果有多个channel可读，则随机公平地选出一个case执行，其他不会执行。*/
	printCase(9)
	c9x, c9y := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		a, b := false, false
		for {
			select {
			case v, ok := <-c9x:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
					if !a {
						a = true
						o <- true
						break
					}
				}
				fmt.Println("c9x", v)
			case v, ok := <-c9y:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
					if !b {
						b = true
						o <- true
						break
					}
				}
				fmt.Println("c9y", v)
			}
		}
	}()
	c9x <- 1
	c9y <- "hi"
	c9x <- 3
	c9y <- "hello"
	close(c9x)
	close(c9y)
	for i := 0; i < 2; i++ {
		<-o
	}

	// case 10: 空select{}会阻塞?, 可用于阻塞main函数不退出?
	// 经测试,最近版本(1.6)不能使用此方法,报错: goroutine 1 [select (no cases)]:

	printCase(10)
	//select()

	// case 11: 为select设置超时
	printCase(11)
	c11 := make(chan bool)
	select {
	case v, _ := <-c11:
		fmt.Println(v)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}

	// case 12: 作业
	printCase(12)
	c12 := make(chan string)
	go func() { //Pingpong
		//i := 0
		for {
			fmt.Println(<-c12)
			//c12 <- fmt.Sprintf("From Pingpong: Hi, #%d", 2)
			//i++
		}
	}()

	for i := 0; i < 10; i++ {
		c12 <- fmt.Sprintf("From main: Hello, #%d", i)
		//fmt.Println(<-c12)

	}
	close(c12)

}

func printCase(cas int) {
	fmt.Println("------------- case ", cas, "-------------")
}

func Go(cas int) {
	fmt.Println("case", cas, "Go Go Go")
}

func Go2(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go2", index, a)
	if index == 9 {
		c <- true
	}
}

// 每个任务都写入channel方案
func Go3(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go3", index, a)
	c <- true
}

// sync解决方案, 注意WaitGroup是传递的指针
func Go4(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println("Go4", index, a)
	wg.Done()
}

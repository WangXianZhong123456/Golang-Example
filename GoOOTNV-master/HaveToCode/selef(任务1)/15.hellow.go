package main

import (
	"fmt"
	"time"
)

func PingPong1(s []int) {
	s = append(s, 3)
}

func PingPong2(s []int) []int {
	s = append(s, 3)
	return s
}
func main() {
	s1 := make([]int, 0)
	PingPong1(s1)
	fmt.Println(s1)

	s2 := make([]int, 0)
	s2 = PingPong2(s2)
	fmt.Println(s2)

	t1 := time.Now()
	fmt.Println(t1.Format(time.RFC3339))
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
	fmt.Println(t1.Format("Mon Jan _2 15:04:06 2006")) //输出时间倒退了

	s := []string{"a", "b", "c"}
	fmt.Println(s)

	// case 4, 闭包问题,下面匿名函数输出的都是"c"
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(1 * time.Second)

	s = []string{"a", "b", "c"}
	fmt.Println(s)
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}

	time.Sleep(1 * time.Second)
}

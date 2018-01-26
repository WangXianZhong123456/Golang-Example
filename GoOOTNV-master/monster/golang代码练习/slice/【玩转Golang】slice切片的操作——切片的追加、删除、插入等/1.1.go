//1,声明变量，go自动初始化为nil，长度：0，地址：0，nil
package _玩转Golang_slice切片的操作__切片的追加_删除_插入等

import "fmt"

func main() {
	var ss []string
	fmt.Printf("length:%v \taddr:%p \tisnil:%v", len(ss), ss, ss == nil)
}

//---
//Running...
//
//length:0     addr:0x0     isnil:true
//Success: process exited with code 0.

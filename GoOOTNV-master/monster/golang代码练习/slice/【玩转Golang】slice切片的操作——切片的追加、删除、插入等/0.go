// slice 指针会因为append 发生改变吗？容量超过时会。
package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	fmt.Printf("slice指针:%p \t newslice指针:%p", slice, newSlice)
	// newSlice 的长度是2，容量是4
	//append10 会改变底层数组的值 , slice 也会随之改变
	fmt.Printf("newSlice: %#v \t slice :%#v\n", newSlice, slice)
	newSlice = append(newSlice, 10)
	fmt.Printf("slice指针:%p \t newslice指针:%p", slice, newSlice)
	fmt.Printf("newSlice: %#v \t slice :%#v\n", newSlice, slice)

	slice = append(slice, 11)
	//*****当slice 添加数据时容量超过，底层指向的数组会拷贝一份，slice指针发生改变，此时再改变newslice的值，不会影响slice的值
	fmt.Printf("slice指针:%p \t newslice指针:%p", slice, newSlice)
	fmt.Printf("newSlice: %#v \t slice :%#v\n", newSlice, slice)

	newSlice = append(newSlice, 11)
	fmt.Printf("slice指针:%p \t newslice指针:%p", slice, newSlice)
	fmt.Printf("newSlice: %#v \t slice :%#v\n", newSlice, slice)

}

package main

import "fmt"

//　2,切片的追加，删除，插入操作

func main(){
	var ss []string;
	fmt.Printf("[ local print ]\t:\t length:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("func print",ss,"\n")
	//切片尾部追加元素append elemnt
	for i:=0;i<10;i++{
		ss=append(ss,fmt.Sprintf("s%d",i));
	}
	fmt.Printf("[ local print ]\t:\tlength:%v\taddr:%p\tisnil:%v\n",len(ss),ss, ss==nil)
	print("after append",ss,"\n")
	//删除切片元素remove element at index
	//这部分操作在music app 中涉及。golang中slice的删除需要手动，可以自己封装
	index:=5;
	ss=append(ss[:index],ss[index+1:]...)
	print("after delete",ss,"\n")
	//在切片中间插入元素insert element at index;
	//注意：保存后部剩余元素，必须新建一个临时切片
	rear:=append([]string{},ss[index:]...)
	ss = append(ss[0:index],"insertElement")
	ss = append(ss,rear...)
	print("after insert",ss,"\n")
}
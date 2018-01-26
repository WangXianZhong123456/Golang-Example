//二、初始大小和容量
package main

import "fmt"

//当我们使用make初始化切片的时候，必须给出size。go语言的书上一般都会告诉我们，当切片有足够大小的时候，append操作是非常快的。但是当给出初始大小后，我们得到的实际上是一个含有这个size数量切片类型的空元素，看例子

func main(){
	var ss=make([]string,10);
	ss=append(ss,"last");
	print("after append",ss,"\n")

	var ss2=make([]int,10);
	print("before append",ss2,"\n")
	ss2=append(ss2,1);
	print("after append",ss2,"\n")
	for i,e:=range ss2{
		fmt.Println("i:",i,"e:",e)
	}

}
//---
//Running...
//
//[         after append ]    :    length:11    addr:0xc20804c000    isnil:false

//实际上，此时我们应该先用下标为切片元素负值。但是如果我们既想有好的效率，有想继续使用append函数而不想区分是否有空的元素，此时就要请出make的第三个参数，容量，也就是我们通过传递给make，0的大小和足够大的容量数值就行了。
//var ss=make([]string,0,10);
//ss=append(ss,"last");
//print("after append",ss)
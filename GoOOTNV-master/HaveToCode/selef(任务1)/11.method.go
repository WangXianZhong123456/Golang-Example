package main

import "fmt"

type A struct {
	Name string
	age  int // 如果定义的常量、变量、类型、接口、结构、函数等的名称是大写字母开头表示能被其它包访问或调用（相当于public），
	//非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）
}

type B struct {
	Name string
}

type TZ int

func main() {
	a := A{Name: "haha", age: 1} //值类型只改变区域内的值,不会全局改变
	a.Println()
	fmt.Println(a)

	b := B{Name: "meiyou"} //引用类型地址和值都会改变,会全局改变
	b.Print()
	fmt.Println(b)
	/*y := test(111)
	fmt.Println(y)*/
	var tz TZ = 10
	tz.Print()
	(*TZ).Print(&tz) //第二种调用方法
	fmt.Println(tz)
}

func (a A) Println() {
	a.Name = "AA"
	fmt.Println(a)
}

func (b *B) Print() {
	b.Name = "nimei"
	fmt.Println(*b)
}

func (a *TZ) Print() {
	fmt.Println("TZ")
}

/*func test(a int) int {
	return a
}*/

package main

import "fmt"

func main() {
	a := 10
	b := 20
	X(a, b, 50, 60) // 值传递
	fmt.Println("值传递 = ", a, b)

	s := []int{10, 20}
	Y(s) // 引用传递
	fmt.Println("引用传递 = ", s)

	Z(&a) // 地址传递
	fmt.Println("地址传递 = ", a)

	ff := func() {
		fmt.Println("匿名函数")
	}
	ff()

	//闭包函数
	fff := closure(100)
	fmt.Println("closure+1 = ", fff(1))
	fmt.Println("closure+2 = ", fff(2))

	fmt.Println("a")
	/*1、作为异常处理和文件关闭的处理函数。
	2、匿名函数和命名函数都可以作为defer的延迟函数。
	3、多个defer语句之间的顺序是先进后出的。
	4、defer语句可以有返回值，但是这个返回值是没有意义的。
	5、匿名函数的操作对象如果会被返回，则该defer是会对返回值有影响的。*/
	defer fmt.Println("defer=b") // 先进后出
	defer fmt.Println("defer=c")
	defer fmt.Println("defer=d")

	for i := 0; i < 3; i++ {
		defer fmt.Println("for=", i)
	}

	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("for defer=", i*10) //注意: 打印值不变,都是30
		}()
	}

	C()
	panicFunc()
	C()

}

// 函数返回值写法
// 无返回值: func A() {...}
// 单返回值: func A() int {...}
// 多返回值1(完全体): func A() (x int, y int, z int){...}
// 多返回值2(同类型简写): func A() (x,y,z int){...} --> 函数体自带 x,y,z变量, return可简写成return
// 多返回值2(不写名称只写类型): func A() (int, int, int){...} --> 函数体必须指定return变量名(return a,b,c)
// 不定长变参: func A(a ...int){...} --> 变参必须写在最后, 变参是值拷贝, 不会修改原来的值
func X(s ...int) {
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
}

// 引用
func Y(s []int) {
	s[0] = 1
	s[1] = 2
	fmt.Println(s)
}

// 指针参数
func Z(a *int) {
	*a = 2
	fmt.Println(*a)
}

// 闭包函数
func closure(x int) func(int) int {
	fmt.Printf("closure地址=%p\n", &x)
	return func(y int) int {
		fmt.Printf("closure地址=%p\n", &y)
		return x + y
	}
}

func panicFunc() {
	// 注册defer函数
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

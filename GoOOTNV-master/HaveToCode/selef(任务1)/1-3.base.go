//当前程序包
package main

// 导入其它包
import "fmt"
import "math"

// 全局变量不能省略var关键字
var PI = 3.14

// 全局编程声明与赋值
var name = "gohlep"

// 一般类型声明
type newType int

// 结构体声明
type gopher struct{}

// 接口声明
type golang interface{}

type byte int8
type rune int32
type 文本 string

func main() {
	var a [1]bool
	var b 文本
	b = "中文类型名"
	var c int = 1
	d := false
	var x, y, z = 1, 2, 3
	xx, _, zz := 3, 2, 1
	var aa float32 = 100.1
	fmt.Println("Hello Go")
	bb := int(aa) // 强制转换
	a[0] = true
	fmt.Println("bb=", bb)
	fmt.Println("a=", a[0])
	fmt.Println("math.MinInt8=", math.MinInt8)
	fmt.Println("math.MaxInt32=", math.MaxInt32)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(xx)
	fmt.Println(zz)

}

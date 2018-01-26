package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Person2 struct {
	Name    string
	Age     int
	Contact struct {
		Phone   string
		Address string
	}
}
type Person3 struct {
	string
	int
}
type human struct {
	Sex int
}
type teacher struct {
	Name string
	Age  int
	Sex  string
	human
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := Person{
		Name: "wxz",
		Age:  27}
	fmt.Println(a)

	A(a)
	fmt.Println(a)
	// 引用传递
	B(&a)
	fmt.Println(a)

	C(&a)
	fmt.Println(a)
	fmt.Println("--------------")

	// 直接使用地址
	b := &Person{
		Name: "joe",
		Age:  19}
	b.Name = "john" // 即使b是地址, 也是使用.符号进行操作
	fmt.Println(b)
	B(b)
	fmt.Println(b)
	C(b)
	fmt.Println(b)
	fmt.Println("--------------")

	p2 := Person2{Name: "王贤忠", Age: 27}
	p2.Contact.Phone = "123456"
	p2.Contact.Address = "湖南郴州市宜章县"
	fmt.Println(p2)

	fmt.Println("**************")
	p3 := Person3{"没有属性的结构", 1}
	fmt.Println(p3)

	// 嵌入结构
	tea := teacher{
		Name:  "ttt",
		Age:   30,
		human: human{Sex: 1},
	}

	fmt.Println(tea)

	stu := student{
		Name:  "I am Student",
		Age:   27,
		human: human{Sex: 100},
	}
	fmt.Println(stu) /*输出的顺序默认是结构体定义的顺序*/
}

func A(per Person) {
	per.Age = 17
	fmt.Println("A", per)
}

func B(per *Person) {
	per.Age = 13
	fmt.Println("A", per)
}

func C(per *Person) {
	per.Age = 15
	fmt.Println("C", per)
}

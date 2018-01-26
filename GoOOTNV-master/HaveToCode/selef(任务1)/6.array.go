package main

import "fmt"

func main() {
	a := [20]int{1: 100, 19: 1} //{}里的冒号左边是下标，右边是值
	fmt.Println(a)
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	c := [...]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Println(c)
	d := [...]int{29: 1}
	fmt.Println(d)
	var p *[30]int = &d
	fmt.Println(p)
	x, y := 1, 2
	q := [...]*int{&x, &y}
	fmt.Println(q)
	//fmt.Println(*q[0], *q[1])

	r := [...]*[5]int{&b, &c} // b和c数组都是相同类型的数组(长度一样)
	fmt.Println("r=", r)

	aa := new([10]int) // new出来的数组，返回的是数组指针，先不用分配地址 相当于 var aa *[20]int = &a
	aa[1] = 2
	fmt.Println("aa = ", *aa)

	bb := [10]int{}
	bb[1] = 2
	fmt.Println("bb = ", bb)

	aaa := [2][3]int{
		{1, 2, 3}, // 注意,最后一个}不能单独占一行
		{4, 5, 6}}
	fmt.Println("aaa = ", aaa)

	bbb := [2][3]int{
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println("bbb = ", bbb)

	ccc := [...][3]int{ // NOTE: 第二维不能用...省略
		{1: 1},
		{2: 2}} // 注意,最后一个}不能单独占一行
	fmt.Println("ccc = ", ccc)

	ppp := [...]int{5, 2, 6, 7, 8}
	fmt.Println("ppp = ", ppp)
	cnt := len(ppp)
	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if ppp[i] > ppp[j] {
				temp := ppp[i]
				ppp[i] = ppp[j]
				ppp[j] = temp
			}
		}
	}
	fmt.Println("ppp sort = ", ppp)

}

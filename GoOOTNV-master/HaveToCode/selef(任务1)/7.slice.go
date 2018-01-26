package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		v := 1
		fmt.Println(&v)
	}

	// 使用数组下表切片
	arr := [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr = ", arr)
	fmt.Println(arr[5:])
	fmt.Println(arr[:5])

	// 使用make创建切片
	s := make([]int, 3, 10) // cap()函数返回的是数组切片分配的空间大小
	fmt.Println(s, len(s), cap(s))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	sb := a[3:5]
	fmt.Println("sa = ", string(sa))
	fmt.Println("sb = ", string(sb))
	fmt.Println("sa[3:5] = ", string(sa[3:5]))
	fmt.Println("sa[3:9] = ", string(sa[3:9]))

	ss := make([]int, 3, 6)
	fmt.Println("%v %p---\n", ss, ss)
	ss = append(ss, 1, 2, 3) // 只有切片才有append方法
	fmt.Printf("%v %p\n", ss, ss)
	ss = append(ss, 1, 2, 3) // 超过容量最大值,重新分配了内存
	fmt.Printf("%v %p\n", ss, ss)

	aaa := []int{1, 2, 3, 4, 5, 6}
	sss1 := aaa[1:3]
	sss2 := aaa[2:5]
	fmt.Println("aaa[1:3] = ", sss1, "aaa[2:5] = ", sss2)
	fmt.Printf("%p--- %p---\n", sss1, sss2)
	sss2 = append(sss2, 1) // 在改变sss1之前对sss2增加元素, 但是没有超出容量, 不重新分配内存
	fmt.Printf("%p--- %p---\n", sss1, sss2)
	//sss1[1] = 9 // 可能改变共同部分(sss2没有重新分配的话), 也可能不改变共同部分(如果sss2重新分配了就没有共同部分了)
	fmt.Println(sss1, sss2)

	// 切片拷贝,以2个切片中短的切片为准, 从左到右尽可能拷贝最多个元素, 直到其中一个切片结束,丢弃多余的部分
	x := []int{1, 2, 3, 4, 5, 6}
	y := []int{7, 8, 9, 10}
	// copy(y, x)
	copy(y, x[1:])
	fmt.Println(y)

	// z := x[0:len(x)]
	// z := x[0:]
	z := x[:]
	fmt.Println(z)
}

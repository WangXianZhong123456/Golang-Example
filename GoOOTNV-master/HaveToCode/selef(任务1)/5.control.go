package main

import "fmt"

func main() {
	/*fmt.Println("Hello Go")*/
	a := 1
	//	a += 1
	var p *int = &a
	fmt.Println("*p=", (*p))

	a++
	a = 10
	if a, b := 1, 2; a > 0 {
		fmt.Println("a=", a, "b=", b)
	}
	fmt.Println("a=", a)

	a = 1
	// 类似无限循环
	for {
		a++
		if a > 3 {
			break
		}
		fmt.Println("in for,a=", a)
	}
	fmt.Println("Over")

	a = 1
	// 类似while循环
	for a <= 3 {
		a++
		fmt.Println("in for,a=", a)

	}
	fmt.Println("Over")

	a = 1
	for i := 0; i < 3; i++ {
		a++
		fmt.Println("in for,a=", a)
	}
	fmt.Println("Over")

	fmt.Println("Switch 1")

	a = 2
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	default:
		fmt.Println("None")
	}

	fmt.Println("Switch 2")
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}

	fmt.Println("Switch 3")
	/*只要有匹配到默认跳出break,如果想继续往下执行加fallthrough*/
	switch s := 1; {
	case s >= 0:
		fmt.Println("s=0")
		fallthrough // case执行中断后，如果需要继续执行下一个case块的内容，在下一个case块结尾执行fallthrough并且可以在fallthrough前使用break语句阻止。但不继续继续后续case块。
	case s >= 1:
		fmt.Println("s=1")
		fallthrough
	case s >= 2:
		fmt.Println("s=2")
		fallthrough
	default:
		fmt.Println("None")

	}

BREAK_LABEL:
	for {
		for i := 1; i < 10; i++ {
			if i > 3 {
				break BREAK_LABEL
			}
		}
	}
	fmt.Println("break to BREAK_LABEL")

CONTINUE_LABEL:
	for i := 0; i < 10; i++ {
		for {
			continue CONTINUE_LABEL
			fmt.Println("i=", i)

		}
	}
	fmt.Println("continue to CONTINUE_LABEL")

GOTO_LABEL:
	for i := 0; i < 10; i++ {
		for {
			goto GOTO_LABEL
			fmt.Println("i=", i)
		}
	}
	fmt.Println("goto to GOTO_LABEL")
}

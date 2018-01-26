// copy的使用
//在使用copy复制切片之前，要保证目标切片有足够的大小，注意是大小，而不是容量，还是看例子
package main

import "fmt"

func main() {
	var sa = make ([]string,0);
	for i:=0;i<10;i++{
		sa=append(sa,fmt.Sprintf("%v",i))

	}
	var da =make([]string,0,10);
	var cc=0;
	cc= copy(da,sa);
	fmt.Printf("copy to da(len=%d)\t%v\n",len(da),da)
	da = make([]string,5)
	cc=copy(da,sa);
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n",len(da),cc,da)
	da = make([]string,10)
	cc =copy(da,sa);
	fmt.Printf("copy to da(len=%d)\tcopied=%d\t%v\n",len(da),cc,da)

}

//---
//Running...
//
//copy to da(len=0)    []
//copy to da(len=5)    copied=5    [0 1 2 3 4]
//copy to da(len=10)    copied=10    [0 1 2 3 4 5 6 7 8 9]
//从上面运行结果，明显看出，目标切片大小0，容量10，copy不能复制。目标切片大小小于源切片大小，copy就按照目标切片大小复制，不会报错。


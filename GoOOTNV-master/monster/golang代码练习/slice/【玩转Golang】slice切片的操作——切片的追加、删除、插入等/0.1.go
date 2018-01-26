//0.1append操作亦然会在需要的时候构造新的切片，不过是将地址都保存到了sa中，因此我们通过该指针始终可以访问到真正的数据。
package main

import "fmt"

func main() {
	var osa = make([]string, 0)
	sa := &osa
	for i := 0; i < 10; i++ {
		*sa = append(*sa, fmt.Sprintf("%v", i))
		fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, sa, sa)
	}
	fmt.Printf("addr of osa:%p,\taddr:%p \t content:%v\n", osa, sa, sa)

}

//---
//Running...
//
//addr of osa:0xc20800a220,    addr:0xc20801e020      content:&[0]
//addr of osa:0xc20801e0a0,    addr:0xc20801e020      content:&[0 1]
//addr of osa:0xc20803e0c0,    addr:0xc20801e020      content:&[0 1 2]
//addr of osa:0xc20803e0c0,    addr:0xc20801e020      content:&[0 1 2 3]
//addr of osa:0xc208050080,    addr:0xc20801e020      content:&[0 1 2 3 4]
//addr of osa:0xc208050080,    addr:0xc20801e020      content:&[0 1 2 3 4 5]
//addr of osa:0xc208050080,    addr:0xc20801e020      content:&[0 1 2 3 4 5 6]
//addr of osa:0xc208050080,    addr:0xc20801e020      content:&[0 1 2 3 4 5 6 7]
//addr of osa:0xc208052000,    addr:0xc20801e020      content:&[0 1 2 3 4 5 6 7 8]
//addr of osa:0xc208052000,    addr:0xc20801e020      content:&[0 1 2 3 4 5 6 7 8 9]
//addr of osa:0xc208052000,    addr:0xc20801e020      content:&[0 1 2 3 4 5 6 7 8 9]

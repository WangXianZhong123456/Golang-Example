package bubblesort

func BubbleSort(values []int) {
	flag := true
	// 冒泡排序，只优化了外层排序
	// 具体可参考http://blog.csdn.net/yanxiaolx/article/details/51622286
	for i := 0; i < len(values); i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}

		}
		if flag == true {
			break
		}
	}
}

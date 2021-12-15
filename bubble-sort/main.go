package main

import (
	"fmt"
)

/*
1.确定下标返回[0,n-1]，[0,n-2]
2.相邻位置交换
*/
func main() {

	// 随机一个0-50的随机数组
	ints := []int{4, 2, 3, 1, 2, 3, 4, 5}
	//ints := []int{}
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 100; i++ {
	//	r := rand.Intn(50)
	//	ints = append(ints, r)
	//}

	// 冒泡排序
	// [0,n-1] 范围 ints[i]>ints[i+1] 则交换，保证最后一位是最大的
	// [0,n-2] 范围 ints[i]>ints[i+1] 则交换，保证最后一位是最大的
	// [0,n-3] 范围 ints[i]>ints[i+1] 则交换，保证最后一位是最大的
	// [0,n-4] 范围 ints[i]>ints[i+1] 则交换，保证最后一位是最大的
	// [0,n-5] 范围 ints[i]>ints[i+1] 则交换，保证最后一位是最大的
	length := len(ints)
	times := 0
	for j := length - 1; j > 0; j-- {
		times++
		for i := 0; i < j; i++ {
			// 0,1
			// 1,2
			// j-1,j
			if ints[i] > ints[i+1] {
				ints[i], ints[i+1] = ints[i+1], ints[i]
			}
		}
	}
	fmt.Println(fmt.Sprintf("times:%v", times))
	fmt.Println(fmt.Sprintf("%v", ints))

	// 优化
	// 对于本身有序，或者部分有序的情况，
	// 前一次排序如果没有发生交换，则说明已经是有序数组，则不需要进行下一次的冒泡排序了
	// 4,2,3,1,2,3,4,5
	// 第一次排序后：[2 3 1 2 3 4 4 5]
	// 第二次排序后：[2 1 2 3 3 4 4 5]
	// 第三次排序后：[2 1 2 3 3 4 4 5]
	// 第四次排序后：[1 2 2 3 3 4 4 5] 下次排序则不会发生交换，说明已经是有序的了

	ints1 := []int{4, 2, 3, 1, 2, 3, 4, 5}
	times1 := 0
	for j := 0; j < length; j++ {
		times1++
		swapFlag := false
		for i := 1; i < length-j; i++ {
			// 0,1
			// 1,2
			// j-1,j
			if ints1[i-1] > ints1[i] {
				ints1[i], ints1[i-1] = ints1[i-1], ints1[i]
				swapFlag = true
			}
		}
		if !swapFlag {
			break
		}

	}
	// 相比优化前少执行3次
	fmt.Println(fmt.Sprintf("times1:%v", times1))

}

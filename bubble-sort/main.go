package main

import (
	"fmt"
	"math/rand"


)

func main() {

	// 随机一个100-1000的随机数组
	ints := []int{}
	for i := 0; i < 100; i++ {
		r := rand.Intn(900) + 100
		ints = append(ints, r)
	}
	fmt.Println(fmt.Sprintf("%v", ints))

	// 冒泡排序
	length := len(ints)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	fmt.Println(fmt.Sprintf("%v", ints))










}

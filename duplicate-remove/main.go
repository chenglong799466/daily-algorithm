package main

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cast"
)

func main() {


	// 随机一个0-1000的随机数组
	var ints []int
	for i := 0; i < 100; i++ {
		r := rand.Intn(10)
		ints = append(ints, r)
	}
	fmt.Println(fmt.Sprintf("%v", ints))

	// 利用mapkey唯一，分组
	group(ints)

	//
	unique(ints)

}

func group(ints []int) map[string][]int {
	m := make(map[string][]int, 0)
	for _, item := range ints {
		value, ok := m[cast.ToString(item)]
		if !ok {
			m[cast.ToString(item)] = value
		}
		value = append(value, item)
		m[cast.ToString(item)] = value
	}
	fmt.Println(fmt.Sprintf("%v", m))
	return m
}

func unique(ints []int) []int {
	//
	var temps []int
	for i := range ints {
		repeat := false
		for j := i + 1; j < len(ints); j++ {
			if ints[i] == ints[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			temps = append(temps, ints[i])
		}
	}
	fmt.Println(fmt.Sprintf("%v", temps))
	return temps
}
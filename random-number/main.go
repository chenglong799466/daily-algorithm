package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 1.不加种子，不加范围。
	// 生成的随机数是固定数，多次运行结果是一样的。
	for i := 0; i < 5; i++ {
		r := rand.Int()
		fmt.Printf("%d,", r)
	}
	fmt.Println("end")


	// 2.仅指定范围，种子数未设定
	// 生成的随机数也是固定数，多次运行结果也是一样的。
	for i := 0; i < 5; i++ {
		r := rand.Intn(100)
		fmt.Printf("%d,", r)

	}
	fmt.Println("end")

	// 3.指定范围，种子数固定
	// 生成的随机数也是固定数，多次运行结果也是一样的。
	rand.Seed(100)
	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d,", r)
	}
	fmt.Println("end")

	// 4.指定范围，种子数不固定
	// 生成的随机数不固定，多次运行结果不同
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r := rand.Intn(10)
		fmt.Printf("%d,", r)
	}
	fmt.Println("end")

	// 4.指定范围100-1000随机数
	// 在原来的（0-900）随机数上加100
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r := rand.Intn(900) + 100
		fmt.Printf("%d,", r)
	}
	fmt.Println("end")

}

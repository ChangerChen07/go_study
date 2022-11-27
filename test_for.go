package main

import "fmt"

func f1() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 省略初始条件
func f2() {
	i := 1
	for ; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

// 省略初始条件和结束条件
func f3() {
	i := 1
	for i <= 10 {
		fmt.Printf("i: %v\n", i)
		i++
	}
}

// 永真循环
func f4() {
	for {
		fmt.Println("永远执行")
	}
}
func main() {
	f3()
}

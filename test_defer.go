package main

import "fmt"

func f1() {
	// 延迟执行,先进后出
	fmt.Println("start")
	defer fmt.Println("step1")
	defer fmt.Println("step2")
	defer fmt.Println("step3")
	fmt.Println("end")
}
func main() {
	f1()
}

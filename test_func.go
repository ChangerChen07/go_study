/* package main

import "fmt"

func sum(a int, b int) (r int) {
	r = a + b
	return r
}

func f1(s []int, a int) {
	// 若是指针类型的参数，则会影响函数外部的变量值
	s[0] = 1000
	a = 100
}

func f2(a ...int) {
	// 可变参数
	for _, v := range a {
		fmt.Printf("v: %v\n", v)
	}
}

func f3(name string, ok bool, args ...int) {
	// 前面固定，后面可变参数
	fmt.Printf("name: %v\n", name)
	fmt.Printf("ok: %v\n", ok)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {
	// r := sum(10, 20)
	// fmt.Printf("r: %v\n", r)
	s := []int{1, 2, 3, 4}
	a := 9
	f1(s, a)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("a: %v\n", a)

	f2(1, 2, 3, 3, 4, 41, 243, 2)
}
*/
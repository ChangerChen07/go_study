/* // 闭包=函数+引用环境
package main

import (
	"fmt"
	"strings"
)

func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// f函数周期内，x的值不会变
	f := add()
	r := f(10)
	fmt.Printf("r: %v\n", r)
	r = f(20)
	fmt.Printf("r: %v\n", r)
	fmt.Println("----------------")
	f1 := add()
	r1 := f1(100)
	fmt.Printf("r1: %v\n", r1)
	r = f(30)
	fmt.Printf("r: %v\n", r)

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
	f1, f2 := calc(100)
	fmt.Println(f1(1), f2(10))
	fmt.Println(f1(2), f2(20))
}
*/
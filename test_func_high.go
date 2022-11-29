/* package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func test(name string, f func(string)) {
	// 函数作为另一个函数的参数
	f(name)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func cal(oprator string) func(int, int) int {
	// 函数作为另一个函数的返回值
	switch oprator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	test("Changer", sayHello)

	f := cal("-")
	r := f(10, 100)
	fmt.Printf("r: %v\n", r)
	r = cal("+")(100, 1000)
	fmt.Printf("r: %v\n", r)
}
*/
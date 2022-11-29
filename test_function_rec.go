/* package main

import "fmt"

func f1(n int) int {
	// 阶乘
	var r int = 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}

func f2(n int) int {
	// 递归实现阶乘
	if n == 1 {
		// 返回条件
		return n
	} else {
		// 自己调用自己
		return n * f2(n-1)
	}
}

func f3(n int) int {
	// 斐波那契数列 f(n)=f(n-2)+f(n-1)，f(2)=f(1)=1
	if n == 1 || n == 2 {
		return 1
	} else {
		return f3(n-2) + f3(n-1)
	}
}

func main() {
	i := f3(15)
	fmt.Printf("i: %v\n", i)
}
*/
/* package main

import "fmt"

func f1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

func f2() {
	// 单独break没有意义，结合fallthrough使用，在fallthrough前break
	i := 1
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}

func f3() {
	// break到标签
MYLABEL:
	for i := 0; i <= 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break MYLABEL
		}
	}
	fmt.Println("结束。。。。。。。")
}

func f4() {
MYLABEL:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v: %v\n", i, j)
			if i >= 2 && j >= 2 {
				break MYLABEL
			}
		}
	}
}
func main() {
	f4()
}
*/
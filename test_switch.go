/* package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("周末")
	default:
		fmt.Println("非法输入")
	}

	score := 0
	switch {
	case score < 60:
		fmt.Println("不及格")
	case score >= 80 && score <= 90:
		fmt.Println("优秀")
	default:
		fmt.Println("其他")
	}

	num := 100
	switch num {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
		// fallthrough
	case 300:
		fmt.Println("300")
	}
}
*/
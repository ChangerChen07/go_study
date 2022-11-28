/* package main

import "fmt"

func f1() {
	i := 1
	if i >= 2 {
		fmt.Printf("i: %v\n", i)
	} else {
		goto END
	}
END:
	fmt.Println("结束")
}

func f2() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%v: %v\n", i, j)
			if i == 2 && j == 2 {
				goto END
			}
		}
	}
END:
	fmt.Println("结束for")
}
func main() {
	f2()
}
*/
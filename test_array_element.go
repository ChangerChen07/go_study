/* package main

import "fmt"

func f1() {
	var a1 = [4]int{321, 3231, 12, 3}
	// var a2 = [3]string{"sfds", "adf", "sfdsa"}
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[2]: %v\n", a1[2])
	fmt.Println("---------" + "修改后" + "---------")
	a1[0] = 10
	a1[2] = 11
	fmt.Printf("a1[0]: %v\n", a1[0])
	fmt.Printf("a1[2]: %v\n", a1[2])
}

func f2() {
	// for range
	var v1 = [3]int{1, 10, 100}
	for i, v := range v1 {
		fmt.Printf("v1[%v]: %v\n", i, v)
	}

	for i := 0; i < len(v1); i++ {
		fmt.Printf("v1[%v]: %v\n", i, v1[i])
	}
}

func main() {
	f2()
}
*/
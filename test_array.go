/* package main

import "fmt"

func f1() {
	var a1 = [...]string{"Sf", "saf", "sfs"}
	var a2 [4]int
	a1 = [3]string{"sf", "saf", "sf00"}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a2: %T\n", a2)
}

func f2() {
	// 指定index值
	var a1 = [3]string{2: "sfasf"}
	var a2 = [4]int{2: 32, 1: 323}
	var a3 = [...]int{1: 32, 10: 1323, 4: 231}
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a3: %v\n", len(a3))
	a3[6] = 99
	fmt.Printf("a3: %v\n", a3)
}
func main() {
	f2()
}
*/
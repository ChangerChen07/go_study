package main

import "fmt"

// add
func f1() {
	var s1 []int
	s1 = append(s1, 123)
	s1 = append(s1, 99)
	s1 = append(s1, 123)
	fmt.Printf("s1: %v\n", s1)
}

// del
func f2() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append(s1[:2], s1[3:]...)
	fmt.Printf("s2: %v\n", s2)
}

// copy
func f3() {
	s1 := []int{12, 23, 123, 32}
	var s2 = []int{0, 0, 0, 0, 0} //要赋予大于等于的长度才能copy
	s3 := make([]int, 5)          //长度要大于等于原切片才能copy
	copy(s2, s1)
	copy(s3, s1)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	s2[1] = 1
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)
	s3[3] = 10
	fmt.Printf("s3: %v\n", s3)
}
func main() {
	f3()
}

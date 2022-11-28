/* package main

import "fmt"

func f1() {
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s1 type: %T\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s2 tyoe: %T\n", s2)
	fmt.Println("没有打印")
}

func f2() {
	var s1 = make([]int, 2)
	fmt.Printf("s1: %v\n", s1)
	// s1 = append(s1, 10)
	// fmt.Printf("s1: %v\n", s1)
}

func f3() {
	var s1 = []int{2, 31, 12}
	fmt.Printf("len(s1): %v\n", len(s1)) //长度
	fmt.Printf("cap(s1): %v\n", cap(s1)) //容量
	s1 = append(s1, 123)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("len(s1): %v\n", len(s1)) //长度
	fmt.Printf("cap(s1): %v\n", cap(s1)) //容量
}

func f4() {
	// 前闭后合
	s1 := [4]int{1, 2, 3, 4}
	s2 := s1[:3]
	s3 := s1[:]
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s2 type: %T\n", s2)
	s2 = append(s2, 123, 12312, 123123)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("len(s2): %v\n", len(s2))
	fmt.Printf("s3: %v\n", s3)
}
func main() {
	f4()
}
*/
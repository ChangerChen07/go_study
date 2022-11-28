/* package main

import "fmt"

// 数组
func f1() {
	x := [5]int{2, 3, 4, 213, 3123}
	for i, v := range x {
		fmt.Printf("index %v=%v\n", i, v)
	}
}

// 切片
func f2() {
	x := []string{"#21", "#213", "4324", "eqwe"}
	for _, v := range x {
		fmt.Printf("v: %v\n", v)
	}
}

// map
func f3() {
	m1 := make(map[string]string, 0)
	m1["name"] = "Changer"
	m1["age"] = "36"
	m1["gender"] = "fmale"
	m2 := make(map[string]string, 0)
	m2["name"] = "Sally"
	m2["age"] = "35"
	m2["gender"] = "male"
	p := make(map[string]map[string]string, 0)
	p["person1"] = m1
	p["person2"] = m2
	for _, v := range p {
		fmt.Printf("v: %v\n", v)
		for key, val := range v {
			fmt.Printf("%v: %v\n", key, val)
		}
	}

	// var m map[string]int
	// 直接初始化
	m := map[string]int{"age": 60, "hight": 187, "weight": 77}

	// make初始化
	// m = make(map[string]int)
	// m["age"] = 30
	// m["hight"] = 178
	// m["weight"] = 73
	for key, v := range m {
		fmt.Printf("%v: %v\n", key, v)
	}

}

// 字符串
func f4() {
	s := "Hello World!"
	for _, v := range s {
		fmt.Printf("v: %c\n", v)
	}
}
func main() {
	f4()
}
*/
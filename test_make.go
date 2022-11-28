/* package main

import "fmt"

func f1() {
	var m1 = map[string]string{
		"name":   "Changer",
		"age":    "36",
		"gender": "fmale",
	}
	fmt.Printf("m1: %v\n", m1)
	var m2 map[string]int
	// m2 = make(map[string]int)
	fmt.Printf("m2: %v\n", m2)
	// m2["grade"] = 1
	// m2["level"] = 3
	// fmt.Printf("m2: %v\n", m2)
}

func f2() {
	// 判断key值是否存在
	var m1 = map[string]string{"name": "Changer", "age": "36", "gender": "fmale"}
	v, ok := m1["name"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	fmt.Println("------------------")
	v, ok = m1["email"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
}

func f3() {
	var m1 = map[string]int{"1": 1, "2": 2, "3": 3}
	for k := range m1 {
		fmt.Printf("k: %v\n", k)
	}

	for k, v := range m1 {
		fmt.Printf("%v: %v\n", k, v)
	}
}
func main() {
	f3()
}
*/
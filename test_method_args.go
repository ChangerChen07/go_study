/* package main

import "fmt"

type Person struct {
	id   int
	name string
}

func f1() {
	p1 := Person{
		0,
		"Changer",
	}
	p2 := &Person{
		1,
		"Sally",
	}
	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("p1 Type: %T\n", p1)
	fmt.Printf("p2: %v\n", p2)
	fmt.Printf("p2 Type: %T\n", p2)
}

// 值类型复制副本
func (per Person) showPerson1() {
	per.name = "Tom1"
	// fmt.Printf("per: %v\n", per.name)
}

func (per *Person) showPerson2() {
	per.name = "Tom2"
}
func main() {
	p1 := Person{
		0,
		"Tom",
	}
	p2 := &Person{
		1,
		"Tom",
	}
	p1.showPerson1()
	fmt.Printf("p1: %v\n", p1)
	p2.showPerson2()
	fmt.Printf("p2: %v\n", *p2)
}
*/
/* package main

import (
	"fmt"
)

func f1() {
	var s string
	s = "Changer"
	var p *string
	p = &s
	fmt.Printf("s: %v\n", s)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("*p: %v\n", *p)
}

func f2() {
	type Person struct {
		id, age int
		name    string
	}
	var tom = Person{
		0,
		20,
		"Tom",
	}
	var p_person *Person
	p_person = &tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("*p_person: %v\n", *p_person)
	fmt.Printf("p_person.age: %v\n", p_person.age)
	fmt.Printf("p_person.name: %v\n", p_person.name)

	var pname *string = &tom.name
	var pname_p *string = &p_person.name
	fmt.Printf("pname: %v\n", pname)
	fmt.Printf("pname_p: %v\n", pname_p)

	p_person.name = "Changer"
	fmt.Printf("p_person: %v\n", p_person)
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("pname: %v\n", pname)
	fmt.Printf("pname_p: %v\n", pname_p)
}

func f3() {
	// new关键字定义
	type Person struct {
		id, age int
		name    string
	}
	var tom = new(Person)
	tom.id = 0
	tom.age = 20
	tom.name = "Tom"
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom: %p\n", tom)
	fmt.Printf("tom type: %T\n", tom)
}
func main() {
	f3()
}
*/
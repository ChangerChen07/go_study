package main

import "fmt"

type Person struct {
	id, age int
	name    string
}

func showNewPerson(per *Person) {
	per.id = 2
	per.age = 30
	per.name = "Changer"
	fmt.Printf("per: %v\n", per)
}

func showPerson(per Person) {
	// 只拷贝副本
	per.id = 2
	per.age = 30
	per.name = "Changer"
	fmt.Printf("per: %v\n", per)
}
func main() {
	var tom = Person{
		0,
		20,
		"Tom",
	}
	fmt.Printf("tom: %v\n", tom)

	showPerson(tom)
	fmt.Printf("tom: %v\n", tom)

	showNewPerson(&tom)
	fmt.Printf("tom: %v\n", tom)
}

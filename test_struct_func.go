/* package main

import "fmt"

type Person struct {
	id, age int
	name    string
}

func (per Person) eat() {
	fmt.Printf("%v eat.....\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("%v sleep.........\n", per.name)
}

type Customer struct {
	id   int
	name string
}

func (customer Customer) login(name string, passwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name != "Hellen" && passwd != "123" {
		return false
	}
	return true
}
func main() {
	var tom = Person{}
	tom.name = "Tom"
	tom.eat()
	tom.sleep()

	var he = Customer{}
	he.id = 0
	he.name = "Hellen"
	b := he.login("Hellen", "123")
	fmt.Printf("b: %v\n", b)
}
*/
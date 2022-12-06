/* package main

import "fmt"

//OCP,open-close ，开-闭设计原则。开放扩展，关闭修改
type Pet interface {
	eat()
	sleep()
}

type Dog struct {
	// 首字母小写是私有，其他包不可访问；若首字母为大写，为公有，其他包可以访问
	name string
	age  int
}

func (d *Dog) eat() {
	fmt.Println("dog eat....")
}

func (d *Dog) sleep() {
	fmt.Println("dog sleep....")
}

type Cat struct {
	name string
	age  int
}

func (c *Cat) eat() {
	fmt.Println("cat eat....")
}

func (c *Cat) sleep() {
	fmt.Println("cat sleep....")
}

type Person struct {
	name string
}

func (p *Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}
func main() {
	tom := Person{name: "Tom"}
	dahuang := &Dog{name: "大黄",
		age: 3,
	}
	huahua := &Cat{
		name: "花花",
		age:  2,
	}
	tom.care(dahuang)
	tom.care(huahua)
}
*/
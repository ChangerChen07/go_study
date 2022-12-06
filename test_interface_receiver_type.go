/* package main

import "fmt"

type Dog struct {
	name string
	age  int
}

type Cat struct {
	name string
	age  int
}

func (d *Dog) eat(w string) string {
	d.name = "宠物" + d.name
	fmt.Printf("d.name: %v 吃 %v\n", d.name, w)
	return "宠物狗(" + d.name + ")吃东西。。。"
}

func (c *Cat) eat(w string) string {
	c.name = "宠物" + c.name
	fmt.Printf("d.name: %v 吃 %v\n", c.name, w)
	return "猫猫(" + c.name + ")吃东西。。。"
}

type PET interface {
	eat(string) string
}

func main() {
	luni := &Dog{
		name: "鲁尼",
		age:  3,
	}
	huaer := &Cat{
		name: "花花",
		age:  5,
	}

	r := luni.eat("骨头")
	fmt.Printf("r: %v\n", r)

	r = huaer.eat("鱼")
	fmt.Printf("r: %v\n", r)

	pets := make([]PET, 2)
	pets[0] = luni
	pets[1] = huaer

	fmt.Printf("pets: %v\n", pets)
}
*/
package main

import "fmt"

// 通过嵌套的方式实现继承
type Animal struct {
	weight, age, hight int
	where              string
}

func (a Animal) eat() {
	fmt.Println("吃东西")
}
func (a Animal) sleep() {
	fmt.Println("睡觉")
}

type Dog struct {
	Animal
	name string
}

type Cat struct {
	Animal
	name string
}

func main() {
	dahuang := Dog{
		Animal{20,
			3,
			50,
			"中国"},
		"大黄",
	}
	dahuang.eat()
	dahuang.sleep()
	fmt.Printf("dahuang.name: %v\n", dahuang.name)

	huahua := Cat{
		Animal{weight: 10, age: 5, hight: 10, where: "法国"},
		"花花",
	}
	fmt.Printf("huahua: %v\n", huahua)
}

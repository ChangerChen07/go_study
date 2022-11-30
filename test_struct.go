/* package main

import "fmt"

type Person struct {
	id    int
	name  string
	age   int
	email string
}

type Work struct {
	// 相同类型可以放在同一行定义
	id, priod, acount int
	name, mode        string
	who               Person
}

func main() {
	var tom Person
	fmt.Printf("tom: %v\n", tom)
	tom.id = 0
	tom.name = "Tom"
	tom.age = 36
	tom.email = "tom@qq.com"
	fmt.Printf("tom:%T, %v\n", tom, tom)

	var w1 Work
	w1.who = tom
	fmt.Printf("w1: %v\n", w1)

	// 匿名结构体
	var me struct {
		id, age     int
		name, email string
	}
	me.id = 1
	me.age = 36
	me.name = "Changer"
	me.email = "Changerqi@163.com"
	fmt.Printf("me: %v\n", me)
	fmt.Printf("me.email: %v\n", me.email)

}
*/
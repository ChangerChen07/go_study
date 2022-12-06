package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person, error) {
	if name == "" {
		return nil, fmt.Errorf("姓名不能为空")
	}
	if age <= 0 {
		return nil, fmt.Errorf("年龄不能小于0")
	}
	return &Person{name: name, age: age}, nil
}
func main() {
	tom, err := NewPerson("Tom", 0)
	if err == nil {
		fmt.Printf("tom.name: %v\n", tom.name)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

package user

import "fmt"

func Hello() string {
	return "Hello"
}

// 首字母小写包内可见
func sayhi() {
	fmt.Println("Say Hi!")
}

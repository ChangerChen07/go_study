/* package main

import "fmt"

func main() {
	var (
		name  string
		age   int
		email string
	)
	fmt.Println("请输入name,age,email，用空格分隔")
	fmt.Scan(&name, &age, &email)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
	// 此方法定义的变量只作用在if里面的区域
	if age := 10; age >= 18 {
		fmt.Printf("已成年，age=%v\n", age)
	} else {
		fmt.Printf("未成年，age=%v\n", age)
	}
	// fmt.Printf("%v\n",age) 报错
}
*/
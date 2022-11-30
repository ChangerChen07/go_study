/* package main

import "fmt"

func main() {
	// &取值，*指针访问值
	// var ip *int
	// fmt.Printf("ip: %v\n", ip)
	// fmt.Printf("ip: %T\n", ip)
	i := 100
	ip := &i //取i的地址
	fmt.Printf("i: %v\n", i)
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("*ip: %v\n", *ip)

	ip2 := &ip
	fmt.Printf("ip2: %v\n", ip2)
	fmt.Printf("*ip2: %v\n", *ip2)

	var a [3]int = [3]int{1, 2, 3}
	var ap [3]*int
	for i := 0; i < len(a); i++ {
		ap[i] = &a[i]
	}
	fmt.Printf("ap: %v\n", ap)
	fmt.Printf("a: %v\n", a)

}
*/
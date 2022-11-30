/* // 执行顺序 变量初始化>init>main
package main

import "fmt"

func init() {
	// 自动执行，不能被调用
	fmt.Println("init.......")
	j = 10
}

var j int
var s int = initVar()

func initVar() int {
	fmt.Println("initVar.....")
	return 100
}
func main() {
	fmt.Printf("s: %v\n", s)
	fmt.Printf("j: %v\n", j)
	fmt.Println("main.......")
}
*/
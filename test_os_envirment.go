/* package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取环境变量
	s := os.Environ()
	fmt.Printf("s: %v\n", s)

	//获得某个环境变量
	s2 := os.Getenv("GOPATH")
	fmt.Printf("s2: %v\n", s2)

	// 设置环境变量
	os.Setenv("ENV1", "TEST")
	os.Setenv("ENV2", "env2")
	fmt.Printf("os.Getenv(\"ENV1\"): %v\n", os.Getenv("ENV1"))

	s3, b := os.LookupEnv("aaaa")
	if b {
		fmt.Printf("s3: %v\n", s3)
	} else {
		fmt.Println("aaaa不存在")
	}

	// 替换
	os.Setenv("ENV1", "test2")
	os.Setenv("ENV2", "test23")
	fmt.Println(os.ExpandEnv("$ENV1 and ${ENV2} is replaced")) //获取环境变量值通过$var 获取

	// os.Clearenv() 清空所有环境变量
}
*/
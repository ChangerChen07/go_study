/* package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile() {
	//若文件存在，会覆盖
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func createDir() {
	// err := os.Mkdir("test", os.ModePerm)
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }
	err2 := os.MkdirAll("test/1/2", os.ModePerm) //若目录存在，不做任何操作，返回nil

	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 删除目录或者文件
func delFile() {
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.RemoveAll("test")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 获取目录信息
func getPWD() {
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("e:/")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}
func main() {
	getPWD()
}
*/
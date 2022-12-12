/* package main

import (
	"fmt"
	"os"
)

func openCloseFIle() {
	f1, _ := os.Open("test.txt")
	fmt.Printf("f1.Name(): %v\n", f1.Name())

	f2, _ := os.OpenFile("test1.txt", os.O_RDWR|os.O_CREATE, 0755)
	fmt.Printf("f2.Name(): %v\n", f2.Name())

	err1 := f1.Close()
	fmt.Printf("err1: %v\n", err1)
	err2 := f2.Close()
	fmt.Printf("err2: %v\n", err2)
}

func readOps() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer f.Close()
	var buf = make([]byte, 4)
	// n, _ := f.ReadAt(buf, 5) 指定从哪里开始读取
	f.Seek(-5, 2) //第一个参数是相对位置，第二个参数是相对哪里计算：0开头，1是相对当前位置，2是文件结尾
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))

	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	for {
	// 		buf := make([]byte, 4)
	// 		n, ferr := f.Read(buf)
	// 		if ferr == io.EOF {
	// 			fmt.Printf("ferr: %v\n", ferr)
	// 			break
	// 		} else {
	// 			fmt.Printf("n: %v\n", n)
	// 			fmt.Printf("string(buf): %v\n", string(buf))
	// 		}
	// 	}

	// }
	// f.Close()

}

func readDirs() {
	f, _ := os.Open("a")
	dr, err := f.ReadDir(-1) //或者直接 dr,err:=os.ReadDir("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, v := range dr {
		fmt.Printf("v.Name(): %v\n", v.Name())
		if v.IsDir() {
			fmt.Printf(" %v是文件夹\n", v.Name())
		} else {
			fmt.Printf(" %v是文件\n", v.Name())
		}
	}
}
func main() {
	// openCloseFIle()
	readOps()
	// readDirs()
}
*/
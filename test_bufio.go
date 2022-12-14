package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	// 读出来放到缓存里操作
	// r := strings.NewReader("test buffer reader")
	r, _ := os.Open("test.txt")
	defer r.Close()
	r2 := bufio.NewReader(r)
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)

}

func test2() {
	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=")
	br := bufio.NewReader(r) //缓冲区
	p := make([]byte, 10)    //创建内存空间为10的字节sliceb

	for {
		n, err := br.Read(p) //从缓存区读到内存空间中
		if err == io.EOF {
			break
		}
		fmt.Printf("string(p): %v\n", string(p[:n]))
	}
}

func test3() {
	r := strings.NewReader("ABCDEFG")
	r2 := bufio.NewReader(r)

	b, _ := r2.ReadByte()
	fmt.Printf("b: %c\n", b)

	b, _ = r2.ReadByte()
	fmt.Printf("b: %c\n", b)

	r2.UnreadByte()
	b, _ = r2.ReadByte()
	fmt.Printf("b: %c\n", b)
}

func test4() {
	// 针对不是英文的读取单字节
	r := strings.NewReader("高龄信息科技股份有限公司！")
	r2 := bufio.NewReader(r)

	r3, size, _ := r2.ReadRune()
	fmt.Printf("r3: %c size:%d\n", r3, size)

	r3, size, _ = r2.ReadRune()
	fmt.Printf("r3: %c size:%d\n", r3, size)

	r2.UnreadRune()
	r3, size, _ = r2.ReadRune()
	fmt.Printf("r3: %c size:%d\n", r3, size)
}

func test5() {
	r := strings.NewReader("ABC,DEF,GHI,JKL")
	r2 := bufio.NewReader(r)

	// line, _ := r2.ReadSlice(',')
	// line, _ := r2.ReadBytes(',')
	line, _ := r2.ReadString(',')
	fmt.Printf("line: %q\n", line)

	line, _ = r2.ReadString(',')
	fmt.Printf("line: %q\n", line)

	line, _ = r2.ReadString(',')
	fmt.Printf("line: %q\n", line)
}

func test6() {
	r := strings.NewReader("ABCDEFGHIJKLMN")
	r2 := bufio.NewReader(r)
	// b := bytes.NewBuffer(make([]byte, 0))

	f, _ := os.OpenFile("test3.txt", os.O_RDWR|os.O_CREATE, 755)
	defer f.Close()
	r2.WriteTo(f)
	// r2.WriteTo(b)
	// fmt.Printf("b: %s\n", b)

}
func main() {
	// test1()
	test6()
}

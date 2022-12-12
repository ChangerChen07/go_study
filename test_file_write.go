/* package main

import (
	"fmt"
	"os"
)

func write() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 766)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Write([]byte("Hello World!"))
	f.Close()
}

func writeString() {
	f, err := os.OpenFile("a.txt", os.O_APPEND, 766)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteString("Hello World!!")
	f.Close()
}

func writeAt() {
	f, err := os.OpenFile("a.txt", os.O_WRONLY, 766)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	n, wrr := f.WriteAt([]byte("CCCC"), 2)
	fmt.Printf("n: %v\n", n)
	if wrr != nil {
		fmt.Printf("wrr: %v\n", wrr)
	}
	f.Close()
}

func main() {
	writeAt()
}
*/
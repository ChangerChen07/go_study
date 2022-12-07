package main

import (
	"fmt"
	"runtime"
)

func showMsg(s int) {
	fmt.Printf("s: %v\n", s)
	// time.Sleep(time.Millisecond * 100)
}
func main() {
	for i := 0; i < 5; i++ {

		go showMsg(i)
		runtime.Gosched()
		fmt.Printf("i: %v\n", i)
	}
	// for i := 0; i < 5; i++ {
	// 	runtime.Gosched()
	// 	fmt.Printf("i: %v\n", i)
	// }
	// runtime.Gosched()
	fmt.Println("main end.......")
}

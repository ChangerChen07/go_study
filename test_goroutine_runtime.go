/* package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wp sync.WaitGroup

func showMsg(s int) {
	fmt.Printf("s: %v\n", s)
	// time.Sleep(time.Millisecond * 100)
}
func show(i int) {
	defer wp.Done()
	if i >= 5 {
		runtime.Goexit()

	}
	fmt.Printf("show i: %v\n", i)
}

func showA() {

	for i := 0; i < 10; i++ {
		fmt.Printf("A i: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}

}
func showB() {

	for i := 0; i < 10; i++ {
		fmt.Printf("B i: %v\n", i)
		time.Sleep(time.Millisecond * 100)
	}

}
func main() {
	fmt.Printf("cpus:%v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go showA()

	go showB()

	time.Sleep(time.Second)
	fmt.Println("main end........")

	// for i := 0; i < 10; i++ {
	// 	go show(i)
	// 	wp.Add(1)
	// }
	// wp.Wait()
	// fmt.Println("main end..........")

	// for i := 0; i < 5; i++ {

	// 	go showMsg(i)
	// 	runtime.Gosched()
	// 	fmt.Printf("i: %v\n", i)
	// }
	// // for i := 0; i < 5; i++ {
	// // 	runtime.Gosched()
	// // 	fmt.Printf("i: %v\n", i)
	// // }
	// // runtime.Gosched()
	// fmt.Println("main end.......")
}
*/
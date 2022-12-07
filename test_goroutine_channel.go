/* package main

import (
	"fmt"
	"time"
)

var values = make(chan int, 5)

func send(value int) {
	// rand.Seed(time.Now().UnixNano())
	// value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	values <- value
}
func showMsg(value int) {
	// value := <-values
	fmt.Printf("get: %v\n", value)
}

func main() {
	// unbuffer := make(chan int)
	// buffed:=make(chan int,5)
	defer close(values)
	for i := 0; i < 10; i++ {
		go send(i)
		value := <-values
		go showMsg(value)
	}

	time.Sleep(time.Millisecond * 10000)
	fmt.Println("main end")
}
*/
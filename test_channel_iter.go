/* package main

import "fmt"

func main() {
	var c = make(chan int)
	// var r int
	go func() {
		defer close(c)
		for i := 0; i < 20; i++ {
			c <- i
		}
		// close(c) //不关闭，死锁(deadlock)
	}()
	// 死循环读取
	// for {
	// 	if r, ok := <-c; ok {
	// 		fmt.Printf("r: %v\n", r)
	// 	} else {
	// 		break
	// 	}
	// }

	for v := range c {
		fmt.Printf("v: %v\n", v)
	}
}
*/
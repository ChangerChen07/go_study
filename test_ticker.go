/* package main

import (
	"fmt"
	"time"
)

func test1() {
	var count int
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("ticker 2s......")
		count++
		fmt.Printf("count: %v\n", count)
		if count >= 3 {
			ticker.Stop()
			fmt.Println("ticker stopped....")
		}
	}
}

func main() {
	// test1() 函数跟下面这段不能放在一起执行，Why?
	var chanInt = make(chan int)
	ticker2 := time.NewTicker(time.Second * 1)
	go func() {
		for _ = range ticker2.C {
			select {
			case chanInt <- 1:
				fmt.Println("send 1")
			case chanInt <- 2:
				fmt.Println("send 2")
			case chanInt <- 3:
				fmt.Println("send 3")
			}
		}
	}()

	sum := 0
	for v := range chanInt {
		fmt.Printf("rec: %v\n", v)
		sum += v
		fmt.Printf("sum: %v\n", sum)
		if sum >= 30 {
			ticker2.Stop()
			break
		}
	}
}
*/
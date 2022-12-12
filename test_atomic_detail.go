/* package main

import (
	"fmt"
	"sync/atomic"
)

func test_add_sub() {
	// 自动上锁，防止多线程导致加减混乱
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -2)
	fmt.Printf("i: %v\n", i)
}

func test_load_store() {
	var i int32 = 100
	atomic.LoadInt32(&i) //read
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i, 200) //write
	fmt.Printf("i: %v\n", i)
}

func test_cas() {
	var i int32 = 200
	b := atomic.CompareAndSwapInt32(&i, 200, 100)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("i: %v\n", i)
}
func main() {
	test_cas()
}
*/
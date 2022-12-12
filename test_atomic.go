/* package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int32 = 100
var m = 100
var wp sync.WaitGroup

func add() {
	defer wp.Done()
	atomic.AddInt32(&i, 1)
}

func sub() {
	defer wp.Done()
	atomic.AddInt32(&i, -1)
}

func add1() {
	m++
}

func sub1() {
	m--
}
func main() {
	for a := 0; a <= 100; a++ {
		// go add()
		// wp.Add(1)
		// go sub()
		// wp.Add(1)
		go add1()
		go sub1()
	}

	// wp.Wait()
	// fmt.Printf("i: %v\n", i)
	// time.Sleep(time.Second * 1)
	fmt.Printf("m: %v\n", m)
}
*/
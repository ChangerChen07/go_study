/* package main

import (
	"fmt"
	"sync"
)

var m int = 100
var wp sync.WaitGroup
var lock sync.Mutex //添加进程锁,结合waitgroup，m支部混乱;若不添加，最后m值混乱

func add() {
	defer wp.Done()
	lock.Lock()
	m += 1
	fmt.Printf("m++: %v\n", m)
	lock.Unlock()
}
func sub() {
	defer wp.Done()
	lock.Lock()
	m -= 1
	fmt.Printf("m--: %v\n", m)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		go add()
		wp.Add(1)
		go sub()
		wp.Add(1)
	}
	wp.Wait()
	fmt.Printf("end m: %v\n", m)
}
*/
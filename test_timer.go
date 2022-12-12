/* package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)

	now := time.Now()
	fmt.Printf("now: %v\n", now)
	// t1 := <-timer.C //阻塞的，直到指定时间到
	// fmt.Printf("t1: %T\n", t1)
	// <-timer.C
	// fmt.Printf("time.Now(): %v\n", time.Now())
	// fmt.Println("2s后")

	// time.Sleep(time.Second * 2)
	// fmt.Println("2s后")

	// <-time.After(time.Second * 2) //After函数返回值是chan time,timer.c?
	// fmt.Println("2s后")
	go func() {
		<-timer.C
		fmt.Println("timer expired")
	}()
	stop := timer.Stop()
	if stop {
		fmt.Println("timer is stopped.")
	}
	<-time.After(time.Second * 4)
	fmt.Println("main end...")

	fmt.Println("Before")
	timer2 := time.NewTimer(time.Second * 5)
	timer2.Reset(time.Second * 2) //重置timer时间
	<-timer2.C
	fmt.Println("After")
}
*/
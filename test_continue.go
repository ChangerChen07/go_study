/* package main

import "fmt"

func f1() {
	//跳到标签 MYLABEL
MYLABEL:
	for i := 0; i < 6; i++ {
		// MYLABEL:
		for j := 0; j < 6; j++ {
			if i == 2 && j == 2 {
				continue MYLABEL
				// continue
			}
			fmt.Printf("%v, %v\n", i, j)
		}
	}
}
func main() {
	f1()
}
*/
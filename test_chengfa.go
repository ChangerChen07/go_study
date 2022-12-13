package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 9; i++ {
		// var s string
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
			// if j > i {
			// 	break
			// }
			// s += strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(i*j) + " "
		}
		// s = s + "\n"
		fmt.Println()
	}
}

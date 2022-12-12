package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 10; i++ {
		var s string
		for j := 1; j < 10; j++ {
			if j > i {
				break
			}
			s += strconv.Itoa(j) + "*" + strconv.Itoa(i) + "=" + strconv.Itoa(i*j) + " "
		}
		s = s + "\n"
		fmt.Println(s)
	}
}

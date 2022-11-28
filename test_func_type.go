/* package main

import "fmt"

type ft func(int, int) int

func sum(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	// var ff ft
	var fm ft
	// ff = sum
	ff := sum
	fmt.Printf("ff: %T\n", ff)
	r := ff(1, 2)
	fmt.Printf("r: %v\n", r)

	fm = max
	fmt.Printf("fm: %T\n", fm)
	r = fm(10, 20)
	fmt.Printf("r: %v\n", r)
}
*/
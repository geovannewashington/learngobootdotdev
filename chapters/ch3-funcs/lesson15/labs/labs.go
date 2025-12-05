package main

import (
	"fmt"
)

func sub(x, y int) int {
	return x - y
}

func square(x int) int {
	return x * x
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	first_result := arithmetic(a, b)
	second_result := arithmetic(first_result, c)
	return second_result
}
func main() {
	x := aggregate(2, 3, 3, sub)
	// y := aggregate(2, 3, 3, square)
	fmt.Printf("%d\n", x)
}

package main

import (
	"fmt"
)

func main() {
	a, b := 2, 3
	swap(&a, &b)
	fmt.Printf("%d, %d\n", a, b)
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

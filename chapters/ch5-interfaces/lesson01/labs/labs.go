package main

import (
	"fmt"
)

type MathUtils interface {
	// - some math operations
	add(a, b int) int 
	square(x int) int 
}

type SomeValues struct {
	a int
	b int
	x int
} 

// -- some methods for 'SomeValues'
func (s SomeValues) add(a, b int) int {
	return a + b;
}

func (s SomeValues) square(x int) int {
	return x * x;
}

func printOperation(values SomeValues) {
	fmt.Printf("Add: %v\n", values.add(values.a, values.b)) // 13
	fmt.Printf("Square: %v\n", values.square(values.x))     // 16
}

func main() {
	values := SomeValues{
		a : 5,
		b : 8,
		x : 4,
	}	
	
	fmt.Println("Testing...");	
	printOperation(values)
}

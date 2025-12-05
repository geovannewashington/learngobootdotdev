package main

import (
	"fmt"
)

func proccessOperation(operationFunc func(int) int, value int) int {
	return operationFunc(value)
}

func main() {
	// -- using an anonymous function example		
	result := proccessOperation(func (x int) int {
		return x * x
	}, 5)	
	fmt.Printf("%d\n", result)
}

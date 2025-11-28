package main

import (
	"fmt"
	"math/rand"
)

func get_length() float64 {
	return rand.Float64() * 1
}

func main() {
	// The initial statement on if conditionals
	if length := get_length(); length < 0.5 {
		fmt.Println("Invalid email");
	}
}



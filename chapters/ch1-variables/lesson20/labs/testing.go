package main

import (
	"fmt"
)

func main() {
	count, err := fmt.Printf("urmom\n");

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Successfully printed characters: %v\n", count) // %v is like a generic format specifier
}

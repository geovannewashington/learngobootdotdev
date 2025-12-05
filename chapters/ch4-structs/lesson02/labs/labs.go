package main

import (
	"fmt"
)

// struct
type Car struct {
	brand string
}

func main() { 
	key := "brand"
	myCar := Car{ brand: "BMW" }

	fmt.Printf("%s\n", myCar.["brand"])
}

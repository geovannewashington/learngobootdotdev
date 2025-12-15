package main 

import (
	"fmt"
)

func main() {
	x := 10.523
	s := fmt.Sprintf("I am %d years old", int(x))
	fmt.Println(s)
}

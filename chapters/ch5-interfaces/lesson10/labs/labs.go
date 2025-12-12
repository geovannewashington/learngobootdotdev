package main

import (
	"fmt"
)

type shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius;
}

func processShape(s shape) {
	isCircle := false
	foo, ok := s.(Circle)	
	
	if ok {
		isCircle = true
	}
	
	if isCircle {
		fmt.Printf("Circle's area: %v\n", foo.radius)
	}
}

func main() {	
	s := Circle{ radius: 5 } 
	processShape(s)	
}

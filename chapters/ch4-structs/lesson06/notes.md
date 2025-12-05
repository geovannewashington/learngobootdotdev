# Struct Methods in Go

Go support methods that can be defined on structs.
Methods are just functions that have a receiver.
A receiver is a special parameter that syntactically goes before the name of the function.

Example:

```go
type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
// rect is the struct
// r is the placeholder
func (r rect) area() int {
  return r.width * r.height
}

var r = rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50
```

The generic syntax for declaring a method in golang:

```plaintext
func (receiverName ReceiverType) MethodName(parameters) (returns) {
    // method body
}
```

receiverName: An identifier used to refer to the instance of the type within the method's body.
Similar to this or self in other languages.

Another exaple:

```go
package main

import "fmt"

type Circle struct {
	Radius float64
}

// Value receiver method
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Pointer receiver method
func (c *Circle) SetRadius(newRadius float64) {
	c.Radius = newRadius
}

func main() {
	myCircle := Circle{Radius: 5}
	fmt.Println("Initial Area:", myCircle.Area()) // Output: Initial Area: 78.5

	myCircle.SetRadius(10)
	fmt.Println("New Area:", myCircle.Area()) // Output: New Area: 314
}
```

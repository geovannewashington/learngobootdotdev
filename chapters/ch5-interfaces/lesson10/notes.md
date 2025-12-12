# Type Assertions in Go

## What is type assertion:

A programming mechanism where a developer explicitly tells the compiler to treat a value as a
specific type, even if the compiler cannot automatically infer it. It is essentially a way of saying,
"Trust me, I know better about the type of this value than you do

In Go, a
type assertion is used to extract an interface variable's underlying concrete value or to check if
ft implements a specific interface. The syntax is i.(T), where i is an interface value and T is the type being asserted

Exaple:

```Go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    // Check if i holds a string
    if s, ok := i.(string); ok {
        fmt.Println("String value:", s)
    } else {
        fmt.Println("Not a string")
    }
    // Output: String value: hello

    // Check for a different type (fails safely)
    if f, ok := i.(float64); ok {
        fmt.Println("Float value:", f)
    } else {
        fmt.Println("Type assertion failed: not a float64")
    }
    // Output: Type assertion failed: not a float64
}
```

Note: we can use the `%T` format specifier to find out the type of a variable.
Note that type assertions in Go only works on interfaces.

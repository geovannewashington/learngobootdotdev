# Type Switches

A `type switch` makes it easy to do several type assertions in a series.

Example:

```Go
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}
```

Another Example:

```Go
package main

import "fmt"

func printType(i interface{}) {
    // The 'v' variable inside each case block will have the specific type declared in that case.
    switch v := i.(type) {
    case int:
        fmt.Printf("It's an integer, double its value: %d\n", v * 2)
    case string:
        fmt.Printf("It's a string, its length is: %d\n", len(v))
    case bool:
        fmt.Printf("It's a boolean with value: %t\n", v)
    case nil:
        fmt.Println("It's nil.")
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    printType(42)
    printType("hello")
    printType(true)
    printType(nil)
    printType(1.23) // Triggers the default case
}
```

Note that an empty interface can used to receive any value, it servers as the `any` value in other languages.

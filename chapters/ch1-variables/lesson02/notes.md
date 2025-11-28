# Notes on Lesson 2

The lesson introduced the most verbose way of declaring variables in Go.
In practice, explicit type annotations are rarely needed because Go performs type inference automatically
using short variable declarations (:=).

## `var` vs `:=`

Using the var keyword, you can declare a variable without an initial value.
In that case, the variable is assigned its zero value, which depends on the type (e.g., 0 for numeric
types, "" for strings, false for booleans, nil for pointers and slices.

Example:

```go
var n int    // zero value: 0
var s string // zero value: ""
var ok bool  // zero value: false
```

Important rule:

- `var` can be used anywhere (including at the package level).
- `:=` is allowed only inside function bodies.

## Format Specifiers (fmt)

- `%v` prints a value in its default format.
- `%q` prints a string as a double-quoted, escaped string literal.

Example:

```go

package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "hello\nworld"
	s3 := `She said, "Hello!"`

	fmt.Printf("%s\n", s1) // hello
	fmt.Printf("%q\n", s1) // "hello"

	fmt.Printf("%s\n", s2) // hello
	                       // world
	fmt.Printf("%q\n", s2) // "hello\nworld"

	fmt.Printf("%s\n", s3) // She said, "Hello!"
	fmt.Printf("%q\n", s3) // "She said, \"Hello!\""
}
```

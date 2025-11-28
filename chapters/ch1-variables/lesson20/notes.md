## Formatting Strings in Go

`fmt.Printf()` behaves almost like in C.
it return two values: err and and integer respectivelly.

this integer, just like in C, is the number of successfully printed characters...

`fmt.Sprintf()` on the other hand, returns the formated string...

Example:

```Go
name := "tpose"
s := fmt.Sprintf("hello, %v\n", name)
// s is: hello, tpose
```

## Named Return Values

When we `return` from a function explicit values like:
`return x` or `return y`, this is called a named retur value.

There is another concept called `naked return` which is a return statement without any arguments
Naked returns should be used only in short functions. They can harm readibility in longer functions.

These two code snippters are equivalent:

```go
func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // naked return: automatically returns x and y
}
```

```go
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}
```

## Named Return Values

Named return is when we name the return on the function signature, just like this:
`func calculator(a, b int) (mut, div int, err error)`

This is best practice when the function return many different things.
It's like a way of seld-documenting the code

On the other hand we should just use `naked return` when the function is small.

These two code snippters are equivalent:

```go
func getCoords() (x, y int) {
	// x and y are initialized with zero values

	return // implicity return: automatically returns x and y
}
```

```go
func getCoords() (int, int) {
	var x int
	var y int
	return x, y
}
```
